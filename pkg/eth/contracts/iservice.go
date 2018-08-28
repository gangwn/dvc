// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IServiceABI is the input ABI used to generate the binding from.
const IServiceABI = "[]"

// IServiceBin is the compiled bytecode used for deploying new contracts.
const IServiceBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a7230582071a5e33f8a3b315198f19dd3aa0914464288f3f907589a2462272a6f47f7545e0029`

// DeployIService deploys a new Ethereum contract, binding an instance of IService to it.
func DeployIService(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IService, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IServiceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IService{IServiceCaller: IServiceCaller{contract: contract}, IServiceTransactor: IServiceTransactor{contract: contract}, IServiceFilterer: IServiceFilterer{contract: contract}}, nil
}

// IService is an auto generated Go binding around an Ethereum contract.
type IService struct {
	IServiceCaller     // Read-only binding to the contract
	IServiceTransactor // Write-only binding to the contract
	IServiceFilterer   // Log filterer for contract events
}

// IServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IServiceSession struct {
	Contract     *IService         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IServiceCallerSession struct {
	Contract *IServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IServiceTransactorSession struct {
	Contract     *IServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IServiceRaw struct {
	Contract *IService // Generic contract binding to access the raw methods on
}

// IServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IServiceCallerRaw struct {
	Contract *IServiceCaller // Generic read-only contract binding to access the raw methods on
}

// IServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IServiceTransactorRaw struct {
	Contract *IServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIService creates a new instance of IService, bound to a specific deployed contract.
func NewIService(address common.Address, backend bind.ContractBackend) (*IService, error) {
	contract, err := bindIService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IService{IServiceCaller: IServiceCaller{contract: contract}, IServiceTransactor: IServiceTransactor{contract: contract}, IServiceFilterer: IServiceFilterer{contract: contract}}, nil
}

// NewIServiceCaller creates a new read-only instance of IService, bound to a specific deployed contract.
func NewIServiceCaller(address common.Address, caller bind.ContractCaller) (*IServiceCaller, error) {
	contract, err := bindIService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IServiceCaller{contract: contract}, nil
}

// NewIServiceTransactor creates a new write-only instance of IService, bound to a specific deployed contract.
func NewIServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*IServiceTransactor, error) {
	contract, err := bindIService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IServiceTransactor{contract: contract}, nil
}

// NewIServiceFilterer creates a new log filterer instance of IService, bound to a specific deployed contract.
func NewIServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*IServiceFilterer, error) {
	contract, err := bindIService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IServiceFilterer{contract: contract}, nil
}

// bindIService binds a generic wrapper to an already deployed contract.
func bindIService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IService *IServiceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IService.Contract.IServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IService *IServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IService.Contract.IServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IService *IServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IService.Contract.IServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IService *IServiceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IService *IServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IService *IServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IService.Contract.contract.Transact(opts, method, params...)
}
