// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bepop

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

// BepopMetaData contains all meta data concerning the Bepop contract.
var BepopMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BepopABI is the input ABI used to generate the binding from.
// Deprecated: Use BepopMetaData.ABI instead.
var BepopABI = BepopMetaData.ABI

// Bepop is an auto generated Go binding around an Ethereum contract.
type Bepop struct {
	BepopCaller     // Read-only binding to the contract
	BepopTransactor // Write-only binding to the contract
	BepopFilterer   // Log filterer for contract events
}

// BepopCaller is an auto generated read-only Go binding around an Ethereum contract.
type BepopCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BepopTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BepopTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BepopFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BepopFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BepopSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BepopSession struct {
	Contract     *Bepop            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BepopCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BepopCallerSession struct {
	Contract *BepopCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BepopTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BepopTransactorSession struct {
	Contract     *BepopTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BepopRaw is an auto generated low-level Go binding around an Ethereum contract.
type BepopRaw struct {
	Contract *Bepop // Generic contract binding to access the raw methods on
}

// BepopCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BepopCallerRaw struct {
	Contract *BepopCaller // Generic read-only contract binding to access the raw methods on
}

// BepopTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BepopTransactorRaw struct {
	Contract *BepopTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBepop creates a new instance of Bepop, bound to a specific deployed contract.
func NewBepop(address common.Address, backend bind.ContractBackend) (*Bepop, error) {
	contract, err := bindBepop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bepop{BepopCaller: BepopCaller{contract: contract}, BepopTransactor: BepopTransactor{contract: contract}, BepopFilterer: BepopFilterer{contract: contract}}, nil
}

// NewBepopCaller creates a new read-only instance of Bepop, bound to a specific deployed contract.
func NewBepopCaller(address common.Address, caller bind.ContractCaller) (*BepopCaller, error) {
	contract, err := bindBepop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BepopCaller{contract: contract}, nil
}

// NewBepopTransactor creates a new write-only instance of Bepop, bound to a specific deployed contract.
func NewBepopTransactor(address common.Address, transactor bind.ContractTransactor) (*BepopTransactor, error) {
	contract, err := bindBepop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BepopTransactor{contract: contract}, nil
}

// NewBepopFilterer creates a new log filterer instance of Bepop, bound to a specific deployed contract.
func NewBepopFilterer(address common.Address, filterer bind.ContractFilterer) (*BepopFilterer, error) {
	contract, err := bindBepop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BepopFilterer{contract: contract}, nil
}

// bindBepop binds a generic wrapper to an already deployed contract.
func bindBepop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BepopMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bepop *BepopRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bepop.Contract.BepopCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bepop *BepopRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bepop.Contract.BepopTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bepop *BepopRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bepop.Contract.BepopTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bepop *BepopCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bepop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bepop *BepopTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bepop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bepop *BepopTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bepop.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bepop *BepopCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bepop.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bepop *BepopSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bepop.Contract.BalanceOf(&_Bepop.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bepop *BepopCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bepop.Contract.BalanceOf(&_Bepop.CallOpts, owner)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Bepop *BepopTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bepop.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Bepop *BepopSession) Deposit() (*types.Transaction, error) {
	return _Bepop.Contract.Deposit(&_Bepop.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Bepop *BepopTransactorSession) Deposit() (*types.Transaction, error) {
	return _Bepop.Contract.Deposit(&_Bepop.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_Bepop *BepopTransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Bepop.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_Bepop *BepopSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _Bepop.Contract.Withdraw(&_Bepop.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_Bepop *BepopTransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _Bepop.Contract.Withdraw(&_Bepop.TransactOpts, wad)
}
