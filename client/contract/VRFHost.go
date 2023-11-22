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

// VRFHostRound is an auto generated low-level Go binding around an user-defined struct.
type VRFHostRound struct {
	Proposer         common.Address
	RandomNumber     *big.Int
	RandomNumberHash [32]byte
	State            uint8
	BlockHeight      *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getCurrentRandomNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPreviousRandomNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"name\":\"getRound\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"randomNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"randomNumberHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumVRFHost.RoundState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"internalType\":\"structVRFHost.Round\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"setRandomNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"verifyProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"verifySignature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040526001805463ffffffff19168117905534801561001e575f80fd5b506040805160a0810182525f81527ff4af66a743b631375dbe319fa457aa07e08b0401f9a822f6c797722938143db260208201527f4d5f7dd6652bc469c41cc53d0c515d0254469f469272633b5b34ad32e2a10a0c918101919091526060810160028152436020918201525f808052815281517fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b0319166001600160a01b03909216919091178155908201517fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb65560408201517fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb75560608201517fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb8805460ff1916600183600281111561015c5761015c610173565b021790555060808201518160040155905050610187565b634e487b7160e01b5f52602160045260245ffd5b61072c806101945f395ff3fe608060405234801561000f575f80fd5b5060043610610085575f3560e01c80636e7b0fed116100585780636e7b0fed146100df5780636ffd39e8146100fd578063969578691461011d578063fb2c76cb14610148575f80fd5b8063209652551461008957806334d5bae31461009f578063380507b3146100c257806347e40553146100d7575b5f80fd5b60015b6040519081526020015b60405180910390f35b6100b26100ad36600461056d565b610150565b6040519015158152602001610096565b6100d56100d036600461056d565b6101b5565b005b6100d561035f565b6001805463ffffffff165f908152602081905260409020015461008c565b61011061010b36600461059d565b6103eb565b60405161009691906105d4565b61013061012b366004610632565b6104a2565b6040516001600160a01b039091168152602001610096565b61008c610523565b600180545f91829161018e918391829161016f9163ffffffff1661067e565b63ffffffff1681526020019081526020015f20600201548686866104a2565b90506001600160a01b038116156101a95760019150506101ae565b5f9150505b9392505050565b600180545f916101f291839182916101d3919063ffffffff1661067e565b63ffffffff1681526020019081526020015f20600201548585856104a2565b90506001600160a01b03811633146102445760405162461bcd60e51b815260206004820152601060248201526f57726f6e67207369676e61747572652160801b60448201526064015b60405180910390fd5b6001805463ffffffff165f908152602081905260409020608085811b9085901c1791600382015460ff16600281111561027f5761027f6105c0565b036102e85780600101548211156102e85760405162461bcd60e51b815260206004820152602760248201527f57726f6e67207369676e6174757265206f72206e756d626572206973206e6f746044820152662076616c69642160c81b606482015260840161023b565b608084811c9086901b17600182018190556040805160208101929092520160408051808303601f190181529190528051602090910120600282015560038101805460ff1916600117905580546001600160a01b03939093166001600160a01b03199093169290921782555043600490910155505050565b60015463ffffffff165f9081526020819052604090206004810154610386906005906106a2565b4311801561039d575080546001600160a01b031633145b6103a5575f80fd5b60038101805460ff191660021790556001805463ffffffff16905f6103c9836106bb565b91906101000a81548163ffffffff021916908363ffffffff1602179055505050565b6040805160a0810182525f8082526020820181905291810182905260608101829052608081019190915263ffffffff82165f9081526020818152604091829020825160a08101845281546001600160a01b0316815260018201549281019290925260028082015493830193909352600381015491929091606084019160ff9091169081111561047c5761047c6105c0565b600281111561048d5761048d6105c0565b81526020016004820154815250509050919050565b5f601b8460ff1610156104bd576104ba601b856106dd565b93505b604080515f8082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa15801561050e573d5f803e3d5ffd5b5050604051601f190151979650505050505050565b600180545f918291829161053c9163ffffffff1661067e565b63ffffffff1681526020019081526020015f2060010154905090565b803560ff81168114610568575f80fd5b919050565b5f805f6060848603121561057f575f80fd5b61058884610558565b95602085013595506040909401359392505050565b5f602082840312156105ad575f80fd5b813563ffffffff811681146101ae575f80fd5b634e487b7160e01b5f52602160045260245ffd5b81516001600160a01b031681526020808301519082015260408083015190820152606082015160a08201906003811061061b57634e487b7160e01b5f52602160045260245ffd5b806060840152506080830151608083015292915050565b5f805f8060808587031215610645575f80fd5b8435935061065560208601610558565b93969395505050506040820135916060013590565b634e487b7160e01b5f52601160045260245ffd5b63ffffffff82811682821603908082111561069b5761069b61066a565b5092915050565b808201808211156106b5576106b561066a565b92915050565b5f63ffffffff8083168181036106d3576106d361066a565b6001019392505050565b60ff81811683821601908111156106b5576106b561066a56fea264697066735822122089ae75720f71ce42c11f8536de41d491aeead30a8c2ed68bac060163dff54ce264736f6c63430008170033",
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

// GetRound is a free data retrieval call binding the contract method 0x6ffd39e8.
//
// Solidity: function getRound(uint32 id) view returns((address,uint256,bytes32,uint8,uint256))
func (_Contract *ContractCaller) GetRound(opts *bind.CallOpts, id uint32) (VRFHostRound, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRound", id)

	if err != nil {
		return *new(VRFHostRound), err
	}

	out0 := *abi.ConvertType(out[0], new(VRFHostRound)).(*VRFHostRound)

	return out0, err

}

// GetRound is a free data retrieval call binding the contract method 0x6ffd39e8.
//
// Solidity: function getRound(uint32 id) view returns((address,uint256,bytes32,uint8,uint256))
func (_Contract *ContractSession) GetRound(id uint32) (VRFHostRound, error) {
	return _Contract.Contract.GetRound(&_Contract.CallOpts, id)
}

// GetRound is a free data retrieval call binding the contract method 0x6ffd39e8.
//
// Solidity: function getRound(uint32 id) view returns((address,uint256,bytes32,uint8,uint256))
func (_Contract *ContractCallerSession) GetRound(id uint32) (VRFHostRound, error) {
	return _Contract.Contract.GetRound(&_Contract.CallOpts, id)
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

// NextRound is a paid mutator transaction binding the contract method 0x47e40553.
//
// Solidity: function nextRound() returns()
func (_Contract *ContractTransactor) NextRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "nextRound")
}

// NextRound is a paid mutator transaction binding the contract method 0x47e40553.
//
// Solidity: function nextRound() returns()
func (_Contract *ContractSession) NextRound() (*types.Transaction, error) {
	return _Contract.Contract.NextRound(&_Contract.TransactOpts)
}

// NextRound is a paid mutator transaction binding the contract method 0x47e40553.
//
// Solidity: function nextRound() returns()
func (_Contract *ContractTransactorSession) NextRound() (*types.Transaction, error) {
	return _Contract.Contract.NextRound(&_Contract.TransactOpts)
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
