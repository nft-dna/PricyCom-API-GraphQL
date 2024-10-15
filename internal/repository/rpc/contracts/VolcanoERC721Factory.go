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

// VolcanoERC721TradablecontractERC721Options is an auto generated low-level Go binding around an user-defined struct.
type VolcanoERC721TradablecontractERC721Options struct {
	BaseUri       string
	UseDecimalUri bool
	BaseUriExt    string
	MaxItems      *big.Int
	MintStartTime *big.Int
	MintStopTime  *big.Int
	RevealTime    *big.Int
	PreRevealUri  string
	ContractUri   string
}

// VolcanoERC721FactoryMetaData contains all meta data concerning the VolcanoERC721Factory contract.
var VolcanoERC721FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auction\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marketplace\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bundleMarketplace\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_platformContractFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_platformMintFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platform\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isprivate\",\"type\":\"bool\"}],\"name\":\"ContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"}],\"name\":\"ContractDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bundleMarketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_isprivate\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_mintFee\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"_creatorFeePerc\",\"type\":\"uint96\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"useDecimalUri\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"baseUriExt\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxItems\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintStopTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"preRevealUri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contractUri\",\"type\":\"string\"}],\"internalType\":\"structVolcanoERC721Tradable.contractERC721Options\",\"name\":\"_options\",\"type\":\"tuple\"}],\"name\":\"createNFTContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"}],\"name\":\"disableTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isprivate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformContractFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformMintFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isprivate\",\"type\":\"bool\"}],\"name\":\"registerTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auction\",\"type\":\"address\"}],\"name\":\"updateAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bundleMarketplace\",\"type\":\"address\"}],\"name\":\"updateBundleMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"updateFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplace\",\"type\":\"address\"}],\"name\":\"updateMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformContractFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformContractFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformMintFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformMintFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VolcanoERC721FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use VolcanoERC721FactoryMetaData.ABI instead.
var VolcanoERC721FactoryABI = VolcanoERC721FactoryMetaData.ABI

// VolcanoERC721Factory is an auto generated Go binding around an Ethereum contract.
type VolcanoERC721Factory struct {
	VolcanoERC721FactoryCaller     // Read-only binding to the contract
	VolcanoERC721FactoryTransactor // Write-only binding to the contract
	VolcanoERC721FactoryFilterer   // Log filterer for contract events
}

// VolcanoERC721FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VolcanoERC721FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC721FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VolcanoERC721FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC721FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolcanoERC721FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC721FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolcanoERC721FactorySession struct {
	Contract     *VolcanoERC721Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VolcanoERC721FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolcanoERC721FactoryCallerSession struct {
	Contract *VolcanoERC721FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// VolcanoERC721FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolcanoERC721FactoryTransactorSession struct {
	Contract     *VolcanoERC721FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// VolcanoERC721FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VolcanoERC721FactoryRaw struct {
	Contract *VolcanoERC721Factory // Generic contract binding to access the raw methods on
}

// VolcanoERC721FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolcanoERC721FactoryCallerRaw struct {
	Contract *VolcanoERC721FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// VolcanoERC721FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolcanoERC721FactoryTransactorRaw struct {
	Contract *VolcanoERC721FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolcanoERC721Factory creates a new instance of VolcanoERC721Factory, bound to a specific deployed contract.
func NewVolcanoERC721Factory(address common.Address, backend bind.ContractBackend) (*VolcanoERC721Factory, error) {
	contract, err := bindVolcanoERC721Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721Factory{VolcanoERC721FactoryCaller: VolcanoERC721FactoryCaller{contract: contract}, VolcanoERC721FactoryTransactor: VolcanoERC721FactoryTransactor{contract: contract}, VolcanoERC721FactoryFilterer: VolcanoERC721FactoryFilterer{contract: contract}}, nil
}

