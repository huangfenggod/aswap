// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

// DOFABI is the input ABI used to generate the binding from.
const DOFABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_tradeLock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"turnTradeLock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DOF is an auto generated Go binding around an Ethereum contract.
type DOF struct {
	DOFCaller     // Read-only binding to the contract
	DOFTransactor // Write-only binding to the contract
}

// DOFCaller is an auto generated read-only Go binding around an Ethereum contract.
type DOFCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOFTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DOFTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOFSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DOFSession struct {
	Contract     *DOF              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DOFCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DOFCallerSession struct {
	Contract *DOFCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DOFTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DOFTransactorSession struct {
	Contract     *DOFTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DOFRaw is an auto generated low-level Go binding around an Ethereum contract.
type DOFRaw struct {
	Contract *DOF // Generic contract binding to access the raw methods on
}

// DOFCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DOFCallerRaw struct {
	Contract *DOFCaller // Generic read-only contract binding to access the raw methods on
}

// DOFTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DOFTransactorRaw struct {
	Contract *DOFTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDOF creates a new instance of DOF, bound to a specific deployed contract.
func NewDOF(address common.Address, backend bind.ContractBackend) (*DOF, error) {
	contract, err := bindDOF(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DOF{DOFCaller: DOFCaller{contract: contract}, DOFTransactor: DOFTransactor{contract: contract}}, nil
}

// NewDOFCaller creates a new read-only instance of DOF, bound to a specific deployed contract.
func NewDOFCaller(address common.Address, caller bind.ContractCaller) (*DOFCaller, error) {
	contract, err := bindDOF(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DOFCaller{contract: contract}, nil
}

// NewDOFTransactor creates a new write-only instance of DOF, bound to a specific deployed contract.
func NewDOFTransactor(address common.Address, transactor bind.ContractTransactor) (*DOFTransactor, error) {
	contract, err := bindDOF(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DOFTransactor{contract: contract}, nil
}

// bindDOF binds a generic wrapper to an already deployed contract.
func bindDOF(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DOFABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor,nil), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DOF *DOFRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DOF.Contract.DOFCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DOF *DOFRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOF.Contract.DOFTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DOF *DOFRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DOF.Contract.DOFTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DOF *DOFCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DOF.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DOF *DOFTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOF.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DOF *DOFTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DOF.Contract.contract.Transact(opts, method, params...)
}

// _owner is a paid mutator transaction binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() returns(address)
func (_DOF *DOFTransactor) _owner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "_owner")
}

// _owner is a paid mutator transaction binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() returns(address)
func (_DOF *DOFSession) _owner() (*types.Transaction, error) {
	return _DOF.Contract._owner(&_DOF.TransactOpts)
}

// _owner is a paid mutator transaction binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() returns(address)
func (_DOF *DOFTransactorSession) _owner() (*types.Transaction, error) {
	return _DOF.Contract._owner(&_DOF.TransactOpts)
}

// _tradeLock is a paid mutator transaction binding the contract method 0xc4bc471d.
//
// Solidity: function _tradeLock() returns(bool)
func (_DOF *DOFTransactor) _tradeLock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "_tradeLock")
}

// _tradeLock is a paid mutator transaction binding the contract method 0xc4bc471d.
//
// Solidity: function _tradeLock() returns(bool)
func (_DOF *DOFSession) _tradeLock() (*types.Transaction, error) {
	return _DOF.Contract._tradeLock(&_DOF.TransactOpts)
}

// _tradeLock is a paid mutator transaction binding the contract method 0xc4bc471d.
//
// Solidity: function _tradeLock() returns(bool)
func (_DOF *DOFTransactorSession) _tradeLock() (*types.Transaction, error) {
	return _DOF.Contract._tradeLock(&_DOF.TransactOpts)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(account address) returns(bool)
func (_DOF *DOFTransactor) AddBlackList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "addBlackList", account)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(account address) returns(bool)
func (_DOF *DOFSession) AddBlackList(account common.Address) (*types.Transaction, error) {
	return _DOF.Contract.AddBlackList(&_DOF.TransactOpts, account)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(account address) returns(bool)
func (_DOF *DOFTransactorSession) AddBlackList(account common.Address) (*types.Transaction, error) {
	return _DOF.Contract.AddBlackList(&_DOF.TransactOpts, account)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0xe7cd4a04.
//
// Solidity: function addWhiteList(account address) returns(bool)
func (_DOF *DOFTransactor) AddWhiteList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "addWhiteList", account)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0xe7cd4a04.
//
// Solidity: function addWhiteList(account address) returns(bool)
func (_DOF *DOFSession) AddWhiteList(account common.Address) (*types.Transaction, error) {
	return _DOF.Contract.AddWhiteList(&_DOF.TransactOpts, account)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0xe7cd4a04.
//
// Solidity: function addWhiteList(account address) returns(bool)
func (_DOF *DOFTransactorSession) AddWhiteList(account common.Address) (*types.Transaction, error) {
	return _DOF.Contract.AddWhiteList(&_DOF.TransactOpts, account)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) returns(uint256)
func (_DOF *DOFTransactor) Allowance(opts *bind.TransactOpts, owner common.Address, spender common.Address) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "allowance", owner, spender)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) returns(uint256)
func (_DOF *DOFSession) Allowance(owner common.Address, spender common.Address) (*types.Transaction, error) {
	return _DOF.Contract.Allowance(&_DOF.TransactOpts, owner, spender)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) returns(uint256)
