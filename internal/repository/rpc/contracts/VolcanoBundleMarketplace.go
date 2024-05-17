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

// VolcanoBundleMarketplaceMetaData contains all meta data concerning the VolcanoBundleMarketplace contract.
var VolcanoBundleMarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"bundleID\",\"type\":\"string\"}],\"name\":\"ItemCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"bundleID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startingTime\",\"type\":\"uint256\"}],\"name\":\"ItemListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"bundleID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"payTokenPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ItemSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"bundleID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nft\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenId\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"quantity\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"ItemUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"bundleID\",\"type\":\"string\"}],\"name\":\"OfferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"bundleID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"OfferCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"name\":\"UpdatePlatformFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"UpdatePlatformFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"acceptOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressRegistry\",\"outputs\":[{\"internalType\":\"contractIVolcanoAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"}],\"name\":\"buyItem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"}],\"name\":\"cancelListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"}],\"name\":\"cancelOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"createOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"}],\"name\":\"getListing\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startingTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_nftAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingTime\",\"type\":\"uint256\"}],\"name\":\"listItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startingTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"offers\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"updateAddressRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_bundleID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"updateListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"updatePlatformFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"validateItemSold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VolcanoBundleMarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use VolcanoBundleMarketplaceMetaData.ABI instead.
var VolcanoBundleMarketplaceABI = VolcanoBundleMarketplaceMetaData.ABI

// VolcanoBundleMarketplace is an auto generated Go binding around an Ethereum contract.
type VolcanoBundleMarketplace struct {
	VolcanoBundleMarketplaceCaller     // Read-only binding to the contract
	VolcanoBundleMarketplaceTransactor // Write-only binding to the contract
	VolcanoBundleMarketplaceFilterer   // Log filterer for contract events
}

// VolcanoBundleMarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type VolcanoBundleMarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoBundleMarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VolcanoBundleMarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoBundleMarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolcanoBundleMarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoBundleMarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolcanoBundleMarketplaceSession struct {
	Contract     *VolcanoBundleMarketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VolcanoBundleMarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolcanoBundleMarketplaceCallerSession struct {
	Contract *VolcanoBundleMarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// VolcanoBundleMarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolcanoBundleMarketplaceTransactorSession struct {
	Contract     *VolcanoBundleMarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// VolcanoBundleMarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type VolcanoBundleMarketplaceRaw struct {
	Contract *VolcanoBundleMarketplace // Generic contract binding to access the raw methods on
}

// VolcanoBundleMarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolcanoBundleMarketplaceCallerRaw struct {
	Contract *VolcanoBundleMarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// VolcanoBundleMarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolcanoBundleMarketplaceTransactorRaw struct {
	Contract *VolcanoBundleMarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolcanoBundleMarketplace creates a new instance of VolcanoBundleMarketplace, bound to a specific deployed contract.
func NewVolcanoBundleMarketplace(address common.Address, backend bind.ContractBackend) (*VolcanoBundleMarketplace, error) {
	contract, err := bindVolcanoBundleMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplace{VolcanoBundleMarketplaceCaller: VolcanoBundleMarketplaceCaller{contract: contract}, VolcanoBundleMarketplaceTransactor: VolcanoBundleMarketplaceTransactor{contract: contract}, VolcanoBundleMarketplaceFilterer: VolcanoBundleMarketplaceFilterer{contract: contract}}, nil
}

// NewVolcanoBundleMarketplaceCaller creates a new read-only instance of VolcanoBundleMarketplace, bound to a specific deployed contract.
func NewVolcanoBundleMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*VolcanoBundleMarketplaceCaller, error) {
	contract, err := bindVolcanoBundleMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceCaller{contract: contract}, nil
}

// NewVolcanoBundleMarketplaceTransactor creates a new write-only instance of VolcanoBundleMarketplace, bound to a specific deployed contract.
func NewVolcanoBundleMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*VolcanoBundleMarketplaceTransactor, error) {
	contract, err := bindVolcanoBundleMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceTransactor{contract: contract}, nil
}

// NewVolcanoBundleMarketplaceFilterer creates a new log filterer instance of VolcanoBundleMarketplace, bound to a specific deployed contract.
func NewVolcanoBundleMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*VolcanoBundleMarketplaceFilterer, error) {
	contract, err := bindVolcanoBundleMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceFilterer{contract: contract}, nil
}

// bindVolcanoBundleMarketplace binds a generic wrapper to an already deployed contract.
func bindVolcanoBundleMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VolcanoBundleMarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoBundleMarketplace.Contract.VolcanoBundleMarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.VolcanoBundleMarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.VolcanoBundleMarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoBundleMarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.contract.Transact(opts, method, params...)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCaller) AddressRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoBundleMarketplace.contract.Call(opts, &out, "addressRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) AddressRegistry() (common.Address, error) {
	return _VolcanoBundleMarketplace.Contract.AddressRegistry(&_VolcanoBundleMarketplace.CallOpts)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCallerSession) AddressRegistry() (common.Address, error) {
	return _VolcanoBundleMarketplace.Contract.AddressRegistry(&_VolcanoBundleMarketplace.CallOpts)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCaller) FeeReceipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoBundleMarketplace.contract.Call(opts, &out, "feeReceipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) FeeReceipient() (common.Address, error) {
	return _VolcanoBundleMarketplace.Contract.FeeReceipient(&_VolcanoBundleMarketplace.CallOpts)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCallerSession) FeeReceipient() (common.Address, error) {
	return _VolcanoBundleMarketplace.Contract.FeeReceipient(&_VolcanoBundleMarketplace.CallOpts)
}

