// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getCurrentRandomNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPreviousRandomNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"setRandomNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"verifyProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"verifySignature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040525f805463ffffffff191681557ff4af66a743b631375dbe319fa457aa07e08b0401f9a822f6c797722938143db26001557f4d5f7dd6652bc469c41cc53d0c515d0254469f469272633b5b34ad32e2a10a0c6002556003556004805460ff19169055348015610070575f80fd5b5061038b8061007e5f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c8063209652551461006457806334d5bae31461007a578063380507b31461009d5780636e7b0fed146100b257806396957869146100ba578063fb2c76cb146100e5575b5f80fd5b60015b6040519081526020015b60405180910390f35b61008d6100883660046102c2565b6100ed565b6040519015158152602001610071565b6100b06100ab3660046102c2565b610124565b005b600354610067565b6100cd6100c83660046102f2565b6101ff565b6040516001600160a01b039091168152602001610071565b600154610067565b5f806100fd6002548686866101ff565b90506001600160a01b0381161561011857600191505061011d565b5f9150505b9392505050565b5f6101308484846100ed565b9050806101775760405162461bcd60e51b815260206004820152601060248201526f57726f6e67207369676e61747572652160801b60448201526064015b60405180910390fd5b600454608084811b9084901c179060ff16156101ef576003548111156101ef5760405162461bcd60e51b815260206004820152602760248201527f57726f6e67207369676e6174757265206f72206e756d626572206973206e6f746044820152662076616c69642160c81b606482015260840161016e565b5050608091821b911c1760035550565b5f808560405160200161021491815260200190565b604051602081830303815290604052805190602001209050601b8560ff16101561024657610243601b8661032a565b94505b604080515f8082526020820180845284905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610297573d5f803e3d5ffd5b5050604051601f19015198975050505050505050565b803560ff811681146102bd575f80fd5b919050565b5f805f606084860312156102d4575f80fd5b6102dd846102ad565b95602085013595506040909401359392505050565b5f805f8060808587031215610305575f80fd5b84359350610315602086016102ad565b93969395505050506040820135916060013590565b60ff818116838216019081111561034f57634e487b7160e01b5f52601160045260245ffd5b9291505056fea264697066735822122080acda7b798afd481eb15bae6de7fcd702b37b227e70c1ac7d5187f697a5d08f64736f6c63430008170033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentRandomNumber is a free data retrieval call binding the contract method 0x6e7b0fed.
//
// Solidity: function getCurrentRandomNumber() view returns(uint256)
func (_Contract *ContractCaller) GetCurrentRandomNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCurrentRandomNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRandomNumber is a free data retrieval call binding the contract method 0x6e7b0fed.
//
// Solidity: function getCurrentRandomNumber() view returns(uint256)
func (_Contract *ContractSession) GetCurrentRandomNumber() (*big.Int, error) {
	return _Contract.Contract.GetCurrentRandomNumber(&_Contract.CallOpts)
}

// GetCurrentRandomNumber is a free data retrieval call binding the contract method 0x6e7b0fed.
//
// Solidity: function getCurrentRandomNumber() view returns(uint256)
func (_Contract *ContractCallerSession) GetCurrentRandomNumber() (*big.Int, error) {
	return _Contract.Contract.GetCurrentRandomNumber(&_Contract.CallOpts)
}

// GetPreviousRandomNumber is a free data retrieval call binding the contract method 0xfb2c76cb.
//
// Solidity: function getPreviousRandomNumber() view returns(uint256)
func (_Contract *ContractCaller) GetPreviousRandomNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPreviousRandomNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreviousRandomNumber is a free data retrieval call binding the contract method 0xfb2c76cb.
//
// Solidity: function getPreviousRandomNumber() view returns(uint256)
func (_Contract *ContractSession) GetPreviousRandomNumber() (*big.Int, error) {
	return _Contract.Contract.GetPreviousRandomNumber(&_Contract.CallOpts)
}

