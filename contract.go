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
const ContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getNumApprovers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getApprovedRole\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getApprover\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSpecification\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"addRequiredRole\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getApprovalsStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getApplicationName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRequiredRole\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"changeStatus\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getApproveRC\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_requestor\",\"type\":\"address\"},{\"name\":\"_size\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"numReqApprovers\",\"type\":\"uint256\"}],\"name\":\"ApplicationCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_size\",\"type\":\"uint256\"}],\"name\":\"WrongApplicationSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Killed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"AddingRequiredRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"numReqApprover\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"RequiredRoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"CheckingRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"RoleAlreadyExists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"AddingRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"numApprovedRole\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"AddingApprover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"numApprovers\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"ApproverAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ApplicationApproved\",\"type\":\"event\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
const ContractBin = `606060405234156200001057600080fd5b60405162001cfd38038062001cfd833981016040528080519060200190919080519060200190919080518201919050505b6000600c60016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600481905550600060068190555080600b9080519060200190620000eb92919062000271565b5060006005819055506000600c60006101000a81548160ff021916908360068111156200011457fe5b02179055506001821415620001565760016007600001819055506004600760010181905550610400600760020181905550600160076003018190555062000200565b60028214156200019357600260076000018190555060086007600101819055506108006007600201819055506002600760030181905550620001ff565b6003821415620001d057600460076000018190555060106007600101819055506110006007600201819055506004600760030181905550620001fe565b817f1e371dab69817be62ab25b28af940aff33a252f63235a64a43e4ae5020c59f3960405160405180910390a25b5b5b6005546000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f525a1f964d375777dfd5b80ed4ed417220ab8bcdb3cd53ce50e4435b7773f44c60405160405180910390a35b50505062000320565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002b457805160ff1916838001178555620002e5565b82800160010185558215620002e5579182015b82811115620002e4578251825591602001919060010190620002c7565b5b509050620002f49190620002f8565b5090565b6200031d91905b8082111562000319576000816000905550600101620002ff565b5090565b90565b6119cd80620003306000396000f300606060405236156100ce576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806313abc4f4146100d357806328b553f6146100fc57806341c0e1b5146101995780634e69d560146101ae5780637157e493146101d757806372e75458146102e1578063778f3de01461031f578063946a0c06146103bf5780639addff571461041c578063a3f5b1bf14610453578063c0a7772d146104a8578063d632e07f14610537578063e8025d77146105d4578063f39ac26c146105f7575b600080fd5b34156100de57600080fd5b6100e6610620565b6040518082815260200191505060405180910390f35b341561010757600080fd5b61011d600480803590602001909190505061062b565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561015e5780820151818401525b602081019050610142565b50505050905090810190601f16801561018b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101a457600080fd5b6101ac6106e7565b005b34156101b957600080fd5b6101c16107dd565b6040518082815260200191505060405180910390f35b34156101e257600080fd5b6101f86004808035906020019091905050610962565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561023d5780820151818401525b602081019050610221565b50505050905090810190601f16801561026a5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156102a45780820151818401525b602081019050610288565b50505050905090810190601f1680156102d15780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34156102ec57600080fd5b6102f4610ae0565b6040518085815260200184815260200183815260200182815260200194505050505060405180910390f35b341561032a57600080fd5b6103bd600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610b0d565b005b34156103ca57600080fd5b61041a600480803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506112f3565b005b341561042757600080fd5b61042f61146e565b60405180848152602001838152602001828152602001935050505060405180910390f35b341561045e57600080fd5b6104666115ec565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156104b357600080fd5b6104bb611616565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104fc5780820151818401525b6020810190506104e0565b50505050905090810190601f1680156105295780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561054257600080fd5b61055860048080359060200190919050506116bf565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105995780820151818401525b60208101905061057d565b50505050905090810190601f1680156105c65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156105df57600080fd5b6105f5600480803590602001909190505061177b565b005b341561060257600080fd5b61060a6118dd565b6040518082815260200191505060405180910390f35b600060045490505b90565b6106336118e8565b600260008381526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106da5780601f106106af576101008083540402835291602001916106da565b820191906000526020600020905b8154815290600101906020018083116106bd57829003601f168201915b505050505090505b919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156107da576000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f4b0bc4f25f8d0b92d2e12b686ba96cd75e4e69325e6cf7b1f3119d14eaf2cbdf60405160405180910390a26000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b60008060068111156107eb57fe5b600c60009054906101000a900460ff16600681111561080657fe5b1415610815576001905061095f565b6001600681111561082257fe5b600c60009054906101000a900460ff16600681111561083d57fe5b141561084c576002905061095f565b6002600681111561085957fe5b600c60009054906101000a900460ff16600681111561087457fe5b1415610883576003905061095f565b6003600681111561089057fe5b600c60009054906101000a900460ff1660068111156108ab57fe5b14156108ba576004905061095f565b600460068111156108c757fe5b600c60009054906101000a900460ff1660068111156108e257fe5b14156108f1576005905061095f565b600560068111156108fe57fe5b600c60009054906101000a900460ff16600681111561091957fe5b1415610928576006905061095f565b60068081111561093457fe5b600c60009054906101000a900460ff16600681111561094f57fe5b141561095e576007905061095f565b5b90565b61096a6118e8565b6109726118e8565b6003600084815260200190815260200160002060010160036000858152602001908152602001600020600001818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a335780601f10610a0857610100808354040283529160200191610a33565b820191906000526020600020905b815481529060010190602001808311610a1657829003601f168201915b50505050509150808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610acf5780601f10610aa457610100808354040283529160200191610acf565b820191906000526020600020905b815481529060010190602001808311610ab257829003601f168201915b50505050509050915091505b915091565b60008060008060076000015460076001015460076002015460076003015493509350935093505b90919293565b6000806000809050836040518082805190602001908083835b602083101515610b4c57805182525b602082019150602081019050602083039250610b26565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020856040518082805190602001908083835b602083101515610bb057805182525b602082019150602081019050602083039250610b8a565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207fb3403027d0ce02c730dc818e7c0d11fa3b3cac8330d1963872096da9d896da9860405160405180910390a3846003600060045481526020019081526020016000206001019080519060200190610c359291906118fc565b50836003600060045481526020019081526020016000206000019080519060200190610c629291906118fc565b5060016004600082825401925050819055506003600060016004540381526020019081526020016000206000016040518082805460018160011615610100020316600290048015610cea5780601f10610cc8576101008083540402835291820191610cea565b820191906000526020600020905b815481529060010190602001808311610cd6575b505091505060405180910390206003600060016004540381526020019081526020016000206001016040518082805460018160011615610100020316600290048015610d6d5780601f10610d4b576101008083540402835291820191610d6d565b820191906000526020600020905b815481529060010190602001808311610d59575b505091505060405180910390206004547f905a46ea8df58076663ef800cf17e413a5dc1fc11b465a1ed579864ccd7db63760405160405180910390a4836040518082805190602001908083835b602083101515610de057805182525b602082019150602081019050602083039250610dba565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207f8fc1bb51d5ce7c93bfc953314351afc9223a322d093523de57d10d2cc68fdd4160405160405180910390a2600092505b600654831015610fef57836040518082805190602001908083835b602083101515610e7f57805182525b602082019150602081019050602083039250610e59565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916600260008581526020019081526020016000206040518082805460018160011615610100020316600290048015610f1e5780601f10610efc576101008083540402835291820191610f1e565b820191906000526020600020905b815481529060010190602001808311610f0a575b50509150506040518091039020600019161415610fe15760019050600260008481526020019081526020016000206040518082805460018160011615610100020316600290048015610fa75780601f10610f85576101008083540402835291820191610fa7565b820191906000526020600020905b815481529060010190602001808311610f93575b505091505060405180910390207f1785dad0ffb50908010ec5e6ae5a882ec193b0b21e039595443cde29a953be7b60405160405180910390a25b5b8280600101935050610e3e565b600092505b6005548310156110fe5760016000848152602001908152602001600020604051808280546001816001161561010002031660029004801561106c5780601f1061104a57610100808354040283529182019161106c565b820191906000526020600020905b815481529060010190602001808311611058575b5050915050604051809103902060001916846040518082805190602001908083835b6020831015156110b457805182525b60208201915060208101905060208303925061108e565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191614156110f057600191505b5b8280600101935050610ff4565b60008114801561110e5750600182145b1561128e57836040518082805190602001908083835b60208310151561114a57805182525b602082019150602081019050602083039250611124565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207f31111fb0ab34fc149a6179f52431820d3929a52a20a79c346754117d42ebe2e060405160405180910390a28360026000600654815260200190815260200160002090805190602001906111cc9291906118fc565b50600160066000828254019250508190555060026000600160065403815260200190815260200160002060405180828054600181600116156101000203166002900480156112515780601f1061122f576101008083540402835291820191611251565b820191906000526020600020905b81548152906001019060200180831161123d575b505091505060405180910390206006547f8a8c6e5d817d6b4b89ece61a1307d0361795dbbf0ef54c2693409c669bf500a960405160405180910390a35b60055460065414156112eb576005600c60006101000a81548160ff021916908360068111156112b957fe5b02179055507f0c8fd27f0e202d53a731dc39238816f51f38331d2e1c6d96ac4f40429a8a3a1a60405160405180910390a15b5b5050505050565b806040518082805190602001908083835b60208310151561132a57805182525b602082019150602081019050602083039250611304565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207fba741c02698f40b54c299b0d26b447d5fb4065d70b357594397f1d8e4c294efc60405160405180910390a28060016000600554815260200190815260200160002090805190602001906113ac9291906118fc565b506001600560008282540192505081905550600160006005548152602001908152602001600020604051808280546001816001161561010002031660029004801561142e5780601f1061140c57610100808354040283529182019161142e565b820191906000526020600020905b81548152906001019060200180831161141a575b505091505060405180910390206005547ff4155a5029e40a4636a842980be6bf1be532789feb40549a24ddc8bcb255913560405160405180910390a35b50565b6000806000806000600681111561148157fe5b600c60009054906101000a900460ff16600681111561149c57fe5b14156114a757600190505b600160068111156114b457fe5b600c60009054906101000a900460ff1660068111156114cf57fe5b14156114da57600290505b600260068111156114e757fe5b600c60009054906101000a900460ff16600681111561150257fe5b141561150d57600390505b6003600681111561151a57fe5b600c60009054906101000a900460ff16600681111561153557fe5b141561154057600490505b6004600681111561154d57fe5b600c60009054906101000a900460ff16600681111561156857fe5b141561157357600590505b6005600681111561158057fe5b600c60009054906101000a900460ff16600681111561159b57fe5b14156115a657600690505b6006808111156115b257fe5b600c60009054906101000a900460ff1660068111156115cd57fe5b14156115d857600790505b806005546006549350935093505b50909192565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b90565b61161e6118e8565b600b8054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156116b45780601f10611689576101008083540402835291602001916116b4565b820191906000526020600020905b81548152906001019060200180831161169757829003601f168201915b505050505090505b90565b6116c76118e8565b600160008381526020019081526020016000208054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561176e5780601f106117435761010080835404028352916020019161176e565b820191906000526020600020905b81548152906001019060200180831161175157829003601f168201915b505050505090505b919050565b60018114156117ad576000600c60006101000a81548160ff021916908360068111156117a357fe5b02179055506118da565b60028114156117df576001600c60006101000a81548160ff021916908360068111156117d557fe5b02179055506118da565b6003811415611811576002600c60006101000a81548160ff0219169083600681111561180757fe5b02179055506118da565b6004811415611843576003600c60006101000a81548160ff0219169083600681111561183957fe5b02179055506118da565b6005811415611875576004600c60006101000a81548160ff0219169083600681111561186b57fe5b02179055506118da565b60068114156118a7576005600c60006101000a81548160ff0219169083600681111561189d57fe5b02179055506118da565b60078114156118d9576006600c60006101000a81548160ff021916908360068111156118cf57fe5b02179055506118da565b5b50565b6000600d5490505b90565b602060405190810160405280600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061193d57805160ff191683800117855561196b565b8280016001018555821561196b579182015b8281111561196a57825182559160200191906001019061194f565b5b509050611978919061197c565b5090565b61199e91905b8082111561199a576000816000905550600101611982565b5090565b905600a165627a7a72305820d71aeb55dfa9d6f161e148e82a9562dd5ca67ab3600e1ebba4f59ef67d7916740029`

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
// Solidity: function getApprover(index uint256) constant returns(string, string)
func (_Contract *ContractCaller) GetApprover(opts *bind.CallOpts, index *big.Int) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Contract.contract.Call(opts, out, "getApprover", index)
	return *ret0, *ret1, err
}

// GetApprover is a free data retrieval call binding the contract method 0x7157e493.
//
// Solidity: function getApprover(index uint256) constant returns(string, string)
func (_Contract *ContractSession) GetApprover(index *big.Int) (string, string, error) {
	return _Contract.Contract.GetApprover(&_Contract.CallOpts, index)
}

// GetApprover is a free data retrieval call binding the contract method 0x7157e493.
//
// Solidity: function getApprover(index uint256) constant returns(string, string)
func (_Contract *ContractCallerSession) GetApprover(index *big.Int) (string, string, error) {
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

// Approve is a paid mutator transaction binding the contract method 0x778f3de0.
//
// Solidity: function approve(_name string, _role string) returns()
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, _name string, _role string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", _name, _role)
}

// Approve is a paid mutator transaction binding the contract method 0x778f3de0.
//
// Solidity: function approve(_name string, _role string) returns()
func (_Contract *ContractSession) Approve(_name string, _role string) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, _name, _role)
}

// Approve is a paid mutator transaction binding the contract method 0x778f3de0.
//
// Solidity: function approve(_name string, _role string) returns()
func (_Contract *ContractTransactorSession) Approve(_name string, _role string) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, _name, _role)
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
