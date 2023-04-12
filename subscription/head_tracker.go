package subscription

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"sync"
	"time"

	"github.com/begmaroman/eth-services/broadcaster"

	"github.com/begmaroman/eth-services/client"
	"github.com/begmaroman/eth-services/store"
	"github.com/begmaroman/eth-services/store/models"
	"github.com/begmaroman/eth-services/types"

	ethereum "github.com/ethereum/go-ethereum"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

//go:generate mockery --name HeadTrackable --output ../internal/mocks/ --case=underscore

// HeadTrackable represents any object that wishes to respond to ethereum events,
// after being attached to HeadTracker.
type HeadTrackable interface {
	Connect(head *models.Head) error
	Disconnect()
	OnNewLongestChain(ctx context.Context, head *models.Head)
}

// headRingBuffer is a small goroutine that sits between the eth client and the
// head tracker and drops the oldest head if necessary in order to keep to a fixed
// queue size (defined by the buffer size of out channel)
type headRingBuffer struct {
	in     <-chan *etypes.Header
	out    chan models.Head
	start  sync.Once
	logger types.Logger
}

func newHeadRingBuffer(in <-chan *etypes.Header, size int, logger types.Logger) (r *headRingBuffer, out chan models.Head) {
	out = make(chan models.Head, size)
	return &headRingBuffer{
		in:     in,
		out:    out,
		start:  sync.Once{},
		logger: logger,
	}, out
}

// Start the headRingBuffer goroutine
// It will be stopped implicitly by closing the in channel
func (r *headRingBuffer) Start() {
	r.start.Do(func() {
		go r.run()
	})
}

func (r *headRingBuffer) run() {
	for h := range r.in {
		if h == nil {
			r.logger.Error("HeadTracker: got nil block header")
			continue
		}

		hInQueue := len(r.out)
		if hInQueue > 0 {
			r.logger.Infof("HeadTracker: Head %v is lagging behind, there are %v more heads in the queue.", h.Number, hInQueue)
		}

		model := models.FromHeader(h)

		select {
		case r.out <- *model:
		default:
			// Need to select/default here because it's conceivable (although
			// improbable) that between the previous select and now, all heads were drained
			// from r.out by another goroutine
			//
			// NOTE: In this unlikely event, we may drop an extra head unnecessarily.
			// The probability of this seems vanishingly small, and only hits
			// if the queue was already full anyway, so we can live with this
			select {
			case dropped := <-r.out:
				r.logger.Errorf("HeadTracker: dropping head %v with hash 0x%x because queue is full.", dropped.Number, h.Hash)
				r.out <- *model
			default:
				r.out <- *model
			}
		}
	}

	close(r.out)
}

// HeadTracker holds and stores the latest block number seen in a thread safe manner.
// Reconstitutes the last block number from the data store on reboot.
type HeadTracker struct {
	callbacks             []HeadTrackable
	inHeaders             chan *etypes.Header
	outHeaders            chan models.Head
	cancelHeadSum         func()
	highestSeenHead       *models.Head
	broadcaster           broadcaster.Broadcaster
	ethClient             client.GethClient
	store                 store.Store
	headMutex             sync.RWMutex
	connected             bool
	sleeper               Sleeper
	done                  chan struct{}
	started               bool
	listenForNewHeadsWg   sync.WaitGroup
	subscriptionSucceeded chan struct{}
	config                *types.Config
	logger                types.Logger
}

// NewHeadTracker instantiates a new HeadTracker using the db to persist new block numbers.
// Can be passed in an optional sleeper object that will dictate how often
// it tries to reconnect.
func NewHeadTracker(
	broadcaster broadcaster.Broadcaster,
	ethClient client.GethClient,
	store store.Store,
	config *types.Config,
	callbacks []HeadTrackable,
	sleepers ...Sleeper,
) *HeadTracker {
	var sleeper Sleeper
	if len(sleepers) > 0 {
		sleeper = sleepers[0]
	} else {
		sleeper = NewBackoffSleeper()
	}
	return &HeadTracker{
		broadcaster: broadcaster,
		ethClient:   ethClient,
		store:       store,
		config:      config,
		logger:      config.Logger,
		callbacks:   callbacks,
		sleeper:     sleeper,
	}
}

