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

// VolcanoAuctionMetaData contains all meta data concerning the VolcanoAuction contract.
var VolcanoAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"payTokenPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winningBid\",\"type\":\"uint256\"}],\"name\":\"AuctionResulted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"PauseToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"UpdateAuctionEndTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"}],\"name\":\"UpdateAuctionReservePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"UpdateAuctionStartTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidWithdrawalLockTime\",\"type\":\"uint256\"}],\"name\":\"UpdateBidWithdrawalLockTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBidIncrement\",\"type\":\"uint256\"}],\"name\":\"UpdateMinBidIncrement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"name\":\"UpdatePlatformFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"UpdatePlatformFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"VolcanoAuctionContractDeployed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressRegistry\",\"outputs\":[{\"internalType\":\"contractIVolcanoAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"resulted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidWithdrawalLockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"minBidReserve\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_endTimestamp\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_resulted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"minBid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getHighestBidder\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lastBidTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestBids\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBidTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBidIncrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bidAmount\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFeeRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"reclaimERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"resultAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleIsPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"updateAddressRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTimestamp\",\"type\":\"uint256\"}],\"name\":\"updateAuctionEndTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"}],\"name\":\"updateAuctionReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"}],\"name\":\"updateAuctionStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bidWithdrawalLockTime\",\"type\":\"uint256\"}],\"name\":\"updateBidWithdrawalLockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBidIncrement\",\"type\":\"uint256\"}],\"name\":\"updateMinBidIncrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"updatePlatformFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VolcanoAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use VolcanoAuctionMetaData.ABI instead.
var VolcanoAuctionABI = VolcanoAuctionMetaData.ABI

// VolcanoAuction is an auto generated Go binding around an Ethereum contract.
type VolcanoAuction struct {
	VolcanoAuctionCaller     // Read-only binding to the contract
	VolcanoAuctionTransactor // Write-only binding to the contract
	VolcanoAuctionFilterer   // Log filterer for contract events
}

// VolcanoAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type VolcanoAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VolcanoAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolcanoAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolcanoAuctionSession struct {
	Contract     *VolcanoAuction   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VolcanoAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolcanoAuctionCallerSession struct {
	Contract *VolcanoAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VolcanoAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolcanoAuctionTransactorSession struct {
	Contract     *VolcanoAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VolcanoAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type VolcanoAuctionRaw struct {
	Contract *VolcanoAuction // Generic contract binding to access the raw methods on
}

// VolcanoAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolcanoAuctionCallerRaw struct {
	Contract *VolcanoAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// VolcanoAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolcanoAuctionTransactorRaw struct {
	Contract *VolcanoAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolcanoAuction creates a new instance of VolcanoAuction, bound to a specific deployed contract.
func NewVolcanoAuction(address common.Address, backend bind.ContractBackend) (*VolcanoAuction, error) {
	contract, err := bindVolcanoAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuction{VolcanoAuctionCaller: VolcanoAuctionCaller{contract: contract}, VolcanoAuctionTransactor: VolcanoAuctionTransactor{contract: contract}, VolcanoAuctionFilterer: VolcanoAuctionFilterer{contract: contract}}, nil
}

// NewVolcanoAuctionCaller creates a new read-only instance of VolcanoAuction, bound to a specific deployed contract.
func NewVolcanoAuctionCaller(address common.Address, caller bind.ContractCaller) (*VolcanoAuctionCaller, error) {
	contract, err := bindVolcanoAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionCaller{contract: contract}, nil
}

// NewVolcanoAuctionTransactor creates a new write-only instance of VolcanoAuction, bound to a specific deployed contract.
func NewVolcanoAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*VolcanoAuctionTransactor, error) {
	contract, err := bindVolcanoAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionTransactor{contract: contract}, nil
}

// NewVolcanoAuctionFilterer creates a new log filterer instance of VolcanoAuction, bound to a specific deployed contract.
func NewVolcanoAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*VolcanoAuctionFilterer, error) {
	contract, err := bindVolcanoAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionFilterer{contract: contract}, nil
}

// bindVolcanoAuction binds a generic wrapper to an already deployed contract.
func bindVolcanoAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VolcanoAuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoAuction *VolcanoAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoAuction.Contract.VolcanoAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoAuction *VolcanoAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.VolcanoAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoAuction *VolcanoAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.VolcanoAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoAuction *VolcanoAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoAuction *VolcanoAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoAuction *VolcanoAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.contract.Transact(opts, method, params...)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_VolcanoAuction *VolcanoAuctionCaller) AddressRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoAuction.contract.Call(opts, &out, "addressRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_VolcanoAuction *VolcanoAuctionSession) AddressRegistry() (common.Address, error) {
	return _VolcanoAuction.Contract.AddressRegistry(&_VolcanoAuction.CallOpts)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_VolcanoAuction *VolcanoAuctionCallerSession) AddressRegistry() (common.Address, error) {
	return _VolcanoAuction.Contract.AddressRegistry(&_VolcanoAuction.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address owner, address payToken, uint256 minBid, uint256 reservePrice, uint256 startTime, uint256 endTime, bool resulted)
func (_VolcanoAuction *VolcanoAuctionCaller) Auctions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	MinBid       *big.Int
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
}, error) {
	var out []interface{}
	err := _VolcanoAuction.contract.Call(opts, &out, "auctions", arg0, arg1)

	outstruct := new(struct {
		Owner        common.Address
		PayToken     common.Address
		MinBid       *big.Int
		ReservePrice *big.Int
		StartTime    *big.Int
		EndTime      *big.Int
		Resulted     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PayToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.MinBid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ReservePrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Resulted = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address owner, address payToken, uint256 minBid, uint256 reservePrice, uint256 startTime, uint256 endTime, bool resulted)
func (_VolcanoAuction *VolcanoAuctionSession) Auctions(arg0 common.Address, arg1 *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	MinBid       *big.Int
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
}, error) {
	return _VolcanoAuction.Contract.Auctions(&_VolcanoAuction.CallOpts, arg0, arg1)
}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address owner, address payToken, uint256 minBid, uint256 reservePrice, uint256 startTime, uint256 endTime, bool resulted)
func (_VolcanoAuction *VolcanoAuctionCallerSession) Auctions(arg0 common.Address, arg1 *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	MinBid       *big.Int
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
}, error) {
	return _VolcanoAuction.Contract.Auctions(&_VolcanoAuction.CallOpts, arg0, arg1)
}

