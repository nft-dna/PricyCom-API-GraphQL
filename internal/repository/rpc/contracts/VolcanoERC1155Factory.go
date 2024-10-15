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

// VolcanoERC1155TradablecontractERC1155Options is an auto generated low-level Go binding around an user-defined struct.
type VolcanoERC1155TradablecontractERC1155Options struct {
	BaseUri        string
	UsebaseUriOnly bool
	UseDecimalUri  bool
	BaseUriExt     string
	MaxItems       *big.Int
	MaxItemSupply  *big.Int
	MintStartTime  *big.Int
	MintStopTime   *big.Int
	RevealTime     *big.Int
	PreRevealUri   string
	ContractUri    string
}

// VolcanoERC1155FactoryMetaData contains all meta data concerning the VolcanoERC1155Factory contract.
var VolcanoERC1155FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auction\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marketplace\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bundleMarketplace\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_platformContractFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_platformMintFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platform\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isprivate\",\"type\":\"bool\"}],\"name\":\"ContractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"}],\"name\":\"ContractDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bundleMarketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_private\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_mintFee\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"_creatorFeePerc\",\"type\":\"uint96\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"usebaseUriOnly\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDecimalUri\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"baseUriExt\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxItems\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxItemSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintStopTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"preRevealUri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contractUri\",\"type\":\"string\"}],\"internalType\":\"structVolcanoERC1155Tradable.contractERC1155Options\",\"name\":\"_options\",\"type\":\"tuple\"}],\"name\":\"createNFTContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"}],\"name\":\"disableTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isprivate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformContractFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformMintFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isprivate\",\"type\":\"bool\"}],\"name\":\"registerTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auction\",\"type\":\"address\"}],\"name\":\"updateAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bundleMarketplace\",\"type\":\"address\"}],\"name\":\"updateBundleMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"updateFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplace\",\"type\":\"address\"}],\"name\":\"updateMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformContractFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformContractFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformMintFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformMintFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VolcanoERC1155FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use VolcanoERC1155FactoryMetaData.ABI instead.
var VolcanoERC1155FactoryABI = VolcanoERC1155FactoryMetaData.ABI

// VolcanoERC1155Factory is an auto generated Go binding around an Ethereum contract.
type VolcanoERC1155Factory struct {
	VolcanoERC1155FactoryCaller     // Read-only binding to the contract
	VolcanoERC1155FactoryTransactor // Write-only binding to the contract
	VolcanoERC1155FactoryFilterer   // Log filterer for contract events
}

// VolcanoERC1155FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VolcanoERC1155FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC1155FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VolcanoERC1155FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC1155FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolcanoERC1155FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC1155FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolcanoERC1155FactorySession struct {
	Contract     *VolcanoERC1155Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VolcanoERC1155FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolcanoERC1155FactoryCallerSession struct {
	Contract *VolcanoERC1155FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// VolcanoERC1155FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolcanoERC1155FactoryTransactorSession struct {
	Contract     *VolcanoERC1155FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// VolcanoERC1155FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VolcanoERC1155FactoryRaw struct {
	Contract *VolcanoERC1155Factory // Generic contract binding to access the raw methods on
}

// VolcanoERC1155FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolcanoERC1155FactoryCallerRaw struct {
	Contract *VolcanoERC1155FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// VolcanoERC1155FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolcanoERC1155FactoryTransactorRaw struct {
	Contract *VolcanoERC1155FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolcanoERC1155Factory creates a new instance of VolcanoERC1155Factory, bound to a specific deployed contract.
func NewVolcanoERC1155Factory(address common.Address, backend bind.ContractBackend) (*VolcanoERC1155Factory, error) {
	contract, err := bindVolcanoERC1155Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155Factory{VolcanoERC1155FactoryCaller: VolcanoERC1155FactoryCaller{contract: contract}, VolcanoERC1155FactoryTransactor: VolcanoERC1155FactoryTransactor{contract: contract}, VolcanoERC1155FactoryFilterer: VolcanoERC1155FactoryFilterer{contract: contract}}, nil
}

// NewVolcanoERC1155FactoryCaller creates a new read-only instance of VolcanoERC1155Factory, bound to a specific deployed contract.
func NewVolcanoERC1155FactoryCaller(address common.Address, caller bind.ContractCaller) (*VolcanoERC1155FactoryCaller, error) {
	contract, err := bindVolcanoERC1155Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155FactoryCaller{contract: contract}, nil
}

