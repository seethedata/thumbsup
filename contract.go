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
const ContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getNumApprovers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCreatedDate\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getApprovedRole\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getApprover\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSpecification\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"setDeployedDate\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"addRequiredRole\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getApprovalsStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDeployedDate\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getApplicationName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequesterName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRequiredRole\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_role\",\"type\":\"string\"},{\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"changeStatus\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_requestor\",\"type\":\"address\"},{\"name\":\"_size\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_requesterName\",\"type\":\"string\"},{\"name\":\"_date\",\"type\":\"string\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"numReqApprovers\",\"type\":\"uint256\"}],\"name\":\"ApplicationCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_size\",\"type\":\"uint256\"}],\"name\":\"WrongApplicationSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Killed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"AddingRequiredRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"numReqApprover\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"RequiredRoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"CheckingRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"RoleAlreadyExists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"AddingRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"numApprovedRole\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"AddingApprover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"numApprovers\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"_role\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"ApproverAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ApplicationApproved\",\"type\":\"event\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
const ContractBin = `606060405234156200000d57fe5b60405162001fe638038062001fe6833981016040528080519060200190919080519060200190919080518201919060200180518201919060200180518201919050505b84600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160019080519060200190620000a9929190620002ac565b506000600581905550600060078190555082600c9080519060200190620000d2929190620002ac565b5060006006819055506001600f8190555080600d9080519060200190620000fb929190620002ac565b50604060405190810160405280600381526020017f4e2f410000000000000000000000000000000000000000000000000000000000815250600e90805190602001906200014a929190620002ac565b506001841415620001885760016008600001819055506004600860010181905550610400600860020181905550600160086003018190555062000235565b6002841415620001c55760026008600001819055506008600860010181905550610800600860020181905550600260086003018190555062000234565b6003841415620002025760046008600001819055506010600860010181905550611000600860020181905550600460086003018190555062000233565b837f1e371dab69817be62ab25b28af940aff33a252f63235a64a43e4ae5020c59f3960405180905060405180910390a25b5b5b600654600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f525a1f964d375777dfd5b80ed4ed417220ab8bcdb3cd53ce50e4435b7773f44c60405180905060405180910390a35b50505050506200035b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002ef57805160ff191683800117855562000320565b8280016001018555821562000320579182015b828111156200031f57825182559160200191906001019062000302565b5b5090506200032f919062000333565b5090565b6200035891905b80821115620003545760008160009055506001016200033a565b5090565b90565b611c7b806200036b6000396000f300606060405236156100ef576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806313abc4f4146100f157806328aea7351461011757806328b553f6146101b057806341c0e1b5146102575780634e69d560146102695780637157e4931461028f57806372e754581461042a57806382756e7a14610465578063946a0c06146104bf5780639addff57146105195780639d37c22a1461054d578063a3f5b1bf146105e6578063c0a7772d14610638578063c0c3887d146106d1578063d632e07f1461076a578063e7e1030a14610811578063e8025d77146108f1575bfe5b34156100f957fe5b610101610911565b6040518082815260200191505060405180910390f35b341561011f57fe5b61012761091c565b6040518080602001828103825283818151815260200191508051906020019080838360008314610176575b80518252602083111561017657602082019150602081019050602083039250610152565b505050905090810190601f1680156101a25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101b857fe5b6101ce60048080359060200190919050506109c5565b604051808060200182810382528381815181526020019150805190602001908083836000831461021d575b80518252602083111561021d576020820191506020810190506020830392506101f9565b505050905090810190601f1680156102495780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561025f57fe5b610267610a81565b005b341561027157fe5b610279610b7d565b6040518082815260200191505060405180910390f35b341561029757fe5b6102ad6004808035906020019091905050610b88565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360008314610304575b805182526020831115610304576020820191506020810190506020830392506102e0565b505050905090810190601f1680156103305780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360008314610378575b80518252602083111561037857602082019150602081019050602083039250610354565b505050905090810190601f1680156103a45780820380516001836020036101000a031916815260200191505b508481038252858181518152602001915080519060200190808383600083146103ec575b8051825260208311156103ec576020820191506020810190506020830392506103c8565b505050905090810190601f1680156104185780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b341561043257fe5b61043a610dc4565b6040518085815260200184815260200183815260200182815260200194505050505060405180910390f35b341561046d57fe5b6104bd600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610df3565b005b34156104c757fe5b610517600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610e0e565b005b341561052157fe5b610529610f8c565b60405180848152602001838152602001828152602001935050505060405180910390f35b341561055557fe5b61055d610fa7565b60405180806020018281038252838181518152602001915080519060200190808383600083146105ac575b8051825260208311156105ac57602082019150602081019050602083039250610588565b505050905090810190601f1680156105d85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156105ee57fe5b6105f6611050565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561064057fe5b61064861107b565b6040518080602001828103825283818151815260200191508051906020019080838360008314610697575b80518252602083111561069757602082019150602081019050602083039250610673565b505050905090810190601f1680156106c35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156106d957fe5b6106e1611124565b6040518080602001828103825283818151815260200191508051906020019080838360008314610730575b8051825260208311156107305760208201915060208101905060208303925061070c565b505050905090810190601f16801561075c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561077257fe5b61078860048080359060200190919050506111cd565b60405180806020018281038252838181518152602001915080519060200190808383600083146107d7575b8051825260208311156107d7576020820191506020810190506020830392506107b3565b505050905090810190601f1680156108035780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561081957fe5b6108ef600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050611289565b005b34156108f957fe5b61090f6004808035906020019091905050611b8b565b005b600060055490505b90565b610924611b96565b600d8054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109ba5780601f1061098f576101008083540402835291602001916109ba565b820191906000526020600020905b81548152906001019060200180831161099d57829003601f168201915b505050505090505b90565b6109cd611b96565b600360008381526020019081526020016000208054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a745780601f10610a4957610100808354040283529160200191610a74565b820191906000526020600020905b815481529060010190602001808311610a5757829003601f168201915b505050505090505b919050565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610b7a57600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f4b0bc4f25f8d0b92d2e12b686ba96cd75e4e69325e6cf7b1f3119d14eaf2cbdf60405180905060405180910390a2600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b6000600f5490505b90565b610b90611b96565b610b98611b96565b610ba0611b96565b600460008581526020019081526020016000206001016004600086815260200190815260200160002060000160046000878152602001908152602001600020600201828054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c775780601f10610c4c57610100808354040283529160200191610c77565b820191906000526020600020905b815481529060010190602001808311610c5a57829003601f168201915b50505050509250818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d135780601f10610ce857610100808354040283529160200191610d13565b820191906000526020600020905b815481529060010190602001808311610cf657829003601f168201915b50505050509150808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610daf5780601f10610d8457610100808354040283529160200191610daf565b820191906000526020600020905b815481529060010190602001808311610d9257829003601f168201915b505050505090509250925092505b9193909250565b600060006000600060086000015460086001015460086002015460086003015493509350935093505b90919293565b80600e9080519060200190610e09929190611baa565b505b50565b806040518082805190602001908083835b60208310610e425780518252602082019150602081019050602083039250610e1f565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207fba741c02698f40b54c299b0d26b447d5fb4065d70b357594397f1d8e4c294efc60405180905060405180910390a2806002600060065481526020019081526020016000209080519060200190610ec7929190611baa565b5060016006600082825401925050819055506002600060065481526020019081526020016000206040518082805460018160011615610100020316600290048015610f495780601f10610f27576101008083540402835291820191610f49565b820191906000526020600020905b815481529060010190602001808311610f35575b505091505060405180910390206006547ff4155a5029e40a4636a842980be6bf1be532789feb40549a24ddc8bcb255913560405180905060405180910390a35b50565b600060006000600f546006546007549250925092505b909192565b610faf611b96565b600e8054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110455780601f1061101a57610100808354040283529160200191611045565b820191906000526020600020905b81548152906001019060200180831161102857829003601f168201915b505050505090505b90565b6000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b90565b611083611b96565b600c8054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111195780601f106110ee57610100808354040283529160200191611119565b820191906000526020600020905b8154815290600101906020018083116110fc57829003601f168201915b505050505090505b90565b61112c611b96565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111c25780601f10611197576101008083540402835291602001916111c2565b820191906000526020600020905b8154815290600101906020018083116111a557829003601f168201915b505050505090505b90565b6111d5611b96565b600260008381526020019081526020016000208054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561127c5780601f106112515761010080835404028352916020019161127c565b820191906000526020600020905b81548152906001019060200180831161125f57829003601f168201915b505050505090505b919050565b60006000600060009050836040518082805190602001908083835b602083106112c757805182526020820191506020810190506020830392506112a4565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020856040518082805190602001908083835b602083106113285780518252602082019150602081019050602083039250611305565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020876040518082805190602001908083835b602083106113895780518252602082019150602081019050602083039250611366565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207f34328b5709ce0e15a7fe963f010fe4fb2b72713edfa4c63e0f24ec697c77b75c60405180905060405180910390a4856004600060055481526020019081526020016000206001019080519060200190611411929190611baa565b5084600460006005548152602001908152602001600020600001908051906020019061143e929190611baa565b5083600460006005548152602001908152602001600020600201908051906020019061146b929190611baa565b50600160056000828254019250508190555060046000600160055403815260200190815260200160002060000160405180828054600181600116156101000203166002900480156114f35780601f106114d15761010080835404028352918201916114f3565b820191906000526020600020905b8154815290600101906020018083116114df575b5050915050604051809103902060046000600160055403815260200190815260200160002060010160405180828054600181600116156101000203166002900480156115765780601f10611554576101008083540402835291820191611576565b820191906000526020600020905b815481529060010190602001808311611562575b505091505060405180910390206005547f73eb4df6129515e874adb10f5eddb8879bbc213d07740c99054de0a386a149a2600460006001600554038152602001908152602001600020600201604051808060200182810382528381815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561164b5780601f106116205761010080835404028352916020019161164b565b820191906000526020600020905b81548152906001019060200180831161162e57829003601f168201915b50509250505060405180910390a4846040518082805190602001908083835b6020831061168d578051825260208201915060208101905060208303925061166a565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207f8fc1bb51d5ce7c93bfc953314351afc9223a322d093523de57d10d2cc68fdd4160405180905060405180910390a2600092505b60075483101561189f57846040518082805190602001908083835b6020831061172c5780518252602082019150602081019050602083039250611709565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020600019166003600085815260200190815260200160002060405180828054600181600116156101000203166002900480156117cb5780601f106117a95761010080835404028352918201916117cb565b820191906000526020600020905b8154815290600101906020018083116117b7575b5050915050604051809103902060001916141561189157600190506003600084815260200190815260200160002060405180828054600181600116156101000203166002900480156118545780601f10611832576101008083540402835291820191611854565b820191906000526020600020905b815481529060010190602001808311611840575b505091505060405180910390207f1785dad0ffb50908010ec5e6ae5a882ec193b0b21e039595443cde29a953be7b60405180905060405180910390a25b5b82806001019350506116ee565b600092505b6006548310156119ab5760026000848152602001908152602001600020604051808280546001816001161561010002031660029004801561191c5780601f106118fa57610100808354040283529182019161191c565b820191906000526020600020905b815481529060010190602001808311611908575b5050915050604051809103902060001916856040518082805190602001908083835b60208310611961578051825260208201915060208101905060208303925061193e565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916141561199d57600191505b5b82806001019350506118a4565b6000811480156119bb5750600182145b15611b3e57846040518082805190602001908083835b602083106119f457805182526020820191506020810190506020830392506119d1565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207f31111fb0ab34fc149a6179f52431820d3929a52a20a79c346754117d42ebe2e060405180905060405180910390a2846003600060075481526020019081526020016000209080519060200190611a79929190611baa565b5060016007600082825401925050819055506003600060016007540381526020019081526020016000206040518082805460018160011615610100020316600290048015611afe5780601f10611adc576101008083540402835291820191611afe565b820191906000526020600020905b815481529060010190602001808311611aea575b505091505060405180910390206007547f8a8c6e5d817d6b4b89ece61a1307d0361795dbbf0ef54c2693409c669bf500a960405180905060405180910390a35b6006546007541415611b82576006600f819055507f0c8fd27f0e202d53a731dc39238816f51f38331d2e1c6d96ac4f40429a8a3a1a60405180905060405180910390a15b5b505050505050565b80600f819055505b50565b602060405190810160405280600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611beb57805160ff1916838001178555611c19565b82800160010185558215611c19579182015b82811115611c18578251825591602001919060010190611bfd565b5b509050611c269190611c2a565b5090565b611c4c91905b80821115611c48576000816000905550600101611c30565b5090565b905600a165627a7a723058205f42a87076099d8476700792997631b3291ddb001766ea8b96f7bcda8c591fff0029`

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _requestor common.Address, _size *big.Int, _name string, _requesterName string, _date string) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend, _requestor, _size, _name, _requesterName, _date)
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

// GetRequesterName is a free data retrieval call binding the contract method 0xc0c3887d.
//
// Solidity: function getRequesterName() constant returns(string)
func (_Contract *ContractCaller) GetRequesterName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getRequesterName")
	return *ret0, err
}

// GetRequesterName is a free data retrieval call binding the contract method 0xc0c3887d.
//
// Solidity: function getRequesterName() constant returns(string)
func (_Contract *ContractSession) GetRequesterName() (string, error) {
	return _Contract.Contract.GetRequesterName(&_Contract.CallOpts)
}

// GetRequesterName is a free data retrieval call binding the contract method 0xc0c3887d.
//
// Solidity: function getRequesterName() constant returns(string)
func (_Contract *ContractCallerSession) GetRequesterName() (string, error) {
	return _Contract.Contract.GetRequesterName(&_Contract.CallOpts)
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
