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

// MetaContractABI is the input ABI used to generate the binding from.
const MetaContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"getApplication\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deleteAllApplications\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberApplications\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"removeApplicationByNumber\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"findApplicationByNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMaxApplications\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"applications\",\"outputs\":[{\"name\":\"contractAddress\",\"type\":\"address\"},{\"name\":\"applicationName\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"addApplication\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_maxApps\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"InitBioMarin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"applicationNumber\",\"type\":\"uint256\"}],\"name\":\"RequestApplicationAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"GetApplicationAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"applicationNumber\",\"type\":\"uint256\"}],\"name\":\"GetApplicationAddressFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"applicationNumber\",\"type\":\"uint256\"}],\"name\":\"DeletingApplicationByNumber\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"ApplicationDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeleteAllApplications\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"applicationNumber\",\"type\":\"uint256\"}],\"name\":\"AllApplicationsDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"AddingApplication\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ShiftingApplications\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ApplicationsShifted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"ApplicationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"SearchingApplication\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"ApplicationFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"ApplicationNotFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"WrongNumber\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Killed\",\"type\":\"event\"}]"

// MetaContract is an auto generated Go binding around an Ethereum contract.
type MetaContract struct {
	MetaContractCaller     // Read-only binding to the contract
	MetaContractTransactor // Write-only binding to the contract
}

// MetaContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaContractSession struct {
	Contract     *MetaContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaContractCallerSession struct {
	Contract *MetaContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MetaContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaContractTransactorSession struct {
	Contract     *MetaContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MetaContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaContractRaw struct {
	Contract *MetaContract // Generic contract binding to access the raw methods on
}

// MetaContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaContractCallerRaw struct {
	Contract *MetaContractCaller // Generic read-only contract binding to access the raw methods on
}

// MetaContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaContractTransactorRaw struct {
	Contract *MetaContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetaContract creates a new instance of MetaContract, bound to a specific deployed contract.
func NewMetaContract(address common.Address, backend bind.ContractBackend) (*MetaContract, error) {
	contract, err := bindMetaContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaContract{MetaContractCaller: MetaContractCaller{contract: contract}, MetaContractTransactor: MetaContractTransactor{contract: contract}}, nil
}

// NewMetaContractCaller creates a new read-only instance of MetaContract, bound to a specific deployed contract.
func NewMetaContractCaller(address common.Address, caller bind.ContractCaller) (*MetaContractCaller, error) {
	contract, err := bindMetaContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MetaContractCaller{contract: contract}, nil
}

// NewMetaContractTransactor creates a new write-only instance of MetaContract, bound to a specific deployed contract.
func NewMetaContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaContractTransactor, error) {
	contract, err := bindMetaContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MetaContractTransactor{contract: contract}, nil
}

// bindMetaContract binds a generic wrapper to an already deployed contract.
func bindMetaContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaContract *MetaContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MetaContract.Contract.MetaContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaContract *MetaContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaContract.Contract.MetaContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaContract *MetaContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaContract.Contract.MetaContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaContract *MetaContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MetaContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaContract *MetaContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaContract *MetaContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaContract.Contract.contract.Transact(opts, method, params...)
}

// Applications is a free data retrieval call binding the contract method 0xdfefadff.
//
// Solidity: function applications( uint256) constant returns(contractAddress address, applicationName string)
func (_MetaContract *MetaContractCaller) Applications(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ContractAddress common.Address
	ApplicationName string
}, error) {
	ret := new(struct {
		ContractAddress common.Address
		ApplicationName string
	})
	out := ret
	err := _MetaContract.contract.Call(opts, out, "applications", arg0)
	return *ret, err
}

// Applications is a free data retrieval call binding the contract method 0xdfefadff.
//
// Solidity: function applications( uint256) constant returns(contractAddress address, applicationName string)
func (_MetaContract *MetaContractSession) Applications(arg0 *big.Int) (struct {
	ContractAddress common.Address
	ApplicationName string
}, error) {
	return _MetaContract.Contract.Applications(&_MetaContract.CallOpts, arg0)
}

