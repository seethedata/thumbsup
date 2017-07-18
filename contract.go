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
const ContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getNumApprovers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getApprovedRole\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getApprover\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSpecification\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"addRequiredRole\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getApprovalsStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getApplicationName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRequiredRole\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_role\",\"type\":\"string\"},{\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"changeStatus\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_requestor\",\"type\":\"address\"},{\"name\":\"_size\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"numReqApprovers\",\"type\":\"uint256\"}],\"name\":\"ApplicationCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_size\",\"type\":\"uint256\"}],\"name\":\"WrongApplicationSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Killed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"AddingRequiredRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"numReqApprover\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"RequiredRoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"CheckingRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"RoleAlreadyExists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"AddingRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"numApprovedRole\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"AddingApprover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"numApprovers\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"ApproverAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ApplicationApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"CalledChangeStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"StatusChanged\",\"type\":\"event\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
const ContractBin = `606060405234156200000d57fe5b604051620017e0380380620017e083398101604090815281516020830151918301519092015b60008054600160a060020a031916600160a060020a038516178155600481905560065580516200006b90600b90602084019062000155565b5060006005556001600c8190558214156200009b57600160078190556004600855610400600955600a5562000110565b8160021415620000bf576002600781905560088055610800600955600a5562000110565b8160031415620000e457600460078190556010600855611000600955600a5562000110565b60405182907f1e371dab69817be62ab25b28af940aff33a252f63235a64a43e4ae5020c59f3990600090a25b5b5b60055460008054604051600160a060020a03909116917f525a1f964d375777dfd5b80ed4ed417220ab8bcdb3cd53ce50e4435b7773f44c91a35b505050620001ff565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200019857805160ff1916838001178555620001c8565b82800160010185558215620001c8579182015b82811115620001c8578251825591602001919060010190620001ab565b5b50620001d7929150620001db565b5090565b620001fc91905b80821115620001d75760008155600101620001e2565b5090565b90565b6115d1806200020f6000396000f300606060405236156100c25763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166313abc4f481146100c457806328b553f6146100e657806341c0e1b5146101795780634e69d5601461018b5780637157e493146101ad57806372e7545814610327578063946a0c061461035d5780639addff57146103b5578063a3f5b1bf146103e3578063c0a7772d1461041c578063d632e07f146104ac578063e7e1030a1461053f578063e8025d7714610611575bfe5b34156100cc57fe5b6100d4610626565b60408051918252519081900360200190f35b34156100ee57fe5b6100f960043561062d565b60408051602080825283518183015283519192839290830191850190808383821561013f575b80518252602083111561013f57601f19909201916020918201910161011f565b505050905090810190601f16801561016b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561018157fe5b6101896106d3565b005b341561019357fe5b6100d4610758565b60408051918252519081900360200190f35b34156101b557fe5b6101c060043561075f565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360008314610213575b80518252602083111561021357601f1990920191602091820191016101f3565b505050905090810190601f16801561023f5780820380516001836020036101000a031916815260200191505b508481038352865181528651602091820191880190808383821561027e575b80518252602083111561027e57601f19909201916020918201910161025e565b505050905090810190601f1680156102aa5780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838382156102e9575b8051825260208311156102e957601f1990920191602091820191016102c9565b505050905090810190601f1680156103155780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b341561032f57fe5b610337610948565b604080519485526020850193909352838301919091526060830152519081900360800190f35b341561036557fe5b610189600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284375094965061095b95505050505050565b005b34156103bd57fe5b6103c5610abf565b60408051938452602084019290925282820152519081900360600190f35b34156103eb57fe5b6103f3610ace565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b341561042457fe5b6100f9610aeb565b60408051602080825283518183015283519192839290830191850190808383821561013f575b80518252602083111561013f57601f19909201916020918201910161011f565b505050905090810190601f16801561016b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156104b457fe5b6100f9600435610b84565b60408051602080825283518183015283519192839290830191850190808383821561013f575b80518252602083111561013f57601f19909201916020918201910161011f565b505050905090810190601f16801561016b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561054757fe5b610189600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284375050604080516020601f89358b0180359182018390048302840183019094528083529799988101979196509182019450925082915084018382808284375050604080516020601f89358b01803591820183900483028401830190945280835297999881019791965091820194509250829150840183828082843750949650610c2b95505050505050565b005b341561061957fe5b610189600435611494565b005b6004545b90565b6106356114f3565b600082815260026020818152604092839020805484516001821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156106c65780601f1061069b576101008083540402835291602001916106c6565b820191906000526020600020905b8154815290600101906020018083116106a957829003601f168201915b505050505090505b919050565b6000543373ffffffffffffffffffffffffffffffffffffffff90811691161415610755576000805460405173ffffffffffffffffffffffffffffffffffffffff909116917f4b0bc4f25f8d0b92d2e12b686ba96cd75e4e69325e6cf7b1f3119d14eaf2cbdf91a260005473ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b600c545b90565b6107676114f3565b61076f6114f3565b6107776114f3565b6000848152600360209081526040918290206001808201805485516002938216156101000260001901909116839004601f8101869004860282018601909652858152909492939184019290918591908301828280156108175780601f106107ec57610100808354040283529160200191610817565b820191906000526020600020905b8154815290600101906020018083116107fa57829003601f168201915b5050855460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959850879450925084019050828280156108a55780601f1061087a576101008083540402835291602001916108a5565b820191906000526020600020905b81548152906001019060200180831161088857829003601f168201915b5050845460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959750869450925084019050828280156109335780601f1061090857610100808354040283529160200191610933565b820191906000526020600020905b81548152906001019060200180831161091657829003601f168201915b505050505090509250925092505b9193909250565b600754600854600954600a545b90919293565b806040518082805190602001908083835b6020831061098b5780518252601f19909201916020918201910161096c565b5181516020939093036101000a60001901801990911692169190911790526040519201829003822093507fba741c02698f40b54c299b0d26b447d5fb4065d70b357594397f1d8e4c294efc92506000919050a260055460009081526001602090815260409091208251610a0092840190611505565b5060016005600082825401925050819055506001600060055481526020019081526020016000206040518082805460018160011615610100020316600290048015610a825780601f10610a60576101008083540402835291820191610a82565b820191906000526020600020905b815481529060010190602001808311610a6e575b505060405190819003812060055490935091507ff4155a5029e40a4636a842980be6bf1be532789feb40549a24ddc8bcb255913590600090a35b50565b600c546005546006545b909192565b60005473ffffffffffffffffffffffffffffffffffffffff165b90565b610af36114f3565b600b805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610b795780601f10610b4e57610100808354040283529160200191610b79565b820191906000526020600020905b815481529060010190602001808311610b5c57829003601f168201915b505050505090505b90565b610b8c6114f3565b60008281526001602081815260409283902080548451600294821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156106c65780601f1061069b576101008083540402835291602001916106c6565b820191906000526020600020905b8154815290600101906020018083116106a957829003601f168201915b505050505090505b919050565b60006000600060009050836040518082805190602001908083835b60208310610c655780518252601f199092019160209182019101610c46565b51815160209384036101000a60001901801990921691161790526040519190930181900381208a519095508a945090928392508401908083835b60208310610cbe5780518252601f199092019160209182019101610c9f565b51815160209384036101000a60001901801990921691161790526040519190930181900381208c519095508c945090928392508401908083835b60208310610d175780518252601f199092019160209182019101610cf8565b5181516020939093036101000a60001901801990911692169190911790526040519201829003822093507f34328b5709ce0e15a7fe963f010fe4fb2b72713edfa4c63e0f24ec697c77b75c92506000919050a460045460009081526003602090815260409091208751610d9292600190920191890190611505565b5060045460009081526003602090815260409091208651610db592880190611505565b5060045460009081526003602090815260409091208551610dde92600290920191870190611505565b5060048054600181810190925560009081526003602052604090819020905181549192909182918491600261010092821615929092026000190116048015610e5d5780601f10610e3b576101008083540402835291820191610e5d565b820191906000526020600020905b815481529060010190602001808311610e49575b505091505060405180910390206003600060016004540381526020019081526020016000206001016040518082805460018160011615610100020316600290048015610ee05780601f10610ebe576101008083540402835291820191610ee0565b820191906000526020600020905b815481529060010190602001808311610ecc575b5050604080519182900382206004546000198082016000908152600360209081529085902081875260029081018054610100600182161502909401909316049086018190529296509094507f73eb4df6129515e874adb10f5eddb8879bbc213d07740c99054de0a386a149a29390929091829182019084908015610fa55780601f10610f7a57610100808354040283529160200191610fa5565b820191906000526020600020905b815481529060010190602001808311610f8857829003601f168201915b50509250505060405180910390a4846040518082805190602001908083835b60208310610fe35780518252601f199092019160209182019101610fc4565b5181516020939093036101000a60001901801990911692169190911790526040519201829003822093507f8fc1bb51d5ce7c93bfc953314351afc9223a322d093523de57d10d2cc68fdd4192506000919050a2600092505b6006548310156111de57846040518082805190602001908083835b602083106110755780518252601f199092019160209182019101611056565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020600019166002600085815260200190815260200160002060405180828054600181600116156101000203166002900480156111145780601f106110f2576101008083540402835291820191611114565b820191906000526020600020905b815481529060010190602001808311611100575b505091505060405180910390206000191614156111d2576001905060026000848152602001908152602001600020604051808280546001816001161561010002031660029004801561119d5780601f1061117b57610100808354040283529182019161119d565b820191906000526020600020905b815481529060010190602001808311611189575b505060405190819003812092507f1785dad0ffb50908010ec5e6ae5a882ec193b0b21e039595443cde29a953be7b9150600090a25b5b60019092019161103b565b600092505b6005548310156112e05760016000848152602001908152602001600020604051808280546001816001161561010002031660029004801561125b5780601f1061123957610100808354040283529182019161125b565b820191906000526020600020905b815481529060010190602001808311611247575b50506040519081900381208851909350889250819060208401908083835b602083106112985780518252601f199092019160209182019101611279565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191614156112d457600191505b5b6001909201916111e3565b801580156112ee5750816001145b1561145057846040518082805190602001908083835b602083106113235780518252601f199092019160209182019101611304565b5181516020939093036101000a60001901801990911692169190911790526040519201829003822093507f31111fb0ab34fc149a6179f52431820d3929a52a20a79c346754117d42ebe2e092506000919050a26006546000908152600260209081526040909120865161139892880190611505565b506006805460018181019092556000908152600260208190526040918290209151825492939092839285926101009083161502600019019091160480156114165780601f106113f4576101008083540402835291820191611416565b820191906000526020600020905b815481529060010190602001808311611402575b505060405190819003812060065490935091507f8a8c6e5d817d6b4b89ece61a1307d0361795dbbf0ef54c2693409c669bf500a990600090a35b600554600654141561148b576006600c556040517f0c8fd27f0e202d53a731dc39238816f51f38331d2e1c6d96ac4f40429a8a3a1a90600090a15b5b505050505050565b60405181907f6c67c6bd783ad5263b3f6c7297ad3b36de1950167ef7b88de39d3a0dd6bf041790600090a2600c81905560405181907ff3daf54aec5860f0f13cb1123778b256daa0b27653fd836691b4bc0eb288c61090600090a25b50565b60408051602081019091526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061154657805160ff1916838001178555611573565b82800160010185558215611573579182015b82811115611573578251825591602001919060010190611558565b5b50611580929150611584565b5090565b61062a91905b80821115611580576000815560010161158a565b5090565b905600a165627a7a723058208c3f8a29ad9bfffc1c45256591ebce3cf4258d7c256fd29090efb730c08979cd0029`

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _requestor common.Address, _size *big.Int, _name string) (common.Address, *types.Transaction, *Contract, error) {
    parsed, err := abi.JSON(strings.NewReader(ContractABI))
    if err != nil {
        return common.Address{}, nil, nil, err
    }
    address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend, _requestor, _size, _name)
    if err != nil {
        return common.Address{}, nil, nil, err
    }
    return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}}, nil
}

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
// Solidity: function getApprovalsStatus() constant returns(uint256, uint256, uint256)
func (_Contract *ContractCaller) GetApprovalsStatus(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
    var (
        ret0 = new(*big.Int)
        ret1 = new(*big.Int)
        ret2 = new(*big.Int)
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
// Solidity: function getApprovalsStatus() constant returns(uint256, uint256, uint256)
func (_Contract *ContractSession) GetApprovalsStatus() (*big.Int, *big.Int, *big.Int, error) {
    return _Contract.Contract.GetApprovalsStatus(&_Contract.CallOpts)
}

// GetApprovalsStatus is a free data retrieval call binding the contract method 0x9addff57.
//
// Solidity: function getApprovalsStatus() constant returns(uint256, uint256, uint256)
func (_Contract *ContractCallerSession) GetApprovalsStatus() (*big.Int, *big.Int, *big.Int, error) {
    return _Contract.Contract.GetApprovalsStatus(&_Contract.CallOpts)
}

// GetApprovedRole is a free data retrieval call binding the contract method 0x28b553f6.
//
// Solidity: function getApprovedRole(index uint256) constant returns(string)
func (_Contract *ContractCaller) GetApprovedRole(opts *bind.CallOpts, index *big.Int) (string, error) {
    var (
        ret0 = new(string)
    )
    out := ret0
    err := _Contract.contract.Call(opts, out, "getApprovedRole", index)
    return *ret0, err
}

// GetApprovedRole is a free data retrieval call binding the contract method 0x28b553f6.
//
// Solidity: function getApprovedRole(index uint256) constant returns(string)
func (_Contract *ContractSession) GetApprovedRole(index *big.Int) (string, error) {
    return _Contract.Contract.GetApprovedRole(&_Contract.CallOpts, index)
}

// GetApprovedRole is a free data retrieval call binding the contract method 0x28b553f6.
//
// Solidity: function getApprovedRole(index uint256) constant returns(string)
func (_Contract *ContractCallerSession) GetApprovedRole(index *big.Int) (string, error) {
    return _Contract.Contract.GetApprovedRole(&_Contract.CallOpts, index)
}

// GetApprover is a free data retrieval call binding the contract method 0x7157e493.
//
// Solidity: function getApprover(index uint256) constant returns(string, string, string)
func (_Contract *ContractCaller) GetApprover(opts *bind.CallOpts, index *big.Int) (string, string, string, error) {
    var (
        ret0 = new(string)
        ret1 = new(string)
        ret2 = new(string)
    )
    out := &[]interface{}{
        ret0,
        ret1,
        ret2,
    }
    err := _Contract.contract.Call(opts, out, "getApprover", index)
    return *ret0, *ret1, *ret2, err
}

// GetApprover is a free data retrieval call binding the contract method 0x7157e493.
//
// Solidity: function getApprover(index uint256) constant returns(string, string, string)
func (_Contract *ContractSession) GetApprover(index *big.Int) (string, string, string, error) {
    return _Contract.Contract.GetApprover(&_Contract.CallOpts, index)
}

// GetApprover is a free data retrieval call binding the contract method 0x7157e493.
//
// Solidity: function getApprover(index uint256) constant returns(string, string, string)
func (_Contract *ContractCallerSession) GetApprover(index *big.Int) (string, string, string, error) {
    return _Contract.Contract.GetApprover(&_Contract.CallOpts, index)
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

// GetRequiredRole is a free data retrieval call binding the contract method 0xd632e07f.
//
// Solidity: function getRequiredRole(index uint256) constant returns(string)
func (_Contract *ContractCaller) GetRequiredRole(opts *bind.CallOpts, index *big.Int) (string, error) {
    var (
        ret0 = new(string)
    )
    out := ret0
    err := _Contract.contract.Call(opts, out, "getRequiredRole", index)
    return *ret0, err
}

// GetRequiredRole is a free data retrieval call binding the contract method 0xd632e07f.
//
// Solidity: function getRequiredRole(index uint256) constant returns(string)
func (_Contract *ContractSession) GetRequiredRole(index *big.Int) (string, error) {
    return _Contract.Contract.GetRequiredRole(&_Contract.CallOpts, index)
}

// GetRequiredRole is a free data retrieval call binding the contract method 0xd632e07f.
//
// Solidity: function getRequiredRole(index uint256) constant returns(string)
func (_Contract *ContractCallerSession) GetRequiredRole(index *big.Int) (string, error) {
    return _Contract.Contract.GetRequiredRole(&_Contract.CallOpts, index)
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

// AddRequiredRole is a paid mutator transaction binding the contract method 0x946a0c06.
//
// Solidity: function addRequiredRole(_role string) returns()
func (_Contract *ContractTransactor) AddRequiredRole(opts *bind.TransactOpts, _role string) (*types.Transaction, error) {
    return _Contract.contract.Transact(opts, "addRequiredRole", _role)
}

// AddRequiredRole is a paid mutator transaction binding the contract method 0x946a0c06.
//
// Solidity: function addRequiredRole(_role string) returns()
func (_Contract *ContractSession) AddRequiredRole(_role string) (*types.Transaction, error) {
    return _Contract.Contract.AddRequiredRole(&_Contract.TransactOpts, _role)
}

// AddRequiredRole is a paid mutator transaction binding the contract method 0x946a0c06.
//
// Solidity: function addRequiredRole(_role string) returns()
func (_Contract *ContractTransactorSession) AddRequiredRole(_role string) (*types.Transaction, error) {
    return _Contract.Contract.AddRequiredRole(&_Contract.TransactOpts, _role)
}

// Approve is a paid mutator transaction binding the contract method 0xe7e1030a.
//
// Solidity: function approve(_name string, _role string, _date string) returns()
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, _name string, _role string, _date string) (*types.Transaction, error) {
    return _Contract.contract.Transact(opts, "approve", _name, _role, _date)
}

// Approve is a paid mutator transaction binding the contract method 0xe7e1030a.
//
// Solidity: function approve(_name string, _role string, _date string) returns()
func (_Contract *ContractSession) Approve(_name string, _role string, _date string) (*types.Transaction, error) {
    return _Contract.Contract.Approve(&_Contract.TransactOpts, _name, _role, _date)
}

// Approve is a paid mutator transaction binding the contract method 0xe7e1030a.
//
// Solidity: function approve(_name string, _role string, _date string) returns()
func (_Contract *ContractTransactorSession) Approve(_name string, _role string, _date string) (*types.Transaction, error) {
    return _Contract.Contract.Approve(&_Contract.TransactOpts, _name, _role, _date)
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
