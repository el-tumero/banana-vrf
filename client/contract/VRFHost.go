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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"_id\",\"type\":\"uint32\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"checkStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRandomNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRoundId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPreviousRandomNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"name\":\"getRound\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"randomNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"randomNumberHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumVRFHost.RoundState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"}],\"internalType\":\"structVRFHost.Round\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isOperatorActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRoundLate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"setRandomNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"verifyProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"verifySignature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526001805463ffffffff19168117905534801561001e575f80fd5b506040805160a0810182525f81527ff4af66a743b631375dbe319fa457aa07e08b0401f9a822f6c797722938143db260208201527f4d5f7dd6652bc469c41cc53d0c515d0254469f469272633b5b34ad32e2a10a0c918101919091526060810160028152436020918201525f808052815281517fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb580546001600160a01b0319166001600160a01b03909216919091178155908201517fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb65560408201517fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb75560608201517fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb8805460ff1916600183600281111561015c5761015c61019f565b02179055506080919091015160049091015560015f908152602052437fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e81556101b3565b634e487b7160e01b5f52602160045260245ffd5b610cb6806101c05f395ff3fe6080604052600436106100d9575f3560e01c80636ffd39e81161007c5780639695786911610057578063969578691461022d578063bed9d86114610264578063eb47353314610278578063fb2c76cb14610297575f80fd5b80636ffd39e8146101b957806380cd39e7146101e557806390d96d76146101f9575f80fd5b806347e40553116100b757806347e405531461014f5780635727e25d146101635780635a627dbc146101875780636e7b0fed1461018f575f80fd5b806320965255146100dd57806334d5bae3146100ff578063380507b31461012e575b5f80fd5b3480156100e8575f80fd5b5060015b6040519081526020015b60405180910390f35b34801561010a575f80fd5b5061011e610119366004610aba565b6102ab565b60405190151581526020016100f6565b348015610139575f80fd5b5061014d610148366004610aba565b610316565b005b34801561015a575f80fd5b5061014d6105ae565b34801561016e575f80fd5b5060015460405163ffffffff90911681526020016100f6565b61014d6106db565b34801561019a575f80fd5b506001805463ffffffff165f90815260208190526040902001546100ec565b3480156101c4575f80fd5b506101d86101d3366004610aea565b61076c565b6040516100f69190610b21565b3480156101f0575f80fd5b5061014d610823565b348015610204575f80fd5b506100ec610213366004610b7f565b6001600160a01b03165f9081526002602052604090205490565b348015610238575f80fd5b5061024c610247366004610ba5565b610864565b6040516001600160a01b0390911681526020016100f6565b34801561026f575f80fd5b5061014d6108e5565b348015610283575f80fd5b5061011e610292366004610b7f565b6109d8565b3480156102a2575f80fd5b506100ec610a6a565b600180545f9182916102ef91839182916102ca9163ffffffff16610bf1565b63ffffffff1663ffffffff1681526020019081526020015f2060020154868686610864565b90506001600160a01b0381161561030a57600191505061030f565b5f9150505b9392505050565b600180545f916103599183918291610334919063ffffffff16610bf1565b63ffffffff1663ffffffff1681526020019081526020015f2060020154858585610864565b90506001600160a01b03811633146103ab5760405162461bcd60e51b815260206004820152601060248201526f57726f6e67207369676e61747572652160801b60448201526064015b60405180910390fd5b6001600160a01b0381165f908152600260205260409020546064106104075760405162461bcd60e51b81526020600482015260126024820152714e6f207265717569726564207374616b652160701b60448201526064016103a2565b6001600160a01b0381165f908152600260205260409020600190810154905460039161043c9163ffffffff9182169116610bf1565b63ffffffff1611806104595750600154600563ffffffff90911611155b6104995760405162461bcd60e51b8152602060048201526011602482015270596f75206e65656420746f20776169742160781b60448201526064016103a2565b6001805463ffffffff165f908152602081905260409020608085811b9085901c1791600382015460ff1660028111156104d4576104d4610b0d565b0361053d57806001015482111561053d5760405162461bcd60e51b815260206004820152602760248201527f57726f6e67207369676e6174757265206f72206e756d626572206973206e6f746044820152662076616c69642160c81b60648201526084016103a2565b608084811c9086901b17600182018190556040805160208101929092520160408051808303601f190181529190528051602090910120600282015560038101805460ff1916600117905580546001600160a01b03939093166001600160a01b03199093169290921790915550505050565b60015463ffffffff165f90815260208190526040902060048101546105d590600590610c15565b431180156105ec575080546001600160a01b031633145b8015610611575060015b600382015460ff16600281111561060f5761060f610b0d565b145b61064e5760405162461bcd60e51b815260206004820152600e60248201526d4e6f74207065726d69747465642160901b60448201526064016103a2565b60038101805460ff191660021790556001805463ffffffff16905f61067283610c2e565b82546101009290920a63ffffffff8181021990931691831602179091556001805482165f9081526020819052604080822043600490910155915491519190921692507f7070e98ba33226dd52c8af328b325cc0a0806df2c2cb2edb42708308e6af745d9190a250565b5f341161071b5760405162461bcd60e51b815260206004820152600e60248201526d4e6f74207065726d69747465642160901b60448201526064016103a2565b335f9081526002602052604081208054349290610739908490610c15565b909155505060018054335f908152600260205260409020909101805463ffffffff191663ffffffff909216919091179055565b6040805160a0810182525f8082526020820181905291810182905260608101829052608081019190915263ffffffff82165f9081526020818152604091829020825160a08101845281546001600160a01b0316815260018201549281019290925260028082015493830193909352600381015491929091606084019160ff909116908111156107fd576107fd610b0d565b600281111561080e5761080e610b0d565b81526020016004820154815250509050919050565b60015463ffffffff165f90815260208190526040902061084560056002610c50565b81600401546108549190610c15565b43118015610611575060016105f6565b5f601b8460ff16101561087f5761087c601b85610c67565b93505b604080515f8082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa1580156108d0573d5f803e3d5ffd5b5050604051601f190151979650505050505050565b335f90815260026020526040902054806109395760405162461bcd60e51b81526020600482015260156024820152744e6f2066756e647320746f2077697468647261772160581b60448201526064016103a2565b335f818152600260205260408082208290555190919083908381818185875af1925050503d805f8114610987576040519150601f19603f3d011682016040523d82523d5f602084013e61098c565b606091505b50509050806109d45760405162461bcd60e51b81526020600482015260146024820152732330b4b632b2103a379039b2b7321022ba3432b960611b60448201526064016103a2565b5050565b6001600160a01b0381165f9081526002602090815260408083208151808301909252805480835260019091015463ffffffff1692820192909252906064108015610a5457506020810151600154600391610a379163ffffffff16610bf1565b63ffffffff161180610a545750600154600563ffffffff90911611155b15610a625750600192915050565b505f92915050565b600180545f9182918291610a839163ffffffff16610bf1565b63ffffffff1663ffffffff1681526020019081526020015f2060010154905090565b803560ff81168114610ab5575f80fd5b919050565b5f805f60608486031215610acc575f80fd5b610ad584610aa5565b95602085013595506040909401359392505050565b5f60208284031215610afa575f80fd5b813563ffffffff8116811461030f575f80fd5b634e487b7160e01b5f52602160045260245ffd5b81516001600160a01b031681526020808301519082015260408083015190820152606082015160a082019060038110610b6857634e487b7160e01b5f52602160045260245ffd5b806060840152506080830151608083015292915050565b5f60208284031215610b8f575f80fd5b81356001600160a01b038116811461030f575f80fd5b5f805f8060808587031215610bb8575f80fd5b84359350610bc860208601610aa5565b93969395505050506040820135916060013590565b634e487b7160e01b5f52601160045260245ffd5b63ffffffff828116828216039080821115610c0e57610c0e610bdd565b5092915050565b80820180821115610c2857610c28610bdd565b92915050565b5f63ffffffff808316818103610c4657610c46610bdd565b6001019392505050565b8082028115828204841417610c2857610c28610bdd565b60ff8181168382160190811115610c2857610c28610bdd56fea2646970667358221220092e020cdefb97ee3476985f3e309ebb081e60060a9e5c18ef4b45fe6421092e64736f6c63430008170033",
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

