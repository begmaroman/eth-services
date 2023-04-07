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

// GethClient is an autogenerated mock type for the GethClient type
type GethClient struct {
	mock.Mock
}

// BalanceAt provides a mock function with given fields: ctx, account, blockNumber
func (_m *GethClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
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

// BlockByNumber provides a mock function with given fields: ctx, number
func (_m *GethClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
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

// CallContract provides a mock function with given fields: ctx, msg, blockNumber
func (_m *GethClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
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
func (_m *GethClient) ChainID(ctx context.Context) (*big.Int, error) {
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

// CodeAt provides a mock function with given fields: ctx, account, blockNumber
func (_m *GethClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
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

// EstimateGas provides a mock function with given fields: ctx, call
func (_m *GethClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
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
func (_m *GethClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
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

// HeaderByNumber provides a mock function with given fields: ctx, n
func (_m *GethClient) HeaderByNumber(ctx context.Context, n *big.Int) (*types.Header, error) {
	ret := _m.Called(ctx, n)

	var r0 *types.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (*types.Header, error)); ok {
		return rf(ctx, n)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Header); ok {
		r0 = rf(ctx, n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PendingCodeAt provides a mock function with given fields: ctx, account
func (_m *GethClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
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
func (_m *GethClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
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

// SendTransaction provides a mock function with given fields: ctx, tx
func (_m *GethClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
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
func (_m *GethClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
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

// SuggestGasPrice provides a mock function with given fields: ctx
func (_m *GethClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
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

// TransactionReceipt provides a mock function with given fields: ctx, txHash
func (_m *GethClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
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

type mockConstructorTestingTNewGethClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewGethClient creates a new instance of GethClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGethClient(t mockConstructorTestingTNewGethClient) *GethClient {
	mock := &GethClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
