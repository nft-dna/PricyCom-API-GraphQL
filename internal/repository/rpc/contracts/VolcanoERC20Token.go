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

// VolcanoERC720TokenMetaData contains all meta data concerning the VolcanoERC720Token contract.
var VolcanoERC720TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initialAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_capAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintBlocksFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_routerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"BlockMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintBlock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBlocksAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBlocksFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBlocksMaxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBlocksSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routerAddress\",\"outputs\":[{\"internalType\":\"contractUniswapRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"updateContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// VolcanoERC720TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use VolcanoERC720TokenMetaData.ABI instead.
var VolcanoERC720TokenABI = VolcanoERC720TokenMetaData.ABI

// VolcanoERC720Token is an auto generated Go binding around an Ethereum contract.
type VolcanoERC720Token struct {
	VolcanoERC720TokenCaller     // Read-only binding to the contract
	VolcanoERC720TokenTransactor // Write-only binding to the contract
	VolcanoERC720TokenFilterer   // Log filterer for contract events
}

// VolcanoERC720TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type VolcanoERC720TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC720TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VolcanoERC720TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC720TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolcanoERC720TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC720TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolcanoERC720TokenSession struct {
	Contract     *VolcanoERC720Token // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VolcanoERC720TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolcanoERC720TokenCallerSession struct {
	Contract *VolcanoERC720TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// VolcanoERC720TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolcanoERC720TokenTransactorSession struct {
	Contract     *VolcanoERC720TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// VolcanoERC720TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type VolcanoERC720TokenRaw struct {
	Contract *VolcanoERC720Token // Generic contract binding to access the raw methods on
}

// VolcanoERC720TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolcanoERC720TokenCallerRaw struct {
	Contract *VolcanoERC720TokenCaller // Generic read-only contract binding to access the raw methods on
}

// VolcanoERC720TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolcanoERC720TokenTransactorRaw struct {
	Contract *VolcanoERC720TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolcanoERC720Token creates a new instance of VolcanoERC720Token, bound to a specific deployed contract.
func NewVolcanoERC720Token(address common.Address, backend bind.ContractBackend) (*VolcanoERC720Token, error) {
	contract, err := bindVolcanoERC720Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720Token{VolcanoERC720TokenCaller: VolcanoERC720TokenCaller{contract: contract}, VolcanoERC720TokenTransactor: VolcanoERC720TokenTransactor{contract: contract}, VolcanoERC720TokenFilterer: VolcanoERC720TokenFilterer{contract: contract}}, nil
}

// NewVolcanoERC720TokenCaller creates a new read-only instance of VolcanoERC720Token, bound to a specific deployed contract.
func NewVolcanoERC720TokenCaller(address common.Address, caller bind.ContractCaller) (*VolcanoERC720TokenCaller, error) {
	contract, err := bindVolcanoERC720Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720TokenCaller{contract: contract}, nil
}

// NewVolcanoERC720TokenTransactor creates a new write-only instance of VolcanoERC720Token, bound to a specific deployed contract.
func NewVolcanoERC720TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*VolcanoERC720TokenTransactor, error) {
	contract, err := bindVolcanoERC720Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720TokenTransactor{contract: contract}, nil
}

// NewVolcanoERC720TokenFilterer creates a new log filterer instance of VolcanoERC720Token, bound to a specific deployed contract.
func NewVolcanoERC720TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*VolcanoERC720TokenFilterer, error) {
	contract, err := bindVolcanoERC720Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720TokenFilterer{contract: contract}, nil
}

