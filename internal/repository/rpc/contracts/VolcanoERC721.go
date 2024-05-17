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

// VolcanoERC721MetaData contains all meta data concerning the VolcanoERC721 contract.
var VolcanoERC721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"name\":\"UpdatePlatformFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"UpdatePlatformFeeRecipient\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_tokenUri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"updatePlatformFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VolcanoERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use VolcanoERC721MetaData.ABI instead.
var VolcanoERC721ABI = VolcanoERC721MetaData.ABI

// VolcanoERC721 is an auto generated Go binding around an Ethereum contract.
type VolcanoERC721 struct {
	VolcanoERC721Caller     // Read-only binding to the contract
	VolcanoERC721Transactor // Write-only binding to the contract
	VolcanoERC721Filterer   // Log filterer for contract events
}

// VolcanoERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type VolcanoERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type VolcanoERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolcanoERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolcanoERC721Session struct {
	Contract     *VolcanoERC721    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VolcanoERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolcanoERC721CallerSession struct {
	Contract *VolcanoERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VolcanoERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolcanoERC721TransactorSession struct {
	Contract     *VolcanoERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VolcanoERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type VolcanoERC721Raw struct {
	Contract *VolcanoERC721 // Generic contract binding to access the raw methods on
}

// VolcanoERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolcanoERC721CallerRaw struct {
	Contract *VolcanoERC721Caller // Generic read-only contract binding to access the raw methods on
}

// VolcanoERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolcanoERC721TransactorRaw struct {
	Contract *VolcanoERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVolcanoERC721 creates a new instance of VolcanoERC721, bound to a specific deployed contract.
func NewVolcanoERC721(address common.Address, backend bind.ContractBackend) (*VolcanoERC721, error) {
	contract, err := bindVolcanoERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721{VolcanoERC721Caller: VolcanoERC721Caller{contract: contract}, VolcanoERC721Transactor: VolcanoERC721Transactor{contract: contract}, VolcanoERC721Filterer: VolcanoERC721Filterer{contract: contract}}, nil
}

// NewVolcanoERC721Caller creates a new read-only instance of VolcanoERC721, bound to a specific deployed contract.
func NewVolcanoERC721Caller(address common.Address, caller bind.ContractCaller) (*VolcanoERC721Caller, error) {
	contract, err := bindVolcanoERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721Caller{contract: contract}, nil
}

// NewVolcanoERC721Transactor creates a new write-only instance of VolcanoERC721, bound to a specific deployed contract.
func NewVolcanoERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*VolcanoERC721Transactor, error) {
	contract, err := bindVolcanoERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721Transactor{contract: contract}, nil
}

// NewVolcanoERC721Filterer creates a new log filterer instance of VolcanoERC721, bound to a specific deployed contract.
func NewVolcanoERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*VolcanoERC721Filterer, error) {
	contract, err := bindVolcanoERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721Filterer{contract: contract}, nil
}

// bindVolcanoERC721 binds a generic wrapper to an already deployed contract.
func bindVolcanoERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VolcanoERC721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC721 *VolcanoERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC721.Contract.VolcanoERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC721 *VolcanoERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.VolcanoERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC721 *VolcanoERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.VolcanoERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC721 *VolcanoERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC721 *VolcanoERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC721 *VolcanoERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_VolcanoERC721 *VolcanoERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_VolcanoERC721 *VolcanoERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VolcanoERC721.Contract.BalanceOf(&_VolcanoERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_VolcanoERC721 *VolcanoERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VolcanoERC721.Contract.BalanceOf(&_VolcanoERC721.CallOpts, owner)
}

// Creators is a free data retrieval call binding the contract method 0xcd53d08e.
//
// Solidity: function creators(uint256 ) view returns(address)
func (_VolcanoERC721 *VolcanoERC721Caller) Creators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "creators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creators is a free data retrieval call binding the contract method 0xcd53d08e.
//
// Solidity: function creators(uint256 ) view returns(address)
func (_VolcanoERC721 *VolcanoERC721Session) Creators(arg0 *big.Int) (common.Address, error) {
	return _VolcanoERC721.Contract.Creators(&_VolcanoERC721.CallOpts, arg0)
}

