// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teko

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

// TekoMetaData contains all meta data concerning the Teko contract.
var TekoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TekoABI is the input ABI used to generate the binding from.
// Deprecated: Use TekoMetaData.ABI instead.
var TekoABI = TekoMetaData.ABI

// Teko is an auto generated Go binding around an Ethereum contract.
type Teko struct {
	TekoCaller     // Read-only binding to the contract
	TekoTransactor // Write-only binding to the contract
	TekoFilterer   // Log filterer for contract events
}

// TekoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TekoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TekoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TekoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TekoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TekoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TekoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TekoSession struct {
	Contract     *Teko             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TekoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TekoCallerSession struct {
	Contract *TekoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TekoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TekoTransactorSession struct {
	Contract     *TekoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TekoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TekoRaw struct {
	Contract *Teko // Generic contract binding to access the raw methods on
}

// TekoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TekoCallerRaw struct {
	Contract *TekoCaller // Generic read-only contract binding to access the raw methods on
}

// TekoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TekoTransactorRaw struct {
	Contract *TekoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeko creates a new instance of Teko, bound to a specific deployed contract.
func NewTeko(address common.Address, backend bind.ContractBackend) (*Teko, error) {
	contract, err := bindTeko(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Teko{TekoCaller: TekoCaller{contract: contract}, TekoTransactor: TekoTransactor{contract: contract}, TekoFilterer: TekoFilterer{contract: contract}}, nil
}

// NewTekoCaller creates a new read-only instance of Teko, bound to a specific deployed contract.
func NewTekoCaller(address common.Address, caller bind.ContractCaller) (*TekoCaller, error) {
	contract, err := bindTeko(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TekoCaller{contract: contract}, nil
}

// NewTekoTransactor creates a new write-only instance of Teko, bound to a specific deployed contract.
func NewTekoTransactor(address common.Address, transactor bind.ContractTransactor) (*TekoTransactor, error) {
	contract, err := bindTeko(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TekoTransactor{contract: contract}, nil
}

// NewTekoFilterer creates a new log filterer instance of Teko, bound to a specific deployed contract.
func NewTekoFilterer(address common.Address, filterer bind.ContractFilterer) (*TekoFilterer, error) {
	contract, err := bindTeko(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TekoFilterer{contract: contract}, nil
}

// bindTeko binds a generic wrapper to an already deployed contract.
func bindTeko(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TekoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Teko *TekoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Teko.Contract.TekoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Teko *TekoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Teko.Contract.TekoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Teko *TekoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Teko.Contract.TekoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Teko *TekoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Teko.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Teko *TekoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Teko.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Teko *TekoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Teko.Contract.contract.Transact(opts, method, params...)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Teko *TekoTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Teko.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Teko *TekoSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Teko.Contract.Mint(&_Teko.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Teko *TekoTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Teko.Contract.Mint(&_Teko.TransactOpts, to, amount)
}
