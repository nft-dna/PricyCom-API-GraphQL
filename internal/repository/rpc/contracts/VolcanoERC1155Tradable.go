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

// VolcanoERC1155TradableMetaData contains all meta data concerning the VolcanoERC1155Tradable contract.
var VolcanoERC1155TradableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_auction\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marketplace\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bundleMarketplace\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintCreatorFee\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"_creatorFeePerc\",\"type\":\"uint96\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isprivate\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"usebaseUriOnly\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"useDecimalUri\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"baseUriExt\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxItems\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxItemSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintStopTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"preRevealUri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contractUri\",\"type\":\"string\"}],\"internalType\":\"structVolcanoERC1155Tradable.contractERC1155Options\",\"name\":\"_options\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"creatorFeePerc\",\"type\":\"uint96\"}],\"name\":\"UpdateCreatorFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"UpdateFeeRecipient\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creatorFeePerc\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isOperator\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isprivate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"itemsSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxItemSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintCreatorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintStopTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revealTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"toHexString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"updateContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_creatorFeePerc\",\"type\":\"uint96\"}],\"name\":\"updateCreatorFeePerc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"updateFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VolcanoERC1155TradableABI is the input ABI used to generate the binding from.
// Deprecated: Use VolcanoERC1155TradableMetaData.ABI instead.
var VolcanoERC1155TradableABI = VolcanoERC1155TradableMetaData.ABI

// VolcanoERC1155Tradable is an auto generated Go binding around an Ethereum contract.
type VolcanoERC1155Tradable struct {
	VolcanoERC1155TradableCaller     // Read-only binding to the contract
	VolcanoERC1155TradableTransactor // Write-only binding to the contract
	VolcanoERC1155TradableFilterer   // Log filterer for contract events
}

