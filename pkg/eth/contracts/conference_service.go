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

// ConferenceServiceABI is the input ABI used to generate the binding from.
const ConferenceServiceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"serviceManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"timeAfter\",\"type\":\"uint256\"}],\"name\":\"getConferences\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"confId\",\"type\":\"string\"}],\"name\":\"ttt\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"confId\",\"type\":\"string\"},{\"name\":\"topic\",\"type\":\"string\"},{\"name\":\"startTime\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"},{\"name\":\"invitees\",\"type\":\"address[]\"}],\"name\":\"scheduleConference\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"serviceManagerAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"creatorAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"confId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"topic\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"invitees\",\"type\":\"address[]\"}],\"name\":\"ConferenceScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"confId\",\"type\":\"string\"}],\"name\":\"Test\",\"type\":\"event\"}]"

// ConferenceService is an auto generated Go binding around an Ethereum contract.
type ConferenceService struct {
	ConferenceServiceCaller     // Read-only binding to the contract
	ConferenceServiceTransactor // Write-only binding to the contract
	ConferenceServiceFilterer   // Log filterer for contract events
}

// ConferenceServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConferenceServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConferenceServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConferenceServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConferenceServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConferenceServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConferenceServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConferenceServiceSession struct {
	Contract     *ConferenceService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConferenceServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConferenceServiceCallerSession struct {
	Contract *ConferenceServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ConferenceServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConferenceServiceTransactorSession struct {
	Contract     *ConferenceServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ConferenceServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConferenceServiceRaw struct {
	Contract *ConferenceService // Generic contract binding to access the raw methods on
}

// ConferenceServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConferenceServiceCallerRaw struct {
	Contract *ConferenceServiceCaller // Generic read-only contract binding to access the raw methods on
}

// ConferenceServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConferenceServiceTransactorRaw struct {
	Contract *ConferenceServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConferenceService creates a new instance of ConferenceService, bound to a specific deployed contract.
func NewConferenceService(address common.Address, backend bind.ContractBackend) (*ConferenceService, error) {
	contract, err := bindConferenceService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConferenceService{ConferenceServiceCaller: ConferenceServiceCaller{contract: contract}, ConferenceServiceTransactor: ConferenceServiceTransactor{contract: contract}, ConferenceServiceFilterer: ConferenceServiceFilterer{contract: contract}}, nil
}

// NewConferenceServiceCaller creates a new read-only instance of ConferenceService, bound to a specific deployed contract.
func NewConferenceServiceCaller(address common.Address, caller bind.ContractCaller) (*ConferenceServiceCaller, error) {
	contract, err := bindConferenceService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConferenceServiceCaller{contract: contract}, nil
}

// NewConferenceServiceTransactor creates a new write-only instance of ConferenceService, bound to a specific deployed contract.
func NewConferenceServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*ConferenceServiceTransactor, error) {
	contract, err := bindConferenceService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConferenceServiceTransactor{contract: contract}, nil
}

// NewConferenceServiceFilterer creates a new log filterer instance of ConferenceService, bound to a specific deployed contract.
func NewConferenceServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*ConferenceServiceFilterer, error) {
	contract, err := bindConferenceService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConferenceServiceFilterer{contract: contract}, nil
}

// bindConferenceService binds a generic wrapper to an already deployed contract.
func bindConferenceService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConferenceServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConferenceService *ConferenceServiceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ConferenceService.Contract.ConferenceServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConferenceService *ConferenceServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConferenceService.Contract.ConferenceServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConferenceService *ConferenceServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConferenceService.Contract.ConferenceServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConferenceService *ConferenceServiceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ConferenceService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConferenceService *ConferenceServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConferenceService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConferenceService *ConferenceServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConferenceService.Contract.contract.Transact(opts, method, params...)
}