// NewVolcanoERC1155FactoryTransactor creates a new write-only instance of VolcanoERC1155Factory, bound to a specific deployed contract.
func NewVolcanoERC1155FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*VolcanoERC1155FactoryTransactor, error) {
	contract, err := bindVolcanoERC1155Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155FactoryTransactor{contract: contract}, nil
}

// NewVolcanoERC1155FactoryFilterer creates a new log filterer instance of VolcanoERC1155Factory, bound to a specific deployed contract.
func NewVolcanoERC1155FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*VolcanoERC1155FactoryFilterer, error) {
	contract, err := bindVolcanoERC1155Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155FactoryFilterer{contract: contract}, nil
}

// bindVolcanoERC1155Factory binds a generic wrapper to an already deployed contract.
func bindVolcanoERC1155Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VolcanoERC1155FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC1155Factory.Contract.VolcanoERC1155FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.VolcanoERC1155FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.VolcanoERC1155FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC1155Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.contract.Transact(opts, method, params...)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCaller) Auction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC1155Factory.contract.Call(opts, &out, "auction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) Auction() (common.Address, error) {
	return _VolcanoERC1155Factory.Contract.Auction(&_VolcanoERC1155Factory.CallOpts)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCallerSession) Auction() (common.Address, error) {
	return _VolcanoERC1155Factory.Contract.Auction(&_VolcanoERC1155Factory.CallOpts)
}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCaller) BundleMarketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC1155Factory.contract.Call(opts, &out, "bundleMarketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) BundleMarketplace() (common.Address, error) {
	return _VolcanoERC1155Factory.Contract.BundleMarketplace(&_VolcanoERC1155Factory.CallOpts)
}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCallerSession) BundleMarketplace() (common.Address, error) {
	return _VolcanoERC1155Factory.Contract.BundleMarketplace(&_VolcanoERC1155Factory.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCaller) Exists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VolcanoERC1155Factory.contract.Call(opts, &out, "exists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) Exists(arg0 common.Address) (bool, error) {
	return _VolcanoERC1155Factory.Contract.Exists(&_VolcanoERC1155Factory.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCallerSession) Exists(arg0 common.Address) (bool, error) {
	return _VolcanoERC1155Factory.Contract.Exists(&_VolcanoERC1155Factory.CallOpts, arg0)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC1155Factory.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) FeeRecipient() (common.Address, error) {
	return _VolcanoERC1155Factory.Contract.FeeRecipient(&_VolcanoERC1155Factory.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCallerSession) FeeRecipient() (common.Address, error) {
	return _VolcanoERC1155Factory.Contract.FeeRecipient(&_VolcanoERC1155Factory.CallOpts)
}

// Isprivate is a free data retrieval call binding the contract method 0x50ac3f47.
//
// Solidity: function isprivate(address ) view returns(bool)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCaller) Isprivate(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VolcanoERC1155Factory.contract.Call(opts, &out, "isprivate", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Isprivate is a free data retrieval call binding the contract method 0x50ac3f47.
//
// Solidity: function isprivate(address ) view returns(bool)
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) Isprivate(arg0 common.Address) (bool, error) {
	return _VolcanoERC1155Factory.Contract.Isprivate(&_VolcanoERC1155Factory.CallOpts, arg0)
}

// Isprivate is a free data retrieval call binding the contract method 0x50ac3f47.
//
// Solidity: function isprivate(address ) view returns(bool)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCallerSession) Isprivate(arg0 common.Address) (bool, error) {
	return _VolcanoERC1155Factory.Contract.Isprivate(&_VolcanoERC1155Factory.CallOpts, arg0)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCaller) Marketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC1155Factory.contract.Call(opts, &out, "marketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) Marketplace() (common.Address, error) {
	return _VolcanoERC1155Factory.Contract.Marketplace(&_VolcanoERC1155Factory.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCallerSession) Marketplace() (common.Address, error) {
	return _VolcanoERC1155Factory.Contract.Marketplace(&_VolcanoERC1155Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC1155Factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) Owner() (common.Address, error) {
	return _VolcanoERC1155Factory.Contract.Owner(&_VolcanoERC1155Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCallerSession) Owner() (common.Address, error) {
	return _VolcanoERC1155Factory.Contract.Owner(&_VolcanoERC1155Factory.CallOpts)
}

// PlatformContractFee is a free data retrieval call binding the contract method 0xb69d4b98.
//
// Solidity: function platformContractFee() view returns(uint256)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCaller) PlatformContractFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC1155Factory.contract.Call(opts, &out, "platformContractFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformContractFee is a free data retrieval call binding the contract method 0xb69d4b98.
//
// Solidity: function platformContractFee() view returns(uint256)
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) PlatformContractFee() (*big.Int, error) {
	return _VolcanoERC1155Factory.Contract.PlatformContractFee(&_VolcanoERC1155Factory.CallOpts)
}

// PlatformContractFee is a free data retrieval call binding the contract method 0xb69d4b98.
//
// Solidity: function platformContractFee() view returns(uint256)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCallerSession) PlatformContractFee() (*big.Int, error) {
	return _VolcanoERC1155Factory.Contract.PlatformContractFee(&_VolcanoERC1155Factory.CallOpts)
}

// PlatformMintFee is a free data retrieval call binding the contract method 0xce922b7b.
//
// Solidity: function platformMintFee() view returns(uint256)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCaller) PlatformMintFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC1155Factory.contract.Call(opts, &out, "platformMintFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformMintFee is a free data retrieval call binding the contract method 0xce922b7b.
//
// Solidity: function platformMintFee() view returns(uint256)
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) PlatformMintFee() (*big.Int, error) {
	return _VolcanoERC1155Factory.Contract.PlatformMintFee(&_VolcanoERC1155Factory.CallOpts)
}

// PlatformMintFee is a free data retrieval call binding the contract method 0xce922b7b.
//
// Solidity: function platformMintFee() view returns(uint256)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryCallerSession) PlatformMintFee() (*big.Int, error) {
	return _VolcanoERC1155Factory.Contract.PlatformMintFee(&_VolcanoERC1155Factory.CallOpts)
}

