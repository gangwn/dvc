// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
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

// UserServiceABI is the input ABI used to generate the binding from.
const UserServiceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"timeAfter\",\"type\":\"uint256\"}],\"name\":\"getConferences\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"users\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"displayName\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UserServiceBin is the compiled bytecode used for deploying new contracts.
const UserServiceBin = `0x608060405234801561001057600080fd5b5061051a806100206000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416635e1a496c8114610050578063a87430ba14610174575b600080fd5b34801561005c57600080fd5b50610068600435610226565b604051808060200187600160a060020a0316600160a060020a0316815260200180602001868152602001858152602001848152602001838103835289818151815260200191508051906020019080838360005b838110156100d35781810151838201526020016100bb565b50505050905090810190601f1680156101005780820380516001836020036101000a031916815260200191505b50838103825287518152875160209182019189019080838360005b8381101561013357818101518382015260200161011b565b50505050905090810190601f1680156101605780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b34801561018057600080fd5b50610195600160a060020a036004351661043d565b6040518083600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101ea5781810151838201526020016101d2565b50505050905090810190601f1680156102175780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b33600090815260208190526040812060609190829082908190819081805b6002830154821015610404576002830180548390811061026057fe5b9060005260206000209060060201600301548a111561027e576103f9565b6002830180548390811061028e57fe5b6000918252602091829020600160069092020181810154600382015460048301546005840154845460408051601f60026000199a851615610100029a909a019093168990049283018a90048a0281018a019091528181529598508897600160a060020a03909516968801959394929391928891908301828280156103535780601f1061032857610100808354040283529160200191610353565b820191906000526020600020905b81548152906001019060200180831161033657829003601f168201915b5050875460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b50899450925084019050828280156103e15780601f106103b6576101008083540402835291602001916103e1565b820191906000526020600020905b8154815290600101906020018083116103c457829003601f168201915b50505050509350985098509850985098509850610431565b600190910190610244565b604080516020818101835260008083528351918201909352828152909a5090985096508795508594508493505b50505091939550919395565b600060208181529181526040908190208054600180830180548551600261010094831615949094026000190190911692909204601f8101879004870283018701909552848252600160a060020a0390921694929390928301828280156104e45780601f106104b9576101008083540402835291602001916104e4565b820191906000526020600020905b8154815290600101906020018083116104c757829003601f168201915b50505050509050825600a165627a7a72305820988f28f53f9a8234da4d9b831aee1f8b0407ad7c11429b7d7b40c2599a34b98e0029`

// DeployUserService deploys a new Ethereum contract, binding an instance of UserService to it.
func DeployUserService(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserService, error) {
	parsed, err := abi.JSON(strings.NewReader(UserServiceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserServiceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserService{UserServiceCaller: UserServiceCaller{contract: contract}, UserServiceTransactor: UserServiceTransactor{contract: contract}, UserServiceFilterer: UserServiceFilterer{contract: contract}}, nil
}

// UserService is an auto generated Go binding around an Ethereum contract.
type UserService struct {
	UserServiceCaller     // Read-only binding to the contract
	UserServiceTransactor // Write-only binding to the contract
	UserServiceFilterer   // Log filterer for contract events
}

// UserServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserServiceSession struct {
	Contract     *UserService      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserServiceCallerSession struct {
	Contract *UserServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// UserServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserServiceTransactorSession struct {
	Contract     *UserServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UserServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserServiceRaw struct {
	Contract *UserService // Generic contract binding to access the raw methods on
}

// UserServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserServiceCallerRaw struct {
	Contract *UserServiceCaller // Generic read-only contract binding to access the raw methods on
}

// UserServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserServiceTransactorRaw struct {
	Contract *UserServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserService creates a new instance of UserService, bound to a specific deployed contract.
func NewUserService(address common.Address, backend bind.ContractBackend) (*UserService, error) {
	contract, err := bindUserService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserService{UserServiceCaller: UserServiceCaller{contract: contract}, UserServiceTransactor: UserServiceTransactor{contract: contract}, UserServiceFilterer: UserServiceFilterer{contract: contract}}, nil
}

// NewUserServiceCaller creates a new read-only instance of UserService, bound to a specific deployed contract.
func NewUserServiceCaller(address common.Address, caller bind.ContractCaller) (*UserServiceCaller, error) {
	contract, err := bindUserService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserServiceCaller{contract: contract}, nil
}

// NewUserServiceTransactor creates a new write-only instance of UserService, bound to a specific deployed contract.
func NewUserServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*UserServiceTransactor, error) {
	contract, err := bindUserService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserServiceTransactor{contract: contract}, nil
}

// NewUserServiceFilterer creates a new log filterer instance of UserService, bound to a specific deployed contract.
func NewUserServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*UserServiceFilterer, error) {
	contract, err := bindUserService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserServiceFilterer{contract: contract}, nil
}

// bindUserService binds a generic wrapper to an already deployed contract.
func bindUserService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserService *UserServiceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserService.Contract.UserServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserService *UserServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserService.Contract.UserServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserService *UserServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserService.Contract.UserServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserService *UserServiceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserService *UserServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserService *UserServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserService.Contract.contract.Transact(opts, method, params...)
}

// GetConferences is a free data retrieval call binding the contract method 0x5e1a496c.
//
// Solidity: function getConferences(timeAfter uint256) constant returns(string, address, string, uint256, uint256, uint256)
func (_UserService *UserServiceCaller) GetConferences(opts *bind.CallOpts, timeAfter *big.Int) (string, common.Address, string, *big.Int, *big.Int, *big.Int, error) {
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
	err := _UserService.contract.Call(opts, out, "getConferences", timeAfter)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetConferences is a free data retrieval call binding the contract method 0x5e1a496c.
//
// Solidity: function getConferences(timeAfter uint256) constant returns(string, address, string, uint256, uint256, uint256)
func (_UserService *UserServiceSession) GetConferences(timeAfter *big.Int) (string, common.Address, string, *big.Int, *big.Int, *big.Int, error) {
	return _UserService.Contract.GetConferences(&_UserService.CallOpts, timeAfter)
}

// GetConferences is a free data retrieval call binding the contract method 0x5e1a496c.
//
// Solidity: function getConferences(timeAfter uint256) constant returns(string, address, string, uint256, uint256, uint256)
func (_UserService *UserServiceCallerSession) GetConferences(timeAfter *big.Int) (string, common.Address, string, *big.Int, *big.Int, *big.Int, error) {
	return _UserService.Contract.GetConferences(&_UserService.CallOpts, timeAfter)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users( address) constant returns(addr address, displayName string)
func (_UserService *UserServiceCaller) Users(opts *bind.CallOpts, arg0 common.Address) (struct {
	Addr        common.Address
	DisplayName string
}, error) {
	ret := new(struct {
		Addr        common.Address
		DisplayName string
	})
	out := ret
	err := _UserService.contract.Call(opts, out, "users", arg0)
	return *ret, err
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users( address) constant returns(addr address, displayName string)
func (_UserService *UserServiceSession) Users(arg0 common.Address) (struct {
	Addr        common.Address
	DisplayName string
}, error) {
	return _UserService.Contract.Users(&_UserService.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users( address) constant returns(addr address, displayName string)
func (_UserService *UserServiceCallerSession) Users(arg0 common.Address) (struct {
	Addr        common.Address
	DisplayName string
}, error) {
	return _UserService.Contract.Users(&_UserService.CallOpts, arg0)
}
