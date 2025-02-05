package client

import (
	"context"
	"math/big"
	"net/url"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"

	esTypes "github.com/begmaroman/eth-services/types"
)

//go:generate mockery --name Client --output ../internal/mocks/ --case=underscore
//go:generate mockery --name GethClient --output ../internal/mocks/ --case=underscore
//go:generate mockery --name RPCClient --output ../internal/mocks/ --case=underscore

// Client is the interface used to interact with an ethereum node.
type Client interface {
	GethClient

	Dial(ctx context.Context) error
	Close()

	Call(result interface{}, method string, args ...interface{}) error
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
}

// GethClient is an interface that represents go-ethereum's own ethclient
// https://github.com/ethereum/go-ethereum/blob/master/ethclient/ethclient.go
type GethClient interface {
	ethereum.TransactionSender
	ethereum.LogFilterer
	ethereum.GasPricer
	ethereum.GasEstimator
	ethereum.ChainReader
	ethereum.PendingStateReader

	ChainID(ctx context.Context) (*big.Int, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error)

	// SuggestGasTipCap retrieves the currently suggested 1559 priority fee to allow
	// a timely execution of a transaction.
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)

	FeeHistory(ctx context.Context, blockCount uint64, lastBlock *big.Int, rewardPercentiles []float64) (*ethereum.FeeHistory, error)
}

// RPCClient is an interface that represents go-ethereum's own rpc.Client.
// https://github.com/ethereum/go-ethereum/blob/master/rpc/client.go
type RPCClient interface {
	Call(result interface{}, method string, args ...interface{}) error
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
	EthSubscribe(ctx context.Context, channel interface{}, args ...interface{}) (ethereum.Subscription, error)
	Close()
}

// Impl implements the ethereum Client interface using a CallerSubscriber instance.
type Impl struct {
	GethClient
	RPCClient
	url                  *url.URL // For reestablishing the connection after a disconnect
	SecondaryGethClients []GethClient
	SecondaryRPCClients  []RPCClient
	secondaryURLs        []*url.URL
	mocked               bool
	logger               esTypes.Logger
}

var _ Client = (*Impl)(nil)

// NewImpl creates a new client implementation
func NewImpl(config *esTypes.Config) (*Impl, error) {
	rpcURL := config.RPCURL
	if rpcURL.Scheme != "ws" && rpcURL.Scheme != "wss" &&
		rpcURL.Scheme != "http" && rpcURL.Scheme != "https" {
		return nil, errors.Errorf("Ethereum URL scheme must be ws(s) or http(s): %s", rpcURL)
	}

	secondaryRPCURLs := config.SecondaryRPCURLs
	for _, url := range secondaryRPCURLs {
		if url.Scheme != "http" && url.Scheme != "https" {
			return nil, errors.Errorf("secondary Ethereum RPC URL scheme must be http(s): %s", url)
		}
	}

	return &Impl{
		url:           rpcURL,
		secondaryURLs: secondaryRPCURLs,
		logger:        config.Logger,
	}, nil
}

func (client *Impl) Dial(ctx context.Context) error {
	client.logger.Debugw("eth.Client#Dial(...)")
	if client.mocked {
		return nil
	} else if client.RPCClient != nil || client.GethClient != nil {
		panic("eth.Client.Dial(...) should only be called once during the application's lifetime.")
	}

	rpcClient, err := rpc.DialContext(ctx, client.url.String())
	if err != nil {
		return err
	}
	client.RPCClient = &rpcClientWrapper{rpcClient}
	client.GethClient = ethclient.NewClient(rpcClient)

	client.SecondaryGethClients = []GethClient{}
	client.SecondaryRPCClients = []RPCClient{}
	for _, url := range client.secondaryURLs {
		secondaryRPCClient, err := rpc.DialContext(ctx, url.String())
		if err != nil {
			return err
		}
		client.SecondaryRPCClients = append(client.SecondaryRPCClients, &rpcClientWrapper{secondaryRPCClient})
		client.SecondaryGethClients = append(client.SecondaryGethClients, ethclient.NewClient(secondaryRPCClient))
	}
	return nil
}

// SendRawTx sends a signed transaction to the transaction pool.
func (client *Impl) SendRawTx(bytes []byte) (common.Hash, error) {
	client.logger.Debugw("eth.Client#SendRawTx(...)",
		"bytes", bytes,
	)
	result := common.Hash{}
	err := client.RPCClient.Call(&result, "eth_sendRawTransaction", hexutil.Encode(bytes))
	return result, err
}

// TransactionReceipt wraps the GethClient's `TransactionReceipt` method so that we can ignore the
// error that arises when we're talking to a Parity node that has no receipt yet.
func (client *Impl) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	client.logger.Debugw("eth.Client#TransactionReceipt(...)",
		"txHash", txHash,
	)
	return client.GethClient.TransactionReceipt(ctx, txHash)
}