// CreateNFTContract is a paid mutator transaction binding the contract method 0x12ad6637.
//
// Solidity: function createNFTContract(string _name, string _symbol, bool _private, uint256 _mintFee, uint96 _creatorFeePerc, address _feeRecipient, (string,bool,bool,string,uint256,uint256,uint256,uint256,uint256,string,string) _options) payable returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactor) CreateNFTContract(opts *bind.TransactOpts, _name string, _symbol string, _private bool, _mintFee *big.Int, _creatorFeePerc *big.Int, _feeRecipient common.Address, _options VolcanoERC1155TradablecontractERC1155Options) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.contract.Transact(opts, "createNFTContract", _name, _symbol, _private, _mintFee, _creatorFeePerc, _feeRecipient, _options)
}

// CreateNFTContract is a paid mutator transaction binding the contract method 0x12ad6637.
//
// Solidity: function createNFTContract(string _name, string _symbol, bool _private, uint256 _mintFee, uint96 _creatorFeePerc, address _feeRecipient, (string,bool,bool,string,uint256,uint256,uint256,uint256,uint256,string,string) _options) payable returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) CreateNFTContract(_name string, _symbol string, _private bool, _mintFee *big.Int, _creatorFeePerc *big.Int, _feeRecipient common.Address, _options VolcanoERC1155TradablecontractERC1155Options) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.CreateNFTContract(&_VolcanoERC1155Factory.TransactOpts, _name, _symbol, _private, _mintFee, _creatorFeePerc, _feeRecipient, _options)
}

// CreateNFTContract is a paid mutator transaction binding the contract method 0x12ad6637.
//
// Solidity: function createNFTContract(string _name, string _symbol, bool _private, uint256 _mintFee, uint96 _creatorFeePerc, address _feeRecipient, (string,bool,bool,string,uint256,uint256,uint256,uint256,uint256,string,string) _options) payable returns(address)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactorSession) CreateNFTContract(_name string, _symbol string, _private bool, _mintFee *big.Int, _creatorFeePerc *big.Int, _feeRecipient common.Address, _options VolcanoERC1155TradablecontractERC1155Options) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.CreateNFTContract(&_VolcanoERC1155Factory.TransactOpts, _name, _symbol, _private, _mintFee, _creatorFeePerc, _feeRecipient, _options)
}

// DisableTokenContract is a paid mutator transaction binding the contract method 0xd39bb16a.
//
// Solidity: function disableTokenContract(address tokenContractAddress) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactor) DisableTokenContract(opts *bind.TransactOpts, tokenContractAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.contract.Transact(opts, "disableTokenContract", tokenContractAddress)
}

// DisableTokenContract is a paid mutator transaction binding the contract method 0xd39bb16a.
//
// Solidity: function disableTokenContract(address tokenContractAddress) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) DisableTokenContract(tokenContractAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.DisableTokenContract(&_VolcanoERC1155Factory.TransactOpts, tokenContractAddress)
}

