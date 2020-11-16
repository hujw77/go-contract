// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MicroPaymentChannelABI is the input ABI used to generate the binding from.
const MicroPaymentChannelABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"claimTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expiration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newExpiration\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sender\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MicroPaymentChannelFuncSigs maps the 4-byte function signature to its string representation.
var MicroPaymentChannelFuncSigs = map[string]string{
	"0e1da6c3": "claimTimeout()",
	"415ffba7": "close(uint256,bytes)",
	"4665096d": "expiration()",
	"9714378c": "extend(uint256)",
	"66d003ac": "recipient()",
	"67e404ce": "sender()",
}

// MicroPaymentChannelBin is the compiled bytecode used for deploying new contracts.
var MicroPaymentChannelBin = "0x60806040526040516104883803806104888339818101604052604081101561002657600080fd5b508051602090910151600080546001600160a01b03199081163317909155600180546001600160a01b0390941693909116929092179091554201600255610416806100726000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630e1da6c314610067578063415ffba7146100715780634665096d1461011e57806366d003ac1461013857806367e404ce1461015c5780639714378c14610164575b600080fd5b61006f610181565b005b61006f6004803603604081101561008757600080fd5b813591908101906040810160208201356401000000008111156100a957600080fd5b8201836020820111156100bb57600080fd5b803590602001918460018302840111640100000000831117156100dd57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061019e945050505050565b610126610211565b60408051918252519081900360200190f35b610140610217565b604080516001600160a01b039092168252519081900360200190f35b610140610226565b61006f6004803603602081101561017a57600080fd5b5035610235565b60025442101561019057600080fd5b6000546001600160a01b0316ff5b6001546001600160a01b031633146101b557600080fd5b6101bf828261025f565b6101c857600080fd5b6001546040516001600160a01b039091169083156108fc029084906000818181858888f19350505050158015610202573d6000803e3d6000fd5b506000546001600160a01b0316ff5b60025481565b6001546001600160a01b031681565b6000546001600160a01b031681565b6000546001600160a01b0316331461024c57600080fd5b600254811161025a57600080fd5b600255565b6000806102ae308560405160200180836001600160a01b03166001600160a01b031660601b815260140182815260200192505050604051602081830303815290604052805190602001206102d9565b6000549091506001600160a01b03166102c7828561032a565b6001600160a01b031614949350505050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a333200000000602080830191909152603c8083019490945282518083039094018452605c909101909152815191012090565b600080600080610339856103b1565b92509250925060018684848460405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa15801561039c573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b600080600083516041146103c457600080fd5b5050506020810151604082015160609092015160001a9290919056fea2646970667358221220d7b5d410f9785815e84b850b81c3b5e35da08fca72f90ab5f120cfe6135763ea64736f6c63430006090033"

// DeployMicroPaymentChannel deploys a new Ethereum contract, binding an instance of MicroPaymentChannel to it.
func DeployMicroPaymentChannel(auth *bind.TransactOpts, backend bind.ContractBackend, _recipient common.Address, duration *big.Int) (common.Address, *types.Transaction, *MicroPaymentChannel, error) {
	parsed, err := abi.JSON(strings.NewReader(MicroPaymentChannelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MicroPaymentChannelBin), backend, _recipient, duration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MicroPaymentChannel{MicroPaymentChannelCaller: MicroPaymentChannelCaller{contract: contract}, MicroPaymentChannelTransactor: MicroPaymentChannelTransactor{contract: contract}, MicroPaymentChannelFilterer: MicroPaymentChannelFilterer{contract: contract}}, nil
}

// MicroPaymentChannel is an auto generated Go binding around an Ethereum contract.
type MicroPaymentChannel struct {
	MicroPaymentChannelCaller     // Read-only binding to the contract
	MicroPaymentChannelTransactor // Write-only binding to the contract
	MicroPaymentChannelFilterer   // Log filterer for contract events
}

// MicroPaymentChannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type MicroPaymentChannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MicroPaymentChannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MicroPaymentChannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MicroPaymentChannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MicroPaymentChannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MicroPaymentChannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MicroPaymentChannelSession struct {
	Contract     *MicroPaymentChannel // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MicroPaymentChannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MicroPaymentChannelCallerSession struct {
	Contract *MicroPaymentChannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// MicroPaymentChannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MicroPaymentChannelTransactorSession struct {
	Contract     *MicroPaymentChannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MicroPaymentChannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type MicroPaymentChannelRaw struct {
	Contract *MicroPaymentChannel // Generic contract binding to access the raw methods on
}

// MicroPaymentChannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MicroPaymentChannelCallerRaw struct {
	Contract *MicroPaymentChannelCaller // Generic read-only contract binding to access the raw methods on
}

// MicroPaymentChannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MicroPaymentChannelTransactorRaw struct {
	Contract *MicroPaymentChannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMicroPaymentChannel creates a new instance of MicroPaymentChannel, bound to a specific deployed contract.
func NewMicroPaymentChannel(address common.Address, backend bind.ContractBackend) (*MicroPaymentChannel, error) {
	contract, err := bindMicroPaymentChannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MicroPaymentChannel{MicroPaymentChannelCaller: MicroPaymentChannelCaller{contract: contract}, MicroPaymentChannelTransactor: MicroPaymentChannelTransactor{contract: contract}, MicroPaymentChannelFilterer: MicroPaymentChannelFilterer{contract: contract}}, nil
}

// NewMicroPaymentChannelCaller creates a new read-only instance of MicroPaymentChannel, bound to a specific deployed contract.
func NewMicroPaymentChannelCaller(address common.Address, caller bind.ContractCaller) (*MicroPaymentChannelCaller, error) {
	contract, err := bindMicroPaymentChannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MicroPaymentChannelCaller{contract: contract}, nil
}

// NewMicroPaymentChannelTransactor creates a new write-only instance of MicroPaymentChannel, bound to a specific deployed contract.
func NewMicroPaymentChannelTransactor(address common.Address, transactor bind.ContractTransactor) (*MicroPaymentChannelTransactor, error) {
	contract, err := bindMicroPaymentChannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MicroPaymentChannelTransactor{contract: contract}, nil
}

// NewMicroPaymentChannelFilterer creates a new log filterer instance of MicroPaymentChannel, bound to a specific deployed contract.
func NewMicroPaymentChannelFilterer(address common.Address, filterer bind.ContractFilterer) (*MicroPaymentChannelFilterer, error) {
	contract, err := bindMicroPaymentChannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MicroPaymentChannelFilterer{contract: contract}, nil
}

// bindMicroPaymentChannel binds a generic wrapper to an already deployed contract.
func bindMicroPaymentChannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MicroPaymentChannelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MicroPaymentChannel *MicroPaymentChannelRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MicroPaymentChannel.Contract.MicroPaymentChannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MicroPaymentChannel *MicroPaymentChannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MicroPaymentChannel.Contract.MicroPaymentChannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MicroPaymentChannel *MicroPaymentChannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MicroPaymentChannel.Contract.MicroPaymentChannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MicroPaymentChannel *MicroPaymentChannelCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MicroPaymentChannel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MicroPaymentChannel *MicroPaymentChannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MicroPaymentChannel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MicroPaymentChannel *MicroPaymentChannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MicroPaymentChannel.Contract.contract.Transact(opts, method, params...)
}

// Expiration is a free data retrieval call binding the contract method 0x4665096d.
//
// Solidity: function expiration() view returns(uint256)
func (_MicroPaymentChannel *MicroPaymentChannelCaller) Expiration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MicroPaymentChannel.contract.Call(opts, out, "expiration")
	return *ret0, err
}

// Expiration is a free data retrieval call binding the contract method 0x4665096d.
//
// Solidity: function expiration() view returns(uint256)
func (_MicroPaymentChannel *MicroPaymentChannelSession) Expiration() (*big.Int, error) {
	return _MicroPaymentChannel.Contract.Expiration(&_MicroPaymentChannel.CallOpts)
}

