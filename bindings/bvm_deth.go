// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// DETHInit is an auto generated low-level Go binding around an user-defined struct.
type DETHInit struct {
	Admin                  common.Address
	L2ShareAddress         common.Address
	Staking                common.Address
	UnstakeRequestsManager common.Address
}

// IDETHBatchMint is an auto generated low-level Go binding around an user-defined struct.
type IDETHBatchMint struct {
	Staker common.Address
	Amount *big.Int
}

// DETHMetaData contains all meta data concerning the DETH contract.
var DETHMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotL2ShareAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotStakingManagerContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotUnstakeRequestsManagerContract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIDETH.BatchMint[]\",\"name\":\"batcher\",\"type\":\"tuple[]\"}],\"name\":\"batchMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2ShareAddress\",\"type\":\"address\"},{\"internalType\":\"contractIStakingManager\",\"name\":\"staking\",\"type\":\"address\"},{\"internalType\":\"contractIUnstakeRequestsManager\",\"name\":\"unstakeRequestsManager\",\"type\":\"address\"}],\"internalType\":\"structDETH.Init\",\"name\":\"init\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ShareAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingContract\",\"outputs\":[{\"internalType\":\"contractIStakingManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeRequestsManagerContract\",\"outputs\":[{\"internalType\":\"contractIUnstakeRequestsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608080604052346100b8577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166100a957506001600160401b036002600160401b031982821601610064575b60405161236b90816100be8239f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1388080610055565b63f92ee8a960e01b8152600490fd5b600080fdfe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a71461168a575080630415cc1514610d5b57806306fdde0314610cb1578063095ea7b314610c8b5780630bec088614610c6257806318160ddd14610c2557806323b872dd14610b40578063248a9ca314610b045780632f2ff15d14610a82578063313ce56714610a665780633644e51514610a4357806336568abe146109fc57806340c10f19146109b257806342966c68146108a557806370a082311461085e5780637ecebe001461080557806384b0196e146106a95780639010d07c1461065557806391d14854146105fb5780639386e1971461054257806395d89b411461045d578063a217fddf14610441578063a9059cbb14610410578063ca15c873146103d7578063d505accf1461024c578063d547741f146101fe578063dd62ed3e146101b5578063ee99205c1461018c5763f14a3f501461015e57600080fd5b34610187576000366003190112610187576001546040516001600160a01b039091168152602090f35b600080fd5b34610187576000366003190112610187576000546040516001600160a01b039091168152602090f35b34610187576040366003190112610187576101ce611733565b6101df6101d9611749565b9161193c565b9060018060a01b03166000526020526020604060002054604051908152f35b346101875760403660031901126101875761024a60043561021d611749565b90806000526000805160206122b6833981519152602052610245600160406000200154611a70565b611ac1565b005b346101875760e036600319011261018757610265611733565b61026d611749565b60443590606435926084359360ff85168503610187578042116103bf5760018060a01b039081831695866000527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040600020908154916001830190556040519260208401927f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98452896040860152858816606086015288608086015260a085015260c084015260c0835260e08301918383106001600160401b038411176103a9576103789361036f93604052519020610348611d67565b906040519161190160f01b83526002830152602282015260c43591604260a4359220611e77565b90929192611f07565b1684810361038b575061024a9350611d09565b60449085604051916325c0072360e11b835260048301526024820152fd5b634e487b7160e01b600052604160045260246000fd5b6024906040519063313c898160e11b82526004820152fd5b34610187576020366003190112610187576004356000526000805160206121d68339815191526020526020604060002054604051908152f35b346101875760403660031901126101875761043661042c611733565b602435903361199b565b602060405160018152f35b3461018757600036600319011261018757602060405160008152f35b3461018757600036600319011261018757604051600080516020612256833981519152805482600061048e8361179f565b92838352602094600191866001821691826000146105205750506001146104d6575b50506104be92500383611712565b6104d260405192828493845283019061175f565b0390f35b8592506000526000805160206122d6833981519152906000915b8583106105085750506104be935082010185806104b0565b805483890185015287945086939092019181016104f0565b92509350506104be94915060ff191682840152151560051b82010185806104b0565b3461018757602080600319360112610187576001600160401b036004358181116101875736602382011215610187578060040135918211610187576024810190602436918460061b010111610187576001805490936001600160a01b0391821633036105e95760005b8481106105b457005b6105bf818686611975565b35908382168203610187576105e38792846105db848a8a611975565b013590611b09565b016105ab565b60405163baa183a760e01b8152600490fd5b3461018757604036600319011261018757610614611749565b6004356000526000805160206122b683398151915260205260406000209060018060a01b0316600052602052602060ff604060002054166040519015158152f35b34610187576040366003190112610187576004356000526000805160206121d683398151915260205260206106906024356040600020611e5f565b905460405160039290921b1c6001600160a01b03168152f35b34610187576000366003190112610187577fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1005415806107dc575b1561079f576106f06117d9565b6106f861189f565b6040516020808201928284106001600160401b038511176103a9579160206107528594610744979660405260008452604051978897600f60f81b895260e0858a015260e089019061175f565b90878203604089015261175f565b91466060870152306080870152600060a087015285830360c0870152519182815201929160005b82811061078857505050500390f35b835185528695509381019392810192600101610779565b60405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b6044820152606490fd5b507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10154156106e3565b34610187576020366003190112610187576001600160a01b03610826611733565b166000527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526020604060002054604051908152f35b34610187576020366003190112610187576001600160a01b0361087f611733565b166000526000805160206122168339815191526020526020604060002054604051908152f35b3461018757602036600319011261018757600254600435906001600160a01b031633036109a05733156109875733600052600080516020612216833981519152806020526040600020549082821061095e579180916000933385526020520360408320557f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace028181540390556040519081527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60203392a3005b5060405163391434e360e21b815233600482015260248101919091526044810191909152606490fd5b604051634b637e8f60e11b815260006004820152602490fd5b60405163359ae1d360e11b8152600490fd5b34610187576040366003190112610187576109cb611733565b6000546001600160a01b031633036109ea5761024a9060243590611b09565b604051631cec3cdb60e31b8152600490fd5b3461018757604036600319011261018757610a15611749565b336001600160a01b03821603610a315761024a90600435611ac1565b60405163334bd91960e11b8152600490fd5b34610187576000366003190112610187576020610a5e611d67565b604051908152f35b3461018757600036600319011261018757602060405160128152f35b3461018757604036600319011261018757600435610a9e611749565b816000526000805160206122b6833981519152602052610ac5600160406000200154611a70565b610acf8183611c7c565b610ad557005b60009182526000805160206121d6833981519152602052604090912061024a916001600160a01b031690611f8c565b34610187576020366003190112610187576004356000526000805160206122b68339815191526020526020600160406000200154604051908152f35b3461018757606036600319011261018757610b59611733565b610b61611749565b60443590610b6e8361193c565b33600052602052604060002054926000198403610b90575b610436935061199b565b828410610bff576001600160a01b03811615610be6573315610bcd578261043694610bba8361193c565b3360005260205203604060002055610b86565b604051634a1406b160e11b815260006004820152602490fd5b60405163e602df0560e01b815260006004820152602490fd5b604051637dc7a0d960e11b81523360048201526024810185905260448101849052606490fd5b346101875760003660031901126101875760207f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0254604051908152f35b34610187576000366003190112610187576002546040516001600160a01b039091168152602090f35b3461018757604036600319011261018757610436610ca7611733565b6024359033611d09565b34610187576000366003190112610187576040516000805160206121f68339815191528054826000610ce28361179f565b9283835260209460019186600182169182600014610520575050600114610d115750506104be92500383611712565b8592506000526000805160206121b6833981519152906000915b858310610d435750506104be935082010185806104b0565b80548389018501528794508693909201918101610d2b565b3461018757608036600319011261018757604051608081018181106001600160401b038211176103a957604052610d90611733565b8152610d9a611749565b60208201526044356001600160a01b03811681036101875760408201526064356001600160a01b03811681036101875760608201526000805160206122f683398151915254906001600160401b0382168015908161167a575b6001149081611670575b159081611667575b506116555760016001600160401b03198316176000805160206122f68339815191525560ff8260401c1615611628575b610e3d611bad565b60405191610e4a836116f7565b60048352630c88aa8960e31b602084015260405191610e68836116f7565b60048352630c88aa8960e31b6020840152610e81611bad565b610e89611bad565b83516001600160401b0381116103a957610eb16000805160206121f68339815191525461179f565b601f81116115b2575b506020601f8211600114611523578192939495600092611518575b50508160011b916000199060031b1c1916176000805160206121f6833981519152555b82516001600160401b0381116103a957610f206000805160206122568339815191525461179f565b601f81116114a2575b506020601f82116001146114145781929394600092611409575b50508160011b916000199060031b1c191617600080516020612256833981519152555b60405191610f73836116f7565b60048352630c88aa8960e31b6020840152610f8c611bad565b60405191610f99836116f7565b60018352603160f81b6020840152610faf611bad565b83516001600160401b0381116103a957610fd76000805160206122368339815191525461179f565b601f8111611393575b506020601f82116001146113045781929394956000926112f9575b50508160011b916000199060031b1c191617600080516020612236833981519152555b8251926001600160401b0384116103a9576110476000805160206122768339815191525461179f565b601f8111611283575b50602090601f85116001146111f15760ff9491600091836111e6575b50508160011b916000199060031b1c191617600080516020612276833981519152555b60007fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1008190557fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1015580516001600160a01b03166110eb81611bdc565b61119f575b50604081810151600080546001600160a01b03199081166001600160a01b03938416179091556020840151600180548316918416919091179055606090930151600280549094169116179091551c161561114657005b68ff0000000000000000196000805160206122f683398151915254166000805160206122f6833981519152557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b600080526000805160206121d68339815191526020526111df907f615f0f9e84155bea8cc509fe18befeb1baf65611e38a6ba60964480fb29dfd44611f8c565b50836110f0565b01519050858061106c565b906000805160206122768339815191526000526000805160206123168339815191529160005b601f198716811061126b575091859160019360ff97601f19811610611252575b505050811b016000805160206122768339815191525561108f565b015160001960f88460031b161c19169055858080611237565b91926020600181928685015181550194019201611217565b600080516020612276833981519152600052601f850160051c60008051602061231683398151915201602086106112e4575b601f820160051c6000805160206123168339815191520181106112d85750611050565b600081556001016112b5565b506000805160206123168339815191526112b5565b015190508580610ffb565b6000805160206122368339815191526000526000805160206122968339815191529060005b601f198416811061137b575060019394959683601f19811610611362575b505050811b016000805160206122368339815191525561101e565b015160001960f88460031b161c19169055858080611347565b9091602060018192858b015181550193019101611329565b600080516020612236833981519152600052601f820160051c60008051602061229683398151915201602083106113f4575b601f820160051c6000805160206122968339815191520181106113e85750610fe0565b600081556001016113c5565b506000805160206122968339815191526113c5565b015190508480610f43565b6000805160206122568339815191526000526000805160206122d68339815191529060005b601f198416811061148a5750600193949583601f19811610611471575b505050811b0160008051602061225683398151915255610f66565b015160001960f88460031b161c19169055848080611456565b9091602060018192858a015181550193019101611439565b600080516020612256833981519152600052601f820160051c6000805160206122d68339815191520160208310611503575b601f820160051c6000805160206122d68339815191520181106114f75750610f29565b600081556001016114d4565b506000805160206122d68339815191526114d4565b015190508580610ed5565b6000805160206121f68339815191526000526000805160206121b68339815191529060005b601f198416811061159a575060019394959683601f19811610611581575b505050811b016000805160206121f683398151915255610ef8565b015160001960f88460031b161c19169055858080611566565b9091602060018192858b015181550193019101611548565b6000805160206121f6833981519152600052601f820160051c6000805160206121b68339815191520160208310611613575b601f820160051c6000805160206121b68339815191520181106116075750610eba565b600081556001016115e4565b506000805160206121b68339815191526115e4565b68ffffffffffffffffff19821668010000000000000001176000805160206122f683398151915255610e35565b60405163f92ee8a960e01b8152600490fd5b90501583610e05565b303b159150610dfd565b604084901c60ff16159150610df3565b34610187576020366003190112610187576004359063ffffffff60e01b821680920361018757602091635a05180f60e01b81149081156116cc575b5015158152f35b637965db0b60e01b8114915081156116e6575b50836116c5565b6301ffc9a760e01b149050836116df565b604081019081106001600160401b038211176103a957604052565b90601f801991011681019081106001600160401b038211176103a957604052565b600435906001600160a01b038216820361018757565b602435906001600160a01b038216820361018757565b919082519283825260005b84811061178b575050826000602080949584010152601f8019910116010190565b60208183018101518483018201520161176a565b90600182811c921680156117cf575b60208310146117b957565b634e487b7160e01b600052602260045260246000fd5b91607f16916117ae565b60405190600082600080516020612236833981519152918254926117fc8461179f565b80845293602091600191828116908115611879575060011461182a575b50505061182892500383611712565b565b60009081526000805160206122968339815191529590935091905b8284106118615750611828945050508101602001388080611819565b85548885018301529485019487945092810192611845565b925050506020925061182894915060ff191682840152151560051b820101388080611819565b60405190600082600080516020612276833981519152918254926118c28461179f565b8084529360209160019182811690811561187957506001146118ed5750505061182892500383611712565b60009081526000805160206123168339815191529590935091905b8284106119245750611828945050508101602001388080611819565b85548885018301529485019487945092810192611908565b6001600160a01b031660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020526040902090565b91908110156119855760061b0190565b634e487b7160e01b600052603260045260246000fd5b916001600160a01b038084169283156109875716928315611a575760009083825260008051602061221683398151915280602052604083205491848310611a24575082847fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef959360409388602097528652038282205586815220818154019055604051908152a3565b60405163391434e360e21b81526001600160a01b0391909116600482015260248101929092525060448101839052606490fd5b60405163ec442f0560e01b815260006004820152602490fd5b806000526000805160206122b683398151915260205260406000203360005260205260ff6040600020541615611aa35750565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b611acb8282611ddc565b9182611ad657505090565b60009182526000805160206121d68339815191526020526040909120611b05916001600160a01b0316906120c5565b5090565b6001600160a01b0316908115611a57577f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02805490828201809211611b97576000927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9260209255848452600080516020612216833981519152825260408420818154019055604051908152a3565b634e487b7160e01b600052601160045260246000fd5b60ff6000805160206122f68339815191525460401c1615611bca57565b604051631afcd79f60e31b8152600490fd5b6001600160a01b031660008181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260408120549091906000805160206122b68339815191529060ff16611c77578280526020526040822081835260205260408220600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b505090565b906000918083526000805160206122b683398151915280602052604084209260018060a01b03169283855260205260ff60408520541615600014611d03578184526020526040832082845260205260408320600160ff198254161790557f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d339380a4600190565b50505090565b916001600160a01b03808416928315610be65716928315610bcd577f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591611d5160209261193c565b85600052825280604060002055604051908152a3565b611d6f612016565b611d77612080565b6040519060208201927f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8452604083015260608201524660808201523060a082015260a0815260c081018181106001600160401b038211176103a95760405251902090565b906000918083526000805160206122b683398151915280602052604084209260018060a01b03169283855260205260ff604085205416600014611d0357818452602052604083208284526020526040832060ff1981541690557ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b339380a4600190565b80548210156119855760005260206000200190600090565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411611efb57926020929160ff608095604051948552168484015260408301526060820152600092839182805260015afa15611eef5780516001600160a01b03811615611ee657918190565b50809160019190565b604051903d90823e3d90fd5b50505060009160039190565b6004811015611f765780611f19575050565b60018103611f335760405163f645eedf60e01b8152600490fd5b60028103611f545760405163fce698f760e01b815260048101839052602490fd5b600314611f5e5750565b602490604051906335e2f38360e21b82526004820152fd5b634e487b7160e01b600052602160045260246000fd5b919060018301600090828252806020526040822054156000146120105784549468010000000000000000861015611ffc5783611fec611fd5886001604098999a01855584611e5f565b819391549060031b91821b91600019901b19161790565b9055549382526020522055600190565b634e487b7160e01b83526041600452602483fd5b50925050565b61201e6117d9565b805190811561202e576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10054801561205b5790565b507fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47090565b61208861189f565b8051908115612098576020012090565b50507fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10154801561205b5790565b906001820190600092818452826020526040842054908115156000146121ae576000199180830181811161219a5782549084820191821161218657818103612151575b5050508054801561213d578201916121208383611e5f565b909182549160031b1b191690555582526020526040812055600190565b634e487b7160e01b86526031600452602486fd5b612171612161611fd59386611e5f565b90549060031b1c92839286611e5f565b90558652846020526040862055388080612108565b634e487b7160e01b88526011600452602488fd5b634e487b7160e01b87526011600452602487fd5b505050509056fe2ae08a8e29253f69ac5d979a101956ab8f8d9d7ded63fa7a83b16fc47648eab0c1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200052c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0352c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10252c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace04a16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10342ad5d3e1f2e6e70edcf6d991b8a3023d3fca8047a131592f9edb9fd9b89d57d02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680046a2803e59a4de4e7a4c574b1243f25977ac4c77d5a1a4a609b5394cebb4a2aaf0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f9ce34815f8e11431c7bb75a8e6886a91478f7ffc1dbb0a98dc240fddd76b75a26469706673582212206347641d11a423645553e2b200c4c2e9f3fb46f5af76e35764fc5952af0049ea64736f6c63430008180033",
}