// CheckStake is a free data retrieval call binding the contract method 0x90d96d76.
//
// Solidity: function checkStake(address operator) view returns(uint256)
func (_Contract *ContractCaller) CheckStake(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "checkStake", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckStake is a free data retrieval call binding the contract method 0x90d96d76.
//
// Solidity: function checkStake(address operator) view returns(uint256)
func (_Contract *ContractSession) CheckStake(operator common.Address) (*big.Int, error) {
	return _Contract.Contract.CheckStake(&_Contract.CallOpts, operator)
}

// CheckStake is a free data retrieval call binding the contract method 0x90d96d76.
//
// Solidity: function checkStake(address operator) view returns(uint256)
func (_Contract *ContractCallerSession) CheckStake(operator common.Address) (*big.Int, error) {
	return _Contract.Contract.CheckStake(&_Contract.CallOpts, operator)
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

// GetCurrentRoundId is a free data retrieval call binding the contract method 0x5727e25d.
//
// Solidity: function getCurrentRoundId() view returns(uint32)
func (_Contract *ContractCaller) GetCurrentRoundId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCurrentRoundId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetCurrentRoundId is a free data retrieval call binding the contract method 0x5727e25d.
//
// Solidity: function getCurrentRoundId() view returns(uint32)
func (_Contract *ContractSession) GetCurrentRoundId() (uint32, error) {
	return _Contract.Contract.GetCurrentRoundId(&_Contract.CallOpts)
}

// GetCurrentRoundId is a free data retrieval call binding the contract method 0x5727e25d.
//
// Solidity: function getCurrentRoundId() view returns(uint32)
func (_Contract *ContractCallerSession) GetCurrentRoundId() (uint32, error) {
	return _Contract.Contract.GetCurrentRoundId(&_Contract.CallOpts)
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

// IsOperatorActive is a free data retrieval call binding the contract method 0xeb473533.
//
// Solidity: function isOperatorActive(address addr) view returns(bool)
func (_Contract *ContractCaller) IsOperatorActive(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isOperatorActive", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorActive is a free data retrieval call binding the contract method 0xeb473533.
//
// Solidity: function isOperatorActive(address addr) view returns(bool)
func (_Contract *ContractSession) IsOperatorActive(addr common.Address) (bool, error) {
	return _Contract.Contract.IsOperatorActive(&_Contract.CallOpts, addr)
}

// IsOperatorActive is a free data retrieval call binding the contract method 0xeb473533.
//
// Solidity: function isOperatorActive(address addr) view returns(bool)
func (_Contract *ContractCallerSession) IsOperatorActive(addr common.Address) (bool, error) {
	return _Contract.Contract.IsOperatorActive(&_Contract.CallOpts, addr)
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

// AddStake is a paid mutator transaction binding the contract method 0x5a627dbc.
//
// Solidity: function addStake() payable returns()
func (_Contract *ContractTransactor) AddStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addStake")
}

// AddStake is a paid mutator transaction binding the contract method 0x5a627dbc.
//
// Solidity: function addStake() payable returns()
func (_Contract *ContractSession) AddStake() (*types.Transaction, error) {
	return _Contract.Contract.AddStake(&_Contract.TransactOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x5a627dbc.
//
// Solidity: function addStake() payable returns()
func (_Contract *ContractTransactorSession) AddStake() (*types.Transaction, error) {
	return _Contract.Contract.AddStake(&_Contract.TransactOpts)
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

// NextRoundLate is a paid mutator transaction binding the contract method 0x80cd39e7.
//
// Solidity: function nextRoundLate() returns()
func (_Contract *ContractTransactor) NextRoundLate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "nextRoundLate")
}

// NextRoundLate is a paid mutator transaction binding the contract method 0x80cd39e7.
//
// Solidity: function nextRoundLate() returns()
func (_Contract *ContractSession) NextRoundLate() (*types.Transaction, error) {
	return _Contract.Contract.NextRoundLate(&_Contract.TransactOpts)
}

// NextRoundLate is a paid mutator transaction binding the contract method 0x80cd39e7.
//
// Solidity: function nextRoundLate() returns()
func (_Contract *ContractTransactorSession) NextRoundLate() (*types.Transaction, error) {
	return _Contract.Contract.NextRoundLate(&_Contract.TransactOpts)
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

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Contract *ContractTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Contract *ContractSession) WithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawStake(&_Contract.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Contract *ContractTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawStake(&_Contract.TransactOpts)
}

// ContractNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the Contract contract.
type ContractNewRoundIterator struct {
	Event *ContractNewRound // Event containing the contract specifics and raw log

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
func (it *ContractNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewRound)
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
		it.Event = new(ContractNewRound)
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
func (it *ContractNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewRound represents a NewRound event raised by the Contract contract.
type ContractNewRound struct {
	Id  uint32
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0x7070e98ba33226dd52c8af328b325cc0a0806df2c2cb2edb42708308e6af745d.
//
// Solidity: event NewRound(uint32 indexed _id)
func (_Contract *ContractFilterer) FilterNewRound(opts *bind.FilterOpts, _id []uint32) (*ContractNewRoundIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewRound", _idRule)
	if err != nil {
		return nil, err
	}
	return &ContractNewRoundIterator{contract: _Contract.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0x7070e98ba33226dd52c8af328b325cc0a0806df2c2cb2edb42708308e6af745d.
//
// Solidity: event NewRound(uint32 indexed _id)
func (_Contract *ContractFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *ContractNewRound, _id []uint32) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewRound", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewRound)
				if err := _Contract.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0x7070e98ba33226dd52c8af328b325cc0a0806df2c2cb2edb42708308e6af745d.
//
// Solidity: event NewRound(uint32 indexed _id)
func (_Contract *ContractFilterer) ParseNewRound(log types.Log) (*ContractNewRound, error) {
	event := new(ContractNewRound)
	if err := _Contract.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
