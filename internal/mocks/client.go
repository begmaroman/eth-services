// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	context "context"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// BalanceAt provides a mock function with given fields: ctx, account, blockNumber
func (_m *Client) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	ret := _m.Called(ctx, account, blockNumber)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) (*big.Int, error)); ok {
		return rf(ctx, account, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) *big.Int); ok {
		r0 = rf(ctx, account, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(ctx, account, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockByHash provides a mock function with given fields: ctx, hash
func (_m *Client) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	ret := _m.Called(ctx, hash)

	var r0 *types.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*types.Block, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Block); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockByNumber provides a mock function with given fields: ctx, number
func (_m *Client) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	ret := _m.Called(ctx, number)

	var r0 *types.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (*types.Block, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Block); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Call provides a mock function with given fields: result, method, args
func (_m *Client) Call(result interface{}, method string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, result, method)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(result, method, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CallContext provides a mock function with given fields: ctx, result, method, args
func (_m *Client) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, result, method)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, result, method, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CallContract provides a mock function with given fields: ctx, msg, blockNumber
func (_m *Client) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, msg, blockNumber)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg, *big.Int) ([]byte, error)); ok {
		return rf(ctx, msg, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg, *big.Int) []byte); ok {
		r0 = rf(ctx, msg, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.CallMsg, *big.Int) error); ok {
		r1 = rf(ctx, msg, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChainID provides a mock function with given fields: ctx
func (_m *Client) ChainID(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *Client) Close() {
	_m.Called()
}

// CodeAt provides a mock function with given fields: ctx, account, blockNumber
func (_m *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, account, blockNumber)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) ([]byte, error)); ok {
		return rf(ctx, account, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) []byte); ok {
		r0 = rf(ctx, account, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(ctx, account, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Dial provides a mock function with given fields: ctx
func (_m *Client) Dial(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EstimateGas provides a mock function with given fields: ctx, call
func (_m *Client) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	ret := _m.Called(ctx, call)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg) (uint64, error)); ok {
		return rf(ctx, call)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg) uint64); ok {
		r0 = rf(ctx, call)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.CallMsg) error); ok {
		r1 = rf(ctx, call)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterLogs provides a mock function with given fields: ctx, q
func (_m *Client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	ret := _m.Called(ctx, q)

	var r0 []types.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery) ([]types.Log, error)); ok {
		return rf(ctx, q)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery) []types.Log); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.FilterQuery) error); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeaderByHash provides a mock function with given fields: ctx, hash
func (_m *Client) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	ret := _m.Called(ctx, hash)

	var r0 *types.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*types.Header, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Header); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeaderByNumber provides a mock function with given fields: ctx, number
func (_m *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ret := _m.Called(ctx, number)

	var r0 *types.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (*types.Header, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Header); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PendingBalanceAt provides a mock function with given fields: ctx, account
func (_m *Client) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	ret := _m.Called(ctx, account)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) (*big.Int, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) *big.Int); ok {
		r0 = rf(ctx, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PendingCodeAt provides a mock function with given fields: ctx, account
func (_m *Client) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	ret := _m.Called(ctx, account)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) ([]byte, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) []byte); ok {
		r0 = rf(ctx, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PendingNonceAt provides a mock function with given fields: ctx, account
func (_m *Client) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	ret := _m.Called(ctx, account)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) (uint64, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) uint64); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PendingStorageAt provides a mock function with given fields: ctx, account, key
func (_m *Client) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	ret := _m.Called(ctx, account, key)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash) ([]byte, error)); ok {
		return rf(ctx, account, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash) []byte); ok {
		r0 = rf(ctx, account, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, common.Hash) error); ok {
		r1 = rf(ctx, account, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PendingTransactionCount provides a mock function with given fields: ctx
func (_m *Client) PendingTransactionCount(ctx context.Context) (uint, error) {
	ret := _m.Called(ctx)

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendRawTx provides a mock function with given fields: bytes
func (_m *Client) SendRawTx(bytes []byte) (common.Hash, error) {
	ret := _m.Called(bytes)

	var r0 common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (common.Hash, error)); ok {
		return rf(bytes)
	}
	if rf, ok := ret.Get(0).(func([]byte) common.Hash); ok {
		r0 = rf(bytes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(bytes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTransaction provides a mock function with given fields: ctx, tx
func (_m *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	ret := _m.Called(ctx, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) error); ok {
		r0 = rf(ctx, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscribeFilterLogs provides a mock function with given fields: ctx, q, ch
func (_m *Client) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	ret := _m.Called(ctx, q, ch)

	var r0 ethereum.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) (ethereum.Subscription, error)); ok {
		return rf(ctx, q, ch)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) ethereum.Subscription); ok {
		r0 = rf(ctx, q, ch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ethereum.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) error); ok {
		r1 = rf(ctx, q, ch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeNewHead provides a mock function with given fields: ctx, ch
func (_m *Client) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	ret := _m.Called(ctx, ch)

	var r0 ethereum.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, chan<- *types.Header) (ethereum.Subscription, error)); ok {
		return rf(ctx, ch)
	}
	if rf, ok := ret.Get(0).(func(context.Context, chan<- *types.Header) ethereum.Subscription); ok {
		r0 = rf(ctx, ch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ethereum.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, chan<- *types.Header) error); ok {
		r1 = rf(ctx, ch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuggestGasPrice provides a mock function with given fields: ctx
func (_m *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionCount provides a mock function with given fields: ctx, blockHash
func (_m *Client) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	ret := _m.Called(ctx, blockHash)

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (uint, error)); ok {
		return rf(ctx, blockHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) uint); ok {
		r0 = rf(ctx, blockHash)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionInBlock provides a mock function with given fields: ctx, blockHash, index
func (_m *Client) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	ret := _m.Called(ctx, blockHash, index)

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, uint) (*types.Transaction, error)); ok {
		return rf(ctx, blockHash, index)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, uint) *types.Transaction); ok {
		r0 = rf(ctx, blockHash, index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, uint) error); ok {
		r1 = rf(ctx, blockHash, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionReceipt provides a mock function with given fields: ctx, txHash
func (_m *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	ret := _m.Called(ctx, txHash)

	var r0 *types.Receipt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*types.Receipt, error)); ok {
		return rf(ctx, txHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Receipt); ok {
		r0 = rf(ctx, txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
