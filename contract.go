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

// ContractBin is the compiled bytecode used for deploying new contracts.
const ContractBin = `606060405234156200000d57fe5b604051620027f1380380620027f183398101604052808051906020019091908051906020019091908051820191906020018051906020019091505b60006000600c60016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826004819055506000600581905550600060068190555083600b9080519060200190620000fc9291906200031e565b506000600c60006101000a81548160ff021916908360068111156200011d57fe5b021790555060018514156200015f576001600760000181905550600460076001018190555061040060076002018190555060016007600301819055506200020c565b60028514156200019c576002600760000181905550600860076001018190555061080060076002018190555060026007600301819055506200020b565b6003851415620001d9576004600760000181905550601060076001018190555061100060076002018190555060046007600301819055506200020a565b847f1e371dab69817be62ab25b28af940aff33a252f63235a64a43e4ae5020c59f3960405180905060405180910390a25b5b5b600090505b6005811015620002a657600182826005811015156200022c57fe5b602002015114156200026a5760016001600083815260200190815260200160002060006101000a81548160ff02191690831515021790555062000297565b60006001600083815260200190815260200160002060006101000a81548160ff0219169083151502179055505b5b808060010191505062000211565b600454600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f525a1f964d375777dfd5b80ed4ed417220ab8bcdb3cd53ce50e4435b7773f44c60405180905060405180910390a35b505050505050620003cd565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200036157805160ff191683800117855562000392565b8280016001018555821562000392579182015b828111156200039157825182559160200191906001019062000374565b5b509050620003a19190620003a5565b5090565b620003ca91905b80821115620003c6576000816000905550600101620003ac565b5090565b90565b61241480620003dd6000396000f300606060405236156100c3576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806313abc4f4146100c557806341c0e1b5146100eb5780634e69d560146100fd57806356c98930146101235780637157e4931461016257806372e75458146101d057806381045ead1461020b5780639addff5714610231578063a3f5b1bf146102d1578063c0a7772d14610323578063daea85c5146103bc578063e8025d77146103f2578063f39ac26c14610412575bfe5b34156100cd57fe5b6100d5610438565b6040518082815260200191505060405180910390f35b34156100f357fe5b6100fb610443565b005b341561010557fe5b61010d61053f565b6040518082815260200191505060405180910390f35b341561012b57fe5b610160600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506106f5565b005b341561016a57fe5b6101806004808035906020019091905050610d5e565b604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390f35b34156101d857fe5b6101e0611302565b6040518085815260200184815260200183815260200182815260200194505050505060405180910390f35b341561021357fe5b61021b611331565b6040518082815260200191505060405180910390f35b341561023957fe5b61024161133c565b6040518084815260200183600560200280838360008314610281575b8051825260208311156102815760208201915060208101905060208303925061025d565b505050905001826005602002808383600083146102bd575b8051825260208311156102bd57602082019150602081019050602083039250610299565b505050905001935050505060405180910390f35b34156102d957fe5b6102e1611cc0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561032b57fe5b610333611ceb565b6040518080602001828103825283818151815260200191508051906020019080838360008314610382575b8051825260208311156103825760208201915060208101905060208303925061035e565b505050905090810190601f1680156103ae5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156103c457fe5b6103f0600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611d94565b005b34156103fa57fe5b610410600480803590602001909190505061223f565b005b341561041a57fe5b6104226123a1565b6040518082815260200191505060405180910390f35b600060045490505b90565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561053c57600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f4b0bc4f25f8d0b92d2e12b686ba96cd75e4e69325e6cf7b1f3119d14eaf2cbdf60405180905060405180910390a2600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b60007fabd95b950242a279866243fa2b8fec5adddf6560d4e1b4f8745cfe7b5778686560405180905060405180910390a16000600681111561057d57fe5b600c60009054906101000a900460ff16600681111561059857fe5b14156105a757600190506106f2565b600160068111156105b457fe5b600c60009054906101000a900460ff1660068111156105cf57fe5b14156105de57600290506106f2565b600260068111156105eb57fe5b600c60009054906101000a900460ff16600681111561060657fe5b141561061557600390506106f2565b6003600681111561062257fe5b600c60009054906101000a900460ff16600681111561063d57fe5b141561064c57600490506106f2565b6004600681111561065957fe5b600c60009054906101000a900460ff16600681111561067457fe5b141561068357600590506106f2565b6005600681111561069057fe5b600c60009054906101000a900460ff1660068111156106ab57fe5b14156106ba57600690506106f2565b600660068111156106c757fe5b600c60009054906101000a900460ff1660068111156106e257fe5b14156106f157600790506106f2565b5b90565b60006004546005548473ffffffffffffffffffffffffffffffffffffffff167f97e8bff8a03a26a59622e8062941ec74afd1386e9e2d758d7766f7865b1df84e60405180905060405180910390a46004546005541015156107a1576004546005548473ffffffffffffffffffffffffffffffffffffffff167f7e0048238e74c403f385c2e9ed10c7f7f03ef80e95f286851be2bf178a551c6960405180905060405180910390a4610d59565b6004546005548473ffffffffffffffffffffffffffffffffffffffff167f954e72672c47e61dd5ba275ec88f2c3938a2d5dc3e9146570a2b75d7f1ae67c560405180905060405180910390a4600090505b6005548110156108bc578273ffffffffffffffffffffffffffffffffffffffff166003600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156108ae578273ffffffffffffffffffffffffffffffffffffffff167f9cdeb1ccdd50e743424ffe1a001f68ec4adb08644d0570f50429f3a75a047a1660405180905060405180910390a2610d59565b5b80806001019150506107f2565b600182141561092e576000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160006101000a81548160ff0219169083600481111561092457fe5b0217905550610b46565b60028214156109a0576001600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160006101000a81548160ff0219169083600481111561099657fe5b0217905550610b45565b6003821415610a12576002600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160006101000a81548160ff02191690836004811115610a0857fe5b0217905550610b44565b6004821415610a84576003600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160006101000a81548160ff02191690836004811115610a7a57fe5b0217905550610b43565b6005821415610af6576004600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160006101000a81548160ff02191690836004811115610aec57fe5b0217905550610b42565b818373ffffffffffffffffffffffffffffffffffffffff167f3a539b0af306e72ec041ee49224212749c209e4161cfe4f88a6447c9a677e79860405180905060405180910390a3610d59565b5b5b5b5b600c60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610be7578273ffffffffffffffffffffffffffffffffffffffff167fd3db92f6a4b866377e026bbd2f5ce5d4b46a3a0685c6fe4ebe5dbb5608e3516860405180905060405180910390a2610d59565b818373ffffffffffffffffffffffffffffffffffffffff167f6afc4ffb48918cd100a7045c5ba97d916ce5ba365a5f3de81d5f04e92d5b7e2160405180905060405180910390a36000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001819055508260036000600554815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060016005540160058190555060045460055460036000600554815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f5aff60c7623a689c6af440f3311639cc7a64be7e50e1bac5a2075750f31fbdfa60405180905060405180910390a45b5b505050565b600060006000600060008510158015610d78575060045485105b1561129157847fc6e7839ade15d86b21e56625de344285a0fd99ac07770b53163856943f34810b60405180905060405180910390a260006004811115610dba57fe5b600260006003600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff166004811115610e4857fe5b1415610e5357600190505b60016004811115610e6057fe5b600260006003600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff166004811115610eee57fe5b1415610ef957600290505b60026004811115610f0657fe5b600260006003600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff166004811115610f9457fe5b1415610f9f57600390505b60036004811115610fac57fe5b600260006003600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff16600481111561103a57fe5b141561104557600490505b6004600481111561105257fe5b600260006003600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff1660048111156110e057fe5b14156110eb57600590505b80600260006003600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001546003600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fed5eb1f7ae9e039c4ded5610848182949b38fc690881e93928a595ab86199e0f60405180905060405180910390a46003600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260006003600089815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154829350935093506112fa565b847fcc0b9efa377244dd0085718fc3aee49a1d3b58bd8aaa9656e77c33581b1b406a60405180905060405180910390a2600c60019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600060008191508090509350935093506112fa565b5b509193909250565b600060006000600060076000015460076001015460076002015460076003015493509350935093505b90919293565b600060055490505b90565b60006113466123ac565b61134e6123ac565b60006000600060006000600681111561136357fe5b600c60009054906101000a900460ff16600681111561137e57fe5b141561138957600192505b6001600681111561139657fe5b600c60009054906101000a900460ff1660068111156113b157fe5b14156113bc57600292505b600260068111156113c957fe5b600c60009054906101000a900460ff1660068111156113e457fe5b14156113ef57600392505b600360068111156113fc57fe5b600c60009054906101000a900460ff16600681111561141757fe5b141561142257600492505b6004600681111561142f57fe5b600c60009054906101000a900460ff16600681111561144a57fe5b141561145557600592505b6005600681111561146257fe5b600c60009054906101000a900460ff16600681111561147d57fe5b141561148857600692505b6006600681111561149557fe5b600c60009054906101000a900460ff1660068111156114b057fe5b14156114bb57600792505b600093505b600584101561153c57600115156001600086815260200190815260200160002060009054906101000a900460ff1615151415611514576001828560058110151561150657fe5b0160005b508190555061152e565b6000828560058110151561152457fe5b0160005b50819055505b5b83806001019450506114c0565b600093505b600454841015611c2f576000600481111561155857fe5b600260006003600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff1660048111156115e657fe5b14156116a9576001600260006003600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154141561168957600181600060058110151561167b57fe5b0160005b50819055506116a4565b600081600060058110151561169a57fe5b0160005b50819055505b611c21565b600160048111156116b657fe5b600260006003600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff16600481111561174457fe5b1415611807576001600260006003600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015414156117e75760018160016005811015156117d957fe5b0160005b5081905550611802565b60008160016005811015156117f857fe5b0160005b50819055505b611c20565b6002600481111561181457fe5b600260006003600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff1660048111156118a257fe5b1415611965576001600260006003600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154141561194557600181600260058110151561193757fe5b0160005b5081905550611960565b600081600260058110151561195657fe5b0160005b50819055505b611c1f565b6003600481111561197257fe5b600260006003600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff166004811115611a0057fe5b1415611ac3576001600260006003600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541415611aa3576001816003600581101515611a9557fe5b0160005b5081905550611abe565b6000816003600581101515611ab457fe5b0160005b50819055505b611c1e565b60046004811115611ad057fe5b600260006003600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff166004811115611b5e57fe5b1415611c1d576001600260006003600088815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541415611c01576001816004600581101515611bf357fe5b0160005b5081905550611c1c565b6000816004600581101515611c1257fe5b0160005b50819055505b5b5b5b5b5b5b8380600101945050611541565b82828281600580602002604051908101604052809291908260058015611c6a576020028201915b815481526020019060010190808311611c56575b5050505050915080600580602002604051908101604052809291908260058015611ca9576020028201915b815481526020019060010190808311611c95575b505050505090509650965096505b50505050909192565b6000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b90565b611cf36123d4565b600b8054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611d895780601f10611d5e57610100808354040283529160200191611d89565b820191906000526020600020905b815481529060010190602001808311611d6c57829003601f168201915b505050505090505b90565b60008173ffffffffffffffffffffffffffffffffffffffff167fdcef5e0f5751f0caf9d03f730740c6ee8f1ca13a020aa793159f4e424c01834f60405180905060405180910390a2600090505b6005548110156121cd576003600082815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f63e5403186eae3a0d060f2b3262d3c505e0a579c22af082afcbe6f62cb6162a460405180905060405180910390a28173ffffffffffffffffffffffffffffffffffffffff166003600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156121bf57806003600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f37b5f5c1aa9e5c018243215c9a3865612c2845a8b1250ad55cd10f4c6fcbf4f860405180905060405180910390a36000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154141561214c576001600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555060016006600082825401925050819055506001600d819055506003600082815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f03a44b14db2ddab320f44e944b9d6ca4163de2f5966ead9de6602eea49c912c160405180905060405180910390a27fe8672c12b66c4cabe8a1d452589f87ad1ad4768cd939311d38eb68db752437e060405180905060405180910390a16004546006541415612105577f0c8fd27f0e202d53a731dc39238816f51f38331d2e1c6d96ac4f40429a8a3a1a60405180905060405180910390a16005600c60006101000a81548160ff021916908360068111156120f357fe5b02179055506002600d8190555061223b565b6004546006547ff6113fea8672aed076e71fdfd14042695c6f29b39c69a9e3b278921b6fc8bcec60405180905060405180910390a36003600d8190555061223b565b61223b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600d819055508173ffffffffffffffffffffffffffffffffffffffff167f164da0e49bdae0061ef3c3a64293025cf7b355fac5b1ed2d995f51cf00e7222660405180905060405180910390a261223b565b5b5b8080600101915050611de1565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe600d819055508173ffffffffffffffffffffffffffffffffffffffff167fbafd0160b32fc86793c9763b4191509af9a94d27b1c0f523fdeb5b2cba7e14ab60405180905060405180910390a25b5050565b6001811415612271576000600c60006101000a81548160ff0219169083600681111561226757fe5b021790555061239e565b60028114156122a3576001600c60006101000a81548160ff0219169083600681111561229957fe5b021790555061239e565b60038114156122d5576002600c60006101000a81548160ff021916908360068111156122cb57fe5b021790555061239e565b6004811415612307576003600c60006101000a81548160ff021916908360068111156122fd57fe5b021790555061239e565b6005811415612339576004600c60006101000a81548160ff0219169083600681111561232f57fe5b021790555061239e565b600681141561236b576005600c60006101000a81548160ff0219169083600681111561236157fe5b021790555061239e565b600781141561239d576006600c60006101000a81548160ff0219169083600681111561239357fe5b021790555061239e565b5b50565b6000600d5490505b90565b60a0604051908101604052806005905b60008152602001906001900390816123bc5790505090565b6020604051908101604052806000815250905600a165627a7a723058209417cc10c3d1bd99e77cb918bc328d1c6eaa74c3b25bf5cc38b309266e63524f0029`

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _requestor common.Address, _size *big.Int, _name string, _numApprovers *big.Int, _reqApprovers []string) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend, _requestor, _size, _name, _numApprovers, _reqApprovers)
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
// Solidity: function addApprover(_approver address, _role string) returns()
func (_Contract *ContractTransactor) AddApprover(opts *bind.TransactOpts, _approver common.Address, _role string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addApprover", _approver, _role)
}

// AddApprover is a paid mutator transaction binding the contract method 0x56c98930.
//
// Solidity: function addApprover(_approver address, _role string) returns()
func (_Contract *ContractSession) AddApprover(_approver common.Address, _role string) (*types.Transaction, error) {
	return _Contract.Contract.AddApprover(&_Contract.TransactOpts, _approver, _role)
}

// AddApprover is a paid mutator transaction binding the contract method 0x56c98930.
//
// Solidity: function addApprover(_approver address, _role string) returns()
func (_Contract *ContractTransactorSession) AddApprover(_approver common.Address, _role string) (*types.Transaction, error) {
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
