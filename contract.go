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
const ContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getNumApprovers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCreatedDate\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getApprovedRole\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getApprover\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSpecification\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"setDeployedDate\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"addRequiredRole\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getApprovalsStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDeployedDate\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getApplicationName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRequiredRole\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_role\",\"type\":\"string\"},{\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"changeStatus\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_requestor\",\"type\":\"address\"},{\"name\":\"_size\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_date\",\"type\":\"string\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"numReqApprovers\",\"type\":\"uint256\"}],\"name\":\"ApplicationCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_size\",\"type\":\"uint256\"}],\"name\":\"WrongApplicationSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Killed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"AddingRequiredRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"numReqApprover\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"RequiredRoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"CheckingRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"RoleAlreadyExists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"AddingRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"numApprovedRole\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"AddingApprover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"numApprovers\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"ApproverAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ApplicationApproved\",\"type\":\"event\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
const ContractBin = `606060405234156200000d57fe5b60405162001e7638038062001e76833981016040528080519060200190919080519060200190919080518201919060200180518201919050505b83600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600481905550600060068190555081600b9080519060200190620000b092919062000289565b5060006005819055506001600e8190555080600c9080519060200190620000d992919062000289565b50604060405190810160405280600381526020017f4e2f410000000000000000000000000000000000000000000000000000000000815250600d90805190602001906200012892919062000289565b506001831415620001665760016007600001819055506004600760010181905550610400600760020181905550600160076003018190555062000213565b6002831415620001a35760026007600001819055506008600760010181905550610800600760020181905550600260076003018190555062000212565b6003831415620001e05760046007600001819055506010600760010181905550611000600760020181905550600460076003018190555062000211565b827f1e371dab69817be62ab25b28af940aff33a252f63235a64a43e4ae5020c59f3960405180905060405180910390a25b5b5b600554600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f525a1f964d375777dfd5b80ed4ed417220ab8bcdb3cd53ce50e4435b7773f44c60405180905060405180910390a35b5050505062000338565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002cc57805160ff1916838001178555620002fd565b82800160010185558215620002fd579182015b82811115620002fc578251825591602001919060010190620002df565b5b5090506200030c919062000310565b5090565b6200033591905b808211156200033157600081600090555060010162000317565b5090565b90565b611b2e80620003486000396000f300606060405236156100e4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806313abc4f4146100e657806328aea7351461010c57806328b553f6146101a557806341c0e1b51461024c5780634e69d5601461025e5780637157e4931461028457806372e754581461041f57806382756e7a1461045a578063946a0c06146104b45780639addff571461050e5780639d37c22a14610542578063a3f5b1bf146105db578063c0a7772d1461062d578063d632e07f146106c6578063e7e1030a1461076d578063e8025d771461084d575bfe5b34156100ee57fe5b6100f661086d565b6040518082815260200191505060405180910390f35b341561011457fe5b61011c610878565b604051808060200182810382528381815181526020019150805190602001908083836000831461016b575b80518252602083111561016b57602082019150602081019050602083039250610147565b505050905090810190601f1680156101975780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101ad57fe5b6101c36004808035906020019091905050610921565b6040518080602001828103825283818151815260200191508051906020019080838360008314610212575b805182526020831115610212576020820191506020810190506020830392506101ee565b505050905090810190601f16801561023e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561025457fe5b61025c6109dd565b005b341561026657fe5b61026e610ad9565b6040518082815260200191505060405180910390f35b341561028c57fe5b6102a26004808035906020019091905050610ae4565b604051808060200180602001806020018481038452878181518152602001915080519060200190808383600083146102f9575b8051825260208311156102f9576020820191506020810190506020830392506102d5565b505050905090810190601f1680156103255780820380516001836020036101000a031916815260200191505b5084810383528681815181526020019150805190602001908083836000831461036d575b80518252602083111561036d57602082019150602081019050602083039250610349565b505050905090810190601f1680156103995780820380516001836020036101000a031916815260200191505b508481038252858181518152602001915080519060200190808383600083146103e1575b8051825260208311156103e1576020820191506020810190506020830392506103bd565b505050905090810190601f16801561040d5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b341561042757fe5b61042f610d20565b6040518085815260200184815260200183815260200182815260200194505050505060405180910390f35b341561046257fe5b6104b2600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610d4f565b005b34156104bc57fe5b61050c600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610d6a565b005b341561051657fe5b61051e610ee8565b60405180848152602001838152602001828152602001935050505060405180910390f35b341561054a57fe5b610552610f03565b60405180806020018281038252838181518152602001915080519060200190808383600083146105a1575b8051825260208311156105a15760208201915060208101905060208303925061057d565b505050905090810190601f1680156105cd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156105e357fe5b6105eb610fac565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561063557fe5b61063d610fd7565b604051808060200182810382528381815181526020019150805190602001908083836000831461068c575b80518252602083111561068c57602082019150602081019050602083039250610668565b505050905090810190601f1680156106b85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156106ce57fe5b6106e46004808035906020019091905050611080565b6040518080602001828103825283818151815260200191508051906020019080838360008314610733575b8051825260208311156107335760208201915060208101905060208303925061070f565b505050905090810190601f16801561075f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561077557fe5b61084b600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190505061113c565b005b341561085557fe5b61086b6004808035906020019091905050611a3e565b005b600060045490505b90565b610880611a49565b600c8054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109165780601f106108eb57610100808354040283529160200191610916565b820191906000526020600020905b8154815290600101906020018083116108f957829003601f168201915b505050505090505b90565b610929611a49565b600260008381526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109d05780601f106109a5576101008083540402835291602001916109d0565b820191906000526020600020905b8154815290600101906020018083116109b357829003601f168201915b505050505090505b919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610ad657600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f4b0bc4f25f8d0b92d2e12b686ba96cd75e4e69325e6cf7b1f3119d14eaf2cbdf60405180905060405180910390a2600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b6000600e5490505b90565b610aec611a49565b610af4611a49565b610afc611a49565b600360008581526020019081526020016000206001016003600086815260200190815260200160002060000160036000878152602001908152602001600020600201828054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bd35780601f10610ba857610100808354040283529160200191610bd3565b820191906000526020600020905b815481529060010190602001808311610bb657829003601f168201915b50505050509250818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c6f5780601f10610c4457610100808354040283529160200191610c6f565b820191906000526020600020905b815481529060010190602001808311610c5257829003601f168201915b50505050509150808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d0b5780601f10610ce057610100808354040283529160200191610d0b565b820191906000526020600020905b815481529060010190602001808311610cee57829003601f168201915b505050505090509250925092505b9193909250565b600060006000600060076000015460076001015460076002015460076003015493509350935093505b90919293565b80600d9080519060200190610d65929190611a5d565b505b50565b806040518082805190602001908083835b60208310610d9e5780518252602082019150602081019050602083039250610d7b565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207fba741c02698f40b54c299b0d26b447d5fb4065d70b357594397f1d8e4c294efc60405180905060405180910390a2806001600060055481526020019081526020016000209080519060200190610e23929190611a5d565b5060016005600082825401925050819055506001600060055481526020019081526020016000206040518082805460018160011615610100020316600290048015610ea55780601f10610e83576101008083540402835291820191610ea5565b820191906000526020600020905b815481529060010190602001808311610e91575b505091505060405180910390206005547ff4155a5029e40a4636a842980be6bf1be532789feb40549a24ddc8bcb255913560405180905060405180910390a35b50565b600060006000600e546005546006549250925092505b909192565b610f0b611a49565b600d8054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610fa15780601f10610f7657610100808354040283529160200191610fa1565b820191906000526020600020905b815481529060010190602001808311610f8457829003601f168201915b505050505090505b90565b6000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b90565b610fdf611a49565b600b8054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110755780601f1061104a57610100808354040283529160200191611075565b820191906000526020600020905b81548152906001019060200180831161105857829003601f168201915b505050505090505b90565b611088611a49565b600160008381526020019081526020016000208054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561112f5780601f106111045761010080835404028352916020019161112f565b820191906000526020600020905b81548152906001019060200180831161111257829003601f168201915b505050505090505b919050565b60006000600060009050836040518082805190602001908083835b6020831061117a5780518252602082019150602081019050602083039250611157565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020856040518082805190602001908083835b602083106111db57805182526020820191506020810190506020830392506111b8565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020876040518082805190602001908083835b6020831061123c5780518252602082019150602081019050602083039250611219565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207f34328b5709ce0e15a7fe963f010fe4fb2b72713edfa4c63e0f24ec697c77b75c60405180905060405180910390a48560036000600454815260200190815260200160002060010190805190602001906112c4929190611a5d565b508460036000600454815260200190815260200160002060000190805190602001906112f1929190611a5d565b5083600360006004548152602001908152602001600020600201908051906020019061131e929190611a5d565b50600160046000828254019250508190555060036000600160045403815260200190815260200160002060000160405180828054600181600116156101000203166002900480156113a65780601f106113845761010080835404028352918201916113a6565b820191906000526020600020905b815481529060010190602001808311611392575b5050915050604051809103902060036000600160045403815260200190815260200160002060010160405180828054600181600116156101000203166002900480156114295780601f10611407576101008083540402835291820191611429565b820191906000526020600020905b815481529060010190602001808311611415575b505091505060405180910390206004547f73eb4df6129515e874adb10f5eddb8879bbc213d07740c99054de0a386a149a260036000600160045403815260200190815260200160002060020160405180806020018281038252838181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156114fe5780601f106114d3576101008083540402835291602001916114fe565b820191906000526020600020905b8154815290600101906020018083116114e157829003601f168201915b50509250505060405180910390a4846040518082805190602001908083835b60208310611540578051825260208201915060208101905060208303925061151d565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207f8fc1bb51d5ce7c93bfc953314351afc9223a322d093523de57d10d2cc68fdd4160405180905060405180910390a2600092505b60065483101561175257846040518082805190602001908083835b602083106115df57805182526020820191506020810190506020830392506115bc565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191660026000858152602001908152602001600020604051808280546001816001161561010002031660029004801561167e5780601f1061165c57610100808354040283529182019161167e565b820191906000526020600020905b81548152906001019060200180831161166a575b5050915050604051809103902060001916141561174457600190506002600084815260200190815260200160002060405180828054600181600116156101000203166002900480156117075780601f106116e5576101008083540402835291820191611707565b820191906000526020600020905b8154815290600101906020018083116116f3575b505091505060405180910390207f1785dad0ffb50908010ec5e6ae5a882ec193b0b21e039595443cde29a953be7b60405180905060405180910390a25b5b82806001019350506115a1565b600092505b60055483101561185e576001600084815260200190815260200160002060405180828054600181600116156101000203166002900480156117cf5780601f106117ad5761010080835404028352918201916117cf565b820191906000526020600020905b8154815290600101906020018083116117bb575b5050915050604051809103902060001916856040518082805190602001908083835b6020831061181457805182526020820191506020810190506020830392506117f1565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916141561185057600191505b5b8280600101935050611757565b60008114801561186e5750600182145b156119f157846040518082805190602001908083835b602083106118a75780518252602082019150602081019050602083039250611884565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207f31111fb0ab34fc149a6179f52431820d3929a52a20a79c346754117d42ebe2e060405180905060405180910390a284600260006006548152602001908152602001600020908051906020019061192c929190611a5d565b50600160066000828254019250508190555060026000600160065403815260200190815260200160002060405180828054600181600116156101000203166002900480156119b15780601f1061198f5761010080835404028352918201916119b1565b820191906000526020600020905b81548152906001019060200180831161199d575b505091505060405180910390206006547f8a8c6e5d817d6b4b89ece61a1307d0361795dbbf0ef54c2693409c669bf500a960405180905060405180910390a35b6005546006541415611a35576006600e819055507f0c8fd27f0e202d53a731dc39238816f51f38331d2e1c6d96ac4f40429a8a3a1a60405180905060405180910390a15b5b505050505050565b80600e819055505b50565b602060405190810160405280600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611a9e57805160ff1916838001178555611acc565b82800160010185558215611acc579182015b82811115611acb578251825591602001919060010190611ab0565b5b509050611ad99190611add565b5090565b611aff91905b80821115611afb576000816000905550600101611ae3565b5090565b905600a165627a7a72305820cf9471571ffc78c1d5c781b3336c4c091c930917d285cdae0284f0c632b6478c0029`

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _requestor common.Address, _size *big.Int, _name string, _date string) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend, _requestor, _size, _name, _date)
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

