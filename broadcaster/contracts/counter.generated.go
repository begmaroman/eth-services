// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_testRange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interval\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"name\":\"Performed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"perform\",\"type\":\"bool\"}],\"name\":\"TriggerPerformance\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"check\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eligible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"}],\"name\":\"perform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"previousPerformBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_testRange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interval\",\"type\":\"uint256\"}],\"name\":\"setSpread\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"perform\",\"type\":\"bool\"}],\"name\":\"trigger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610919380380610919833981810160405281019061003291906100a1565b81600081905550806001819055506000600381905550436002819055506000600481905550600060058190555050506100e1565b600080fd5b6000819050919050565b61007e8161006b565b811461008957600080fd5b50565b60008151905061009b81610075565b92915050565b600080604083850312156100b8576100b7610066565b5b60006100c68582860161008c565b92505060206100d78582860161008c565b9150509250929050565b610829806100f06000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063917d895f11610071578063917d895f14610142578063947a36fb14610160578063bb6ae2cb1461017e578063c64b3bb51461019a578063d832d92f146101cb578063f6c6d4da146101e9576100a9565b80632cb15864146100ae57806361bc221a146100cc5780636250a13a146100ea5780637f407edf14610108578063806b984f14610124575b600080fd5b6100b6610205565b6040516100c391906103e4565b60405180910390f35b6100d461020b565b6040516100e191906103e4565b60405180910390f35b6100f2610211565b6040516100ff91906103e4565b60405180910390f35b610122600480360381019061011d9190610435565b610217565b005b61012c610239565b60405161013991906103e4565b60405180910390f35b61014a61023f565b60405161015791906103e4565b60405180910390f35b610168610245565b60405161017591906103e4565b60405180910390f35b610198600480360381019061019391906104da565b61024b565b005b6101b460048036038101906101af91906104da565b6102e2565b6040516101c29291906105d2565b60405180910390f35b6101d3610344565b6040516101e09190610602565b60405180910390f35b61020360048036038101906101fe9190610649565b61038b565b005b60045481565b60055481565b60005481565b8160008190555080600181905550600060048190555060006005819055505050565b60025481565b60035481565b60015481565b60006004540361025d57436004819055505b43600281905550600160055461027391906106d8565b6005819055503273ffffffffffffffffffffffffffffffffffffffff167fb55f31fdba783b65883517a934423678c525f9b6a83225968ffe3d08399883626004546002546003546005546040516102cd949392919061070c565b60405180910390a26002546003819055505050565b600060606102ee610344565b848481818080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505090509050915091509250929050565b600080600454036103585760019050610388565b600054600454436103699190610751565b1080156103855750600154600254436103829190610751565b10155b90505b90565b7fe5fc199a02ad9a3a02003f2440f5ea46b15b0a069c607e4bbbcde0fa705118b48383836040516103be939291906107c1565b60405180910390a1505050565b6000819050919050565b6103de816103cb565b82525050565b60006020820190506103f960008301846103d5565b92915050565b600080fd5b600080fd5b610412816103cb565b811461041d57600080fd5b50565b60008135905061042f81610409565b92915050565b6000806040838503121561044c5761044b6103ff565b5b600061045a85828601610420565b925050602061046b85828601610420565b9150509250929050565b600080fd5b600080fd5b600080fd5b60008083601f84011261049a57610499610475565b5b8235905067ffffffffffffffff8111156104b7576104b661047a565b5b6020830191508360018202830111156104d3576104d261047f565b5b9250929050565b600080602083850312156104f1576104f06103ff565b5b600083013567ffffffffffffffff81111561050f5761050e610404565b5b61051b85828601610484565b92509250509250929050565b60008115159050919050565b61053c81610527565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561057c578082015181840152602081019050610561565b60008484015250505050565b6000601f19601f8301169050919050565b60006105a482610542565b6105ae818561054d565b93506105be81856020860161055e565b6105c781610588565b840191505092915050565b60006040820190506105e76000830185610533565b81810360208301526105f98184610599565b90509392505050565b60006020820190506106176000830184610533565b92915050565b61062681610527565b811461063157600080fd5b50565b6000813590506106438161061d565b92915050565b600080600060408486031215610662576106616103ff565b5b600084013567ffffffffffffffff8111156106805761067f610404565b5b61068c86828701610484565b9350935050602061069f86828701610634565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006106e3826103cb565b91506106ee836103cb565b9250828201905080821115610706576107056106a9565b5b92915050565b600060808201905061072160008301876103d5565b61072e60208301866103d5565b61073b60408301856103d5565b61074860608301846103d5565b95945050505050565b600061075c826103cb565b9150610767836103cb565b925082820390508181111561077f5761077e6106a9565b5b92915050565b82818337600083830152505050565b60006107a0838561054d565b93506107ad838584610785565b6107b683610588565b840190509392505050565b600060408201905081810360008301526107dc818587610794565b90506107eb6020830184610533565b94935050505056fea2646970667358221220628fde3a30c1cfef2e8ce143839a818c722637062d257f660a3f18b04a706e7d64736f6c63430008130033",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// CounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterMetaData.Bin instead.
var CounterBin = CounterMetaData.Bin

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(auth *bind.TransactOpts, backend bind.ContractBackend, _testRange *big.Int, _interval *big.Int) (common.Address, *types.Transaction, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterBin), backend, _testRange, _interval)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// Check is a free data retrieval call binding the contract method 0xc64b3bb5.
//
// Solidity: function check(bytes data) view returns(bool, bytes)
func (_Counter *CounterCaller) Check(opts *bind.CallOpts, data []byte) (bool, []byte, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "check", data)

	if err != nil {
		return *new(bool), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// Check is a free data retrieval call binding the contract method 0xc64b3bb5.
//
// Solidity: function check(bytes data) view returns(bool, bytes)
func (_Counter *CounterSession) Check(data []byte) (bool, []byte, error) {
	return _Counter.Contract.Check(&_Counter.CallOpts, data)
}

// Check is a free data retrieval call binding the contract method 0xc64b3bb5.
//
// Solidity: function check(bytes data) view returns(bool, bytes)
func (_Counter *CounterCallerSession) Check(data []byte) (bool, []byte, error) {
	return _Counter.Contract.Check(&_Counter.CallOpts, data)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Counter *CounterCaller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Counter *CounterSession) Counter() (*big.Int, error) {
	return _Counter.Contract.Counter(&_Counter.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_Counter *CounterCallerSession) Counter() (*big.Int, error) {
	return _Counter.Contract.Counter(&_Counter.CallOpts)
}

// Eligible is a free data retrieval call binding the contract method 0xd832d92f.
//
// Solidity: function eligible() view returns(bool)
func (_Counter *CounterCaller) Eligible(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "eligible")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Eligible is a free data retrieval call binding the contract method 0xd832d92f.
//
// Solidity: function eligible() view returns(bool)
func (_Counter *CounterSession) Eligible() (bool, error) {
	return _Counter.Contract.Eligible(&_Counter.CallOpts)
}

// Eligible is a free data retrieval call binding the contract method 0xd832d92f.
//
// Solidity: function eligible() view returns(bool)
func (_Counter *CounterCallerSession) Eligible() (bool, error) {
	return _Counter.Contract.Eligible(&_Counter.CallOpts)
}

// InitialBlock is a free data retrieval call binding the contract method 0x2cb15864.
//
// Solidity: function initialBlock() view returns(uint256)
func (_Counter *CounterCaller) InitialBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "initialBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialBlock is a free data retrieval call binding the contract method 0x2cb15864.
//
// Solidity: function initialBlock() view returns(uint256)
func (_Counter *CounterSession) InitialBlock() (*big.Int, error) {
	return _Counter.Contract.InitialBlock(&_Counter.CallOpts)
}

// InitialBlock is a free data retrieval call binding the contract method 0x2cb15864.
//
// Solidity: function initialBlock() view returns(uint256)
func (_Counter *CounterCallerSession) InitialBlock() (*big.Int, error) {
	return _Counter.Contract.InitialBlock(&_Counter.CallOpts)
}

// Interval is a free data retrieval call binding the contract method 0x947a36fb.
//
// Solidity: function interval() view returns(uint256)
func (_Counter *CounterCaller) Interval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "interval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Interval is a free data retrieval call binding the contract method 0x947a36fb.
//
// Solidity: function interval() view returns(uint256)
func (_Counter *CounterSession) Interval() (*big.Int, error) {
	return _Counter.Contract.Interval(&_Counter.CallOpts)
}

// Interval is a free data retrieval call binding the contract method 0x947a36fb.
//
// Solidity: function interval() view returns(uint256)
func (_Counter *CounterCallerSession) Interval() (*big.Int, error) {
	return _Counter.Contract.Interval(&_Counter.CallOpts)
}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_Counter *CounterCaller) LastBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "lastBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_Counter *CounterSession) LastBlock() (*big.Int, error) {
	return _Counter.Contract.LastBlock(&_Counter.CallOpts)
}

