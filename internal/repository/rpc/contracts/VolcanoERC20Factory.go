// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// VolcanoERC720FactoryMetaData contains all meta data concerning the VolcanoERC720Factory contract.
var VolcanoERC720FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_erc20MintFeePerc\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_erc20MintTokenFeePerc\",\"type\":\"uint96\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_platformContractFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_routerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenDisabled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initialAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_capAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintBlocksFee\",\"type\":\"uint256\"}],\"name\":\"createTokenContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"}],\"name\":\"disableTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20MintFeePerc\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20MintTokenFeePerc\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintTokenBlock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformContractFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"}],\"name\":\"registerTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routerAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_erc20MintFeePerc\",\"type\":\"uint96\"}],\"name\":\"updateErc20MintFeePerc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_erc20MintTokenFeePerc\",\"type\":\"uint96\"}],\"name\":\"updateErc20MintTokenFeePerc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"updateFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformContractFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformContractFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_routerAddress\",\"type\":\"address\"}],\"name\":\"updateRouterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VolcanoERC720FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use VolcanoERC720FactoryMetaData.ABI instead.
var VolcanoERC720FactoryABI = VolcanoERC720FactoryMetaData.ABI

// VolcanoERC720Factory is an auto generated Go binding around an Ethereum contract.
type VolcanoERC720Factory struct {
	VolcanoERC720FactoryCaller     // Read-only binding to the contract
	VolcanoERC720FactoryTransactor // Write-only binding to the contract
	VolcanoERC720FactoryFilterer   // Log filterer for contract events
}

// VolcanoERC720FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VolcanoERC720FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC720FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VolcanoERC720FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC720FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolcanoERC720FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC720FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolcanoERC720FactorySession struct {
	Contract     *VolcanoERC720Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VolcanoERC720FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolcanoERC720FactoryCallerSession struct {
	Contract *VolcanoERC720FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// VolcanoERC720FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolcanoERC720FactoryTransactorSession struct {
	Contract     *VolcanoERC720FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// VolcanoERC720FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VolcanoERC720FactoryRaw struct {
	Contract *VolcanoERC720Factory // Generic contract binding to access the raw methods on
}

// VolcanoERC720FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolcanoERC720FactoryCallerRaw struct {
	Contract *VolcanoERC720FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// VolcanoERC720FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolcanoERC720FactoryTransactorRaw struct {
	Contract *VolcanoERC720FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolcanoERC720Factory creates a new instance of VolcanoERC720Factory, bound to a specific deployed contract.
func NewVolcanoERC720Factory(address common.Address, backend bind.ContractBackend) (*VolcanoERC720Factory, error) {
	contract, err := bindVolcanoERC720Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720Factory{VolcanoERC720FactoryCaller: VolcanoERC720FactoryCaller{contract: contract}, VolcanoERC720FactoryTransactor: VolcanoERC720FactoryTransactor{contract: contract}, VolcanoERC720FactoryFilterer: VolcanoERC720FactoryFilterer{contract: contract}}, nil
}

// NewVolcanoERC720FactoryCaller creates a new read-only instance of VolcanoERC720Factory, bound to a specific deployed contract.
func NewVolcanoERC720FactoryCaller(address common.Address, caller bind.ContractCaller) (*VolcanoERC720FactoryCaller, error) {
	contract, err := bindVolcanoERC720Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720FactoryCaller{contract: contract}, nil
}

// NewVolcanoERC720FactoryTransactor creates a new write-only instance of VolcanoERC720Factory, bound to a specific deployed contract.
func NewVolcanoERC720FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*VolcanoERC720FactoryTransactor, error) {
	contract, err := bindVolcanoERC720Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720FactoryTransactor{contract: contract}, nil
}

// NewVolcanoERC720FactoryFilterer creates a new log filterer instance of VolcanoERC720Factory, bound to a specific deployed contract.
func NewVolcanoERC720FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*VolcanoERC720FactoryFilterer, error) {
	contract, err := bindVolcanoERC720Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720FactoryFilterer{contract: contract}, nil
}