// BidWithdrawalLockTime is a free data retrieval call binding the contract method 0x6a5362f6.
//
// Solidity: function bidWithdrawalLockTime() view returns(uint256)
func (_VolcanoAuction *VolcanoAuctionCaller) BidWithdrawalLockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoAuction.contract.Call(opts, &out, "bidWithdrawalLockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BidWithdrawalLockTime is a free data retrieval call binding the contract method 0x6a5362f6.
//
// Solidity: function bidWithdrawalLockTime() view returns(uint256)
func (_VolcanoAuction *VolcanoAuctionSession) BidWithdrawalLockTime() (*big.Int, error) {
	return _VolcanoAuction.Contract.BidWithdrawalLockTime(&_VolcanoAuction.CallOpts)
}

// BidWithdrawalLockTime is a free data retrieval call binding the contract method 0x6a5362f6.
//
// Solidity: function bidWithdrawalLockTime() view returns(uint256)
func (_VolcanoAuction *VolcanoAuctionCallerSession) BidWithdrawalLockTime() (*big.Int, error) {
	return _VolcanoAuction.Contract.BidWithdrawalLockTime(&_VolcanoAuction.CallOpts)
}

// GetAuction is a free data retrieval call binding the contract method 0x762a66a2.
//
// Solidity: function getAuction(address _nftAddress, uint256 _tokenId) view returns(address _owner, address _payToken, uint256 _reservePrice, uint256 _startTime, uint256 _endTime, bool _resulted, uint256 minBid)
func (_VolcanoAuction *VolcanoAuctionCaller) GetAuction(opts *bind.CallOpts, _nftAddress common.Address, _tokenId *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
	MinBid       *big.Int
}, error) {
	var out []interface{}
	err := _VolcanoAuction.contract.Call(opts, &out, "getAuction", _nftAddress, _tokenId)

	outstruct := new(struct {
		Owner        common.Address
		PayToken     common.Address
		ReservePrice *big.Int
		StartTime    *big.Int
		EndTime      *big.Int
		Resulted     bool
		MinBid       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PayToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ReservePrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Resulted = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.MinBid = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAuction is a free data retrieval call binding the contract method 0x762a66a2.
//
// Solidity: function getAuction(address _nftAddress, uint256 _tokenId) view returns(address _owner, address _payToken, uint256 _reservePrice, uint256 _startTime, uint256 _endTime, bool _resulted, uint256 minBid)
func (_VolcanoAuction *VolcanoAuctionSession) GetAuction(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
	MinBid       *big.Int
}, error) {
	return _VolcanoAuction.Contract.GetAuction(&_VolcanoAuction.CallOpts, _nftAddress, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x762a66a2.
//
// Solidity: function getAuction(address _nftAddress, uint256 _tokenId) view returns(address _owner, address _payToken, uint256 _reservePrice, uint256 _startTime, uint256 _endTime, bool _resulted, uint256 minBid)
func (_VolcanoAuction *VolcanoAuctionCallerSession) GetAuction(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Owner        common.Address
	PayToken     common.Address
	ReservePrice *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	Resulted     bool
	MinBid       *big.Int
}, error) {
	return _VolcanoAuction.Contract.GetAuction(&_VolcanoAuction.CallOpts, _nftAddress, _tokenId)
}

// GetHighestBidder is a free data retrieval call binding the contract method 0xa670a9b4.
//
// Solidity: function getHighestBidder(address _nftAddress, uint256 _tokenId) view returns(address _bidder, uint256 _bid, uint256 _lastBidTime)
func (_VolcanoAuction *VolcanoAuctionCaller) GetHighestBidder(opts *bind.CallOpts, _nftAddress common.Address, _tokenId *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	var out []interface{}
	err := _VolcanoAuction.contract.Call(opts, &out, "getHighestBidder", _nftAddress, _tokenId)

	outstruct := new(struct {
		Bidder      common.Address
		Bid         *big.Int
		LastBidTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bidder = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Bid = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastBidTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetHighestBidder is a free data retrieval call binding the contract method 0xa670a9b4.
//
// Solidity: function getHighestBidder(address _nftAddress, uint256 _tokenId) view returns(address _bidder, uint256 _bid, uint256 _lastBidTime)
func (_VolcanoAuction *VolcanoAuctionSession) GetHighestBidder(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _VolcanoAuction.Contract.GetHighestBidder(&_VolcanoAuction.CallOpts, _nftAddress, _tokenId)
}

// GetHighestBidder is a free data retrieval call binding the contract method 0xa670a9b4.
//
// Solidity: function getHighestBidder(address _nftAddress, uint256 _tokenId) view returns(address _bidder, uint256 _bid, uint256 _lastBidTime)
func (_VolcanoAuction *VolcanoAuctionCallerSession) GetHighestBidder(_nftAddress common.Address, _tokenId *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _VolcanoAuction.Contract.GetHighestBidder(&_VolcanoAuction.CallOpts, _nftAddress, _tokenId)
}

// HighestBids is a free data retrieval call binding the contract method 0x6b753f36.
//
// Solidity: function highestBids(address , uint256 ) view returns(address bidder, uint256 bid, uint256 lastBidTime)
func (_VolcanoAuction *VolcanoAuctionCaller) HighestBids(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	var out []interface{}
	err := _VolcanoAuction.contract.Call(opts, &out, "highestBids", arg0, arg1)

	outstruct := new(struct {
		Bidder      common.Address
		Bid         *big.Int
		LastBidTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bidder = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Bid = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastBidTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// HighestBids is a free data retrieval call binding the contract method 0x6b753f36.
//
// Solidity: function highestBids(address , uint256 ) view returns(address bidder, uint256 bid, uint256 lastBidTime)
func (_VolcanoAuction *VolcanoAuctionSession) HighestBids(arg0 common.Address, arg1 *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _VolcanoAuction.Contract.HighestBids(&_VolcanoAuction.CallOpts, arg0, arg1)
}

// HighestBids is a free data retrieval call binding the contract method 0x6b753f36.
//
// Solidity: function highestBids(address , uint256 ) view returns(address bidder, uint256 bid, uint256 lastBidTime)
func (_VolcanoAuction *VolcanoAuctionCallerSession) HighestBids(arg0 common.Address, arg1 *big.Int) (struct {
	Bidder      common.Address
	Bid         *big.Int
	LastBidTime *big.Int
}, error) {
	return _VolcanoAuction.Contract.HighestBids(&_VolcanoAuction.CallOpts, arg0, arg1)
}

// MinBidIncrement is a free data retrieval call binding the contract method 0x335b115e.
//
// Solidity: function minBidIncrement() view returns(uint256)
func (_VolcanoAuction *VolcanoAuctionCaller) MinBidIncrement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoAuction.contract.Call(opts, &out, "minBidIncrement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBidIncrement is a free data retrieval call binding the contract method 0x335b115e.
//
// Solidity: function minBidIncrement() view returns(uint256)
func (_VolcanoAuction *VolcanoAuctionSession) MinBidIncrement() (*big.Int, error) {
	return _VolcanoAuction.Contract.MinBidIncrement(&_VolcanoAuction.CallOpts)
}

// MinBidIncrement is a free data retrieval call binding the contract method 0x335b115e.
//
// Solidity: function minBidIncrement() view returns(uint256)
func (_VolcanoAuction *VolcanoAuctionCallerSession) MinBidIncrement() (*big.Int, error) {
	return _VolcanoAuction.Contract.MinBidIncrement(&_VolcanoAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoAuction *VolcanoAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoAuction *VolcanoAuctionSession) Owner() (common.Address, error) {
	return _VolcanoAuction.Contract.Owner(&_VolcanoAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoAuction *VolcanoAuctionCallerSession) Owner() (common.Address, error) {
	return _VolcanoAuction.Contract.Owner(&_VolcanoAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VolcanoAuction *VolcanoAuctionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VolcanoAuction.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VolcanoAuction *VolcanoAuctionSession) Paused() (bool, error) {
	return _VolcanoAuction.Contract.Paused(&_VolcanoAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VolcanoAuction *VolcanoAuctionCallerSession) Paused() (bool, error) {
	return _VolcanoAuction.Contract.Paused(&_VolcanoAuction.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VolcanoAuction *VolcanoAuctionCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoAuction.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VolcanoAuction *VolcanoAuctionSession) PlatformFee() (*big.Int, error) {
	return _VolcanoAuction.Contract.PlatformFee(&_VolcanoAuction.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_VolcanoAuction *VolcanoAuctionCallerSession) PlatformFee() (*big.Int, error) {
	return _VolcanoAuction.Contract.PlatformFee(&_VolcanoAuction.CallOpts)
}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_VolcanoAuction *VolcanoAuctionCaller) PlatformFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoAuction.contract.Call(opts, &out, "platformFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_VolcanoAuction *VolcanoAuctionSession) PlatformFeeRecipient() (common.Address, error) {
	return _VolcanoAuction.Contract.PlatformFeeRecipient(&_VolcanoAuction.CallOpts)
}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_VolcanoAuction *VolcanoAuctionCallerSession) PlatformFeeRecipient() (common.Address, error) {
	return _VolcanoAuction.Contract.PlatformFeeRecipient(&_VolcanoAuction.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VolcanoAuction *VolcanoAuctionCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VolcanoAuction.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VolcanoAuction *VolcanoAuctionSession) ProxiableUUID() ([32]byte, error) {
	return _VolcanoAuction.Contract.ProxiableUUID(&_VolcanoAuction.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_VolcanoAuction *VolcanoAuctionCallerSession) ProxiableUUID() ([32]byte, error) {
	return _VolcanoAuction.Contract.ProxiableUUID(&_VolcanoAuction.CallOpts)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x859b97fe.
//
// Solidity: function cancelAuction(address _nftAddress, uint256 _tokenId) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) CancelAuction(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "cancelAuction", _nftAddress, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x859b97fe.
//
// Solidity: function cancelAuction(address _nftAddress, uint256 _tokenId) returns()
func (_VolcanoAuction *VolcanoAuctionSession) CancelAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.CancelAuction(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x859b97fe.
//
// Solidity: function cancelAuction(address _nftAddress, uint256 _tokenId) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) CancelAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.CancelAuction(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xab2870e2.
//
// Solidity: function createAuction(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _reservePrice, uint256 _startTimestamp, bool minBidReserve, uint256 _endTimestamp) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) CreateAuction(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _reservePrice *big.Int, _startTimestamp *big.Int, minBidReserve bool, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "createAuction", _nftAddress, _tokenId, _payToken, _reservePrice, _startTimestamp, minBidReserve, _endTimestamp)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xab2870e2.
//
// Solidity: function createAuction(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _reservePrice, uint256 _startTimestamp, bool minBidReserve, uint256 _endTimestamp) returns()
func (_VolcanoAuction *VolcanoAuctionSession) CreateAuction(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _reservePrice *big.Int, _startTimestamp *big.Int, minBidReserve bool, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.CreateAuction(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId, _payToken, _reservePrice, _startTimestamp, minBidReserve, _endTimestamp)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xab2870e2.
//
// Solidity: function createAuction(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _reservePrice, uint256 _startTimestamp, bool minBidReserve, uint256 _endTimestamp) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) CreateAuction(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _reservePrice *big.Int, _startTimestamp *big.Int, minBidReserve bool, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.CreateAuction(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId, _payToken, _reservePrice, _startTimestamp, minBidReserve, _endTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _platformFeeRecipient, uint256 _platformFee) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) Initialize(opts *bind.TransactOpts, _platformFeeRecipient common.Address, _platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "initialize", _platformFeeRecipient, _platformFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _platformFeeRecipient, uint256 _platformFee) returns()
func (_VolcanoAuction *VolcanoAuctionSession) Initialize(_platformFeeRecipient common.Address, _platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.Initialize(&_VolcanoAuction.TransactOpts, _platformFeeRecipient, _platformFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _platformFeeRecipient, uint256 _platformFee) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) Initialize(_platformFeeRecipient common.Address, _platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.Initialize(&_VolcanoAuction.TransactOpts, _platformFeeRecipient, _platformFee)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x04717aca.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId, uint256 _bidAmount) payable returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) PlaceBid(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "placeBid", _nftAddress, _tokenId, _bidAmount)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x04717aca.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId, uint256 _bidAmount) payable returns()
func (_VolcanoAuction *VolcanoAuctionSession) PlaceBid(_nftAddress common.Address, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.PlaceBid(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId, _bidAmount)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x04717aca.
//
// Solidity: function placeBid(address _nftAddress, uint256 _tokenId, uint256 _bidAmount) payable returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) PlaceBid(_nftAddress common.Address, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.PlaceBid(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId, _bidAmount)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address _tokenContract) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) ReclaimERC20(opts *bind.TransactOpts, _tokenContract common.Address) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "reclaimERC20", _tokenContract)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address _tokenContract) returns()
func (_VolcanoAuction *VolcanoAuctionSession) ReclaimERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.ReclaimERC20(&_VolcanoAuction.TransactOpts, _tokenContract)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address _tokenContract) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) ReclaimERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.ReclaimERC20(&_VolcanoAuction.TransactOpts, _tokenContract)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoAuction *VolcanoAuctionSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoAuction.Contract.RenounceOwnership(&_VolcanoAuction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoAuction.Contract.RenounceOwnership(&_VolcanoAuction.TransactOpts)
}

// ResultAuction is a paid mutator transaction binding the contract method 0x5f7063ef.
//
// Solidity: function resultAuction(address _nftAddress, uint256 _tokenId) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) ResultAuction(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "resultAuction", _nftAddress, _tokenId)
}

// ResultAuction is a paid mutator transaction binding the contract method 0x5f7063ef.
//
// Solidity: function resultAuction(address _nftAddress, uint256 _tokenId) returns()
func (_VolcanoAuction *VolcanoAuctionSession) ResultAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.ResultAuction(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId)
}

// ResultAuction is a paid mutator transaction binding the contract method 0x5f7063ef.
//
// Solidity: function resultAuction(address _nftAddress, uint256 _tokenId) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) ResultAuction(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.ResultAuction(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId)
}

// ToggleIsPaused is a paid mutator transaction binding the contract method 0xf9613d7f.
//
// Solidity: function toggleIsPaused() returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) ToggleIsPaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "toggleIsPaused")
}

// ToggleIsPaused is a paid mutator transaction binding the contract method 0xf9613d7f.
//
// Solidity: function toggleIsPaused() returns()
func (_VolcanoAuction *VolcanoAuctionSession) ToggleIsPaused() (*types.Transaction, error) {
	return _VolcanoAuction.Contract.ToggleIsPaused(&_VolcanoAuction.TransactOpts)
}

// ToggleIsPaused is a paid mutator transaction binding the contract method 0xf9613d7f.
//
// Solidity: function toggleIsPaused() returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) ToggleIsPaused() (*types.Transaction, error) {
	return _VolcanoAuction.Contract.ToggleIsPaused(&_VolcanoAuction.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoAuction *VolcanoAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.TransferOwnership(&_VolcanoAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.TransferOwnership(&_VolcanoAuction.TransactOpts, newOwner)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) UpdateAddressRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "updateAddressRegistry", _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_VolcanoAuction *VolcanoAuctionSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdateAddressRegistry(&_VolcanoAuction.TransactOpts, _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdateAddressRegistry(&_VolcanoAuction.TransactOpts, _registry)
}

// UpdateAuctionEndTime is a paid mutator transaction binding the contract method 0x9ad2d9ec.
//
// Solidity: function updateAuctionEndTime(address _nftAddress, uint256 _tokenId, uint256 _endTimestamp) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) UpdateAuctionEndTime(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "updateAuctionEndTime", _nftAddress, _tokenId, _endTimestamp)
}

// UpdateAuctionEndTime is a paid mutator transaction binding the contract method 0x9ad2d9ec.
//
// Solidity: function updateAuctionEndTime(address _nftAddress, uint256 _tokenId, uint256 _endTimestamp) returns()
func (_VolcanoAuction *VolcanoAuctionSession) UpdateAuctionEndTime(_nftAddress common.Address, _tokenId *big.Int, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdateAuctionEndTime(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId, _endTimestamp)
}

// UpdateAuctionEndTime is a paid mutator transaction binding the contract method 0x9ad2d9ec.
//
// Solidity: function updateAuctionEndTime(address _nftAddress, uint256 _tokenId, uint256 _endTimestamp) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) UpdateAuctionEndTime(_nftAddress common.Address, _tokenId *big.Int, _endTimestamp *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdateAuctionEndTime(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId, _endTimestamp)
}

// UpdateAuctionReservePrice is a paid mutator transaction binding the contract method 0xa062feb9.
//
// Solidity: function updateAuctionReservePrice(address _nftAddress, uint256 _tokenId, uint256 _reservePrice) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) UpdateAuctionReservePrice(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _reservePrice *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "updateAuctionReservePrice", _nftAddress, _tokenId, _reservePrice)
}

// UpdateAuctionReservePrice is a paid mutator transaction binding the contract method 0xa062feb9.
//
// Solidity: function updateAuctionReservePrice(address _nftAddress, uint256 _tokenId, uint256 _reservePrice) returns()
func (_VolcanoAuction *VolcanoAuctionSession) UpdateAuctionReservePrice(_nftAddress common.Address, _tokenId *big.Int, _reservePrice *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdateAuctionReservePrice(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId, _reservePrice)
}

// UpdateAuctionReservePrice is a paid mutator transaction binding the contract method 0xa062feb9.
//
// Solidity: function updateAuctionReservePrice(address _nftAddress, uint256 _tokenId, uint256 _reservePrice) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) UpdateAuctionReservePrice(_nftAddress common.Address, _tokenId *big.Int, _reservePrice *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdateAuctionReservePrice(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId, _reservePrice)
}

// UpdateAuctionStartTime is a paid mutator transaction binding the contract method 0x6c58a5a4.
//
// Solidity: function updateAuctionStartTime(address _nftAddress, uint256 _tokenId, uint256 _startTime) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) UpdateAuctionStartTime(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _startTime *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "updateAuctionStartTime", _nftAddress, _tokenId, _startTime)
}

// UpdateAuctionStartTime is a paid mutator transaction binding the contract method 0x6c58a5a4.
//
// Solidity: function updateAuctionStartTime(address _nftAddress, uint256 _tokenId, uint256 _startTime) returns()
func (_VolcanoAuction *VolcanoAuctionSession) UpdateAuctionStartTime(_nftAddress common.Address, _tokenId *big.Int, _startTime *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdateAuctionStartTime(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId, _startTime)
}

// UpdateAuctionStartTime is a paid mutator transaction binding the contract method 0x6c58a5a4.
//
// Solidity: function updateAuctionStartTime(address _nftAddress, uint256 _tokenId, uint256 _startTime) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) UpdateAuctionStartTime(_nftAddress common.Address, _tokenId *big.Int, _startTime *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdateAuctionStartTime(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId, _startTime)
}

// UpdateBidWithdrawalLockTime is a paid mutator transaction binding the contract method 0x623de5ad.
//
// Solidity: function updateBidWithdrawalLockTime(uint256 _bidWithdrawalLockTime) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) UpdateBidWithdrawalLockTime(opts *bind.TransactOpts, _bidWithdrawalLockTime *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "updateBidWithdrawalLockTime", _bidWithdrawalLockTime)
}

// UpdateBidWithdrawalLockTime is a paid mutator transaction binding the contract method 0x623de5ad.
//
// Solidity: function updateBidWithdrawalLockTime(uint256 _bidWithdrawalLockTime) returns()
func (_VolcanoAuction *VolcanoAuctionSession) UpdateBidWithdrawalLockTime(_bidWithdrawalLockTime *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdateBidWithdrawalLockTime(&_VolcanoAuction.TransactOpts, _bidWithdrawalLockTime)
}

// UpdateBidWithdrawalLockTime is a paid mutator transaction binding the contract method 0x623de5ad.
//
// Solidity: function updateBidWithdrawalLockTime(uint256 _bidWithdrawalLockTime) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) UpdateBidWithdrawalLockTime(_bidWithdrawalLockTime *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdateBidWithdrawalLockTime(&_VolcanoAuction.TransactOpts, _bidWithdrawalLockTime)
}

// UpdateMinBidIncrement is a paid mutator transaction binding the contract method 0x083759e0.
//
// Solidity: function updateMinBidIncrement(uint256 _minBidIncrement) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) UpdateMinBidIncrement(opts *bind.TransactOpts, _minBidIncrement *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "updateMinBidIncrement", _minBidIncrement)
}

// UpdateMinBidIncrement is a paid mutator transaction binding the contract method 0x083759e0.
//
// Solidity: function updateMinBidIncrement(uint256 _minBidIncrement) returns()
func (_VolcanoAuction *VolcanoAuctionSession) UpdateMinBidIncrement(_minBidIncrement *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdateMinBidIncrement(&_VolcanoAuction.TransactOpts, _minBidIncrement)
}

// UpdateMinBidIncrement is a paid mutator transaction binding the contract method 0x083759e0.
//
// Solidity: function updateMinBidIncrement(uint256 _minBidIncrement) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) UpdateMinBidIncrement(_minBidIncrement *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdateMinBidIncrement(&_VolcanoAuction.TransactOpts, _minBidIncrement)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) UpdatePlatformFee(opts *bind.TransactOpts, _platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "updatePlatformFee", _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_VolcanoAuction *VolcanoAuctionSession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdatePlatformFee(&_VolcanoAuction.TransactOpts, _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xaa0b5988.
//
// Solidity: function updatePlatformFee(uint256 _platformFee) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) UpdatePlatformFee(_platformFee *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdatePlatformFee(&_VolcanoAuction.TransactOpts, _platformFee)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) UpdatePlatformFeeRecipient(opts *bind.TransactOpts, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "updatePlatformFeeRecipient", _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_VolcanoAuction *VolcanoAuctionSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdatePlatformFeeRecipient(&_VolcanoAuction.TransactOpts, _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpdatePlatformFeeRecipient(&_VolcanoAuction.TransactOpts, _platformFeeRecipient)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VolcanoAuction *VolcanoAuctionSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpgradeTo(&_VolcanoAuction.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpgradeTo(&_VolcanoAuction.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VolcanoAuction *VolcanoAuctionSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpgradeToAndCall(&_VolcanoAuction.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.UpgradeToAndCall(&_VolcanoAuction.TransactOpts, newImplementation, data)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x4cb8ef5b.
//
// Solidity: function withdrawBid(address _nftAddress, uint256 _tokenId) returns()
func (_VolcanoAuction *VolcanoAuctionTransactor) WithdrawBid(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.contract.Transact(opts, "withdrawBid", _nftAddress, _tokenId)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x4cb8ef5b.
//
// Solidity: function withdrawBid(address _nftAddress, uint256 _tokenId) returns()
func (_VolcanoAuction *VolcanoAuctionSession) WithdrawBid(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.WithdrawBid(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId)
}

// WithdrawBid is a paid mutator transaction binding the contract method 0x4cb8ef5b.
//
// Solidity: function withdrawBid(address _nftAddress, uint256 _tokenId) returns()
func (_VolcanoAuction *VolcanoAuctionTransactorSession) WithdrawBid(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _VolcanoAuction.Contract.WithdrawBid(&_VolcanoAuction.TransactOpts, _nftAddress, _tokenId)
}

// VolcanoAuctionAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the VolcanoAuction contract.
type VolcanoAuctionAdminChangedIterator struct {
	Event *VolcanoAuctionAdminChanged // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionAdminChanged)
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
		it.Event = new(VolcanoAuctionAdminChanged)
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
func (it *VolcanoAuctionAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionAdminChanged represents a AdminChanged event raised by the VolcanoAuction contract.
type VolcanoAuctionAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*VolcanoAuctionAdminChangedIterator, error) {

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionAdminChangedIterator{contract: _VolcanoAuction.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionAdminChanged) (event.Subscription, error) {

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionAdminChanged)
				if err := _VolcanoAuction.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseAdminChanged(log types.Log) (*VolcanoAuctionAdminChanged, error) {
	event := new(VolcanoAuctionAdminChanged)
	if err := _VolcanoAuction.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the VolcanoAuction contract.
type VolcanoAuctionAuctionCancelledIterator struct {
	Event *VolcanoAuctionAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionAuctionCancelled)
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
		it.Event = new(VolcanoAuctionAuctionCancelled)
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
func (it *VolcanoAuctionAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionAuctionCancelled represents a AuctionCancelled event raised by the VolcanoAuction contract.
type VolcanoAuctionAuctionCancelled struct {
	NftAddress common.Address
	TokenId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x018b64b6242d32aa550e95d78985b938d71af5b3f10827b0683f55da16393048.
//
// Solidity: event AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterAuctionCancelled(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*VolcanoAuctionAuctionCancelledIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "AuctionCancelled", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionAuctionCancelledIterator{contract: _VolcanoAuction.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x018b64b6242d32aa550e95d78985b938d71af5b3f10827b0683f55da16393048.
//
// Solidity: event AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionAuctionCancelled, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "AuctionCancelled", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionAuctionCancelled)
				if err := _VolcanoAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// ParseAuctionCancelled is a log parse operation binding the contract event 0x018b64b6242d32aa550e95d78985b938d71af5b3f10827b0683f55da16393048.
//
// Solidity: event AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId)
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseAuctionCancelled(log types.Log) (*VolcanoAuctionAuctionCancelled, error) {
	event := new(VolcanoAuctionAuctionCancelled)
	if err := _VolcanoAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the VolcanoAuction contract.
type VolcanoAuctionAuctionCreatedIterator struct {
	Event *VolcanoAuctionAuctionCreated // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionAuctionCreated)
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
		it.Event = new(VolcanoAuctionAuctionCreated)
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
func (it *VolcanoAuctionAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionAuctionCreated represents a AuctionCreated event raised by the VolcanoAuction contract.
type VolcanoAuctionAuctionCreated struct {
	NftAddress common.Address
	TokenId    *big.Int
	PayToken   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xca437d90ed6373b827a01275bd2fdfe7e6406d7ecd400662ee0533d3546ab17a.
//
// Solidity: event AuctionCreated(address indexed nftAddress, uint256 indexed tokenId, address payToken)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterAuctionCreated(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*VolcanoAuctionAuctionCreatedIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "AuctionCreated", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionAuctionCreatedIterator{contract: _VolcanoAuction.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xca437d90ed6373b827a01275bd2fdfe7e6406d7ecd400662ee0533d3546ab17a.
//
// Solidity: event AuctionCreated(address indexed nftAddress, uint256 indexed tokenId, address payToken)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionAuctionCreated, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "AuctionCreated", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionAuctionCreated)
				if err := _VolcanoAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xca437d90ed6373b827a01275bd2fdfe7e6406d7ecd400662ee0533d3546ab17a.
//
// Solidity: event AuctionCreated(address indexed nftAddress, uint256 indexed tokenId, address payToken)
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseAuctionCreated(log types.Log) (*VolcanoAuctionAuctionCreated, error) {
	event := new(VolcanoAuctionAuctionCreated)
	if err := _VolcanoAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionAuctionResultedIterator is returned from FilterAuctionResulted and is used to iterate over the raw logs and unpacked data for AuctionResulted events raised by the VolcanoAuction contract.
type VolcanoAuctionAuctionResultedIterator struct {
	Event *VolcanoAuctionAuctionResulted // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionAuctionResultedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionAuctionResulted)
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
		it.Event = new(VolcanoAuctionAuctionResulted)
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
func (it *VolcanoAuctionAuctionResultedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionAuctionResultedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionAuctionResulted represents a AuctionResulted event raised by the VolcanoAuction contract.
type VolcanoAuctionAuctionResulted struct {
	OldOwner      common.Address
	NftAddress    common.Address
	TokenId       *big.Int
	Winner        common.Address
	PayToken      common.Address
	PayTokenPrice *big.Int
	WinningBid    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionResulted is a free log retrieval operation binding the contract event 0xa717395bf4be1915f0bc8e6cf9b0f526c5ad40a0c3750b709a29a834daf0fd9b.
//
// Solidity: event AuctionResulted(address oldOwner, address indexed nftAddress, uint256 indexed tokenId, address indexed winner, address payToken, int256 payTokenPrice, uint256 winningBid)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterAuctionResulted(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, winner []common.Address) (*VolcanoAuctionAuctionResultedIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "AuctionResulted", nftAddressRule, tokenIdRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionAuctionResultedIterator{contract: _VolcanoAuction.contract, event: "AuctionResulted", logs: logs, sub: sub}, nil
}

// WatchAuctionResulted is a free log subscription operation binding the contract event 0xa717395bf4be1915f0bc8e6cf9b0f526c5ad40a0c3750b709a29a834daf0fd9b.
//
// Solidity: event AuctionResulted(address oldOwner, address indexed nftAddress, uint256 indexed tokenId, address indexed winner, address payToken, int256 payTokenPrice, uint256 winningBid)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchAuctionResulted(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionAuctionResulted, nftAddress []common.Address, tokenId []*big.Int, winner []common.Address) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "AuctionResulted", nftAddressRule, tokenIdRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionAuctionResulted)
				if err := _VolcanoAuction.contract.UnpackLog(event, "AuctionResulted", log); err != nil {
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

// ParseAuctionResulted is a log parse operation binding the contract event 0xa717395bf4be1915f0bc8e6cf9b0f526c5ad40a0c3750b709a29a834daf0fd9b.
//
// Solidity: event AuctionResulted(address oldOwner, address indexed nftAddress, uint256 indexed tokenId, address indexed winner, address payToken, int256 payTokenPrice, uint256 winningBid)
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseAuctionResulted(log types.Log) (*VolcanoAuctionAuctionResulted, error) {
	event := new(VolcanoAuctionAuctionResulted)
	if err := _VolcanoAuction.contract.UnpackLog(event, "AuctionResulted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the VolcanoAuction contract.
type VolcanoAuctionBeaconUpgradedIterator struct {
	Event *VolcanoAuctionBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionBeaconUpgraded)
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
		it.Event = new(VolcanoAuctionBeaconUpgraded)
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
func (it *VolcanoAuctionBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionBeaconUpgraded represents a BeaconUpgraded event raised by the VolcanoAuction contract.
type VolcanoAuctionBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*VolcanoAuctionBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionBeaconUpgradedIterator{contract: _VolcanoAuction.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionBeaconUpgraded)
				if err := _VolcanoAuction.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseBeaconUpgraded(log types.Log) (*VolcanoAuctionBeaconUpgraded, error) {
	event := new(VolcanoAuctionBeaconUpgraded)
	if err := _VolcanoAuction.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the VolcanoAuction contract.
type VolcanoAuctionBidPlacedIterator struct {
	Event *VolcanoAuctionBidPlaced // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionBidPlaced)
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
		it.Event = new(VolcanoAuctionBidPlaced)
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
func (it *VolcanoAuctionBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionBidPlaced represents a BidPlaced event raised by the VolcanoAuction contract.
type VolcanoAuctionBidPlaced struct {
	NftAddress common.Address
	TokenId    *big.Int
	Bidder     common.Address
	Bid        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x0158f5674dc243762459b88cfc91b10d2d1ef9d40821cca978c2b680aa444682.
//
// Solidity: event BidPlaced(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterBidPlaced(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (*VolcanoAuctionBidPlacedIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "BidPlaced", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionBidPlacedIterator{contract: _VolcanoAuction.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x0158f5674dc243762459b88cfc91b10d2d1ef9d40821cca978c2b680aa444682.
//
// Solidity: event BidPlaced(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionBidPlaced, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "BidPlaced", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionBidPlaced)
				if err := _VolcanoAuction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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

// ParseBidPlaced is a log parse operation binding the contract event 0x0158f5674dc243762459b88cfc91b10d2d1ef9d40821cca978c2b680aa444682.
//
// Solidity: event BidPlaced(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseBidPlaced(log types.Log) (*VolcanoAuctionBidPlaced, error) {
	event := new(VolcanoAuctionBidPlaced)
	if err := _VolcanoAuction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionBidRefundedIterator is returned from FilterBidRefunded and is used to iterate over the raw logs and unpacked data for BidRefunded events raised by the VolcanoAuction contract.
type VolcanoAuctionBidRefundedIterator struct {
	Event *VolcanoAuctionBidRefunded // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionBidRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionBidRefunded)
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
		it.Event = new(VolcanoAuctionBidRefunded)
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
func (it *VolcanoAuctionBidRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionBidRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionBidRefunded represents a BidRefunded event raised by the VolcanoAuction contract.
type VolcanoAuctionBidRefunded struct {
	NftAddress common.Address
	TokenId    *big.Int
	Bidder     common.Address
	Bid        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidRefunded is a free log retrieval operation binding the contract event 0x90e20d1ba82eaa07a212267355536f76bc83bf91c81fd41b74283c4082e76952.
//
// Solidity: event BidRefunded(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterBidRefunded(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (*VolcanoAuctionBidRefundedIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "BidRefunded", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionBidRefundedIterator{contract: _VolcanoAuction.contract, event: "BidRefunded", logs: logs, sub: sub}, nil
}

// WatchBidRefunded is a free log subscription operation binding the contract event 0x90e20d1ba82eaa07a212267355536f76bc83bf91c81fd41b74283c4082e76952.
//
// Solidity: event BidRefunded(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchBidRefunded(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionBidRefunded, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "BidRefunded", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionBidRefunded)
				if err := _VolcanoAuction.contract.UnpackLog(event, "BidRefunded", log); err != nil {
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

// ParseBidRefunded is a log parse operation binding the contract event 0x90e20d1ba82eaa07a212267355536f76bc83bf91c81fd41b74283c4082e76952.
//
// Solidity: event BidRefunded(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseBidRefunded(log types.Log) (*VolcanoAuctionBidRefunded, error) {
	event := new(VolcanoAuctionBidRefunded)
	if err := _VolcanoAuction.contract.UnpackLog(event, "BidRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionBidWithdrawnIterator is returned from FilterBidWithdrawn and is used to iterate over the raw logs and unpacked data for BidWithdrawn events raised by the VolcanoAuction contract.
type VolcanoAuctionBidWithdrawnIterator struct {
	Event *VolcanoAuctionBidWithdrawn // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionBidWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionBidWithdrawn)
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
		it.Event = new(VolcanoAuctionBidWithdrawn)
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
func (it *VolcanoAuctionBidWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionBidWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionBidWithdrawn represents a BidWithdrawn event raised by the VolcanoAuction contract.
type VolcanoAuctionBidWithdrawn struct {
	NftAddress common.Address
	TokenId    *big.Int
	Bidder     common.Address
	Bid        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidWithdrawn is a free log retrieval operation binding the contract event 0x867b8ea96dd803063f905a19f8117cbb1866ec7a594dfede75ab4a5235f61d7c.
//
// Solidity: event BidWithdrawn(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterBidWithdrawn(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (*VolcanoAuctionBidWithdrawnIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "BidWithdrawn", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionBidWithdrawnIterator{contract: _VolcanoAuction.contract, event: "BidWithdrawn", logs: logs, sub: sub}, nil
}

// WatchBidWithdrawn is a free log subscription operation binding the contract event 0x867b8ea96dd803063f905a19f8117cbb1866ec7a594dfede75ab4a5235f61d7c.
//
// Solidity: event BidWithdrawn(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchBidWithdrawn(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionBidWithdrawn, nftAddress []common.Address, tokenId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "BidWithdrawn", nftAddressRule, tokenIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionBidWithdrawn)
				if err := _VolcanoAuction.contract.UnpackLog(event, "BidWithdrawn", log); err != nil {
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

// ParseBidWithdrawn is a log parse operation binding the contract event 0x867b8ea96dd803063f905a19f8117cbb1866ec7a594dfede75ab4a5235f61d7c.
//
// Solidity: event BidWithdrawn(address indexed nftAddress, uint256 indexed tokenId, address indexed bidder, uint256 bid)
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseBidWithdrawn(log types.Log) (*VolcanoAuctionBidWithdrawn, error) {
	event := new(VolcanoAuctionBidWithdrawn)
	if err := _VolcanoAuction.contract.UnpackLog(event, "BidWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VolcanoAuction contract.
type VolcanoAuctionInitializedIterator struct {
	Event *VolcanoAuctionInitialized // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionInitialized)
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
		it.Event = new(VolcanoAuctionInitialized)
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
func (it *VolcanoAuctionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionInitialized represents a Initialized event raised by the VolcanoAuction contract.
type VolcanoAuctionInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterInitialized(opts *bind.FilterOpts) (*VolcanoAuctionInitializedIterator, error) {

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionInitializedIterator{contract: _VolcanoAuction.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionInitialized) (event.Subscription, error) {

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionInitialized)
				if err := _VolcanoAuction.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseInitialized(log types.Log) (*VolcanoAuctionInitialized, error) {
	event := new(VolcanoAuctionInitialized)
	if err := _VolcanoAuction.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VolcanoAuction contract.
type VolcanoAuctionOwnershipTransferredIterator struct {
	Event *VolcanoAuctionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionOwnershipTransferred)
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
		it.Event = new(VolcanoAuctionOwnershipTransferred)
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
func (it *VolcanoAuctionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionOwnershipTransferred represents a OwnershipTransferred event raised by the VolcanoAuction contract.
type VolcanoAuctionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VolcanoAuctionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionOwnershipTransferredIterator{contract: _VolcanoAuction.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionOwnershipTransferred)
				if err := _VolcanoAuction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseOwnershipTransferred(log types.Log) (*VolcanoAuctionOwnershipTransferred, error) {
	event := new(VolcanoAuctionOwnershipTransferred)
	if err := _VolcanoAuction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionPauseToggledIterator is returned from FilterPauseToggled and is used to iterate over the raw logs and unpacked data for PauseToggled events raised by the VolcanoAuction contract.
type VolcanoAuctionPauseToggledIterator struct {
	Event *VolcanoAuctionPauseToggled // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionPauseToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionPauseToggled)
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
		it.Event = new(VolcanoAuctionPauseToggled)
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
func (it *VolcanoAuctionPauseToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionPauseToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionPauseToggled represents a PauseToggled event raised by the VolcanoAuction contract.
type VolcanoAuctionPauseToggled struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPauseToggled is a free log retrieval operation binding the contract event 0x9077d36bc00859b5c3f320310707208543dd35092cb0a0fe117d0c6a558b148b.
//
// Solidity: event PauseToggled(bool isPaused)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterPauseToggled(opts *bind.FilterOpts) (*VolcanoAuctionPauseToggledIterator, error) {

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "PauseToggled")
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionPauseToggledIterator{contract: _VolcanoAuction.contract, event: "PauseToggled", logs: logs, sub: sub}, nil
}

// WatchPauseToggled is a free log subscription operation binding the contract event 0x9077d36bc00859b5c3f320310707208543dd35092cb0a0fe117d0c6a558b148b.
//
// Solidity: event PauseToggled(bool isPaused)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchPauseToggled(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionPauseToggled) (event.Subscription, error) {

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "PauseToggled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionPauseToggled)
				if err := _VolcanoAuction.contract.UnpackLog(event, "PauseToggled", log); err != nil {
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

// ParsePauseToggled is a log parse operation binding the contract event 0x9077d36bc00859b5c3f320310707208543dd35092cb0a0fe117d0c6a558b148b.
//
// Solidity: event PauseToggled(bool isPaused)
func (_VolcanoAuction *VolcanoAuctionFilterer) ParsePauseToggled(log types.Log) (*VolcanoAuctionPauseToggled, error) {
	event := new(VolcanoAuctionPauseToggled)
	if err := _VolcanoAuction.contract.UnpackLog(event, "PauseToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the VolcanoAuction contract.
type VolcanoAuctionPausedIterator struct {
	Event *VolcanoAuctionPaused // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionPaused)
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
		it.Event = new(VolcanoAuctionPaused)
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
func (it *VolcanoAuctionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionPaused represents a Paused event raised by the VolcanoAuction contract.
type VolcanoAuctionPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterPaused(opts *bind.FilterOpts) (*VolcanoAuctionPausedIterator, error) {

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionPausedIterator{contract: _VolcanoAuction.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionPaused) (event.Subscription, error) {

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionPaused)
				if err := _VolcanoAuction.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_VolcanoAuction *VolcanoAuctionFilterer) ParsePaused(log types.Log) (*VolcanoAuctionPaused, error) {
	event := new(VolcanoAuctionPaused)
	if err := _VolcanoAuction.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the VolcanoAuction contract.
type VolcanoAuctionUnpausedIterator struct {
	Event *VolcanoAuctionUnpaused // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionUnpaused)
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
		it.Event = new(VolcanoAuctionUnpaused)
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
func (it *VolcanoAuctionUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionUnpaused represents a Unpaused event raised by the VolcanoAuction contract.
type VolcanoAuctionUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VolcanoAuctionUnpausedIterator, error) {

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionUnpausedIterator{contract: _VolcanoAuction.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionUnpaused) (event.Subscription, error) {

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionUnpaused)
				if err := _VolcanoAuction.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseUnpaused(log types.Log) (*VolcanoAuctionUnpaused, error) {
	event := new(VolcanoAuctionUnpaused)
	if err := _VolcanoAuction.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionUpdateAuctionEndTimeIterator is returned from FilterUpdateAuctionEndTime and is used to iterate over the raw logs and unpacked data for UpdateAuctionEndTime events raised by the VolcanoAuction contract.
type VolcanoAuctionUpdateAuctionEndTimeIterator struct {
	Event *VolcanoAuctionUpdateAuctionEndTime // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionUpdateAuctionEndTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionUpdateAuctionEndTime)
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
		it.Event = new(VolcanoAuctionUpdateAuctionEndTime)
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
func (it *VolcanoAuctionUpdateAuctionEndTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionUpdateAuctionEndTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionUpdateAuctionEndTime represents a UpdateAuctionEndTime event raised by the VolcanoAuction contract.
type VolcanoAuctionUpdateAuctionEndTime struct {
	NftAddress common.Address
	TokenId    *big.Int
	EndTime    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateAuctionEndTime is a free log retrieval operation binding the contract event 0xd8aaad775c4c9b21112b5cb22ab4618c74936d873850fbd6f2e691701119ffe1.
//
// Solidity: event UpdateAuctionEndTime(address indexed nftAddress, uint256 indexed tokenId, uint256 endTime)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterUpdateAuctionEndTime(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*VolcanoAuctionUpdateAuctionEndTimeIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "UpdateAuctionEndTime", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionUpdateAuctionEndTimeIterator{contract: _VolcanoAuction.contract, event: "UpdateAuctionEndTime", logs: logs, sub: sub}, nil
}

// WatchUpdateAuctionEndTime is a free log subscription operation binding the contract event 0xd8aaad775c4c9b21112b5cb22ab4618c74936d873850fbd6f2e691701119ffe1.
//
// Solidity: event UpdateAuctionEndTime(address indexed nftAddress, uint256 indexed tokenId, uint256 endTime)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchUpdateAuctionEndTime(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionUpdateAuctionEndTime, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "UpdateAuctionEndTime", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionUpdateAuctionEndTime)
				if err := _VolcanoAuction.contract.UnpackLog(event, "UpdateAuctionEndTime", log); err != nil {
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

// ParseUpdateAuctionEndTime is a log parse operation binding the contract event 0xd8aaad775c4c9b21112b5cb22ab4618c74936d873850fbd6f2e691701119ffe1.
//
// Solidity: event UpdateAuctionEndTime(address indexed nftAddress, uint256 indexed tokenId, uint256 endTime)
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseUpdateAuctionEndTime(log types.Log) (*VolcanoAuctionUpdateAuctionEndTime, error) {
	event := new(VolcanoAuctionUpdateAuctionEndTime)
	if err := _VolcanoAuction.contract.UnpackLog(event, "UpdateAuctionEndTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionUpdateAuctionReservePriceIterator is returned from FilterUpdateAuctionReservePrice and is used to iterate over the raw logs and unpacked data for UpdateAuctionReservePrice events raised by the VolcanoAuction contract.
type VolcanoAuctionUpdateAuctionReservePriceIterator struct {
	Event *VolcanoAuctionUpdateAuctionReservePrice // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionUpdateAuctionReservePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionUpdateAuctionReservePrice)
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
		it.Event = new(VolcanoAuctionUpdateAuctionReservePrice)
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
func (it *VolcanoAuctionUpdateAuctionReservePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionUpdateAuctionReservePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionUpdateAuctionReservePrice represents a UpdateAuctionReservePrice event raised by the VolcanoAuction contract.
type VolcanoAuctionUpdateAuctionReservePrice struct {
	NftAddress   common.Address
	TokenId      *big.Int
	PayToken     common.Address
	ReservePrice *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateAuctionReservePrice is a free log retrieval operation binding the contract event 0x1492815b4a41d34cc0015f8ba8013fa45717ea6d8591db6036bdca706981fe83.
//
// Solidity: event UpdateAuctionReservePrice(address indexed nftAddress, uint256 indexed tokenId, address payToken, uint256 reservePrice)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterUpdateAuctionReservePrice(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*VolcanoAuctionUpdateAuctionReservePriceIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "UpdateAuctionReservePrice", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionUpdateAuctionReservePriceIterator{contract: _VolcanoAuction.contract, event: "UpdateAuctionReservePrice", logs: logs, sub: sub}, nil
}

// WatchUpdateAuctionReservePrice is a free log subscription operation binding the contract event 0x1492815b4a41d34cc0015f8ba8013fa45717ea6d8591db6036bdca706981fe83.
//
// Solidity: event UpdateAuctionReservePrice(address indexed nftAddress, uint256 indexed tokenId, address payToken, uint256 reservePrice)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchUpdateAuctionReservePrice(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionUpdateAuctionReservePrice, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "UpdateAuctionReservePrice", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionUpdateAuctionReservePrice)
				if err := _VolcanoAuction.contract.UnpackLog(event, "UpdateAuctionReservePrice", log); err != nil {
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

// ParseUpdateAuctionReservePrice is a log parse operation binding the contract event 0x1492815b4a41d34cc0015f8ba8013fa45717ea6d8591db6036bdca706981fe83.
//
// Solidity: event UpdateAuctionReservePrice(address indexed nftAddress, uint256 indexed tokenId, address payToken, uint256 reservePrice)
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseUpdateAuctionReservePrice(log types.Log) (*VolcanoAuctionUpdateAuctionReservePrice, error) {
	event := new(VolcanoAuctionUpdateAuctionReservePrice)
	if err := _VolcanoAuction.contract.UnpackLog(event, "UpdateAuctionReservePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionUpdateAuctionStartTimeIterator is returned from FilterUpdateAuctionStartTime and is used to iterate over the raw logs and unpacked data for UpdateAuctionStartTime events raised by the VolcanoAuction contract.
type VolcanoAuctionUpdateAuctionStartTimeIterator struct {
	Event *VolcanoAuctionUpdateAuctionStartTime // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionUpdateAuctionStartTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionUpdateAuctionStartTime)
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
		it.Event = new(VolcanoAuctionUpdateAuctionStartTime)
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
func (it *VolcanoAuctionUpdateAuctionStartTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionUpdateAuctionStartTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionUpdateAuctionStartTime represents a UpdateAuctionStartTime event raised by the VolcanoAuction contract.
type VolcanoAuctionUpdateAuctionStartTime struct {
	NftAddress common.Address
	TokenId    *big.Int
	StartTime  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateAuctionStartTime is a free log retrieval operation binding the contract event 0xf123182d1bfce603eb20a2f74cce94af6542f353a4c2a8d5cfadf41822b32b38.
//
// Solidity: event UpdateAuctionStartTime(address indexed nftAddress, uint256 indexed tokenId, uint256 startTime)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterUpdateAuctionStartTime(opts *bind.FilterOpts, nftAddress []common.Address, tokenId []*big.Int) (*VolcanoAuctionUpdateAuctionStartTimeIterator, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "UpdateAuctionStartTime", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionUpdateAuctionStartTimeIterator{contract: _VolcanoAuction.contract, event: "UpdateAuctionStartTime", logs: logs, sub: sub}, nil
}

// WatchUpdateAuctionStartTime is a free log subscription operation binding the contract event 0xf123182d1bfce603eb20a2f74cce94af6542f353a4c2a8d5cfadf41822b32b38.
//
// Solidity: event UpdateAuctionStartTime(address indexed nftAddress, uint256 indexed tokenId, uint256 startTime)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchUpdateAuctionStartTime(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionUpdateAuctionStartTime, nftAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var nftAddressRule []interface{}
	for _, nftAddressItem := range nftAddress {
		nftAddressRule = append(nftAddressRule, nftAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "UpdateAuctionStartTime", nftAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionUpdateAuctionStartTime)
				if err := _VolcanoAuction.contract.UnpackLog(event, "UpdateAuctionStartTime", log); err != nil {
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

// ParseUpdateAuctionStartTime is a log parse operation binding the contract event 0xf123182d1bfce603eb20a2f74cce94af6542f353a4c2a8d5cfadf41822b32b38.
//
// Solidity: event UpdateAuctionStartTime(address indexed nftAddress, uint256 indexed tokenId, uint256 startTime)
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseUpdateAuctionStartTime(log types.Log) (*VolcanoAuctionUpdateAuctionStartTime, error) {
	event := new(VolcanoAuctionUpdateAuctionStartTime)
	if err := _VolcanoAuction.contract.UnpackLog(event, "UpdateAuctionStartTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionUpdateBidWithdrawalLockTimeIterator is returned from FilterUpdateBidWithdrawalLockTime and is used to iterate over the raw logs and unpacked data for UpdateBidWithdrawalLockTime events raised by the VolcanoAuction contract.
type VolcanoAuctionUpdateBidWithdrawalLockTimeIterator struct {
	Event *VolcanoAuctionUpdateBidWithdrawalLockTime // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionUpdateBidWithdrawalLockTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionUpdateBidWithdrawalLockTime)
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
		it.Event = new(VolcanoAuctionUpdateBidWithdrawalLockTime)
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
func (it *VolcanoAuctionUpdateBidWithdrawalLockTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionUpdateBidWithdrawalLockTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionUpdateBidWithdrawalLockTime represents a UpdateBidWithdrawalLockTime event raised by the VolcanoAuction contract.
type VolcanoAuctionUpdateBidWithdrawalLockTime struct {
	BidWithdrawalLockTime *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdateBidWithdrawalLockTime is a free log retrieval operation binding the contract event 0xf9e4b69944d3fbdd96aedcda02032e2091346ececc3b55a485b40a6dc09bb9e0.
//
// Solidity: event UpdateBidWithdrawalLockTime(uint256 bidWithdrawalLockTime)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterUpdateBidWithdrawalLockTime(opts *bind.FilterOpts) (*VolcanoAuctionUpdateBidWithdrawalLockTimeIterator, error) {

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "UpdateBidWithdrawalLockTime")
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionUpdateBidWithdrawalLockTimeIterator{contract: _VolcanoAuction.contract, event: "UpdateBidWithdrawalLockTime", logs: logs, sub: sub}, nil
}

// WatchUpdateBidWithdrawalLockTime is a free log subscription operation binding the contract event 0xf9e4b69944d3fbdd96aedcda02032e2091346ececc3b55a485b40a6dc09bb9e0.
//
// Solidity: event UpdateBidWithdrawalLockTime(uint256 bidWithdrawalLockTime)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchUpdateBidWithdrawalLockTime(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionUpdateBidWithdrawalLockTime) (event.Subscription, error) {

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "UpdateBidWithdrawalLockTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionUpdateBidWithdrawalLockTime)
				if err := _VolcanoAuction.contract.UnpackLog(event, "UpdateBidWithdrawalLockTime", log); err != nil {
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

// ParseUpdateBidWithdrawalLockTime is a log parse operation binding the contract event 0xf9e4b69944d3fbdd96aedcda02032e2091346ececc3b55a485b40a6dc09bb9e0.
//
// Solidity: event UpdateBidWithdrawalLockTime(uint256 bidWithdrawalLockTime)
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseUpdateBidWithdrawalLockTime(log types.Log) (*VolcanoAuctionUpdateBidWithdrawalLockTime, error) {
	event := new(VolcanoAuctionUpdateBidWithdrawalLockTime)
	if err := _VolcanoAuction.contract.UnpackLog(event, "UpdateBidWithdrawalLockTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionUpdateMinBidIncrementIterator is returned from FilterUpdateMinBidIncrement and is used to iterate over the raw logs and unpacked data for UpdateMinBidIncrement events raised by the VolcanoAuction contract.
type VolcanoAuctionUpdateMinBidIncrementIterator struct {
	Event *VolcanoAuctionUpdateMinBidIncrement // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionUpdateMinBidIncrementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionUpdateMinBidIncrement)
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
		it.Event = new(VolcanoAuctionUpdateMinBidIncrement)
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
func (it *VolcanoAuctionUpdateMinBidIncrementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionUpdateMinBidIncrementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionUpdateMinBidIncrement represents a UpdateMinBidIncrement event raised by the VolcanoAuction contract.
type VolcanoAuctionUpdateMinBidIncrement struct {
	MinBidIncrement *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinBidIncrement is a free log retrieval operation binding the contract event 0x489b0441344cbdcb036bee4857de51567f580e9747166f76b581be803ca45fcb.
//
// Solidity: event UpdateMinBidIncrement(uint256 minBidIncrement)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterUpdateMinBidIncrement(opts *bind.FilterOpts) (*VolcanoAuctionUpdateMinBidIncrementIterator, error) {

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "UpdateMinBidIncrement")
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionUpdateMinBidIncrementIterator{contract: _VolcanoAuction.contract, event: "UpdateMinBidIncrement", logs: logs, sub: sub}, nil
}

// WatchUpdateMinBidIncrement is a free log subscription operation binding the contract event 0x489b0441344cbdcb036bee4857de51567f580e9747166f76b581be803ca45fcb.
//
// Solidity: event UpdateMinBidIncrement(uint256 minBidIncrement)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchUpdateMinBidIncrement(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionUpdateMinBidIncrement) (event.Subscription, error) {

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "UpdateMinBidIncrement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionUpdateMinBidIncrement)
				if err := _VolcanoAuction.contract.UnpackLog(event, "UpdateMinBidIncrement", log); err != nil {
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

// ParseUpdateMinBidIncrement is a log parse operation binding the contract event 0x489b0441344cbdcb036bee4857de51567f580e9747166f76b581be803ca45fcb.
//
// Solidity: event UpdateMinBidIncrement(uint256 minBidIncrement)
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseUpdateMinBidIncrement(log types.Log) (*VolcanoAuctionUpdateMinBidIncrement, error) {
	event := new(VolcanoAuctionUpdateMinBidIncrement)
	if err := _VolcanoAuction.contract.UnpackLog(event, "UpdateMinBidIncrement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionUpdatePlatformFeeIterator is returned from FilterUpdatePlatformFee and is used to iterate over the raw logs and unpacked data for UpdatePlatformFee events raised by the VolcanoAuction contract.
type VolcanoAuctionUpdatePlatformFeeIterator struct {
	Event *VolcanoAuctionUpdatePlatformFee // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionUpdatePlatformFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionUpdatePlatformFee)
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
		it.Event = new(VolcanoAuctionUpdatePlatformFee)
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
func (it *VolcanoAuctionUpdatePlatformFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionUpdatePlatformFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionUpdatePlatformFee represents a UpdatePlatformFee event raised by the VolcanoAuction contract.
type VolcanoAuctionUpdatePlatformFee struct {
	PlatformFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFee is a free log retrieval operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterUpdatePlatformFee(opts *bind.FilterOpts) (*VolcanoAuctionUpdatePlatformFeeIterator, error) {

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionUpdatePlatformFeeIterator{contract: _VolcanoAuction.contract, event: "UpdatePlatformFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFee is a free log subscription operation binding the contract event 0x2644fd26359c107ff7991b6dfc36ce902b334ce4e3891bbecacc5922aa620efa.
//
// Solidity: event UpdatePlatformFee(uint256 platformFee)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchUpdatePlatformFee(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionUpdatePlatformFee) (event.Subscription, error) {

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionUpdatePlatformFee)
				if err := _VolcanoAuction.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
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
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseUpdatePlatformFee(log types.Log) (*VolcanoAuctionUpdatePlatformFee, error) {
	event := new(VolcanoAuctionUpdatePlatformFee)
	if err := _VolcanoAuction.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionUpdatePlatformFeeRecipientIterator is returned from FilterUpdatePlatformFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdatePlatformFeeRecipient events raised by the VolcanoAuction contract.
type VolcanoAuctionUpdatePlatformFeeRecipientIterator struct {
	Event *VolcanoAuctionUpdatePlatformFeeRecipient // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionUpdatePlatformFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionUpdatePlatformFeeRecipient)
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
		it.Event = new(VolcanoAuctionUpdatePlatformFeeRecipient)
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
func (it *VolcanoAuctionUpdatePlatformFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionUpdatePlatformFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionUpdatePlatformFeeRecipient represents a UpdatePlatformFeeRecipient event raised by the VolcanoAuction contract.
type VolcanoAuctionUpdatePlatformFeeRecipient struct {
	PlatformFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFeeRecipient is a free log retrieval operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterUpdatePlatformFeeRecipient(opts *bind.FilterOpts) (*VolcanoAuctionUpdatePlatformFeeRecipientIterator, error) {

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionUpdatePlatformFeeRecipientIterator{contract: _VolcanoAuction.contract, event: "UpdatePlatformFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFeeRecipient is a free log subscription operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchUpdatePlatformFeeRecipient(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionUpdatePlatformFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionUpdatePlatformFeeRecipient)
				if err := _VolcanoAuction.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
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
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseUpdatePlatformFeeRecipient(log types.Log) (*VolcanoAuctionUpdatePlatformFeeRecipient, error) {
	event := new(VolcanoAuctionUpdatePlatformFeeRecipient)
	if err := _VolcanoAuction.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the VolcanoAuction contract.
type VolcanoAuctionUpgradedIterator struct {
	Event *VolcanoAuctionUpgraded // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionUpgraded)
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
		it.Event = new(VolcanoAuctionUpgraded)
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
func (it *VolcanoAuctionUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionUpgraded represents a Upgraded event raised by the VolcanoAuction contract.
type VolcanoAuctionUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VolcanoAuctionUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionUpgradedIterator{contract: _VolcanoAuction.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionUpgraded)
				if err := _VolcanoAuction.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseUpgraded(log types.Log) (*VolcanoAuctionUpgraded, error) {
	event := new(VolcanoAuctionUpgraded)
	if err := _VolcanoAuction.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoAuctionVolcanoAuctionContractDeployedIterator is returned from FilterVolcanoAuctionContractDeployed and is used to iterate over the raw logs and unpacked data for VolcanoAuctionContractDeployed events raised by the VolcanoAuction contract.
type VolcanoAuctionVolcanoAuctionContractDeployedIterator struct {
	Event *VolcanoAuctionVolcanoAuctionContractDeployed // Event containing the contract specifics and raw log

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
func (it *VolcanoAuctionVolcanoAuctionContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoAuctionVolcanoAuctionContractDeployed)
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
		it.Event = new(VolcanoAuctionVolcanoAuctionContractDeployed)
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
func (it *VolcanoAuctionVolcanoAuctionContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoAuctionVolcanoAuctionContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoAuctionVolcanoAuctionContractDeployed represents a VolcanoAuctionContractDeployed event raised by the VolcanoAuction contract.
type VolcanoAuctionVolcanoAuctionContractDeployed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVolcanoAuctionContractDeployed is a free log retrieval operation binding the contract event 0xc915c60f988f2a651fe4d21eaf755271bc215c26f2230dcece96d0253d29ee3b.
//
// Solidity: event VolcanoAuctionContractDeployed()
func (_VolcanoAuction *VolcanoAuctionFilterer) FilterVolcanoAuctionContractDeployed(opts *bind.FilterOpts) (*VolcanoAuctionVolcanoAuctionContractDeployedIterator, error) {

	logs, sub, err := _VolcanoAuction.contract.FilterLogs(opts, "VolcanoAuctionContractDeployed")
	if err != nil {
		return nil, err
	}
	return &VolcanoAuctionVolcanoAuctionContractDeployedIterator{contract: _VolcanoAuction.contract, event: "VolcanoAuctionContractDeployed", logs: logs, sub: sub}, nil
}

// WatchVolcanoAuctionContractDeployed is a free log subscription operation binding the contract event 0xc915c60f988f2a651fe4d21eaf755271bc215c26f2230dcece96d0253d29ee3b.
//
// Solidity: event VolcanoAuctionContractDeployed()
func (_VolcanoAuction *VolcanoAuctionFilterer) WatchVolcanoAuctionContractDeployed(opts *bind.WatchOpts, sink chan<- *VolcanoAuctionVolcanoAuctionContractDeployed) (event.Subscription, error) {

	logs, sub, err := _VolcanoAuction.contract.WatchLogs(opts, "VolcanoAuctionContractDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoAuctionVolcanoAuctionContractDeployed)
				if err := _VolcanoAuction.contract.UnpackLog(event, "VolcanoAuctionContractDeployed", log); err != nil {
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

// ParseVolcanoAuctionContractDeployed is a log parse operation binding the contract event 0xc915c60f988f2a651fe4d21eaf755271bc215c26f2230dcece96d0253d29ee3b.
//
// Solidity: event VolcanoAuctionContractDeployed()
func (_VolcanoAuction *VolcanoAuctionFilterer) ParseVolcanoAuctionContractDeployed(log types.Log) (*VolcanoAuctionVolcanoAuctionContractDeployed, error) {
	event := new(VolcanoAuctionVolcanoAuctionContractDeployed)
	if err := _VolcanoAuction.contract.UnpackLog(event, "VolcanoAuctionContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