// Expiration is a free data retrieval call binding the contract method 0x4665096d.
//
// Solidity: function expiration() view returns(uint256)
func (_MicroPaymentChannel *MicroPaymentChannelCallerSession) Expiration() (*big.Int, error) {
	return _MicroPaymentChannel.Contract.Expiration(&_MicroPaymentChannel.CallOpts)
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_MicroPaymentChannel *MicroPaymentChannelCaller) Recipient(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MicroPaymentChannel.contract.Call(opts, out, "recipient")
	return *ret0, err
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_MicroPaymentChannel *MicroPaymentChannelSession) Recipient() (common.Address, error) {
	return _MicroPaymentChannel.Contract.Recipient(&_MicroPaymentChannel.CallOpts)
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_MicroPaymentChannel *MicroPaymentChannelCallerSession) Recipient() (common.Address, error) {
	return _MicroPaymentChannel.Contract.Recipient(&_MicroPaymentChannel.CallOpts)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_MicroPaymentChannel *MicroPaymentChannelCaller) Sender(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MicroPaymentChannel.contract.Call(opts, out, "sender")
	return *ret0, err
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_MicroPaymentChannel *MicroPaymentChannelSession) Sender() (common.Address, error) {
	return _MicroPaymentChannel.Contract.Sender(&_MicroPaymentChannel.CallOpts)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address)
func (_MicroPaymentChannel *MicroPaymentChannelCallerSession) Sender() (common.Address, error) {
	return _MicroPaymentChannel.Contract.Sender(&_MicroPaymentChannel.CallOpts)
}

// ClaimTimeout is a paid mutator transaction binding the contract method 0x0e1da6c3.
//
// Solidity: function claimTimeout() returns()
func (_MicroPaymentChannel *MicroPaymentChannelTransactor) ClaimTimeout(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MicroPaymentChannel.contract.Transact(opts, "claimTimeout")
}

// ClaimTimeout is a paid mutator transaction binding the contract method 0x0e1da6c3.
//
// Solidity: function claimTimeout() returns()
func (_MicroPaymentChannel *MicroPaymentChannelSession) ClaimTimeout() (*types.Transaction, error) {
	return _MicroPaymentChannel.Contract.ClaimTimeout(&_MicroPaymentChannel.TransactOpts)
}

// ClaimTimeout is a paid mutator transaction binding the contract method 0x0e1da6c3.
//
// Solidity: function claimTimeout() returns()
func (_MicroPaymentChannel *MicroPaymentChannelTransactorSession) ClaimTimeout() (*types.Transaction, error) {
	return _MicroPaymentChannel.Contract.ClaimTimeout(&_MicroPaymentChannel.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x415ffba7.
//
// Solidity: function close(uint256 amount, bytes signature) returns()
func (_MicroPaymentChannel *MicroPaymentChannelTransactor) Close(opts *bind.TransactOpts, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _MicroPaymentChannel.contract.Transact(opts, "close", amount, signature)
}

// Close is a paid mutator transaction binding the contract method 0x415ffba7.
//
// Solidity: function close(uint256 amount, bytes signature) returns()
func (_MicroPaymentChannel *MicroPaymentChannelSession) Close(amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _MicroPaymentChannel.Contract.Close(&_MicroPaymentChannel.TransactOpts, amount, signature)
}

// Close is a paid mutator transaction binding the contract method 0x415ffba7.
//
// Solidity: function close(uint256 amount, bytes signature) returns()
func (_MicroPaymentChannel *MicroPaymentChannelTransactorSession) Close(amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _MicroPaymentChannel.Contract.Close(&_MicroPaymentChannel.TransactOpts, amount, signature)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 newExpiration) returns()
func (_MicroPaymentChannel *MicroPaymentChannelTransactor) Extend(opts *bind.TransactOpts, newExpiration *big.Int) (*types.Transaction, error) {
	return _MicroPaymentChannel.contract.Transact(opts, "extend", newExpiration)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 newExpiration) returns()
func (_MicroPaymentChannel *MicroPaymentChannelSession) Extend(newExpiration *big.Int) (*types.Transaction, error) {
	return _MicroPaymentChannel.Contract.Extend(&_MicroPaymentChannel.TransactOpts, newExpiration)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 newExpiration) returns()
func (_MicroPaymentChannel *MicroPaymentChannelTransactorSession) Extend(newExpiration *big.Int) (*types.Transaction, error) {
	return _MicroPaymentChannel.Contract.Extend(&_MicroPaymentChannel.TransactOpts, newExpiration)
}
