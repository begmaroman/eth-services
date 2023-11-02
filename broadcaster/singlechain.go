package broadcaster

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/event"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

const (
	headsChanSize        = 1000
	blockUpdateThreshold = 2
)

// Client represents the behavior of the on-chain data provider
type Client interface {
	ethereum.LogFilterer
	ethereum.ChainReader
}

// Options contains options to create a broadcaster
type Options struct {
	ChainID    uint64
	ForceBlock uint64
	BlockTime  time.Duration
}

// singleChainBroadcaster implements Broadcaster interface.
// It uses a blockchain node as an event source.
type singleChainBroadcaster struct {
	logger     logrus.FieldLogger
	client     Client
	chainID    uint64
	forceBlock uint64
	blockTime  time.Duration
	sbs        *subscriptions
	stop       chan struct{}
	wg         sync.WaitGroup

	lastHeadLock      sync.Mutex
	lastHead          *big.Int
	lastHeadUpdatedAt time.Time
}

// NewSingleChain is the constructor of singleChainBroadcaster
func NewSingleChain(logger logrus.FieldLogger, client Client, opts Options) (Broadcaster, error) {
	return &singleChainBroadcaster{
		logger:            logger,
		client:            client,
		chainID:           opts.ChainID,
		forceBlock:        opts.ForceBlock,
		blockTime:         opts.BlockTime,
		sbs:               newSubscriptions(),
		stop:              make(chan struct{}),
		lastHead:          big.NewInt(0),
		lastHeadUpdatedAt: time.Now(),
	}, nil
}

// RegisterEventHandler registers the given handler using the given events filters
func (l *singleChainBroadcaster) RegisterEventHandler(id string, chainID uint64, handler HandleEventFunc, opts EventOptions) (func(), error) {
	if chainID != l.chainID {
		return nil, fmt.Errorf("the given chain ID %d does not match with the broadcaster's one %d", chainID, l.chainID)
	}

	l.sbs.addEventSubscription(newEventSubscription(id, handler, opts))

	return func() {
		l.sbs.removeEventSubscriptions(id)
		l.logger.WithField("id", id).Info("subscription has been unregistered")
	}, nil
}

// RegisterBlockHandler registers the given block handler
func (l *singleChainBroadcaster) RegisterBlockHandler(id string, chainID uint64, handler HandleBlockFunc, opts BlockOptions) (func(), error) {
	if chainID != l.chainID {
		return nil, fmt.Errorf("the given chain ID %d does not match with the broadcaster's one %d", chainID, l.chainID)
	}

	l.sbs.addBlockSubscription(newBlockSubscription(id, handler, opts))

	return func() {
		l.sbs.removeBlockSubscriptions(id)
		l.logger.WithField("id", id).Info("subscription has been unregistered")
	}, nil
}

