package broadcaster

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"github.com/begmaroman/eth-services/broadcaster/contracts"
)

const (
	testTriggerPerformanceEventID = "0xe5fc199a02ad9a3a02003f2440f5ea46b15b0a069c607e4bbbcde0fa705118b4"
	testChainID                   = uint64(1337)
	testWfID                      = "test-workflow"
)

func Test_SingleChainBroadcaster_EventTrigger(t *testing.T) {
	ctx := context.Background()
	logger := logrus.New()

	t.Run("successfully trigger subscriber on any event", func(t *testing.T) {
		_, txOpts, simulatedBackend, testContract := initSimulatedBackend(ctx, t)

		var consumed bool
		broadcaster := &singleChainBroadcaster{
			logger:        logger,
			client:        simulatedBackend,
			chainID:       testChainID,
			finalityDepth: big.NewInt(0),
			sbs:           newSubscriptions(),
			stop:          make(chan struct{}),
		}

		_, err := broadcaster.RegisterEventHandler(testWfID, testChainID, func(ctx context.Context, event types.Log) {
			consumed = true
		}, EventOptions{})
		require.NoError(t, err)

		err = broadcaster.Start(ctx)
		require.NoError(t, err)

		tx, err := testContract.Trigger(txOpts, []byte("qwe"), true)
		require.NoError(t, err)

		simulatedBackend.Commit()

		_, err = bind.WaitMined(ctx, simulatedBackend, tx)
		require.NoError(t, err)

		gomega.NewWithT(t).Eventually(func() bool {
			return consumed
		}).Should(gomega.BeTrue())

		err = broadcaster.Stop()
		require.NoError(t, err)
	})

	t.Run("successfully trigger subscriber on any event from the contract", func(t *testing.T) {
		addr, txOpts, simulatedBackend, testContract := initSimulatedBackend(ctx, t)

		var consumed bool
		broadcaster := &singleChainBroadcaster{
			logger:        logger,
			client:        simulatedBackend,
			chainID:       testChainID,
			finalityDepth: big.NewInt(0),
			sbs:           newSubscriptions(),
			stop:          make(chan struct{}),
		}

		_, err := broadcaster.RegisterEventHandler(testWfID, testChainID, func(ctx context.Context, event types.Log) {
			consumed = true
		}, EventOptions{
			Contracts: []common.Address{addr},
		})
		require.NoError(t, err)

		err = broadcaster.Start(ctx)
		require.NoError(t, err)

		tx, err := testContract.Trigger(txOpts, []byte("qwe"), true)
		require.NoError(t, err)

		simulatedBackend.Commit()

		_, err = bind.WaitMined(ctx, simulatedBackend, tx)
		require.NoError(t, err)

		gomega.NewWithT(t).Eventually(func() bool {
			return consumed
		}).Should(gomega.BeTrue())

		err = broadcaster.Stop()
		require.NoError(t, err)
	})

	t.Run("successfully trigger subscriber on TriggerPerformance event from the contract", func(t *testing.T) {
		addr, txOpts, simulatedBackend, testContract := initSimulatedBackend(ctx, t)

		var consumed bool
		broadcaster := &singleChainBroadcaster{
			logger:        logger,
			client:        simulatedBackend,
			chainID:       testChainID,
			finalityDepth: big.NewInt(0),
			sbs:           newSubscriptions(),
			stop:          make(chan struct{}),
		}

		_, err := broadcaster.RegisterEventHandler(testWfID, testChainID, func(ctx context.Context, event types.Log) {
			consumed = true
		}, EventOptions{
			Contracts: []common.Address{addr},
			LogsWithTopics: map[common.Hash][][]common.Hash{
				common.HexToHash(testTriggerPerformanceEventID): {},
			},
		})
		require.NoError(t, err)

		err = broadcaster.Start(ctx)
		require.NoError(t, err)

		tx, err := testContract.Trigger(txOpts, []byte("qwe"), true)
		require.NoError(t, err)

		simulatedBackend.Commit()

		_, err = bind.WaitMined(ctx, simulatedBackend, tx)
		require.NoError(t, err)

		gomega.NewWithT(t).Eventually(func() bool {
			return consumed
		}).Should(gomega.BeTrue())

		err = broadcaster.Stop()
		require.NoError(t, err)
	})

	t.Run("successfully trigger subscriber on TriggerPerformance event from the contract and event value", func(t *testing.T) {
		addr, txOpts, simulatedBackend, testContract := initSimulatedBackend(ctx, t)

		var consumed bool
		broadcaster := &singleChainBroadcaster{
			logger:        logger,
			client:        simulatedBackend,
			chainID:       testChainID,
			finalityDepth: big.NewInt(0),
			sbs:           newSubscriptions(),
			stop:          make(chan struct{}),
		}

		_, err := broadcaster.RegisterEventHandler(testWfID, testChainID, func(ctx context.Context, event types.Log) {
			consumed = true
		}, EventOptions{
			Contracts: []common.Address{addr},
			LogsWithTopics: map[common.Hash][][]common.Hash{
				common.HexToHash(testTriggerPerformanceEventID): {
					{},
					{common.BytesToHash([]byte("qwe"))},
					{},
				},
			},
		})
		require.NoError(t, err)

		err = broadcaster.Start(ctx)
		require.NoError(t, err)

		tx, err := testContract.Trigger(txOpts, []byte("qwe"), true)
		require.NoError(t, err)

		simulatedBackend.Commit()

		_, err = bind.WaitMined(ctx, simulatedBackend, tx)
		require.NoError(t, err)

		gomega.NewWithT(t).Eventually(func() bool {
			return consumed
		}).Should(gomega.BeTrue())

		err = broadcaster.Stop()
		require.NoError(t, err)
	})

	t.Run("successfully trigger subscriber on any event from the contract and event value", func(t *testing.T) {
		addr, txOpts, simulatedBackend, testContract := initSimulatedBackend(ctx, t)

		var consumed bool
		broadcaster := &singleChainBroadcaster{
			logger:        logger,
			client:        simulatedBackend,
			chainID:       testChainID,
			finalityDepth: big.NewInt(0),
			sbs:           newSubscriptions(),
			stop:          make(chan struct{}),
		}

		_, err := broadcaster.RegisterEventHandler(testWfID, testChainID, func(ctx context.Context, event types.Log) {
			consumed = true
		}, EventOptions{
			Contracts: []common.Address{addr},
			LogsWithTopics: map[common.Hash][][]common.Hash{
				zeroHash: {
					{},
					{common.BytesToHash([]byte("qwe"))},
					{},
				},
			},
		})
		require.NoError(t, err)

		err = broadcaster.Start(ctx)
		require.NoError(t, err)

		tx, err := testContract.Trigger(txOpts, []byte("qwe"), true)
		require.NoError(t, err)

		simulatedBackend.Commit()

		_, err = bind.WaitMined(ctx, simulatedBackend, tx)
		require.NoError(t, err)

		gomega.NewWithT(t).Eventually(func() bool {
			return consumed
		}).Should(gomega.BeTrue())

		err = broadcaster.Stop()
		require.NoError(t, err)
	})

	t.Run("successfully trigger subscriber on event value", func(t *testing.T) {
		_, txOpts, simulatedBackend, testContract := initSimulatedBackend(ctx, t)

		var consumed bool
		broadcaster := &singleChainBroadcaster{
			logger:        logger,
			client:        simulatedBackend,
			chainID:       testChainID,
			finalityDepth: big.NewInt(0),
			sbs:           newSubscriptions(),
			stop:          make(chan struct{}),
		}

		_, err := broadcaster.RegisterEventHandler(testWfID, testChainID, func(ctx context.Context, event types.Log) {
			consumed = true
		}, EventOptions{
			LogsWithTopics: map[common.Hash][][]common.Hash{
				zeroHash: {
					{},
					{common.BytesToHash([]byte("qwe"))},
					{},
				},
			},
		})
		require.NoError(t, err)

		err = broadcaster.Start(ctx)
		require.NoError(t, err)

		tx, err := testContract.Trigger(txOpts, []byte("qwe"), true)
		require.NoError(t, err)

		simulatedBackend.Commit()

		_, err = bind.WaitMined(ctx, simulatedBackend, tx)
		require.NoError(t, err)

		gomega.NewWithT(t).Eventually(func() bool {
			return consumed
		}).Should(gomega.BeTrue())

		err = broadcaster.Stop()
		require.NoError(t, err)
	})

	t.Run("wrong contract address provided", func(t *testing.T) {
		_, txOpts, simulatedBackend, testContract := initSimulatedBackend(ctx, t)

		var consumed bool
		broadcaster := &singleChainBroadcaster{
			logger:        logger,
			client:        simulatedBackend,
			chainID:       testChainID,
			finalityDepth: big.NewInt(0),
			sbs:           newSubscriptions(),
			stop:          make(chan struct{}),
		}

		_, err := broadcaster.RegisterEventHandler(testWfID, testChainID, func(ctx context.Context, event types.Log) {
			consumed = true
		}, EventOptions{
			Contracts: []common.Address{common.HexToAddress("0x20dacbf83c5de6658e14cbf7bcae5c15eca2eedecf1c66fbca928e4d351bea0d")},
		})
		require.NoError(t, err)

		err = broadcaster.Start(ctx)
		require.NoError(t, err)

		tx, err := testContract.Trigger(txOpts, []byte("qwe"), true)
		require.NoError(t, err)

		simulatedBackend.Commit()

		_, err = bind.WaitMined(ctx, simulatedBackend, tx)
		require.NoError(t, err)

		gomega.NewWithT(t).Eventually(func() bool {
			return consumed
		}).Should(gomega.BeFalse())

		err = broadcaster.Stop()
		require.NoError(t, err)
	})

	t.Run("wrong topic hash provided", func(t *testing.T) {
		addr, txOpts, simulatedBackend, testContract := initSimulatedBackend(ctx, t)

		var consumed bool
		broadcaster := &singleChainBroadcaster{
			logger:        logger,
			client:        simulatedBackend,
			chainID:       testChainID,
			finalityDepth: big.NewInt(0),
			sbs:           newSubscriptions(),
			stop:          make(chan struct{}),
		}

		_, err := broadcaster.RegisterEventHandler(testWfID, testChainID, func(ctx context.Context, event types.Log) {
			consumed = true
		}, EventOptions{
			Contracts: []common.Address{addr},
			LogsWithTopics: map[common.Hash][][]common.Hash{
				common.HexToHash("0x19dacbf83c5de6658e14cbf7bcae5c15eca2eedecf1c66fbca928e4d351bea0d"): {},
			},
		})
		require.NoError(t, err)

		err = broadcaster.Start(ctx)
		require.NoError(t, err)

		tx, err := testContract.Trigger(txOpts, []byte("qwe"), true)
		require.NoError(t, err)

		simulatedBackend.Commit()

		_, err = bind.WaitMined(ctx, simulatedBackend, tx)
		require.NoError(t, err)

		gomega.NewWithT(t).Eventually(func() bool {
			return consumed
		}).Should(gomega.BeFalse())

		err = broadcaster.Stop()
		require.NoError(t, err)
	})

	// TODO: This logic is not implemented yet
	/*t.Run("wrong topic value provided", func(t *testing.T) {
		addr, txOpts, simulatedBackend, testContract := initSimulatedBackend(ctx, t)

		var consumed bool
		broadcaster := &singleChainBroadcaster{
			logger:           logger,
			client:           simulatedBackend,
			chainID:          testChainID,
			finalityDepth:    big.NewInt(0),
			stop:             make(chan struct{}),
			eventSubscribers: make(map[common.Address]map[common.Hash][]*eventSubscription),
		}

		err := broadcaster.RegisterEventHandler(testChainID, func(ctx context.Context, event types.Log) {
			consumed = true
		}, EventOptions{
			Contracts: []common.Address{addr},
			LogsWithTopics: map[common.Hash][][]common.Hash{
				common.HexToHash(testTriggerPerformanceEventID): {
					{},
					{common.BytesToHash([]byte("wrong"))},
					{},
				},
			},
		})
		require.NoError(t, err)

		err = broadcaster.Start(ctx)
		require.NoError(t, err)

		tx, err := testContract.Trigger(txOpts, []byte("qwe"), true)
		require.NoError(t, err)

		simulatedBackend.Commit()

		_, err = bind.WaitMined(ctx, simulatedBackend, tx)
		require.NoError(t, err)

		gomega.NewWithT(t).Eventually(func() bool {
			return consumed
		}).Should(gomega.BeFalse())

		err = broadcaster.Stop()
		require.NoError(t, err)
	})*/
}