// Applications is a free data retrieval call binding the contract method 0xdfefadff.
//
// Solidity: function applications( uint256) constant returns(contractAddress address, applicationName string)
func (_MetaContract *MetaContractCallerSession) Applications(arg0 *big.Int) (struct {
	ContractAddress common.Address
	ApplicationName string
}, error) {
	return _MetaContract.Contract.Applications(&_MetaContract.CallOpts, arg0)
}

// FindApplicationByNumber is a free data retrieval call binding the contract method 0xbd418d41.
//
// Solidity: function findApplicationByNumber(_number uint256) constant returns(address, string)
func (_MetaContract *MetaContractCaller) FindApplicationByNumber(opts *bind.CallOpts, _number *big.Int) (common.Address, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MetaContract.contract.Call(opts, out, "findApplicationByNumber", _number)
	return *ret0, *ret1, err
}

// FindApplicationByNumber is a free data retrieval call binding the contract method 0xbd418d41.
//
// Solidity: function findApplicationByNumber(_number uint256) constant returns(address, string)
func (_MetaContract *MetaContractSession) FindApplicationByNumber(_number *big.Int) (common.Address, string, error) {
	return _MetaContract.Contract.FindApplicationByNumber(&_MetaContract.CallOpts, _number)
}

// FindApplicationByNumber is a free data retrieval call binding the contract method 0xbd418d41.
//
// Solidity: function findApplicationByNumber(_number uint256) constant returns(address, string)
func (_MetaContract *MetaContractCallerSession) FindApplicationByNumber(_number *big.Int) (common.Address, string, error) {
	return _MetaContract.Contract.FindApplicationByNumber(&_MetaContract.CallOpts, _number)
}

// GetApplication is a free data retrieval call binding the contract method 0x1f2f4bfd.
//
// Solidity: function getApplication(_number uint256) constant returns(address, string)
func (_MetaContract *MetaContractCaller) GetApplication(opts *bind.CallOpts, _number *big.Int) (common.Address, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MetaContract.contract.Call(opts, out, "getApplication", _number)
	return *ret0, *ret1, err
}

// GetApplication is a free data retrieval call binding the contract method 0x1f2f4bfd.
//
// Solidity: function getApplication(_number uint256) constant returns(address, string)
func (_MetaContract *MetaContractSession) GetApplication(_number *big.Int) (common.Address, string, error) {
	return _MetaContract.Contract.GetApplication(&_MetaContract.CallOpts, _number)
}

// GetApplication is a free data retrieval call binding the contract method 0x1f2f4bfd.
//
// Solidity: function getApplication(_number uint256) constant returns(address, string)
func (_MetaContract *MetaContractCallerSession) GetApplication(_number *big.Int) (common.Address, string, error) {
	return _MetaContract.Contract.GetApplication(&_MetaContract.CallOpts, _number)
}

// GetMaxApplications is a free data retrieval call binding the contract method 0xd1595b97.
//
// Solidity: function getMaxApplications() constant returns(uint256)
func (_MetaContract *MetaContractCaller) GetMaxApplications(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MetaContract.contract.Call(opts, out, "getMaxApplications")
	return *ret0, err
}

// GetMaxApplications is a free data retrieval call binding the contract method 0xd1595b97.
//
// Solidity: function getMaxApplications() constant returns(uint256)
func (_MetaContract *MetaContractSession) GetMaxApplications() (*big.Int, error) {
	return _MetaContract.Contract.GetMaxApplications(&_MetaContract.CallOpts)
}

// GetMaxApplications is a free data retrieval call binding the contract method 0xd1595b97.
//
// Solidity: function getMaxApplications() constant returns(uint256)
func (_MetaContract *MetaContractCallerSession) GetMaxApplications() (*big.Int, error) {
	return _MetaContract.Contract.GetMaxApplications(&_MetaContract.CallOpts)
}

// GetNumberApplications is a free data retrieval call binding the contract method 0x66bb6036.
//
// Solidity: function getNumberApplications() constant returns(uint256)
func (_MetaContract *MetaContractCaller) GetNumberApplications(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MetaContract.contract.Call(opts, out, "getNumberApplications")
	return *ret0, err
}

