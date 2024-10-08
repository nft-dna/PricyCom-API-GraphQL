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
}

// VolcanoERC721TradableMetaData contains all meta data concerning the VolcanoERC721Tradable contract.
var VolcanoERC721TradableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_auction\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marketplace\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bundleMarketplace\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintCreatorFee\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"_creatorFeePerc\",\"type\":\"uint96\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isprivate\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"useDecimalUri\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"baseUriExt\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxItems\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintStopTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"preRevealUri\",\"type\":\"string\"}],\"internalType\":\"structVolcanoERC721Tradable.contractERC721Options\",\"name\":\"_options\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"creatorFeePerc\",\"type\":\"uint96\"}],\"name\":\"UpdateCreatorFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"UpdateFeeRecipient\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creatorFeePerc\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isprivate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintCreatorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintStopTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revealTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"toHexString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"updateContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_creatorFeePerc\",\"type\":\"uint96\"}],\"name\":\"updateCreatorFeePerc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"updateFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useBaseUri\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VolcanoERC721TradableABI is the input ABI used to generate the binding from.
// Deprecated: Use VolcanoERC721TradableMetaData.ABI instead.
var VolcanoERC721TradableABI = VolcanoERC721TradableMetaData.ABI

// VolcanoERC721Tradable is an auto generated Go binding around an Ethereum contract.
type VolcanoERC721Tradable struct {
	VolcanoERC721TradableCaller     // Read-only binding to the contract
	VolcanoERC721TradableTransactor // Write-only binding to the contract
	VolcanoERC721TradableFilterer   // Log filterer for contract events
}