func (_DOF *DOFTransactorSession) Allowance(owner common.Address, spender common.Address) (*types.Transaction, error) {
	return _DOF.Contract.Allowance(&_DOF.TransactOpts, owner, spender)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, amount uint256) returns(bool)
func (_DOF *DOFTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, amount uint256) returns(bool)
func (_DOF *DOFSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DOF.Contract.Approve(&_DOF.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, amount uint256) returns(bool)
func (_DOF *DOFTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DOF.Contract.Approve(&_DOF.TransactOpts, spender, amount)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(account address) returns(uint256)
func (_DOF *DOFTransactor) BalanceOf(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "balanceOf", account)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(account address) returns(uint256)
func (_DOF *DOFSession) BalanceOf(account common.Address) (*types.Transaction, error) {
	return _DOF.Contract.BalanceOf(&_DOF.TransactOpts, account)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(account address) returns(uint256)
func (_DOF *DOFTransactorSession) BalanceOf(account common.Address) (*types.Transaction, error) {
	return _DOF.Contract.BalanceOf(&_DOF.TransactOpts, account)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns(bool)
func (_DOF *DOFTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns(bool)
func (_DOF *DOFSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _DOF.Contract.Burn(&_DOF.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(amount uint256) returns(bool)
func (_DOF *DOFTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _DOF.Contract.Burn(&_DOF.TransactOpts, amount)
}

// Decimals is a paid mutator transaction binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_DOF *DOFTransactor) Decimals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "decimals")
}

// Decimals is a paid mutator transaction binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_DOF *DOFSession) Decimals() (*types.Transaction, error) {
	return _DOF.Contract.Decimals(&_DOF.TransactOpts)
}

// Decimals is a paid mutator transaction binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_DOF *DOFTransactorSession) Decimals() (*types.Transaction, error) {
	return _DOF.Contract.Decimals(&_DOF.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_DOF *DOFTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_DOF *DOFSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DOF.Contract.DecreaseAllowance(&_DOF.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_DOF *DOFTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DOF.Contract.DecreaseAllowance(&_DOF.TransactOpts, spender, subtractedValue)
}

// GetOwner is a paid mutator transaction binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() returns(address)
func (_DOF *DOFTransactor) GetOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "getOwner")
}

// GetOwner is a paid mutator transaction binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() returns(address)
func (_DOF *DOFSession) GetOwner() (*types.Transaction, error) {
	return _DOF.Contract.GetOwner(&_DOF.TransactOpts)
}

// GetOwner is a paid mutator transaction binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() returns(address)
func (_DOF *DOFTransactorSession) GetOwner() (*types.Transaction, error) {
	return _DOF.Contract.GetOwner(&_DOF.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_DOF *DOFTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_DOF *DOFSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DOF.Contract.IncreaseAllowance(&_DOF.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_DOF *DOFTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DOF.Contract.IncreaseAllowance(&_DOF.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(account address, amount uint256) returns(bool)
func (_DOF *DOFTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(account address, amount uint256) returns(bool)
func (_DOF *DOFSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DOF.Contract.Mint(&_DOF.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(account address, amount uint256) returns(bool)
func (_DOF *DOFTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DOF.Contract.Mint(&_DOF.TransactOpts, account, amount)
}

// Name is a paid mutator transaction binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_DOF *DOFTransactor) Name(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "name")
}

// Name is a paid mutator transaction binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_DOF *DOFSession) Name() (*types.Transaction, error) {
	return _DOF.Contract.Name(&_DOF.TransactOpts)
}

// Name is a paid mutator transaction binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string)
func (_DOF *DOFTransactorSession) Name() (*types.Transaction, error) {
	return _DOF.Contract.Name(&_DOF.TransactOpts)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(account address) returns(bool)
func (_DOF *DOFTransactor) RemoveBlackList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "removeBlackList", account)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(account address) returns(bool)
func (_DOF *DOFSession) RemoveBlackList(account common.Address) (*types.Transaction, error) {
	return _DOF.Contract.RemoveBlackList(&_DOF.TransactOpts, account)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(account address) returns(bool)
func (_DOF *DOFTransactorSession) RemoveBlackList(account common.Address) (*types.Transaction, error) {
	return _DOF.Contract.RemoveBlackList(&_DOF.TransactOpts, account)
}

// RemoveWhiteList is a paid mutator transaction binding the contract method 0x2042e5c2.
//
// Solidity: function removeWhiteList(account address) returns(bool)
func (_DOF *DOFTransactor) RemoveWhiteList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "removeWhiteList", account)
}

// RemoveWhiteList is a paid mutator transaction binding the contract method 0x2042e5c2.
//
// Solidity: function removeWhiteList(account address) returns(bool)
func (_DOF *DOFSession) RemoveWhiteList(account common.Address) (*types.Transaction, error) {
	return _DOF.Contract.RemoveWhiteList(&_DOF.TransactOpts, account)
}

// RemoveWhiteList is a paid mutator transaction binding the contract method 0x2042e5c2.
//
// Solidity: function removeWhiteList(account address) returns(bool)
func (_DOF *DOFTransactorSession) RemoveWhiteList(account common.Address) (*types.Transaction, error) {
	return _DOF.Contract.RemoveWhiteList(&_DOF.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DOF *DOFTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DOF *DOFSession) RenounceOwnership() (*types.Transaction, error) {
	return _DOF.Contract.RenounceOwnership(&_DOF.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DOF *DOFTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DOF.Contract.RenounceOwnership(&_DOF.TransactOpts)
}

// Symbol is a paid mutator transaction binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_DOF *DOFTransactor) Symbol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "symbol")
}

// Symbol is a paid mutator transaction binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_DOF *DOFSession) Symbol() (*types.Transaction, error) {
	return _DOF.Contract.Symbol(&_DOF.TransactOpts)
}

// Symbol is a paid mutator transaction binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string)
func (_DOF *DOFTransactorSession) Symbol() (*types.Transaction, error) {
	return _DOF.Contract.Symbol(&_DOF.TransactOpts)
}

// TotalSupply is a paid mutator transaction binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_DOF *DOFTransactor) TotalSupply(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "totalSupply")
}

// TotalSupply is a paid mutator transaction binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_DOF *DOFSession) TotalSupply() (*types.Transaction, error) {
	return _DOF.Contract.TotalSupply(&_DOF.TransactOpts)
}

// TotalSupply is a paid mutator transaction binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_DOF *DOFTransactorSession) TotalSupply() (*types.Transaction, error) {
	return _DOF.Contract.TotalSupply(&_DOF.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(recipient address, amount uint256) returns(bool)
func (_DOF *DOFTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(recipient address, amount uint256) returns(bool)
func (_DOF *DOFSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DOF.Contract.Transfer(&_DOF.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(recipient address, amount uint256) returns(bool)
func (_DOF *DOFTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DOF.Contract.Transfer(&_DOF.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(sender address, recipient address, amount uint256) returns(bool)
func (_DOF *DOFTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(sender address, recipient address, amount uint256) returns(bool)
func (_DOF *DOFSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DOF.Contract.TransferFrom(&_DOF.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(sender address, recipient address, amount uint256) returns(bool)
func (_DOF *DOFTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DOF.Contract.TransferFrom(&_DOF.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DOF *DOFTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DOF *DOFSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DOF.Contract.TransferOwnership(&_DOF.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DOF *DOFTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DOF.Contract.TransferOwnership(&_DOF.TransactOpts, newOwner)
}

// TurnTradeLock is a paid mutator transaction binding the contract method 0x9fb4bdd0.
//
// Solidity: function turnTradeLock() returns(bool)
func (_DOF *DOFTransactor) TurnTradeLock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOF.contract.Transact(opts, "turnTradeLock")
}

// TurnTradeLock is a paid mutator transaction binding the contract method 0x9fb4bdd0.
//
// Solidity: function turnTradeLock() returns(bool)
func (_DOF *DOFSession) TurnTradeLock() (*types.Transaction, error) {
	return _DOF.Contract.TurnTradeLock(&_DOF.TransactOpts)
}

// TurnTradeLock is a paid mutator transaction binding the contract method 0x9fb4bdd0.
//
// Solidity: function turnTradeLock() returns(bool)
func (_DOF *DOFTransactorSession) TurnTradeLock() (*types.Transaction, error) {
	return _DOF.Contract.TurnTradeLock(&_DOF.TransactOpts)
}
