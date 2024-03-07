// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todo

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

// TodoTask is an auto generated low-level Go binding around an user-defined struct.
type TodoTask struct {
	Content string
	Status  bool
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b503360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506110928061005d5f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80630f560cd7146100645780634cc82215146100825780638da5cb5b1461009e5780639507d39a146100bc578063b0c8f9dc146100ec578063f745630f14610108575b5f80fd5b61006c610124565b604051610079919061089e565b60405180910390f35b61009c60048036038101906100979190610902565b610280565b005b6100a66103da565b6040516100b3919061096c565b60405180910390f35b6100d660048036038101906100d19190610902565b6103ff565b6040516100e391906109bf565b60405180910390f35b61010660048036038101906101019190610b0b565b610539565b005b610122600480360381019061011d9190610b52565b610607565b005b606060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461017e575f80fd5b5f805480602002602001604051908101604052809291908181526020015f905b82821015610277578382905f5260205f2090600202016040518060400160405290815f820180546101ce90610bd9565b80601f01602080910402602001604051908101604052809291908181526020018280546101fa90610bd9565b80156102455780601f1061021c57610100808354040283529160200191610245565b820191905f5260205f20905b81548152906001019060200180831161022857829003601f168201915b50505050508152602001600182015f9054906101000a900460ff1615151515815250508152602001906001019061019e565b50505050905090565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102d8575f80fd5b5f8190505b60015f805490506102ee9190610c36565b81101561038d575f6001826103039190610c69565b8154811061031457610313610c9c565b5b905f5260205f2090600202015f828154811061033357610332610c9c565b5b905f5260205f2090600202015f8201815f0190816103519190610e7b565b50600182015f9054906101000a900460ff16816001015f6101000a81548160ff02191690831515021790555090505080806001019150506102dd565b505f80548061039f5761039e610f60565b5b600190038181905f5260205f2090600202015f8082015f6103c09190610692565b600182015f6101000a81549060ff02191690555050905550565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6104076106cf565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461045f575f80fd5b5f828154811061047257610471610c9c565b5b905f5260205f2090600202016040518060400160405290815f8201805461049890610bd9565b80601f01602080910402602001604051908101604052809291908181526020018280546104c490610bd9565b801561050f5780601f106104e65761010080835404028352916020019161050f565b820191905f5260205f20905b8154815290600101906020018083116104f257829003601f168201915b50505050508152602001600182015f9054906101000a900460ff1615151515815250509050919050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610591575f80fd5b5f60405180604001604052808381526020015f1515815250908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f0190816105e29190610f8d565b506020820151816001015f6101000a81548160ff021916908315150217905550505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461065f575f80fd5b805f838154811061067357610672610c9c565b5b905f5260205f2090600202015f01908161068d9190610f8d565b505050565b50805461069e90610bd9565b5f825580601f106106af57506106cc565b601f0160209004905f5260205f20908101906106cb91906106ea565b5b50565b6040518060400160405280606081526020015f151581525090565b5b80821115610701575f815f9055506001016106eb565b5090565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561076557808201518184015260208101905061074a565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61078a8261072e565b6107948185610738565b93506107a4818560208601610748565b6107ad81610770565b840191505092915050565b5f8115159050919050565b6107cc816107b8565b82525050565b5f604083015f8301518482035f8601526107ec8282610780565b915050602083015161080160208601826107c3565b508091505092915050565b5f61081783836107d2565b905092915050565b5f602082019050919050565b5f61083582610705565b61083f818561070f565b9350836020820285016108518561071f565b805f5b8581101561088c578484038952815161086d858261080c565b94506108788361081f565b925060208a01995050600181019050610854565b50829750879550505050505092915050565b5f6020820190508181035f8301526108b6818461082b565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b6108e1816108cf565b81146108eb575f80fd5b50565b5f813590506108fc816108d8565b92915050565b5f60208284031215610917576109166108c7565b5b5f610924848285016108ee565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6109568261092d565b9050919050565b6109668161094c565b82525050565b5f60208201905061097f5f83018461095d565b92915050565b5f604083015f8301518482035f86015261099f8282610780565b91505060208301516109b460208601826107c3565b508091505092915050565b5f6020820190508181035f8301526109d78184610985565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610a1d82610770565b810181811067ffffffffffffffff82111715610a3c57610a3b6109e7565b5b80604052505050565b5f610a4e6108be565b9050610a5a8282610a14565b919050565b5f67ffffffffffffffff821115610a7957610a786109e7565b5b610a8282610770565b9050602081019050919050565b828183375f83830152505050565b5f610aaf610aaa84610a5f565b610a45565b905082815260208101848484011115610acb57610aca6109e3565b5b610ad6848285610a8f565b509392505050565b5f82601f830112610af257610af16109df565b5b8135610b02848260208601610a9d565b91505092915050565b5f60208284031215610b2057610b1f6108c7565b5b5f82013567ffffffffffffffff811115610b3d57610b3c6108cb565b5b610b4984828501610ade565b91505092915050565b5f8060408385031215610b6857610b676108c7565b5b5f610b75858286016108ee565b925050602083013567ffffffffffffffff811115610b9657610b956108cb565b5b610ba285828601610ade565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610bf057607f821691505b602082108103610c0357610c02610bac565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610c40826108cf565b9150610c4b836108cf565b9250828203905081811115610c6357610c62610c09565b5b92915050565b5f610c73826108cf565b9150610c7e836108cf565b9250828201905080821115610c9657610c95610c09565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81549050610cd781610bd9565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610d3a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610cff565b610d448683610cff565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610d7f610d7a610d75846108cf565b610d5c565b6108cf565b9050919050565b5f819050919050565b610d9883610d65565b610dac610da482610d86565b848454610d0b565b825550505050565b5f90565b610dc0610db4565b610dcb818484610d8f565b505050565b5b81811015610dee57610de35f82610db8565b600181019050610dd1565b5050565b601f821115610e3357610e0481610cde565b610e0d84610cf0565b81016020851015610e1c578190505b610e30610e2885610cf0565b830182610dd0565b50505b505050565b5f82821c905092915050565b5f610e535f1984600802610e38565b1980831691505092915050565b5f610e6b8383610e44565b9150826002028217905092915050565b818103610e89575050610f5e565b610e9282610cc9565b67ffffffffffffffff811115610eab57610eaa6109e7565b5b610eb58254610bd9565b610ec0828285610df2565b5f601f831160018114610eed575f8415610edb578287015490505b610ee58582610e60565b865550610f57565b601f198416610efb87610cde565b9650610f0686610cde565b5f5b82811015610f2d57848901548255600182019150600185019450602081019050610f08565b86831015610f4a5784890154610f46601f891682610e44565b8355505b6001600288020188555050505b5050505050505b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b610f968261072e565b67ffffffffffffffff811115610faf57610fae6109e7565b5b610fb98254610bd9565b610fc4828285610df2565b5f60209050601f831160018114610ff5575f8415610fe3578287015190505b610fed8582610e60565b865550611054565b601f19841661100386610cde565b5f5b8281101561102a57848901518255600182019150602085019450602081019050611005565b868310156110475784890151611043601f891682610e44565b8355505b6001600288020188555050505b50505050505056fea264697066735822122049472d7476062965d1f0f558ccc7211a81ab868e6407b22345b512736413ead464736f6c63430008180033",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 id) view returns((string,bool))
func (_Todo *TodoCaller) Get(opts *bind.CallOpts, id *big.Int) (TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "get", id)

	if err != nil {
		return *new(TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new(TodoTask)).(*TodoTask)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 id) view returns((string,bool))
func (_Todo *TodoSession) Get(id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, id)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 id) view returns((string,bool))
func (_Todo *TodoCallerSession) Get(id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, id)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCaller) List(opts *bind.CallOpts) ([]TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTask)).(*[]TodoTask)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCallerSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCallerSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoTransactor) Add(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "add", _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoSession) Add(_content string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _content) returns()
func (_Todo *TodoTransactorSession) Add(_content string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _content)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactor) Remove(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "remove", _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactorSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactor) Update(opts *bind.TransactOpts, _id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "update", _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactorSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}