// Start retrieves the last persisted block number from the HeadTracker,
// subscribes to new heads, and if successful fires Connect on the
// HeadTrackable argument.
func (ht *HeadTracker) Start() error {
	ht.headMutex.Lock()
	defer ht.headMutex.Unlock()

	if ht.started {
		return nil
	}

	if err := ht.setHighestSeenHeadFromDB(); err != nil {
		return err
	}
	if ht.highestSeenHead != nil {
		ht.logger.Debug("Tracking logs from last block ", ht.highestSeenHead.ToInt(), " with hash ", ht.highestSeenHead.Hash.Hex())
	}

	ht.done = make(chan struct{})
	ht.subscriptionSucceeded = make(chan struct{})

	ht.listenForNewHeadsWg.Add(1)
	go ht.listenForNewHeads()

	ht.started = true
	return nil
}

// Stop unsubscribes all connections and fires Disconnect.
func (ht *HeadTracker) Stop() error {
	ht.headMutex.Lock()

	if !ht.started {
		ht.headMutex.Unlock()
		return nil
	}

	if ht.connected {
		ht.connected = false
		ht.disconnect()
	}
	ht.logger.Infof(fmt.Sprintf("Head tracker disconnecting from %v", ht.config.RPCURL))
	close(ht.done)
	close(ht.subscriptionSucceeded)
	ht.started = false
	ht.headMutex.Unlock()

	ht.listenForNewHeadsWg.Wait()
	return nil
}

// Save updates the latest block number, if indeed the latest, and persists
// this number in case of reboot. Thread safe.
func (ht *HeadTracker) Save(h *models.Head) error {
	ht.headMutex.Lock()
	if h.GreaterThan(ht.highestSeenHead) {
		ht.highestSeenHead = h
	}
	ht.headMutex.Unlock()

	err := ht.store.InsertHead(h)
	if err != nil {
		return err
	}
	return ht.store.TrimOldHeads(ht.config.HeadTrackerHistoryDepth)
}

// HighestSeenHead returns the block header with the highest number that has been seen, or nil
func (ht *HeadTracker) HighestSeenHead() *models.Head {
	ht.headMutex.RLock()
	defer ht.headMutex.RUnlock()

	if ht.highestSeenHead == nil {
		return nil
	}
	h := *ht.highestSeenHead
	return &h
}

// Connected returns whether or not this HeadTracker is connected.
func (ht *HeadTracker) Connected() bool {
	ht.headMutex.RLock()
	defer ht.headMutex.RUnlock()

	return ht.connected
}

// ExportedDone exports the done channel for testing
func (ht *HeadTracker) ExportedDone() chan struct{} {
	return ht.done
}

func (ht *HeadTracker) connect(bn *models.Head) {
	for _, trackable := range ht.callbacks {
		err := trackable.Connect(bn)
		if err != nil {
			ht.logger.Warn("Error connecting", err)
		}
	}
}

func (ht *HeadTracker) disconnect() {
	for _, trackable := range ht.callbacks {
		trackable.Disconnect()
	}
}

func (ht *HeadTracker) listenForNewHeads() {
	defer ht.listenForNewHeadsWg.Done()
	defer func() {
		err := ht.unsubscribeFromHead()
		if err != nil {
			ht.logger.Errorf("Failed when unsubscribe from head %w", err)
		}
	}()

	for {
		if !ht.subscribe() {
			return
		}
		if err := ht.receiveHeaders(); err != nil {
			ht.logger.Errorw(fmt.Sprintf("Error in new head subscription, unsubscribed: %s", err.Error()), "err", err)
			continue
		} else {
			return
		}
	}
}