// LastBlock is a free data retrieval call binding the contract method 0x806b984f.
//
// Solidity: function lastBlock() view returns(uint256)
func (_Counter *CounterCallerSession) LastBlock() (*big.Int, error) {
	return _Counter.Contract.LastBlock(&_Counter.CallOpts)
}

// PreviousPerformBlock is a free data retrieval call binding the contract method 0x917d895f.
//
// Solidity: function previousPerformBlock() view returns(uint256)
func (_Counter *CounterCaller) PreviousPerformBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "previousPerformBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviousPerformBlock is a free data retrieval call binding the contract method 0x917d895f.
//
// Solidity: function previousPerformBlock() view returns(uint256)
func (_Counter *CounterSession) PreviousPerformBlock() (*big.Int, error) {
	return _Counter.Contract.PreviousPerformBlock(&_Counter.CallOpts)
}

// PreviousPerformBlock is a free data retrieval call binding the contract method 0x917d895f.
//
// Solidity: function previousPerformBlock() view returns(uint256)
func (_Counter *CounterCallerSession) PreviousPerformBlock() (*big.Int, error) {
	return _Counter.Contract.PreviousPerformBlock(&_Counter.CallOpts)
}

// TestRange is a free data retrieval call binding the contract method 0x6250a13a.
//
// Solidity: function testRange() view returns(uint256)
func (_Counter *CounterCaller) TestRange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "testRange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestRange is a free data retrieval call binding the contract method 0x6250a13a.
//
// Solidity: function testRange() view returns(uint256)
func (_Counter *CounterSession) TestRange() (*big.Int, error) {
	return _Counter.Contract.TestRange(&_Counter.CallOpts)
}

