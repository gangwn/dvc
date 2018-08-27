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

// WebexMeetingABI is the input ABI used to generate the binding from.
const WebexMeetingABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getMeetingKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getMeetingSite\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"startMeeting\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"meetingtopic\",\"type\":\"string\"},{\"name\":\"starttime\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"scheduleMeeting\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mMeetingKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"confId\",\"type\":\"string\"}],\"name\":\"MeetingStarted\",\"type\":\"event\"}]"

// WebexMeetingBin is the compiled bytecode used for deploying new contracts.
const WebexMeetingBin = `0x608060405234801561001057600080fd5b506106ce806100206000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416635d3033eb81146100715780636742a096146100fb578063b0161e6f14610110578063bbda530014610125578063cd1624a614610187575b600080fd5b34801561007d57600080fd5b5061008661019c565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c05781810151838201526020016100a8565b50505050905090810190601f1680156100ed5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561010757600080fd5b50610086610233565b34801561011c57600080fd5b50610086610293565b34801561013157600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100869436949293602493928401919081908401838280828437509497505084359550505060209092013591506102ca9050565b34801561019357600080fd5b5061008661055b565b60008054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156102285780601f106101fd57610100808354040283529160200191610228565b820191906000526020600020905b81548152906001019060200180831161020b57829003601f168201915b505050505090505b90565b60018054604080516020601f600260001961010087891615020190951694909404938401819004810282018101909252828152606093909290918301828280156102285780601f106101fd57610100808354040283529160200191610228565b60408051808201909152600d81527f5374617274204d656574696e6700000000000000000000000000000000000000602082015290565b60606102d5846105e9565b1561034157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f6d656574696e6720746f706963206973206e756c6c0000000000000000000000604482015290519081900360640190fd5b835161035490600290602087019061060a565b50600082116103c457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f6d656574696e67206475726174696f6e206973206d757374203e203000000000604482015290519081900360640190fd5b60038290556000831161043857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f6d656574696e6720737461727474696d65206973206d757374203e2030000000604482015290519081900360640190fd5b60048390556040805180820190915260098082527f333438333937353134000000000000000000000000000000000000000000000060209092019182526104819160009161060a565b5060408051808201909152600c8082527f676f2e77656265782e636f6d000000000000000000000000000000000000000060209092019182526104c69160019161060a565b506000805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561054d5780601f106105225761010080835404028352916020019161054d565b820191906000526020600020905b81548152906001019060200180831161053057829003601f168201915b505050505090509392505050565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105e15780601f106105b6576101008083540402835291602001916105e1565b820191906000526020600020905b8154815290600101906020018083116105c457829003601f168201915b505050505081565b8051600090829015156105ff5760019150610604565b600091505b50919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061064b57805160ff1916838001178555610678565b82800160010185558215610678579182015b8281111561067857825182559160200191906001019061065d565b50610684929150610688565b5090565b61023091905b80821115610684576000815560010161068e5600a165627a7a7230582041cf4ba8d1599dad4f412f249d0a8f658a3bbbb955968b6c579aced4a63b70e50029`

