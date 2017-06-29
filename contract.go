// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getNumApprovers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_approver\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"uint256\"}],\"name\":\"addApprover\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getApprover\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSpecification\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getApprovalsStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256[5]\"},{\"name\":\"\",\"type\":\"uint256[5]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getApplicationName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_approver\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"changeStatus\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getApproveRC\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_requestor\",\"type\":\"address\"},{\"name\":\"_size\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_numApprovers\",\"type\":\"uint256\"},{\"name\":\"_reqApprovers\",\"type\":\"uint256[5]\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"numApprovers\",\"type\":\"uint256\"}],\"name\":\"ApplicationCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ApplicationAlreadyApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"TooManyApprovers\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"AddedApprover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"HasApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ApproverNotFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"GetApproverByNumber\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"decision\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"ApproverByNumber\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"GetNumberApprovers\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"WrongApplicationSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"WrongIndex\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_type\",\"type\":\"uint256\"}],\"name\":\"WrongApproverType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"WrongApproverAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"GetStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Killed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"AddingApprover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"CanAddApprover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"title\",\"type\":\"uint256\"}],\"name\":\"TitleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ApproverAlreadyExists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ApprovalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"CheckingApprover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"FoundApprover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ApplicationApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CheckingApprovals\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approvals\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"approvers\",\"type\":\"uint256\"}],\"name\":\"MissingApprovals\",\"type\":\"event\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
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
	contract, err := bindContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

// GetApplicationName is a free data retrieval call binding the contract method 0xc0a7772d.
//
// Solidity: function getApplicationName() constant returns(string)
func (_Contract *ContractCaller) GetApplicationName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getApplicationName")
	return *ret0, err
}

// GetApplicationName is a free data retrieval call binding the contract method 0xc0a7772d.
//
// Solidity: function getApplicationName() constant returns(string)
func (_Contract *ContractSession) GetApplicationName() (string, error) {
	return _Contract.Contract.GetApplicationName(&_Contract.CallOpts)
}

// GetApplicationName is a free data retrieval call binding the contract method 0xc0a7772d.
//
// Solidity: function getApplicationName() constant returns(string)
func (_Contract *ContractCallerSession) GetApplicationName() (string, error) {
	return _Contract.Contract.GetApplicationName(&_Contract.CallOpts)
}

