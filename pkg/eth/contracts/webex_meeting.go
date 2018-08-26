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
const WebexMeetingABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getMeetingKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getMeetingSite\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"string\"}],\"name\":\"getServiceURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startMeeting\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"meetingtopic\",\"type\":\"string\"},{\"name\":\"starttime\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"scheduleMeeting\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mMeetingKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"calculateSum\",\"outputs\":[{\"name\":\"sum\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"value\",\"type\":\"string\"}],\"name\":\"startMeetingEvent\",\"type\":\"event\"}]"

// WebexMeetingBin is the compiled bytecode used for deploying new contracts.
const WebexMeetingBin = `0x608060405234801561001057600080fd5b506108f5806100206000396000f3006080604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416635d3033eb81146100875780636742a09614610111578063675a0c1e14610126578063b0161e6f1461017f578063bbda530014610194578063cd1624a6146101f6578063ea4097bd1461020b575b600080fd5b34801561009357600080fd5b5061009c610238565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100d65781810151838201526020016100be565b50505050905090810190601f1680156101035780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561011d57600080fd5b5061009c6102cf565b34801561013257600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261009c94369492936024939284019190819084018382808284375094975061032f9650505050505050565b34801561018b57600080fd5b5061009c610423565b3480156101a057600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261009c9436949293602493928401919081908401838280828437509497505084359550505060209092013591506104f79050565b34801561020257600080fd5b5061009c610788565b34801561021757600080fd5b50610226600435602435610816565b60408051918252519081900360200190f35b60008054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156102c45780601f10610299576101008083540402835291602001916102c4565b820191906000526020600020905b8154815290600101906020018083116102a757829003601f168201915b505050505090505b90565b60018054604080516020601f600260001961010087891615020190951694909404938401819004810282018101909252828152606093909290918301828280156102c45780601f10610299576101008083540402835291602001916102c4565b60606005826040518082805190602001908083835b602083106103635780518252601f199092019160209182019101610344565b518151600019602094850361010090810a820192831692199390931691909117909252949092019687526040805197889003820188208054601f60026001831615909802909501169590950492830182900482028801820190528187529294509250508301828280156104175780601f106103ec57610100808354040283529160200191610417565b820191906000526020600020905b8154815290600101906020018083116103fa57829003601f168201915b50505050509050919050565b604080518082018252600581527f73746172740000000000000000000000000000000000000000000000000000006020808301918252835181815283519181019190915282516060947f917d34685b354a00269a44827ad3b6c1547987122a8623d84f0090686d3859289385939283928301919080838360005b838110156104b557818101518382015260200161049d565b50505050905090810190601f1680156104e25780820380516001836020036101000a031916815260200191505b509250505060405180910390a18091505b5090565b60606105028461081a565b1561056e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f6d656574696e6720746f706963206973206e756c6c0000000000000000000000604482015290519081900360640190fd5b835161058190600290602087019061083b565b50600082116105f157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f6d656574696e67206475726174696f6e206973206d757374203e203000000000604482015290519081900360640190fd5b60038290556000831161066557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f6d656574696e6720737461727474696d65206973206d757374203e2030000000604482015290519081900360640190fd5b60048390556040805180820190915260098082527f333438333937353134000000000000000000000000000000000000000000000060209092019182526106ae9160009161083b565b5060408051808201909152600c8082527f676f2e77656265782e636f6d000000000000000000000000000000000000000060209092019182526106f39160019161083b565b506000805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561077a5780601f1061074f5761010080835404028352916020019161077a565b820191906000526020600020905b81548152906001019060200180831161075d57829003601f168201915b505050505090509392505050565b6000805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561080e5780601f106107e35761010080835404028352916020019161080e565b820191906000526020600020905b8154815290600101906020018083116107f157829003601f168201915b505050505081565b0190565b8051600090829015156108305760019150610835565b600091505b50919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061087c57805160ff19168380011785556108a9565b828001600101855582156108a9579182015b828111156108a957825182559160200191906001019061088e565b506104f3926102cc9250905b808211156104f357600081556001016108b55600a165627a7a72305820050065636c896570abab003256fa994f6583793eee04a2ed36f1f03209183de90029`

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

// CalculateSum is a free data retrieval call binding the contract method 0xea4097bd.
//
// Solidity: function calculateSum(a uint256, b uint256) constant returns(sum uint256)
func (_WebexMeeting *WebexMeetingCaller) CalculateSum(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WebexMeeting.contract.Call(opts, out, "calculateSum", a, b)
	return *ret0, err
}

// CalculateSum is a free data retrieval call binding the contract method 0xea4097bd.
//
// Solidity: function calculateSum(a uint256, b uint256) constant returns(sum uint256)
func (_WebexMeeting *WebexMeetingSession) CalculateSum(a *big.Int, b *big.Int) (*big.Int, error) {
	return _WebexMeeting.Contract.CalculateSum(&_WebexMeeting.CallOpts, a, b)
}

