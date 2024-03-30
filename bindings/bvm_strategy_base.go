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

// StrategyBaseMetaData contains all meta data concerning the StrategyBase contract.
var StrategyBaseMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ETHBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WETHBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"weth\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"newShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_stakingWeth\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_pauser\",\"type\":\"address\",\"internalType\":\"contractIL2Pauser\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauser\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIL2Pauser\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"relayer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToStaking\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToStakingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakingToShares\",\"inputs\":[{\"name\":\"amountStaking\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakingToSharesView\",\"inputs\":[{\"name\":\"amountStaking\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakingWeth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferETHToL2DappLinkBridge\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1StakingManagerAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transferWETHToL2DappLinkBridge\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1StakingManagerAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wethAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"userStaking\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userStakingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weth\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608080604052346100b8577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166100a957506001600160401b036002600160401b031982821601610064575b6040516112a290816100be8239f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1388080610055565b63f92ee8a960e01b8152600490fd5b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806310ed193a1461014d5780631a51cf75146101485780633055a78c1461014357806339b70e381461013e5780633a98ef391461013957806347e7ef241461013457806356a7117b14610125578063821460f5146101025780638406c0791461012f578063893d4d9e1461012a5780639910a665146101255780639d9cc41a146101205780639fd0506d1461011b578063a841ec3914610116578063ab5921e114610111578063ce7c2ac21461010c578063d9caed1214610107578063f5f1b29c14610102578063f8c8765e146100fd5763fc1d98fb146100f857600080fd5b610851565b6106f5565b6103af565b610601565b6105df565b61050e565b6104b4565b61048b565b610407565b610391565b61014d565b6103de565b610275565b610257565b61022e565b610212565b6101bb565b346101a55760203660031901126101a5576003546103e8908181018091116101a057610177610eb2565b9182018092116101a057602091610193610198926004356108ca565b6108dd565b604051908152f35b61087a565b600080fd5b6001600160a01b038116036101a557565b60a03660031901126101a55760206102086044356101d8816101aa565b6064356101e4816101aa565b6101f960018060a01b036002541633146108fd565b608435916024356004356109d9565b6040519015158152f35b346101a55760003660031901126101a557602047604051908152f35b346101a55760003660031901126101a5576000546040516001600160a01b039091168152602090f35b346101a55760003660031901126101a5576020600354604051908152f35b346101a55760403660031901126101a55760048035610293816101aa565b6024359060206102c160018060a01b036102b281600054163314610b1d565b8554166001600160a01b031690565b604051631c187a5960e11b815294859182905afa90811561038c576102f56102fa926103599560009161035d575b50610b80565b610f6d565b6103496103446103326003549361019361031386610890565b61032c83610327610322610eb2565b610890565b610bcc565b926108ca565b809361033f821515610bd9565b61089f565b600355565b6040519081529081906020820190565b0390f35b61037f915060203d602011610385575b610377818361099c565b810190610b68565b386102ef565b503d61036d565b6109cd565b346101a55760203660031901126101a5576020610198600435610c3c565b346101a55760203660031901126101a55760206101986103d96004356103d4816101aa565b610d0a565b610c3c565b346101a55760003660031901126101a5576002546040516001600160a01b039091168152602090f35b346101a55760003660031901126101a5576001546040516370a0823160e01b815230600482015290602090829060249082906001600160a01b03165afa801561038c5760209160009161045e575b50604051908152f35b61047e9150823d8411610484575b610476818361099c565b8101906109be565b38610455565b503d61046c565b346101a55760003660031901126101a5576004546040516001600160a01b039091168152602090f35b60c03660031901126101a55760206102086044356104d1816101aa565b6064356104dd816101aa565b608435906104ea826101aa565b6104ff60018060a01b036002541633146108fd565b60a43592602435600435610c6b565b346101a55760003660031901126101a557604080519061052d8261095f565b604d825260207f4261736520537472617465677920696d706c656d656e746174696f6e20746f2060208401527f696e68657269742066726f6d20666f72206d6f726520636f6d706c657820696d828401526c706c656d656e746174696f6e7360981b606084015281519283916020835281519182602085015260005b8381106105c9575050600083830185015250601f01601f19168101030190f35b81810183015187820187015286945082016105a9565b346101a55760203660031901126101a55760206101986004356103d4816101aa565b346101a55760603660031901126101a5576004803561061f816101aa565b60243561062b816101aa565b60443591602061065960018060a01b0361064a81600054163314610b1d565b8654166001600160a01b031690565b604051632602edb160e01b815295869182905afa93841561038c576106d49461068a916000916106d6575b50610d69565b61069382611018565b6106cf610344600354946106a986821115610db5565b6106c96106b587610890565b610193836106c4610322610eb2565b6108ca565b95610bcc565b6110ce565b005b6106ef915060203d60201161038557610377818361099c565b38610684565b346101a55760803660031901126101a557600435610712816101aa565b6024359061071f826101aa565b60443561072b816101aa565b60643590610738826101aa565b60008051602061124d833981519152549367ffffffffffffffff60ff8660401c1615951680159081610849575b600114908161083f575b159081610836575b506108245760008051602061124d833981519152805467ffffffffffffffff191660011790556107ab93856107fa57610e3d565b6107b157005b60008051602061124d833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b60008051602061124d833981519152805460ff60401b191668010000000000000000179055610e3d565b60405163f92ee8a960e01b8152600490fd5b90501538610777565b303b15915061076f565b869150610765565b346101a55760003660031901126101a5576001546040516001600160a01b039091168152602090f35b634e487b7160e01b600052601160045260246000fd5b906103e882018092116101a057565b919082018092116101a057565b906801bc16d674ec800000918281029281840414901517156101a057565b818102929181159184041417156101a057565b81156108e7570490565b634e487b7160e01b600052601260045260246000fd5b1561090457565b60405162461bcd60e51b815260206004820152601860248201527f5374726174656779426173652e6f6e6c7952656c6179657200000000000000006044820152606490fd5b634e487b7160e01b600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761097b57604052565b610949565b60e0810190811067ffffffffffffffff82111761097b57604052565b90601f8019910116810190811067ffffffffffffffff82111761097b57604052565b908160209103126101a5575190565b6040513d6000823e3d90fd5b929390916109fd6109f160015460018060a01b031690565b6001600160a01b031690565b6040516370a0823160e01b808252306004830152919690916020919082846024818c5afa93841561038c57600094610afe575b506801bc16d674ec80000080941015610a5157505050505050505050600090565b604051908152306004820152978290899060249082905afa90811561038c57610a8e610ad994610acb93610ade9b600091610ae1575b50046108ac565b604051630136884760e71b94810194909452602484019890985260448301969096526001600160a01b039095166064820152949385906084820190565b03601f19810186528561099c565b610f1a565b90565b610af89150863d881161048457610476818361099c565b38610a87565b610b16919450833d851161048457610476818361099c565b9238610a30565b15610b2457565b606460405162461bcd60e51b815260206004820152602060248201527f5374726174656779426173652e6f6e6c7953747261746567794d616e616765726044820152fd5b908160209103126101a5575180151581036101a55790565b15610b8757565b60405162461bcd60e51b815260206004820152601b60248201527f5374726174656779426173653a6465706f7369742070617573656400000000006044820152606490fd5b919082039182116101a057565b15610be057565b60405162461bcd60e51b815260206004820152602e60248201527f5374726174656779426173652e6465706f7369743a206e65775368617265732060448201526d63616e6e6f74206265207a65726f60901b6064820152608490fd5b6003546103e88082018092116101a057610c54610eb2565b9081018091116101a057610ade92610193916108ca565b93909491926801bc16d674ec80000091824711610c8e5750505050505050600090565b8247048381029381850414901517156101a057610ade96604051966303b0230f60e41b60208901526024880152604487015260018060a01b03809216606487015216608485015260a484015260a48352610ce783610980565b3491610f1a565b67ffffffffffffffff811161097b57601f01601f191660200190565b600054604051633d3f06c960e11b81526001600160a01b0392831660048201523060248201529160209183916044918391165afa90811561038c57600091610d50575090565b610ade915060203d60201161048457610476818361099c565b15610d7057565b60405162461bcd60e51b815260206004820152601c60248201527f5374726174656779426173653a776974686472617720706175736564000000006044820152606490fd5b15610dbc57565b60405162461bcd60e51b815260206004820152604d60248201527f5374726174656779426173652e77697468647261773a20616d6f756e7453686160448201527f726573206d757374206265206c657373207468616e206f7220657175616c207460648201526c6f20746f74616c53686172657360981b608482015260a490fd5b9260ff60008051602061124d8339815191525460401c1615610ea05760018060a01b0392838092816bffffffffffffffffffffffff60a01b9716876001541617600155168560045416176004551683600054161760005516906002541617600255565b604051631afcd79f60e31b8152600490fd5b6001546040516370a0823160e01b815230600482015290602090829060249082906001600160a01b03165afa90811561038c57600091610efb575b504781018091116101a05790565b610f14915060203d60201161048457610476818361099c565b38610eed565b929060061b622673c001603f5a0210610f3e576000928392602083519301915af190565b6308c379a06000526020805278185361666543616c6c3a204e6f7420656e6f756768206761736058526064601cfd5b6001546001600160a01b0391821691168114908115610ffa575b5015610f8f57565b60405162461bcd60e51b815260206004820152603a60248201527f5374726174656779426173652e6465706f7369743a2043616e206f6e6c79206460448201527f65706f736974207374616b696e675765746820616e64206574680000000000006064820152608490fd5b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee91501438610f87565b6001546001600160a01b03918216911681149081156110b0575b501561103a57565b60405162461bcd60e51b815260206004820152604260248201527f5374726174656779426173652e77697468647261773a2043616e206f6e6c792060448201527f776974686472617720746865207374726174656779207765746820616e6420656064820152610e8d60f31b608482015260a490fd5b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee91501438611032565b906001600160a01b0390811673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee810361111b5750600080938193829383918315611111575b1690f11561038c57565b6108fc9250611107565b9261118692600092839260405191602083019363a9059cbb60e01b85521660248301526044820152604481526111508161095f565b519082865af13d156111e1573d9061116782610cee565b91611175604051938461099c565b82523d6000602084013e5b836111e9565b80519081151591826111bf575b505061119c5750565b604051635274afe760e01b81526001600160a01b03919091166004820152602490fd5b6111da9250906020806111d6938301019101610b68565b1590565b3880611193565b606090611180565b9061121057508051156111fe57805190602001fd5b604051630a12f52160e11b8152600490fd5b81511580611243575b611221575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561121956fef0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a2646970667358221220f1a28e9f923ab85ed69fc6c79f8f2eecf9938698d8f496623d5801f39f20622264736f6c63430008180033",
}

// StrategyBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyBaseMetaData.ABI instead.
var StrategyBaseABI = StrategyBaseMetaData.ABI

// StrategyBaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StrategyBaseMetaData.Bin instead.
var StrategyBaseBin = StrategyBaseMetaData.Bin

// DeployStrategyBase deploys a new Ethereum contract, binding an instance of StrategyBase to it.
func DeployStrategyBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StrategyBase, error) {
	parsed, err := StrategyBaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StrategyBaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StrategyBase{StrategyBaseCaller: StrategyBaseCaller{contract: contract}, StrategyBaseTransactor: StrategyBaseTransactor{contract: contract}, StrategyBaseFilterer: StrategyBaseFilterer{contract: contract}}, nil
}

// StrategyBase is an auto generated Go binding around an Ethereum contract.
type StrategyBase struct {
	StrategyBaseCaller     // Read-only binding to the contract
	StrategyBaseTransactor // Write-only binding to the contract
	StrategyBaseFilterer   // Log filterer for contract events
}

// StrategyBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategyBaseSession struct {
	Contract     *StrategyBase     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrategyBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyBaseCallerSession struct {
	Contract *StrategyBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StrategyBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyBaseTransactorSession struct {
	Contract     *StrategyBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StrategyBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyBaseRaw struct {
	Contract *StrategyBase // Generic contract binding to access the raw methods on
}

// StrategyBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyBaseCallerRaw struct {
	Contract *StrategyBaseCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyBaseTransactorRaw struct {
	Contract *StrategyBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategyBase creates a new instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBase(address common.Address, backend bind.ContractBackend) (*StrategyBase, error) {
	contract, err := bindStrategyBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrategyBase{StrategyBaseCaller: StrategyBaseCaller{contract: contract}, StrategyBaseTransactor: StrategyBaseTransactor{contract: contract}, StrategyBaseFilterer: StrategyBaseFilterer{contract: contract}}, nil
}

// NewStrategyBaseCaller creates a new read-only instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBaseCaller(address common.Address, caller bind.ContractCaller) (*StrategyBaseCaller, error) {
	contract, err := bindStrategyBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseCaller{contract: contract}, nil
}

// NewStrategyBaseTransactor creates a new write-only instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyBaseTransactor, error) {
	contract, err := bindStrategyBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTransactor{contract: contract}, nil
}

// NewStrategyBaseFilterer creates a new log filterer instance of StrategyBase, bound to a specific deployed contract.
func NewStrategyBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyBaseFilterer, error) {
	contract, err := bindStrategyBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseFilterer{contract: contract}, nil
}