// NewVolcanoERC721FactoryCaller creates a new read-only instance of VolcanoERC721Factory, bound to a specific deployed contract.
func NewVolcanoERC721FactoryCaller(address common.Address, caller bind.ContractCaller) (*VolcanoERC721FactoryCaller, error) {
	contract, err := bindVolcanoERC721Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721FactoryCaller{contract: contract}, nil
}

// NewVolcanoERC721FactoryTransactor creates a new write-only instance of VolcanoERC721Factory, bound to a specific deployed contract.
func NewVolcanoERC721FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*VolcanoERC721FactoryTransactor, error) {
	contract, err := bindVolcanoERC721Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721FactoryTransactor{contract: contract}, nil
}

// NewVolcanoERC721FactoryFilterer creates a new log filterer instance of VolcanoERC721Factory, bound to a specific deployed contract.
func NewVolcanoERC721FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*VolcanoERC721FactoryFilterer, error) {
	contract, err := bindVolcanoERC721Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721FactoryFilterer{contract: contract}, nil
}

// bindVolcanoERC721Factory binds a generic wrapper to an already deployed contract.
func bindVolcanoERC721Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VolcanoERC721FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC721Factory *VolcanoERC721FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC721Factory.Contract.VolcanoERC721FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC721Factory *VolcanoERC721FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.VolcanoERC721FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC721Factory *VolcanoERC721FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.VolcanoERC721FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC721Factory *VolcanoERC721FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC721Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.contract.Transact(opts, method, params...)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCaller) Auction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC721Factory.contract.Call(opts, &out, "auction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) Auction() (common.Address, error) {
	return _VolcanoERC721Factory.Contract.Auction(&_VolcanoERC721Factory.CallOpts)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCallerSession) Auction() (common.Address, error) {
	return _VolcanoERC721Factory.Contract.Auction(&_VolcanoERC721Factory.CallOpts)
}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCaller) BundleMarketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC721Factory.contract.Call(opts, &out, "bundleMarketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) BundleMarketplace() (common.Address, error) {
	return _VolcanoERC721Factory.Contract.BundleMarketplace(&_VolcanoERC721Factory.CallOpts)
}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCallerSession) BundleMarketplace() (common.Address, error) {
	return _VolcanoERC721Factory.Contract.BundleMarketplace(&_VolcanoERC721Factory.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCaller) Exists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VolcanoERC721Factory.contract.Call(opts, &out, "exists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) Exists(arg0 common.Address) (bool, error) {
	return _VolcanoERC721Factory.Contract.Exists(&_VolcanoERC721Factory.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCallerSession) Exists(arg0 common.Address) (bool, error) {
	return _VolcanoERC721Factory.Contract.Exists(&_VolcanoERC721Factory.CallOpts, arg0)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC721Factory.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) FeeRecipient() (common.Address, error) {
	return _VolcanoERC721Factory.Contract.FeeRecipient(&_VolcanoERC721Factory.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCallerSession) FeeRecipient() (common.Address, error) {
	return _VolcanoERC721Factory.Contract.FeeRecipient(&_VolcanoERC721Factory.CallOpts)
}

// Isprivate is a free data retrieval call binding the contract method 0x50ac3f47.
//
// Solidity: function isprivate(address ) view returns(bool)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCaller) Isprivate(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VolcanoERC721Factory.contract.Call(opts, &out, "isprivate", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Isprivate is a free data retrieval call binding the contract method 0x50ac3f47.
//
// Solidity: function isprivate(address ) view returns(bool)
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) Isprivate(arg0 common.Address) (bool, error) {
	return _VolcanoERC721Factory.Contract.Isprivate(&_VolcanoERC721Factory.CallOpts, arg0)
}

// Isprivate is a free data retrieval call binding the contract method 0x50ac3f47.
//
// Solidity: function isprivate(address ) view returns(bool)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCallerSession) Isprivate(arg0 common.Address) (bool, error) {
	return _VolcanoERC721Factory.Contract.Isprivate(&_VolcanoERC721Factory.CallOpts, arg0)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCaller) Marketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC721Factory.contract.Call(opts, &out, "marketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) Marketplace() (common.Address, error) {
	return _VolcanoERC721Factory.Contract.Marketplace(&_VolcanoERC721Factory.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCallerSession) Marketplace() (common.Address, error) {
	return _VolcanoERC721Factory.Contract.Marketplace(&_VolcanoERC721Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC721Factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) Owner() (common.Address, error) {
	return _VolcanoERC721Factory.Contract.Owner(&_VolcanoERC721Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCallerSession) Owner() (common.Address, error) {
	return _VolcanoERC721Factory.Contract.Owner(&_VolcanoERC721Factory.CallOpts)
}

// PlatformContractFee is a free data retrieval call binding the contract method 0xb69d4b98.
//
// Solidity: function platformContractFee() view returns(uint256)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCaller) PlatformContractFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721Factory.contract.Call(opts, &out, "platformContractFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformContractFee is a free data retrieval call binding the contract method 0xb69d4b98.
//
// Solidity: function platformContractFee() view returns(uint256)
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) PlatformContractFee() (*big.Int, error) {
	return _VolcanoERC721Factory.Contract.PlatformContractFee(&_VolcanoERC721Factory.CallOpts)
}

// PlatformContractFee is a free data retrieval call binding the contract method 0xb69d4b98.
//
// Solidity: function platformContractFee() view returns(uint256)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCallerSession) PlatformContractFee() (*big.Int, error) {
	return _VolcanoERC721Factory.Contract.PlatformContractFee(&_VolcanoERC721Factory.CallOpts)
}

// PlatformMintFee is a free data retrieval call binding the contract method 0xce922b7b.
//
// Solidity: function platformMintFee() view returns(uint256)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCaller) PlatformMintFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721Factory.contract.Call(opts, &out, "platformMintFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformMintFee is a free data retrieval call binding the contract method 0xce922b7b.
//
// Solidity: function platformMintFee() view returns(uint256)
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) PlatformMintFee() (*big.Int, error) {
	return _VolcanoERC721Factory.Contract.PlatformMintFee(&_VolcanoERC721Factory.CallOpts)
}

// PlatformMintFee is a free data retrieval call binding the contract method 0xce922b7b.
//
// Solidity: function platformMintFee() view returns(uint256)
func (_VolcanoERC721Factory *VolcanoERC721FactoryCallerSession) PlatformMintFee() (*big.Int, error) {
	return _VolcanoERC721Factory.Contract.PlatformMintFee(&_VolcanoERC721Factory.CallOpts)
}

// CreateNFTContract is a paid mutator transaction binding the contract method 0x1087acfb.
//
// Solidity: function createNFTContract(string _name, string _symbol, bool _isprivate, uint256 _mintFee, uint96 _creatorFeePerc, address _feeRecipient, (string,bool,string,uint256,uint256,uint256,uint256,string,string) _options) payable returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactor) CreateNFTContract(opts *bind.TransactOpts, _name string, _symbol string, _isprivate bool, _mintFee *big.Int, _creatorFeePerc *big.Int, _feeRecipient common.Address, _options VolcanoERC721TradablecontractERC721Options) (*types.Transaction, error) {
	return _VolcanoERC721Factory.contract.Transact(opts, "createNFTContract", _name, _symbol, _isprivate, _mintFee, _creatorFeePerc, _feeRecipient, _options)
}