// DETHABI is the input ABI used to generate the binding from.
// Deprecated: Use DETHMetaData.ABI instead.
var DETHABI = DETHMetaData.ABI

// DETHBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DETHMetaData.Bin instead.
var DETHBin = DETHMetaData.Bin

// DeployDETH deploys a new Ethereum contract, binding an instance of DETH to it.
func DeployDETH(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DETH, error) {
	parsed, err := DETHMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DETHBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DETH{DETHCaller: DETHCaller{contract: contract}, DETHTransactor: DETHTransactor{contract: contract}, DETHFilterer: DETHFilterer{contract: contract}}, nil
}

// DETH is an auto generated Go binding around an Ethereum contract.
type DETH struct {
	DETHCaller     // Read-only binding to the contract
	DETHTransactor // Write-only binding to the contract
	DETHFilterer   // Log filterer for contract events
}

// DETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type DETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DETHSession struct {
	Contract     *DETH             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DETHCallerSession struct {
	Contract *DETHCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DETHTransactorSession struct {
	Contract     *DETHTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type DETHRaw struct {
	Contract *DETH // Generic contract binding to access the raw methods on
}

// DETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DETHCallerRaw struct {
	Contract *DETHCaller // Generic read-only contract binding to access the raw methods on
}

// DETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DETHTransactorRaw struct {
	Contract *DETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDETH creates a new instance of DETH, bound to a specific deployed contract.
func NewDETH(address common.Address, backend bind.ContractBackend) (*DETH, error) {
	contract, err := bindDETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DETH{DETHCaller: DETHCaller{contract: contract}, DETHTransactor: DETHTransactor{contract: contract}, DETHFilterer: DETHFilterer{contract: contract}}, nil
}

// NewDETHCaller creates a new read-only instance of DETH, bound to a specific deployed contract.
func NewDETHCaller(address common.Address, caller bind.ContractCaller) (*DETHCaller, error) {
	contract, err := bindDETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DETHCaller{contract: contract}, nil
}

// NewDETHTransactor creates a new write-only instance of DETH, bound to a specific deployed contract.
func NewDETHTransactor(address common.Address, transactor bind.ContractTransactor) (*DETHTransactor, error) {
	contract, err := bindDETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DETHTransactor{contract: contract}, nil
}

// NewDETHFilterer creates a new log filterer instance of DETH, bound to a specific deployed contract.
func NewDETHFilterer(address common.Address, filterer bind.ContractFilterer) (*DETHFilterer, error) {
	contract, err := bindDETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DETHFilterer{contract: contract}, nil
}

// bindDETH binds a generic wrapper to an already deployed contract.
func bindDETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DETHMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DETH *DETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DETH.Contract.DETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DETH *DETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DETH.Contract.DETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DETH *DETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DETH.Contract.DETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DETH *DETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DETH *DETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DETH *DETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DETH.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DETH *DETHCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DETH *DETHSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DETH.Contract.DEFAULTADMINROLE(&_DETH.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DETH *DETHCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DETH.Contract.DEFAULTADMINROLE(&_DETH.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DETH *DETHCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DETH *DETHSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _DETH.Contract.DOMAINSEPARATOR(&_DETH.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_DETH *DETHCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _DETH.Contract.DOMAINSEPARATOR(&_DETH.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DETH *DETHCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DETH *DETHSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DETH.Contract.Allowance(&_DETH.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_DETH *DETHCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DETH.Contract.Allowance(&_DETH.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DETH *DETHCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DETH *DETHSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _DETH.Contract.BalanceOf(&_DETH.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_DETH *DETHCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _DETH.Contract.BalanceOf(&_DETH.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DETH *DETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DETH *DETHSession) Decimals() (uint8, error) {
	return _DETH.Contract.Decimals(&_DETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DETH *DETHCallerSession) Decimals() (uint8, error) {
	return _DETH.Contract.Decimals(&_DETH.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_DETH *DETHCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_DETH *DETHSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _DETH.Contract.Eip712Domain(&_DETH.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_DETH *DETHCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _DETH.Contract.Eip712Domain(&_DETH.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DETH *DETHCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DETH *DETHSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DETH.Contract.GetRoleAdmin(&_DETH.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DETH *DETHCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DETH.Contract.GetRoleAdmin(&_DETH.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_DETH *DETHCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_DETH *DETHSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _DETH.Contract.GetRoleMember(&_DETH.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_DETH *DETHCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _DETH.Contract.GetRoleMember(&_DETH.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_DETH *DETHCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_DETH *DETHSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _DETH.Contract.GetRoleMemberCount(&_DETH.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_DETH *DETHCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _DETH.Contract.GetRoleMemberCount(&_DETH.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DETH *DETHCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DETH *DETHSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DETH.Contract.HasRole(&_DETH.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DETH *DETHCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DETH.Contract.HasRole(&_DETH.CallOpts, role, account)
}

// L2ShareAddress is a free data retrieval call binding the contract method 0xf14a3f50.
//
// Solidity: function l2ShareAddress() view returns(address)
func (_DETH *DETHCaller) L2ShareAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "l2ShareAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ShareAddress is a free data retrieval call binding the contract method 0xf14a3f50.
//
// Solidity: function l2ShareAddress() view returns(address)
func (_DETH *DETHSession) L2ShareAddress() (common.Address, error) {
	return _DETH.Contract.L2ShareAddress(&_DETH.CallOpts)
}

// L2ShareAddress is a free data retrieval call binding the contract method 0xf14a3f50.
//
// Solidity: function l2ShareAddress() view returns(address)
func (_DETH *DETHCallerSession) L2ShareAddress() (common.Address, error) {
	return _DETH.Contract.L2ShareAddress(&_DETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DETH *DETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DETH *DETHSession) Name() (string, error) {
	return _DETH.Contract.Name(&_DETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DETH *DETHCallerSession) Name() (string, error) {
	return _DETH.Contract.Name(&_DETH.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_DETH *DETHCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_DETH *DETHSession) Nonces(owner common.Address) (*big.Int, error) {
	return _DETH.Contract.Nonces(&_DETH.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_DETH *DETHCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _DETH.Contract.Nonces(&_DETH.CallOpts, owner)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_DETH *DETHCaller) StakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "stakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_DETH *DETHSession) StakingContract() (common.Address, error) {
	return _DETH.Contract.StakingContract(&_DETH.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_DETH *DETHCallerSession) StakingContract() (common.Address, error) {
	return _DETH.Contract.StakingContract(&_DETH.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DETH *DETHCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DETH *DETHSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DETH.Contract.SupportsInterface(&_DETH.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DETH *DETHCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DETH.Contract.SupportsInterface(&_DETH.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DETH *DETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DETH *DETHSession) Symbol() (string, error) {
	return _DETH.Contract.Symbol(&_DETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DETH *DETHCallerSession) Symbol() (string, error) {
	return _DETH.Contract.Symbol(&_DETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DETH *DETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DETH *DETHSession) TotalSupply() (*big.Int, error) {
	return _DETH.Contract.TotalSupply(&_DETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DETH *DETHCallerSession) TotalSupply() (*big.Int, error) {
	return _DETH.Contract.TotalSupply(&_DETH.CallOpts)
}

// UnstakeRequestsManagerContract is a free data retrieval call binding the contract method 0x0bec0886.
//
// Solidity: function unstakeRequestsManagerContract() view returns(address)
func (_DETH *DETHCaller) UnstakeRequestsManagerContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DETH.contract.Call(opts, &out, "unstakeRequestsManagerContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnstakeRequestsManagerContract is a free data retrieval call binding the contract method 0x0bec0886.
//
// Solidity: function unstakeRequestsManagerContract() view returns(address)
func (_DETH *DETHSession) UnstakeRequestsManagerContract() (common.Address, error) {
	return _DETH.Contract.UnstakeRequestsManagerContract(&_DETH.CallOpts)
}

// UnstakeRequestsManagerContract is a free data retrieval call binding the contract method 0x0bec0886.
//
// Solidity: function unstakeRequestsManagerContract() view returns(address)
func (_DETH *DETHCallerSession) UnstakeRequestsManagerContract() (common.Address, error) {
	return _DETH.Contract.UnstakeRequestsManagerContract(&_DETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_DETH *DETHTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _DETH.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_DETH *DETHSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _DETH.Contract.Approve(&_DETH.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_DETH *DETHTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _DETH.Contract.Approve(&_DETH.TransactOpts, spender, value)
}

// BatchMint is a paid mutator transaction binding the contract method 0x9386e197.
//
// Solidity: function batchMint((address,uint256)[] batcher) returns()
func (_DETH *DETHTransactor) BatchMint(opts *bind.TransactOpts, batcher []IDETHBatchMint) (*types.Transaction, error) {
	return _DETH.contract.Transact(opts, "batchMint", batcher)
}

// BatchMint is a paid mutator transaction binding the contract method 0x9386e197.
//
// Solidity: function batchMint((address,uint256)[] batcher) returns()
func (_DETH *DETHSession) BatchMint(batcher []IDETHBatchMint) (*types.Transaction, error) {
	return _DETH.Contract.BatchMint(&_DETH.TransactOpts, batcher)
}

// BatchMint is a paid mutator transaction binding the contract method 0x9386e197.
//
// Solidity: function batchMint((address,uint256)[] batcher) returns()
func (_DETH *DETHTransactorSession) BatchMint(batcher []IDETHBatchMint) (*types.Transaction, error) {
	return _DETH.Contract.BatchMint(&_DETH.TransactOpts, batcher)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_DETH *DETHTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _DETH.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_DETH *DETHSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _DETH.Contract.Burn(&_DETH.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_DETH *DETHTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _DETH.Contract.Burn(&_DETH.TransactOpts, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DETH *DETHTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DETH.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DETH *DETHSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DETH.Contract.GrantRole(&_DETH.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DETH *DETHTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DETH.Contract.GrantRole(&_DETH.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x0415cc15.
//
// Solidity: function initialize((address,address,address,address) init) returns()
func (_DETH *DETHTransactor) Initialize(opts *bind.TransactOpts, init DETHInit) (*types.Transaction, error) {
	return _DETH.contract.Transact(opts, "initialize", init)
}

// Initialize is a paid mutator transaction binding the contract method 0x0415cc15.
//
// Solidity: function initialize((address,address,address,address) init) returns()
func (_DETH *DETHSession) Initialize(init DETHInit) (*types.Transaction, error) {
	return _DETH.Contract.Initialize(&_DETH.TransactOpts, init)
}

// Initialize is a paid mutator transaction binding the contract method 0x0415cc15.
//
// Solidity: function initialize((address,address,address,address) init) returns()
func (_DETH *DETHTransactorSession) Initialize(init DETHInit) (*types.Transaction, error) {
	return _DETH.Contract.Initialize(&_DETH.TransactOpts, init)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address staker, uint256 amount) returns()
func (_DETH *DETHTransactor) Mint(opts *bind.TransactOpts, staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DETH.contract.Transact(opts, "mint", staker, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address staker, uint256 amount) returns()
func (_DETH *DETHSession) Mint(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DETH.Contract.Mint(&_DETH.TransactOpts, staker, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address staker, uint256 amount) returns()
func (_DETH *DETHTransactorSession) Mint(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DETH.Contract.Mint(&_DETH.TransactOpts, staker, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_DETH *DETHTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DETH.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_DETH *DETHSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DETH.Contract.Permit(&_DETH.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_DETH *DETHTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _DETH.Contract.Permit(&_DETH.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DETH *DETHTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DETH.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DETH *DETHSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DETH.Contract.RenounceRole(&_DETH.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_DETH *DETHTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _DETH.Contract.RenounceRole(&_DETH.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DETH *DETHTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DETH.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DETH *DETHSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DETH.Contract.RevokeRole(&_DETH.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DETH *DETHTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DETH.Contract.RevokeRole(&_DETH.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_DETH *DETHTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DETH.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_DETH *DETHSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DETH.Contract.Transfer(&_DETH.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_DETH *DETHTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DETH.Contract.Transfer(&_DETH.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_DETH *DETHTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DETH.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_DETH *DETHSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DETH.Contract.TransferFrom(&_DETH.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_DETH *DETHTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DETH.Contract.TransferFrom(&_DETH.TransactOpts, from, to, value)
}

// DETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DETH contract.
type DETHApprovalIterator struct {
	Event *DETHApproval // Event containing the contract specifics and raw log

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
func (it *DETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DETHApproval)
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
		it.Event = new(DETHApproval)
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
func (it *DETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DETHApproval represents a Approval event raised by the DETH contract.
type DETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DETH *DETHFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*DETHApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DETH.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &DETHApprovalIterator{contract: _DETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DETH *DETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DETHApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _DETH.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DETHApproval)
				if err := _DETH.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_DETH *DETHFilterer) ParseApproval(log types.Log) (*DETHApproval, error) {
	event := new(DETHApproval)
	if err := _DETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DETHEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the DETH contract.
type DETHEIP712DomainChangedIterator struct {
	Event *DETHEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *DETHEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DETHEIP712DomainChanged)
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
		it.Event = new(DETHEIP712DomainChanged)
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
func (it *DETHEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DETHEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DETHEIP712DomainChanged represents a EIP712DomainChanged event raised by the DETH contract.
type DETHEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_DETH *DETHFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*DETHEIP712DomainChangedIterator, error) {

	logs, sub, err := _DETH.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &DETHEIP712DomainChangedIterator{contract: _DETH.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_DETH *DETHFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *DETHEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _DETH.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DETHEIP712DomainChanged)
				if err := _DETH.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_DETH *DETHFilterer) ParseEIP712DomainChanged(log types.Log) (*DETHEIP712DomainChanged, error) {
	event := new(DETHEIP712DomainChanged)
	if err := _DETH.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DETHInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DETH contract.
type DETHInitializedIterator struct {
	Event *DETHInitialized // Event containing the contract specifics and raw log

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
func (it *DETHInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DETHInitialized)
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
		it.Event = new(DETHInitialized)
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
func (it *DETHInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DETHInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DETHInitialized represents a Initialized event raised by the DETH contract.
type DETHInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DETH *DETHFilterer) FilterInitialized(opts *bind.FilterOpts) (*DETHInitializedIterator, error) {

	logs, sub, err := _DETH.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DETHInitializedIterator{contract: _DETH.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DETH *DETHFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DETHInitialized) (event.Subscription, error) {

	logs, sub, err := _DETH.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DETHInitialized)
				if err := _DETH.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DETH *DETHFilterer) ParseInitialized(log types.Log) (*DETHInitialized, error) {
	event := new(DETHInitialized)
	if err := _DETH.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DETHRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the DETH contract.
type DETHRoleAdminChangedIterator struct {
	Event *DETHRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DETHRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DETHRoleAdminChanged)
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
		it.Event = new(DETHRoleAdminChanged)
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
func (it *DETHRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DETHRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DETHRoleAdminChanged represents a RoleAdminChanged event raised by the DETH contract.
type DETHRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DETH *DETHFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DETHRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _DETH.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DETHRoleAdminChangedIterator{contract: _DETH.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DETH *DETHFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DETHRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _DETH.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DETHRoleAdminChanged)
				if err := _DETH.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DETH *DETHFilterer) ParseRoleAdminChanged(log types.Log) (*DETHRoleAdminChanged, error) {
	event := new(DETHRoleAdminChanged)
	if err := _DETH.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DETHRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the DETH contract.
type DETHRoleGrantedIterator struct {
	Event *DETHRoleGranted // Event containing the contract specifics and raw log

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
func (it *DETHRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DETHRoleGranted)
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
		it.Event = new(DETHRoleGranted)
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
func (it *DETHRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DETHRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DETHRoleGranted represents a RoleGranted event raised by the DETH contract.
type DETHRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DETH *DETHFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DETHRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DETH.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DETHRoleGrantedIterator{contract: _DETH.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DETH *DETHFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DETHRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DETH.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DETHRoleGranted)
				if err := _DETH.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DETH *DETHFilterer) ParseRoleGranted(log types.Log) (*DETHRoleGranted, error) {
	event := new(DETHRoleGranted)
	if err := _DETH.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DETHRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the DETH contract.
type DETHRoleRevokedIterator struct {
	Event *DETHRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DETHRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DETHRoleRevoked)
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
		it.Event = new(DETHRoleRevoked)
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
func (it *DETHRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DETHRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DETHRoleRevoked represents a RoleRevoked event raised by the DETH contract.
type DETHRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DETH *DETHFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DETHRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DETH.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DETHRoleRevokedIterator{contract: _DETH.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DETH *DETHFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DETHRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DETH.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DETHRoleRevoked)
				if err := _DETH.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DETH *DETHFilterer) ParseRoleRevoked(log types.Log) (*DETHRoleRevoked, error) {
	event := new(DETHRoleRevoked)
	if err := _DETH.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DETH contract.
type DETHTransferIterator struct {
	Event *DETHTransfer // Event containing the contract specifics and raw log

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
func (it *DETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DETHTransfer)
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
		it.Event = new(DETHTransfer)
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
func (it *DETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DETHTransfer represents a Transfer event raised by the DETH contract.
type DETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DETH *DETHFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*DETHTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DETH.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DETHTransferIterator{contract: _DETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DETH *DETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DETHTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DETH.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DETHTransfer)
				if err := _DETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_DETH *DETHFilterer) ParseTransfer(log types.Log) (*DETHTransfer, error) {
	event := new(DETHTransfer)
	if err := _DETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