// GetConferences is a free data retrieval call binding the contract method 0x5e1a496c.
//
// Solidity: function getConferences(timeAfter uint256) constant returns(string, address, string, uint256, uint256, uint256)
func (_ConferenceService *ConferenceServiceCaller) GetConferences(opts *bind.CallOpts, timeAfter *big.Int) (string, common.Address, string, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(common.Address)
		ret2 = new(string)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _ConferenceService.contract.Call(opts, out, "getConferences", timeAfter)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetConferences is a free data retrieval call binding the contract method 0x5e1a496c.
//
// Solidity: function getConferences(timeAfter uint256) constant returns(string, address, string, uint256, uint256, uint256)
func (_ConferenceService *ConferenceServiceSession) GetConferences(timeAfter *big.Int) (string, common.Address, string, *big.Int, *big.Int, *big.Int, error) {
	return _ConferenceService.Contract.GetConferences(&_ConferenceService.CallOpts, timeAfter)
}

// GetConferences is a free data retrieval call binding the contract method 0x5e1a496c.
//
// Solidity: function getConferences(timeAfter uint256) constant returns(string, address, string, uint256, uint256, uint256)
func (_ConferenceService *ConferenceServiceCallerSession) GetConferences(timeAfter *big.Int) (string, common.Address, string, *big.Int, *big.Int, *big.Int, error) {
	return _ConferenceService.Contract.GetConferences(&_ConferenceService.CallOpts, timeAfter)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() constant returns(address)
func (_ConferenceService *ConferenceServiceCaller) ServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ConferenceService.contract.Call(opts, out, "serviceManager")
	return *ret0, err
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() constant returns(address)
func (_ConferenceService *ConferenceServiceSession) ServiceManager() (common.Address, error) {
	return _ConferenceService.Contract.ServiceManager(&_ConferenceService.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() constant returns(address)
func (_ConferenceService *ConferenceServiceCallerSession) ServiceManager() (common.Address, error) {
	return _ConferenceService.Contract.ServiceManager(&_ConferenceService.CallOpts)
}

// Ttt is a free data retrieval call binding the contract method 0x88581c9a.
//
// Solidity: function ttt(confId string) constant returns(string)
func (_ConferenceService *ConferenceServiceCaller) Ttt(opts *bind.CallOpts, confId string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ConferenceService.contract.Call(opts, out, "ttt", confId)
	return *ret0, err
}

// Ttt is a free data retrieval call binding the contract method 0x88581c9a.
//
// Solidity: function ttt(confId string) constant returns(string)
func (_ConferenceService *ConferenceServiceSession) Ttt(confId string) (string, error) {
	return _ConferenceService.Contract.Ttt(&_ConferenceService.CallOpts, confId)
}

// Ttt is a free data retrieval call binding the contract method 0x88581c9a.
//
// Solidity: function ttt(confId string) constant returns(string)
func (_ConferenceService *ConferenceServiceCallerSession) Ttt(confId string) (string, error) {
	return _ConferenceService.Contract.Ttt(&_ConferenceService.CallOpts, confId)
}

// ScheduleConference is a paid mutator transaction binding the contract method 0xec35ea15.
//
// Solidity: function scheduleConference(confId string, topic string, startTime uint256, duration uint256, invitees address[]) returns()
func (_ConferenceService *ConferenceServiceTransactor) ScheduleConference(opts *bind.TransactOpts, confId string, topic string, startTime *big.Int, duration *big.Int, invitees []common.Address) (*types.Transaction, error) {
	return _ConferenceService.contract.Transact(opts, "scheduleConference", confId, topic, startTime, duration, invitees)
}

// ScheduleConference is a paid mutator transaction binding the contract method 0xec35ea15.
//
// Solidity: function scheduleConference(confId string, topic string, startTime uint256, duration uint256, invitees address[]) returns()
func (_ConferenceService *ConferenceServiceSession) ScheduleConference(confId string, topic string, startTime *big.Int, duration *big.Int, invitees []common.Address) (*types.Transaction, error) {
	return _ConferenceService.Contract.ScheduleConference(&_ConferenceService.TransactOpts, confId, topic, startTime, duration, invitees)
}

// ScheduleConference is a paid mutator transaction binding the contract method 0xec35ea15.
//
// Solidity: function scheduleConference(confId string, topic string, startTime uint256, duration uint256, invitees address[]) returns()
func (_ConferenceService *ConferenceServiceTransactorSession) ScheduleConference(confId string, topic string, startTime *big.Int, duration *big.Int, invitees []common.Address) (*types.Transaction, error) {
	return _ConferenceService.Contract.ScheduleConference(&_ConferenceService.TransactOpts, confId, topic, startTime, duration, invitees)
}

// ConferenceServiceConferenceScheduledIterator is returned from FilterConferenceScheduled and is used to iterate over the raw logs and unpacked data for ConferenceScheduled events raised by the ConferenceService contract.
type ConferenceServiceConferenceScheduledIterator struct {
	Event *ConferenceServiceConferenceScheduled // Event containing the contract specifics and raw log

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
func (it *ConferenceServiceConferenceScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConferenceServiceConferenceScheduled)
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
		it.Event = new(ConferenceServiceConferenceScheduled)
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
func (it *ConferenceServiceConferenceScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConferenceServiceConferenceScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConferenceServiceConferenceScheduled represents a ConferenceScheduled event raised by the ConferenceService contract.
type ConferenceServiceConferenceScheduled struct {
	CreatorAddr common.Address
	ConfId      string
	Topic       string
	StartTime   *big.Int
	Duration    *big.Int
	Invitees    []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConferenceScheduled is a free log retrieval operation binding the contract event 0x9687a96d2065c77c0926ca0180854e8adce339c57e9e6e9b5b7e0a2a5386a87f.
//
// Solidity: e ConferenceScheduled(creatorAddr address, confId string, topic string, startTime uint256, duration uint256, invitees address[])
func (_ConferenceService *ConferenceServiceFilterer) FilterConferenceScheduled(opts *bind.FilterOpts) (*ConferenceServiceConferenceScheduledIterator, error) {

	logs, sub, err := _ConferenceService.contract.FilterLogs(opts, "ConferenceScheduled")
	if err != nil {
		return nil, err
	}
	return &ConferenceServiceConferenceScheduledIterator{contract: _ConferenceService.contract, event: "ConferenceScheduled", logs: logs, sub: sub}, nil
}

// WatchConferenceScheduled is a free log subscription operation binding the contract event 0x9687a96d2065c77c0926ca0180854e8adce339c57e9e6e9b5b7e0a2a5386a87f.
//
// Solidity: e ConferenceScheduled(creatorAddr address, confId string, topic string, startTime uint256, duration uint256, invitees address[])
func (_ConferenceService *ConferenceServiceFilterer) WatchConferenceScheduled(opts *bind.WatchOpts, sink chan<- *ConferenceServiceConferenceScheduled) (event.Subscription, error) {

	logs, sub, err := _ConferenceService.contract.WatchLogs(opts, "ConferenceScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConferenceServiceConferenceScheduled)
				if err := _ConferenceService.contract.UnpackLog(event, "ConferenceScheduled", log); err != nil {
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

// ConferenceServiceTestIterator is returned from FilterTest and is used to iterate over the raw logs and unpacked data for Test events raised by the ConferenceService contract.
type ConferenceServiceTestIterator struct {
	Event *ConferenceServiceTest // Event containing the contract specifics and raw log

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
func (it *ConferenceServiceTestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConferenceServiceTest)
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
		it.Event = new(ConferenceServiceTest)
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
func (it *ConferenceServiceTestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConferenceServiceTestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConferenceServiceTest represents a Test event raised by the ConferenceService contract.
type ConferenceServiceTest struct {
	ConfId string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTest is a free log retrieval operation binding the contract event 0x00cb39d6c2c520f0597db0021367767c48fef2964cf402d3c9e9d4df12e43964.
//
// Solidity: e Test(confId string)
func (_ConferenceService *ConferenceServiceFilterer) FilterTest(opts *bind.FilterOpts) (*ConferenceServiceTestIterator, error) {

	logs, sub, err := _ConferenceService.contract.FilterLogs(opts, "Test")
	if err != nil {
		return nil, err
	}
	return &ConferenceServiceTestIterator{contract: _ConferenceService.contract, event: "Test", logs: logs, sub: sub}, nil
}

// WatchTest is a free log subscription operation binding the contract event 0x00cb39d6c2c520f0597db0021367767c48fef2964cf402d3c9e9d4df12e43964.
//
// Solidity: e Test(confId string)
func (_ConferenceService *ConferenceServiceFilterer) WatchTest(opts *bind.WatchOpts, sink chan<- *ConferenceServiceTest) (event.Subscription, error) {

	logs, sub, err := _ConferenceService.contract.WatchLogs(opts, "Test")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConferenceServiceTest)
				if err := _ConferenceService.contract.UnpackLog(event, "Test", log); err != nil {
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