func Test_SingleChainBroadcaster_BlockTrigger(t *testing.T) {
	ctx := context.Background()
	logger := logrus.New()

	t.Run("each block trigger", func(t *testing.T) {
		_, _, simulatedBackend, _ := initSimulatedBackend(ctx, t)

		broadcaster := &singleChainBroadcaster{
			logger:        logger,
			client:        simulatedBackend,
			chainID:       testChainID,
			finalityDepth: big.NewInt(0),
			sbs:           newSubscriptions(),
			stop:          make(chan struct{}),
		}

		var consumed int
		_, err := broadcaster.RegisterBlockHandler(testWfID, testChainID, func(ctx context.Context, header types.Header) {
			consumed++
		}, BlockOptions{
			Number: EachBlockNumber(),
		})
		require.NoError(t, err)

		err = broadcaster.Start(ctx)
		require.NoError(t, err)

		// Produce 3 blocks and expect 3 executions
		simulatedBackend.Commit()
		simulatedBackend.Commit()
		simulatedBackend.Commit()

		gomega.NewWithT(t).Eventually(func() bool {
			return consumed == 3
		}).Should(gomega.BeTrue())

		err = broadcaster.Stop()
		require.NoError(t, err)
	})

	t.Run("each second block trigger", func(t *testing.T) {
		_, _, simulatedBackend, _ := initSimulatedBackend(ctx, t)

		broadcaster := &singleChainBroadcaster{
			logger:        logger,
			client:        simulatedBackend,
			chainID:       testChainID,
			finalityDepth: big.NewInt(0),
			sbs:           newSubscriptions(),
			stop:          make(chan struct{}),
		}

		var consumed int
		_, err := broadcaster.RegisterBlockHandler(testWfID, testChainID, func(ctx context.Context, header types.Header) {
			consumed++
		}, BlockOptions{
			Number: BlockNumberMod(3),
		})
		require.NoError(t, err)

		err = broadcaster.Start(ctx)
		require.NoError(t, err)

		// Produce 5 more blocks to have 2 executions
		simulatedBackend.Commit()
		simulatedBackend.Commit()
		simulatedBackend.Commit()
		simulatedBackend.Commit()
		simulatedBackend.Commit()

		gomega.NewWithT(t).Eventually(func() bool {
			return consumed == 2
		}).Should(gomega.BeTrue())

		err = broadcaster.Stop()
		require.NoError(t, err)
	})
}

func initSimulatedBackend(ctx context.Context, t *testing.T) (common.Address, *bind.TransactOpts, *backends.SimulatedBackend, *contracts.Counter) {
	t.Helper()

	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)
	txOpts, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(0).SetUint64(testChainID))
	require.NoError(t, err)

	simulatedBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		addr: {Balance: big.NewInt(params.Ether * 2)},
	}, 10000000)

	addr, tx, testContract, err := contracts.DeployCounter(txOpts, simulatedBackend, big.NewInt(1000), big.NewInt(1))
	require.NoError(t, err)

	simulatedBackend.Commit()

	_, err = bind.WaitMined(ctx, simulatedBackend, tx)
	require.NoError(t, err)

	return addr, txOpts, simulatedBackend, testContract
}
