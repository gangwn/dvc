// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// CCSServiceABI is the input ABI used to generate the binding from.
const CCSServiceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"confId\",\"type\":\"string\"}],\"name\":\"getCCS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"serviceManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"confId\",\"type\":\"string\"},{\"name\":\"ccsAddress\",\"type\":\"address\"}],\"name\":\"newJob\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ip\",\"type\":\"string\"},{\"name\":\"port\",\"type\":\"int256\"}],\"name\":\"registerCCS\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"serviceManagerAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"confId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"ccsAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"ip\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"port\",\"type\":\"int256\"}],\"name\":\"NewJobCreated\",\"type\":\"event\"}]"

// CCSService is an auto generated Go binding around an Ethereum contract.
type CCSService struct {
	CCSServiceCaller     // Read-only binding to the contract
	CCSServiceTransactor // Write-only binding to the contract
	CCSServiceFilterer   // Log filterer for contract events
}

// CCSServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type CCSServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CCSServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CCSServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CCSServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CCSServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CCSServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CCSServiceSession struct {
	Contract     *CCSService       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CCSServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CCSServiceCallerSession struct {
	Contract *CCSServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CCSServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CCSServiceTransactorSession struct {
	Contract     *CCSServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CCSServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type CCSServiceRaw struct {
	Contract *CCSService // Generic contract binding to access the raw methods on
}

// CCSServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CCSServiceCallerRaw struct {
	Contract *CCSServiceCaller // Generic read-only contract binding to access the raw methods on
}

// CCSServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CCSServiceTransactorRaw struct {
	Contract *CCSServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCCSService creates a new instance of CCSService, bound to a specific deployed contract.
func NewCCSService(address common.Address, backend bind.ContractBackend) (*CCSService, error) {
	contract, err := bindCCSService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCSService{CCSServiceCaller: CCSServiceCaller{contract: contract}, CCSServiceTransactor: CCSServiceTransactor{contract: contract}, CCSServiceFilterer: CCSServiceFilterer{contract: contract}}, nil
}

// NewCCSServiceCaller creates a new read-only instance of CCSService, bound to a specific deployed contract.
func NewCCSServiceCaller(address common.Address, caller bind.ContractCaller) (*CCSServiceCaller, error) {
	contract, err := bindCCSService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCSServiceCaller{contract: contract}, nil
}

// NewCCSServiceTransactor creates a new write-only instance of CCSService, bound to a specific deployed contract.
func NewCCSServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*CCSServiceTransactor, error) {
	contract, err := bindCCSService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCSServiceTransactor{contract: contract}, nil
}

// NewCCSServiceFilterer creates a new log filterer instance of CCSService, bound to a specific deployed contract.
func NewCCSServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*CCSServiceFilterer, error) {
	contract, err := bindCCSService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCSServiceFilterer{contract: contract}, nil
}

// bindCCSService binds a generic wrapper to an already deployed contract.
func bindCCSService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CCSServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CCSService *CCSServiceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CCSService.Contract.CCSServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CCSService *CCSServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCSService.Contract.CCSServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CCSService *CCSServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCSService.Contract.CCSServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CCSService *CCSServiceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CCSService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CCSService *CCSServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCSService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CCSService *CCSServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCSService.Contract.contract.Transact(opts, method, params...)
}

// GetCCS is a free data retrieval call binding the contract method 0x339fcb3f.
//
// Solidity: function getCCS(confId string) constant returns(address, string, int256)
func (_CCSService *CCSServiceCaller) GetCCS(opts *bind.CallOpts, confId string) (common.Address, string, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _CCSService.contract.Call(opts, out, "getCCS", confId)
	return *ret0, *ret1, *ret2, err
}

// GetCCS is a free data retrieval call binding the contract method 0x339fcb3f.
//
// Solidity: function getCCS(confId string) constant returns(address, string, int256)
func (_CCSService *CCSServiceSession) GetCCS(confId string) (common.Address, string, *big.Int, error) {
	return _CCSService.Contract.GetCCS(&_CCSService.CallOpts, confId)
}