// CreateNFTContract is a paid mutator transaction binding the contract method 0x1087acfb.
//
// Solidity: function createNFTContract(string _name, string _symbol, bool _isprivate, uint256 _mintFee, uint96 _creatorFeePerc, address _feeRecipient, (string,bool,string,uint256,uint256,uint256,uint256,string,string) _options) payable returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) CreateNFTContract(_name string, _symbol string, _isprivate bool, _mintFee *big.Int, _creatorFeePerc *big.Int, _feeRecipient common.Address, _options VolcanoERC721TradablecontractERC721Options) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.CreateNFTContract(&_VolcanoERC721Factory.TransactOpts, _name, _symbol, _isprivate, _mintFee, _creatorFeePerc, _feeRecipient, _options)
}

// CreateNFTContract is a paid mutator transaction binding the contract method 0x1087acfb.
//
// Solidity: function createNFTContract(string _name, string _symbol, bool _isprivate, uint256 _mintFee, uint96 _creatorFeePerc, address _feeRecipient, (string,bool,string,uint256,uint256,uint256,uint256,string,string) _options) payable returns(address)
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactorSession) CreateNFTContract(_name string, _symbol string, _isprivate bool, _mintFee *big.Int, _creatorFeePerc *big.Int, _feeRecipient common.Address, _options VolcanoERC721TradablecontractERC721Options) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.CreateNFTContract(&_VolcanoERC721Factory.TransactOpts, _name, _symbol, _isprivate, _mintFee, _creatorFeePerc, _feeRecipient, _options)
}

