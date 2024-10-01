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

// VolcanoMarketplaceMetaData contains all meta data concerning the VolcanoMarketplace contract.
var VolcanoMarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ItemCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startingTime\",\"type\":\"uint256\"}],\"name\":\"ItemListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"payTokenPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"}],\"name\":\"ItemSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"ItemUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OfferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"OfferCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"platformFee\",\"type\":\"uint16\"}],\"name\":\"UpdatePlatformFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"UpdatePlatformFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"acceptOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressRegistry\",\"outputs\":[{\"internalType\":\"contractIVolcanoAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"buyItem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collectionRoyalties\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"royalty\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"createOffer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingTime\",\"type\":\"uint256\"}],\"name\":\"listItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startingTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"offers\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_royalty\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"registerCollectionRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_royalty\",\"type\":\"uint16\"}],\"name\":\"registerRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"royalties\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"updateAddressRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"updateListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_platformFee\",\"type\":\"uint16\"}],\"name\":\"updatePlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"updatePlatformFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"validateItemSold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// VolcanoMarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use VolcanoMarketplaceMetaData.ABI instead.
var VolcanoMarketplaceABI = VolcanoMarketplaceMetaData.ABI

// VolcanoMarketplace is an auto generated Go binding around an Ethereum contract.
type VolcanoMarketplace struct {
	VolcanoMarketplaceCaller     // Read-only binding to the contract
	VolcanoMarketplaceTransactor // Write-only binding to the contract
	VolcanoMarketplaceFilterer   // Log filterer for contract events
}

// VolcanoMarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type VolcanoMarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoMarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VolcanoMarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoMarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolcanoMarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoMarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolcanoMarketplaceSession struct {
	Contract     *VolcanoMarketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VolcanoMarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolcanoMarketplaceCallerSession struct {
	Contract *VolcanoMarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// VolcanoMarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolcanoMarketplaceTransactorSession struct {
	Contract     *VolcanoMarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// VolcanoMarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type VolcanoMarketplaceRaw struct {
	Contract *VolcanoMarketplace // Generic contract binding to access the raw methods on
}

// VolcanoMarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolcanoMarketplaceCallerRaw struct {
	Contract *VolcanoMarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// VolcanoMarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolcanoMarketplaceTransactorRaw struct {
	Contract *VolcanoMarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolcanoMarketplace creates a new instance of VolcanoMarketplace, bound to a specific deployed contract.
func NewVolcanoMarketplace(address common.Address, backend bind.ContractBackend) (*VolcanoMarketplace, error) {
	contract, err := bindVolcanoMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplace{VolcanoMarketplaceCaller: VolcanoMarketplaceCaller{contract: contract}, VolcanoMarketplaceTransactor: VolcanoMarketplaceTransactor{contract: contract}, VolcanoMarketplaceFilterer: VolcanoMarketplaceFilterer{contract: contract}}, nil
}

// NewVolcanoMarketplaceCaller creates a new read-only instance of VolcanoMarketplace, bound to a specific deployed contract.
func NewVolcanoMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*VolcanoMarketplaceCaller, error) {
	contract, err := bindVolcanoMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceCaller{contract: contract}, nil
}

// NewVolcanoMarketplaceTransactor creates a new write-only instance of VolcanoMarketplace, bound to a specific deployed contract.
func NewVolcanoMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*VolcanoMarketplaceTransactor, error) {
	contract, err := bindVolcanoMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceTransactor{contract: contract}, nil
}

// NewVolcanoMarketplaceFilterer creates a new log filterer instance of VolcanoMarketplace, bound to a specific deployed contract.
func NewVolcanoMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*VolcanoMarketplaceFilterer, error) {
	contract, err := bindVolcanoMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceFilterer{contract: contract}, nil
}

// bindVolcanoMarketplace binds a generic wrapper to an already deployed contract.
func bindVolcanoMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VolcanoMarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoMarketplace *VolcanoMarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoMarketplace.Contract.VolcanoMarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoMarketplace *VolcanoMarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.VolcanoMarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoMarketplace *VolcanoMarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.VolcanoMarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoMarketplace *VolcanoMarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoMarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.contract.Transact(opts, method, params...)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_VolcanoMarketplace *VolcanoMarketplaceCaller) AddressRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoMarketplace.contract.Call(opts, &out, "addressRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_VolcanoMarketplace *VolcanoMarketplaceSession) AddressRegistry() (common.Address, error) {
	return _VolcanoMarketplace.Contract.AddressRegistry(&_VolcanoMarketplace.CallOpts)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_VolcanoMarketplace *VolcanoMarketplaceCallerSession) AddressRegistry() (common.Address, error) {
	return _VolcanoMarketplace.Contract.AddressRegistry(&_VolcanoMarketplace.CallOpts)
}