// GetCreatedDate is a free data retrieval call binding the contract method 0x28aea735.
//
// Solidity: function getCreatedDate() constant returns(string)
func (_Contract *ContractCaller) GetCreatedDate(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getCreatedDate")
	return *ret0, err
}

// GetCreatedDate is a free data retrieval call binding the contract method 0x28aea735.
//
// Solidity: function getCreatedDate() constant returns(string)
func (_Contract *ContractSession) GetCreatedDate() (string, error) {
	return _Contract.Contract.GetCreatedDate(&_Contract.CallOpts)
}

// GetCreatedDate is a free data retrieval call binding the contract method 0x28aea735.
//
// Solidity: function getCreatedDate() constant returns(string)
func (_Contract *ContractCallerSession) GetCreatedDate() (string, error) {
	return _Contract.Contract.GetCreatedDate(&_Contract.CallOpts)
}

// GetDeployedDate is a free data retrieval call binding the contract method 0x9d37c22a.
//
// Solidity: function getDeployedDate() constant returns(string)
func (_Contract *ContractCaller) GetDeployedDate(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getDeployedDate")
	return *ret0, err
}

// GetDeployedDate is a free data retrieval call binding the contract method 0x9d37c22a.
//
// Solidity: function getDeployedDate() constant returns(string)
func (_Contract *ContractSession) GetDeployedDate() (string, error) {
	return _Contract.Contract.GetDeployedDate(&_Contract.CallOpts)
}