// DisableTokenContract is a paid mutator transaction binding the contract method 0xd39bb16a.
//
// Solidity: function disableTokenContract(address tokenContractAddress) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactor) DisableTokenContract(opts *bind.TransactOpts, tokenContractAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.contract.Transact(opts, "disableTokenContract", tokenContractAddress)
}

// DisableTokenContract is a paid mutator transaction binding the contract method 0xd39bb16a.
//
// Solidity: function disableTokenContract(address tokenContractAddress) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) DisableTokenContract(tokenContractAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.DisableTokenContract(&_VolcanoERC721Factory.TransactOpts, tokenContractAddress)
}

// DisableTokenContract is a paid mutator transaction binding the contract method 0xd39bb16a.
//
// Solidity: function disableTokenContract(address tokenContractAddress) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactorSession) DisableTokenContract(tokenContractAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.DisableTokenContract(&_VolcanoERC721Factory.TransactOpts, tokenContractAddress)
}

// RegisterTokenContract is a paid mutator transaction binding the contract method 0x9c7d1a4e.
//
// Solidity: function registerTokenContract(address tokenContractAddress, bool _isprivate) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactor) RegisterTokenContract(opts *bind.TransactOpts, tokenContractAddress common.Address, _isprivate bool) (*types.Transaction, error) {
	return _VolcanoERC721Factory.contract.Transact(opts, "registerTokenContract", tokenContractAddress, _isprivate)
}

// RegisterTokenContract is a paid mutator transaction binding the contract method 0x9c7d1a4e.
//
// Solidity: function registerTokenContract(address tokenContractAddress, bool _isprivate) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) RegisterTokenContract(tokenContractAddress common.Address, _isprivate bool) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.RegisterTokenContract(&_VolcanoERC721Factory.TransactOpts, tokenContractAddress, _isprivate)
}

// RegisterTokenContract is a paid mutator transaction binding the contract method 0x9c7d1a4e.
//
// Solidity: function registerTokenContract(address tokenContractAddress, bool _isprivate) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactorSession) RegisterTokenContract(tokenContractAddress common.Address, _isprivate bool) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.RegisterTokenContract(&_VolcanoERC721Factory.TransactOpts, tokenContractAddress, _isprivate)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC721Factory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.RenounceOwnership(&_VolcanoERC721Factory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.RenounceOwnership(&_VolcanoERC721Factory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.TransferOwnership(&_VolcanoERC721Factory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.TransferOwnership(&_VolcanoERC721Factory.TransactOpts, newOwner)
}

// UpdateAuction is a paid mutator transaction binding the contract method 0x3402a036.
//
// Solidity: function updateAuction(address _auction) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactor) UpdateAuction(opts *bind.TransactOpts, _auction common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.contract.Transact(opts, "updateAuction", _auction)
}