// bindVolcanoERC720Factory binds a generic wrapper to an already deployed contract.
func bindVolcanoERC720Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VolcanoERC720FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC720Factory *VolcanoERC720FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC720Factory.Contract.VolcanoERC720FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC720Factory *VolcanoERC720FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.VolcanoERC720FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC720Factory *VolcanoERC720FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.VolcanoERC720FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC720Factory *VolcanoERC720FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC720Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.contract.Transact(opts, method, params...)
}

// Erc20MintFeePerc is a free data retrieval call binding the contract method 0xf4143c6d.
//
// Solidity: function erc20MintFeePerc() view returns(uint96)
func (_VolcanoERC720Factory *VolcanoERC720FactoryCaller) Erc20MintFeePerc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC720Factory.contract.Call(opts, &out, "erc20MintFeePerc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Erc20MintFeePerc is a free data retrieval call binding the contract method 0xf4143c6d.
//
// Solidity: function erc20MintFeePerc() view returns(uint96)
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) Erc20MintFeePerc() (*big.Int, error) {
	return _VolcanoERC720Factory.Contract.Erc20MintFeePerc(&_VolcanoERC720Factory.CallOpts)
}

// Erc20MintFeePerc is a free data retrieval call binding the contract method 0xf4143c6d.
//
// Solidity: function erc20MintFeePerc() view returns(uint96)
func (_VolcanoERC720Factory *VolcanoERC720FactoryCallerSession) Erc20MintFeePerc() (*big.Int, error) {
	return _VolcanoERC720Factory.Contract.Erc20MintFeePerc(&_VolcanoERC720Factory.CallOpts)
}

// Erc20MintTokenFeePerc is a free data retrieval call binding the contract method 0x02b333f2.
//
// Solidity: function erc20MintTokenFeePerc() view returns(uint96)
func (_VolcanoERC720Factory *VolcanoERC720FactoryCaller) Erc20MintTokenFeePerc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC720Factory.contract.Call(opts, &out, "erc20MintTokenFeePerc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Erc20MintTokenFeePerc is a free data retrieval call binding the contract method 0x02b333f2.
//
// Solidity: function erc20MintTokenFeePerc() view returns(uint96)
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) Erc20MintTokenFeePerc() (*big.Int, error) {
	return _VolcanoERC720Factory.Contract.Erc20MintTokenFeePerc(&_VolcanoERC720Factory.CallOpts)
}