// TestRange is a free data retrieval call binding the contract method 0x6250a13a.
//
// Solidity: function testRange() view returns(uint256)
func (_Counter *CounterCallerSession) TestRange() (*big.Int, error) {
	return _Counter.Contract.TestRange(&_Counter.CallOpts)
}

// Perform is a paid mutator transaction binding the contract method 0xbb6ae2cb.
//
// Solidity: function perform(bytes performData) returns()
func (_Counter *CounterTransactor) Perform(opts *bind.TransactOpts, performData []byte) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "perform", performData)
}

// Perform is a paid mutator transaction binding the contract method 0xbb6ae2cb.
//
// Solidity: function perform(bytes performData) returns()
func (_Counter *CounterSession) Perform(performData []byte) (*types.Transaction, error) {
	return _Counter.Contract.Perform(&_Counter.TransactOpts, performData)
}

// Perform is a paid mutator transaction binding the contract method 0xbb6ae2cb.
//
// Solidity: function perform(bytes performData) returns()
func (_Counter *CounterTransactorSession) Perform(performData []byte) (*types.Transaction, error) {
	return _Counter.Contract.Perform(&_Counter.TransactOpts, performData)
}

// SetSpread is a paid mutator transaction binding the contract method 0x7f407edf.
//
// Solidity: function setSpread(uint256 _testRange, uint256 _interval) returns()
func (_Counter *CounterTransactor) SetSpread(opts *bind.TransactOpts, _testRange *big.Int, _interval *big.Int) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "setSpread", _testRange, _interval)
}

// SetSpread is a paid mutator transaction binding the contract method 0x7f407edf.
//
// Solidity: function setSpread(uint256 _testRange, uint256 _interval) returns()
func (_Counter *CounterSession) SetSpread(_testRange *big.Int, _interval *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.SetSpread(&_Counter.TransactOpts, _testRange, _interval)
}

// SetSpread is a paid mutator transaction binding the contract method 0x7f407edf.
//
// Solidity: function setSpread(uint256 _testRange, uint256 _interval) returns()
func (_Counter *CounterTransactorSession) SetSpread(_testRange *big.Int, _interval *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.SetSpread(&_Counter.TransactOpts, _testRange, _interval)
}

// Trigger is a paid mutator transaction binding the contract method 0xf6c6d4da.
//
// Solidity: function trigger(bytes data, bool perform) returns()
func (_Counter *CounterTransactor) Trigger(opts *bind.TransactOpts, data []byte, perform bool) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "trigger", data, perform)
}

// Trigger is a paid mutator transaction binding the contract method 0xf6c6d4da.
//
// Solidity: function trigger(bytes data, bool perform) returns()
func (_Counter *CounterSession) Trigger(data []byte, perform bool) (*types.Transaction, error) {
	return _Counter.Contract.Trigger(&_Counter.TransactOpts, data, perform)
}

// Trigger is a paid mutator transaction binding the contract method 0xf6c6d4da.
//
// Solidity: function trigger(bytes data, bool perform) returns()
func (_Counter *CounterTransactorSession) Trigger(data []byte, perform bool) (*types.Transaction, error) {
	return _Counter.Contract.Trigger(&_Counter.TransactOpts, data, perform)
}