// CalculateSum is a free data retrieval call binding the contract method 0xea4097bd.
//
// Solidity: function calculateSum(a uint256, b uint256) constant returns(sum uint256)
func (_WebexMeeting *WebexMeetingCallerSession) CalculateSum(a *big.Int, b *big.Int) (*big.Int, error) {
	return _WebexMeeting.Contract.CalculateSum(&_WebexMeeting.CallOpts, a, b)
}

// GetServiceURI is a free data retrieval call binding the contract method 0x675a0c1e.
//
// Solidity: function getServiceURI(_addr string) constant returns(string)
func (_WebexMeeting *WebexMeetingCaller) GetServiceURI(opts *bind.CallOpts, _addr string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WebexMeeting.contract.Call(opts, out, "getServiceURI", _addr)
	return *ret0, err
}

// GetServiceURI is a free data retrieval call binding the contract method 0x675a0c1e.
//
// Solidity: function getServiceURI(_addr string) constant returns(string)
func (_WebexMeeting *WebexMeetingSession) GetServiceURI(_addr string) (string, error) {
	return _WebexMeeting.Contract.GetServiceURI(&_WebexMeeting.CallOpts, _addr)
}

// GetServiceURI is a free data retrieval call binding the contract method 0x675a0c1e.
//
// Solidity: function getServiceURI(_addr string) constant returns(string)
func (_WebexMeeting *WebexMeetingCallerSession) GetServiceURI(_addr string) (string, error) {
	return _WebexMeeting.Contract.GetServiceURI(&_WebexMeeting.CallOpts, _addr)
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

// StartMeeting is a free data retrieval call binding the contract method 0xb0161e6f.
//
// Solidity: function startMeeting() constant returns(string)
func (_WebexMeeting *WebexMeetingCaller) StartMeeting(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WebexMeeting.contract.Call(opts, out, "startMeeting")
	return *ret0, err
}

// StartMeeting is a free data retrieval call binding the contract method 0xb0161e6f.
//
// Solidity: function startMeeting() constant returns(string)
func (_WebexMeeting *WebexMeetingSession) StartMeeting() (string, error) {
	return _WebexMeeting.Contract.StartMeeting(&_WebexMeeting.CallOpts)
}

// StartMeeting is a free data retrieval call binding the contract method 0xb0161e6f.
//
// Solidity: function startMeeting() constant returns(string)
func (_WebexMeeting *WebexMeetingCallerSession) StartMeeting() (string, error) {
	return _WebexMeeting.Contract.StartMeeting(&_WebexMeeting.CallOpts)
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

// WebexMeetingStartMeetingEventIterator is returned from FilterStartMeetingEvent and is used to iterate over the raw logs and unpacked data for StartMeetingEvent events raised by the WebexMeeting contract.
type WebexMeetingStartMeetingEventIterator struct {
	Event *WebexMeetingStartMeetingEvent // Event containing the contract specifics and raw log

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
func (it *WebexMeetingStartMeetingEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WebexMeetingStartMeetingEvent)
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
		it.Event = new(WebexMeetingStartMeetingEvent)
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
func (it *WebexMeetingStartMeetingEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WebexMeetingStartMeetingEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WebexMeetingStartMeetingEvent represents a StartMeetingEvent event raised by the WebexMeeting contract.
type WebexMeetingStartMeetingEvent struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStartMeetingEvent is a free log retrieval operation binding the contract event 0x917d34685b354a00269a44827ad3b6c1547987122a8623d84f0090686d385928.
//
// Solidity: e startMeetingEvent(value string)
func (_WebexMeeting *WebexMeetingFilterer) FilterStartMeetingEvent(opts *bind.FilterOpts) (*WebexMeetingStartMeetingEventIterator, error) {

	logs, sub, err := _WebexMeeting.contract.FilterLogs(opts, "startMeetingEvent")
	if err != nil {
		return nil, err
	}
	return &WebexMeetingStartMeetingEventIterator{contract: _WebexMeeting.contract, event: "startMeetingEvent", logs: logs, sub: sub}, nil
}

// WatchStartMeetingEvent is a free log subscription operation binding the contract event 0x917d34685b354a00269a44827ad3b6c1547987122a8623d84f0090686d385928.
//
// Solidity: e startMeetingEvent(value string)
func (_WebexMeeting *WebexMeetingFilterer) WatchStartMeetingEvent(opts *bind.WatchOpts, sink chan<- *WebexMeetingStartMeetingEvent) (event.Subscription, error) {

	logs, sub, err := _WebexMeeting.contract.WatchLogs(opts, "startMeetingEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WebexMeetingStartMeetingEvent)
				if err := _WebexMeeting.contract.UnpackLog(event, "startMeetingEvent", log); err != nil {
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