// subscribe periodically attempts to connect to the ethereum node via websocket.
// It returns true on success, and false if cut short by a done request and did not connect.
func (ht *HeadTracker) subscribe() bool {
	ht.sleeper.Reset()
	for {
		err := ht.unsubscribeFromHead()
		if err != nil {
			ht.logger.Errorf("Failed when unsubscribe from head %w", err)
			return false
		}

		ht.logger.Info("Connecting to ethereum node ", ht.config.RPCURL, " in ", ht.sleeper.Duration())

		select {
		case <-ht.done:
			return false
		case <-time.After(ht.sleeper.After()):
			if err = ht.subscribeToHead(); err != nil {
				ht.logger.Warnw(fmt.Sprintf("Failed to connect to ethereum node %v", ht.config.RPCURL), "err", err)
			} else {
				ht.logger.Info("Connected to ethereum node ", ht.config.RPCURL)
				return true
			}
		}
	}
}

// This should be safe to run concurrently across multiple nodes connected to the same database
func (ht *HeadTracker) receiveHeaders() error {
	for {
		select {
		case <-ht.done:
			return nil
		case blockHeader, open := <-ht.outHeaders:
			if !open {
				return errors.New("HeadTracker: outHeaders prematurely closed")
			}
			ctx, cancel := context.WithTimeout(context.Background(), ht.totalNewHeadTimeBudget())
			if err := ht.handleNewHead(ctx, &blockHeader); err != nil {
				cancel()
				return err
			}
			cancel()
		}
	}
}

func (ht *HeadTracker) handleNewHead(ctx context.Context, head *models.Head) error {
	defer func(start time.Time, number int64) {
		elapsed := time.Since(start)
		if elapsed > ht.callbackExecutionThreshold() {
			ht.logger.Warnw(fmt.Sprintf("HeadTracker finished processing head %v in %s which exceeds callback execution threshold of %s", number, elapsed.String(), ht.callbackExecutionThreshold().String()), "blockNumber", number, "time", elapsed, "id", "head_tracker")
		} else {
			ht.logger.Debugw(fmt.Sprintf("HeadTracker finished processing head %v in %s", number, elapsed.String()), "blockNumber", number, "time", elapsed, "id", "head_tracker")
		}
	}(time.Now(), head.Number)
	prevHead := ht.HighestSeenHead()

	ht.logger.Debugw("Received new head", "blockHeight", head.ToInt(), "blockHash", head.Hash)

	if err := ht.Save(head); err != nil {
		return err
	}

	if prevHead == nil || head.Number > prevHead.Number {
		return ht.handleNewHighestHead(head)
	}
	if head.Number == prevHead.Number {
		if head.Hash != prevHead.Hash {
			ht.logger.Debugw("HeadTracker: got duplicate head", "blockNum", head.Number, "gotHead", head.Hash.Hex(), "highestSeenHead", ht.highestSeenHead.Hash.Hex())
		} else {
			ht.logger.Debugw("HeadTracker: head already in the database", "gotHead", head.Hash.Hex())
		}
	} else {
		ht.logger.Debugw("HeadTracker: got out of order head", "blockNum", head.Number, "gotHead", head.Hash.Hex(), "highestSeenHead", ht.highestSeenHead.Number)
	}
	return nil
}

func (ht *HeadTracker) handleNewHighestHead(head *models.Head) error {
	// NOTE: We must set a hard time limit on this, backfilling heads should
	// not block the head tracker
	ctx, cancel := context.WithTimeout(context.Background(), ht.backfillTimeBudget())
	defer cancel()

	headWithChain, err := ht.GetChainWithBackfill(ctx, head, ht.config.FinalityDepth)
	if err != nil {
		return err
	}

	ht.onNewLongestChain(ctx, headWithChain)
	return nil
}