// CounterPerformedIterator is returned from FilterPerformed and is used to iterate over the raw logs and unpacked data for Performed events raised by the Counter contract.
type CounterPerformedIterator struct {
	Event *CounterPerformed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CounterPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterPerformed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CounterPerformed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CounterPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterPerformed represents a Performed event raised by the Counter contract.
type CounterPerformed struct {
	From          common.Address
	InitialBlock  *big.Int
	LastBlock     *big.Int
	PreviousBlock *big.Int
	Counter       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPerformed is a free log retrieval operation binding the contract event 0xb55f31fdba783b65883517a934423678c525f9b6a83225968ffe3d0839988362.
//
// Solidity: event Performed(address indexed from, uint256 initialBlock, uint256 lastBlock, uint256 previousBlock, uint256 counter)
func (_Counter *CounterFilterer) FilterPerformed(opts *bind.FilterOpts, from []common.Address) (*CounterPerformedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Counter.contract.FilterLogs(opts, "Performed", fromRule)
	if err != nil {
		return nil, err
	}
	return &CounterPerformedIterator{contract: _Counter.contract, event: "Performed", logs: logs, sub: sub}, nil
}

// WatchPerformed is a free log subscription operation binding the contract event 0xb55f31fdba783b65883517a934423678c525f9b6a83225968ffe3d0839988362.
//
// Solidity: event Performed(address indexed from, uint256 initialBlock, uint256 lastBlock, uint256 previousBlock, uint256 counter)
func (_Counter *CounterFilterer) WatchPerformed(opts *bind.WatchOpts, sink chan<- *CounterPerformed, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Counter.contract.WatchLogs(opts, "Performed", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterPerformed)
				if err := _Counter.contract.UnpackLog(event, "Performed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePerformed is a log parse operation binding the contract event 0xb55f31fdba783b65883517a934423678c525f9b6a83225968ffe3d0839988362.
//
// Solidity: event Performed(address indexed from, uint256 initialBlock, uint256 lastBlock, uint256 previousBlock, uint256 counter)
func (_Counter *CounterFilterer) ParsePerformed(log types.Log) (*CounterPerformed, error) {
	event := new(CounterPerformed)
	if err := _Counter.contract.UnpackLog(event, "Performed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterTriggerPerformanceIterator is returned from FilterTriggerPerformance and is used to iterate over the raw logs and unpacked data for TriggerPerformance events raised by the Counter contract.
type CounterTriggerPerformanceIterator struct {
	Event *CounterTriggerPerformance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CounterTriggerPerformanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterTriggerPerformance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CounterTriggerPerformance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CounterTriggerPerformanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterTriggerPerformanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterTriggerPerformance represents a TriggerPerformance event raised by the Counter contract.
type CounterTriggerPerformance struct {
	Data    []byte
	Perform bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTriggerPerformance is a free log retrieval operation binding the contract event 0xe5fc199a02ad9a3a02003f2440f5ea46b15b0a069c607e4bbbcde0fa705118b4.
//
// Solidity: event TriggerPerformance(bytes data, bool perform)
func (_Counter *CounterFilterer) FilterTriggerPerformance(opts *bind.FilterOpts) (*CounterTriggerPerformanceIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "TriggerPerformance")
	if err != nil {
		return nil, err
	}
	return &CounterTriggerPerformanceIterator{contract: _Counter.contract, event: "TriggerPerformance", logs: logs, sub: sub}, nil
}

// WatchTriggerPerformance is a free log subscription operation binding the contract event 0xe5fc199a02ad9a3a02003f2440f5ea46b15b0a069c607e4bbbcde0fa705118b4.
//
// Solidity: event TriggerPerformance(bytes data, bool perform)
func (_Counter *CounterFilterer) WatchTriggerPerformance(opts *bind.WatchOpts, sink chan<- *CounterTriggerPerformance) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "TriggerPerformance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterTriggerPerformance)
				if err := _Counter.contract.UnpackLog(event, "TriggerPerformance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTriggerPerformance is a log parse operation binding the contract event 0xe5fc199a02ad9a3a02003f2440f5ea46b15b0a069c607e4bbbcde0fa705118b4.
//
// Solidity: event TriggerPerformance(bytes data, bool perform)
func (_Counter *CounterFilterer) ParseTriggerPerformance(log types.Log) (*CounterTriggerPerformance, error) {
	event := new(CounterTriggerPerformance)
	if err := _Counter.contract.UnpackLog(event, "TriggerPerformance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