// bindStrategyBase binds a generic wrapper to an already deployed contract.
func bindStrategyBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrategyBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyBase *StrategyBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyBase.Contract.StrategyBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyBase *StrategyBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBase.Contract.StrategyBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyBase *StrategyBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyBase.Contract.StrategyBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyBase *StrategyBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyBase *StrategyBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyBase *StrategyBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyBase.Contract.contract.Transact(opts, method, params...)
}

// ETHBalance is a free data retrieval call binding the contract method 0x3055a78c.
//
// Solidity: function ETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) ETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "ETHBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ETHBalance is a free data retrieval call binding the contract method 0x3055a78c.
//
// Solidity: function ETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) ETHBalance() (*big.Int, error) {
	return _StrategyBase.Contract.ETHBalance(&_StrategyBase.CallOpts)
}

// ETHBalance is a free data retrieval call binding the contract method 0x3055a78c.
//
// Solidity: function ETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) ETHBalance() (*big.Int, error) {
	return _StrategyBase.Contract.ETHBalance(&_StrategyBase.CallOpts)
}

// WETHBalance is a free data retrieval call binding the contract method 0x9d9cc41a.
//
// Solidity: function WETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) WETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "WETHBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WETHBalance is a free data retrieval call binding the contract method 0x9d9cc41a.
//
// Solidity: function WETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) WETHBalance() (*big.Int, error) {
	return _StrategyBase.Contract.WETHBalance(&_StrategyBase.CallOpts)
}