// GetListing is a free data retrieval call binding the contract method 0x49deb977.
//
// Solidity: function getListing(address _owner, string _bundleID) view returns(address[] nfts, uint256[] tokenIds, uint256[] quantities, uint256 price, uint256 startingTime)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCaller) GetListing(opts *bind.CallOpts, _owner common.Address, _bundleID string) (struct {
	Nfts         []common.Address
	TokenIds     []*big.Int
	Quantities   []*big.Int
	Price        *big.Int
	StartingTime *big.Int
}, error) {
	var out []interface{}
	err := _VolcanoBundleMarketplace.contract.Call(opts, &out, "getListing", _owner, _bundleID)

	outstruct := new(struct {
		Nfts         []common.Address
		TokenIds     []*big.Int
		Quantities   []*big.Int
		Price        *big.Int
		StartingTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nfts = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.TokenIds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Quantities = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartingTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetListing is a free data retrieval call binding the contract method 0x49deb977.
//
// Solidity: function getListing(address _owner, string _bundleID) view returns(address[] nfts, uint256[] tokenIds, uint256[] quantities, uint256 price, uint256 startingTime)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) GetListing(_owner common.Address, _bundleID string) (struct {
	Nfts         []common.Address
	TokenIds     []*big.Int
	Quantities   []*big.Int
	Price        *big.Int
	StartingTime *big.Int
}, error) {
	return _VolcanoBundleMarketplace.Contract.GetListing(&_VolcanoBundleMarketplace.CallOpts, _owner, _bundleID)
}

// GetListing is a free data retrieval call binding the contract method 0x49deb977.
//
// Solidity: function getListing(address _owner, string _bundleID) view returns(address[] nfts, uint256[] tokenIds, uint256[] quantities, uint256 price, uint256 startingTime)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCallerSession) GetListing(_owner common.Address, _bundleID string) (struct {
	Nfts         []common.Address
	TokenIds     []*big.Int
	Quantities   []*big.Int
	Price        *big.Int
	StartingTime *big.Int
}, error) {
	return _VolcanoBundleMarketplace.Contract.GetListing(&_VolcanoBundleMarketplace.CallOpts, _owner, _bundleID)
}

// Listings is a free data retrieval call binding the contract method 0xaa3a6b36.
//
// Solidity: function listings(address , bytes32 ) view returns(address payToken, uint256 price, uint256 startingTime)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCaller) Listings(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	PayToken     common.Address
	Price        *big.Int
	StartingTime *big.Int
}, error) {
	var out []interface{}
	err := _VolcanoBundleMarketplace.contract.Call(opts, &out, "listings", arg0, arg1)

	outstruct := new(struct {
		PayToken     common.Address
		Price        *big.Int
		StartingTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PayToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartingTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xaa3a6b36.
//
// Solidity: function listings(address , bytes32 ) view returns(address payToken, uint256 price, uint256 startingTime)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) Listings(arg0 common.Address, arg1 [32]byte) (struct {
	PayToken     common.Address
	Price        *big.Int
	StartingTime *big.Int
}, error) {
	return _VolcanoBundleMarketplace.Contract.Listings(&_VolcanoBundleMarketplace.CallOpts, arg0, arg1)
}

// Listings is a free data retrieval call binding the contract method 0xaa3a6b36.
//
// Solidity: function listings(address , bytes32 ) view returns(address payToken, uint256 price, uint256 startingTime)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCallerSession) Listings(arg0 common.Address, arg1 [32]byte) (struct {
	PayToken     common.Address
	Price        *big.Int
	StartingTime *big.Int
}, error) {
	return _VolcanoBundleMarketplace.Contract.Listings(&_VolcanoBundleMarketplace.CallOpts, arg0, arg1)
}