// CollectionRoyalties is a free data retrieval call binding the contract method 0x57b275b4.
//
// Solidity: function collectionRoyalties(address ) view returns(uint16 royalty, address creator, address feeRecipient)
func (_VolcanoMarketplace *VolcanoMarketplaceCaller) CollectionRoyalties(opts *bind.CallOpts, arg0 common.Address) (struct {
	Royalty      uint16
	Creator      common.Address
	FeeRecipient common.Address
}, error) {
	var out []interface{}
	err := _VolcanoMarketplace.contract.Call(opts, &out, "collectionRoyalties", arg0)

	outstruct := new(struct {
		Royalty      uint16
		Creator      common.Address
		FeeRecipient common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Royalty = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.Creator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.FeeRecipient = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CollectionRoyalties is a free data retrieval call binding the contract method 0x57b275b4.
//
// Solidity: function collectionRoyalties(address ) view returns(uint16 royalty, address creator, address feeRecipient)
func (_VolcanoMarketplace *VolcanoMarketplaceSession) CollectionRoyalties(arg0 common.Address) (struct {
	Royalty      uint16
	Creator      common.Address
	FeeRecipient common.Address
}, error) {
	return _VolcanoMarketplace.Contract.CollectionRoyalties(&_VolcanoMarketplace.CallOpts, arg0)
}

// CollectionRoyalties is a free data retrieval call binding the contract method 0x57b275b4.
//
// Solidity: function collectionRoyalties(address ) view returns(uint16 royalty, address creator, address feeRecipient)
func (_VolcanoMarketplace *VolcanoMarketplaceCallerSession) CollectionRoyalties(arg0 common.Address) (struct {
	Royalty      uint16
	Creator      common.Address
	FeeRecipient common.Address
}, error) {
	return _VolcanoMarketplace.Contract.CollectionRoyalties(&_VolcanoMarketplace.CallOpts, arg0)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoMarketplace *VolcanoMarketplaceCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoMarketplace.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoMarketplace *VolcanoMarketplaceSession) FeeRecipient() (common.Address, error) {
	return _VolcanoMarketplace.Contract.FeeRecipient(&_VolcanoMarketplace.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoMarketplace *VolcanoMarketplaceCallerSession) FeeRecipient() (common.Address, error) {
	return _VolcanoMarketplace.Contract.FeeRecipient(&_VolcanoMarketplace.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _payToken) view returns(int256)
func (_VolcanoMarketplace *VolcanoMarketplaceCaller) GetPrice(opts *bind.CallOpts, _payToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoMarketplace.contract.Call(opts, &out, "getPrice", _payToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _payToken) view returns(int256)
func (_VolcanoMarketplace *VolcanoMarketplaceSession) GetPrice(_payToken common.Address) (*big.Int, error) {
	return _VolcanoMarketplace.Contract.GetPrice(&_VolcanoMarketplace.CallOpts, _payToken)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _payToken) view returns(int256)
func (_VolcanoMarketplace *VolcanoMarketplaceCallerSession) GetPrice(_payToken common.Address) (*big.Int, error) {
	return _VolcanoMarketplace.Contract.GetPrice(&_VolcanoMarketplace.CallOpts, _payToken)
}

// Listings is a free data retrieval call binding the contract method 0x6bd3a64b.
//
// Solidity: function listings(address , uint256 , address ) view returns(uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_VolcanoMarketplace *VolcanoMarketplaceCaller) Listings(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	StartingTime *big.Int
}, error) {
	var out []interface{}
	err := _VolcanoMarketplace.contract.Call(opts, &out, "listings", arg0, arg1, arg2)

	outstruct := new(struct {
		Quantity     *big.Int
		PayToken     common.Address
		PricePerItem *big.Int
		StartingTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Quantity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PayToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.PricePerItem = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartingTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0x6bd3a64b.
//
// Solidity: function listings(address , uint256 , address ) view returns(uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_VolcanoMarketplace *VolcanoMarketplaceSession) Listings(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	StartingTime *big.Int
}, error) {
	return _VolcanoMarketplace.Contract.Listings(&_VolcanoMarketplace.CallOpts, arg0, arg1, arg2)
}

// Listings is a free data retrieval call binding the contract method 0x6bd3a64b.
//
// Solidity: function listings(address , uint256 , address ) view returns(uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_VolcanoMarketplace *VolcanoMarketplaceCallerSession) Listings(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	StartingTime *big.Int
}, error) {
	return _VolcanoMarketplace.Contract.Listings(&_VolcanoMarketplace.CallOpts, arg0, arg1, arg2)
}

// Minters is a free data retrieval call binding the contract method 0x78746d0d.
//
// Solidity: function minters(address , uint256 ) view returns(address)
func (_VolcanoMarketplace *VolcanoMarketplaceCaller) Minters(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VolcanoMarketplace.contract.Call(opts, &out, "minters", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minters is a free data retrieval call binding the contract method 0x78746d0d.
//
// Solidity: function minters(address , uint256 ) view returns(address)
func (_VolcanoMarketplace *VolcanoMarketplaceSession) Minters(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _VolcanoMarketplace.Contract.Minters(&_VolcanoMarketplace.CallOpts, arg0, arg1)
}

// Minters is a free data retrieval call binding the contract method 0x78746d0d.
//
// Solidity: function minters(address , uint256 ) view returns(address)
func (_VolcanoMarketplace *VolcanoMarketplaceCallerSession) Minters(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _VolcanoMarketplace.Contract.Minters(&_VolcanoMarketplace.CallOpts, arg0, arg1)
}

// Offers is a free data retrieval call binding the contract method 0xd3f494cc.
//
// Solidity: function offers(address , uint256 , address ) view returns(address payToken, uint256 quantity, uint256 pricePerItem, uint256 deadline)
func (_VolcanoMarketplace *VolcanoMarketplaceCaller) Offers(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	PayToken     common.Address
	Quantity     *big.Int
	PricePerItem *big.Int
	Deadline     *big.Int
}, error) {
	var out []interface{}
	err := _VolcanoMarketplace.contract.Call(opts, &out, "offers", arg0, arg1, arg2)

	outstruct := new(struct {
		PayToken     common.Address
		Quantity     *big.Int
		PricePerItem *big.Int
		Deadline     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PayToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Quantity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PricePerItem = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Offers is a free data retrieval call binding the contract method 0xd3f494cc.
//
// Solidity: function offers(address , uint256 , address ) view returns(address payToken, uint256 quantity, uint256 pricePerItem, uint256 deadline)
func (_VolcanoMarketplace *VolcanoMarketplaceSession) Offers(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	PayToken     common.Address
	Quantity     *big.Int
	PricePerItem *big.Int
	Deadline     *big.Int
}, error) {
	return _VolcanoMarketplace.Contract.Offers(&_VolcanoMarketplace.CallOpts, arg0, arg1, arg2)
}

// Offers is a free data retrieval call binding the contract method 0xd3f494cc.
//
// Solidity: function offers(address , uint256 , address ) view returns(address payToken, uint256 quantity, uint256 pricePerItem, uint256 deadline)
func (_VolcanoMarketplace *VolcanoMarketplaceCallerSession) Offers(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	PayToken     common.Address
	Quantity     *big.Int
	PricePerItem *big.Int
	Deadline     *big.Int
}, error) {
	return _VolcanoMarketplace.Contract.Offers(&_VolcanoMarketplace.CallOpts, arg0, arg1, arg2)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoMarketplace *VolcanoMarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoMarketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoMarketplace *VolcanoMarketplaceSession) Owner() (common.Address, error) {
	return _VolcanoMarketplace.Contract.Owner(&_VolcanoMarketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoMarketplace *VolcanoMarketplaceCallerSession) Owner() (common.Address, error) {
	return _VolcanoMarketplace.Contract.Owner(&_VolcanoMarketplace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VolcanoMarketplace *VolcanoMarketplaceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VolcanoMarketplace.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VolcanoMarketplace *VolcanoMarketplaceSession) Paused() (bool, error) {
	return _VolcanoMarketplace.Contract.Paused(&_VolcanoMarketplace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VolcanoMarketplace *VolcanoMarketplaceCallerSession) Paused() (bool, error) {
	return _VolcanoMarketplace.Contract.Paused(&_VolcanoMarketplace.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VolcanoMarketplace *VolcanoMarketplaceCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoMarketplace.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VolcanoMarketplace *VolcanoMarketplaceSession) PlatformFee() (*big.Int, error) {
	return _VolcanoMarketplace.Contract.PlatformFee(&_VolcanoMarketplace.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VolcanoMarketplace *VolcanoMarketplaceCallerSession) PlatformFee() (*big.Int, error) {
	return _VolcanoMarketplace.Contract.PlatformFee(&_VolcanoMarketplace.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VolcanoMarketplace *VolcanoMarketplaceCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VolcanoMarketplace.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VolcanoMarketplace *VolcanoMarketplaceSession) ProxiableUUID() ([32]byte, error) {
	return _VolcanoMarketplace.Contract.ProxiableUUID(&_VolcanoMarketplace.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VolcanoMarketplace *VolcanoMarketplaceCallerSession) ProxiableUUID() ([32]byte, error) {
	return _VolcanoMarketplace.Contract.ProxiableUUID(&_VolcanoMarketplace.CallOpts)
}

// Royalties is a free data retrieval call binding the contract method 0xe1e549c4.
//
// Solidity: function royalties(address , uint256 ) view returns(uint16)
func (_VolcanoMarketplace *VolcanoMarketplaceCaller) Royalties(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (uint16, error) {
	var out []interface{}
	err := _VolcanoMarketplace.contract.Call(opts, &out, "royalties", arg0, arg1)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Royalties is a free data retrieval call binding the contract method 0xe1e549c4.
//
// Solidity: function royalties(address , uint256 ) view returns(uint16)
func (_VolcanoMarketplace *VolcanoMarketplaceSession) Royalties(arg0 common.Address, arg1 *big.Int) (uint16, error) {
	return _VolcanoMarketplace.Contract.Royalties(&_VolcanoMarketplace.CallOpts, arg0, arg1)
}

// Royalties is a free data retrieval call binding the contract method 0xe1e549c4.
//
// Solidity: function royalties(address , uint256 ) view returns(uint16)
func (_VolcanoMarketplace *VolcanoMarketplaceCallerSession) Royalties(arg0 common.Address, arg1 *big.Int) (uint16, error) {
	return _VolcanoMarketplace.Contract.Royalties(&_VolcanoMarketplace.CallOpts, arg0, arg1)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x3bbb2806.
//
// Solidity: function acceptOffer(address _nftAddress, uint256 _tokenId, address _creator) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) AcceptOffer(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "acceptOffer", _nftAddress, _tokenId, _creator)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x3bbb2806.
//
// Solidity: function acceptOffer(address _nftAddress, uint256 _tokenId, address _creator) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) AcceptOffer(_nftAddress common.Address, _tokenId *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.AcceptOffer(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId, _creator)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x3bbb2806.
//
// Solidity: function acceptOffer(address _nftAddress, uint256 _tokenId, address _creator) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) AcceptOffer(_nftAddress common.Address, _tokenId *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.AcceptOffer(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId, _creator)
}

// BuyItem is a paid mutator transaction binding the contract method 0x85f9d345.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _payToken, address _owner) payable returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) BuyItem(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _owner common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "buyItem", _nftAddress, _tokenId, _payToken, _owner)
}

// BuyItem is a paid mutator transaction binding the contract method 0x85f9d345.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _payToken, address _owner) payable returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) BuyItem(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _owner common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.BuyItem(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _owner)
}

// BuyItem is a paid mutator transaction binding the contract method 0x85f9d345.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _payToken, address _owner) payable returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) BuyItem(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _owner common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.BuyItem(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _owner)
}

// CancelListing is a paid mutator transaction binding the contract method 0xb2ddee06.
//
// Solidity: function cancelListing(address _nftAddress, uint256 _tokenId) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) CancelListing(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "cancelListing", _nftAddress, _tokenId)
}

// CancelListing is a paid mutator transaction binding the contract method 0xb2ddee06.
//
// Solidity: function cancelListing(address _nftAddress, uint256 _tokenId) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) CancelListing(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.CancelListing(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId)
}

// CancelListing is a paid mutator transaction binding the contract method 0xb2ddee06.
//
// Solidity: function cancelListing(address _nftAddress, uint256 _tokenId) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) CancelListing(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.CancelListing(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x058a56ac.
//
// Solidity: function cancelOffer(address _nftAddress, uint256 _tokenId) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) CancelOffer(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "cancelOffer", _nftAddress, _tokenId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x058a56ac.
//
// Solidity: function cancelOffer(address _nftAddress, uint256 _tokenId) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) CancelOffer(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.CancelOffer(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x058a56ac.
//
// Solidity: function cancelOffer(address _nftAddress, uint256 _tokenId) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) CancelOffer(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.CancelOffer(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x42cbb4b8.
//
// Solidity: function createOffer(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _quantity, uint256 _pricePerItem, uint256 _deadline) payable returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) CreateOffer(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _quantity *big.Int, _pricePerItem *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "createOffer", _nftAddress, _tokenId, _payToken, _quantity, _pricePerItem, _deadline)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x42cbb4b8.
//
// Solidity: function createOffer(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _quantity, uint256 _pricePerItem, uint256 _deadline) payable returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) CreateOffer(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _quantity *big.Int, _pricePerItem *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.CreateOffer(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _quantity, _pricePerItem, _deadline)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x42cbb4b8.
//
// Solidity: function createOffer(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _quantity, uint256 _pricePerItem, uint256 _deadline) payable returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) CreateOffer(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _quantity *big.Int, _pricePerItem *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.CreateOffer(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _quantity, _pricePerItem, _deadline)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _feeRecipient, uint256 _platformFee) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) Initialize(opts *bind.TransactOpts, _feeRecipient common.Address, _platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "initialize", _feeRecipient, _platformFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _feeRecipient, uint256 _platformFee) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) Initialize(_feeRecipient common.Address, _platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.Initialize(&_VolcanoMarketplace.TransactOpts, _feeRecipient, _platformFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _feeRecipient, uint256 _platformFee) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) Initialize(_feeRecipient common.Address, _platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.Initialize(&_VolcanoMarketplace.TransactOpts, _feeRecipient, _platformFee)
}

// ListItem is a paid mutator transaction binding the contract method 0x3fc1cc26.
//
// Solidity: function listItem(address _nftAddress, uint256 _tokenId, uint256 _quantity, address _payToken, uint256 _pricePerItem, uint256 _startingTime) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) ListItem(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int, _payToken common.Address, _pricePerItem *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "listItem", _nftAddress, _tokenId, _quantity, _payToken, _pricePerItem, _startingTime)
}

// ListItem is a paid mutator transaction binding the contract method 0x3fc1cc26.
//
// Solidity: function listItem(address _nftAddress, uint256 _tokenId, uint256 _quantity, address _payToken, uint256 _pricePerItem, uint256 _startingTime) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) ListItem(_nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int, _payToken common.Address, _pricePerItem *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.ListItem(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId, _quantity, _payToken, _pricePerItem, _startingTime)
}

// ListItem is a paid mutator transaction binding the contract method 0x3fc1cc26.
//
// Solidity: function listItem(address _nftAddress, uint256 _tokenId, uint256 _quantity, address _payToken, uint256 _pricePerItem, uint256 _startingTime) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) ListItem(_nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int, _payToken common.Address, _pricePerItem *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.ListItem(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId, _quantity, _payToken, _pricePerItem, _startingTime)
}

// RegisterCollectionRoyalty is a paid mutator transaction binding the contract method 0xc1661a9a.
//
// Solidity: function registerCollectionRoyalty(address _nftAddress, address _creator, uint16 _royalty, address _feeRecipient) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) RegisterCollectionRoyalty(opts *bind.TransactOpts, _nftAddress common.Address, _creator common.Address, _royalty uint16, _feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "registerCollectionRoyalty", _nftAddress, _creator, _royalty, _feeRecipient)
}

// RegisterCollectionRoyalty is a paid mutator transaction binding the contract method 0xc1661a9a.
//
// Solidity: function registerCollectionRoyalty(address _nftAddress, address _creator, uint16 _royalty, address _feeRecipient) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) RegisterCollectionRoyalty(_nftAddress common.Address, _creator common.Address, _royalty uint16, _feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.RegisterCollectionRoyalty(&_VolcanoMarketplace.TransactOpts, _nftAddress, _creator, _royalty, _feeRecipient)
}

// RegisterCollectionRoyalty is a paid mutator transaction binding the contract method 0xc1661a9a.
//
// Solidity: function registerCollectionRoyalty(address _nftAddress, address _creator, uint16 _royalty, address _feeRecipient) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) RegisterCollectionRoyalty(_nftAddress common.Address, _creator common.Address, _royalty uint16, _feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.RegisterCollectionRoyalty(&_VolcanoMarketplace.TransactOpts, _nftAddress, _creator, _royalty, _feeRecipient)
}

// RegisterRoyalty is a paid mutator transaction binding the contract method 0xf3880b6e.
//
// Solidity: function registerRoyalty(address _nftAddress, uint256 _tokenId, uint16 _royalty) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) RegisterRoyalty(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _royalty uint16) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "registerRoyalty", _nftAddress, _tokenId, _royalty)
}

// RegisterRoyalty is a paid mutator transaction binding the contract method 0xf3880b6e.
//
// Solidity: function registerRoyalty(address _nftAddress, uint256 _tokenId, uint16 _royalty) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) RegisterRoyalty(_nftAddress common.Address, _tokenId *big.Int, _royalty uint16) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.RegisterRoyalty(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId, _royalty)
}