// WETHBalance is a free data retrieval call binding the contract method 0x9d9cc41a.
//
// Solidity: function WETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) WETHBalance() (*big.Int, error) {
	return _StrategyBase.Contract.WETHBalance(&_StrategyBase.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_StrategyBase *StrategyBaseCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_StrategyBase *StrategyBaseSession) Explanation() (string, error) {
	return _StrategyBase.Contract.Explanation(&_StrategyBase.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_StrategyBase *StrategyBaseCallerSession) Explanation() (string, error) {
	return _StrategyBase.Contract.Explanation(&_StrategyBase.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_StrategyBase *StrategyBaseCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_StrategyBase *StrategyBaseSession) Pauser() (common.Address, error) {
	return _StrategyBase.Contract.Pauser(&_StrategyBase.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) Pauser() (common.Address, error) {
	return _StrategyBase.Contract.Pauser(&_StrategyBase.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_StrategyBase *StrategyBaseCaller) Relayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "relayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_StrategyBase *StrategyBaseSession) Relayer() (common.Address, error) {
	return _StrategyBase.Contract.Relayer(&_StrategyBase.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) Relayer() (common.Address, error) {
	return _StrategyBase.Contract.Relayer(&_StrategyBase.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) Shares(user common.Address) (*big.Int, error) {
	return _StrategyBase.Contract.Shares(&_StrategyBase.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _StrategyBase.Contract.Shares(&_StrategyBase.CallOpts, user)
}

// SharesToStaking is a free data retrieval call binding the contract method 0x56a7117b.
//
// Solidity: function sharesToStaking(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) SharesToStaking(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "sharesToStaking", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToStaking is a free data retrieval call binding the contract method 0x56a7117b.
//
// Solidity: function sharesToStaking(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) SharesToStaking(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToStaking(&_StrategyBase.CallOpts, amountShares)
}

// SharesToStaking is a free data retrieval call binding the contract method 0x56a7117b.
//
// Solidity: function sharesToStaking(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) SharesToStaking(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToStaking(&_StrategyBase.CallOpts, amountShares)
}

// SharesToStakingView is a free data retrieval call binding the contract method 0x9910a665.
//
// Solidity: function sharesToStakingView(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) SharesToStakingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "sharesToStakingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToStakingView is a free data retrieval call binding the contract method 0x9910a665.
//
// Solidity: function sharesToStakingView(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) SharesToStakingView(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToStakingView(&_StrategyBase.CallOpts, amountShares)
}

// SharesToStakingView is a free data retrieval call binding the contract method 0x9910a665.
//
// Solidity: function sharesToStakingView(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) SharesToStakingView(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToStakingView(&_StrategyBase.CallOpts, amountShares)
}

// StakingToShares is a free data retrieval call binding the contract method 0x893d4d9e.
//
// Solidity: function stakingToShares(uint256 amountStaking) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) StakingToShares(opts *bind.CallOpts, amountStaking *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "stakingToShares", amountStaking)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingToShares is a free data retrieval call binding the contract method 0x893d4d9e.
//
// Solidity: function stakingToShares(uint256 amountStaking) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) StakingToShares(amountStaking *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.StakingToShares(&_StrategyBase.CallOpts, amountStaking)
}

// StakingToShares is a free data retrieval call binding the contract method 0x893d4d9e.
//
// Solidity: function stakingToShares(uint256 amountStaking) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) StakingToShares(amountStaking *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.StakingToShares(&_StrategyBase.CallOpts, amountStaking)
}