// Offers is a free data retrieval call binding the contract method 0xa5851902.
//
// Solidity: function offers(bytes32 , address ) view returns(address payToken, uint256 price, uint256 deadline)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCaller) Offers(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (struct {
	PayToken common.Address
	Price    *big.Int
	Deadline *big.Int
}, error) {
	var out []interface{}
	err := _VolcanoBundleMarketplace.contract.Call(opts, &out, "offers", arg0, arg1)

	outstruct := new(struct {
		PayToken common.Address
		Price    *big.Int
		Deadline *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PayToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Offers is a free data retrieval call binding the contract method 0xa5851902.
//
// Solidity: function offers(bytes32 , address ) view returns(address payToken, uint256 price, uint256 deadline)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) Offers(arg0 [32]byte, arg1 common.Address) (struct {
	PayToken common.Address
	Price    *big.Int
	Deadline *big.Int
}, error) {
	return _VolcanoBundleMarketplace.Contract.Offers(&_VolcanoBundleMarketplace.CallOpts, arg0, arg1)
}

// Offers is a free data retrieval call binding the contract method 0xa5851902.
//
// Solidity: function offers(bytes32 , address ) view returns(address payToken, uint256 price, uint256 deadline)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCallerSession) Offers(arg0 [32]byte, arg1 common.Address) (struct {
	PayToken common.Address
	Price    *big.Int
	Deadline *big.Int
}, error) {
	return _VolcanoBundleMarketplace.Contract.Offers(&_VolcanoBundleMarketplace.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoBundleMarketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) Owner() (common.Address, error) {
	return _VolcanoBundleMarketplace.Contract.Owner(&_VolcanoBundleMarketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCallerSession) Owner() (common.Address, error) {
	return _VolcanoBundleMarketplace.Contract.Owner(&_VolcanoBundleMarketplace.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0xfb8ad6ff.
//
// Solidity: function owners(bytes32 ) view returns(address)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCaller) Owners(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _VolcanoBundleMarketplace.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0xfb8ad6ff.
//
// Solidity: function owners(bytes32 ) view returns(address)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) Owners(arg0 [32]byte) (common.Address, error) {
	return _VolcanoBundleMarketplace.Contract.Owners(&_VolcanoBundleMarketplace.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0xfb8ad6ff.
//
// Solidity: function owners(bytes32 ) view returns(address)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCallerSession) Owners(arg0 [32]byte) (common.Address, error) {
	return _VolcanoBundleMarketplace.Contract.Owners(&_VolcanoBundleMarketplace.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VolcanoBundleMarketplace.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) Paused() (bool, error) {
	return _VolcanoBundleMarketplace.Contract.Paused(&_VolcanoBundleMarketplace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCallerSession) Paused() (bool, error) {
	return _VolcanoBundleMarketplace.Contract.Paused(&_VolcanoBundleMarketplace.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoBundleMarketplace.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) PlatformFee() (*big.Int, error) {
	return _VolcanoBundleMarketplace.Contract.PlatformFee(&_VolcanoBundleMarketplace.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCallerSession) PlatformFee() (*big.Int, error) {
	return _VolcanoBundleMarketplace.Contract.PlatformFee(&_VolcanoBundleMarketplace.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VolcanoBundleMarketplace.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) ProxiableUUID() ([32]byte, error) {
	return _VolcanoBundleMarketplace.Contract.ProxiableUUID(&_VolcanoBundleMarketplace.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceCallerSession) ProxiableUUID() ([32]byte, error) {
	return _VolcanoBundleMarketplace.Contract.ProxiableUUID(&_VolcanoBundleMarketplace.CallOpts)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x84328e20.
//
// Solidity: function acceptOffer(string _bundleID, address _creator) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) AcceptOffer(opts *bind.TransactOpts, _bundleID string, _creator common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "acceptOffer", _bundleID, _creator)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x84328e20.
//
// Solidity: function acceptOffer(string _bundleID, address _creator) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) AcceptOffer(_bundleID string, _creator common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.AcceptOffer(&_VolcanoBundleMarketplace.TransactOpts, _bundleID, _creator)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x84328e20.
//
// Solidity: function acceptOffer(string _bundleID, address _creator) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) AcceptOffer(_bundleID string, _creator common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.AcceptOffer(&_VolcanoBundleMarketplace.TransactOpts, _bundleID, _creator)
}

// BuyItem is a paid mutator transaction binding the contract method 0x9d8820c4.
//
// Solidity: function buyItem(string _bundleID, address _payToken) payable returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) BuyItem(opts *bind.TransactOpts, _bundleID string, _payToken common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "buyItem", _bundleID, _payToken)
}

// BuyItem is a paid mutator transaction binding the contract method 0x9d8820c4.
//
// Solidity: function buyItem(string _bundleID, address _payToken) payable returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) BuyItem(_bundleID string, _payToken common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.BuyItem(&_VolcanoBundleMarketplace.TransactOpts, _bundleID, _payToken)
}

// BuyItem is a paid mutator transaction binding the contract method 0x9d8820c4.
//
// Solidity: function buyItem(string _bundleID, address _payToken) payable returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) BuyItem(_bundleID string, _payToken common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.BuyItem(&_VolcanoBundleMarketplace.TransactOpts, _bundleID, _payToken)
}

// CancelListing is a paid mutator transaction binding the contract method 0x943f2261.
//
// Solidity: function cancelListing(string _bundleID) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) CancelListing(opts *bind.TransactOpts, _bundleID string) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "cancelListing", _bundleID)
}

// CancelListing is a paid mutator transaction binding the contract method 0x943f2261.
//
// Solidity: function cancelListing(string _bundleID) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) CancelListing(_bundleID string) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.CancelListing(&_VolcanoBundleMarketplace.TransactOpts, _bundleID)
}

// CancelListing is a paid mutator transaction binding the contract method 0x943f2261.
//
// Solidity: function cancelListing(string _bundleID) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) CancelListing(_bundleID string) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.CancelListing(&_VolcanoBundleMarketplace.TransactOpts, _bundleID)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x2d63ce27.
//
// Solidity: function cancelOffer(string _bundleID) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) CancelOffer(opts *bind.TransactOpts, _bundleID string) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "cancelOffer", _bundleID)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x2d63ce27.
//
// Solidity: function cancelOffer(string _bundleID) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) CancelOffer(_bundleID string) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.CancelOffer(&_VolcanoBundleMarketplace.TransactOpts, _bundleID)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x2d63ce27.
//
// Solidity: function cancelOffer(string _bundleID) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) CancelOffer(_bundleID string) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.CancelOffer(&_VolcanoBundleMarketplace.TransactOpts, _bundleID)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xe831eebf.
//
// Solidity: function createOffer(string _bundleID, address _payToken, uint256 _price, uint256 _deadline) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) CreateOffer(opts *bind.TransactOpts, _bundleID string, _payToken common.Address, _price *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "createOffer", _bundleID, _payToken, _price, _deadline)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xe831eebf.
//
// Solidity: function createOffer(string _bundleID, address _payToken, uint256 _price, uint256 _deadline) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) CreateOffer(_bundleID string, _payToken common.Address, _price *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.CreateOffer(&_VolcanoBundleMarketplace.TransactOpts, _bundleID, _payToken, _price, _deadline)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xe831eebf.
//
// Solidity: function createOffer(string _bundleID, address _payToken, uint256 _price, uint256 _deadline) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) CreateOffer(_bundleID string, _payToken common.Address, _price *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.CreateOffer(&_VolcanoBundleMarketplace.TransactOpts, _bundleID, _payToken, _price, _deadline)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _feeRecipient, uint256 _platformFee) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) Initialize(opts *bind.TransactOpts, _feeRecipient common.Address, _platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "initialize", _feeRecipient, _platformFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _feeRecipient, uint256 _platformFee) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) Initialize(_feeRecipient common.Address, _platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.Initialize(&_VolcanoBundleMarketplace.TransactOpts, _feeRecipient, _platformFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _feeRecipient, uint256 _platformFee) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) Initialize(_feeRecipient common.Address, _platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.Initialize(&_VolcanoBundleMarketplace.TransactOpts, _feeRecipient, _platformFee)
}

// ListItem is a paid mutator transaction binding the contract method 0x4994bfe5.
//
// Solidity: function listItem(string _bundleID, address[] _nftAddresses, uint256[] _tokenIds, uint256[] _quantities, address _payToken, uint256 _price, uint256 _startingTime) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) ListItem(opts *bind.TransactOpts, _bundleID string, _nftAddresses []common.Address, _tokenIds []*big.Int, _quantities []*big.Int, _payToken common.Address, _price *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "listItem", _bundleID, _nftAddresses, _tokenIds, _quantities, _payToken, _price, _startingTime)
}

// ListItem is a paid mutator transaction binding the contract method 0x4994bfe5.
//
// Solidity: function listItem(string _bundleID, address[] _nftAddresses, uint256[] _tokenIds, uint256[] _quantities, address _payToken, uint256 _price, uint256 _startingTime) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) ListItem(_bundleID string, _nftAddresses []common.Address, _tokenIds []*big.Int, _quantities []*big.Int, _payToken common.Address, _price *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.ListItem(&_VolcanoBundleMarketplace.TransactOpts, _bundleID, _nftAddresses, _tokenIds, _quantities, _payToken, _price, _startingTime)
}

// ListItem is a paid mutator transaction binding the contract method 0x4994bfe5.
//
// Solidity: function listItem(string _bundleID, address[] _nftAddresses, uint256[] _tokenIds, uint256[] _quantities, address _payToken, uint256 _price, uint256 _startingTime) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) ListItem(_bundleID string, _nftAddresses []common.Address, _tokenIds []*big.Int, _quantities []*big.Int, _payToken common.Address, _price *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.ListItem(&_VolcanoBundleMarketplace.TransactOpts, _bundleID, _nftAddresses, _tokenIds, _quantities, _payToken, _price, _startingTime)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.RenounceOwnership(&_VolcanoBundleMarketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.RenounceOwnership(&_VolcanoBundleMarketplace.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.TransferOwnership(&_VolcanoBundleMarketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.TransferOwnership(&_VolcanoBundleMarketplace.TransactOpts, newOwner)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) UpdateAddressRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "updateAddressRegistry", _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.UpdateAddressRegistry(&_VolcanoBundleMarketplace.TransactOpts, _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.UpdateAddressRegistry(&_VolcanoBundleMarketplace.TransactOpts, _registry)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xab0962dd.
//
// Solidity: function updateListing(string _bundleID, address _payToken, uint256 _newPrice) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) UpdateListing(opts *bind.TransactOpts, _bundleID string, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "updateListing", _bundleID, _payToken, _newPrice)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xab0962dd.
//
// Solidity: function updateListing(string _bundleID, address _payToken, uint256 _newPrice) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) UpdateListing(_bundleID string, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.UpdateListing(&_VolcanoBundleMarketplace.TransactOpts, _bundleID, _payToken, _newPrice)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xab0962dd.
//
// Solidity: function updateListing(string _bundleID, address _payToken, uint256 _newPrice) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) UpdateListing(_bundleID string, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.UpdateListing(&_VolcanoBundleMarketplace.TransactOpts, _bundleID, _payToken, _newPrice)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) UpdatePlatformFee(opts *bind.TransactOpts, _platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "updatePlatformFee", _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.UpdatePlatformFee(&_VolcanoBundleMarketplace.TransactOpts, _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.UpdatePlatformFee(&_VolcanoBundleMarketplace.TransactOpts, _platformFee)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) UpdatePlatformFeeRecipient(opts *bind.TransactOpts, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "updatePlatformFeeRecipient", _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.UpdatePlatformFeeRecipient(&_VolcanoBundleMarketplace.TransactOpts, _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.UpdatePlatformFeeRecipient(&_VolcanoBundleMarketplace.TransactOpts, _platformFeeRecipient)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.UpgradeTo(&_VolcanoBundleMarketplace.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.UpgradeTo(&_VolcanoBundleMarketplace.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.UpgradeToAndCall(&_VolcanoBundleMarketplace.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.UpgradeToAndCall(&_VolcanoBundleMarketplace.TransactOpts, newImplementation, data)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0x27c1f4f4.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, uint256 _quantity) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactor) ValidateItemSold(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.contract.Transact(opts, "validateItemSold", _nftAddress, _tokenId, _quantity)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0x27c1f4f4.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, uint256 _quantity) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceSession) ValidateItemSold(_nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.ValidateItemSold(&_VolcanoBundleMarketplace.TransactOpts, _nftAddress, _tokenId, _quantity)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0x27c1f4f4.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, uint256 _quantity) returns()
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceTransactorSession) ValidateItemSold(_nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int) (*types.Transaction, error) {
	return _VolcanoBundleMarketplace.Contract.ValidateItemSold(&_VolcanoBundleMarketplace.TransactOpts, _nftAddress, _tokenId, _quantity)
}

// VolcanoBundleMarketplaceAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceAdminChangedIterator struct {
	Event *VolcanoBundleMarketplaceAdminChanged // Event containing the contract specifics and raw log

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
func (it *VolcanoBundleMarketplaceAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoBundleMarketplaceAdminChanged)
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
		it.Event = new(VolcanoBundleMarketplaceAdminChanged)
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
func (it *VolcanoBundleMarketplaceAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoBundleMarketplaceAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoBundleMarketplaceAdminChanged represents a AdminChanged event raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*VolcanoBundleMarketplaceAdminChangedIterator, error) {

	logs, sub, err := _VolcanoBundleMarketplace.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceAdminChangedIterator{contract: _VolcanoBundleMarketplace.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *VolcanoBundleMarketplaceAdminChanged) (event.Subscription, error) {

	logs, sub, err := _VolcanoBundleMarketplace.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoBundleMarketplaceAdminChanged)
				if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) ParseAdminChanged(log types.Log) (*VolcanoBundleMarketplaceAdminChanged, error) {
	event := new(VolcanoBundleMarketplaceAdminChanged)
	if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoBundleMarketplaceBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceBeaconUpgradedIterator struct {
	Event *VolcanoBundleMarketplaceBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *VolcanoBundleMarketplaceBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoBundleMarketplaceBeaconUpgraded)
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
		it.Event = new(VolcanoBundleMarketplaceBeaconUpgraded)
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
func (it *VolcanoBundleMarketplaceBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoBundleMarketplaceBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoBundleMarketplaceBeaconUpgraded represents a BeaconUpgraded event raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*VolcanoBundleMarketplaceBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceBeaconUpgradedIterator{contract: _VolcanoBundleMarketplace.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *VolcanoBundleMarketplaceBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoBundleMarketplaceBeaconUpgraded)
				if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) ParseBeaconUpgraded(log types.Log) (*VolcanoBundleMarketplaceBeaconUpgraded, error) {
	event := new(VolcanoBundleMarketplaceBeaconUpgraded)
	if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoBundleMarketplaceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceInitializedIterator struct {
	Event *VolcanoBundleMarketplaceInitialized // Event containing the contract specifics and raw log

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
func (it *VolcanoBundleMarketplaceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoBundleMarketplaceInitialized)
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
		it.Event = new(VolcanoBundleMarketplaceInitialized)
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
func (it *VolcanoBundleMarketplaceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoBundleMarketplaceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoBundleMarketplaceInitialized represents a Initialized event raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) FilterInitialized(opts *bind.FilterOpts) (*VolcanoBundleMarketplaceInitializedIterator, error) {

	logs, sub, err := _VolcanoBundleMarketplace.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceInitializedIterator{contract: _VolcanoBundleMarketplace.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VolcanoBundleMarketplaceInitialized) (event.Subscription, error) {

	logs, sub, err := _VolcanoBundleMarketplace.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoBundleMarketplaceInitialized)
				if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) ParseInitialized(log types.Log) (*VolcanoBundleMarketplaceInitialized, error) {
	event := new(VolcanoBundleMarketplaceInitialized)
	if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoBundleMarketplaceItemCanceledIterator is returned from FilterItemCanceled and is used to iterate over the raw logs and unpacked data for ItemCanceled events raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceItemCanceledIterator struct {
	Event *VolcanoBundleMarketplaceItemCanceled // Event containing the contract specifics and raw log

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
func (it *VolcanoBundleMarketplaceItemCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoBundleMarketplaceItemCanceled)
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
		it.Event = new(VolcanoBundleMarketplaceItemCanceled)
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
func (it *VolcanoBundleMarketplaceItemCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoBundleMarketplaceItemCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoBundleMarketplaceItemCanceled represents a ItemCanceled event raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceItemCanceled struct {
	Owner    common.Address
	BundleID string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterItemCanceled is a free log retrieval operation binding the contract event 0x71a0e856c3470e868cc72d9c038e190004e167ea6490f2063682cd37679a302a.
//
// Solidity: event ItemCanceled(address indexed owner, string bundleID)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) FilterItemCanceled(opts *bind.FilterOpts, owner []common.Address) (*VolcanoBundleMarketplaceItemCanceledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.FilterLogs(opts, "ItemCanceled", ownerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceItemCanceledIterator{contract: _VolcanoBundleMarketplace.contract, event: "ItemCanceled", logs: logs, sub: sub}, nil
}

// WatchItemCanceled is a free log subscription operation binding the contract event 0x71a0e856c3470e868cc72d9c038e190004e167ea6490f2063682cd37679a302a.
//
// Solidity: event ItemCanceled(address indexed owner, string bundleID)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) WatchItemCanceled(opts *bind.WatchOpts, sink chan<- *VolcanoBundleMarketplaceItemCanceled, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.WatchLogs(opts, "ItemCanceled", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoBundleMarketplaceItemCanceled)
				if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "ItemCanceled", log); err != nil {
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

// ParseItemCanceled is a log parse operation binding the contract event 0x71a0e856c3470e868cc72d9c038e190004e167ea6490f2063682cd37679a302a.
//
// Solidity: event ItemCanceled(address indexed owner, string bundleID)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) ParseItemCanceled(log types.Log) (*VolcanoBundleMarketplaceItemCanceled, error) {
	event := new(VolcanoBundleMarketplaceItemCanceled)
	if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "ItemCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoBundleMarketplaceItemListedIterator is returned from FilterItemListed and is used to iterate over the raw logs and unpacked data for ItemListed events raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceItemListedIterator struct {
	Event *VolcanoBundleMarketplaceItemListed // Event containing the contract specifics and raw log

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
func (it *VolcanoBundleMarketplaceItemListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoBundleMarketplaceItemListed)
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
		it.Event = new(VolcanoBundleMarketplaceItemListed)
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
func (it *VolcanoBundleMarketplaceItemListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoBundleMarketplaceItemListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoBundleMarketplaceItemListed represents a ItemListed event raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceItemListed struct {
	Owner        common.Address
	BundleID     string
	PayToken     common.Address
	Price        *big.Int
	StartingTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterItemListed is a free log retrieval operation binding the contract event 0x1dd1f9de6505bd1f4c42382f88f7e9475dedc3c10ff5bd77cf14490a5d74fab2.
//
// Solidity: event ItemListed(address indexed owner, string bundleID, address payToken, uint256 price, uint256 startingTime)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) FilterItemListed(opts *bind.FilterOpts, owner []common.Address) (*VolcanoBundleMarketplaceItemListedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.FilterLogs(opts, "ItemListed", ownerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceItemListedIterator{contract: _VolcanoBundleMarketplace.contract, event: "ItemListed", logs: logs, sub: sub}, nil
}

// WatchItemListed is a free log subscription operation binding the contract event 0x1dd1f9de6505bd1f4c42382f88f7e9475dedc3c10ff5bd77cf14490a5d74fab2.
//
// Solidity: event ItemListed(address indexed owner, string bundleID, address payToken, uint256 price, uint256 startingTime)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) WatchItemListed(opts *bind.WatchOpts, sink chan<- *VolcanoBundleMarketplaceItemListed, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.WatchLogs(opts, "ItemListed", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoBundleMarketplaceItemListed)
				if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "ItemListed", log); err != nil {
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

// ParseItemListed is a log parse operation binding the contract event 0x1dd1f9de6505bd1f4c42382f88f7e9475dedc3c10ff5bd77cf14490a5d74fab2.
//
// Solidity: event ItemListed(address indexed owner, string bundleID, address payToken, uint256 price, uint256 startingTime)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) ParseItemListed(log types.Log) (*VolcanoBundleMarketplaceItemListed, error) {
	event := new(VolcanoBundleMarketplaceItemListed)
	if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "ItemListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoBundleMarketplaceItemSoldIterator is returned from FilterItemSold and is used to iterate over the raw logs and unpacked data for ItemSold events raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceItemSoldIterator struct {
	Event *VolcanoBundleMarketplaceItemSold // Event containing the contract specifics and raw log

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
func (it *VolcanoBundleMarketplaceItemSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoBundleMarketplaceItemSold)
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
		it.Event = new(VolcanoBundleMarketplaceItemSold)
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
func (it *VolcanoBundleMarketplaceItemSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoBundleMarketplaceItemSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoBundleMarketplaceItemSold represents a ItemSold event raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceItemSold struct {
	Seller        common.Address
	Buyer         common.Address
	BundleID      string
	PayToken      common.Address
	PayTokenPrice *big.Int
	Price         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterItemSold is a free log retrieval operation binding the contract event 0x586a7960ce4f3631680ea77efd69fa758b24388a9a5183f452209aa9f1c74bfd.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, string bundleID, address payToken, int256 payTokenPrice, uint256 price)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) FilterItemSold(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address) (*VolcanoBundleMarketplaceItemSoldIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.FilterLogs(opts, "ItemSold", sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceItemSoldIterator{contract: _VolcanoBundleMarketplace.contract, event: "ItemSold", logs: logs, sub: sub}, nil
}

// WatchItemSold is a free log subscription operation binding the contract event 0x586a7960ce4f3631680ea77efd69fa758b24388a9a5183f452209aa9f1c74bfd.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, string bundleID, address payToken, int256 payTokenPrice, uint256 price)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) WatchItemSold(opts *bind.WatchOpts, sink chan<- *VolcanoBundleMarketplaceItemSold, seller []common.Address, buyer []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.WatchLogs(opts, "ItemSold", sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoBundleMarketplaceItemSold)
				if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "ItemSold", log); err != nil {
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

// ParseItemSold is a log parse operation binding the contract event 0x586a7960ce4f3631680ea77efd69fa758b24388a9a5183f452209aa9f1c74bfd.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, string bundleID, address payToken, int256 payTokenPrice, uint256 price)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) ParseItemSold(log types.Log) (*VolcanoBundleMarketplaceItemSold, error) {
	event := new(VolcanoBundleMarketplaceItemSold)
	if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "ItemSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoBundleMarketplaceItemUpdatedIterator is returned from FilterItemUpdated and is used to iterate over the raw logs and unpacked data for ItemUpdated events raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceItemUpdatedIterator struct {
	Event *VolcanoBundleMarketplaceItemUpdated // Event containing the contract specifics and raw log

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
func (it *VolcanoBundleMarketplaceItemUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoBundleMarketplaceItemUpdated)
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
		it.Event = new(VolcanoBundleMarketplaceItemUpdated)
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
func (it *VolcanoBundleMarketplaceItemUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoBundleMarketplaceItemUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoBundleMarketplaceItemUpdated represents a ItemUpdated event raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceItemUpdated struct {
	Owner    common.Address
	BundleID string
	Nft      []common.Address
	TokenId  []*big.Int
	Quantity []*big.Int
	PayToken common.Address
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterItemUpdated is a free log retrieval operation binding the contract event 0xbf004834f59fa5d831659d6ca458fe7e8675daefebc459ac53a3cab2b6816bf4.
//
// Solidity: event ItemUpdated(address indexed owner, string bundleID, address[] nft, uint256[] tokenId, uint256[] quantity, address payToken, uint256 newPrice)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) FilterItemUpdated(opts *bind.FilterOpts, owner []common.Address) (*VolcanoBundleMarketplaceItemUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.FilterLogs(opts, "ItemUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceItemUpdatedIterator{contract: _VolcanoBundleMarketplace.contract, event: "ItemUpdated", logs: logs, sub: sub}, nil
}

// WatchItemUpdated is a free log subscription operation binding the contract event 0xbf004834f59fa5d831659d6ca458fe7e8675daefebc459ac53a3cab2b6816bf4.
//
// Solidity: event ItemUpdated(address indexed owner, string bundleID, address[] nft, uint256[] tokenId, uint256[] quantity, address payToken, uint256 newPrice)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) WatchItemUpdated(opts *bind.WatchOpts, sink chan<- *VolcanoBundleMarketplaceItemUpdated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.WatchLogs(opts, "ItemUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoBundleMarketplaceItemUpdated)
				if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "ItemUpdated", log); err != nil {
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

// ParseItemUpdated is a log parse operation binding the contract event 0xbf004834f59fa5d831659d6ca458fe7e8675daefebc459ac53a3cab2b6816bf4.
//
// Solidity: event ItemUpdated(address indexed owner, string bundleID, address[] nft, uint256[] tokenId, uint256[] quantity, address payToken, uint256 newPrice)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) ParseItemUpdated(log types.Log) (*VolcanoBundleMarketplaceItemUpdated, error) {
	event := new(VolcanoBundleMarketplaceItemUpdated)
	if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "ItemUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoBundleMarketplaceOfferCanceledIterator is returned from FilterOfferCanceled and is used to iterate over the raw logs and unpacked data for OfferCanceled events raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceOfferCanceledIterator struct {
	Event *VolcanoBundleMarketplaceOfferCanceled // Event containing the contract specifics and raw log

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
func (it *VolcanoBundleMarketplaceOfferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoBundleMarketplaceOfferCanceled)
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
		it.Event = new(VolcanoBundleMarketplaceOfferCanceled)
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
func (it *VolcanoBundleMarketplaceOfferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoBundleMarketplaceOfferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoBundleMarketplaceOfferCanceled represents a OfferCanceled event raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceOfferCanceled struct {
	Creator  common.Address
	BundleID string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOfferCanceled is a free log retrieval operation binding the contract event 0x7d7a082d5e9005cb920fdc04023f730c7776927c8a728347783ac7a8d2b31e7c.
//
// Solidity: event OfferCanceled(address indexed creator, string bundleID)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) FilterOfferCanceled(opts *bind.FilterOpts, creator []common.Address) (*VolcanoBundleMarketplaceOfferCanceledIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.FilterLogs(opts, "OfferCanceled", creatorRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceOfferCanceledIterator{contract: _VolcanoBundleMarketplace.contract, event: "OfferCanceled", logs: logs, sub: sub}, nil
}

// WatchOfferCanceled is a free log subscription operation binding the contract event 0x7d7a082d5e9005cb920fdc04023f730c7776927c8a728347783ac7a8d2b31e7c.
//
// Solidity: event OfferCanceled(address indexed creator, string bundleID)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) WatchOfferCanceled(opts *bind.WatchOpts, sink chan<- *VolcanoBundleMarketplaceOfferCanceled, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.WatchLogs(opts, "OfferCanceled", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoBundleMarketplaceOfferCanceled)
				if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
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

// ParseOfferCanceled is a log parse operation binding the contract event 0x7d7a082d5e9005cb920fdc04023f730c7776927c8a728347783ac7a8d2b31e7c.
//
// Solidity: event OfferCanceled(address indexed creator, string bundleID)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) ParseOfferCanceled(log types.Log) (*VolcanoBundleMarketplaceOfferCanceled, error) {
	event := new(VolcanoBundleMarketplaceOfferCanceled)
	if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoBundleMarketplaceOfferCreatedIterator is returned from FilterOfferCreated and is used to iterate over the raw logs and unpacked data for OfferCreated events raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceOfferCreatedIterator struct {
	Event *VolcanoBundleMarketplaceOfferCreated // Event containing the contract specifics and raw log

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
func (it *VolcanoBundleMarketplaceOfferCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoBundleMarketplaceOfferCreated)
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
		it.Event = new(VolcanoBundleMarketplaceOfferCreated)
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
func (it *VolcanoBundleMarketplaceOfferCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoBundleMarketplaceOfferCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoBundleMarketplaceOfferCreated represents a OfferCreated event raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceOfferCreated struct {
	Creator  common.Address
	BundleID string
	PayToken common.Address
	Price    *big.Int
	Deadline *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOfferCreated is a free log retrieval operation binding the contract event 0x6c59d59bc0c5e094ffd184dd57ea02b8f8e2e2299fd34ffebdd4888864b547b9.
//
// Solidity: event OfferCreated(address indexed creator, string bundleID, address payToken, uint256 price, uint256 deadline)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) FilterOfferCreated(opts *bind.FilterOpts, creator []common.Address) (*VolcanoBundleMarketplaceOfferCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.FilterLogs(opts, "OfferCreated", creatorRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceOfferCreatedIterator{contract: _VolcanoBundleMarketplace.contract, event: "OfferCreated", logs: logs, sub: sub}, nil
}

// WatchOfferCreated is a free log subscription operation binding the contract event 0x6c59d59bc0c5e094ffd184dd57ea02b8f8e2e2299fd34ffebdd4888864b547b9.
//
// Solidity: event OfferCreated(address indexed creator, string bundleID, address payToken, uint256 price, uint256 deadline)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) WatchOfferCreated(opts *bind.WatchOpts, sink chan<- *VolcanoBundleMarketplaceOfferCreated, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.WatchLogs(opts, "OfferCreated", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoBundleMarketplaceOfferCreated)
				if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "OfferCreated", log); err != nil {
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

// ParseOfferCreated is a log parse operation binding the contract event 0x6c59d59bc0c5e094ffd184dd57ea02b8f8e2e2299fd34ffebdd4888864b547b9.
//
// Solidity: event OfferCreated(address indexed creator, string bundleID, address payToken, uint256 price, uint256 deadline)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) ParseOfferCreated(log types.Log) (*VolcanoBundleMarketplaceOfferCreated, error) {
	event := new(VolcanoBundleMarketplaceOfferCreated)
	if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "OfferCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoBundleMarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceOwnershipTransferredIterator struct {
	Event *VolcanoBundleMarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VolcanoBundleMarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoBundleMarketplaceOwnershipTransferred)
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
		it.Event = new(VolcanoBundleMarketplaceOwnershipTransferred)
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
func (it *VolcanoBundleMarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoBundleMarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoBundleMarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VolcanoBundleMarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceOwnershipTransferredIterator{contract: _VolcanoBundleMarketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VolcanoBundleMarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoBundleMarketplaceOwnershipTransferred)
				if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*VolcanoBundleMarketplaceOwnershipTransferred, error) {
	event := new(VolcanoBundleMarketplaceOwnershipTransferred)
	if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoBundleMarketplacePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplacePausedIterator struct {
	Event *VolcanoBundleMarketplacePaused // Event containing the contract specifics and raw log

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
func (it *VolcanoBundleMarketplacePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoBundleMarketplacePaused)
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
		it.Event = new(VolcanoBundleMarketplacePaused)
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
func (it *VolcanoBundleMarketplacePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoBundleMarketplacePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoBundleMarketplacePaused represents a Paused event raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplacePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) FilterPaused(opts *bind.FilterOpts) (*VolcanoBundleMarketplacePausedIterator, error) {

	logs, sub, err := _VolcanoBundleMarketplace.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplacePausedIterator{contract: _VolcanoBundleMarketplace.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VolcanoBundleMarketplacePaused) (event.Subscription, error) {

	logs, sub, err := _VolcanoBundleMarketplace.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoBundleMarketplacePaused)
				if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) ParsePaused(log types.Log) (*VolcanoBundleMarketplacePaused, error) {
	event := new(VolcanoBundleMarketplacePaused)
	if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoBundleMarketplaceUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceUnpausedIterator struct {
	Event *VolcanoBundleMarketplaceUnpaused // Event containing the contract specifics and raw log

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
func (it *VolcanoBundleMarketplaceUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoBundleMarketplaceUnpaused)
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
		it.Event = new(VolcanoBundleMarketplaceUnpaused)
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
func (it *VolcanoBundleMarketplaceUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoBundleMarketplaceUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoBundleMarketplaceUnpaused represents a Unpaused event raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VolcanoBundleMarketplaceUnpausedIterator, error) {

	logs, sub, err := _VolcanoBundleMarketplace.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceUnpausedIterator{contract: _VolcanoBundleMarketplace.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VolcanoBundleMarketplaceUnpaused) (event.Subscription, error) {

	logs, sub, err := _VolcanoBundleMarketplace.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoBundleMarketplaceUnpaused)
				if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) ParseUnpaused(log types.Log) (*VolcanoBundleMarketplaceUnpaused, error) {
	event := new(VolcanoBundleMarketplaceUnpaused)
	if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoBundleMarketplaceUpdatePlatformFeeIterator is returned from FilterUpdatePlatformFee and is used to iterate over the raw logs and unpacked data for UpdatePlatformFee events raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceUpdatePlatformFeeIterator struct {
	Event *VolcanoBundleMarketplaceUpdatePlatformFee // Event containing the contract specifics and raw log

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
func (it *VolcanoBundleMarketplaceUpdatePlatformFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoBundleMarketplaceUpdatePlatformFee)
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
		it.Event = new(VolcanoBundleMarketplaceUpdatePlatformFee)
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
func (it *VolcanoBundleMarketplaceUpdatePlatformFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoBundleMarketplaceUpdatePlatformFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoBundleMarketplaceUpdatePlatformFee represents a UpdatePlatformFee event raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceUpdatePlatformFee struct {
	PlatformFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFee is a free log retrieval operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) FilterUpdatePlatformFee(opts *bind.FilterOpts) (*VolcanoBundleMarketplaceUpdatePlatformFeeIterator, error) {

	logs, sub, err := _VolcanoBundleMarketplace.contract.FilterLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceUpdatePlatformFeeIterator{contract: _VolcanoBundleMarketplace.contract, event: "UpdatePlatformFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFee is a free log subscription operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) WatchUpdatePlatformFee(opts *bind.WatchOpts, sink chan<- *VolcanoBundleMarketplaceUpdatePlatformFee) (event.Subscription, error) {

	logs, sub, err := _VolcanoBundleMarketplace.contract.WatchLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoBundleMarketplaceUpdatePlatformFee)
				if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
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
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) ParseUpdatePlatformFee(log types.Log) (*VolcanoBundleMarketplaceUpdatePlatformFee, error) {
	event := new(VolcanoBundleMarketplaceUpdatePlatformFee)
	if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoBundleMarketplaceUpdatePlatformFeeRecipientIterator is returned from FilterUpdatePlatformFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdatePlatformFeeRecipient events raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceUpdatePlatformFeeRecipientIterator struct {
	Event *VolcanoBundleMarketplaceUpdatePlatformFeeRecipient // Event containing the contract specifics and raw log

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
func (it *VolcanoBundleMarketplaceUpdatePlatformFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoBundleMarketplaceUpdatePlatformFeeRecipient)
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
		it.Event = new(VolcanoBundleMarketplaceUpdatePlatformFeeRecipient)
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
func (it *VolcanoBundleMarketplaceUpdatePlatformFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoBundleMarketplaceUpdatePlatformFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoBundleMarketplaceUpdatePlatformFeeRecipient represents a UpdatePlatformFeeRecipient event raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceUpdatePlatformFeeRecipient struct {
	PlatformFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFeeRecipient is a free log retrieval operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) FilterUpdatePlatformFeeRecipient(opts *bind.FilterOpts) (*VolcanoBundleMarketplaceUpdatePlatformFeeRecipientIterator, error) {

	logs, sub, err := _VolcanoBundleMarketplace.contract.FilterLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceUpdatePlatformFeeRecipientIterator{contract: _VolcanoBundleMarketplace.contract, event: "UpdatePlatformFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFeeRecipient is a free log subscription operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) WatchUpdatePlatformFeeRecipient(opts *bind.WatchOpts, sink chan<- *VolcanoBundleMarketplaceUpdatePlatformFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _VolcanoBundleMarketplace.contract.WatchLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoBundleMarketplaceUpdatePlatformFeeRecipient)
				if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
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
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) ParseUpdatePlatformFeeRecipient(log types.Log) (*VolcanoBundleMarketplaceUpdatePlatformFeeRecipient, error) {
	event := new(VolcanoBundleMarketplaceUpdatePlatformFeeRecipient)
	if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoBundleMarketplaceUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceUpgradedIterator struct {
	Event *VolcanoBundleMarketplaceUpgraded // Event containing the contract specifics and raw log

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
func (it *VolcanoBundleMarketplaceUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoBundleMarketplaceUpgraded)
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
		it.Event = new(VolcanoBundleMarketplaceUpgraded)
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
func (it *VolcanoBundleMarketplaceUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoBundleMarketplaceUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoBundleMarketplaceUpgraded represents a Upgraded event raised by the VolcanoBundleMarketplace contract.
type VolcanoBundleMarketplaceUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VolcanoBundleMarketplaceUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoBundleMarketplaceUpgradedIterator{contract: _VolcanoBundleMarketplace.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VolcanoBundleMarketplaceUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VolcanoBundleMarketplace.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoBundleMarketplaceUpgraded)
				if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VolcanoBundleMarketplace *VolcanoBundleMarketplaceFilterer) ParseUpgraded(log types.Log) (*VolcanoBundleMarketplaceUpgraded, error) {
	event := new(VolcanoBundleMarketplaceUpgraded)
	if err := _VolcanoBundleMarketplace.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