// UpdateAuction is a paid mutator transaction binding the contract method 0x3402a036.
//
// Solidity: function updateAuction(address _auction) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) UpdateAuction(_auction common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.UpdateAuction(&_VolcanoERC721Factory.TransactOpts, _auction)
}

// UpdateAuction is a paid mutator transaction binding the contract method 0x3402a036.
//
// Solidity: function updateAuction(address _auction) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactorSession) UpdateAuction(_auction common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.UpdateAuction(&_VolcanoERC721Factory.TransactOpts, _auction)
}

// UpdateBundleMarketplace is a paid mutator transaction binding the contract method 0x9d0e24a7.
//
// Solidity: function updateBundleMarketplace(address _bundleMarketplace) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactor) UpdateBundleMarketplace(opts *bind.TransactOpts, _bundleMarketplace common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.contract.Transact(opts, "updateBundleMarketplace", _bundleMarketplace)
}

// UpdateBundleMarketplace is a paid mutator transaction binding the contract method 0x9d0e24a7.
//
// Solidity: function updateBundleMarketplace(address _bundleMarketplace) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) UpdateBundleMarketplace(_bundleMarketplace common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.UpdateBundleMarketplace(&_VolcanoERC721Factory.TransactOpts, _bundleMarketplace)
}

// UpdateBundleMarketplace is a paid mutator transaction binding the contract method 0x9d0e24a7.
//
// Solidity: function updateBundleMarketplace(address _bundleMarketplace) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactorSession) UpdateBundleMarketplace(_bundleMarketplace common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.UpdateBundleMarketplace(&_VolcanoERC721Factory.TransactOpts, _bundleMarketplace)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactor) UpdateFeeRecipient(opts *bind.TransactOpts, _feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.contract.Transact(opts, "updateFeeRecipient", _feeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) UpdateFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.UpdateFeeRecipient(&_VolcanoERC721Factory.TransactOpts, _feeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactorSession) UpdateFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.UpdateFeeRecipient(&_VolcanoERC721Factory.TransactOpts, _feeRecipient)
}

// UpdateMarketplace is a paid mutator transaction binding the contract method 0xeb636317.
//
// Solidity: function updateMarketplace(address _marketplace) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactor) UpdateMarketplace(opts *bind.TransactOpts, _marketplace common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.contract.Transact(opts, "updateMarketplace", _marketplace)
}

// UpdateMarketplace is a paid mutator transaction binding the contract method 0xeb636317.
//
// Solidity: function updateMarketplace(address _marketplace) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) UpdateMarketplace(_marketplace common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.UpdateMarketplace(&_VolcanoERC721Factory.TransactOpts, _marketplace)
}

// UpdateMarketplace is a paid mutator transaction binding the contract method 0xeb636317.
//
// Solidity: function updateMarketplace(address _marketplace) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactorSession) UpdateMarketplace(_marketplace common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.UpdateMarketplace(&_VolcanoERC721Factory.TransactOpts, _marketplace)
}

// UpdatePlatformContractFee is a paid mutator transaction binding the contract method 0xc57482bb.
//
// Solidity: function updatePlatformContractFee(uint256 _platformContractFee) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactor) UpdatePlatformContractFee(opts *bind.TransactOpts, _platformContractFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Factory.contract.Transact(opts, "updatePlatformContractFee", _platformContractFee)
}

// UpdatePlatformContractFee is a paid mutator transaction binding the contract method 0xc57482bb.
//
// Solidity: function updatePlatformContractFee(uint256 _platformContractFee) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) UpdatePlatformContractFee(_platformContractFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.UpdatePlatformContractFee(&_VolcanoERC721Factory.TransactOpts, _platformContractFee)
}

// UpdatePlatformContractFee is a paid mutator transaction binding the contract method 0xc57482bb.
//
// Solidity: function updatePlatformContractFee(uint256 _platformContractFee) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactorSession) UpdatePlatformContractFee(_platformContractFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.UpdatePlatformContractFee(&_VolcanoERC721Factory.TransactOpts, _platformContractFee)
}