// StakingToSharesView is a free data retrieval call binding the contract method 0x10ed193a.
//
// Solidity: function stakingToSharesView(uint256 amountStaking) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) StakingToSharesView(opts *bind.CallOpts, amountStaking *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "stakingToSharesView", amountStaking)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingToSharesView is a free data retrieval call binding the contract method 0x10ed193a.
//
// Solidity: function stakingToSharesView(uint256 amountStaking) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) StakingToSharesView(amountStaking *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.StakingToSharesView(&_StrategyBase.CallOpts, amountStaking)
}

// StakingToSharesView is a free data retrieval call binding the contract method 0x10ed193a.
//
// Solidity: function stakingToSharesView(uint256 amountStaking) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) StakingToSharesView(amountStaking *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.StakingToSharesView(&_StrategyBase.CallOpts, amountStaking)
}

// StakingWeth is a free data retrieval call binding the contract method 0xfc1d98fb.
//
// Solidity: function stakingWeth() view returns(address)
func (_StrategyBase *StrategyBaseCaller) StakingWeth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "stakingWeth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingWeth is a free data retrieval call binding the contract method 0xfc1d98fb.
//
// Solidity: function stakingWeth() view returns(address)
func (_StrategyBase *StrategyBaseSession) StakingWeth() (common.Address, error) {
	return _StrategyBase.Contract.StakingWeth(&_StrategyBase.CallOpts)
}