// GetNumberApplications is a free data retrieval call binding the contract method 0x66bb6036.
//
// Solidity: function getNumberApplications() constant returns(uint256)
func (_MetaContract *MetaContractSession) GetNumberApplications() (*big.Int, error) {
	return _MetaContract.Contract.GetNumberApplications(&_MetaContract.CallOpts)
}

// GetNumberApplications is a free data retrieval call binding the contract method 0x66bb6036.
//
// Solidity: function getNumberApplications() constant returns(uint256)
func (_MetaContract *MetaContractCallerSession) GetNumberApplications() (*big.Int, error) {
	return _MetaContract.Contract.GetNumberApplications(&_MetaContract.CallOpts)
}

// AddApplication is a paid mutator transaction binding the contract method 0xf09ce682.
//
// Solidity: function addApplication(_address address, _name string) returns()
func (_MetaContract *MetaContractTransactor) AddApplication(opts *bind.TransactOpts, _address common.Address, _name string) (*types.Transaction, error) {
	return _MetaContract.contract.Transact(opts, "addApplication", _address, _name)
}

// AddApplication is a paid mutator transaction binding the contract method 0xf09ce682.
//
// Solidity: function addApplication(_address address, _name string) returns()
func (_MetaContract *MetaContractSession) AddApplication(_address common.Address, _name string) (*types.Transaction, error) {
	return _MetaContract.Contract.AddApplication(&_MetaContract.TransactOpts, _address, _name)
}

// AddApplication is a paid mutator transaction binding the contract method 0xf09ce682.
//
// Solidity: function addApplication(_address address, _name string) returns()
func (_MetaContract *MetaContractTransactorSession) AddApplication(_address common.Address, _name string) (*types.Transaction, error) {
	return _MetaContract.Contract.AddApplication(&_MetaContract.TransactOpts, _address, _name)
}

// DeleteAllApplications is a paid mutator transaction binding the contract method 0x5abb56ba.
//
// Solidity: function deleteAllApplications() returns()
func (_MetaContract *MetaContractTransactor) DeleteAllApplications(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaContract.contract.Transact(opts, "deleteAllApplications")
}

// DeleteAllApplications is a paid mutator transaction binding the contract method 0x5abb56ba.
//
// Solidity: function deleteAllApplications() returns()
func (_MetaContract *MetaContractSession) DeleteAllApplications() (*types.Transaction, error) {
	return _MetaContract.Contract.DeleteAllApplications(&_MetaContract.TransactOpts)
}

// DeleteAllApplications is a paid mutator transaction binding the contract method 0x5abb56ba.
//
// Solidity: function deleteAllApplications() returns()
func (_MetaContract *MetaContractTransactorSession) DeleteAllApplications() (*types.Transaction, error) {
	return _MetaContract.Contract.DeleteAllApplications(&_MetaContract.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_MetaContract *MetaContractTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaContract.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_MetaContract *MetaContractSession) Kill() (*types.Transaction, error) {
	return _MetaContract.Contract.Kill(&_MetaContract.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_MetaContract *MetaContractTransactorSession) Kill() (*types.Transaction, error) {
	return _MetaContract.Contract.Kill(&_MetaContract.TransactOpts)
}

// RemoveApplicationByNumber is a paid mutator transaction binding the contract method 0xb06383ed.
//
// Solidity: function removeApplicationByNumber(_number uint256) returns()
func (_MetaContract *MetaContractTransactor) RemoveApplicationByNumber(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _MetaContract.contract.Transact(opts, "removeApplicationByNumber", _number)
}

// RemoveApplicationByNumber is a paid mutator transaction binding the contract method 0xb06383ed.
//
// Solidity: function removeApplicationByNumber(_number uint256) returns()
func (_MetaContract *MetaContractSession) RemoveApplicationByNumber(_number *big.Int) (*types.Transaction, error) {
	return _MetaContract.Contract.RemoveApplicationByNumber(&_MetaContract.TransactOpts, _number)
}

// RemoveApplicationByNumber is a paid mutator transaction binding the contract method 0xb06383ed.
//
// Solidity: function removeApplicationByNumber(_number uint256) returns()
func (_MetaContract *MetaContractTransactorSession) RemoveApplicationByNumber(_number *big.Int) (*types.Transaction, error) {
	return _MetaContract.Contract.RemoveApplicationByNumber(&_MetaContract.TransactOpts, _number)
}
