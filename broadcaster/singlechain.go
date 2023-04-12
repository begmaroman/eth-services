package broadcaster

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

const (
	headsChanSize = 1000
)

// Client represents the behavior of the on-chain data provider
type Client interface {
	ethereum.LogFilterer
	ethereum.ChainReader
}

// Options contains options to create a broadcaster
type Options struct {
	ChainID       uint64
	FinalityDepth uint64
	ForceBlock    uint64
}

// singleChainBroadcaster implements Broadcaster interface.
// It uses a blockchain node as an event source.
type singleChainBroadcaster struct {
	logger        logrus.FieldLogger
	client        Client
	chainID       uint64
	finalityDepth *big.Int
	forceBlock    uint64
	sbs           *subscriptions
	stop          chan struct{}
	wg            sync.WaitGroup
}

// NewSingleChain is the constructor of singleChainBroadcaster
func NewSingleChain(logger logrus.FieldLogger, client Client, opts Options) (Broadcaster, error) {
	return &singleChainBroadcaster{
		logger:        logger,
		client:        client,
		chainID:       opts.ChainID,
		finalityDepth: big.NewInt(0).SetUint64(opts.FinalityDepth),
		forceBlock:    opts.ForceBlock,
		sbs:           newSubscriptions(),
		stop:          make(chan struct{}),
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
		l.logger.WithField("id", id).Info("subscription ha been unregistered")
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
	sub, err := l.client.SubscribeNewHead(ctx, ch)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe on new heads")
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				if err = ctx.Err(); err != nil {
					l.logger.WithError(err).Error("failed to get logs due to failed context")
				}
				return
			case err = <-sub.Err():
				l.logger.WithError(err).Error("failed to get logs due to failed subscription")
				sub, err = l.client.SubscribeNewHead(ctx, ch)
				if err != nil {
					l.logger.WithError(err).Fatal("failed to subscribe on new heads")
				}
			case <-l.stop:
				return
			case head := <-ch:
				targetBlock := big.NewInt(0)
				if l.forceBlock > 0 {
					targetBlock = targetBlock.SetUint64(l.forceBlock)
				} else {
					targetBlock = targetBlock.Sub(head.Number, l.finalityDepth)
				}

				logger := l.logger.WithFields(logrus.Fields{
					"block":        head.Number.String(),
					"target_block": targetBlock.String(),
				})
				logger.Debug("got new block")

				var errGroup errgroup.Group

				// Call block subscribers
				errGroup.Go(func() error {
					targetHeader := head
					if targetBlock.Cmp(targetHeader.Number) != 0 {
						targetHeader, err = l.client.HeaderByNumber(ctx, targetBlock)
						if err != nil {
							return errors.Wrap(err, "failed to get a target header")
						}
					}

					l.handleBlock(ctx, *targetHeader)

					return nil
				})

				// Call event subscribers
				errGroup.Go(func() error {
					filters := l.sbs.buildFilters()
					filters.FromBlock = targetBlock
					filters.ToBlock = targetBlock

					// Fetch logs from chain
					var logs []types.Log
					if logs, err = l.client.FilterLogs(ctx, filters); err != nil {
						return errors.Wrap(err, "failed to filter logs for the forced block")
					}

					if len(logs) == 0 {
						logger.Debug("no events for the block")
						return nil
					}
					logger.WithField("logs", len(logs)).Debug("found some logs")

					for _, log := range logs {
						if log.Removed {
							continue
						}

						l.handleEvent(ctx, log)
					}

					return nil
				})

				if err = errGroup.Wait(); err != nil {
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