// Start starts broadcasting messages
func (l *singleChainBroadcaster) Start(ctx context.Context) error {
	// Initialize a subscription
	ch := make(chan *types.Header, headsChanSize)

	sub := event.Resubscribe(2*time.Second, func(ctx context.Context) (event.Subscription, error) {
		resubscribeNewHeadsSubscriptionCounter.WithLabelValues(big.NewInt(0).SetUint64(l.chainID).String()).Inc()

		return l.client.SubscribeNewHead(ctx, ch)
	})

	// Handle new heads
	go func() {
		for {
			select {
			case <-ctx.Done():
				sub.Unsubscribe()

				l.logger.Info("stopping new head subscription due to canceled context")

				if err := ctx.Err(); err != nil {
					l.logger.WithError(err).Error("context cancelled with error")
				}

				return
			case err := <-sub.Err():
				if err == nil {
					continue
				}

				failedSubscribeNewHeadCounter.WithLabelValues(big.NewInt(0).SetUint64(l.chainID).String()).Inc()

				l.logger.WithError(err).Error("failed to get heads due to failed subscription")
			case <-l.stop:
				sub.Unsubscribe()

				return
			case head := <-ch:
				targetBlock := head.Number
				if l.forceBlock > 0 {
					targetBlock = targetBlock.SetUint64(l.forceBlock)
				}

				// Check if this head has been proceeded already
				if targetBlock.Cmp(l.lastHead) <= 0 {
					continue
				}

				// Update the last handled head
				l.lastHeadLock.Lock()
				l.lastHead = new(big.Int).Set(targetBlock)
				l.lastHeadUpdatedAt = time.Now()
				l.lastHeadLock.Unlock()

				logger := l.logger.WithField("block", targetBlock.String())
				logger.Debug("got new block")

				var errGroup errgroup.Group

				// Call block subscribers
				if l.sbs.existBlockSubscribers() {
					errGroup.Go(func() error {
						targetHeader := head
						if targetBlock.Cmp(targetHeader.Number) != 0 {
							var err error
							if targetHeader, err = l.client.HeaderByNumber(ctx, targetBlock); err != nil {
								return errors.Wrap(err, "failed to get a target header")
							}
						}

						logger.Debug("found head subscribers for the current block")

						l.handleBlock(ctx, *targetHeader)

						return nil
					})
				}

				// Call event subscribers
				if l.sbs.existEventSubscribers() {
					errGroup.Go(func() error {
						filters := l.sbs.buildFilters()
						filters.FromBlock = targetBlock
						filters.ToBlock = targetBlock

						// Fetch logs from chain
						logs, err := l.client.FilterLogs(ctx, filters)
						if err != nil {
							return errors.Wrap(err, "failed to filter logs for the forced block")
						}

						if len(logs) == 0 {
							logger.Debug("no events for the block")
							return nil
						}

						logger.WithField("logs", len(logs)).Debug("found some events to be handled")

						for _, log := range logs {
							if log.Removed {
								continue
							}

							l.handleEvent(ctx, log)
						}

						return nil
					})
				}

				if err := errGroup.Wait(); err != nil {
					logger.Error(err)
				}
			}
		}
	}()

	return nil
}

// Stop stops broadcasting
func (l *singleChainBroadcaster) Stop() error {
	l.stop <- struct{}{}
	l.wg.Wait()
	return nil
}

// Healthcheck performs a healthcheck
func (l *singleChainBroadcaster) Healthcheck(ctx context.Context) error {
	if stuck, lastUpdate := l.isHeadsSubscriptionStuck(); stuck {
		failedHealthcheckCounter.WithLabelValues(big.NewInt(0).SetUint64(l.chainID).String()).Inc()

		return fmt.Errorf("new head is missing for %s for chain %d", lastUpdate, l.chainID)
	}

	return nil
}

func (l *singleChainBroadcaster) isHeadsSubscriptionStuck() (bool, time.Duration) {
	lastUpdate := time.Now().Sub(l.lastHeadUpdatedAt)
	return lastUpdate > l.blockTime*blockUpdateThreshold*2, lastUpdate
}

// handleEvent handles the given event
func (l *singleChainBroadcaster) handleEvent(ctx context.Context, event types.Log) {
	sbs := l.sbs.getEventSubscriptions(event)
	if len(sbs) == 0 {
		return
	}

	for _, s := range sbs {
		l.wg.Add(1)
		go func(s *eventSubscription) {
			if err := s.execute(ctx, event); err != nil {
				l.logger.WithError(err).Debug("failed to execute subscriber on event")
			}
			l.wg.Done()
		}(s)
	}
}

// handleEvent handles the given event
func (l *singleChainBroadcaster) handleBlock(ctx context.Context, header types.Header) {
	sbs := l.sbs.getBlockSubscriptions(header)
	if len(sbs) == 0 {
		return
	}

	for _, s := range sbs {
		l.wg.Add(1)
		go func(s *blockSubscription) {
			if err := s.execute(ctx, header); err != nil {
				l.logger.WithError(err).Debug("failed to execute subscriber on block")
			}
			l.wg.Done()
		}(s)
	}
}