// Erc20MintTokenFeePerc is a free data retrieval call binding the contract method 0x02b333f2.
//
// Solidity: function erc20MintTokenFeePerc() view returns(uint96)
func (_VolcanoERC720Factory *VolcanoERC720FactoryCallerSession) Erc20MintTokenFeePerc() (*big.Int, error) {
	return _VolcanoERC720Factory.Contract.Erc20MintTokenFeePerc(&_VolcanoERC720Factory.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_VolcanoERC720Factory *VolcanoERC720FactoryCaller) Exists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VolcanoERC720Factory.contract.Call(opts, &out, "exists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) Exists(arg0 common.Address) (bool, error) {
	return _VolcanoERC720Factory.Contract.Exists(&_VolcanoERC720Factory.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_VolcanoERC720Factory *VolcanoERC720FactoryCallerSession) Exists(arg0 common.Address) (bool, error) {
	return _VolcanoERC720Factory.Contract.Exists(&_VolcanoERC720Factory.CallOpts, arg0)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC720Factory *VolcanoERC720FactoryCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC720Factory.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) FeeRecipient() (common.Address, error) {
	return _VolcanoERC720Factory.Contract.FeeRecipient(&_VolcanoERC720Factory.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC720Factory *VolcanoERC720FactoryCallerSession) FeeRecipient() (common.Address, error) {
	return _VolcanoERC720Factory.Contract.FeeRecipient(&_VolcanoERC720Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC720Factory *VolcanoERC720FactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC720Factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) Owner() (common.Address, error) {
	return _VolcanoERC720Factory.Contract.Owner(&_VolcanoERC720Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC720Factory *VolcanoERC720FactoryCallerSession) Owner() (common.Address, error) {
	return _VolcanoERC720Factory.Contract.Owner(&_VolcanoERC720Factory.CallOpts)
}

// PlatformContractFee is a free data retrieval call binding the contract method 0xb69d4b98.
//
// Solidity: function platformContractFee() view returns(uint256)
func (_VolcanoERC720Factory *VolcanoERC720FactoryCaller) PlatformContractFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC720Factory.contract.Call(opts, &out, "platformContractFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformContractFee is a free data retrieval call binding the contract method 0xb69d4b98.
//
// Solidity: function platformContractFee() view returns(uint256)
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) PlatformContractFee() (*big.Int, error) {
	return _VolcanoERC720Factory.Contract.PlatformContractFee(&_VolcanoERC720Factory.CallOpts)
}

// PlatformContractFee is a free data retrieval call binding the contract method 0xb69d4b98.
//
// Solidity: function platformContractFee() view returns(uint256)
func (_VolcanoERC720Factory *VolcanoERC720FactoryCallerSession) PlatformContractFee() (*big.Int, error) {
	return _VolcanoERC720Factory.Contract.PlatformContractFee(&_VolcanoERC720Factory.CallOpts)
}

// RouterAddress is a free data retrieval call binding the contract method 0x3268cc56.
//
// Solidity: function routerAddress() view returns(address)
func (_VolcanoERC720Factory *VolcanoERC720FactoryCaller) RouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC720Factory.contract.Call(opts, &out, "routerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RouterAddress is a free data retrieval call binding the contract method 0x3268cc56.
//
// Solidity: function routerAddress() view returns(address)
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) RouterAddress() (common.Address, error) {
	return _VolcanoERC720Factory.Contract.RouterAddress(&_VolcanoERC720Factory.CallOpts)
}

// RouterAddress is a free data retrieval call binding the contract method 0x3268cc56.
//
// Solidity: function routerAddress() view returns(address)
func (_VolcanoERC720Factory *VolcanoERC720FactoryCallerSession) RouterAddress() (common.Address, error) {
	return _VolcanoERC720Factory.Contract.RouterAddress(&_VolcanoERC720Factory.CallOpts)
}

// CreateTokenContract is a paid mutator transaction binding the contract method 0xc3d28c65.
//
// Solidity: function createTokenContract(string _name, string _symbol, string _uri, address _initialReceiver, uint256 _initialAmount, uint256 _capAmount, uint256 _mintBlocks, uint256 _mintBlocksFee) payable returns(address)
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactor) CreateTokenContract(opts *bind.TransactOpts, _name string, _symbol string, _uri string, _initialReceiver common.Address, _initialAmount *big.Int, _capAmount *big.Int, _mintBlocks *big.Int, _mintBlocksFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Factory.contract.Transact(opts, "createTokenContract", _name, _symbol, _uri, _initialReceiver, _initialAmount, _capAmount, _mintBlocks, _mintBlocksFee)
}

// CreateTokenContract is a paid mutator transaction binding the contract method 0xc3d28c65.
//
// Solidity: function createTokenContract(string _name, string _symbol, string _uri, address _initialReceiver, uint256 _initialAmount, uint256 _capAmount, uint256 _mintBlocks, uint256 _mintBlocksFee) payable returns(address)
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) CreateTokenContract(_name string, _symbol string, _uri string, _initialReceiver common.Address, _initialAmount *big.Int, _capAmount *big.Int, _mintBlocks *big.Int, _mintBlocksFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.CreateTokenContract(&_VolcanoERC720Factory.TransactOpts, _name, _symbol, _uri, _initialReceiver, _initialAmount, _capAmount, _mintBlocks, _mintBlocksFee)
}

// CreateTokenContract is a paid mutator transaction binding the contract method 0xc3d28c65.
//
// Solidity: function createTokenContract(string _name, string _symbol, string _uri, address _initialReceiver, uint256 _initialAmount, uint256 _capAmount, uint256 _mintBlocks, uint256 _mintBlocksFee) payable returns(address)
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactorSession) CreateTokenContract(_name string, _symbol string, _uri string, _initialReceiver common.Address, _initialAmount *big.Int, _capAmount *big.Int, _mintBlocks *big.Int, _mintBlocksFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.CreateTokenContract(&_VolcanoERC720Factory.TransactOpts, _name, _symbol, _uri, _initialReceiver, _initialAmount, _capAmount, _mintBlocks, _mintBlocksFee)
}

// DisableTokenContract is a paid mutator transaction binding the contract method 0xd39bb16a.
//
// Solidity: function disableTokenContract(address tokenContractAddress) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactor) DisableTokenContract(opts *bind.TransactOpts, tokenContractAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.contract.Transact(opts, "disableTokenContract", tokenContractAddress)
}

// DisableTokenContract is a paid mutator transaction binding the contract method 0xd39bb16a.
//
// Solidity: function disableTokenContract(address tokenContractAddress) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) DisableTokenContract(tokenContractAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.DisableTokenContract(&_VolcanoERC720Factory.TransactOpts, tokenContractAddress)
}