// totalNewHeadTimeBudget is the timeout on the shared context for all
// requests triggered by a new head
//
// These values are chosen to be roughly 2 * block time (to give some leeway
// for temporary overload). They are by no means set in stone and may require
// adjustment based on real world feedback.
func (ht *HeadTracker) totalNewHeadTimeBudget() time.Duration {
	return 2 * ht.config.BlockTime
}

// Maximum time we are allowed to spend backfilling heads. This should be
// somewhat shorter than the average time between heads to ensure we
// don't starve the runqueue.
func (ht *HeadTracker) backfillTimeBudget() time.Duration {
	return time.Duration(7 * float64(ht.config.BlockTime) / 10)
}

// If total callback execution time exceeds this threshold we consider this to
// be a problem and will log a warning.
// Here we set it to the average time between blocks.
func (ht *HeadTracker) callbackExecutionThreshold() time.Duration {
	return ht.config.BlockTime
}

// GetChainWithBackfill returns a chain of the given length, backfilling any
// heads that may be missing from the database
func (ht *HeadTracker) GetChainWithBackfill(ctx context.Context, head *models.Head, depth int64) (*models.Head, error) {
	ctx, cancel := context.WithTimeout(ctx, ht.backfillTimeBudget())
	defer cancel()

	head, err := ht.store.Chain(head.Hash, depth)
	if err != nil {
		return head, errors.Wrap(err, "GetChainWithBackfill failed fetching chain")
	}
	if head.ChainLength() >= depth {
		return head, nil
	}
	baseHeight := int64(head.Number) - (int64(depth) - 1)
	if baseHeight < 0 {
		baseHeight = 0
	}

	if err := ht.backfill(ctx, head.EarliestInChain(), baseHeight); err != nil {
		return head, errors.Wrap(err, "GetChainWithBackfill failed backfilling")
	}
	return ht.store.Chain(head.Hash, depth)
}

// backfill fetches all missing heads up until the base height
func (ht *HeadTracker) backfill(ctx context.Context, head *models.Head, baseHeight int64) error {
	if head.Number <= baseHeight {
		return nil
	}
	mark := time.Now()
	fetched := 0
	defer func() {
		var headNumber int64 = 0
		if head != nil {
			headNumber = head.Number
		}
		ht.logger.Debugw("HeadTracker: finished backfill",
			"fetched", fetched,
			"blockNumber", headNumber,
			"time", time.Since(mark),
			"id", "head_tracker",
			"n", headNumber-baseHeight,
			"fromBlockHeight", baseHeight,
			"toBlockHeight", headNumber-1)
	}()

	for i := int64(head.Number - 1); i >= int64(baseHeight); i-- {
		// NOTE: Sequential requests here mean it's a potential performance bottleneck, be aware!
		existingHead, err := ht.store.HeadByHash(head.ParentHash)
		if err != nil {
			return errors.Wrap(err, "HeadByHash failed")
		}
		if existingHead != nil {
			head = existingHead
			continue
		}
		head, err = ht.fetchAndSaveHead(ctx, uint64(i))
		fetched++
		if err != nil {
			if errors.Cause(err) == ethereum.NotFound {
				ht.logger.Errorw("HeadTracker: backfill failed to fetch head (not found), chain will be truncated for this head", "headNum", i)
			} else if errors.Cause(err) == context.DeadlineExceeded {
				ht.logger.Infow("HeadTracker: backfill deadline exceeded, chain will be truncated for this head", "headNum", i)
			} else {
				ht.logger.Errorw("HeadTracker: backfill encountered unknown error, chain will be truncated for this head", "headNum", i, "err", err)
			}
			break
		}
	}
	return nil
}

func (ht *HeadTracker) fetchAndSaveHead(ctx context.Context, n uint64) (*models.Head, error) {
	ht.logger.Debugw("HeadTracker: fetching head", "blockHeight", n)

	head, err := ht.ethClient.HeaderByNumber(ctx, new(big.Int).SetUint64(n))
	if err != nil {
		return nil, err
	} else if head == nil {
		return nil, errors.New("got nil head")
	}

	model := models.FromHeader(head)

	// Store header
	if err = ht.store.InsertHead(model); err != nil {
		return nil, err
	}

	return model, nil
}