// GetApprovalsStatus is a free data retrieval call binding the contract method 0x9addff57.
//
// Solidity: function getApprovalsStatus() constant returns(uint256, uint256[5], uint256[5])
func (_Contract *ContractCaller) GetApprovalsStatus(opts *bind.CallOpts) (*big.Int, [5]*big.Int, [5]*big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([5]*big.Int)
		ret2 = new([5]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Contract.contract.Call(opts, out, "getApprovalsStatus")
	return *ret0, *ret1, *ret2, err
}

// GetApprovalsStatus is a free data retrieval call binding the contract method 0x9addff57.
//
// Solidity: function getApprovalsStatus() constant returns(uint256, uint256[5], uint256[5])
func (_Contract *ContractSession) GetApprovalsStatus() (*big.Int, [5]*big.Int, [5]*big.Int, error) {
	return _Contract.Contract.GetApprovalsStatus(&_Contract.CallOpts)
}

// GetApprovalsStatus is a free data retrieval call binding the contract method 0x9addff57.
//
// Solidity: function getApprovalsStatus() constant returns(uint256, uint256[5], uint256[5])
func (_Contract *ContractCallerSession) GetApprovalsStatus() (*big.Int, [5]*big.Int, [5]*big.Int, error) {
	return _Contract.Contract.GetApprovalsStatus(&_Contract.CallOpts)
}

// GetApproveRC is a free data retrieval call binding the contract method 0xf39ac26c.
//
// Solidity: function getApproveRC() constant returns(int256)
func (_Contract *ContractCaller) GetApproveRC(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getApproveRC")
	return *ret0, err
}

// GetApproveRC is a free data retrieval call binding the contract method 0xf39ac26c.
//
// Solidity: function getApproveRC() constant returns(int256)
func (_Contract *ContractSession) GetApproveRC() (*big.Int, error) {
	return _Contract.Contract.GetApproveRC(&_Contract.CallOpts)
}

// GetApproveRC is a free data retrieval call binding the contract method 0xf39ac26c.
//
// Solidity: function getApproveRC() constant returns(int256)
func (_Contract *ContractCallerSession) GetApproveRC() (*big.Int, error) {
	return _Contract.Contract.GetApproveRC(&_Contract.CallOpts)
}

// GetApprover is a free data retrieval call binding the contract method 0x7157e493.
//
// Solidity: function getApprover(i uint256) constant returns(address, uint256, uint256)
func (_Contract *ContractCaller) GetApprover(opts *bind.CallOpts, i *big.Int) (common.Address, *big.Int, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Contract.contract.Call(opts, out, "getApprover", i)
	return *ret0, *ret1, *ret2, err
}

// GetApprover is a free data retrieval call binding the contract method 0x7157e493.
//
// Solidity: function getApprover(i uint256) constant returns(address, uint256, uint256)
func (_Contract *ContractSession) GetApprover(i *big.Int) (common.Address, *big.Int, *big.Int, error) {
	return _Contract.Contract.GetApprover(&_Contract.CallOpts, i)
}

// GetApprover is a free data retrieval call binding the contract method 0x7157e493.
//
// Solidity: function getApprover(i uint256) constant returns(address, uint256, uint256)
func (_Contract *ContractCallerSession) GetApprover(i *big.Int) (common.Address, *big.Int, *big.Int, error) {
	return _Contract.Contract.GetApprover(&_Contract.CallOpts, i)
}

// GetIndex is a free data retrieval call binding the contract method 0x81045ead.
//
// Solidity: function getIndex() constant returns(uint256)
func (_Contract *ContractCaller) GetIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getIndex")
	return *ret0, err
}

// GetIndex is a free data retrieval call binding the contract method 0x81045ead.
//
// Solidity: function getIndex() constant returns(uint256)
func (_Contract *ContractSession) GetIndex() (*big.Int, error) {
	return _Contract.Contract.GetIndex(&_Contract.CallOpts)
}

// GetIndex is a free data retrieval call binding the contract method 0x81045ead.
//
// Solidity: function getIndex() constant returns(uint256)
func (_Contract *ContractCallerSession) GetIndex() (*big.Int, error) {
	return _Contract.Contract.GetIndex(&_Contract.CallOpts)
}

// GetNumApprovers is a free data retrieval call binding the contract method 0x13abc4f4.
//
// Solidity: function getNumApprovers() constant returns(uint256)
func (_Contract *ContractCaller) GetNumApprovers(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getNumApprovers")
	return *ret0, err
}

// GetNumApprovers is a free data retrieval call binding the contract method 0x13abc4f4.
//
// Solidity: function getNumApprovers() constant returns(uint256)
func (_Contract *ContractSession) GetNumApprovers() (*big.Int, error) {
	return _Contract.Contract.GetNumApprovers(&_Contract.CallOpts)
}

// GetNumApprovers is a free data retrieval call binding the contract method 0x13abc4f4.
//
// Solidity: function getNumApprovers() constant returns(uint256)
func (_Contract *ContractCallerSession) GetNumApprovers() (*big.Int, error) {
	return _Contract.Contract.GetNumApprovers(&_Contract.CallOpts)
}

// GetRequestor is a free data retrieval call binding the contract method 0xa3f5b1bf.
//
// Solidity: function getRequestor() constant returns(address)
func (_Contract *ContractCaller) GetRequestor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getRequestor")
	return *ret0, err
}