// DisableTokenContract is a paid mutator transaction binding the contract method 0xd39bb16a.
//
// Solidity: function disableTokenContract(address tokenContractAddress) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactorSession) DisableTokenContract(tokenContractAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.DisableTokenContract(&_VolcanoERC720Factory.TransactOpts, tokenContractAddress)
}

// MintTokenBlock is a paid mutator transaction binding the contract method 0xa8bef5bf.
//
// Solidity: function mintTokenBlock(address tokenContractAddress, address to) payable returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactor) MintTokenBlock(opts *bind.TransactOpts, tokenContractAddress common.Address, to common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.contract.Transact(opts, "mintTokenBlock", tokenContractAddress, to)
}

// MintTokenBlock is a paid mutator transaction binding the contract method 0xa8bef5bf.
//
// Solidity: function mintTokenBlock(address tokenContractAddress, address to) payable returns()
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) MintTokenBlock(tokenContractAddress common.Address, to common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.MintTokenBlock(&_VolcanoERC720Factory.TransactOpts, tokenContractAddress, to)
}

// MintTokenBlock is a paid mutator transaction binding the contract method 0xa8bef5bf.
//
// Solidity: function mintTokenBlock(address tokenContractAddress, address to) payable returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactorSession) MintTokenBlock(tokenContractAddress common.Address, to common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.MintTokenBlock(&_VolcanoERC720Factory.TransactOpts, tokenContractAddress, to)
}

// RegisterTokenContract is a paid mutator transaction binding the contract method 0xf7b0770b.
//
// Solidity: function registerTokenContract(address tokenContractAddress) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactor) RegisterTokenContract(opts *bind.TransactOpts, tokenContractAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.contract.Transact(opts, "registerTokenContract", tokenContractAddress)
}

// RegisterTokenContract is a paid mutator transaction binding the contract method 0xf7b0770b.
//
// Solidity: function registerTokenContract(address tokenContractAddress) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) RegisterTokenContract(tokenContractAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.RegisterTokenContract(&_VolcanoERC720Factory.TransactOpts, tokenContractAddress)
}

// RegisterTokenContract is a paid mutator transaction binding the contract method 0xf7b0770b.
//
// Solidity: function registerTokenContract(address tokenContractAddress) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactorSession) RegisterTokenContract(tokenContractAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.RegisterTokenContract(&_VolcanoERC720Factory.TransactOpts, tokenContractAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC720Factory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.RenounceOwnership(&_VolcanoERC720Factory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.RenounceOwnership(&_VolcanoERC720Factory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.TransferOwnership(&_VolcanoERC720Factory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.TransferOwnership(&_VolcanoERC720Factory.TransactOpts, newOwner)
}

