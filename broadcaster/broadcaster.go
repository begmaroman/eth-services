package broadcaster

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// HandleEventFunc is the signature of a function to handle events
type HandleEventFunc func(ctx context.Context, event types.Log)

// HandleBlockFunc is the signature of a function to handle heads
type HandleBlockFunc func(ctx context.Context, header types.Header)

// EventOptions contains options to filter events
type EventOptions struct {
	// Contracts contains contract addresses to listen events from.
	Contracts []common.Address

	// Event types to receive, with value filter for each field in the event
	// No filter or an empty filter for a given field position mean: all values allowed
	// the key should be a result of AbigenLog.Topic() call
	// topic => topicValueFilters
	LogsWithTopics map[common.Hash][][]common.Hash
}

// BlockOptions contains options to filter blocks
type BlockOptions struct {
	// Number returns true if the block with the given number should be handled
	Number func(number uint64) bool
}

// Broadcaster represents a behavior of events broadcaster.
type Broadcaster interface {
	// RegisterEventHandler registers the given handler for specific events based on the given filters.
	RegisterEventHandler(id string, chainID uint64, handler HandleEventFunc, opts EventOptions) (func(), error)

	// RegisterBlockHandler registers the given handler for blocks based on the given filters.
	RegisterBlockHandler(id string, chainID uint64, handler HandleBlockFunc, opts BlockOptions) (func(), error)

	// Start starts broadcasting on-chain data to subscribers
	Start(ctx context.Context) error

	// Stop stops broadcasting
	Stop() error

	// Healthcheck performs a healthcheck of a broadcaster
	Healthcheck(ctx context.Context) error
}

// BlockNumberMod returns true if the block number % n is 0
func BlockNumberMod(n uint64) func(number uint64) bool {
	return func(number uint64) bool {
		return number%n == 0
	}
}

// EachBlockNumber returns true
func EachBlockNumber() func(number uint64) bool {
	return func(number uint64) bool {
		return true
	}
}