// Creators is a free data retrieval call binding the contract method 0xcd53d08e.
//
// Solidity: function creators(uint256 ) view returns(address)
func (_VolcanoERC721 *VolcanoERC721CallerSession) Creators(arg0 *big.Int) (common.Address, error) {
	return _VolcanoERC721.Contract.Creators(&_VolcanoERC721.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) view returns(bool)
func (_VolcanoERC721 *VolcanoERC721Caller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "exists", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) view returns(bool)
func (_VolcanoERC721 *VolcanoERC721Session) Exists(_tokenId *big.Int) (bool, error) {
	return _VolcanoERC721.Contract.Exists(&_VolcanoERC721.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 _tokenId) view returns(bool)
func (_VolcanoERC721 *VolcanoERC721CallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _VolcanoERC721.Contract.Exists(&_VolcanoERC721.CallOpts, _tokenId)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_VolcanoERC721 *VolcanoERC721Caller) FeeReceipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "feeReceipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_VolcanoERC721 *VolcanoERC721Session) FeeReceipient() (common.Address, error) {
	return _VolcanoERC721.Contract.FeeReceipient(&_VolcanoERC721.CallOpts)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_VolcanoERC721 *VolcanoERC721CallerSession) FeeReceipient() (common.Address, error) {
	return _VolcanoERC721.Contract.FeeReceipient(&_VolcanoERC721.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_VolcanoERC721 *VolcanoERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_VolcanoERC721 *VolcanoERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _VolcanoERC721.Contract.GetApproved(&_VolcanoERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_VolcanoERC721 *VolcanoERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _VolcanoERC721.Contract.GetApproved(&_VolcanoERC721.CallOpts, tokenId)
}

// IsApproved is a free data retrieval call binding the contract method 0x56c31637.
//
// Solidity: function isApproved(uint256 _tokenId, address _operator) view returns(bool)
func (_VolcanoERC721 *VolcanoERC721Caller) IsApproved(opts *bind.CallOpts, _tokenId *big.Int, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "isApproved", _tokenId, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x56c31637.
//
// Solidity: function isApproved(uint256 _tokenId, address _operator) view returns(bool)
func (_VolcanoERC721 *VolcanoERC721Session) IsApproved(_tokenId *big.Int, _operator common.Address) (bool, error) {
	return _VolcanoERC721.Contract.IsApproved(&_VolcanoERC721.CallOpts, _tokenId, _operator)
}

// IsApproved is a free data retrieval call binding the contract method 0x56c31637.
//
// Solidity: function isApproved(uint256 _tokenId, address _operator) view returns(bool)
func (_VolcanoERC721 *VolcanoERC721CallerSession) IsApproved(_tokenId *big.Int, _operator common.Address) (bool, error) {
	return _VolcanoERC721.Contract.IsApproved(&_VolcanoERC721.CallOpts, _tokenId, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_VolcanoERC721 *VolcanoERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_VolcanoERC721 *VolcanoERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _VolcanoERC721.Contract.IsApprovedForAll(&_VolcanoERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_VolcanoERC721 *VolcanoERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _VolcanoERC721.Contract.IsApprovedForAll(&_VolcanoERC721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VolcanoERC721 *VolcanoERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VolcanoERC721 *VolcanoERC721Session) Name() (string, error) {
	return _VolcanoERC721.Contract.Name(&_VolcanoERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VolcanoERC721 *VolcanoERC721CallerSession) Name() (string, error) {
	return _VolcanoERC721.Contract.Name(&_VolcanoERC721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC721 *VolcanoERC721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC721 *VolcanoERC721Session) Owner() (common.Address, error) {
	return _VolcanoERC721.Contract.Owner(&_VolcanoERC721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC721 *VolcanoERC721CallerSession) Owner() (common.Address, error) {
	return _VolcanoERC721.Contract.Owner(&_VolcanoERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_VolcanoERC721 *VolcanoERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_VolcanoERC721 *VolcanoERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _VolcanoERC721.Contract.OwnerOf(&_VolcanoERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_VolcanoERC721 *VolcanoERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _VolcanoERC721.Contract.OwnerOf(&_VolcanoERC721.CallOpts, tokenId)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VolcanoERC721 *VolcanoERC721Caller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VolcanoERC721 *VolcanoERC721Session) PlatformFee() (*big.Int, error) {
	return _VolcanoERC721.Contract.PlatformFee(&_VolcanoERC721.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VolcanoERC721 *VolcanoERC721CallerSession) PlatformFee() (*big.Int, error) {
	return _VolcanoERC721.Contract.PlatformFee(&_VolcanoERC721.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VolcanoERC721 *VolcanoERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VolcanoERC721 *VolcanoERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VolcanoERC721.Contract.SupportsInterface(&_VolcanoERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VolcanoERC721 *VolcanoERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VolcanoERC721.Contract.SupportsInterface(&_VolcanoERC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VolcanoERC721 *VolcanoERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VolcanoERC721 *VolcanoERC721Session) Symbol() (string, error) {
	return _VolcanoERC721.Contract.Symbol(&_VolcanoERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VolcanoERC721 *VolcanoERC721CallerSession) Symbol() (string, error) {
	return _VolcanoERC721.Contract.Symbol(&_VolcanoERC721.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_VolcanoERC721 *VolcanoERC721Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_VolcanoERC721 *VolcanoERC721Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _VolcanoERC721.Contract.TokenByIndex(&_VolcanoERC721.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_VolcanoERC721 *VolcanoERC721CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _VolcanoERC721.Contract.TokenByIndex(&_VolcanoERC721.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_VolcanoERC721 *VolcanoERC721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_VolcanoERC721 *VolcanoERC721Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _VolcanoERC721.Contract.TokenOfOwnerByIndex(&_VolcanoERC721.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_VolcanoERC721 *VolcanoERC721CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _VolcanoERC721.Contract.TokenOfOwnerByIndex(&_VolcanoERC721.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_VolcanoERC721 *VolcanoERC721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_VolcanoERC721 *VolcanoERC721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _VolcanoERC721.Contract.TokenURI(&_VolcanoERC721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_VolcanoERC721 *VolcanoERC721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _VolcanoERC721.Contract.TokenURI(&_VolcanoERC721.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VolcanoERC721 *VolcanoERC721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC721.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VolcanoERC721 *VolcanoERC721Session) TotalSupply() (*big.Int, error) {
	return _VolcanoERC721.Contract.TotalSupply(&_VolcanoERC721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VolcanoERC721 *VolcanoERC721CallerSession) TotalSupply() (*big.Int, error) {
	return _VolcanoERC721.Contract.TotalSupply(&_VolcanoERC721.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_VolcanoERC721 *VolcanoERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_VolcanoERC721 *VolcanoERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.Approve(&_VolcanoERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_VolcanoERC721 *VolcanoERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.Approve(&_VolcanoERC721.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_VolcanoERC721 *VolcanoERC721Transactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_VolcanoERC721 *VolcanoERC721Session) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.Burn(&_VolcanoERC721.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_VolcanoERC721 *VolcanoERC721TransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.Burn(&_VolcanoERC721.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address _beneficiary, string _tokenUri) payable returns(uint256)
func (_VolcanoERC721 *VolcanoERC721Transactor) Mint(opts *bind.TransactOpts, _beneficiary common.Address, _tokenUri string) (*types.Transaction, error) {
	return _VolcanoERC721.contract.Transact(opts, "mint", _beneficiary, _tokenUri)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address _beneficiary, string _tokenUri) payable returns(uint256)
func (_VolcanoERC721 *VolcanoERC721Session) Mint(_beneficiary common.Address, _tokenUri string) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.Mint(&_VolcanoERC721.TransactOpts, _beneficiary, _tokenUri)
}

// Mint is a paid mutator transaction binding the contract method 0xd0def521.
//
// Solidity: function mint(address _beneficiary, string _tokenUri) payable returns(uint256)
func (_VolcanoERC721 *VolcanoERC721TransactorSession) Mint(_beneficiary common.Address, _tokenUri string) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.Mint(&_VolcanoERC721.TransactOpts, _beneficiary, _tokenUri)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC721 *VolcanoERC721Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC721.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC721 *VolcanoERC721Session) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC721.Contract.RenounceOwnership(&_VolcanoERC721.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC721 *VolcanoERC721TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC721.Contract.RenounceOwnership(&_VolcanoERC721.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_VolcanoERC721 *VolcanoERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_VolcanoERC721 *VolcanoERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.SafeTransferFrom(&_VolcanoERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_VolcanoERC721 *VolcanoERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.SafeTransferFrom(&_VolcanoERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_VolcanoERC721 *VolcanoERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _VolcanoERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_VolcanoERC721 *VolcanoERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.SafeTransferFrom0(&_VolcanoERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_VolcanoERC721 *VolcanoERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.SafeTransferFrom0(&_VolcanoERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_VolcanoERC721 *VolcanoERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _VolcanoERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_VolcanoERC721 *VolcanoERC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.SetApprovalForAll(&_VolcanoERC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_VolcanoERC721 *VolcanoERC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.SetApprovalForAll(&_VolcanoERC721.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_VolcanoERC721 *VolcanoERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_VolcanoERC721 *VolcanoERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.TransferFrom(&_VolcanoERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_VolcanoERC721 *VolcanoERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.TransferFrom(&_VolcanoERC721.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC721 *VolcanoERC721Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC721.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC721 *VolcanoERC721Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.TransferOwnership(&_VolcanoERC721.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC721 *VolcanoERC721TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.TransferOwnership(&_VolcanoERC721.TransactOpts, newOwner)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_VolcanoERC721 *VolcanoERC721Transactor) UpdatePlatformFee(opts *bind.TransactOpts, _platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721.contract.Transact(opts, "updatePlatformFee", _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_VolcanoERC721 *VolcanoERC721Session) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.UpdatePlatformFee(&_VolcanoERC721.TransactOpts, _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_VolcanoERC721 *VolcanoERC721TransactorSession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.UpdatePlatformFee(&_VolcanoERC721.TransactOpts, _platformFee)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_VolcanoERC721 *VolcanoERC721Transactor) UpdatePlatformFeeRecipient(opts *bind.TransactOpts, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC721.contract.Transact(opts, "updatePlatformFeeRecipient", _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_VolcanoERC721 *VolcanoERC721Session) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.UpdatePlatformFeeRecipient(&_VolcanoERC721.TransactOpts, _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_VolcanoERC721 *VolcanoERC721TransactorSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC721.Contract.UpdatePlatformFeeRecipient(&_VolcanoERC721.TransactOpts, _platformFeeRecipient)
}

// VolcanoERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VolcanoERC721 contract.
type VolcanoERC721ApprovalIterator struct {
	Event *VolcanoERC721Approval // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721Approval)
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
		it.Event = new(VolcanoERC721Approval)
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
func (it *VolcanoERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721Approval represents a Approval event raised by the VolcanoERC721 contract.
type VolcanoERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_VolcanoERC721 *VolcanoERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*VolcanoERC721ApprovalIterator, error) {

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

	logs, sub, err := _VolcanoERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721ApprovalIterator{contract: _VolcanoERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_VolcanoERC721 *VolcanoERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VolcanoERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _VolcanoERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721Approval)
				if err := _VolcanoERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_VolcanoERC721 *VolcanoERC721Filterer) ParseApproval(log types.Log) (*VolcanoERC721Approval, error) {
	event := new(VolcanoERC721Approval)
	if err := _VolcanoERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the VolcanoERC721 contract.
type VolcanoERC721ApprovalForAllIterator struct {
	Event *VolcanoERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721ApprovalForAll)
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
		it.Event = new(VolcanoERC721ApprovalForAll)
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
func (it *VolcanoERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721ApprovalForAll represents a ApprovalForAll event raised by the VolcanoERC721 contract.
type VolcanoERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_VolcanoERC721 *VolcanoERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*VolcanoERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _VolcanoERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721ApprovalForAllIterator{contract: _VolcanoERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_VolcanoERC721 *VolcanoERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *VolcanoERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _VolcanoERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721ApprovalForAll)
				if err := _VolcanoERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_VolcanoERC721 *VolcanoERC721Filterer) ParseApprovalForAll(log types.Log) (*VolcanoERC721ApprovalForAll, error) {
	event := new(VolcanoERC721ApprovalForAll)
	if err := _VolcanoERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721BatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the VolcanoERC721 contract.
type VolcanoERC721BatchMetadataUpdateIterator struct {
	Event *VolcanoERC721BatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721BatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721BatchMetadataUpdate)
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
		it.Event = new(VolcanoERC721BatchMetadataUpdate)
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
func (it *VolcanoERC721BatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721BatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721BatchMetadataUpdate represents a BatchMetadataUpdate event raised by the VolcanoERC721 contract.
type VolcanoERC721BatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_VolcanoERC721 *VolcanoERC721Filterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*VolcanoERC721BatchMetadataUpdateIterator, error) {

	logs, sub, err := _VolcanoERC721.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721BatchMetadataUpdateIterator{contract: _VolcanoERC721.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_VolcanoERC721 *VolcanoERC721Filterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *VolcanoERC721BatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC721.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721BatchMetadataUpdate)
				if err := _VolcanoERC721.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_VolcanoERC721 *VolcanoERC721Filterer) ParseBatchMetadataUpdate(log types.Log) (*VolcanoERC721BatchMetadataUpdate, error) {
	event := new(VolcanoERC721BatchMetadataUpdate)
	if err := _VolcanoERC721.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721MetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the VolcanoERC721 contract.
type VolcanoERC721MetadataUpdateIterator struct {
	Event *VolcanoERC721MetadataUpdate // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721MetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721MetadataUpdate)
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
		it.Event = new(VolcanoERC721MetadataUpdate)
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
func (it *VolcanoERC721MetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721MetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721MetadataUpdate represents a MetadataUpdate event raised by the VolcanoERC721 contract.
type VolcanoERC721MetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_VolcanoERC721 *VolcanoERC721Filterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*VolcanoERC721MetadataUpdateIterator, error) {

	logs, sub, err := _VolcanoERC721.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721MetadataUpdateIterator{contract: _VolcanoERC721.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_VolcanoERC721 *VolcanoERC721Filterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *VolcanoERC721MetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC721.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721MetadataUpdate)
				if err := _VolcanoERC721.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_VolcanoERC721 *VolcanoERC721Filterer) ParseMetadataUpdate(log types.Log) (*VolcanoERC721MetadataUpdate, error) {
	event := new(VolcanoERC721MetadataUpdate)
	if err := _VolcanoERC721.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721MintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the VolcanoERC721 contract.
type VolcanoERC721MintedIterator struct {
	Event *VolcanoERC721Minted // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721MintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721Minted)
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
		it.Event = new(VolcanoERC721Minted)
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
func (it *VolcanoERC721MintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721MintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721Minted represents a Minted event raised by the VolcanoERC721 contract.
type VolcanoERC721Minted struct {
	TokenId     *big.Int
	Beneficiary common.Address
	TokenUri    string
	Minter      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x997115af5924f5e38964c6d65c804d4cb85129b65e62eb20a8ca6329dbe57e18.
//
// Solidity: event Minted(uint256 tokenId, address beneficiary, string tokenUri, address minter)
func (_VolcanoERC721 *VolcanoERC721Filterer) FilterMinted(opts *bind.FilterOpts) (*VolcanoERC721MintedIterator, error) {

	logs, sub, err := _VolcanoERC721.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721MintedIterator{contract: _VolcanoERC721.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x997115af5924f5e38964c6d65c804d4cb85129b65e62eb20a8ca6329dbe57e18.
//
// Solidity: event Minted(uint256 tokenId, address beneficiary, string tokenUri, address minter)
func (_VolcanoERC721 *VolcanoERC721Filterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *VolcanoERC721Minted) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC721.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721Minted)
				if err := _VolcanoERC721.contract.UnpackLog(event, "Minted", log); err != nil {
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
func (_VolcanoERC721 *VolcanoERC721Filterer) ParseMinted(log types.Log) (*VolcanoERC721Minted, error) {
	event := new(VolcanoERC721Minted)
	if err := _VolcanoERC721.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VolcanoERC721 contract.
type VolcanoERC721OwnershipTransferredIterator struct {
	Event *VolcanoERC721OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721OwnershipTransferred)
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
		it.Event = new(VolcanoERC721OwnershipTransferred)
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
func (it *VolcanoERC721OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721OwnershipTransferred represents a OwnershipTransferred event raised by the VolcanoERC721 contract.
type VolcanoERC721OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC721 *VolcanoERC721Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VolcanoERC721OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC721.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721OwnershipTransferredIterator{contract: _VolcanoERC721.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC721 *VolcanoERC721Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VolcanoERC721OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC721.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721OwnershipTransferred)
				if err := _VolcanoERC721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VolcanoERC721 *VolcanoERC721Filterer) ParseOwnershipTransferred(log types.Log) (*VolcanoERC721OwnershipTransferred, error) {
	event := new(VolcanoERC721OwnershipTransferred)
	if err := _VolcanoERC721.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VolcanoERC721 contract.
type VolcanoERC721TransferIterator struct {
	Event *VolcanoERC721Transfer // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721Transfer)
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
		it.Event = new(VolcanoERC721Transfer)
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
func (it *VolcanoERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721Transfer represents a Transfer event raised by the VolcanoERC721 contract.
type VolcanoERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_VolcanoERC721 *VolcanoERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*VolcanoERC721TransferIterator, error) {

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

	logs, sub, err := _VolcanoERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721TransferIterator{contract: _VolcanoERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_VolcanoERC721 *VolcanoERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VolcanoERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _VolcanoERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721Transfer)
				if err := _VolcanoERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_VolcanoERC721 *VolcanoERC721Filterer) ParseTransfer(log types.Log) (*VolcanoERC721Transfer, error) {
	event := new(VolcanoERC721Transfer)
	if err := _VolcanoERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721UpdatePlatformFeeIterator is returned from FilterUpdatePlatformFee and is used to iterate over the raw logs and unpacked data for UpdatePlatformFee events raised by the VolcanoERC721 contract.
type VolcanoERC721UpdatePlatformFeeIterator struct {
	Event *VolcanoERC721UpdatePlatformFee // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721UpdatePlatformFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721UpdatePlatformFee)
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
		it.Event = new(VolcanoERC721UpdatePlatformFee)
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
func (it *VolcanoERC721UpdatePlatformFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721UpdatePlatformFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721UpdatePlatformFee represents a UpdatePlatformFee event raised by the VolcanoERC721 contract.
type VolcanoERC721UpdatePlatformFee struct {
	PlatformFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFee is a free log retrieval operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_VolcanoERC721 *VolcanoERC721Filterer) FilterUpdatePlatformFee(opts *bind.FilterOpts) (*VolcanoERC721UpdatePlatformFeeIterator, error) {

	logs, sub, err := _VolcanoERC721.contract.FilterLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721UpdatePlatformFeeIterator{contract: _VolcanoERC721.contract, event: "UpdatePlatformFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFee is a free log subscription operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_VolcanoERC721 *VolcanoERC721Filterer) WatchUpdatePlatformFee(opts *bind.WatchOpts, sink chan<- *VolcanoERC721UpdatePlatformFee) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC721.contract.WatchLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721UpdatePlatformFee)
				if err := _VolcanoERC721.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
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

// ParseUpdatePlatformFee is a log parse operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_VolcanoERC721 *VolcanoERC721Filterer) ParseUpdatePlatformFee(log types.Log) (*VolcanoERC721UpdatePlatformFee, error) {
	event := new(VolcanoERC721UpdatePlatformFee)
	if err := _VolcanoERC721.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC721UpdatePlatformFeeRecipientIterator is returned from FilterUpdatePlatformFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdatePlatformFeeRecipient events raised by the VolcanoERC721 contract.
type VolcanoERC721UpdatePlatformFeeRecipientIterator struct {
	Event *VolcanoERC721UpdatePlatformFeeRecipient // Event containing the contract specifics and raw log

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
func (it *VolcanoERC721UpdatePlatformFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC721UpdatePlatformFeeRecipient)
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
		it.Event = new(VolcanoERC721UpdatePlatformFeeRecipient)
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
func (it *VolcanoERC721UpdatePlatformFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC721UpdatePlatformFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC721UpdatePlatformFeeRecipient represents a UpdatePlatformFeeRecipient event raised by the VolcanoERC721 contract.
type VolcanoERC721UpdatePlatformFeeRecipient struct {
	PlatformFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFeeRecipient is a free log retrieval operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_VolcanoERC721 *VolcanoERC721Filterer) FilterUpdatePlatformFeeRecipient(opts *bind.FilterOpts) (*VolcanoERC721UpdatePlatformFeeRecipientIterator, error) {

	logs, sub, err := _VolcanoERC721.contract.FilterLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC721UpdatePlatformFeeRecipientIterator{contract: _VolcanoERC721.contract, event: "UpdatePlatformFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFeeRecipient is a free log subscription operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_VolcanoERC721 *VolcanoERC721Filterer) WatchUpdatePlatformFeeRecipient(opts *bind.WatchOpts, sink chan<- *VolcanoERC721UpdatePlatformFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC721.contract.WatchLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC721UpdatePlatformFeeRecipient)
				if err := _VolcanoERC721.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
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

// ParseUpdatePlatformFeeRecipient is a log parse operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_VolcanoERC721 *VolcanoERC721Filterer) ParseUpdatePlatformFeeRecipient(log types.Log) (*VolcanoERC721UpdatePlatformFeeRecipient, error) {
	event := new(VolcanoERC721UpdatePlatformFeeRecipient)
	if err := _VolcanoERC721.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