// GetRequestor is a free data retrieval call binding the contract method 0xa3f5b1bf.
//
// Solidity: function getRequestor() constant returns(address)
func (_Contract *ContractSession) GetRequestor() (common.Address, error) {
	return _Contract.Contract.GetRequestor(&_Contract.CallOpts)
}

// GetRequestor is a free data retrieval call binding the contract method 0xa3f5b1bf.
//
// Solidity: function getRequestor() constant returns(address)
func (_Contract *ContractCallerSession) GetRequestor() (common.Address, error) {
	return _Contract.Contract.GetRequestor(&_Contract.CallOpts)
}

// GetSpecification is a free data retrieval call binding the contract method 0x72e75458.
//
// Solidity: function getSpecification() constant returns(uint256, uint256, uint256, uint256)
func (_Contract *ContractCaller) GetSpecification(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Contract.contract.Call(opts, out, "getSpecification")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetSpecification is a free data retrieval call binding the contract method 0x72e75458.
//
// Solidity: function getSpecification() constant returns(uint256, uint256, uint256, uint256)
func (_Contract *ContractSession) GetSpecification() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.GetSpecification(&_Contract.CallOpts)
}

// GetSpecification is a free data retrieval call binding the contract method 0x72e75458.
//
// Solidity: function getSpecification() constant returns(uint256, uint256, uint256, uint256)
func (_Contract *ContractCallerSession) GetSpecification() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Contract.Contract.GetSpecification(&_Contract.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(uint256)
func (_Contract *ContractCaller) GetStatus(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getStatus")
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(uint256)
func (_Contract *ContractSession) GetStatus() (*big.Int, error) {
	return _Contract.Contract.GetStatus(&_Contract.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() constant returns(uint256)
func (_Contract *ContractCallerSession) GetStatus() (*big.Int, error) {
	return _Contract.Contract.GetStatus(&_Contract.CallOpts)
}

// AddApprover is a paid mutator transaction binding the contract method 0x56c98930.
//
// Solidity: function addApprover(_approver address, _role uint256) returns()
func (_Contract *ContractTransactor) AddApprover(opts *bind.TransactOpts, _approver common.Address, _role *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addApprover", _approver, _role)
}

// AddApprover is a paid mutator transaction binding the contract method 0x56c98930.
//
// Solidity: function addApprover(_approver address, _role uint256) returns()
func (_Contract *ContractSession) AddApprover(_approver common.Address, _role *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddApprover(&_Contract.TransactOpts, _approver, _role)
}

// AddApprover is a paid mutator transaction binding the contract method 0x56c98930.
//
// Solidity: function addApprover(_approver address, _role uint256) returns()
func (_Contract *ContractTransactorSession) AddApprover(_approver common.Address, _role *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.AddApprover(&_Contract.TransactOpts, _approver, _role)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(_approver address) returns()
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, _approver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", _approver)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(_approver address) returns()
func (_Contract *ContractSession) Approve(_approver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, _approver)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(_approver address) returns()
func (_Contract *ContractTransactorSession) Approve(_approver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, _approver)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0xe8025d77.
//
// Solidity: function changeStatus(_status uint256) returns()
func (_Contract *ContractTransactor) ChangeStatus(opts *bind.TransactOpts, _status *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "changeStatus", _status)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0xe8025d77.
//
// Solidity: function changeStatus(_status uint256) returns()
func (_Contract *ContractSession) ChangeStatus(_status *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ChangeStatus(&_Contract.TransactOpts, _status)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0xe8025d77.
//
// Solidity: function changeStatus(_status uint256) returns()
func (_Contract *ContractTransactorSession) ChangeStatus(_status *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ChangeStatus(&_Contract.TransactOpts, _status)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Contract *ContractTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Contract *ContractSession) Kill() (*types.Transaction, error) {
	return _Contract.Contract.Kill(&_Contract.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Contract *ContractTransactorSession) Kill() (*types.Transaction, error) {
	return _Contract.Contract.Kill(&_Contract.TransactOpts)
}