// DisableTokenContract is a paid mutator transaction binding the contract method 0xd39bb16a.
//
// Solidity: function disableTokenContract(address tokenContractAddress) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactorSession) DisableTokenContract(tokenContractAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.DisableTokenContract(&_VolcanoERC1155Factory.TransactOpts, tokenContractAddress)
}

// RegisterTokenContract is a paid mutator transaction binding the contract method 0x9c7d1a4e.
//
// Solidity: function registerTokenContract(address tokenContractAddress, bool _isprivate) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactor) RegisterTokenContract(opts *bind.TransactOpts, tokenContractAddress common.Address, _isprivate bool) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.contract.Transact(opts, "registerTokenContract", tokenContractAddress, _isprivate)
}

// RegisterTokenContract is a paid mutator transaction binding the contract method 0x9c7d1a4e.
//
// Solidity: function registerTokenContract(address tokenContractAddress, bool _isprivate) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) RegisterTokenContract(tokenContractAddress common.Address, _isprivate bool) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.RegisterTokenContract(&_VolcanoERC1155Factory.TransactOpts, tokenContractAddress, _isprivate)
}

// RegisterTokenContract is a paid mutator transaction binding the contract method 0x9c7d1a4e.
//
// Solidity: function registerTokenContract(address tokenContractAddress, bool _isprivate) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactorSession) RegisterTokenContract(tokenContractAddress common.Address, _isprivate bool) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.RegisterTokenContract(&_VolcanoERC1155Factory.TransactOpts, tokenContractAddress, _isprivate)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.RenounceOwnership(&_VolcanoERC1155Factory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.RenounceOwnership(&_VolcanoERC1155Factory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.TransferOwnership(&_VolcanoERC1155Factory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.TransferOwnership(&_VolcanoERC1155Factory.TransactOpts, newOwner)
}

// UpdateAuction is a paid mutator transaction binding the contract method 0x3402a036.
//
// Solidity: function updateAuction(address _auction) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactor) UpdateAuction(opts *bind.TransactOpts, _auction common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.contract.Transact(opts, "updateAuction", _auction)
}

// UpdateAuction is a paid mutator transaction binding the contract method 0x3402a036.
//
// Solidity: function updateAuction(address _auction) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) UpdateAuction(_auction common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.UpdateAuction(&_VolcanoERC1155Factory.TransactOpts, _auction)
}

// UpdateAuction is a paid mutator transaction binding the contract method 0x3402a036.
//
// Solidity: function updateAuction(address _auction) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactorSession) UpdateAuction(_auction common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.UpdateAuction(&_VolcanoERC1155Factory.TransactOpts, _auction)
}

// UpdateBundleMarketplace is a paid mutator transaction binding the contract method 0x9d0e24a7.
//
// Solidity: function updateBundleMarketplace(address _bundleMarketplace) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactor) UpdateBundleMarketplace(opts *bind.TransactOpts, _bundleMarketplace common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.contract.Transact(opts, "updateBundleMarketplace", _bundleMarketplace)
}

// UpdateBundleMarketplace is a paid mutator transaction binding the contract method 0x9d0e24a7.
//
// Solidity: function updateBundleMarketplace(address _bundleMarketplace) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) UpdateBundleMarketplace(_bundleMarketplace common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.UpdateBundleMarketplace(&_VolcanoERC1155Factory.TransactOpts, _bundleMarketplace)
}

// UpdateBundleMarketplace is a paid mutator transaction binding the contract method 0x9d0e24a7.
//
// Solidity: function updateBundleMarketplace(address _bundleMarketplace) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactorSession) UpdateBundleMarketplace(_bundleMarketplace common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.UpdateBundleMarketplace(&_VolcanoERC1155Factory.TransactOpts, _bundleMarketplace)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactor) UpdateFeeRecipient(opts *bind.TransactOpts, _feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.contract.Transact(opts, "updateFeeRecipient", _feeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) UpdateFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.UpdateFeeRecipient(&_VolcanoERC1155Factory.TransactOpts, _feeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactorSession) UpdateFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.UpdateFeeRecipient(&_VolcanoERC1155Factory.TransactOpts, _feeRecipient)
}

// UpdateMarketplace is a paid mutator transaction binding the contract method 0xeb636317.
//
// Solidity: function updateMarketplace(address _marketplace) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactor) UpdateMarketplace(opts *bind.TransactOpts, _marketplace common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.contract.Transact(opts, "updateMarketplace", _marketplace)
}