func (ht *HeadTracker) onNewLongestChain(ctx context.Context, headWithChain *models.Head) {
	ht.headMutex.Lock()
	defer ht.headMutex.Unlock()

	ht.logger.Debugw("HeadTracker initiating callbacks",
		"headNum", headWithChain.Number,
		"chainLength", headWithChain.ChainLength(),
		"numCallbacks", len(ht.callbacks),
	)

	ht.concurrentlyExecuteCallbacks(ctx, headWithChain)
}

func (ht *HeadTracker) concurrentlyExecuteCallbacks(ctx context.Context, headWithChain *models.Head) {
	wg := sync.WaitGroup{}
	wg.Add(len(ht.callbacks))
	for idx, trackable := range ht.callbacks {
		go func(i int, t HeadTrackable) {
			start := time.Now()
			t.OnNewLongestChain(ctx, headWithChain)
			elapsed := time.Since(start)
			ht.logger.Debugw(fmt.Sprintf("HeadTracker: finished callback %v in %s", i, elapsed), "callbackType", reflect.TypeOf(t), "callbackIdx", i, "blockNumber", headWithChain.Number, "time", elapsed, "id", "head_tracker")
			wg.Done()
		}(idx, trackable)
	}
	wg.Wait()
}

func (ht *HeadTracker) subscribeToHead() error {
	ht.headMutex.Lock()
	defer ht.headMutex.Unlock()

	ht.inHeaders = make(chan *etypes.Header)
	var rb *headRingBuffer
	rb, ht.outHeaders = newHeadRingBuffer(ht.inHeaders, int(ht.config.HeadTrackerMaxBufferSize), ht.logger)
	// It will autostop when we close inHeaders channel
	rb.Start()

	cancelHeadSum, err := ht.broadcaster.RegisterBlockHandler("", 0, func(ctx context.Context, header etypes.Header) {
		ht.inHeaders <- &header
	}, broadcaster.BlockOptions{
		Number: broadcaster.EachBlockNumber(),
	})
	if err != nil {
		return errors.Wrap(err, "Broadcaster#RegisterBlockHandler")
	}

	if err = verifyEthereumChainID(ht); err != nil {
		return errors.Wrap(err, "verifyEthereumChainID failed")
	}

	ht.cancelHeadSum = cancelHeadSum
	ht.connected = true

	ht.connect(ht.highestSeenHead)
	return nil
}

func (ht *HeadTracker) unsubscribeFromHead() error {
	ht.headMutex.Lock()
	defer ht.headMutex.Unlock()

	if !ht.connected {
		return nil
	}

	timedUnsubscribe(ht.cancelHeadSum, ht.logger)

	ht.connected = false
	ht.disconnect()
	close(ht.inHeaders)
	// Drain channel and wait for ringbuffer to close it
	for range ht.outHeaders {
	}
	return nil
}

func (ht *HeadTracker) setHighestSeenHeadFromDB() error {
	head, err := ht.store.LastHead()
	if err != nil {
		return err
	}
	ht.highestSeenHead = head
	return nil
}

// verifyEthereumChainID checks whether or not the ChainID from the config matches the ChainID
// reported by the Ethereum node.
func verifyEthereumChainID(ht *HeadTracker) error {
	ethereumChainID, err := ht.ethClient.ChainID(context.TODO())
	if err != nil {
		return err
	}

	if ethereumChainID.Cmp(ht.config.ChainID) != 0 {
		return fmt.Errorf(
			"ethereum ChainID doesn't match configured ChainID: config ID=%d, eth RPC ID=%d",
			ht.config.ChainID,
			ethereumChainID,
		)
	}
	return nil
}