// DeployWebexMeeting deploys a new Ethereum contract, binding an instance of WebexMeeting to it.
func DeployWebexMeeting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WebexMeeting, error) {
	parsed, err := abi.JSON(strings.NewReader(WebexMeetingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WebexMeetingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WebexMeeting{WebexMeetingCaller: WebexMeetingCaller{contract: contract}, WebexMeetingTransactor: WebexMeetingTransactor{contract: contract}, WebexMeetingFilterer: WebexMeetingFilterer{contract: contract}}, nil
}

// WebexMeeting is an auto generated Go binding around an Ethereum contract.
type WebexMeeting struct {
	WebexMeetingCaller     // Read-only binding to the contract
	WebexMeetingTransactor // Write-only binding to the contract
	WebexMeetingFilterer   // Log filterer for contract events
}

// WebexMeetingCaller is an auto generated read-only Go binding around an Ethereum contract.
type WebexMeetingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WebexMeetingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WebexMeetingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WebexMeetingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WebexMeetingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WebexMeetingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WebexMeetingSession struct {
	Contract     *WebexMeeting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WebexMeetingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WebexMeetingCallerSession struct {
	Contract *WebexMeetingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WebexMeetingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WebexMeetingTransactorSession struct {
	Contract     *WebexMeetingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WebexMeetingRaw is an auto generated low-level Go binding around an Ethereum contract.
type WebexMeetingRaw struct {
	Contract *WebexMeeting // Generic contract binding to access the raw methods on
}

// WebexMeetingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WebexMeetingCallerRaw struct {
	Contract *WebexMeetingCaller // Generic read-only contract binding to access the raw methods on
}

// WebexMeetingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WebexMeetingTransactorRaw struct {
	Contract *WebexMeetingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWebexMeeting creates a new instance of WebexMeeting, bound to a specific deployed contract.
func NewWebexMeeting(address common.Address, backend bind.ContractBackend) (*WebexMeeting, error) {
	contract, err := bindWebexMeeting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WebexMeeting{WebexMeetingCaller: WebexMeetingCaller{contract: contract}, WebexMeetingTransactor: WebexMeetingTransactor{contract: contract}, WebexMeetingFilterer: WebexMeetingFilterer{contract: contract}}, nil
}

// NewWebexMeetingCaller creates a new read-only instance of WebexMeeting, bound to a specific deployed contract.
func NewWebexMeetingCaller(address common.Address, caller bind.ContractCaller) (*WebexMeetingCaller, error) {
	contract, err := bindWebexMeeting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WebexMeetingCaller{contract: contract}, nil
}

// NewWebexMeetingTransactor creates a new write-only instance of WebexMeeting, bound to a specific deployed contract.
func NewWebexMeetingTransactor(address common.Address, transactor bind.ContractTransactor) (*WebexMeetingTransactor, error) {
	contract, err := bindWebexMeeting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WebexMeetingTransactor{contract: contract}, nil
}

// NewWebexMeetingFilterer creates a new log filterer instance of WebexMeeting, bound to a specific deployed contract.
func NewWebexMeetingFilterer(address common.Address, filterer bind.ContractFilterer) (*WebexMeetingFilterer, error) {
	contract, err := bindWebexMeeting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WebexMeetingFilterer{contract: contract}, nil
}

// bindWebexMeeting binds a generic wrapper to an already deployed contract.
func bindWebexMeeting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WebexMeetingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WebexMeeting *WebexMeetingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WebexMeeting.Contract.WebexMeetingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WebexMeeting *WebexMeetingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WebexMeeting.Contract.WebexMeetingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WebexMeeting *WebexMeetingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WebexMeeting.Contract.WebexMeetingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WebexMeeting *WebexMeetingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WebexMeeting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WebexMeeting *WebexMeetingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WebexMeeting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WebexMeeting *WebexMeetingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WebexMeeting.Contract.contract.Transact(opts, method, params...)
}

// MMeetingKey is a free data retrieval call binding the contract method 0xcd1624a6.
//
// Solidity: function mMeetingKey() constant returns(string)
func (_WebexMeeting *WebexMeetingCaller) MMeetingKey(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WebexMeeting.contract.Call(opts, out, "mMeetingKey")
	return *ret0, err
}

// MMeetingKey is a free data retrieval call binding the contract method 0xcd1624a6.
//
// Solidity: function mMeetingKey() constant returns(string)
func (_WebexMeeting *WebexMeetingSession) MMeetingKey() (string, error) {
	return _WebexMeeting.Contract.MMeetingKey(&_WebexMeeting.CallOpts)
}

// MMeetingKey is a free data retrieval call binding the contract method 0xcd1624a6.
//
// Solidity: function mMeetingKey() constant returns(string)
func (_WebexMeeting *WebexMeetingCallerSession) MMeetingKey() (string, error) {
	return _WebexMeeting.Contract.MMeetingKey(&_WebexMeeting.CallOpts)
}

// GetMeetingKey is a paid mutator transaction binding the contract method 0x5d3033eb.
//
// Solidity: function getMeetingKey() returns(string)
func (_WebexMeeting *WebexMeetingTransactor) GetMeetingKey(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WebexMeeting.contract.Transact(opts, "getMeetingKey")
}

// GetMeetingKey is a paid mutator transaction binding the contract method 0x5d3033eb.
//
// Solidity: function getMeetingKey() returns(string)
func (_WebexMeeting *WebexMeetingSession) GetMeetingKey() (*types.Transaction, error) {
	return _WebexMeeting.Contract.GetMeetingKey(&_WebexMeeting.TransactOpts)
}

// GetMeetingKey is a paid mutator transaction binding the contract method 0x5d3033eb.
//
// Solidity: function getMeetingKey() returns(string)
func (_WebexMeeting *WebexMeetingTransactorSession) GetMeetingKey() (*types.Transaction, error) {
	return _WebexMeeting.Contract.GetMeetingKey(&_WebexMeeting.TransactOpts)
}

// GetMeetingSite is a paid mutator transaction binding the contract method 0x6742a096.
//
// Solidity: function getMeetingSite() returns(string)
func (_WebexMeeting *WebexMeetingTransactor) GetMeetingSite(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WebexMeeting.contract.Transact(opts, "getMeetingSite")
}

// GetMeetingSite is a paid mutator transaction binding the contract method 0x6742a096.
//
// Solidity: function getMeetingSite() returns(string)
func (_WebexMeeting *WebexMeetingSession) GetMeetingSite() (*types.Transaction, error) {
	return _WebexMeeting.Contract.GetMeetingSite(&_WebexMeeting.TransactOpts)
}

// GetMeetingSite is a paid mutator transaction binding the contract method 0x6742a096.
//
// Solidity: function getMeetingSite() returns(string)
func (_WebexMeeting *WebexMeetingTransactorSession) GetMeetingSite() (*types.Transaction, error) {
	return _WebexMeeting.Contract.GetMeetingSite(&_WebexMeeting.TransactOpts)
}

// ScheduleMeeting is a paid mutator transaction binding the contract method 0xbbda5300.
//
// Solidity: function scheduleMeeting(meetingtopic string, starttime uint256, duration uint256) returns(string)
func (_WebexMeeting *WebexMeetingTransactor) ScheduleMeeting(opts *bind.TransactOpts, meetingtopic string, starttime *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _WebexMeeting.contract.Transact(opts, "scheduleMeeting", meetingtopic, starttime, duration)
}

// ScheduleMeeting is a paid mutator transaction binding the contract method 0xbbda5300.
//
// Solidity: function scheduleMeeting(meetingtopic string, starttime uint256, duration uint256) returns(string)
func (_WebexMeeting *WebexMeetingSession) ScheduleMeeting(meetingtopic string, starttime *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _WebexMeeting.Contract.ScheduleMeeting(&_WebexMeeting.TransactOpts, meetingtopic, starttime, duration)
}

// ScheduleMeeting is a paid mutator transaction binding the contract method 0xbbda5300.
//
// Solidity: function scheduleMeeting(meetingtopic string, starttime uint256, duration uint256) returns(string)
func (_WebexMeeting *WebexMeetingTransactorSession) ScheduleMeeting(meetingtopic string, starttime *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _WebexMeeting.Contract.ScheduleMeeting(&_WebexMeeting.TransactOpts, meetingtopic, starttime, duration)
}

// StartMeeting is a paid mutator transaction binding the contract method 0xb0161e6f.
//
// Solidity: function startMeeting() returns(string)
func (_WebexMeeting *WebexMeetingTransactor) StartMeeting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WebexMeeting.contract.Transact(opts, "startMeeting")
}

// StartMeeting is a paid mutator transaction binding the contract method 0xb0161e6f.
//
// Solidity: function startMeeting() returns(string)
func (_WebexMeeting *WebexMeetingSession) StartMeeting() (*types.Transaction, error) {
	return _WebexMeeting.Contract.StartMeeting(&_WebexMeeting.TransactOpts)
}

// StartMeeting is a paid mutator transaction binding the contract method 0xb0161e6f.
//
// Solidity: function startMeeting() returns(string)
func (_WebexMeeting *WebexMeetingTransactorSession) StartMeeting() (*types.Transaction, error) {
	return _WebexMeeting.Contract.StartMeeting(&_WebexMeeting.TransactOpts)
}

// WebexMeetingMeetingStartedIterator is returned from FilterMeetingStarted and is used to iterate over the raw logs and unpacked data for MeetingStarted events raised by the WebexMeeting contract.
type WebexMeetingMeetingStartedIterator struct {
	Event *WebexMeetingMeetingStarted // Event containing the contract specifics and raw log

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
func (it *WebexMeetingMeetingStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WebexMeetingMeetingStarted)
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
		it.Event = new(WebexMeetingMeetingStarted)
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
func (it *WebexMeetingMeetingStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WebexMeetingMeetingStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WebexMeetingMeetingStarted represents a MeetingStarted event raised by the WebexMeeting contract.
type WebexMeetingMeetingStarted struct {
	ConfId string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMeetingStarted is a free log retrieval operation binding the contract event 0x4aa2cd46efd0afd0b0b844ae5e7ac5b3baea9976ee7985a912e0606661281116.
//
// Solidity: e MeetingStarted(confId string)
func (_WebexMeeting *WebexMeetingFilterer) FilterMeetingStarted(opts *bind.FilterOpts) (*WebexMeetingMeetingStartedIterator, error) {

	logs, sub, err := _WebexMeeting.contract.FilterLogs(opts, "MeetingStarted")
	if err != nil {
		return nil, err
	}
	return &WebexMeetingMeetingStartedIterator{contract: _WebexMeeting.contract, event: "MeetingStarted", logs: logs, sub: sub}, nil
}

// WatchMeetingStarted is a free log subscription operation binding the contract event 0x4aa2cd46efd0afd0b0b844ae5e7ac5b3baea9976ee7985a912e0606661281116.
//
// Solidity: e MeetingStarted(confId string)
func (_WebexMeeting *WebexMeetingFilterer) WatchMeetingStarted(opts *bind.WatchOpts, sink chan<- *WebexMeetingMeetingStarted) (event.Subscription, error) {

	logs, sub, err := _WebexMeeting.contract.WatchLogs(opts, "MeetingStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WebexMeetingMeetingStarted)
				if err := _WebexMeeting.contract.UnpackLog(event, "MeetingStarted", log); err != nil {
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