// VolcanoERC1155TradableCaller is an auto generated read-only Go binding around an Ethereum contract.
type VolcanoERC1155TradableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC1155TradableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VolcanoERC1155TradableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC1155TradableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolcanoERC1155TradableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC1155TradableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolcanoERC1155TradableSession struct {
	Contract     *VolcanoERC1155Tradable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VolcanoERC1155TradableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolcanoERC1155TradableCallerSession struct {
	Contract *VolcanoERC1155TradableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// VolcanoERC1155TradableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolcanoERC1155TradableTransactorSession struct {
	Contract     *VolcanoERC1155TradableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// VolcanoERC1155TradableRaw is an auto generated low-level Go binding around an Ethereum contract.
type VolcanoERC1155TradableRaw struct {
	Contract *VolcanoERC1155Tradable // Generic contract binding to access the raw methods on
}

// VolcanoERC1155TradableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolcanoERC1155TradableCallerRaw struct {
	Contract *VolcanoERC1155TradableCaller // Generic read-only contract binding to access the raw methods on
}

// VolcanoERC1155TradableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolcanoERC1155TradableTransactorRaw struct {
	Contract *VolcanoERC1155TradableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolcanoERC1155Tradable creates a new instance of VolcanoERC1155Tradable, bound to a specific deployed contract.
func NewVolcanoERC1155Tradable(address common.Address, backend bind.ContractBackend) (*VolcanoERC1155Tradable, error) {
	contract, err := bindVolcanoERC1155Tradable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155Tradable{VolcanoERC1155TradableCaller: VolcanoERC1155TradableCaller{contract: contract}, VolcanoERC1155TradableTransactor: VolcanoERC1155TradableTransactor{contract: contract}, VolcanoERC1155TradableFilterer: VolcanoERC1155TradableFilterer{contract: contract}}, nil
}

// NewVolcanoERC1155TradableCaller creates a new read-only instance of VolcanoERC1155Tradable, bound to a specific deployed contract.
func NewVolcanoERC1155TradableCaller(address common.Address, caller bind.ContractCaller) (*VolcanoERC1155TradableCaller, error) {
	contract, err := bindVolcanoERC1155Tradable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155TradableCaller{contract: contract}, nil
}

// NewVolcanoERC1155TradableTransactor creates a new write-only instance of VolcanoERC1155Tradable, bound to a specific deployed contract.
func NewVolcanoERC1155TradableTransactor(address common.Address, transactor bind.ContractTransactor) (*VolcanoERC1155TradableTransactor, error) {
	contract, err := bindVolcanoERC1155Tradable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155TradableTransactor{contract: contract}, nil
}

// NewVolcanoERC1155TradableFilterer creates a new log filterer instance of VolcanoERC1155Tradable, bound to a specific deployed contract.
func NewVolcanoERC1155TradableFilterer(address common.Address, filterer bind.ContractFilterer) (*VolcanoERC1155TradableFilterer, error) {
	contract, err := bindVolcanoERC1155Tradable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155TradableFilterer{contract: contract}, nil
}

// bindVolcanoERC1155Tradable binds a generic wrapper to an already deployed contract.
func bindVolcanoERC1155Tradable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VolcanoERC1155TradableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC1155Tradable.Contract.VolcanoERC1155TradableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.VolcanoERC1155TradableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.VolcanoERC1155TradableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC1155Tradable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.BalanceOf(&_VolcanoERC1155Tradable.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.BalanceOf(&_VolcanoERC1155Tradable.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.BalanceOfBatch(&_VolcanoERC1155Tradable.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.BalanceOfBatch(&_VolcanoERC1155Tradable.CallOpts, accounts, ids)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) ContractURI() (string, error) {
	return _VolcanoERC1155Tradable.Contract.ContractURI(&_VolcanoERC1155Tradable.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) ContractURI() (string, error) {
	return _VolcanoERC1155Tradable.Contract.ContractURI(&_VolcanoERC1155Tradable.CallOpts)
}

// CreatorFeePerc is a free data retrieval call binding the contract method 0xa7050d63.
//
// Solidity: function creatorFeePerc() view returns(uint96)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) CreatorFeePerc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "creatorFeePerc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatorFeePerc is a free data retrieval call binding the contract method 0xa7050d63.
//
// Solidity: function creatorFeePerc() view returns(uint96)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) CreatorFeePerc() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.CreatorFeePerc(&_VolcanoERC1155Tradable.CallOpts)
}

// CreatorFeePerc is a free data retrieval call binding the contract method 0xa7050d63.
//
// Solidity: function creatorFeePerc() view returns(uint96)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) CreatorFeePerc() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.CreatorFeePerc(&_VolcanoERC1155Tradable.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) Exists(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "exists", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) Exists(id *big.Int) (bool, error) {
	return _VolcanoERC1155Tradable.Contract.Exists(&_VolcanoERC1155Tradable.CallOpts, id)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) Exists(id *big.Int) (bool, error) {
	return _VolcanoERC1155Tradable.Contract.Exists(&_VolcanoERC1155Tradable.CallOpts, id)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) Factory() (common.Address, error) {
	return _VolcanoERC1155Tradable.Contract.Factory(&_VolcanoERC1155Tradable.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) Factory() (common.Address, error) {
	return _VolcanoERC1155Tradable.Contract.Factory(&_VolcanoERC1155Tradable.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) FeeRecipient() (common.Address, error) {
	return _VolcanoERC1155Tradable.Contract.FeeRecipient(&_VolcanoERC1155Tradable.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) FeeRecipient() (common.Address, error) {
	return _VolcanoERC1155Tradable.Contract.FeeRecipient(&_VolcanoERC1155Tradable.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool isOperator)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool isOperator)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _VolcanoERC1155Tradable.Contract.IsApprovedForAll(&_VolcanoERC1155Tradable.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool isOperator)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _VolcanoERC1155Tradable.Contract.IsApprovedForAll(&_VolcanoERC1155Tradable.CallOpts, _owner, _operator)
}

// Isprivate is a free data retrieval call binding the contract method 0xfcbbbd1a.
//
// Solidity: function isprivate() view returns(bool)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) Isprivate(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "isprivate")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Isprivate is a free data retrieval call binding the contract method 0xfcbbbd1a.
//
// Solidity: function isprivate() view returns(bool)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) Isprivate() (bool, error) {
	return _VolcanoERC1155Tradable.Contract.Isprivate(&_VolcanoERC1155Tradable.CallOpts)
}

// Isprivate is a free data retrieval call binding the contract method 0xfcbbbd1a.
//
// Solidity: function isprivate() view returns(bool)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) Isprivate() (bool, error) {
	return _VolcanoERC1155Tradable.Contract.Isprivate(&_VolcanoERC1155Tradable.CallOpts)
}

// ItemsSupply is a free data retrieval call binding the contract method 0x49ffa7d3.
//
// Solidity: function itemsSupply() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) ItemsSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "itemsSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ItemsSupply is a free data retrieval call binding the contract method 0x49ffa7d3.
//
// Solidity: function itemsSupply() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) ItemsSupply() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.ItemsSupply(&_VolcanoERC1155Tradable.CallOpts)
}

// ItemsSupply is a free data retrieval call binding the contract method 0x49ffa7d3.
//
// Solidity: function itemsSupply() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) ItemsSupply() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.ItemsSupply(&_VolcanoERC1155Tradable.CallOpts)
}

