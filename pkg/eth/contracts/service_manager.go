// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// ServiceManagerABI is the input ABI used to generate the binding from.
const ServiceManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"serviceName\",\"type\":\"string\"},{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"registerService\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"getService\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ServiceManager is an auto generated Go binding around an Ethereum contract.
type ServiceManager struct {
	ServiceManagerCaller     // Read-only binding to the contract
	ServiceManagerTransactor // Write-only binding to the contract
	ServiceManagerFilterer   // Log filterer for contract events
}

// ServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ServiceManagerSession struct {
	Contract     *ServiceManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ServiceManagerCallerSession struct {
	Contract *ServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ServiceManagerTransactorSession struct {
	Contract     *ServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ServiceManagerRaw struct {
	Contract *ServiceManager // Generic contract binding to access the raw methods on
}

// ServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ServiceManagerCallerRaw struct {
	Contract *ServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ServiceManagerTransactorRaw struct {
	Contract *ServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewServiceManager creates a new instance of ServiceManager, bound to a specific deployed contract.
func NewServiceManager(address common.Address, backend bind.ContractBackend) (*ServiceManager, error) {
	contract, err := bindServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ServiceManager{ServiceManagerCaller: ServiceManagerCaller{contract: contract}, ServiceManagerTransactor: ServiceManagerTransactor{contract: contract}, ServiceManagerFilterer: ServiceManagerFilterer{contract: contract}}, nil
}

// NewServiceManagerCaller creates a new read-only instance of ServiceManager, bound to a specific deployed contract.
func NewServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ServiceManagerCaller, error) {
	contract, err := bindServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerCaller{contract: contract}, nil
}

// NewServiceManagerTransactor creates a new write-only instance of ServiceManager, bound to a specific deployed contract.
func NewServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ServiceManagerTransactor, error) {
	contract, err := bindServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerTransactor{contract: contract}, nil
}

// NewServiceManagerFilterer creates a new log filterer instance of ServiceManager, bound to a specific deployed contract.
func NewServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ServiceManagerFilterer, error) {
	contract, err := bindServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerFilterer{contract: contract}, nil
}

// bindServiceManager binds a generic wrapper to an already deployed contract.
func bindServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ServiceManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServiceManager *ServiceManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ServiceManager.Contract.ServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServiceManager *ServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceManager.Contract.ServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServiceManager *ServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServiceManager.Contract.ServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServiceManager *ServiceManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServiceManager *ServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServiceManager *ServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServiceManager.Contract.contract.Transact(opts, method, params...)
}

// GetService is a free data retrieval call binding the contract method 0x794758be.
//
// Solidity: function getService(serviceName string) constant returns(address)
func (_ServiceManager *ServiceManagerCaller) GetService(opts *bind.CallOpts, serviceName string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ServiceManager.contract.Call(opts, out, "getService", serviceName)
	return *ret0, err
}

// GetService is a free data retrieval call binding the contract method 0x794758be.
//
// Solidity: function getService(serviceName string) constant returns(address)
func (_ServiceManager *ServiceManagerSession) GetService(serviceName string) (common.Address, error) {
	return _ServiceManager.Contract.GetService(&_ServiceManager.CallOpts, serviceName)
}

// GetService is a free data retrieval call binding the contract method 0x794758be.
//
// Solidity: function getService(serviceName string) constant returns(address)
func (_ServiceManager *ServiceManagerCallerSession) GetService(serviceName string) (common.Address, error) {
	return _ServiceManager.Contract.GetService(&_ServiceManager.CallOpts, serviceName)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ServiceManager *ServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ServiceManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ServiceManager *ServiceManagerSession) Owner() (common.Address, error) {
	return _ServiceManager.Contract.Owner(&_ServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ServiceManager *ServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ServiceManager.Contract.Owner(&_ServiceManager.CallOpts)
}

// RegisterService is a paid mutator transaction binding the contract method 0x3573c000.
//
// Solidity: function registerService(serviceName string, contractAddress address) returns()
func (_ServiceManager *ServiceManagerTransactor) RegisterService(opts *bind.TransactOpts, serviceName string, contractAddress common.Address) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "registerService", serviceName, contractAddress)
}

// RegisterService is a paid mutator transaction binding the contract method 0x3573c000.
//
// Solidity: function registerService(serviceName string, contractAddress address) returns()
func (_ServiceManager *ServiceManagerSession) RegisterService(serviceName string, contractAddress common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.RegisterService(&_ServiceManager.TransactOpts, serviceName, contractAddress)
}

// RegisterService is a paid mutator transaction binding the contract method 0x3573c000.
//
// Solidity: function registerService(serviceName string, contractAddress address) returns()
func (_ServiceManager *ServiceManagerTransactorSession) RegisterService(serviceName string, contractAddress common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.RegisterService(&_ServiceManager.TransactOpts, serviceName, contractAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ServiceManager *ServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ServiceManager *ServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ServiceManager.Contract.RenounceOwnership(&_ServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ServiceManager *ServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ServiceManager.Contract.RenounceOwnership(&_ServiceManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ServiceManager *ServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ServiceManager *ServiceManagerSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.TransferOwnership(&_ServiceManager.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ServiceManager *ServiceManagerTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.TransferOwnership(&_ServiceManager.TransactOpts, _newOwner)
}

// ServiceManagerOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the ServiceManager contract.
type ServiceManagerOwnershipRenouncedIterator struct {
	Event *ServiceManagerOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *ServiceManagerOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerOwnershipRenounced)
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
		it.Event = new(ServiceManagerOwnershipRenounced)
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
func (it *ServiceManagerOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerOwnershipRenounced represents a OwnershipRenounced event raised by the ServiceManager contract.
type ServiceManagerOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_ServiceManager *ServiceManagerFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ServiceManagerOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerOwnershipRenouncedIterator{contract: _ServiceManager.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_ServiceManager *ServiceManagerFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ServiceManagerOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerOwnershipRenounced)
				if err := _ServiceManager.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ServiceManager contract.
type ServiceManagerOwnershipTransferredIterator struct {
	Event *ServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerOwnershipTransferred)
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
		it.Event = new(ServiceManagerOwnershipTransferred)
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
func (it *ServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ServiceManager contract.
type ServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ServiceManager *ServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerOwnershipTransferredIterator{contract: _ServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ServiceManager *ServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerOwnershipTransferred)
				if err := _ServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
