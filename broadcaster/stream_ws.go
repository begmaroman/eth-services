package broadcaster

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type wsHeadStreamer struct {
	logger  logrus.FieldLogger
	client  Client
	chainID uint64
	stop    chan struct{}

	lastBlockNumber     *big.Int
	lastBlockNumberLock sync.Mutex

	headersChan chan *types.Header
}

func NewWSHeadStreamer(
	logger logrus.FieldLogger,
	client Client,
	chainID uint64,
) HeadStreamer {
	return &wsHeadStreamer{
		logger:          logger,
		client:          client,
		chainID:         chainID,
		stop:            make(chan struct{}),
		lastBlockNumber: big.NewInt(0),
		headersChan:     make(chan *types.Header, headersChanCap),
	}
}

func (ws *wsHeadStreamer) Start(ctx context.Context) {
	var ch chan *types.Header

	// Initialize a subscription
	sub := event.ResubscribeErr(5*time.Second, func(ctx context.Context, err error) (event.Subscription, error) {
		if err != nil {
			ws.logger.WithError(err).Error("resubscribing new heads with error")

			failedSubscribeNewHeadCounter.WithLabelValues(big.NewInt(0).SetUint64(ws.chainID).String()).Inc()
		}

		resubscribeNewHeadsSubscriptionCounter.WithLabelValues(big.NewInt(0).SetUint64(ws.chainID).String()).Inc()

		ch = make(chan *types.Header, headsChanSize)
		return ws.client.SubscribeNewHead(ctx, ch)
	})

	go func() {
		for {
			select {
			case <-ctx.Done():
				sub.Unsubscribe()
				return
			case err := <-sub.Err():
				if err == nil {
					continue
				}

				failedSubscribeNewHeadCounter.WithLabelValues(big.NewInt(0).SetUint64(ws.chainID).String()).Inc()

				ws.logger.WithError(err).Error("failed to get heads due to failed subscription")
			case <-ws.stop:
				sub.Unsubscribe()

				return
			case header := <-ch:
				ws.lastBlockNumberLock.Lock()
				lastHeaderNumber := ws.lastBlockNumber.Uint64()
				ws.lastBlockNumberLock.Unlock()
				currentHeaderNumber := header.Number.Uint64()

				// if the difference between the last fetched block and the current one is more than 1, fallback headers
				if blockDiff := currentHeaderNumber - lastHeaderNumber; lastHeaderNumber > 0 && blockDiff > 1 {
					var errGrp errgroup.Group

					for blockToFetch := lastHeaderNumber + 1; blockToFetch < currentHeaderNumber-1; blockToFetch++ {
						blockToFetch := blockToFetch
						errGrp.Go(func() error {
							fallbackHeader, err := ws.client.HeaderByNumber(ctx, big.NewInt(0).SetUint64(blockToFetch))
							if err != nil {
								return err
							}

							select {
							case ws.headersChan <- fallbackHeader:
							default:
								ws.logger.Warn("headers channel is full, skipping header")
							}

							return nil
						})
					}

					if err := errGrp.Wait(); err != nil {
						ws.logger.WithError(err).Error("failed to fallback headers")
					}
				}

				// Send the latest header to the stream
				select {
				case ws.headersChan <- header:
				default:
					ws.logger.Warn("headers channel is full, skipping header")
				}

				// Update the last handled block number
				ws.lastBlockNumberLock.Lock()
				ws.lastBlockNumber = big.NewInt(0).Set(header.Number)
				ws.lastBlockNumberLock.Unlock()
			}
		}
	}()
}

func (ws *wsHeadStreamer) Stop() {
	ws.stop <- struct{}{}
}

func (ws *wsHeadStreamer) Next() *types.Header {
	return <-ws.headersChan
}