// MaxItemSupply is a free data retrieval call binding the contract method 0xdd8d4434.
//
// Solidity: function maxItemSupply() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) MaxItemSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "maxItemSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxItemSupply is a free data retrieval call binding the contract method 0xdd8d4434.
//
// Solidity: function maxItemSupply() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) MaxItemSupply() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.MaxItemSupply(&_VolcanoERC1155Tradable.CallOpts)
}

// MaxItemSupply is a free data retrieval call binding the contract method 0xdd8d4434.
//
// Solidity: function maxItemSupply() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) MaxItemSupply() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.MaxItemSupply(&_VolcanoERC1155Tradable.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) MaxSupply() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.MaxSupply(&_VolcanoERC1155Tradable.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) MaxSupply() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.MaxSupply(&_VolcanoERC1155Tradable.CallOpts)
}

// MintCreatorFee is a free data retrieval call binding the contract method 0x9237d5a5.
//
// Solidity: function mintCreatorFee() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) MintCreatorFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "mintCreatorFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintCreatorFee is a free data retrieval call binding the contract method 0x9237d5a5.
//
// Solidity: function mintCreatorFee() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) MintCreatorFee() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.MintCreatorFee(&_VolcanoERC1155Tradable.CallOpts)
}

// MintCreatorFee is a free data retrieval call binding the contract method 0x9237d5a5.
//
// Solidity: function mintCreatorFee() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) MintCreatorFee() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.MintCreatorFee(&_VolcanoERC1155Tradable.CallOpts)
}

// MintStartTime is a free data retrieval call binding the contract method 0x931e2e49.
//
// Solidity: function mintStartTime() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) MintStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "mintStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintStartTime is a free data retrieval call binding the contract method 0x931e2e49.
//
// Solidity: function mintStartTime() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) MintStartTime() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.MintStartTime(&_VolcanoERC1155Tradable.CallOpts)
}

// MintStartTime is a free data retrieval call binding the contract method 0x931e2e49.
//
// Solidity: function mintStartTime() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) MintStartTime() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.MintStartTime(&_VolcanoERC1155Tradable.CallOpts)
}

// MintStopTime is a free data retrieval call binding the contract method 0x95c8b27a.
//
// Solidity: function mintStopTime() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) MintStopTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "mintStopTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintStopTime is a free data retrieval call binding the contract method 0x95c8b27a.
//
// Solidity: function mintStopTime() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) MintStopTime() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.MintStopTime(&_VolcanoERC1155Tradable.CallOpts)
}