// StakingWeth is a free data retrieval call binding the contract method 0xfc1d98fb.
//
// Solidity: function stakingWeth() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) StakingWeth() (common.Address, error) {
	return _StrategyBase.Contract.StakingWeth(&_StrategyBase.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyBase *StrategyBaseCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyBase *StrategyBaseSession) StrategyManager() (common.Address, error) {
	return _StrategyBase.Contract.StrategyManager(&_StrategyBase.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) StrategyManager() (common.Address, error) {
	return _StrategyBase.Contract.StrategyManager(&_StrategyBase.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) TotalShares() (*big.Int, error) {
	return _StrategyBase.Contract.TotalShares(&_StrategyBase.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) TotalShares() (*big.Int, error) {
	return _StrategyBase.Contract.TotalShares(&_StrategyBase.CallOpts)
}

// UserStakingView is a free data retrieval call binding the contract method 0xf5f1b29c.
//
// Solidity: function userStakingView(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) UserStakingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "userStakingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserStakingView is a free data retrieval call binding the contract method 0xf5f1b29c.
//
// Solidity: function userStakingView(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) UserStakingView(user common.Address) (*big.Int, error) {
	return _StrategyBase.Contract.UserStakingView(&_StrategyBase.CallOpts, user)
}

// UserStakingView is a free data retrieval call binding the contract method 0xf5f1b29c.
//
// Solidity: function userStakingView(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) UserStakingView(user common.Address) (*big.Int, error) {
	return _StrategyBase.Contract.UserStakingView(&_StrategyBase.CallOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address weth, uint256 amount) returns(uint256 newShares)
func (_StrategyBase *StrategyBaseTransactor) Deposit(opts *bind.TransactOpts, weth common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "deposit", weth, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address weth, uint256 amount) returns(uint256 newShares)
func (_StrategyBase *StrategyBaseSession) Deposit(weth common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Deposit(&_StrategyBase.TransactOpts, weth, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address weth, uint256 amount) returns(uint256 newShares)
func (_StrategyBase *StrategyBaseTransactorSession) Deposit(weth common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Deposit(&_StrategyBase.TransactOpts, weth, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _stakingWeth, address _relayer, address _strategyManager, address _pauser) returns()
func (_StrategyBase *StrategyBaseTransactor) Initialize(opts *bind.TransactOpts, _stakingWeth common.Address, _relayer common.Address, _strategyManager common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "initialize", _stakingWeth, _relayer, _strategyManager, _pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _stakingWeth, address _relayer, address _strategyManager, address _pauser) returns()
func (_StrategyBase *StrategyBaseSession) Initialize(_stakingWeth common.Address, _relayer common.Address, _strategyManager common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.Initialize(&_StrategyBase.TransactOpts, _stakingWeth, _relayer, _strategyManager, _pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _stakingWeth, address _relayer, address _strategyManager, address _pauser) returns()
func (_StrategyBase *StrategyBaseTransactorSession) Initialize(_stakingWeth common.Address, _relayer common.Address, _strategyManager common.Address, _pauser common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.Initialize(&_StrategyBase.TransactOpts, _stakingWeth, _relayer, _strategyManager, _pauser)
}

// TransferETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0x1a51cf75.
//
// Solidity: function transferETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, address l1StakingManagerAddr, uint256 gasLimit) payable returns(bool)
func (_StrategyBase *StrategyBaseTransactor) TransferETHToL2DappLinkBridge(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, l1StakingManagerAddr common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "transferETHToL2DappLinkBridge", sourceChainId, destChainId, bridge, l1StakingManagerAddr, gasLimit)
}

// TransferETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0x1a51cf75.
//
// Solidity: function transferETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, address l1StakingManagerAddr, uint256 gasLimit) payable returns(bool)
func (_StrategyBase *StrategyBaseSession) TransferETHToL2DappLinkBridge(sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, l1StakingManagerAddr common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.TransferETHToL2DappLinkBridge(&_StrategyBase.TransactOpts, sourceChainId, destChainId, bridge, l1StakingManagerAddr, gasLimit)
}

// TransferETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0x1a51cf75.
//
// Solidity: function transferETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, address l1StakingManagerAddr, uint256 gasLimit) payable returns(bool)
func (_StrategyBase *StrategyBaseTransactorSession) TransferETHToL2DappLinkBridge(sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, l1StakingManagerAddr common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.TransferETHToL2DappLinkBridge(&_StrategyBase.TransactOpts, sourceChainId, destChainId, bridge, l1StakingManagerAddr, gasLimit)
}

// TransferWETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0xa841ec39.
//
// Solidity: function transferWETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, address l1StakingManagerAddr, address wethAddress, uint256 gasLimit) payable returns(bool)
func (_StrategyBase *StrategyBaseTransactor) TransferWETHToL2DappLinkBridge(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, l1StakingManagerAddr common.Address, wethAddress common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "transferWETHToL2DappLinkBridge", sourceChainId, destChainId, bridge, l1StakingManagerAddr, wethAddress, gasLimit)
}

// TransferWETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0xa841ec39.
//
// Solidity: function transferWETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, address l1StakingManagerAddr, address wethAddress, uint256 gasLimit) payable returns(bool)
func (_StrategyBase *StrategyBaseSession) TransferWETHToL2DappLinkBridge(sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, l1StakingManagerAddr common.Address, wethAddress common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.TransferWETHToL2DappLinkBridge(&_StrategyBase.TransactOpts, sourceChainId, destChainId, bridge, l1StakingManagerAddr, wethAddress, gasLimit)
}

// TransferWETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0xa841ec39.
//
// Solidity: function transferWETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, address l1StakingManagerAddr, address wethAddress, uint256 gasLimit) payable returns(bool)
func (_StrategyBase *StrategyBaseTransactorSession) TransferWETHToL2DappLinkBridge(sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, l1StakingManagerAddr common.Address, wethAddress common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.TransferWETHToL2DappLinkBridge(&_StrategyBase.TransactOpts, sourceChainId, destChainId, bridge, l1StakingManagerAddr, wethAddress, gasLimit)
}

// UserStaking is a paid mutator transaction binding the contract method 0x821460f5.
//
// Solidity: function userStaking(address user) returns(uint256)
func (_StrategyBase *StrategyBaseTransactor) UserStaking(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "userStaking", user)
}

// UserStaking is a paid mutator transaction binding the contract method 0x821460f5.
//
// Solidity: function userStaking(address user) returns(uint256)
func (_StrategyBase *StrategyBaseSession) UserStaking(user common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.UserStaking(&_StrategyBase.TransactOpts, user)
}

// UserStaking is a paid mutator transaction binding the contract method 0x821460f5.
//
// Solidity: function userStaking(address user) returns(uint256)
func (_StrategyBase *StrategyBaseTransactorSession) UserStaking(user common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.UserStaking(&_StrategyBase.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address weth, uint256 amountShares) returns()
func (_StrategyBase *StrategyBaseTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, weth common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "withdraw", recipient, weth, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address weth, uint256 amountShares) returns()
func (_StrategyBase *StrategyBaseSession) Withdraw(recipient common.Address, weth common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Withdraw(&_StrategyBase.TransactOpts, recipient, weth, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address weth, uint256 amountShares) returns()
func (_StrategyBase *StrategyBaseTransactorSession) Withdraw(recipient common.Address, weth common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Withdraw(&_StrategyBase.TransactOpts, recipient, weth, amountShares)
}

// StrategyBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StrategyBase contract.
type StrategyBaseInitializedIterator struct {
	Event *StrategyBaseInitialized // Event containing the contract specifics and raw log

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
func (it *StrategyBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseInitialized)
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
		it.Event = new(StrategyBaseInitialized)
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
func (it *StrategyBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseInitialized represents a Initialized event raised by the StrategyBase contract.
type StrategyBaseInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StrategyBase *StrategyBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*StrategyBaseInitializedIterator, error) {

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseInitializedIterator{contract: _StrategyBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StrategyBase *StrategyBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StrategyBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseInitialized)
				if err := _StrategyBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StrategyBase *StrategyBaseFilterer) ParseInitialized(log types.Log) (*StrategyBaseInitialized, error) {
	event := new(StrategyBaseInitialized)
	if err := _StrategyBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