// UpdateErc20MintFeePerc is a paid mutator transaction binding the contract method 0x6be852ac.
//
// Solidity: function updateErc20MintFeePerc(uint96 _erc20MintFeePerc) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactor) UpdateErc20MintFeePerc(opts *bind.TransactOpts, _erc20MintFeePerc *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Factory.contract.Transact(opts, "updateErc20MintFeePerc", _erc20MintFeePerc)
}

// UpdateErc20MintFeePerc is a paid mutator transaction binding the contract method 0x6be852ac.
//
// Solidity: function updateErc20MintFeePerc(uint96 _erc20MintFeePerc) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) UpdateErc20MintFeePerc(_erc20MintFeePerc *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.UpdateErc20MintFeePerc(&_VolcanoERC720Factory.TransactOpts, _erc20MintFeePerc)
}

// UpdateErc20MintFeePerc is a paid mutator transaction binding the contract method 0x6be852ac.
//
// Solidity: function updateErc20MintFeePerc(uint96 _erc20MintFeePerc) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactorSession) UpdateErc20MintFeePerc(_erc20MintFeePerc *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.UpdateErc20MintFeePerc(&_VolcanoERC720Factory.TransactOpts, _erc20MintFeePerc)
}

// UpdateErc20MintTokenFeePerc is a paid mutator transaction binding the contract method 0xceb7ff64.
//
// Solidity: function updateErc20MintTokenFeePerc(uint96 _erc20MintTokenFeePerc) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactor) UpdateErc20MintTokenFeePerc(opts *bind.TransactOpts, _erc20MintTokenFeePerc *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Factory.contract.Transact(opts, "updateErc20MintTokenFeePerc", _erc20MintTokenFeePerc)
}

// UpdateErc20MintTokenFeePerc is a paid mutator transaction binding the contract method 0xceb7ff64.
//
// Solidity: function updateErc20MintTokenFeePerc(uint96 _erc20MintTokenFeePerc) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) UpdateErc20MintTokenFeePerc(_erc20MintTokenFeePerc *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.UpdateErc20MintTokenFeePerc(&_VolcanoERC720Factory.TransactOpts, _erc20MintTokenFeePerc)
}

// UpdateErc20MintTokenFeePerc is a paid mutator transaction binding the contract method 0xceb7ff64.
//
// Solidity: function updateErc20MintTokenFeePerc(uint96 _erc20MintTokenFeePerc) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactorSession) UpdateErc20MintTokenFeePerc(_erc20MintTokenFeePerc *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.UpdateErc20MintTokenFeePerc(&_VolcanoERC720Factory.TransactOpts, _erc20MintTokenFeePerc)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactor) UpdateFeeRecipient(opts *bind.TransactOpts, _feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.contract.Transact(opts, "updateFeeRecipient", _feeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) UpdateFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.UpdateFeeRecipient(&_VolcanoERC720Factory.TransactOpts, _feeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactorSession) UpdateFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.UpdateFeeRecipient(&_VolcanoERC720Factory.TransactOpts, _feeRecipient)
}

// UpdatePlatformContractFee is a paid mutator transaction binding the contract method 0xc57482bb.
//
// Solidity: function updatePlatformContractFee(uint256 _platformContractFee) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactor) UpdatePlatformContractFee(opts *bind.TransactOpts, _platformContractFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Factory.contract.Transact(opts, "updatePlatformContractFee", _platformContractFee)
}

// UpdatePlatformContractFee is a paid mutator transaction binding the contract method 0xc57482bb.
//
// Solidity: function updatePlatformContractFee(uint256 _platformContractFee) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) UpdatePlatformContractFee(_platformContractFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.UpdatePlatformContractFee(&_VolcanoERC720Factory.TransactOpts, _platformContractFee)
}