// RegisterRoyalty is a paid mutator transaction binding the contract method 0xf3880b6e.
//
// Solidity: function registerRoyalty(address _nftAddress, uint256 _tokenId, uint16 _royalty) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) RegisterRoyalty(_nftAddress common.Address, _tokenId *big.Int, _royalty uint16) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.RegisterRoyalty(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId, _royalty)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.RenounceOwnership(&_VolcanoMarketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.RenounceOwnership(&_VolcanoMarketplace.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.TransferOwnership(&_VolcanoMarketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.TransferOwnership(&_VolcanoMarketplace.TransactOpts, newOwner)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) UpdateAddressRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "updateAddressRegistry", _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.UpdateAddressRegistry(&_VolcanoMarketplace.TransactOpts, _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.UpdateAddressRegistry(&_VolcanoMarketplace.TransactOpts, _registry)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xe984f2eb.
//
// Solidity: function updateListing(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _newPrice) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) UpdateListing(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "updateListing", _nftAddress, _tokenId, _payToken, _newPrice)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xe984f2eb.
//
// Solidity: function updateListing(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _newPrice) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) UpdateListing(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.UpdateListing(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _newPrice)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xe984f2eb.
//
// Solidity: function updateListing(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _newPrice) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) UpdateListing(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.UpdateListing(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _newPrice)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xafb06952.
//
// Solidity: function updatePlatformFee(uint16 _platformFee) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) UpdatePlatformFee(opts *bind.TransactOpts, _platformFee uint16) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "updatePlatformFee", _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xafb06952.
//
// Solidity: function updatePlatformFee(uint16 _platformFee) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) UpdatePlatformFee(_platformFee uint16) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.UpdatePlatformFee(&_VolcanoMarketplace.TransactOpts, _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xafb06952.
//
// Solidity: function updatePlatformFee(uint16 _platformFee) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) UpdatePlatformFee(_platformFee uint16) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.UpdatePlatformFee(&_VolcanoMarketplace.TransactOpts, _platformFee)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) UpdatePlatformFeeRecipient(opts *bind.TransactOpts, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "updatePlatformFeeRecipient", _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.UpdatePlatformFeeRecipient(&_VolcanoMarketplace.TransactOpts, _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.UpdatePlatformFeeRecipient(&_VolcanoMarketplace.TransactOpts, _platformFeeRecipient)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.UpgradeTo(&_VolcanoMarketplace.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.UpgradeTo(&_VolcanoMarketplace.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.UpgradeToAndCall(&_VolcanoMarketplace.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.UpgradeToAndCall(&_VolcanoMarketplace.TransactOpts, newImplementation, data)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0xe940ebeb.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, address _seller, address _buyer) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) ValidateItemSold(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _seller common.Address, _buyer common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.Transact(opts, "validateItemSold", _nftAddress, _tokenId, _seller, _buyer)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0xe940ebeb.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, address _seller, address _buyer) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) ValidateItemSold(_nftAddress common.Address, _tokenId *big.Int, _seller common.Address, _buyer common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.ValidateItemSold(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId, _seller, _buyer)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0xe940ebeb.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, address _seller, address _buyer) returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) ValidateItemSold(_nftAddress common.Address, _tokenId *big.Int, _seller common.Address, _buyer common.Address) (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.ValidateItemSold(&_VolcanoMarketplace.TransactOpts, _nftAddress, _tokenId, _seller, _buyer)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoMarketplace.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VolcanoMarketplace *VolcanoMarketplaceSession) Receive() (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.Receive(&_VolcanoMarketplace.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VolcanoMarketplace *VolcanoMarketplaceTransactorSession) Receive() (*types.Transaction, error) {
	return _VolcanoMarketplace.Contract.Receive(&_VolcanoMarketplace.TransactOpts)
}

// VolcanoMarketplaceAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceAdminChangedIterator struct {
	Event *VolcanoMarketplaceAdminChanged // Event containing the contract specifics and raw log

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
func (it *VolcanoMarketplaceAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoMarketplaceAdminChanged)
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
		it.Event = new(VolcanoMarketplaceAdminChanged)
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
func (it *VolcanoMarketplaceAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoMarketplaceAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoMarketplaceAdminChanged represents a AdminChanged event raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*VolcanoMarketplaceAdminChangedIterator, error) {

	logs, sub, err := _VolcanoMarketplace.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceAdminChangedIterator{contract: _VolcanoMarketplace.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *VolcanoMarketplaceAdminChanged) (event.Subscription, error) {

	logs, sub, err := _VolcanoMarketplace.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoMarketplaceAdminChanged)
				if err := _VolcanoMarketplace.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) ParseAdminChanged(log types.Log) (*VolcanoMarketplaceAdminChanged, error) {
	event := new(VolcanoMarketplaceAdminChanged)
	if err := _VolcanoMarketplace.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoMarketplaceBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceBeaconUpgradedIterator struct {
	Event *VolcanoMarketplaceBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *VolcanoMarketplaceBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoMarketplaceBeaconUpgraded)
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
		it.Event = new(VolcanoMarketplaceBeaconUpgraded)
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
func (it *VolcanoMarketplaceBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoMarketplaceBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoMarketplaceBeaconUpgraded represents a BeaconUpgraded event raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*VolcanoMarketplaceBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceBeaconUpgradedIterator{contract: _VolcanoMarketplace.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *VolcanoMarketplaceBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoMarketplaceBeaconUpgraded)
				if err := _VolcanoMarketplace.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) ParseBeaconUpgraded(log types.Log) (*VolcanoMarketplaceBeaconUpgraded, error) {
	event := new(VolcanoMarketplaceBeaconUpgraded)
	if err := _VolcanoMarketplace.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoMarketplaceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceInitializedIterator struct {
	Event *VolcanoMarketplaceInitialized // Event containing the contract specifics and raw log

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
func (it *VolcanoMarketplaceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoMarketplaceInitialized)
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
		it.Event = new(VolcanoMarketplaceInitialized)
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
func (it *VolcanoMarketplaceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoMarketplaceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoMarketplaceInitialized represents a Initialized event raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) FilterInitialized(opts *bind.FilterOpts) (*VolcanoMarketplaceInitializedIterator, error) {

	logs, sub, err := _VolcanoMarketplace.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceInitializedIterator{contract: _VolcanoMarketplace.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VolcanoMarketplaceInitialized) (event.Subscription, error) {

	logs, sub, err := _VolcanoMarketplace.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoMarketplaceInitialized)
				if err := _VolcanoMarketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) ParseInitialized(log types.Log) (*VolcanoMarketplaceInitialized, error) {
	event := new(VolcanoMarketplaceInitialized)
	if err := _VolcanoMarketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoMarketplaceItemCanceledIterator is returned from FilterItemCanceled and is used to iterate over the raw logs and unpacked data for ItemCanceled events raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceItemCanceledIterator struct {
	Event *VolcanoMarketplaceItemCanceled // Event containing the contract specifics and raw log

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
func (it *VolcanoMarketplaceItemCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoMarketplaceItemCanceled)
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
		it.Event = new(VolcanoMarketplaceItemCanceled)
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
func (it *VolcanoMarketplaceItemCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoMarketplaceItemCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoMarketplaceItemCanceled represents a ItemCanceled event raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceItemCanceled struct {
	Owner   common.Address
	Nft     common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterItemCanceled is a free log retrieval operation binding the contract event 0x9ba1a3cb55ce8d63d072a886f94d2a744f50cddf82128e897d0661f5ec623158.
//
// Solidity: event ItemCanceled(address indexed owner, address indexed nft, uint256 tokenId)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) FilterItemCanceled(opts *bind.FilterOpts, owner []common.Address, nft []common.Address) (*VolcanoMarketplaceItemCanceledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.FilterLogs(opts, "ItemCanceled", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceItemCanceledIterator{contract: _VolcanoMarketplace.contract, event: "ItemCanceled", logs: logs, sub: sub}, nil
}

// WatchItemCanceled is a free log subscription operation binding the contract event 0x9ba1a3cb55ce8d63d072a886f94d2a744f50cddf82128e897d0661f5ec623158.
//
// Solidity: event ItemCanceled(address indexed owner, address indexed nft, uint256 tokenId)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) WatchItemCanceled(opts *bind.WatchOpts, sink chan<- *VolcanoMarketplaceItemCanceled, owner []common.Address, nft []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.WatchLogs(opts, "ItemCanceled", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoMarketplaceItemCanceled)
				if err := _VolcanoMarketplace.contract.UnpackLog(event, "ItemCanceled", log); err != nil {
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

// ParseItemCanceled is a log parse operation binding the contract event 0x9ba1a3cb55ce8d63d072a886f94d2a744f50cddf82128e897d0661f5ec623158.
//
// Solidity: event ItemCanceled(address indexed owner, address indexed nft, uint256 tokenId)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) ParseItemCanceled(log types.Log) (*VolcanoMarketplaceItemCanceled, error) {
	event := new(VolcanoMarketplaceItemCanceled)
	if err := _VolcanoMarketplace.contract.UnpackLog(event, "ItemCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoMarketplaceItemListedIterator is returned from FilterItemListed and is used to iterate over the raw logs and unpacked data for ItemListed events raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceItemListedIterator struct {
	Event *VolcanoMarketplaceItemListed // Event containing the contract specifics and raw log

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
func (it *VolcanoMarketplaceItemListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoMarketplaceItemListed)
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
		it.Event = new(VolcanoMarketplaceItemListed)
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
func (it *VolcanoMarketplaceItemListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoMarketplaceItemListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoMarketplaceItemListed represents a ItemListed event raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceItemListed struct {
	Owner        common.Address
	Nft          common.Address
	TokenId      *big.Int
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	StartingTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterItemListed is a free log retrieval operation binding the contract event 0xa0294f02f8ad82fe4744717b0f953a105547196cd3c67056200c1a4ae3aa2629.
//
// Solidity: event ItemListed(address indexed owner, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) FilterItemListed(opts *bind.FilterOpts, owner []common.Address, nft []common.Address) (*VolcanoMarketplaceItemListedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.FilterLogs(opts, "ItemListed", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceItemListedIterator{contract: _VolcanoMarketplace.contract, event: "ItemListed", logs: logs, sub: sub}, nil
}

// WatchItemListed is a free log subscription operation binding the contract event 0xa0294f02f8ad82fe4744717b0f953a105547196cd3c67056200c1a4ae3aa2629.
//
// Solidity: event ItemListed(address indexed owner, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) WatchItemListed(opts *bind.WatchOpts, sink chan<- *VolcanoMarketplaceItemListed, owner []common.Address, nft []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.WatchLogs(opts, "ItemListed", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoMarketplaceItemListed)
				if err := _VolcanoMarketplace.contract.UnpackLog(event, "ItemListed", log); err != nil {
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

// ParseItemListed is a log parse operation binding the contract event 0xa0294f02f8ad82fe4744717b0f953a105547196cd3c67056200c1a4ae3aa2629.
//
// Solidity: event ItemListed(address indexed owner, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) ParseItemListed(log types.Log) (*VolcanoMarketplaceItemListed, error) {
	event := new(VolcanoMarketplaceItemListed)
	if err := _VolcanoMarketplace.contract.UnpackLog(event, "ItemListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoMarketplaceItemSoldIterator is returned from FilterItemSold and is used to iterate over the raw logs and unpacked data for ItemSold events raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceItemSoldIterator struct {
	Event *VolcanoMarketplaceItemSold // Event containing the contract specifics and raw log

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
func (it *VolcanoMarketplaceItemSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoMarketplaceItemSold)
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
		it.Event = new(VolcanoMarketplaceItemSold)
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
func (it *VolcanoMarketplaceItemSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoMarketplaceItemSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoMarketplaceItemSold represents a ItemSold event raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceItemSold struct {
	Seller        common.Address
	Buyer         common.Address
	Nft           common.Address
	TokenId       *big.Int
	Quantity      *big.Int
	PayToken      common.Address
	PayTokenPrice *big.Int
	PricePerItem  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterItemSold is a free log retrieval operation binding the contract event 0x949d1413baca5c0e4ab96b0198d536cac8cdcc17cb909b9ea24594f42ed9fa0d.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, int256 payTokenPrice, uint256 pricePerItem)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) FilterItemSold(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, nft []common.Address) (*VolcanoMarketplaceItemSoldIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.FilterLogs(opts, "ItemSold", sellerRule, buyerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceItemSoldIterator{contract: _VolcanoMarketplace.contract, event: "ItemSold", logs: logs, sub: sub}, nil
}

// WatchItemSold is a free log subscription operation binding the contract event 0x949d1413baca5c0e4ab96b0198d536cac8cdcc17cb909b9ea24594f42ed9fa0d.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, int256 payTokenPrice, uint256 pricePerItem)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) WatchItemSold(opts *bind.WatchOpts, sink chan<- *VolcanoMarketplaceItemSold, seller []common.Address, buyer []common.Address, nft []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.WatchLogs(opts, "ItemSold", sellerRule, buyerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoMarketplaceItemSold)
				if err := _VolcanoMarketplace.contract.UnpackLog(event, "ItemSold", log); err != nil {
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

// ParseItemSold is a log parse operation binding the contract event 0x949d1413baca5c0e4ab96b0198d536cac8cdcc17cb909b9ea24594f42ed9fa0d.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, int256 payTokenPrice, uint256 pricePerItem)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) ParseItemSold(log types.Log) (*VolcanoMarketplaceItemSold, error) {
	event := new(VolcanoMarketplaceItemSold)
	if err := _VolcanoMarketplace.contract.UnpackLog(event, "ItemSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoMarketplaceItemUpdatedIterator is returned from FilterItemUpdated and is used to iterate over the raw logs and unpacked data for ItemUpdated events raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceItemUpdatedIterator struct {
	Event *VolcanoMarketplaceItemUpdated // Event containing the contract specifics and raw log

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
func (it *VolcanoMarketplaceItemUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoMarketplaceItemUpdated)
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
		it.Event = new(VolcanoMarketplaceItemUpdated)
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
func (it *VolcanoMarketplaceItemUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoMarketplaceItemUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoMarketplaceItemUpdated represents a ItemUpdated event raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceItemUpdated struct {
	Owner    common.Address
	Nft      common.Address
	TokenId  *big.Int
	PayToken common.Address
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterItemUpdated is a free log retrieval operation binding the contract event 0x60a11f1619b1716bc2857bf610d4bc631336e14d197025fd5875c1aca1ac7cbd.
//
// Solidity: event ItemUpdated(address indexed owner, address indexed nft, uint256 tokenId, address payToken, uint256 newPrice)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) FilterItemUpdated(opts *bind.FilterOpts, owner []common.Address, nft []common.Address) (*VolcanoMarketplaceItemUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.FilterLogs(opts, "ItemUpdated", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceItemUpdatedIterator{contract: _VolcanoMarketplace.contract, event: "ItemUpdated", logs: logs, sub: sub}, nil
}

// WatchItemUpdated is a free log subscription operation binding the contract event 0x60a11f1619b1716bc2857bf610d4bc631336e14d197025fd5875c1aca1ac7cbd.
//
// Solidity: event ItemUpdated(address indexed owner, address indexed nft, uint256 tokenId, address payToken, uint256 newPrice)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) WatchItemUpdated(opts *bind.WatchOpts, sink chan<- *VolcanoMarketplaceItemUpdated, owner []common.Address, nft []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.WatchLogs(opts, "ItemUpdated", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoMarketplaceItemUpdated)
				if err := _VolcanoMarketplace.contract.UnpackLog(event, "ItemUpdated", log); err != nil {
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

// ParseItemUpdated is a log parse operation binding the contract event 0x60a11f1619b1716bc2857bf610d4bc631336e14d197025fd5875c1aca1ac7cbd.
//
// Solidity: event ItemUpdated(address indexed owner, address indexed nft, uint256 tokenId, address payToken, uint256 newPrice)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) ParseItemUpdated(log types.Log) (*VolcanoMarketplaceItemUpdated, error) {
	event := new(VolcanoMarketplaceItemUpdated)
	if err := _VolcanoMarketplace.contract.UnpackLog(event, "ItemUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoMarketplaceOfferCanceledIterator is returned from FilterOfferCanceled and is used to iterate over the raw logs and unpacked data for OfferCanceled events raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceOfferCanceledIterator struct {
	Event *VolcanoMarketplaceOfferCanceled // Event containing the contract specifics and raw log

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
func (it *VolcanoMarketplaceOfferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoMarketplaceOfferCanceled)
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
		it.Event = new(VolcanoMarketplaceOfferCanceled)
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
func (it *VolcanoMarketplaceOfferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoMarketplaceOfferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoMarketplaceOfferCanceled represents a OfferCanceled event raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceOfferCanceled struct {
	Creator common.Address
	Nft     common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferCanceled is a free log retrieval operation binding the contract event 0xc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a7.
//
// Solidity: event OfferCanceled(address indexed creator, address indexed nft, uint256 tokenId)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) FilterOfferCanceled(opts *bind.FilterOpts, creator []common.Address, nft []common.Address) (*VolcanoMarketplaceOfferCanceledIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.FilterLogs(opts, "OfferCanceled", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceOfferCanceledIterator{contract: _VolcanoMarketplace.contract, event: "OfferCanceled", logs: logs, sub: sub}, nil
}

// WatchOfferCanceled is a free log subscription operation binding the contract event 0xc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a7.
//
// Solidity: event OfferCanceled(address indexed creator, address indexed nft, uint256 tokenId)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) WatchOfferCanceled(opts *bind.WatchOpts, sink chan<- *VolcanoMarketplaceOfferCanceled, creator []common.Address, nft []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.WatchLogs(opts, "OfferCanceled", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoMarketplaceOfferCanceled)
				if err := _VolcanoMarketplace.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
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

// ParseOfferCanceled is a log parse operation binding the contract event 0xc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a7.
//
// Solidity: event OfferCanceled(address indexed creator, address indexed nft, uint256 tokenId)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) ParseOfferCanceled(log types.Log) (*VolcanoMarketplaceOfferCanceled, error) {
	event := new(VolcanoMarketplaceOfferCanceled)
	if err := _VolcanoMarketplace.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoMarketplaceOfferCreatedIterator is returned from FilterOfferCreated and is used to iterate over the raw logs and unpacked data for OfferCreated events raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceOfferCreatedIterator struct {
	Event *VolcanoMarketplaceOfferCreated // Event containing the contract specifics and raw log

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
func (it *VolcanoMarketplaceOfferCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoMarketplaceOfferCreated)
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
		it.Event = new(VolcanoMarketplaceOfferCreated)
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
func (it *VolcanoMarketplaceOfferCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoMarketplaceOfferCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoMarketplaceOfferCreated represents a OfferCreated event raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceOfferCreated struct {
	Creator      common.Address
	Nft          common.Address
	TokenId      *big.Int
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	Deadline     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOfferCreated is a free log retrieval operation binding the contract event 0x89f255157c655b5155655107b77c620998e5ad4e7485d749e4e6d7ddb63e70f6.
//
// Solidity: event OfferCreated(address indexed creator, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 deadline)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) FilterOfferCreated(opts *bind.FilterOpts, creator []common.Address, nft []common.Address) (*VolcanoMarketplaceOfferCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.FilterLogs(opts, "OfferCreated", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceOfferCreatedIterator{contract: _VolcanoMarketplace.contract, event: "OfferCreated", logs: logs, sub: sub}, nil
}

// WatchOfferCreated is a free log subscription operation binding the contract event 0x89f255157c655b5155655107b77c620998e5ad4e7485d749e4e6d7ddb63e70f6.
//
// Solidity: event OfferCreated(address indexed creator, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 deadline)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) WatchOfferCreated(opts *bind.WatchOpts, sink chan<- *VolcanoMarketplaceOfferCreated, creator []common.Address, nft []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.WatchLogs(opts, "OfferCreated", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoMarketplaceOfferCreated)
				if err := _VolcanoMarketplace.contract.UnpackLog(event, "OfferCreated", log); err != nil {
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

// ParseOfferCreated is a log parse operation binding the contract event 0x89f255157c655b5155655107b77c620998e5ad4e7485d749e4e6d7ddb63e70f6.
//
// Solidity: event OfferCreated(address indexed creator, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 deadline)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) ParseOfferCreated(log types.Log) (*VolcanoMarketplaceOfferCreated, error) {
	event := new(VolcanoMarketplaceOfferCreated)
	if err := _VolcanoMarketplace.contract.UnpackLog(event, "OfferCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoMarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceOwnershipTransferredIterator struct {
	Event *VolcanoMarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VolcanoMarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoMarketplaceOwnershipTransferred)
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
		it.Event = new(VolcanoMarketplaceOwnershipTransferred)
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
func (it *VolcanoMarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoMarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoMarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VolcanoMarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceOwnershipTransferredIterator{contract: _VolcanoMarketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VolcanoMarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoMarketplaceOwnershipTransferred)
				if err := _VolcanoMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*VolcanoMarketplaceOwnershipTransferred, error) {
	event := new(VolcanoMarketplaceOwnershipTransferred)
	if err := _VolcanoMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoMarketplacePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the VolcanoMarketplace contract.
type VolcanoMarketplacePausedIterator struct {
	Event *VolcanoMarketplacePaused // Event containing the contract specifics and raw log

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
func (it *VolcanoMarketplacePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoMarketplacePaused)
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
		it.Event = new(VolcanoMarketplacePaused)
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
func (it *VolcanoMarketplacePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoMarketplacePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoMarketplacePaused represents a Paused event raised by the VolcanoMarketplace contract.
type VolcanoMarketplacePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) FilterPaused(opts *bind.FilterOpts) (*VolcanoMarketplacePausedIterator, error) {

	logs, sub, err := _VolcanoMarketplace.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplacePausedIterator{contract: _VolcanoMarketplace.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VolcanoMarketplacePaused) (event.Subscription, error) {

	logs, sub, err := _VolcanoMarketplace.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoMarketplacePaused)
				if err := _VolcanoMarketplace.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) ParsePaused(log types.Log) (*VolcanoMarketplacePaused, error) {
	event := new(VolcanoMarketplacePaused)
	if err := _VolcanoMarketplace.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoMarketplaceUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceUnpausedIterator struct {
	Event *VolcanoMarketplaceUnpaused // Event containing the contract specifics and raw log

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
func (it *VolcanoMarketplaceUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoMarketplaceUnpaused)
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
		it.Event = new(VolcanoMarketplaceUnpaused)
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
func (it *VolcanoMarketplaceUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoMarketplaceUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoMarketplaceUnpaused represents a Unpaused event raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VolcanoMarketplaceUnpausedIterator, error) {

	logs, sub, err := _VolcanoMarketplace.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceUnpausedIterator{contract: _VolcanoMarketplace.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VolcanoMarketplaceUnpaused) (event.Subscription, error) {

	logs, sub, err := _VolcanoMarketplace.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoMarketplaceUnpaused)
				if err := _VolcanoMarketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) ParseUnpaused(log types.Log) (*VolcanoMarketplaceUnpaused, error) {
	event := new(VolcanoMarketplaceUnpaused)
	if err := _VolcanoMarketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoMarketplaceUpdatePlatformFeeIterator is returned from FilterUpdatePlatformFee and is used to iterate over the raw logs and unpacked data for UpdatePlatformFee events raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceUpdatePlatformFeeIterator struct {
	Event *VolcanoMarketplaceUpdatePlatformFee // Event containing the contract specifics and raw log

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
func (it *VolcanoMarketplaceUpdatePlatformFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoMarketplaceUpdatePlatformFee)
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
		it.Event = new(VolcanoMarketplaceUpdatePlatformFee)
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
func (it *VolcanoMarketplaceUpdatePlatformFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoMarketplaceUpdatePlatformFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoMarketplaceUpdatePlatformFee represents a UpdatePlatformFee event raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceUpdatePlatformFee struct {
	PlatformFee uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFee is a free log retrieval operation binding the contract event 0xcabead39256298cdf7d1228eee68b47bc54a2192283d8d93574da08d2d7ac303.
//
// Solidity: event UpdatePlatformFee(uint16 platformFee)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) FilterUpdatePlatformFee(opts *bind.FilterOpts) (*VolcanoMarketplaceUpdatePlatformFeeIterator, error) {

	logs, sub, err := _VolcanoMarketplace.contract.FilterLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceUpdatePlatformFeeIterator{contract: _VolcanoMarketplace.contract, event: "UpdatePlatformFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFee is a free log subscription operation binding the contract event 0xcabead39256298cdf7d1228eee68b47bc54a2192283d8d93574da08d2d7ac303.
//
// Solidity: event UpdatePlatformFee(uint16 platformFee)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) WatchUpdatePlatformFee(opts *bind.WatchOpts, sink chan<- *VolcanoMarketplaceUpdatePlatformFee) (event.Subscription, error) {

	logs, sub, err := _VolcanoMarketplace.contract.WatchLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoMarketplaceUpdatePlatformFee)
				if err := _VolcanoMarketplace.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
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

// ParseUpdatePlatformFee is a log parse operation binding the contract event 0xcabead39256298cdf7d1228eee68b47bc54a2192283d8d93574da08d2d7ac303.
//
// Solidity: event UpdatePlatformFee(uint16 platformFee)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) ParseUpdatePlatformFee(log types.Log) (*VolcanoMarketplaceUpdatePlatformFee, error) {
	event := new(VolcanoMarketplaceUpdatePlatformFee)
	if err := _VolcanoMarketplace.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoMarketplaceUpdatePlatformFeeRecipientIterator is returned from FilterUpdatePlatformFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdatePlatformFeeRecipient events raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceUpdatePlatformFeeRecipientIterator struct {
	Event *VolcanoMarketplaceUpdatePlatformFeeRecipient // Event containing the contract specifics and raw log

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
func (it *VolcanoMarketplaceUpdatePlatformFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoMarketplaceUpdatePlatformFeeRecipient)
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
		it.Event = new(VolcanoMarketplaceUpdatePlatformFeeRecipient)
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
func (it *VolcanoMarketplaceUpdatePlatformFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoMarketplaceUpdatePlatformFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoMarketplaceUpdatePlatformFeeRecipient represents a UpdatePlatformFeeRecipient event raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceUpdatePlatformFeeRecipient struct {
	PlatformFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFeeRecipient is a free log retrieval operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) FilterUpdatePlatformFeeRecipient(opts *bind.FilterOpts) (*VolcanoMarketplaceUpdatePlatformFeeRecipientIterator, error) {

	logs, sub, err := _VolcanoMarketplace.contract.FilterLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceUpdatePlatformFeeRecipientIterator{contract: _VolcanoMarketplace.contract, event: "UpdatePlatformFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFeeRecipient is a free log subscription operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) WatchUpdatePlatformFeeRecipient(opts *bind.WatchOpts, sink chan<- *VolcanoMarketplaceUpdatePlatformFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _VolcanoMarketplace.contract.WatchLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoMarketplaceUpdatePlatformFeeRecipient)
				if err := _VolcanoMarketplace.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
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
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) ParseUpdatePlatformFeeRecipient(log types.Log) (*VolcanoMarketplaceUpdatePlatformFeeRecipient, error) {
	event := new(VolcanoMarketplaceUpdatePlatformFeeRecipient)
	if err := _VolcanoMarketplace.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoMarketplaceUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceUpgradedIterator struct {
	Event *VolcanoMarketplaceUpgraded // Event containing the contract specifics and raw log

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
func (it *VolcanoMarketplaceUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoMarketplaceUpgraded)
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
		it.Event = new(VolcanoMarketplaceUpgraded)
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
func (it *VolcanoMarketplaceUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoMarketplaceUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoMarketplaceUpgraded represents a Upgraded event raised by the VolcanoMarketplace contract.
type VolcanoMarketplaceUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VolcanoMarketplaceUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoMarketplaceUpgradedIterator{contract: _VolcanoMarketplace.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VolcanoMarketplaceUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VolcanoMarketplace.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoMarketplaceUpgraded)
				if err := _VolcanoMarketplace.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_VolcanoMarketplace *VolcanoMarketplaceFilterer) ParseUpgraded(log types.Log) (*VolcanoMarketplaceUpgraded, error) {
	event := new(VolcanoMarketplaceUpgraded)
	if err := _VolcanoMarketplace.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