// GetCCS is a free data retrieval call binding the contract method 0x339fcb3f.
//
// Solidity: function getCCS(confId string) constant returns(address, string, int256)
func (_CCSService *CCSServiceCallerSession) GetCCS(confId string) (common.Address, string, *big.Int, error) {
	return _CCSService.Contract.GetCCS(&_CCSService.CallOpts, confId)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() constant returns(address)
func (_CCSService *CCSServiceCaller) ServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CCSService.contract.Call(opts, out, "serviceManager")
	return *ret0, err
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() constant returns(address)
func (_CCSService *CCSServiceSession) ServiceManager() (common.Address, error) {
	return _CCSService.Contract.ServiceManager(&_CCSService.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() constant returns(address)
func (_CCSService *CCSServiceCallerSession) ServiceManager() (common.Address, error) {
	return _CCSService.Contract.ServiceManager(&_CCSService.CallOpts)
}

// NewJob is a paid mutator transaction binding the contract method 0xefe4c5cc.
//
// Solidity: function newJob(confId string, ccsAddress address) returns()
func (_CCSService *CCSServiceTransactor) NewJob(opts *bind.TransactOpts, confId string, ccsAddress common.Address) (*types.Transaction, error) {
	return _CCSService.contract.Transact(opts, "newJob", confId, ccsAddress)
}

// NewJob is a paid mutator transaction binding the contract method 0xefe4c5cc.
//
// Solidity: function newJob(confId string, ccsAddress address) returns()
func (_CCSService *CCSServiceSession) NewJob(confId string, ccsAddress common.Address) (*types.Transaction, error) {
	return _CCSService.Contract.NewJob(&_CCSService.TransactOpts, confId, ccsAddress)
}

// NewJob is a paid mutator transaction binding the contract method 0xefe4c5cc.
//
// Solidity: function newJob(confId string, ccsAddress address) returns()
func (_CCSService *CCSServiceTransactorSession) NewJob(confId string, ccsAddress common.Address) (*types.Transaction, error) {
	return _CCSService.Contract.NewJob(&_CCSService.TransactOpts, confId, ccsAddress)
}

// RegisterCCS is a paid mutator transaction binding the contract method 0xfc14ce50.
//
// Solidity: function registerCCS(ip string, port int256) returns()
func (_CCSService *CCSServiceTransactor) RegisterCCS(opts *bind.TransactOpts, ip string, port *big.Int) (*types.Transaction, error) {
	return _CCSService.contract.Transact(opts, "registerCCS", ip, port)
}

// RegisterCCS is a paid mutator transaction binding the contract method 0xfc14ce50.
//
// Solidity: function registerCCS(ip string, port int256) returns()
func (_CCSService *CCSServiceSession) RegisterCCS(ip string, port *big.Int) (*types.Transaction, error) {
	return _CCSService.Contract.RegisterCCS(&_CCSService.TransactOpts, ip, port)
}

// RegisterCCS is a paid mutator transaction binding the contract method 0xfc14ce50.
//
// Solidity: function registerCCS(ip string, port int256) returns()
func (_CCSService *CCSServiceTransactorSession) RegisterCCS(ip string, port *big.Int) (*types.Transaction, error) {
	return _CCSService.Contract.RegisterCCS(&_CCSService.TransactOpts, ip, port)
}

// CCSServiceNewJobCreatedIterator is returned from FilterNewJobCreated and is used to iterate over the raw logs and unpacked data for NewJobCreated events raised by the CCSService contract.
type CCSServiceNewJobCreatedIterator struct {
	Event *CCSServiceNewJobCreated // Event containing the contract specifics and raw log

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
func (it *CCSServiceNewJobCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCSServiceNewJobCreated)
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
		it.Event = new(CCSServiceNewJobCreated)
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
func (it *CCSServiceNewJobCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CCSServiceNewJobCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CCSServiceNewJobCreated represents a NewJobCreated event raised by the CCSService contract.
type CCSServiceNewJobCreated struct {
	ConfId     string
	CcsAddress common.Address
	Ip         string
	Port       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewJobCreated is a free log retrieval operation binding the contract event 0x8fdb8b620029e6ca95a9d02e67bbc42a839ad5c1500657de42e4561d4bbf3349.
//
// Solidity: e NewJobCreated(confId string, ccsAddress address, ip string, port int256)
func (_CCSService *CCSServiceFilterer) FilterNewJobCreated(opts *bind.FilterOpts) (*CCSServiceNewJobCreatedIterator, error) {

	logs, sub, err := _CCSService.contract.FilterLogs(opts, "NewJobCreated")
	if err != nil {
		return nil, err
	}
	return &CCSServiceNewJobCreatedIterator{contract: _CCSService.contract, event: "NewJobCreated", logs: logs, sub: sub}, nil
}

// WatchNewJobCreated is a free log subscription operation binding the contract event 0x8fdb8b620029e6ca95a9d02e67bbc42a839ad5c1500657de42e4561d4bbf3349.
//
// Solidity: e NewJobCreated(confId string, ccsAddress address, ip string, port int256)
func (_CCSService *CCSServiceFilterer) WatchNewJobCreated(opts *bind.WatchOpts, sink chan<- *CCSServiceNewJobCreated) (event.Subscription, error) {

	logs, sub, err := _CCSService.contract.WatchLogs(opts, "NewJobCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CCSServiceNewJobCreated)
				if err := _CCSService.contract.UnpackLog(event, "NewJobCreated", log); err != nil {
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