// UpdatePlatformMintFee is a paid mutator transaction binding the contract method 0x61d2beae.
//
// Solidity: function updatePlatformMintFee(uint256 _platformMintFee) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactor) UpdatePlatformMintFee(opts *bind.TransactOpts, _platformMintFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Factory.contract.Transact(opts, "updatePlatformMintFee", _platformMintFee)
}

// UpdatePlatformMintFee is a paid mutator transaction binding the contract method 0x61d2beae.
//
// Solidity: function updatePlatformMintFee(uint256 _platformMintFee) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactorySession) UpdatePlatformMintFee(_platformMintFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.UpdatePlatformMintFee(&_VolcanoERC721Factory.TransactOpts, _platformMintFee)
}

// UpdatePlatformMintFee is a paid mutator transaction binding the contract method 0x61d2beae.
//
// Solidity: function updatePlatformMintFee(uint256 _platformMintFee) returns()
func (_VolcanoERC721Factory *VolcanoERC721FactoryTransactorSession) UpdatePlatformMintFee(_platformMintFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Factory.Contract.UpdatePlatformMintFee(&_VolcanoERC721Factory.TransactOpts, _platformMintFee)
}

// VolcanoERC721FactoryContractCreatedIterator is returned from FilterContractCreated and is used to iterate over the raw logs and unpacked data for ContractCreated events raised by the VolcanoERC721Factory contract.
type VolcanoERC721FactoryContractCreatedIterator struct {
	Event *VolcanoERC721FactoryContractCreated // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721FactoryContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721FactoryContractCreated)
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
		it.Event = new(VolcanoERC721FactoryContractCreated)
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
func (it *VolcanoERC721FactoryContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721FactoryContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721FactoryContractCreated represents a ContractCreated event raised by the VolcanoERC721Factory contract.
type VolcanoERC721FactoryContractCreated struct {
	Platform  common.Address
	Nft       common.Address
	Isprivate bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContractCreated is a free log retrieval operation binding the contract event 0xbbe78d8749296d5db350f8a76bbc6f2f85649a62516790be41dc6a67393aecde.
//
// Solidity: event ContractCreated(address platform, address nft, bool isprivate)
func (_VolcanoERC721Factory *VolcanoERC721FactoryFilterer) FilterContractCreated(opts *bind.FilterOpts) (*VolcanoERC721FactoryContractCreatedIterator, error) {

	logs, sub, err := _VolcanoERC721Factory.contract.FilterLogs(opts, "ContractCreated")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721FactoryContractCreatedIterator{contract: _VolcanoERC721Factory.contract, event: "ContractCreated", logs: logs, sub: sub}, nil
}

// WatchContractCreated is a free log subscription operation binding the contract event 0xbbe78d8749296d5db350f8a76bbc6f2f85649a62516790be41dc6a67393aecde.
//
// Solidity: event ContractCreated(address platform, address nft, bool isprivate)
func (_VolcanoERC721Factory *VolcanoERC721FactoryFilterer) WatchContractCreated(opts *bind.WatchOpts, sink chan<- *VolcanoERC721FactoryContractCreated) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC721Factory.contract.WatchLogs(opts, "ContractCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721FactoryContractCreated)
				if err := _VolcanoERC721Factory.contract.UnpackLog(event, "ContractCreated", log); err != nil {
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

// ParseContractCreated is a log parse operation binding the contract event 0xbbe78d8749296d5db350f8a76bbc6f2f85649a62516790be41dc6a67393aecde.
//
// Solidity: event ContractCreated(address platform, address nft, bool isprivate)
func (_VolcanoERC721Factory *VolcanoERC721FactoryFilterer) ParseContractCreated(log types.Log) (*VolcanoERC721FactoryContractCreated, error) {
	event := new(VolcanoERC721FactoryContractCreated)
	if err := _VolcanoERC721Factory.contract.UnpackLog(event, "ContractCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721FactoryContractDisabledIterator is returned from FilterContractDisabled and is used to iterate over the raw logs and unpacked data for ContractDisabled events raised by the VolcanoERC721Factory contract.
type VolcanoERC721FactoryContractDisabledIterator struct {
	Event *VolcanoERC721FactoryContractDisabled // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721FactoryContractDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721FactoryContractDisabled)
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
		it.Event = new(VolcanoERC721FactoryContractDisabled)
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
func (it *VolcanoERC721FactoryContractDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721FactoryContractDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721FactoryContractDisabled represents a ContractDisabled event raised by the VolcanoERC721Factory contract.
type VolcanoERC721FactoryContractDisabled struct {
	Caller common.Address
	Nft    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterContractDisabled is a free log retrieval operation binding the contract event 0x0d60ee093b891105dfb73abfbe0421cc1a07500f5c1776ec7d189e6c7f6d6340.
//
// Solidity: event ContractDisabled(address caller, address nft)
func (_VolcanoERC721Factory *VolcanoERC721FactoryFilterer) FilterContractDisabled(opts *bind.FilterOpts) (*VolcanoERC721FactoryContractDisabledIterator, error) {

	logs, sub, err := _VolcanoERC721Factory.contract.FilterLogs(opts, "ContractDisabled")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721FactoryContractDisabledIterator{contract: _VolcanoERC721Factory.contract, event: "ContractDisabled", logs: logs, sub: sub}, nil
}

// WatchContractDisabled is a free log subscription operation binding the contract event 0x0d60ee093b891105dfb73abfbe0421cc1a07500f5c1776ec7d189e6c7f6d6340.
//
// Solidity: event ContractDisabled(address caller, address nft)
func (_VolcanoERC721Factory *VolcanoERC721FactoryFilterer) WatchContractDisabled(opts *bind.WatchOpts, sink chan<- *VolcanoERC721FactoryContractDisabled) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC721Factory.contract.WatchLogs(opts, "ContractDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721FactoryContractDisabled)
				if err := _VolcanoERC721Factory.contract.UnpackLog(event, "ContractDisabled", log); err != nil {
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

// ParseContractDisabled is a log parse operation binding the contract event 0x0d60ee093b891105dfb73abfbe0421cc1a07500f5c1776ec7d189e6c7f6d6340.
//
// Solidity: event ContractDisabled(address caller, address nft)
func (_VolcanoERC721Factory *VolcanoERC721FactoryFilterer) ParseContractDisabled(log types.Log) (*VolcanoERC721FactoryContractDisabled, error) {
	event := new(VolcanoERC721FactoryContractDisabled)
	if err := _VolcanoERC721Factory.contract.UnpackLog(event, "ContractDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721FactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VolcanoERC721Factory contract.
type VolcanoERC721FactoryOwnershipTransferredIterator struct {
	Event *VolcanoERC721FactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721FactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721FactoryOwnershipTransferred)
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
		it.Event = new(VolcanoERC721FactoryOwnershipTransferred)
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
func (it *VolcanoERC721FactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721FactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721FactoryOwnershipTransferred represents a OwnershipTransferred event raised by the VolcanoERC721Factory contract.
type VolcanoERC721FactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC721Factory *VolcanoERC721FactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VolcanoERC721FactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC721Factory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721FactoryOwnershipTransferredIterator{contract: _VolcanoERC721Factory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC721Factory *VolcanoERC721FactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VolcanoERC721FactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC721Factory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721FactoryOwnershipTransferred)
				if err := _VolcanoERC721Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VolcanoERC721Factory *VolcanoERC721FactoryFilterer) ParseOwnershipTransferred(log types.Log) (*VolcanoERC721FactoryOwnershipTransferred, error) {
	event := new(VolcanoERC721FactoryOwnershipTransferred)
	if err := _VolcanoERC721Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