// VolcanoERC721TradableCaller is an auto generated read-only Go binding around an Ethereum contract.
type VolcanoERC721TradableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC721TradableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VolcanoERC721TradableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC721TradableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolcanoERC721TradableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC721TradableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolcanoERC721TradableSession struct {
	Contract     *VolcanoERC721Tradable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VolcanoERC721TradableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolcanoERC721TradableCallerSession struct {
	Contract *VolcanoERC721TradableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// VolcanoERC721TradableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolcanoERC721TradableTransactorSession struct {
	Contract     *VolcanoERC721TradableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// VolcanoERC721TradableRaw is an auto generated low-level Go binding around an Ethereum contract.
type VolcanoERC721TradableRaw struct {
	Contract *VolcanoERC721Tradable // Generic contract binding to access the raw methods on
}

// VolcanoERC721TradableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolcanoERC721TradableCallerRaw struct {
	Contract *VolcanoERC721TradableCaller // Generic read-only contract binding to access the raw methods on
}

// VolcanoERC721TradableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolcanoERC721TradableTransactorRaw struct {
	Contract *VolcanoERC721TradableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolcanoERC721Tradable creates a new instance of VolcanoERC721Tradable, bound to a specific deployed contract.
func NewVolcanoERC721Tradable(address common.Address, backend bind.ContractBackend) (*VolcanoERC721Tradable, error) {
	contract, err := bindVolcanoERC721Tradable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721Tradable{VolcanoERC721TradableCaller: VolcanoERC721TradableCaller{contract: contract}, VolcanoERC721TradableTransactor: VolcanoERC721TradableTransactor{contract: contract}, VolcanoERC721TradableFilterer: VolcanoERC721TradableFilterer{contract: contract}}, nil
}

// NewVolcanoERC721TradableCaller creates a new read-only instance of VolcanoERC721Tradable, bound to a specific deployed contract.
func NewVolcanoERC721TradableCaller(address common.Address, caller bind.ContractCaller) (*VolcanoERC721TradableCaller, error) {
	contract, err := bindVolcanoERC721Tradable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721TradableCaller{contract: contract}, nil
}

// NewVolcanoERC721TradableTransactor creates a new write-only instance of VolcanoERC721Tradable, bound to a specific deployed contract.
func NewVolcanoERC721TradableTransactor(address common.Address, transactor bind.ContractTransactor) (*VolcanoERC721TradableTransactor, error) {
	contract, err := bindVolcanoERC721Tradable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721TradableTransactor{contract: contract}, nil
}

// NewVolcanoERC721TradableFilterer creates a new log filterer instance of VolcanoERC721Tradable, bound to a specific deployed contract.
func NewVolcanoERC721TradableFilterer(address common.Address, filterer bind.ContractFilterer) (*VolcanoERC721TradableFilterer, error) {
	contract, err := bindVolcanoERC721Tradable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721TradableFilterer{contract: contract}, nil
}

// bindVolcanoERC721Tradable binds a generic wrapper to an already deployed contract.
func bindVolcanoERC721Tradable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VolcanoERC721TradableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC721Tradable *VolcanoERC721TradableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC721Tradable.Contract.VolcanoERC721TradableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC721Tradable *VolcanoERC721TradableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.VolcanoERC721TradableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC721Tradable *VolcanoERC721TradableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.VolcanoERC721TradableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC721Tradable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.BalanceOf(&_VolcanoERC721Tradable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.BalanceOf(&_VolcanoERC721Tradable.CallOpts, owner)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) ContractURI() (string, error) {
	return _VolcanoERC721Tradable.Contract.ContractURI(&_VolcanoERC721Tradable.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) ContractURI() (string, error) {
	return _VolcanoERC721Tradable.Contract.ContractURI(&_VolcanoERC721Tradable.CallOpts)
}

// CreatorFeePerc is a free data retrieval call binding the contract method 0xa7050d63.
//
// Solidity: function creatorFeePerc() view returns(uint96)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) CreatorFeePerc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "creatorFeePerc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatorFeePerc is a free data retrieval call binding the contract method 0xa7050d63.
//
// Solidity: function creatorFeePerc() view returns(uint96)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) CreatorFeePerc() (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.CreatorFeePerc(&_VolcanoERC721Tradable.CallOpts)
}

// CreatorFeePerc is a free data retrieval call binding the contract method 0xa7050d63.
//
// Solidity: function creatorFeePerc() view returns(uint96)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) CreatorFeePerc() (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.CreatorFeePerc(&_VolcanoERC721Tradable.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "exists", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) Exists(_tokenId *big.Int) (bool, error) {
	return _VolcanoERC721Tradable.Contract.Exists(&_VolcanoERC721Tradable.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _VolcanoERC721Tradable.Contract.Exists(&_VolcanoERC721Tradable.CallOpts, _tokenId)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) Factory() (common.Address, error) {
	return _VolcanoERC721Tradable.Contract.Factory(&_VolcanoERC721Tradable.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) Factory() (common.Address, error) {
	return _VolcanoERC721Tradable.Contract.Factory(&_VolcanoERC721Tradable.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) FeeRecipient() (common.Address, error) {
	return _VolcanoERC721Tradable.Contract.FeeRecipient(&_VolcanoERC721Tradable.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) FeeRecipient() (common.Address, error) {
	return _VolcanoERC721Tradable.Contract.FeeRecipient(&_VolcanoERC721Tradable.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _VolcanoERC721Tradable.Contract.GetApproved(&_VolcanoERC721Tradable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _VolcanoERC721Tradable.Contract.GetApproved(&_VolcanoERC721Tradable.CallOpts, tokenId)
}

// IsApproved is a free data retrieval call binding the contract method 0x56c31637.
//
// Solidity: function isApproved(uint256 _tokenId, address _operator) view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) IsApproved(opts *bind.CallOpts, _tokenId *big.Int, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "isApproved", _tokenId, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x56c31637.
//
// Solidity: function isApproved(uint256 _tokenId, address _operator) view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) IsApproved(_tokenId *big.Int, _operator common.Address) (bool, error) {
	return _VolcanoERC721Tradable.Contract.IsApproved(&_VolcanoERC721Tradable.CallOpts, _tokenId, _operator)
}

// IsApproved is a free data retrieval call binding the contract method 0x56c31637.
//
// Solidity: function isApproved(uint256 _tokenId, address _operator) view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) IsApproved(_tokenId *big.Int, _operator common.Address) (bool, error) {
	return _VolcanoERC721Tradable.Contract.IsApproved(&_VolcanoERC721Tradable.CallOpts, _tokenId, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _VolcanoERC721Tradable.Contract.IsApprovedForAll(&_VolcanoERC721Tradable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _VolcanoERC721Tradable.Contract.IsApprovedForAll(&_VolcanoERC721Tradable.CallOpts, owner, operator)
}

// Isprivate is a free data retrieval call binding the contract method 0xfcbbbd1a.
//
// Solidity: function isprivate() view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) Isprivate(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "isprivate")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Isprivate is a free data retrieval call binding the contract method 0xfcbbbd1a.
//
// Solidity: function isprivate() view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) Isprivate() (bool, error) {
	return _VolcanoERC721Tradable.Contract.Isprivate(&_VolcanoERC721Tradable.CallOpts)
}

// Isprivate is a free data retrieval call binding the contract method 0xfcbbbd1a.
//
// Solidity: function isprivate() view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) Isprivate() (bool, error) {
	return _VolcanoERC721Tradable.Contract.Isprivate(&_VolcanoERC721Tradable.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) MaxSupply() (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.MaxSupply(&_VolcanoERC721Tradable.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) MaxSupply() (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.MaxSupply(&_VolcanoERC721Tradable.CallOpts)
}

// MintCreatorFee is a free data retrieval call binding the contract method 0x9237d5a5.
//
// Solidity: function mintCreatorFee() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) MintCreatorFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "mintCreatorFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintCreatorFee is a free data retrieval call binding the contract method 0x9237d5a5.
//
// Solidity: function mintCreatorFee() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) MintCreatorFee() (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.MintCreatorFee(&_VolcanoERC721Tradable.CallOpts)
}

// MintCreatorFee is a free data retrieval call binding the contract method 0x9237d5a5.
//
// Solidity: function mintCreatorFee() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) MintCreatorFee() (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.MintCreatorFee(&_VolcanoERC721Tradable.CallOpts)
}

// MintStartTime is a free data retrieval call binding the contract method 0x931e2e49.
//
// Solidity: function mintStartTime() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) MintStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "mintStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintStartTime is a free data retrieval call binding the contract method 0x931e2e49.
//
// Solidity: function mintStartTime() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) MintStartTime() (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.MintStartTime(&_VolcanoERC721Tradable.CallOpts)
}

// MintStartTime is a free data retrieval call binding the contract method 0x931e2e49.
//
// Solidity: function mintStartTime() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) MintStartTime() (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.MintStartTime(&_VolcanoERC721Tradable.CallOpts)
}

// MintStopTime is a free data retrieval call binding the contract method 0x95c8b27a.
//
// Solidity: function mintStopTime() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) MintStopTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "mintStopTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintStopTime is a free data retrieval call binding the contract method 0x95c8b27a.
//
// Solidity: function mintStopTime() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) MintStopTime() (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.MintStopTime(&_VolcanoERC721Tradable.CallOpts)
}

// MintStopTime is a free data retrieval call binding the contract method 0x95c8b27a.
//
// Solidity: function mintStopTime() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) MintStopTime() (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.MintStopTime(&_VolcanoERC721Tradable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) Name() (string, error) {
	return _VolcanoERC721Tradable.Contract.Name(&_VolcanoERC721Tradable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) Name() (string, error) {
	return _VolcanoERC721Tradable.Contract.Name(&_VolcanoERC721Tradable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) Owner() (common.Address, error) {
	return _VolcanoERC721Tradable.Contract.Owner(&_VolcanoERC721Tradable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) Owner() (common.Address, error) {
	return _VolcanoERC721Tradable.Contract.Owner(&_VolcanoERC721Tradable.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _VolcanoERC721Tradable.Contract.OwnerOf(&_VolcanoERC721Tradable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _VolcanoERC721Tradable.Contract.OwnerOf(&_VolcanoERC721Tradable.CallOpts, tokenId)
}

// RevealTime is a free data retrieval call binding the contract method 0xba829d71.
//
// Solidity: function revealTime() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) RevealTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "revealTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevealTime is a free data retrieval call binding the contract method 0xba829d71.
//
// Solidity: function revealTime() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) RevealTime() (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.RevealTime(&_VolcanoERC721Tradable.CallOpts)
}

// RevealTime is a free data retrieval call binding the contract method 0xba829d71.
//
// Solidity: function revealTime() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) RevealTime() (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.RevealTime(&_VolcanoERC721Tradable.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _VolcanoERC721Tradable.Contract.RoyaltyInfo(&_VolcanoERC721Tradable.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _VolcanoERC721Tradable.Contract.RoyaltyInfo(&_VolcanoERC721Tradable.CallOpts, tokenId, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VolcanoERC721Tradable.Contract.SupportsInterface(&_VolcanoERC721Tradable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VolcanoERC721Tradable.Contract.SupportsInterface(&_VolcanoERC721Tradable.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) Symbol() (string, error) {
	return _VolcanoERC721Tradable.Contract.Symbol(&_VolcanoERC721Tradable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) Symbol() (string, error) {
	return _VolcanoERC721Tradable.Contract.Symbol(&_VolcanoERC721Tradable.CallOpts)
}

// ToHexString is a free data retrieval call binding the contract method 0x63e1cbea.
//
// Solidity: function toHexString(uint256 value, uint256 length) pure returns(string)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) ToHexString(opts *bind.CallOpts, value *big.Int, length *big.Int) (string, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "toHexString", value, length)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToHexString is a free data retrieval call binding the contract method 0x63e1cbea.
//
// Solidity: function toHexString(uint256 value, uint256 length) pure returns(string)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) ToHexString(value *big.Int, length *big.Int) (string, error) {
	return _VolcanoERC721Tradable.Contract.ToHexString(&_VolcanoERC721Tradable.CallOpts, value, length)
}

// ToHexString is a free data retrieval call binding the contract method 0x63e1cbea.
//
// Solidity: function toHexString(uint256 value, uint256 length) pure returns(string)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) ToHexString(value *big.Int, length *big.Int) (string, error) {
	return _VolcanoERC721Tradable.Contract.ToHexString(&_VolcanoERC721Tradable.CallOpts, value, length)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.TokenByIndex(&_VolcanoERC721Tradable.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.TokenByIndex(&_VolcanoERC721Tradable.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.TokenOfOwnerByIndex(&_VolcanoERC721Tradable.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.TokenOfOwnerByIndex(&_VolcanoERC721Tradable.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) TokenURI(tokenId *big.Int) (string, error) {
	return _VolcanoERC721Tradable.Contract.TokenURI(&_VolcanoERC721Tradable.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _VolcanoERC721Tradable.Contract.TokenURI(&_VolcanoERC721Tradable.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) TotalSupply() (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.TotalSupply(&_VolcanoERC721Tradable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) TotalSupply() (*big.Int, error) {
	return _VolcanoERC721Tradable.Contract.TotalSupply(&_VolcanoERC721Tradable.CallOpts)
}

// UseBaseUri is a free data retrieval call binding the contract method 0xa81d9b36.
//
// Solidity: function useBaseUri() view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCaller) UseBaseUri(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VolcanoERC721Tradable.contract.Call(opts, &out, "useBaseUri")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseBaseUri is a free data retrieval call binding the contract method 0xa81d9b36.
//
// Solidity: function useBaseUri() view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) UseBaseUri() (bool, error) {
	return _VolcanoERC721Tradable.Contract.UseBaseUri(&_VolcanoERC721Tradable.CallOpts)
}

// UseBaseUri is a free data retrieval call binding the contract method 0xa81d9b36.
//
// Solidity: function useBaseUri() view returns(bool)
func (_VolcanoERC721Tradable *VolcanoERC721TradableCallerSession) UseBaseUri() (bool, error) {
	return _VolcanoERC721Tradable.Contract.UseBaseUri(&_VolcanoERC721Tradable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.Approve(&_VolcanoERC721Tradable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.Approve(&_VolcanoERC721Tradable.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.Burn(&_VolcanoERC721Tradable.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.Burn(&_VolcanoERC721Tradable.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address to, string uri) payable returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactor) Mint(opts *bind.TransactOpts, to common.Address, uri string) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.contract.Transact(opts, "mint", to, uri)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address to, string uri) payable returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) Mint(to common.Address, uri string) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.Mint(&_VolcanoERC721Tradable.TransactOpts, to, uri)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address to, string uri) payable returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactorSession) Mint(to common.Address, uri string) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.Mint(&_VolcanoERC721Tradable.TransactOpts, to, uri)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.RenounceOwnership(&_VolcanoERC721Tradable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.RenounceOwnership(&_VolcanoERC721Tradable.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.SafeTransferFrom(&_VolcanoERC721Tradable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.SafeTransferFrom(&_VolcanoERC721Tradable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.SafeTransferFrom0(&_VolcanoERC721Tradable.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.SafeTransferFrom0(&_VolcanoERC721Tradable.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.SetApprovalForAll(&_VolcanoERC721Tradable.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.SetApprovalForAll(&_VolcanoERC721Tradable.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.TransferFrom(&_VolcanoERC721Tradable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.TransferFrom(&_VolcanoERC721Tradable.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.TransferOwnership(&_VolcanoERC721Tradable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.TransferOwnership(&_VolcanoERC721Tradable.TransactOpts, newOwner)
}

// UpdateContractURI is a paid mutator transaction binding the contract method 0x7e5b1e24.
//
// Solidity: function updateContractURI(string _uri) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactor) UpdateContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.contract.Transact(opts, "updateContractURI", _uri)
}

// UpdateContractURI is a paid mutator transaction binding the contract method 0x7e5b1e24.
//
// Solidity: function updateContractURI(string _uri) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) UpdateContractURI(_uri string) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.UpdateContractURI(&_VolcanoERC721Tradable.TransactOpts, _uri)
}

// UpdateContractURI is a paid mutator transaction binding the contract method 0x7e5b1e24.
//
// Solidity: function updateContractURI(string _uri) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactorSession) UpdateContractURI(_uri string) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.UpdateContractURI(&_VolcanoERC721Tradable.TransactOpts, _uri)
}

// UpdateCreatorFeePerc is a paid mutator transaction binding the contract method 0x19794ea6.
//
// Solidity: function updateCreatorFeePerc(uint96 _creatorFeePerc) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactor) UpdateCreatorFeePerc(opts *bind.TransactOpts, _creatorFeePerc *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.contract.Transact(opts, "updateCreatorFeePerc", _creatorFeePerc)
}

// UpdateCreatorFeePerc is a paid mutator transaction binding the contract method 0x19794ea6.
//
// Solidity: function updateCreatorFeePerc(uint96 _creatorFeePerc) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) UpdateCreatorFeePerc(_creatorFeePerc *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.UpdateCreatorFeePerc(&_VolcanoERC721Tradable.TransactOpts, _creatorFeePerc)
}

// UpdateCreatorFeePerc is a paid mutator transaction binding the contract method 0x19794ea6.
//
// Solidity: function updateCreatorFeePerc(uint96 _creatorFeePerc) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactorSession) UpdateCreatorFeePerc(_creatorFeePerc *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.UpdateCreatorFeePerc(&_VolcanoERC721Tradable.TransactOpts, _creatorFeePerc)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactor) UpdateFeeRecipient(opts *bind.TransactOpts, _feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.contract.Transact(opts, "updateFeeRecipient", _feeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableSession) UpdateFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.UpdateFeeRecipient(&_VolcanoERC721Tradable.TransactOpts, _feeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_VolcanoERC721Tradable *VolcanoERC721TradableTransactorSession) UpdateFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC721Tradable.Contract.UpdateFeeRecipient(&_VolcanoERC721Tradable.TransactOpts, _feeRecipient)
}

// VolcanoERC721TradableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableApprovalIterator struct {
	Event *VolcanoERC721TradableApproval // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721TradableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721TradableApproval)
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
		it.Event = new(VolcanoERC721TradableApproval)
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
func (it *VolcanoERC721TradableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721TradableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721TradableApproval represents a Approval event raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*VolcanoERC721TradableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VolcanoERC721Tradable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721TradableApprovalIterator{contract: _VolcanoERC721Tradable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VolcanoERC721TradableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VolcanoERC721Tradable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721TradableApproval)
				if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) ParseApproval(log types.Log) (*VolcanoERC721TradableApproval, error) {
	event := new(VolcanoERC721TradableApproval)
	if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721TradableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableApprovalForAllIterator struct {
	Event *VolcanoERC721TradableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721TradableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721TradableApprovalForAll)
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
		it.Event = new(VolcanoERC721TradableApprovalForAll)
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
func (it *VolcanoERC721TradableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721TradableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721TradableApprovalForAll represents a ApprovalForAll event raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*VolcanoERC721TradableApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _VolcanoERC721Tradable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721TradableApprovalForAllIterator{contract: _VolcanoERC721Tradable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *VolcanoERC721TradableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _VolcanoERC721Tradable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721TradableApprovalForAll)
				if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) ParseApprovalForAll(log types.Log) (*VolcanoERC721TradableApprovalForAll, error) {
	event := new(VolcanoERC721TradableApprovalForAll)
	if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721TradableBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableBatchMetadataUpdateIterator struct {
	Event *VolcanoERC721TradableBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721TradableBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721TradableBatchMetadataUpdate)
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
		it.Event = new(VolcanoERC721TradableBatchMetadataUpdate)
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
func (it *VolcanoERC721TradableBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721TradableBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721TradableBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*VolcanoERC721TradableBatchMetadataUpdateIterator, error) {

	logs, sub, err := _VolcanoERC721Tradable.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721TradableBatchMetadataUpdateIterator{contract: _VolcanoERC721Tradable.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *VolcanoERC721TradableBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC721Tradable.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721TradableBatchMetadataUpdate)
				if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) ParseBatchMetadataUpdate(log types.Log) (*VolcanoERC721TradableBatchMetadataUpdate, error) {
	event := new(VolcanoERC721TradableBatchMetadataUpdate)
	if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721TradableMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableMetadataUpdateIterator struct {
	Event *VolcanoERC721TradableMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721TradableMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721TradableMetadataUpdate)
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
		it.Event = new(VolcanoERC721TradableMetadataUpdate)
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
func (it *VolcanoERC721TradableMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721TradableMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721TradableMetadataUpdate represents a MetadataUpdate event raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*VolcanoERC721TradableMetadataUpdateIterator, error) {

	logs, sub, err := _VolcanoERC721Tradable.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721TradableMetadataUpdateIterator{contract: _VolcanoERC721Tradable.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *VolcanoERC721TradableMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC721Tradable.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721TradableMetadataUpdate)
				if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) ParseMetadataUpdate(log types.Log) (*VolcanoERC721TradableMetadataUpdate, error) {
	event := new(VolcanoERC721TradableMetadataUpdate)
	if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721TradableMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableMintedIterator struct {
	Event *VolcanoERC721TradableMinted // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721TradableMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721TradableMinted)
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
		it.Event = new(VolcanoERC721TradableMinted)
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
func (it *VolcanoERC721TradableMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721TradableMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721TradableMinted represents a Minted event raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableMinted struct {
	TokenId     *big.Int
	Beneficiary common.Address
	TokenUri    string
	Minter      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x997115af5924f5e38964c6d65c804d4cb85129b65e62eb20a8ca6329dbe57e18.
//
// Solidity: event Minted(uint256 tokenId, address beneficiary, string tokenUri, address minter)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) FilterMinted(opts *bind.FilterOpts) (*VolcanoERC721TradableMintedIterator, error) {

	logs, sub, err := _VolcanoERC721Tradable.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721TradableMintedIterator{contract: _VolcanoERC721Tradable.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x997115af5924f5e38964c6d65c804d4cb85129b65e62eb20a8ca6329dbe57e18.
//
// Solidity: event Minted(uint256 tokenId, address beneficiary, string tokenUri, address minter)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *VolcanoERC721TradableMinted) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC721Tradable.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721TradableMinted)
				if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x997115af5924f5e38964c6d65c804d4cb85129b65e62eb20a8ca6329dbe57e18.
//
// Solidity: event Minted(uint256 tokenId, address beneficiary, string tokenUri, address minter)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) ParseMinted(log types.Log) (*VolcanoERC721TradableMinted, error) {
	event := new(VolcanoERC721TradableMinted)
	if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721TradableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableOwnershipTransferredIterator struct {
	Event *VolcanoERC721TradableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721TradableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721TradableOwnershipTransferred)
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
		it.Event = new(VolcanoERC721TradableOwnershipTransferred)
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
func (it *VolcanoERC721TradableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721TradableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721TradableOwnershipTransferred represents a OwnershipTransferred event raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VolcanoERC721TradableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC721Tradable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721TradableOwnershipTransferredIterator{contract: _VolcanoERC721Tradable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VolcanoERC721TradableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC721Tradable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721TradableOwnershipTransferred)
				if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) ParseOwnershipTransferred(log types.Log) (*VolcanoERC721TradableOwnershipTransferred, error) {
	event := new(VolcanoERC721TradableOwnershipTransferred)
	if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721TradableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableTransferIterator struct {
	Event *VolcanoERC721TradableTransfer // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721TradableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721TradableTransfer)
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
		it.Event = new(VolcanoERC721TradableTransfer)
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
func (it *VolcanoERC721TradableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721TradableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721TradableTransfer represents a Transfer event raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*VolcanoERC721TradableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VolcanoERC721Tradable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721TradableTransferIterator{contract: _VolcanoERC721Tradable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VolcanoERC721TradableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VolcanoERC721Tradable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721TradableTransfer)
				if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) ParseTransfer(log types.Log) (*VolcanoERC721TradableTransfer, error) {
	event := new(VolcanoERC721TradableTransfer)
	if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721TradableUpdateCreatorFeeIterator is returned from FilterUpdateCreatorFee and is used to iterate over the raw logs and unpacked data for UpdateCreatorFee events raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableUpdateCreatorFeeIterator struct {
	Event *VolcanoERC721TradableUpdateCreatorFee // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721TradableUpdateCreatorFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721TradableUpdateCreatorFee)
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
		it.Event = new(VolcanoERC721TradableUpdateCreatorFee)
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
func (it *VolcanoERC721TradableUpdateCreatorFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721TradableUpdateCreatorFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721TradableUpdateCreatorFee represents a UpdateCreatorFee event raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableUpdateCreatorFee struct {
	CreatorFeePerc *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateCreatorFee is a free log retrieval operation binding the contract event 0xc3d04e260fe05d096e73692c9096281fdaf2535704474c7a7bcc02ca2f1def15.
//
// Solidity: event UpdateCreatorFee(uint96 creatorFeePerc)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) FilterUpdateCreatorFee(opts *bind.FilterOpts) (*VolcanoERC721TradableUpdateCreatorFeeIterator, error) {

	logs, sub, err := _VolcanoERC721Tradable.contract.FilterLogs(opts, "UpdateCreatorFee")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721TradableUpdateCreatorFeeIterator{contract: _VolcanoERC721Tradable.contract, event: "UpdateCreatorFee", logs: logs, sub: sub}, nil
}

// WatchUpdateCreatorFee is a free log subscription operation binding the contract event 0xc3d04e260fe05d096e73692c9096281fdaf2535704474c7a7bcc02ca2f1def15.
//
// Solidity: event UpdateCreatorFee(uint96 creatorFeePerc)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) WatchUpdateCreatorFee(opts *bind.WatchOpts, sink chan<- *VolcanoERC721TradableUpdateCreatorFee) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC721Tradable.contract.WatchLogs(opts, "UpdateCreatorFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721TradableUpdateCreatorFee)
				if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "UpdateCreatorFee", log); err != nil {
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

// ParseUpdateCreatorFee is a log parse operation binding the contract event 0xc3d04e260fe05d096e73692c9096281fdaf2535704474c7a7bcc02ca2f1def15.
//
// Solidity: event UpdateCreatorFee(uint96 creatorFeePerc)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) ParseUpdateCreatorFee(log types.Log) (*VolcanoERC721TradableUpdateCreatorFee, error) {
	event := new(VolcanoERC721TradableUpdateCreatorFee)
	if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "UpdateCreatorFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721TradableUpdateFeeRecipientIterator is returned from FilterUpdateFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdateFeeRecipient events raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableUpdateFeeRecipientIterator struct {
	Event *VolcanoERC721TradableUpdateFeeRecipient // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721TradableUpdateFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721TradableUpdateFeeRecipient)
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
		it.Event = new(VolcanoERC721TradableUpdateFeeRecipient)
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
func (it *VolcanoERC721TradableUpdateFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721TradableUpdateFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721TradableUpdateFeeRecipient represents a UpdateFeeRecipient event raised by the VolcanoERC721Tradable contract.
type VolcanoERC721TradableUpdateFeeRecipient struct {
	FeeRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateFeeRecipient is a free log retrieval operation binding the contract event 0x6632de8ab33c46549f7bb29f647ea0d751157b25fe6a14b1bcc7527cdfbeb79c.
//
// Solidity: event UpdateFeeRecipient(address feeRecipient)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) FilterUpdateFeeRecipient(opts *bind.FilterOpts) (*VolcanoERC721TradableUpdateFeeRecipientIterator, error) {

	logs, sub, err := _VolcanoERC721Tradable.contract.FilterLogs(opts, "UpdateFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721TradableUpdateFeeRecipientIterator{contract: _VolcanoERC721Tradable.contract, event: "UpdateFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdateFeeRecipient is a free log subscription operation binding the contract event 0x6632de8ab33c46549f7bb29f647ea0d751157b25fe6a14b1bcc7527cdfbeb79c.
//
// Solidity: event UpdateFeeRecipient(address feeRecipient)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) WatchUpdateFeeRecipient(opts *bind.WatchOpts, sink chan<- *VolcanoERC721TradableUpdateFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC721Tradable.contract.WatchLogs(opts, "UpdateFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721TradableUpdateFeeRecipient)
				if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "UpdateFeeRecipient", log); err != nil {
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

// ParseUpdateFeeRecipient is a log parse operation binding the contract event 0x6632de8ab33c46549f7bb29f647ea0d751157b25fe6a14b1bcc7527cdfbeb79c.
//
// Solidity: event UpdateFeeRecipient(address feeRecipient)
func (_VolcanoERC721Tradable *VolcanoERC721TradableFilterer) ParseUpdateFeeRecipient(log types.Log) (*VolcanoERC721TradableUpdateFeeRecipient, error) {
	event := new(VolcanoERC721TradableUpdateFeeRecipient)
	if err := _VolcanoERC721Tradable.contract.UnpackLog(event, "UpdateFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
