package broadcaster

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type longPollingHeadStreamer struct {
	logger  logrus.FieldLogger
	client  Client
	chainID uint64

	longPollingTicker *time.Ticker

	lastBlockNumber     *big.Int
	lastBlockNumberLock sync.Mutex

	headersChan chan *types.Header
}

func NewLongPollingHeadStreamer(
	logger logrus.FieldLogger,
	client Client,
	blockTime time.Duration,
	chainID uint64,
) HeadStreamer {
	return &longPollingHeadStreamer{
		logger:            logger,
		client:            client,
		chainID:           chainID,
		longPollingTicker: time.NewTicker(blockTime),
		lastBlockNumber:   big.NewInt(0),
		headersChan:       make(chan *types.Header, headersChanCap),
	}
}

func (lp *longPollingHeadStreamer) Start(ctx context.Context) {
	go func() {
		for range lp.longPollingTicker.C {
			select {
			case <-ctx.Done():
				lp.longPollingTicker.Stop()
				return
			default:
				// Get the latest block
				header, err := lp.client.HeaderByNumber(ctx, nil)
				if err != nil {
					lp.logger.WithError(err).Error("failed to get header by number")
					failedGetHeaderByNumberCounter.WithLabelValues(big.NewInt(0).SetUint64(lp.chainID).String()).Inc()
					continue
				}

				lp.lastBlockNumberLock.Lock()
				lastHeaderNumber := lp.lastBlockNumber.Uint64()
				lp.lastBlockNumberLock.Unlock()
				currentHeaderNumber := header.Number.Uint64()

				if currentHeaderNumber <= lastHeaderNumber {
					continue
				}

				// if the difference between the last fetched block and the current one is more than 1, fallback headers
				if blockDiff := currentHeaderNumber - lastHeaderNumber; lastHeaderNumber > 0 && blockDiff > 1 {
					var errGrp errgroup.Group

					for blockToFetch := lastHeaderNumber + 1; blockToFetch < currentHeaderNumber-1; blockToFetch++ {
						blockToFetch := blockToFetch
						errGrp.Go(func() error {
							fallbackHeader, err := lp.client.HeaderByNumber(ctx, big.NewInt(0).SetUint64(blockToFetch))
							if err != nil {
								failedGetHeaderByNumberCounter.WithLabelValues(big.NewInt(0).SetUint64(lp.chainID).String()).Inc()
								return err
							}

							select {
							case lp.headersChan <- fallbackHeader:
							default:
								lp.logger.Warn("headers channel is full, skipping header")
							}

							return nil
						})
					}

					if err = errGrp.Wait(); err != nil {
						lp.logger.WithError(err).Error("failed to fallback headers")
					}
				}

				// Send the latest header to the stream
				select {
				case lp.headersChan <- header:
				default:
					lp.logger.Warn("headers channel is full, skipping header")
				}

				// Update the last handled block number
				lp.lastBlockNumberLock.Lock()
				lp.lastBlockNumber = big.NewInt(0).Set(header.Number)
				lp.lastBlockNumberLock.Unlock()
			}
		}
	}()
}

func (lp *longPollingHeadStreamer) Stop() {
	lp.longPollingTicker.Stop()
}

func (lp *longPollingHeadStreamer) Next() *types.Header {
	return <-lp.headersChan
}