// UpdatePlatformContractFee is a paid mutator transaction binding the contract method 0xc57482bb.
//
// Solidity: function updatePlatformContractFee(uint256 _platformContractFee) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactorSession) UpdatePlatformContractFee(_platformContractFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.UpdatePlatformContractFee(&_VolcanoERC720Factory.TransactOpts, _platformContractFee)
}

// UpdateRouterAddress is a paid mutator transaction binding the contract method 0x00e6be7b.
//
// Solidity: function updateRouterAddress(address _routerAddress) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactor) UpdateRouterAddress(opts *bind.TransactOpts, _routerAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.contract.Transact(opts, "updateRouterAddress", _routerAddress)
}

// UpdateRouterAddress is a paid mutator transaction binding the contract method 0x00e6be7b.
//
// Solidity: function updateRouterAddress(address _routerAddress) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactorySession) UpdateRouterAddress(_routerAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.UpdateRouterAddress(&_VolcanoERC720Factory.TransactOpts, _routerAddress)
}

// UpdateRouterAddress is a paid mutator transaction binding the contract method 0x00e6be7b.
//
// Solidity: function updateRouterAddress(address _routerAddress) returns()
func (_VolcanoERC720Factory *VolcanoERC720FactoryTransactorSession) UpdateRouterAddress(_routerAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Factory.Contract.UpdateRouterAddress(&_VolcanoERC720Factory.TransactOpts, _routerAddress)
}