// GetDeployedDate is a free data retrieval call binding the contract method 0x9d37c22a.
//
// Solidity: function getDeployedDate() constant returns(string)
func (_Contract *ContractCallerSession) GetDeployedDate() (string, error) {
	return _Contract.Contract.GetDeployedDate(&_Contract.CallOpts)
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

// SetDeployedDate is a paid mutator transaction binding the contract method 0x82756e7a.
//
// Solidity: function setDeployedDate(_date string) returns()
func (_Contract *ContractTransactor) SetDeployedDate(opts *bind.TransactOpts, _date string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setDeployedDate", _date)
}

// SetDeployedDate is a paid mutator transaction binding the contract method 0x82756e7a.
//
// Solidity: function setDeployedDate(_date string) returns()
func (_Contract *ContractSession) SetDeployedDate(_date string) (*types.Transaction, error) {
	return _Contract.Contract.SetDeployedDate(&_Contract.TransactOpts, _date)
}

// SetDeployedDate is a paid mutator transaction binding the contract method 0x82756e7a.
//
// Solidity: function setDeployedDate(_date string) returns()
func (_Contract *ContractTransactorSession) SetDeployedDate(_date string) (*types.Transaction, error) {
	return _Contract.Contract.SetDeployedDate(&_Contract.TransactOpts, _date)
}