// UpdateMarketplace is a paid mutator transaction binding the contract method 0xeb636317.
//
// Solidity: function updateMarketplace(address _marketplace) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) UpdateMarketplace(_marketplace common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.UpdateMarketplace(&_VolcanoERC1155Factory.TransactOpts, _marketplace)
}

// UpdateMarketplace is a paid mutator transaction binding the contract method 0xeb636317.
//
// Solidity: function updateMarketplace(address _marketplace) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactorSession) UpdateMarketplace(_marketplace common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.UpdateMarketplace(&_VolcanoERC1155Factory.TransactOpts, _marketplace)
}

// UpdatePlatformContractFee is a paid mutator transaction binding the contract method 0xc57482bb.
//
// Solidity: function updatePlatformContractFee(uint256 _platformContractFee) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactor) UpdatePlatformContractFee(opts *bind.TransactOpts, _platformContractFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.contract.Transact(opts, "updatePlatformContractFee", _platformContractFee)
}

// UpdatePlatformContractFee is a paid mutator transaction binding the contract method 0xc57482bb.
//
// Solidity: function updatePlatformContractFee(uint256 _platformContractFee) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) UpdatePlatformContractFee(_platformContractFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.UpdatePlatformContractFee(&_VolcanoERC1155Factory.TransactOpts, _platformContractFee)
}

// UpdatePlatformContractFee is a paid mutator transaction binding the contract method 0xc57482bb.
//
// Solidity: function updatePlatformContractFee(uint256 _platformContractFee) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactorSession) UpdatePlatformContractFee(_platformContractFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.UpdatePlatformContractFee(&_VolcanoERC1155Factory.TransactOpts, _platformContractFee)
}

// UpdatePlatformMintFee is a paid mutator transaction binding the contract method 0x61d2beae.
//
// Solidity: function updatePlatformMintFee(uint256 _platformMintFee) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactor) UpdatePlatformMintFee(opts *bind.TransactOpts, _platformMintFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.contract.Transact(opts, "updatePlatformMintFee", _platformMintFee)
}

// UpdatePlatformMintFee is a paid mutator transaction binding the contract method 0x61d2beae.
//
// Solidity: function updatePlatformMintFee(uint256 _platformMintFee) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactorySession) UpdatePlatformMintFee(_platformMintFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.UpdatePlatformMintFee(&_VolcanoERC1155Factory.TransactOpts, _platformMintFee)
}

// UpdatePlatformMintFee is a paid mutator transaction binding the contract method 0x61d2beae.
//
// Solidity: function updatePlatformMintFee(uint256 _platformMintFee) returns()
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryTransactorSession) UpdatePlatformMintFee(_platformMintFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC1155Factory.Contract.UpdatePlatformMintFee(&_VolcanoERC1155Factory.TransactOpts, _platformMintFee)
}