// GetPreviousRandomNumber is a free data retrieval call binding the contract method 0xfb2c76cb.
//
// Solidity: function getPreviousRandomNumber() view returns(uint256)
func (_Contract *ContractCallerSession) GetPreviousRandomNumber() (*big.Int, error) {
	return _Contract.Contract.GetPreviousRandomNumber(&_Contract.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() pure returns(uint256)
func (_Contract *ContractCaller) GetValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() pure returns(uint256)
func (_Contract *ContractSession) GetValue() (*big.Int, error) {
	return _Contract.Contract.GetValue(&_Contract.CallOpts)
}

// GetValue is a free data retrieval call binding the contract method 0x20965255.
//
// Solidity: function getValue() pure returns(uint256)
func (_Contract *ContractCallerSession) GetValue() (*big.Int, error) {
	return _Contract.Contract.GetValue(&_Contract.CallOpts)
}

// VerifyProposal is a free data retrieval call binding the contract method 0x34d5bae3.
//
// Solidity: function verifyProposal(uint8 _v, bytes32 _r, bytes32 _s) view returns(bool)
func (_Contract *ContractCaller) VerifyProposal(opts *bind.CallOpts, _v uint8, _r [32]byte, _s [32]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verifyProposal", _v, _r, _s)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProposal is a free data retrieval call binding the contract method 0x34d5bae3.
//
// Solidity: function verifyProposal(uint8 _v, bytes32 _r, bytes32 _s) view returns(bool)
func (_Contract *ContractSession) VerifyProposal(_v uint8, _r [32]byte, _s [32]byte) (bool, error) {
	return _Contract.Contract.VerifyProposal(&_Contract.CallOpts, _v, _r, _s)
}

// VerifyProposal is a free data retrieval call binding the contract method 0x34d5bae3.
//
// Solidity: function verifyProposal(uint8 _v, bytes32 _r, bytes32 _s) view returns(bool)
func (_Contract *ContractCallerSession) VerifyProposal(_v uint8, _r [32]byte, _s [32]byte) (bool, error) {
	return _Contract.Contract.VerifyProposal(&_Contract.CallOpts, _v, _r, _s)
}

// VerifySignature is a free data retrieval call binding the contract method 0x96957869.
//
// Solidity: function verifySignature(bytes32 message, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_Contract *ContractCaller) VerifySignature(opts *bind.CallOpts, message [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verifySignature", message, _v, _r, _s)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifySignature is a free data retrieval call binding the contract method 0x96957869.
//
// Solidity: function verifySignature(bytes32 message, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_Contract *ContractSession) VerifySignature(message [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	return _Contract.Contract.VerifySignature(&_Contract.CallOpts, message, _v, _r, _s)
}

// VerifySignature is a free data retrieval call binding the contract method 0x96957869.
//
// Solidity: function verifySignature(bytes32 message, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_Contract *ContractCallerSession) VerifySignature(message [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	return _Contract.Contract.VerifySignature(&_Contract.CallOpts, message, _v, _r, _s)
}

// SetRandomNumber is a paid mutator transaction binding the contract method 0x380507b3.
//
// Solidity: function setRandomNumber(uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Contract *ContractTransactor) SetRandomNumber(opts *bind.TransactOpts, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setRandomNumber", _v, _r, _s)
}

// SetRandomNumber is a paid mutator transaction binding the contract method 0x380507b3.
//
// Solidity: function setRandomNumber(uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Contract *ContractSession) SetRandomNumber(_v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetRandomNumber(&_Contract.TransactOpts, _v, _r, _s)
}

// SetRandomNumber is a paid mutator transaction binding the contract method 0x380507b3.
//
// Solidity: function setRandomNumber(uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Contract *ContractTransactorSession) SetRandomNumber(_v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetRandomNumber(&_Contract.TransactOpts, _v, _r, _s)
}
