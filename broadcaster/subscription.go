package broadcaster

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type subscriptions struct {
	eventSubscribersLock sync.Mutex
	eventSubscribers     map[common.Address]map[common.Hash][]*eventSubscription

	blockSubscribersLock sync.Mutex
	blockSubscribers     []*blockSubscription
}

func newSubscriptions() *subscriptions {
	return &subscriptions{
		eventSubscribers: make(map[common.Address]map[common.Hash][]*eventSubscription),
		blockSubscribers: make([]*blockSubscription, 0),
	}
}

func (s *subscriptions) addBlockSubscription(bs *blockSubscription) {
	s.blockSubscribersLock.Lock()
	s.blockSubscribers = append(s.blockSubscribers, bs)
	s.blockSubscribersLock.Unlock()
}

func (s *subscriptions) getBlockSubscriptions(header types.Header) []*blockSubscription {
	s.blockSubscribersLock.Lock()

	var sbs []*blockSubscription

	for _, sb := range s.blockSubscribers {
		if sb.opts.Number != nil {
			if !sb.opts.Number(header.Number.Uint64()) {
				continue
			}
		}

		sbs = append(sbs, sb)
	}

	s.blockSubscribersLock.Unlock()

	return sbs
}

func (s *subscriptions) removeBlockSubscriptions(id string) {
	s.blockSubscribersLock.Lock()
	currentSubscribers := make([]*blockSubscription, len(s.blockSubscribers))
	copy(currentSubscribers, s.blockSubscribers)
	s.blockSubscribers = make([]*blockSubscription, 0)
	s.blockSubscribersLock.Unlock()

	for _, bs := range currentSubscribers {
		if bs.id != id {
			s.addBlockSubscription(bs)
		}
	}
}

func (s *subscriptions) addEventSubscription(es *eventSubscription) {
	s.eventSubscribersLock.Lock()
	defer s.eventSubscribersLock.Unlock()

	addrs := es.opts.Contracts
	if len(addrs) == 0 {
		addrs = []common.Address{zeroAddr}
	}

	for _, addr := range addrs {
		if _, ok := s.eventSubscribers[addr]; !ok {
			s.eventSubscribers[addr] = make(map[common.Hash][]*eventSubscription)
		}

		if es.opts.LogsWithTopics != nil {
			// TODO: Check values as well
			for topic := range es.opts.LogsWithTopics {
				if _, ok := s.eventSubscribers[addr][topic]; !ok {
					s.eventSubscribers[addr][topic] = make([]*eventSubscription, 0)
				}

				s.eventSubscribers[addr][topic] = append(s.eventSubscribers[addr][topic], es)
			}
		} else {
			if _, ok := s.eventSubscribers[addr][zeroHash]; !ok {
				s.eventSubscribers[addr][zeroHash] = make([]*eventSubscription, 0)
			}

			s.eventSubscribers[addr][zeroHash] = append(s.eventSubscribers[addr][zeroHash], es)
		}
	}
}

func (s *subscriptions) getEventSubscriptions(event types.Log) []*eventSubscription {
	s.eventSubscribersLock.Lock()

	var sbs []*eventSubscription

	// Exact match both by address and topic
	subs, ok := s.eventSubscribers[event.Address][event.Topics[0]]
	if ok {
		sbs = append(sbs, subs...)
	}

	// Match by topic
	subs, ok = s.eventSubscribers[zeroAddr][event.Topics[0]]
	if ok {
		sbs = append(sbs, subs...)
	}

	// Match by contract address
	subs, ok = s.eventSubscribers[event.Address][zeroHash]
	if ok {
		sbs = append(sbs, subs...)
	}

	// Match by any event
	subs, ok = s.eventSubscribers[zeroAddr][zeroHash]
	if ok {
		sbs = append(sbs, subs...)
	}

	s.eventSubscribersLock.Unlock()

	var filteredSubscriptions []*eventSubscription
	for _, sb := range sbs {
		if len(sb.opts.LogsWithTopics) > 0 && len(event.Topics) > 1 {
			topicValues := event.Topics[1:]
			if !filtersContainValues(topicValues, sb.opts.LogsWithTopics[event.Topics[0]]) {
				continue
			}
		}

		filteredSubscriptions = append(filteredSubscriptions, sb)
	}

	return filteredSubscriptions
}

func (s *subscriptions) removeEventSubscriptions(id string) {
	s.eventSubscribersLock.Lock()
	currentSubscribers := make(map[common.Address]map[common.Hash][]*eventSubscription)
	for key, val := range s.eventSubscribers {
		currentSubscribers[key] = val
	}
	s.eventSubscribers = make(map[common.Address]map[common.Hash][]*eventSubscription)
	s.eventSubscribersLock.Unlock()

	for _, subss := range currentSubscribers {
		for _, subs := range subss {
			for _, sub := range subs {
				if sub.id != id {
					s.addEventSubscription(sub)
				}
			}
		}
	}
}

func (s *subscriptions) buildFilters() ethereum.FilterQuery {
	s.eventSubscribersLock.Lock()
	var cntrcts []common.Address
	var topics []common.Hash
	for contract, topicsMap := range s.eventSubscribers {
		for topic := range topicsMap {
			if topic != zeroHash {
				topics = append(topics, topic)
			}
		}

		if contract != zeroAddr {
			cntrcts = append(cntrcts, contract)
		}
	}
	s.eventSubscribersLock.Unlock()

	return ethereum.FilterQuery{
		Addresses: cntrcts,
		Topics:    [][]common.Hash{topics},
	}
}

type eventSubscription struct {
	id         string
	inProgress int32
	handler    HandleEventFunc
	opts       EventOptions
}

func newEventSubscription(id string, handler HandleEventFunc, opts EventOptions) *eventSubscription {
	return &eventSubscription{
		id:      id,
		handler: handler,
		opts:    opts,
	}
}

func (s *eventSubscription) execute(ctx context.Context, event types.Log) error {
	if s.inProgress == 1 {
		return errors.New("subscriber is in progress")
	}

	atomic.StoreInt32(&s.inProgress, 1)
	s.handler(ctx, event)
	atomic.StoreInt32(&s.inProgress, 0)
	return nil
}

type blockSubscription struct {
	id         string
	inProgress int32
	handler    HandleBlockFunc
	opts       BlockOptions
}

func newBlockSubscription(id string, handler HandleBlockFunc, opts BlockOptions) *blockSubscription {
	return &blockSubscription{
		id:      id,
		handler: handler,
		opts:    opts,
	}
}

func (s *blockSubscription) execute(ctx context.Context, header types.Header) error {
	if s.inProgress == 1 {
		return errors.New("subscriber is in progress")
	}

	atomic.StoreInt32(&s.inProgress, 1)
	s.handler(ctx, header)
	atomic.StoreInt32(&s.inProgress, 0)
	return nil
}