// VolcanoERC1155FactoryContractCreatedIterator is returned from FilterContractCreated and is used to iterate over the raw logs and unpacked data for ContractCreated events raised by the VolcanoERC1155Factory contract.
type VolcanoERC1155FactoryContractCreatedIterator struct {
	Event *VolcanoERC1155FactoryContractCreated // Event containing the contract specifics and raw log

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
func (it *VolcanoERC1155FactoryContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC1155FactoryContractCreated)
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
		it.Event = new(VolcanoERC1155FactoryContractCreated)
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
func (it *VolcanoERC1155FactoryContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC1155FactoryContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC1155FactoryContractCreated represents a ContractCreated event raised by the VolcanoERC1155Factory contract.
type VolcanoERC1155FactoryContractCreated struct {
	Platform  common.Address
	Nft       common.Address
	Isprivate bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContractCreated is a free log retrieval operation binding the contract event 0xbbe78d8749296d5db350f8a76bbc6f2f85649a62516790be41dc6a67393aecde.
//
// Solidity: event ContractCreated(address platform, address nft, bool isprivate)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryFilterer) FilterContractCreated(opts *bind.FilterOpts) (*VolcanoERC1155FactoryContractCreatedIterator, error) {

	logs, sub, err := _VolcanoERC1155Factory.contract.FilterLogs(opts, "ContractCreated")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155FactoryContractCreatedIterator{contract: _VolcanoERC1155Factory.contract, event: "ContractCreated", logs: logs, sub: sub}, nil
}

// WatchContractCreated is a free log subscription operation binding the contract event 0xbbe78d8749296d5db350f8a76bbc6f2f85649a62516790be41dc6a67393aecde.
//
// Solidity: event ContractCreated(address platform, address nft, bool isprivate)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryFilterer) WatchContractCreated(opts *bind.WatchOpts, sink chan<- *VolcanoERC1155FactoryContractCreated) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC1155Factory.contract.WatchLogs(opts, "ContractCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC1155FactoryContractCreated)
				if err := _VolcanoERC1155Factory.contract.UnpackLog(event, "ContractCreated", log); err != nil {
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
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryFilterer) ParseContractCreated(log types.Log) (*VolcanoERC1155FactoryContractCreated, error) {
	event := new(VolcanoERC1155FactoryContractCreated)
	if err := _VolcanoERC1155Factory.contract.UnpackLog(event, "ContractCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC1155FactoryContractDisabledIterator is returned from FilterContractDisabled and is used to iterate over the raw logs and unpacked data for ContractDisabled events raised by the VolcanoERC1155Factory contract.
type VolcanoERC1155FactoryContractDisabledIterator struct {
	Event *VolcanoERC1155FactoryContractDisabled // Event containing the contract specifics and raw log

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
func (it *VolcanoERC1155FactoryContractDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC1155FactoryContractDisabled)
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
		it.Event = new(VolcanoERC1155FactoryContractDisabled)
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
func (it *VolcanoERC1155FactoryContractDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC1155FactoryContractDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC1155FactoryContractDisabled represents a ContractDisabled event raised by the VolcanoERC1155Factory contract.
type VolcanoERC1155FactoryContractDisabled struct {
	Caller common.Address
	Nft    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterContractDisabled is a free log retrieval operation binding the contract event 0x0d60ee093b891105dfb73abfbe0421cc1a07500f5c1776ec7d189e6c7f6d6340.
//
// Solidity: event ContractDisabled(address caller, address nft)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryFilterer) FilterContractDisabled(opts *bind.FilterOpts) (*VolcanoERC1155FactoryContractDisabledIterator, error) {

	logs, sub, err := _VolcanoERC1155Factory.contract.FilterLogs(opts, "ContractDisabled")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155FactoryContractDisabledIterator{contract: _VolcanoERC1155Factory.contract, event: "ContractDisabled", logs: logs, sub: sub}, nil
}

// WatchContractDisabled is a free log subscription operation binding the contract event 0x0d60ee093b891105dfb73abfbe0421cc1a07500f5c1776ec7d189e6c7f6d6340.
//
// Solidity: event ContractDisabled(address caller, address nft)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryFilterer) WatchContractDisabled(opts *bind.WatchOpts, sink chan<- *VolcanoERC1155FactoryContractDisabled) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC1155Factory.contract.WatchLogs(opts, "ContractDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC1155FactoryContractDisabled)
				if err := _VolcanoERC1155Factory.contract.UnpackLog(event, "ContractDisabled", log); err != nil {
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
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryFilterer) ParseContractDisabled(log types.Log) (*VolcanoERC1155FactoryContractDisabled, error) {
	event := new(VolcanoERC1155FactoryContractDisabled)
	if err := _VolcanoERC1155Factory.contract.UnpackLog(event, "ContractDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC1155FactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VolcanoERC1155Factory contract.
type VolcanoERC1155FactoryOwnershipTransferredIterator struct {
	Event *VolcanoERC1155FactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VolcanoERC1155FactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC1155FactoryOwnershipTransferred)
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
		it.Event = new(VolcanoERC1155FactoryOwnershipTransferred)
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
func (it *VolcanoERC1155FactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC1155FactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC1155FactoryOwnershipTransferred represents a OwnershipTransferred event raised by the VolcanoERC1155Factory contract.
type VolcanoERC1155FactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VolcanoERC1155FactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC1155Factory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155FactoryOwnershipTransferredIterator{contract: _VolcanoERC1155Factory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VolcanoERC1155FactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC1155Factory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC1155FactoryOwnershipTransferred)
				if err := _VolcanoERC1155Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VolcanoERC1155Factory *VolcanoERC1155FactoryFilterer) ParseOwnershipTransferred(log types.Log) (*VolcanoERC1155FactoryOwnershipTransferred, error) {
	event := new(VolcanoERC1155FactoryOwnershipTransferred)
	if err := _VolcanoERC1155Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