// VolcanoERC720FactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VolcanoERC720Factory contract.
type VolcanoERC720FactoryOwnershipTransferredIterator struct {
	Event *VolcanoERC720FactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VolcanoERC720FactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC720FactoryOwnershipTransferred)
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
		it.Event = new(VolcanoERC720FactoryOwnershipTransferred)
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
func (it *VolcanoERC720FactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC720FactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC720FactoryOwnershipTransferred represents a OwnershipTransferred event raised by the VolcanoERC720Factory contract.
type VolcanoERC720FactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC720Factory *VolcanoERC720FactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VolcanoERC720FactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC720Factory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720FactoryOwnershipTransferredIterator{contract: _VolcanoERC720Factory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC720Factory *VolcanoERC720FactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VolcanoERC720FactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC720Factory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC720FactoryOwnershipTransferred)
				if err := _VolcanoERC720Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC720Factory *VolcanoERC720FactoryFilterer) ParseOwnershipTransferred(log types.Log) (*VolcanoERC720FactoryOwnershipTransferred, error) {
	event := new(VolcanoERC720FactoryOwnershipTransferred)
	if err := _VolcanoERC720Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC720FactoryTokenCreatedIterator is returned from FilterTokenCreated and is used to iterate over the raw logs and unpacked data for TokenCreated events raised by the VolcanoERC720Factory contract.
type VolcanoERC720FactoryTokenCreatedIterator struct {
	Event *VolcanoERC720FactoryTokenCreated // Event containing the contract specifics and raw log

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
func (it *VolcanoERC720FactoryTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC720FactoryTokenCreated)
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
		it.Event = new(VolcanoERC720FactoryTokenCreated)
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
func (it *VolcanoERC720FactoryTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC720FactoryTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC720FactoryTokenCreated represents a TokenCreated event raised by the VolcanoERC720Factory contract.
type VolcanoERC720FactoryTokenCreated struct {
	Caller common.Address
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenCreated is a free log retrieval operation binding the contract event 0xd5f9bdf12adf29dab0248c349842c3822d53ae2bb4f36352f301630d018c8139.
//
// Solidity: event TokenCreated(address caller, address token)
func (_VolcanoERC720Factory *VolcanoERC720FactoryFilterer) FilterTokenCreated(opts *bind.FilterOpts) (*VolcanoERC720FactoryTokenCreatedIterator, error) {

	logs, sub, err := _VolcanoERC720Factory.contract.FilterLogs(opts, "TokenCreated")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720FactoryTokenCreatedIterator{contract: _VolcanoERC720Factory.contract, event: "TokenCreated", logs: logs, sub: sub}, nil
}

// WatchTokenCreated is a free log subscription operation binding the contract event 0xd5f9bdf12adf29dab0248c349842c3822d53ae2bb4f36352f301630d018c8139.
//
// Solidity: event TokenCreated(address caller, address token)
func (_VolcanoERC720Factory *VolcanoERC720FactoryFilterer) WatchTokenCreated(opts *bind.WatchOpts, sink chan<- *VolcanoERC720FactoryTokenCreated) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC720Factory.contract.WatchLogs(opts, "TokenCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC720FactoryTokenCreated)
				if err := _VolcanoERC720Factory.contract.UnpackLog(event, "TokenCreated", log); err != nil {
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

// ParseTokenCreated is a log parse operation binding the contract event 0xd5f9bdf12adf29dab0248c349842c3822d53ae2bb4f36352f301630d018c8139.
//
// Solidity: event TokenCreated(address caller, address token)
func (_VolcanoERC720Factory *VolcanoERC720FactoryFilterer) ParseTokenCreated(log types.Log) (*VolcanoERC720FactoryTokenCreated, error) {
	event := new(VolcanoERC720FactoryTokenCreated)
	if err := _VolcanoERC720Factory.contract.UnpackLog(event, "TokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC720FactoryTokenDisabledIterator is returned from FilterTokenDisabled and is used to iterate over the raw logs and unpacked data for TokenDisabled events raised by the VolcanoERC720Factory contract.
type VolcanoERC720FactoryTokenDisabledIterator struct {
	Event *VolcanoERC720FactoryTokenDisabled // Event containing the contract specifics and raw log

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
func (it *VolcanoERC720FactoryTokenDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC720FactoryTokenDisabled)
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
		it.Event = new(VolcanoERC720FactoryTokenDisabled)
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
func (it *VolcanoERC720FactoryTokenDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC720FactoryTokenDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC720FactoryTokenDisabled represents a TokenDisabled event raised by the VolcanoERC720Factory contract.
type VolcanoERC720FactoryTokenDisabled struct {
	Caller common.Address
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenDisabled is a free log retrieval operation binding the contract event 0x69db9e9764da61b6681b89fb0f5be018506e0342036e56b59c3015a620d60977.
//
// Solidity: event TokenDisabled(address caller, address token)
func (_VolcanoERC720Factory *VolcanoERC720FactoryFilterer) FilterTokenDisabled(opts *bind.FilterOpts) (*VolcanoERC720FactoryTokenDisabledIterator, error) {

	logs, sub, err := _VolcanoERC720Factory.contract.FilterLogs(opts, "TokenDisabled")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720FactoryTokenDisabledIterator{contract: _VolcanoERC720Factory.contract, event: "TokenDisabled", logs: logs, sub: sub}, nil
}

// WatchTokenDisabled is a free log subscription operation binding the contract event 0x69db9e9764da61b6681b89fb0f5be018506e0342036e56b59c3015a620d60977.
//
// Solidity: event TokenDisabled(address caller, address token)
func (_VolcanoERC720Factory *VolcanoERC720FactoryFilterer) WatchTokenDisabled(opts *bind.WatchOpts, sink chan<- *VolcanoERC720FactoryTokenDisabled) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC720Factory.contract.WatchLogs(opts, "TokenDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC720FactoryTokenDisabled)
				if err := _VolcanoERC720Factory.contract.UnpackLog(event, "TokenDisabled", log); err != nil {
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

// ParseTokenDisabled is a log parse operation binding the contract event 0x69db9e9764da61b6681b89fb0f5be018506e0342036e56b59c3015a620d60977.
//
// Solidity: event TokenDisabled(address caller, address token)
func (_VolcanoERC720Factory *VolcanoERC720FactoryFilterer) ParseTokenDisabled(log types.Log) (*VolcanoERC720FactoryTokenDisabled, error) {
	event := new(VolcanoERC720FactoryTokenDisabled)
	if err := _VolcanoERC720Factory.contract.UnpackLog(event, "TokenDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