// MintStopTime is a free data retrieval call binding the contract method 0x95c8b27a.
//
// Solidity: function mintStopTime() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) MintStopTime() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.MintStopTime(&_VolcanoERC1155Tradable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) Name() (string, error) {
	return _VolcanoERC1155Tradable.Contract.Name(&_VolcanoERC1155Tradable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) Name() (string, error) {
	return _VolcanoERC1155Tradable.Contract.Name(&_VolcanoERC1155Tradable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) Owner() (common.Address, error) {
	return _VolcanoERC1155Tradable.Contract.Owner(&_VolcanoERC1155Tradable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) Owner() (common.Address, error) {
	return _VolcanoERC1155Tradable.Contract.Owner(&_VolcanoERC1155Tradable.CallOpts)
}

// RevealTime is a free data retrieval call binding the contract method 0xba829d71.
//
// Solidity: function revealTime() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) RevealTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "revealTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevealTime is a free data retrieval call binding the contract method 0xba829d71.
//
// Solidity: function revealTime() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) RevealTime() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.RevealTime(&_VolcanoERC1155Tradable.CallOpts)
}

// RevealTime is a free data retrieval call binding the contract method 0xba829d71.
//
// Solidity: function revealTime() view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) RevealTime() (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.RevealTime(&_VolcanoERC1155Tradable.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

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
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.RoyaltyInfo(&_VolcanoERC1155Tradable.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.RoyaltyInfo(&_VolcanoERC1155Tradable.CallOpts, tokenId, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VolcanoERC1155Tradable.Contract.SupportsInterface(&_VolcanoERC1155Tradable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VolcanoERC1155Tradable.Contract.SupportsInterface(&_VolcanoERC1155Tradable.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) Symbol() (string, error) {
	return _VolcanoERC1155Tradable.Contract.Symbol(&_VolcanoERC1155Tradable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) Symbol() (string, error) {
	return _VolcanoERC1155Tradable.Contract.Symbol(&_VolcanoERC1155Tradable.CallOpts)
}

// ToHexString is a free data retrieval call binding the contract method 0x63e1cbea.
//
// Solidity: function toHexString(uint256 value, uint256 length) pure returns(string)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) ToHexString(opts *bind.CallOpts, value *big.Int, length *big.Int) (string, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "toHexString", value, length)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToHexString is a free data retrieval call binding the contract method 0x63e1cbea.
//
// Solidity: function toHexString(uint256 value, uint256 length) pure returns(string)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) ToHexString(value *big.Int, length *big.Int) (string, error) {
	return _VolcanoERC1155Tradable.Contract.ToHexString(&_VolcanoERC1155Tradable.CallOpts, value, length)
}

// ToHexString is a free data retrieval call binding the contract method 0x63e1cbea.
//
// Solidity: function toHexString(uint256 value, uint256 length) pure returns(string)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) ToHexString(value *big.Int, length *big.Int) (string, error) {
	return _VolcanoERC1155Tradable.Contract.ToHexString(&_VolcanoERC1155Tradable.CallOpts, value, length)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) TotalSupply(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "totalSupply", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.TotalSupply(&_VolcanoERC1155Tradable.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _VolcanoERC1155Tradable.Contract.TotalSupply(&_VolcanoERC1155Tradable.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _VolcanoERC1155Tradable.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) Uri(tokenId *big.Int) (string, error) {
	return _VolcanoERC1155Tradable.Contract.Uri(&_VolcanoERC1155Tradable.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _VolcanoERC1155Tradable.Contract.Uri(&_VolcanoERC1155Tradable.CallOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactor) Burn(opts *bind.TransactOpts, account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.contract.Transact(opts, "burn", account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.Burn(&_VolcanoERC1155Tradable.TransactOpts, account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactorSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.Burn(&_VolcanoERC1155Tradable.TransactOpts, account, id, value)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactor) BurnBatch(opts *bind.TransactOpts, account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.contract.Transact(opts, "burnBatch", account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.BurnBatch(&_VolcanoERC1155Tradable.TransactOpts, account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactorSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.BurnBatch(&_VolcanoERC1155Tradable.TransactOpts, account, ids, values)
}

// Mint is a paid mutator transaction binding the contract method 0x94d008ef.
//
// Solidity: function mint(address account, uint256 amount, bytes data) payable returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.contract.Transact(opts, "mint", account, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x94d008ef.
//
// Solidity: function mint(address account, uint256 amount, bytes data) payable returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) Mint(account common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.Mint(&_VolcanoERC1155Tradable.TransactOpts, account, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x94d008ef.
//
// Solidity: function mint(address account, uint256 amount, bytes data) payable returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactorSession) Mint(account common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.Mint(&_VolcanoERC1155Tradable.TransactOpts, account, amount, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.RenounceOwnership(&_VolcanoERC1155Tradable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.RenounceOwnership(&_VolcanoERC1155Tradable.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.SafeBatchTransferFrom(&_VolcanoERC1155Tradable.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.SafeBatchTransferFrom(&_VolcanoERC1155Tradable.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.SafeTransferFrom(&_VolcanoERC1155Tradable.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.SafeTransferFrom(&_VolcanoERC1155Tradable.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.SetApprovalForAll(&_VolcanoERC1155Tradable.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.SetApprovalForAll(&_VolcanoERC1155Tradable.TransactOpts, operator, approved)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.TransferOwnership(&_VolcanoERC1155Tradable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.TransferOwnership(&_VolcanoERC1155Tradable.TransactOpts, newOwner)
}

// UpdateContractURI is a paid mutator transaction binding the contract method 0x7e5b1e24.
//
// Solidity: function updateContractURI(string _uri) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactor) UpdateContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.contract.Transact(opts, "updateContractURI", _uri)
}

// UpdateContractURI is a paid mutator transaction binding the contract method 0x7e5b1e24.
//
// Solidity: function updateContractURI(string _uri) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) UpdateContractURI(_uri string) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.UpdateContractURI(&_VolcanoERC1155Tradable.TransactOpts, _uri)
}

// UpdateContractURI is a paid mutator transaction binding the contract method 0x7e5b1e24.
//
// Solidity: function updateContractURI(string _uri) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactorSession) UpdateContractURI(_uri string) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.UpdateContractURI(&_VolcanoERC1155Tradable.TransactOpts, _uri)
}

// UpdateCreatorFeePerc is a paid mutator transaction binding the contract method 0x19794ea6.
//
// Solidity: function updateCreatorFeePerc(uint96 _creatorFeePerc) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactor) UpdateCreatorFeePerc(opts *bind.TransactOpts, _creatorFeePerc *big.Int) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.contract.Transact(opts, "updateCreatorFeePerc", _creatorFeePerc)
}

// UpdateCreatorFeePerc is a paid mutator transaction binding the contract method 0x19794ea6.
//
// Solidity: function updateCreatorFeePerc(uint96 _creatorFeePerc) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) UpdateCreatorFeePerc(_creatorFeePerc *big.Int) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.UpdateCreatorFeePerc(&_VolcanoERC1155Tradable.TransactOpts, _creatorFeePerc)
}

// UpdateCreatorFeePerc is a paid mutator transaction binding the contract method 0x19794ea6.
//
// Solidity: function updateCreatorFeePerc(uint96 _creatorFeePerc) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactorSession) UpdateCreatorFeePerc(_creatorFeePerc *big.Int) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.UpdateCreatorFeePerc(&_VolcanoERC1155Tradable.TransactOpts, _creatorFeePerc)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactor) UpdateFeeRecipient(opts *bind.TransactOpts, _feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.contract.Transact(opts, "updateFeeRecipient", _feeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableSession) UpdateFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.UpdateFeeRecipient(&_VolcanoERC1155Tradable.TransactOpts, _feeRecipient)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address _feeRecipient) returns()
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableTransactorSession) UpdateFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC1155Tradable.Contract.UpdateFeeRecipient(&_VolcanoERC1155Tradable.TransactOpts, _feeRecipient)
}

// VolcanoERC1155TradableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableApprovalForAllIterator struct {
	Event *VolcanoERC1155TradableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *VolcanoERC1155TradableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC1155TradableApprovalForAll)
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
		it.Event = new(VolcanoERC1155TradableApprovalForAll)
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
func (it *VolcanoERC1155TradableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC1155TradableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC1155TradableApprovalForAll represents a ApprovalForAll event raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*VolcanoERC1155TradableApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _VolcanoERC1155Tradable.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155TradableApprovalForAllIterator{contract: _VolcanoERC1155Tradable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *VolcanoERC1155TradableApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _VolcanoERC1155Tradable.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC1155TradableApprovalForAll)
				if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) ParseApprovalForAll(log types.Log) (*VolcanoERC1155TradableApprovalForAll, error) {
	event := new(VolcanoERC1155TradableApprovalForAll)
	if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC1155TradableMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableMintedIterator struct {
	Event *VolcanoERC1155TradableMinted // Event containing the contract specifics and raw log

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
func (it *VolcanoERC1155TradableMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC1155TradableMinted)
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
		it.Event = new(VolcanoERC1155TradableMinted)
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
func (it *VolcanoERC1155TradableMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC1155TradableMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC1155TradableMinted represents a Minted event raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableMinted struct {
	TokenId     *big.Int
	Amount      *big.Int
	Beneficiary common.Address
	Uri         string
	Data        []byte
	Minter      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0xd1390903c6e0d21c236d4f87b8f6b8ed0c107dd8ab79fc5aeb0b654bb99039bd.
//
// Solidity: event Minted(uint256 tokenId, uint256 amount, address beneficiary, string uri, bytes data, address minter)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) FilterMinted(opts *bind.FilterOpts) (*VolcanoERC1155TradableMintedIterator, error) {

	logs, sub, err := _VolcanoERC1155Tradable.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155TradableMintedIterator{contract: _VolcanoERC1155Tradable.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0xd1390903c6e0d21c236d4f87b8f6b8ed0c107dd8ab79fc5aeb0b654bb99039bd.
//
// Solidity: event Minted(uint256 tokenId, uint256 amount, address beneficiary, string uri, bytes data, address minter)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *VolcanoERC1155TradableMinted) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC1155Tradable.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC1155TradableMinted)
				if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0xd1390903c6e0d21c236d4f87b8f6b8ed0c107dd8ab79fc5aeb0b654bb99039bd.
//
// Solidity: event Minted(uint256 tokenId, uint256 amount, address beneficiary, string uri, bytes data, address minter)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) ParseMinted(log types.Log) (*VolcanoERC1155TradableMinted, error) {
	event := new(VolcanoERC1155TradableMinted)
	if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC1155TradableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableOwnershipTransferredIterator struct {
	Event *VolcanoERC1155TradableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VolcanoERC1155TradableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC1155TradableOwnershipTransferred)
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
		it.Event = new(VolcanoERC1155TradableOwnershipTransferred)
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
func (it *VolcanoERC1155TradableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC1155TradableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC1155TradableOwnershipTransferred represents a OwnershipTransferred event raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VolcanoERC1155TradableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC1155Tradable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155TradableOwnershipTransferredIterator{contract: _VolcanoERC1155Tradable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VolcanoERC1155TradableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC1155Tradable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC1155TradableOwnershipTransferred)
				if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) ParseOwnershipTransferred(log types.Log) (*VolcanoERC1155TradableOwnershipTransferred, error) {
	event := new(VolcanoERC1155TradableOwnershipTransferred)
	if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC1155TradableTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableTransferBatchIterator struct {
	Event *VolcanoERC1155TradableTransferBatch // Event containing the contract specifics and raw log

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
func (it *VolcanoERC1155TradableTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC1155TradableTransferBatch)
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
		it.Event = new(VolcanoERC1155TradableTransferBatch)
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
func (it *VolcanoERC1155TradableTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC1155TradableTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC1155TradableTransferBatch represents a TransferBatch event raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*VolcanoERC1155TradableTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VolcanoERC1155Tradable.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155TradableTransferBatchIterator{contract: _VolcanoERC1155Tradable.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *VolcanoERC1155TradableTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VolcanoERC1155Tradable.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC1155TradableTransferBatch)
				if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) ParseTransferBatch(log types.Log) (*VolcanoERC1155TradableTransferBatch, error) {
	event := new(VolcanoERC1155TradableTransferBatch)
	if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC1155TradableTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableTransferSingleIterator struct {
	Event *VolcanoERC1155TradableTransferSingle // Event containing the contract specifics and raw log

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
func (it *VolcanoERC1155TradableTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC1155TradableTransferSingle)
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
		it.Event = new(VolcanoERC1155TradableTransferSingle)
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
func (it *VolcanoERC1155TradableTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC1155TradableTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC1155TradableTransferSingle represents a TransferSingle event raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*VolcanoERC1155TradableTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VolcanoERC1155Tradable.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155TradableTransferSingleIterator{contract: _VolcanoERC1155Tradable.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *VolcanoERC1155TradableTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VolcanoERC1155Tradable.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC1155TradableTransferSingle)
				if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) ParseTransferSingle(log types.Log) (*VolcanoERC1155TradableTransferSingle, error) {
	event := new(VolcanoERC1155TradableTransferSingle)
	if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC1155TradableURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableURIIterator struct {
	Event *VolcanoERC1155TradableURI // Event containing the contract specifics and raw log

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
func (it *VolcanoERC1155TradableURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC1155TradableURI)
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
		it.Event = new(VolcanoERC1155TradableURI)
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
func (it *VolcanoERC1155TradableURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC1155TradableURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC1155TradableURI represents a URI event raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*VolcanoERC1155TradableURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _VolcanoERC1155Tradable.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155TradableURIIterator{contract: _VolcanoERC1155Tradable.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *VolcanoERC1155TradableURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _VolcanoERC1155Tradable.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC1155TradableURI)
				if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) ParseURI(log types.Log) (*VolcanoERC1155TradableURI, error) {
	event := new(VolcanoERC1155TradableURI)
	if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC1155TradableUpdateCreatorFeeIterator is returned from FilterUpdateCreatorFee and is used to iterate over the raw logs and unpacked data for UpdateCreatorFee events raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableUpdateCreatorFeeIterator struct {
	Event *VolcanoERC1155TradableUpdateCreatorFee // Event containing the contract specifics and raw log

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
func (it *VolcanoERC1155TradableUpdateCreatorFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC1155TradableUpdateCreatorFee)
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
		it.Event = new(VolcanoERC1155TradableUpdateCreatorFee)
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
func (it *VolcanoERC1155TradableUpdateCreatorFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC1155TradableUpdateCreatorFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC1155TradableUpdateCreatorFee represents a UpdateCreatorFee event raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableUpdateCreatorFee struct {
	CreatorFeePerc *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateCreatorFee is a free log retrieval operation binding the contract event 0xc3d04e260fe05d096e73692c9096281fdaf2535704474c7a7bcc02ca2f1def15.
//
// Solidity: event UpdateCreatorFee(uint96 creatorFeePerc)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) FilterUpdateCreatorFee(opts *bind.FilterOpts) (*VolcanoERC1155TradableUpdateCreatorFeeIterator, error) {

	logs, sub, err := _VolcanoERC1155Tradable.contract.FilterLogs(opts, "UpdateCreatorFee")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155TradableUpdateCreatorFeeIterator{contract: _VolcanoERC1155Tradable.contract, event: "UpdateCreatorFee", logs: logs, sub: sub}, nil
}

// WatchUpdateCreatorFee is a free log subscription operation binding the contract event 0xc3d04e260fe05d096e73692c9096281fdaf2535704474c7a7bcc02ca2f1def15.
//
// Solidity: event UpdateCreatorFee(uint96 creatorFeePerc)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) WatchUpdateCreatorFee(opts *bind.WatchOpts, sink chan<- *VolcanoERC1155TradableUpdateCreatorFee) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC1155Tradable.contract.WatchLogs(opts, "UpdateCreatorFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC1155TradableUpdateCreatorFee)
				if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "UpdateCreatorFee", log); err != nil {
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
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) ParseUpdateCreatorFee(log types.Log) (*VolcanoERC1155TradableUpdateCreatorFee, error) {
	event := new(VolcanoERC1155TradableUpdateCreatorFee)
	if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "UpdateCreatorFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC1155TradableUpdateFeeRecipientIterator is returned from FilterUpdateFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdateFeeRecipient events raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableUpdateFeeRecipientIterator struct {
	Event *VolcanoERC1155TradableUpdateFeeRecipient // Event containing the contract specifics and raw log

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
func (it *VolcanoERC1155TradableUpdateFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC1155TradableUpdateFeeRecipient)
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
		it.Event = new(VolcanoERC1155TradableUpdateFeeRecipient)
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
func (it *VolcanoERC1155TradableUpdateFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC1155TradableUpdateFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC1155TradableUpdateFeeRecipient represents a UpdateFeeRecipient event raised by the VolcanoERC1155Tradable contract.
type VolcanoERC1155TradableUpdateFeeRecipient struct {
	FeeRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateFeeRecipient is a free log retrieval operation binding the contract event 0x6632de8ab33c46549f7bb29f647ea0d751157b25fe6a14b1bcc7527cdfbeb79c.
//
// Solidity: event UpdateFeeRecipient(address feeRecipient)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) FilterUpdateFeeRecipient(opts *bind.FilterOpts) (*VolcanoERC1155TradableUpdateFeeRecipientIterator, error) {

	logs, sub, err := _VolcanoERC1155Tradable.contract.FilterLogs(opts, "UpdateFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC1155TradableUpdateFeeRecipientIterator{contract: _VolcanoERC1155Tradable.contract, event: "UpdateFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdateFeeRecipient is a free log subscription operation binding the contract event 0x6632de8ab33c46549f7bb29f647ea0d751157b25fe6a14b1bcc7527cdfbeb79c.
//
// Solidity: event UpdateFeeRecipient(address feeRecipient)
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) WatchUpdateFeeRecipient(opts *bind.WatchOpts, sink chan<- *VolcanoERC1155TradableUpdateFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC1155Tradable.contract.WatchLogs(opts, "UpdateFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC1155TradableUpdateFeeRecipient)
				if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "UpdateFeeRecipient", log); err != nil {
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
func (_VolcanoERC1155Tradable *VolcanoERC1155TradableFilterer) ParseUpdateFeeRecipient(log types.Log) (*VolcanoERC1155TradableUpdateFeeRecipient, error) {
	event := new(VolcanoERC1155TradableUpdateFeeRecipient)
	if err := _VolcanoERC1155Tradable.contract.UnpackLog(event, "UpdateFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