func (client *Impl) ChainID(ctx context.Context) (*big.Int, error) {
	client.logger.Debugw("eth.Client#ChainID(...)")
	return client.GethClient.ChainID(ctx)
}

// SendTransaction also uses the secondary HTTP RPC URL if set
func (client *Impl) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	client.logger.Debugw("eth.Client#SendTransaction(...)",
		"tx", tx,
	)

	var wg sync.WaitGroup
	for _, gethClient := range client.SecondaryGethClients {
		// Parallel send to secondary node
		client.logger.Tracew("eth.SecondaryClient#SendTransaction(...)", "tx", tx)

		wg.Add(1)
		go func(gethClient GethClient) {
			defer wg.Done()

			err := NewSendError(gethClient.SendTransaction(ctx, tx))
			if err == nil || err.IsNonceTooLowError() || err.IsTransactionAlreadyInMempool() {
				// Nonce too low or transaction known errors are expected since
				// the primary SendTransaction may well have succeeded already
				return
			}
			client.logger.Warnf("secondary eth client returned error", "err", err, "tx", tx)
		}(gethClient)
	}
	wg.Wait()

	return client.GethClient.SendTransaction(ctx, tx)
}

func (client *Impl) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	client.logger.Debugw("eth.Client#PendingNonceAt(...)",
		"account", account,
	)
	return client.GethClient.PendingNonceAt(ctx, account)
}

func (client *Impl) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	client.logger.Debugw("eth.Client#PendingCodeAt(...)",
		"account", account,
	)
	return client.GethClient.PendingCodeAt(ctx, account)
}

func (client *Impl) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	client.logger.Debugw("eth.Client#EstimateGas(...)",
		"call", call,
	)
	return client.GethClient.EstimateGas(ctx, call)
}

func (client *Impl) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	client.logger.Debugw("eth.Client#SuggestGasPrice()")
	return client.GethClient.SuggestGasPrice(ctx)
}

func (client *Impl) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	client.logger.Debugw("eth.Client#BlockByNumber(...)",
		"number", number,
	)
	return client.GethClient.BlockByNumber(ctx, number)
}

func (client *Impl) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	client.logger.Debugw("eth.Client#HeaderByNumber(...)",
		"number", number,
	)
	return client.GethClient.HeaderByNumber(ctx, number)
}

func (client *Impl) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	client.logger.Debug("eth.Client#BalanceAt(...)",
		"account", account,
		"blockNumber", blockNumber,
	)
	return client.GethClient.BalanceAt(ctx, account, blockNumber)
}

func (client *Impl) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	client.logger.Debugw("eth.Client#FilterLogs(...)",
		"q", q,
	)
	return client.GethClient.FilterLogs(ctx, q)
}

func (client *Impl) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	if client.url.Scheme != "ws" && client.url.Scheme != "wss" {
		return nil, errors.Errorf("subscriptions allowed for ws(s) only: %s", client.url)
	}

	client.logger.Debugw("eth.Client#SubscribeFilterLogs(...)",
		"q", q,
	)

	return client.GethClient.SubscribeFilterLogs(ctx, q, ch)
}

func (client *Impl) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	if client.url.Scheme != "ws" && client.url.Scheme != "wss" {
		return nil, errors.Errorf("subscriptions allowed for ws(s) only: %s", client.url)
	}

	client.logger.Debugw("eth.Client#SubscribeNewHead(...)")

	return client.GethClient.SubscribeNewHead(ctx, ch)
}

// SuggestGasTipCap retrieves the currently suggested 1559 priority fee to allow
// a timely execution of a transaction.
func (client *Impl) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	client.logger.Debugw("eth.Client#SuggestGasTipCap(...)")
	return client.GethClient.SuggestGasTipCap(ctx)
}

// FeeHistory returns the collection of historical gas information
func (client *Impl) FeeHistory(ctx context.Context, blockCount uint64, lastBlock *big.Int, rewardPercentiles []float64) (*ethereum.FeeHistory, error) {
	client.logger.Debugw("eth.Client#FeeHistory(...)")
	return client.GethClient.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
}

// TODO: remove this wrapper type once EthMock is no longer in use in upstream.
type rpcClientWrapper struct {
	*rpc.Client
}

func (w *rpcClientWrapper) EthSubscribe(ctx context.Context, channel interface{}, args ...interface{}) (ethereum.Subscription, error) {
	return w.Client.EthSubscribe(ctx, channel, args...)
}

func (client *Impl) Call(result interface{}, method string, args ...interface{}) error {
	client.logger.Debugw("eth.Client#Call(...)",
		"method", method,
		"args", args,
	)
	return client.RPCClient.Call(result, method, args...)
}

func (client *Impl) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	client.logger.Debugw("eth.Client#Call(...)",
		"method", method,
		"args", args,
	)
	return client.RPCClient.CallContext(ctx, result, method, args...)
}