// bindVolcanoERC720Token binds a generic wrapper to an already deployed contract.
func bindVolcanoERC720Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VolcanoERC720TokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC720Token *VolcanoERC720TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC720Token.Contract.VolcanoERC720TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC720Token *VolcanoERC720TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.VolcanoERC720TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC720Token *VolcanoERC720TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.VolcanoERC720TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC720Token *VolcanoERC720TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC720Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _VolcanoERC720Token.Contract.DOMAINSEPARATOR(&_VolcanoERC720Token.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _VolcanoERC720Token.Contract.DOMAINSEPARATOR(&_VolcanoERC720Token.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VolcanoERC720Token.Contract.Allowance(&_VolcanoERC720Token.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VolcanoERC720Token.Contract.Allowance(&_VolcanoERC720Token.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VolcanoERC720Token.Contract.BalanceOf(&_VolcanoERC720Token.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VolcanoERC720Token.Contract.BalanceOf(&_VolcanoERC720Token.CallOpts, account)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) Cap() (*big.Int, error) {
	return _VolcanoERC720Token.Contract.Cap(&_VolcanoERC720Token.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) Cap() (*big.Int, error) {
	return _VolcanoERC720Token.Contract.Cap(&_VolcanoERC720Token.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) ContractURI() (string, error) {
	return _VolcanoERC720Token.Contract.ContractURI(&_VolcanoERC720Token.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) ContractURI() (string, error) {
	return _VolcanoERC720Token.Contract.ContractURI(&_VolcanoERC720Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) Decimals() (uint8, error) {
	return _VolcanoERC720Token.Contract.Decimals(&_VolcanoERC720Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) Decimals() (uint8, error) {
	return _VolcanoERC720Token.Contract.Decimals(&_VolcanoERC720Token.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _VolcanoERC720Token.Contract.Eip712Domain(&_VolcanoERC720Token.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _VolcanoERC720Token.Contract.Eip712Domain(&_VolcanoERC720Token.CallOpts)
}

// InitialReserves is a free data retrieval call binding the contract method 0x87890455.
//
// Solidity: function initialReserves() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) InitialReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "initialReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialReserves is a free data retrieval call binding the contract method 0x87890455.
//
// Solidity: function initialReserves() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) InitialReserves() (*big.Int, error) {
	return _VolcanoERC720Token.Contract.InitialReserves(&_VolcanoERC720Token.CallOpts)
}

// InitialReserves is a free data retrieval call binding the contract method 0x87890455.
//
// Solidity: function initialReserves() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) InitialReserves() (*big.Int, error) {
	return _VolcanoERC720Token.Contract.InitialReserves(&_VolcanoERC720Token.CallOpts)
}

// MintBlocksAmount is a free data retrieval call binding the contract method 0x92b52b8e.
//
// Solidity: function mintBlocksAmount() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) MintBlocksAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "mintBlocksAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintBlocksAmount is a free data retrieval call binding the contract method 0x92b52b8e.
//
// Solidity: function mintBlocksAmount() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) MintBlocksAmount() (*big.Int, error) {
	return _VolcanoERC720Token.Contract.MintBlocksAmount(&_VolcanoERC720Token.CallOpts)
}

// MintBlocksAmount is a free data retrieval call binding the contract method 0x92b52b8e.
//
// Solidity: function mintBlocksAmount() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) MintBlocksAmount() (*big.Int, error) {
	return _VolcanoERC720Token.Contract.MintBlocksAmount(&_VolcanoERC720Token.CallOpts)
}

// MintBlocksFee is a free data retrieval call binding the contract method 0x63441ee2.
//
// Solidity: function mintBlocksFee() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) MintBlocksFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "mintBlocksFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintBlocksFee is a free data retrieval call binding the contract method 0x63441ee2.
//
// Solidity: function mintBlocksFee() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) MintBlocksFee() (*big.Int, error) {
	return _VolcanoERC720Token.Contract.MintBlocksFee(&_VolcanoERC720Token.CallOpts)
}

// MintBlocksFee is a free data retrieval call binding the contract method 0x63441ee2.
//
// Solidity: function mintBlocksFee() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) MintBlocksFee() (*big.Int, error) {
	return _VolcanoERC720Token.Contract.MintBlocksFee(&_VolcanoERC720Token.CallOpts)
}

// MintBlocksMaxSupply is a free data retrieval call binding the contract method 0xbbb1fd51.
//
// Solidity: function mintBlocksMaxSupply() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) MintBlocksMaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "mintBlocksMaxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintBlocksMaxSupply is a free data retrieval call binding the contract method 0xbbb1fd51.
//
// Solidity: function mintBlocksMaxSupply() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) MintBlocksMaxSupply() (*big.Int, error) {
	return _VolcanoERC720Token.Contract.MintBlocksMaxSupply(&_VolcanoERC720Token.CallOpts)
}

// MintBlocksMaxSupply is a free data retrieval call binding the contract method 0xbbb1fd51.
//
// Solidity: function mintBlocksMaxSupply() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) MintBlocksMaxSupply() (*big.Int, error) {
	return _VolcanoERC720Token.Contract.MintBlocksMaxSupply(&_VolcanoERC720Token.CallOpts)
}

// MintBlocksSupply is a free data retrieval call binding the contract method 0xcf82f26d.
//
// Solidity: function mintBlocksSupply() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) MintBlocksSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "mintBlocksSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintBlocksSupply is a free data retrieval call binding the contract method 0xcf82f26d.
//
// Solidity: function mintBlocksSupply() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) MintBlocksSupply() (*big.Int, error) {
	return _VolcanoERC720Token.Contract.MintBlocksSupply(&_VolcanoERC720Token.CallOpts)
}

// MintBlocksSupply is a free data retrieval call binding the contract method 0xcf82f26d.
//
// Solidity: function mintBlocksSupply() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) MintBlocksSupply() (*big.Int, error) {
	return _VolcanoERC720Token.Contract.MintBlocksSupply(&_VolcanoERC720Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) Name() (string, error) {
	return _VolcanoERC720Token.Contract.Name(&_VolcanoERC720Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) Name() (string, error) {
	return _VolcanoERC720Token.Contract.Name(&_VolcanoERC720Token.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) Nonces(owner common.Address) (*big.Int, error) {
	return _VolcanoERC720Token.Contract.Nonces(&_VolcanoERC720Token.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _VolcanoERC720Token.Contract.Nonces(&_VolcanoERC720Token.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) Owner() (common.Address, error) {
	return _VolcanoERC720Token.Contract.Owner(&_VolcanoERC720Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) Owner() (common.Address, error) {
	return _VolcanoERC720Token.Contract.Owner(&_VolcanoERC720Token.CallOpts)
}

// RouterAddress is a free data retrieval call binding the contract method 0x3268cc56.
//
// Solidity: function routerAddress() view returns(address)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) RouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "routerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RouterAddress is a free data retrieval call binding the contract method 0x3268cc56.
//
// Solidity: function routerAddress() view returns(address)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) RouterAddress() (common.Address, error) {
	return _VolcanoERC720Token.Contract.RouterAddress(&_VolcanoERC720Token.CallOpts)
}

// RouterAddress is a free data retrieval call binding the contract method 0x3268cc56.
//
// Solidity: function routerAddress() view returns(address)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) RouterAddress() (common.Address, error) {
	return _VolcanoERC720Token.Contract.RouterAddress(&_VolcanoERC720Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) Symbol() (string, error) {
	return _VolcanoERC720Token.Contract.Symbol(&_VolcanoERC720Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) Symbol() (string, error) {
	return _VolcanoERC720Token.Contract.Symbol(&_VolcanoERC720Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC720Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) TotalSupply() (*big.Int, error) {
	return _VolcanoERC720Token.Contract.TotalSupply(&_VolcanoERC720Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VolcanoERC720Token *VolcanoERC720TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _VolcanoERC720Token.Contract.TotalSupply(&_VolcanoERC720Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VolcanoERC720Token *VolcanoERC720TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.Approve(&_VolcanoERC720Token.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.Approve(&_VolcanoERC720Token.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.Burn(&_VolcanoERC720Token.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.Burn(&_VolcanoERC720Token.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.BurnFrom(&_VolcanoERC720Token.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.BurnFrom(&_VolcanoERC720Token.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VolcanoERC720Token *VolcanoERC720TokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.DecreaseAllowance(&_VolcanoERC720Token.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.DecreaseAllowance(&_VolcanoERC720Token.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VolcanoERC720Token *VolcanoERC720TokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.IncreaseAllowance(&_VolcanoERC720Token.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.IncreaseAllowance(&_VolcanoERC720Token.TransactOpts, spender, addedValue)
}

// MintBlock is a paid mutator transaction binding the contract method 0xf35abe09.
//
// Solidity: function mintBlock(address to) payable returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactor) MintBlock(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Token.contract.Transact(opts, "mintBlock", to)
}

// MintBlock is a paid mutator transaction binding the contract method 0xf35abe09.
//
// Solidity: function mintBlock(address to) payable returns()
func (_VolcanoERC720Token *VolcanoERC720TokenSession) MintBlock(to common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.MintBlock(&_VolcanoERC720Token.TransactOpts, to)
}

// MintBlock is a paid mutator transaction binding the contract method 0xf35abe09.
//
// Solidity: function mintBlock(address to) payable returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorSession) MintBlock(to common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.MintBlock(&_VolcanoERC720Token.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VolcanoERC720Token.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.Permit(&_VolcanoERC720Token.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.Permit(&_VolcanoERC720Token.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC720Token.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC720Token *VolcanoERC720TokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.RenounceOwnership(&_VolcanoERC720Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.RenounceOwnership(&_VolcanoERC720Token.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VolcanoERC720Token *VolcanoERC720TokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.Transfer(&_VolcanoERC720Token.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.Transfer(&_VolcanoERC720Token.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VolcanoERC720Token *VolcanoERC720TokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VolcanoERC720Token *VolcanoERC720TokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.TransferFrom(&_VolcanoERC720Token.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.TransferFrom(&_VolcanoERC720Token.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.TransferOwnership(&_VolcanoERC720Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.TransferOwnership(&_VolcanoERC720Token.TransactOpts, newOwner)
}

// UpdateContractURI is a paid mutator transaction binding the contract method 0x7e5b1e24.
//
// Solidity: function updateContractURI(string _uri) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactor) UpdateContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _VolcanoERC720Token.contract.Transact(opts, "updateContractURI", _uri)
}

// UpdateContractURI is a paid mutator transaction binding the contract method 0x7e5b1e24.
//
// Solidity: function updateContractURI(string _uri) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenSession) UpdateContractURI(_uri string) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.UpdateContractURI(&_VolcanoERC720Token.TransactOpts, _uri)
}

// UpdateContractURI is a paid mutator transaction binding the contract method 0x7e5b1e24.
//
// Solidity: function updateContractURI(string _uri) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorSession) UpdateContractURI(_uri string) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.UpdateContractURI(&_VolcanoERC720Token.TransactOpts, _uri)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactor) WithdrawETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.contract.Transact(opts, "withdrawETH", amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.WithdrawETH(&_VolcanoERC720Token.TransactOpts, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.WithdrawETH(&_VolcanoERC720Token.TransactOpts, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x315a095d.
//
// Solidity: function withdrawTokens(uint256 amount) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactor) WithdrawTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.contract.Transact(opts, "withdrawTokens", amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x315a095d.
//
// Solidity: function withdrawTokens(uint256 amount) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenSession) WithdrawTokens(amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.WithdrawTokens(&_VolcanoERC720Token.TransactOpts, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x315a095d.
//
// Solidity: function withdrawTokens(uint256 amount) returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorSession) WithdrawTokens(amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.WithdrawTokens(&_VolcanoERC720Token.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC720Token.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VolcanoERC720Token *VolcanoERC720TokenSession) Receive() (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.Receive(&_VolcanoERC720Token.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VolcanoERC720Token *VolcanoERC720TokenTransactorSession) Receive() (*types.Transaction, error) {
	return _VolcanoERC720Token.Contract.Receive(&_VolcanoERC720Token.TransactOpts)
}

// VolcanoERC720TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VolcanoERC720Token contract.
type VolcanoERC720TokenApprovalIterator struct {
	Event *VolcanoERC720TokenApproval // Event containing the contract specifics and raw log

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
func (it *VolcanoERC720TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC720TokenApproval)
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
		it.Event = new(VolcanoERC720TokenApproval)
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
func (it *VolcanoERC720TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC720TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC720TokenApproval represents a Approval event raised by the VolcanoERC720Token contract.
type VolcanoERC720TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VolcanoERC720Token *VolcanoERC720TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VolcanoERC720TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VolcanoERC720Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720TokenApprovalIterator{contract: _VolcanoERC720Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VolcanoERC720Token *VolcanoERC720TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VolcanoERC720TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VolcanoERC720Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC720TokenApproval)
				if err := _VolcanoERC720Token.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VolcanoERC720Token *VolcanoERC720TokenFilterer) ParseApproval(log types.Log) (*VolcanoERC720TokenApproval, error) {
	event := new(VolcanoERC720TokenApproval)
	if err := _VolcanoERC720Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC720TokenBlockMintedIterator is returned from FilterBlockMinted and is used to iterate over the raw logs and unpacked data for BlockMinted events raised by the VolcanoERC720Token contract.
type VolcanoERC720TokenBlockMintedIterator struct {
	Event *VolcanoERC720TokenBlockMinted // Event containing the contract specifics and raw log

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
func (it *VolcanoERC720TokenBlockMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC720TokenBlockMinted)
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
		it.Event = new(VolcanoERC720TokenBlockMinted)
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
func (it *VolcanoERC720TokenBlockMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC720TokenBlockMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC720TokenBlockMinted represents a BlockMinted event raised by the VolcanoERC720Token contract.
type VolcanoERC720TokenBlockMinted struct {
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlockMinted is a free log retrieval operation binding the contract event 0x0e9a1ec107d573764d20047a2ac52b16c549f7366e25e13cc1c8437d3fe98b5d.
//
// Solidity: event BlockMinted(address receiver)
func (_VolcanoERC720Token *VolcanoERC720TokenFilterer) FilterBlockMinted(opts *bind.FilterOpts) (*VolcanoERC720TokenBlockMintedIterator, error) {

	logs, sub, err := _VolcanoERC720Token.contract.FilterLogs(opts, "BlockMinted")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720TokenBlockMintedIterator{contract: _VolcanoERC720Token.contract, event: "BlockMinted", logs: logs, sub: sub}, nil
}

// WatchBlockMinted is a free log subscription operation binding the contract event 0x0e9a1ec107d573764d20047a2ac52b16c549f7366e25e13cc1c8437d3fe98b5d.
//
// Solidity: event BlockMinted(address receiver)
func (_VolcanoERC720Token *VolcanoERC720TokenFilterer) WatchBlockMinted(opts *bind.WatchOpts, sink chan<- *VolcanoERC720TokenBlockMinted) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC720Token.contract.WatchLogs(opts, "BlockMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC720TokenBlockMinted)
				if err := _VolcanoERC720Token.contract.UnpackLog(event, "BlockMinted", log); err != nil {
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

// ParseBlockMinted is a log parse operation binding the contract event 0x0e9a1ec107d573764d20047a2ac52b16c549f7366e25e13cc1c8437d3fe98b5d.
//
// Solidity: event BlockMinted(address receiver)
func (_VolcanoERC720Token *VolcanoERC720TokenFilterer) ParseBlockMinted(log types.Log) (*VolcanoERC720TokenBlockMinted, error) {
	event := new(VolcanoERC720TokenBlockMinted)
	if err := _VolcanoERC720Token.contract.UnpackLog(event, "BlockMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC720TokenEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the VolcanoERC720Token contract.
type VolcanoERC720TokenEIP712DomainChangedIterator struct {
	Event *VolcanoERC720TokenEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *VolcanoERC720TokenEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC720TokenEIP712DomainChanged)
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
		it.Event = new(VolcanoERC720TokenEIP712DomainChanged)
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
func (it *VolcanoERC720TokenEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC720TokenEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC720TokenEIP712DomainChanged represents a EIP712DomainChanged event raised by the VolcanoERC720Token contract.
type VolcanoERC720TokenEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_VolcanoERC720Token *VolcanoERC720TokenFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*VolcanoERC720TokenEIP712DomainChangedIterator, error) {

	logs, sub, err := _VolcanoERC720Token.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720TokenEIP712DomainChangedIterator{contract: _VolcanoERC720Token.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_VolcanoERC720Token *VolcanoERC720TokenFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *VolcanoERC720TokenEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC720Token.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC720TokenEIP712DomainChanged)
				if err := _VolcanoERC720Token.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_VolcanoERC720Token *VolcanoERC720TokenFilterer) ParseEIP712DomainChanged(log types.Log) (*VolcanoERC720TokenEIP712DomainChanged, error) {
	event := new(VolcanoERC720TokenEIP712DomainChanged)
	if err := _VolcanoERC720Token.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC720TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VolcanoERC720Token contract.
type VolcanoERC720TokenOwnershipTransferredIterator struct {
	Event *VolcanoERC720TokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VolcanoERC720TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC720TokenOwnershipTransferred)
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
		it.Event = new(VolcanoERC720TokenOwnershipTransferred)
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
func (it *VolcanoERC720TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC720TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC720TokenOwnershipTransferred represents a OwnershipTransferred event raised by the VolcanoERC720Token contract.
type VolcanoERC720TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC720Token *VolcanoERC720TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VolcanoERC720TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC720Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720TokenOwnershipTransferredIterator{contract: _VolcanoERC720Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC720Token *VolcanoERC720TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VolcanoERC720TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC720Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC720TokenOwnershipTransferred)
				if err := _VolcanoERC720Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VolcanoERC720Token *VolcanoERC720TokenFilterer) ParseOwnershipTransferred(log types.Log) (*VolcanoERC720TokenOwnershipTransferred, error) {
	event := new(VolcanoERC720TokenOwnershipTransferred)
	if err := _VolcanoERC720Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC720TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VolcanoERC720Token contract.
type VolcanoERC720TokenTransferIterator struct {
	Event *VolcanoERC720TokenTransfer // Event containing the contract specifics and raw log

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
func (it *VolcanoERC720TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC720TokenTransfer)
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
		it.Event = new(VolcanoERC720TokenTransfer)
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
func (it *VolcanoERC720TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC720TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC720TokenTransfer represents a Transfer event raised by the VolcanoERC720Token contract.
type VolcanoERC720TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VolcanoERC720Token *VolcanoERC720TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VolcanoERC720TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VolcanoERC720Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC720TokenTransferIterator{contract: _VolcanoERC720Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VolcanoERC720Token *VolcanoERC720TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VolcanoERC720TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VolcanoERC720Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC720TokenTransfer)
				if err := _VolcanoERC720Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VolcanoERC720Token *VolcanoERC720TokenFilterer) ParseTransfer(log types.Log) (*VolcanoERC720TokenTransfer, error) {
	event := new(VolcanoERC720TokenTransfer)
	if err := _VolcanoERC720Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
