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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"newShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_relayer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"relayer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenETHBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenWETHBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferETHToL2DappLinkBridge\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transferWETHToL2DappLinkBridge\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100d0565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006e5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cd5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611ccb806100df6000396000f3fe6080604052600436106101815760003560e01c80637a8b2637116100d1578063ac36440a1161008a578063e3dae51c11610064578063e3dae51c1461043f578063f3e738751461045f578063f8c8765e1461047f578063fabc1cbc1461049f57600080fd5b8063ac36440a146103ea578063ce7c2ac2146103ff578063d9caed121461041f57600080fd5b80637a8b2637146103285780638406c07914610348578063886f1195146103685780638c871019146103885780638f6a6240146103a8578063ab5921e1146103c857600080fd5b806347e7ef241161013e578063553ca5f811610118578063553ca5f8146102b2578063595c6a67146102d25780635ac86ab7146102e75780635c975abb1461031357600080fd5b806347e7ef241461026c57806349fcf4a71461028c5780634fcffb8a1461029f57600080fd5b806310d67a2f14610186578063136439dd146101a85780632495a599146101c857806339b70e38146102055780633a98ef39146102255780634102e0c214610249575b600080fd5b34801561019257600080fd5b506101a66101a1366004611915565b6104bf565b005b3480156101b457600080fd5b506101a66101c3366004611932565b610579565b3480156101d457600080fd5b506033546101e8906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561021157600080fd5b506032546101e8906001600160a01b031681565b34801561023157600080fd5b5061023b60355481565b6040519081526020016101fc565b61025c61025736600461194b565b6106b8565b60405190151581526020016101fc565b34801561027857600080fd5b5061023b61028736600461198a565b6108a0565b34801561029857600080fd5b504761023b565b61025c6102ad36600461194b565b610a28565b3480156102be57600080fd5b5061023b6102cd366004611915565b610b22565b3480156102de57600080fd5b506101a6610b36565b3480156102f357600080fd5b5061025c6103023660046119b6565b6001805460ff9092161b9081161490565b34801561031f57600080fd5b5060015461023b565b34801561033457600080fd5b5061023b610343366004611932565b610bfd565b34801561035457600080fd5b506034546101e8906001600160a01b031681565b34801561037457600080fd5b506000546101e8906001600160a01b031681565b34801561039457600080fd5b5061023b6103a3366004611932565b610c40565b3480156103b457600080fd5b5061023b6103c3366004611915565b610c4b565b3480156103d457600080fd5b506103dd610c59565b6040516101fc91906119fd565b3480156103f657600080fd5b5061023b610c79565b34801561040b57600080fd5b5061023b61041a366004611915565b610ceb565b34801561042b57600080fd5b506101a661043a366004611a30565b610d60565b34801561044b57600080fd5b5061023b61045a366004611932565b610f0a565b34801561046b57600080fd5b5061023b61047a366004611932565b610f43565b34801561048b57600080fd5b506101a661049a366004611a71565b610f4e565b3480156104ab57600080fd5b506101a66104ba366004611932565b61108f565b60008054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610510573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105349190611acd565b6001600160a01b0316336001600160a01b03161461056d5760405162461bcd60e51b815260040161056490611aea565b60405180910390fd5b610576816111e9565b50565b60005460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156105c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e59190611b34565b6106015760405162461bcd60e51b815260040161056490611b56565b6001548181161461067a5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610564565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b6034546000906001600160a01b031633146107105760405162461bcd60e51b815260206004820152601860248201527729ba3930ba32b3bca130b9b29737b7363ca932b630bcb2b960411b6044820152606401610564565b6033546040516370a0823160e01b81523060048201526801bc16d674ec800000916001600160a01b0316906370a0823190602401602060405180830381865afa158015610761573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107859190611b9e565b10610894576033546040516370a0823160e01b81523060048201526000916801bc16d674ec800000916001600160a01b03909116906370a0823190602401602060405180830381865afa1580156107e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108049190611b9e565b61080e9190611bcd565b610821906801bc16d674ec800000611bef565b60405160248101889052604481018790526001600160a01b03861660648201526084810182905290915060009061088a9086908690349060a40160408051601f198184030181529190526020810180516001600160e01b0316637e3c2a0b60e01b1790526112e0565b9250610898915050565b5060005b949350505050565b6001805460009182918116036108f45760405162461bcd60e51b815260206004820152601960248201527814185d5cd8589b194e881a5b99195e081a5cc81c185d5cd959603a1b6044820152606401610564565b6032546001600160a01b0316331461094e5760405162461bcd60e51b815260206004820181905260248201527f5374726174656779426173652e6f6e6c7953747261746567794d616e616765726044820152606401610564565b610958848461133e565b60355460006109696103e883611c06565b905060006103e86109786113ec565b6109829190611c06565b905060006109908783611c19565b90508061099d8489611bef565b6109a79190611bcd565b955085600003610a105760405162461bcd60e51b815260206004820152602e60248201527f5374726174656779426173652e6465706f7369743a206e65775368617265732060448201526d63616e6e6f74206265207a65726f60901b6064820152608401610564565b610a1a8685611c06565b603555505050505092915050565b6034546000906001600160a01b03163314610a805760405162461bcd60e51b815260206004820152601860248201527729ba3930ba32b3bca130b9b29737b7363ca932b630bcb2b960411b6044820152606401610564565b6801bc16d674ec800000471115610894576000610aa66801bc16d674ec80000047611bcd565b610ab9906801bc16d674ec800000611bef565b60405160248101889052604481018790526001600160a01b03861660648201526084810182905290915060009061088a9086908690349060a40160408051601f198184030181529190526020810180516001600160e01b03166371bbb61760e01b1790526112e0565b6000610b3061034383610ceb565b92915050565b60005460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610b7e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ba29190611b34565b610bbe5760405162461bcd60e51b815260040161056490611b56565b600019600181905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6000806103e8603554610c109190611c06565b905060006103e8610c1f6113ec565b610c299190611c06565b905081610c368583611bef565b6108989190611bcd565b6000610b3082610f0a565b6000610b3061047a83610ceb565b60606040518060800160405280604d8152602001611c49604d9139905090565b6033546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a0823190602401602060405180830381865afa158015610cc2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce69190611b9e565b905090565b603254604051633d3f06c960e11b81526001600160a01b0383811660048301523060248301526000921690637a7e0d9290604401602060405180830381865afa158015610d3c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b309190611b9e565b60018054600290811603610db25760405162461bcd60e51b815260206004820152601960248201527814185d5cd8589b194e881a5b99195e081a5cc81c185d5cd959603a1b6044820152606401610564565b6032546001600160a01b03163314610e0c5760405162461bcd60e51b815260206004820181905260248201527f5374726174656779426173652e6f6e6c7953747261746567794d616e616765726044820152606401610564565b610e17848484611467565b60355480831115610ea65760405162461bcd60e51b815260206004820152604d60248201527f5374726174656779426173652e77697468647261773a20616d6f756e7453686160448201527f726573206d757374206265206c657373207468616e206f7220657175616c207460648201526c6f20746f74616c53686172657360981b608482015260a401610564565b6000610eb46103e883611c06565b905060006103e8610ec36113ec565b610ecd9190611c06565b9050600082610edc8784611bef565b610ee69190611bcd565b9050610ef28685611c19565b603555610f00888883611522565b5050505050505050565b6000806103e8603554610f1d9190611c06565b905060006103e8610f2c6113ec565b610f369190611c06565b905080610c368386611bef565b6000610b3082610bfd565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff16600081158015610f945750825b905060008267ffffffffffffffff166001148015610fb15750303b155b905081158015610fbf575080155b15610fdd5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561100757845460ff60401b1916600160401b1785555b6110118988611597565b603280546001600160a01b038089166001600160a01b0319928316179092556034805491821691909216179055831561108457845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b60008054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111049190611acd565b6001600160a01b0316336001600160a01b0316146111345760405162461bcd60e51b815260040161056490611aea565b6001541981196001541916146111b25760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610564565b600181905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016106ad565b6001600160a01b0381166112775760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610564565b600054604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1600080546001600160a01b0319166001600160a01b0392909216919091179055565b60008060006112f08660006115c5565b905080611326576308c379a06000526020805278185361666543616c6c3a204e6f7420656e6f756768206761736058526064601cfd5b600080855160208701888b5af1979650505050505050565b6033546001600160a01b038381169116148061137657506001600160a01b03821673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee145b6113e85760405162461bcd60e51b815260206004820152603e60248201527f5374726174656779426173652e6465706f7369743a2043616e206f6e6c79206460448201527f65706f73697420756e6465726c79696e67546f6b656e20616e642065746800006064820152608401610564565b5050565b6033546040516370a0823160e01b815230600482015260009147916001600160a01b03909116906370a0823190602401602060405180830381865afa158015611439573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061145d9190611b9e565b610ce69190611c06565b6033546001600160a01b038381169116148061149f57506001600160a01b03821673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee145b61151d5760405162461bcd60e51b815260206004820152604360248201527f5374726174656779426173652e77697468647261773a2043616e206f6e6c792060448201527f77697468647261772074686520737472617465677920746f6b656e20616e64206064820152620cae8d60eb1b608482015260a401610564565b505050565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeed196001600160a01b03831601611583576040516001600160a01b0384169082156108fc029083906000818181858888f1935050505015801561157d573d6000803e3d6000fd5b50505050565b61151d6001600160a01b03831684836115e3565b61159f611635565b603380546001600160a01b0319166001600160a01b0384161790556113e8816000611680565b600080603f83619c4001026040850201603f5a021015949350505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b17905261151d908490611766565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661167e57604051631afcd79f60e31b815260040160405180910390fd5b565b6000546001600160a01b03161580156116a157506001600160a01b03821615155b6117235760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610564565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a26113e8826111e9565b600061177b6001600160a01b038416836117c9565b905080516000141580156117a057508080602001905181019061179e9190611b34565b155b1561151d57604051635274afe760e01b81526001600160a01b0384166004820152602401610564565b60606117d7838360006117de565b9392505050565b6060814710156118035760405163cd78605960e01b8152306004820152602401610564565b600080856001600160a01b0316848660405161181f9190611c2c565b60006040518083038185875af1925050503d806000811461185c576040519150601f19603f3d011682016040523d82523d6000602084013e611861565b606091505b509150915061187186838361187b565b9695505050505050565b6060826118905761188b826118d7565b6117d7565b81511580156118a757506001600160a01b0384163b155b156118d057604051639996b31560e01b81526001600160a01b0385166004820152602401610564565b50806117d7565b8051156118e75780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b6001600160a01b038116811461057657600080fd5b60006020828403121561192757600080fd5b81356117d781611900565b60006020828403121561194457600080fd5b5035919050565b6000806000806080858703121561196157600080fd5b8435935060208501359250604085013561197a81611900565b9396929550929360600135925050565b6000806040838503121561199d57600080fd5b82356119a881611900565b946020939093013593505050565b6000602082840312156119c857600080fd5b813560ff811681146117d757600080fd5b60005b838110156119f45781810151838201526020016119dc565b50506000910152565b6020815260008251806020840152611a1c8160408501602087016119d9565b601f01601f19169190910160400192915050565b600080600060608486031215611a4557600080fd5b8335611a5081611900565b92506020840135611a6081611900565b929592945050506040919091013590565b60008060008060808587031215611a8757600080fd5b8435611a9281611900565b93506020850135611aa281611900565b92506040850135611ab281611900565b91506060850135611ac281611900565b939692955090935050565b600060208284031215611adf57600080fd5b81516117d781611900565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215611b4657600080fd5b815180151581146117d757600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b600060208284031215611bb057600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600082611bea57634e487b7160e01b600052601260045260246000fd5b500490565b8082028115828204841417610b3057610b30611bb7565b80820180821115610b3057610b30611bb7565b81810381811115610b3057610b30611bb7565b60008251611c3e8184602087016119d9565b919091019291505056fe4261736520537472617465677920696d706c656d656e746174696f6e20746f20696e68657269742066726f6d20666f72206d6f726520636f6d706c657820696d706c656d656e746174696f6e73a264697066735822122021ea4a923fbd43e3bf1660602e02840e40fd484bdd0f2d55ff60f5a834fa853264736f6c63430008180033",
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

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyBase *StrategyBaseCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyBase *StrategyBaseSession) Paused(index uint8) (bool, error) {
	return _StrategyBase.Contract.Paused(&_StrategyBase.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyBase *StrategyBaseCallerSession) Paused(index uint8) (bool, error) {
	return _StrategyBase.Contract.Paused(&_StrategyBase.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) Paused0() (*big.Int, error) {
	return _StrategyBase.Contract.Paused0(&_StrategyBase.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) Paused0() (*big.Int, error) {
	return _StrategyBase.Contract.Paused0(&_StrategyBase.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyBase *StrategyBaseCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyBase *StrategyBaseSession) PauserRegistry() (common.Address, error) {
	return _StrategyBase.Contract.PauserRegistry(&_StrategyBase.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) PauserRegistry() (common.Address, error) {
	return _StrategyBase.Contract.PauserRegistry(&_StrategyBase.CallOpts)
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

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) SharesToUnderlying(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "sharesToUnderlying", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToUnderlying(&_StrategyBase.CallOpts, amountShares)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToUnderlying(&_StrategyBase.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToUnderlyingView(&_StrategyBase.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.SharesToUnderlyingView(&_StrategyBase.CallOpts, amountShares)
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

// TokenETHBalance is a free data retrieval call binding the contract method 0x49fcf4a7.
//
// Solidity: function tokenETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) TokenETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "tokenETHBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenETHBalance is a free data retrieval call binding the contract method 0x49fcf4a7.
//
// Solidity: function tokenETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) TokenETHBalance() (*big.Int, error) {
	return _StrategyBase.Contract.TokenETHBalance(&_StrategyBase.CallOpts)
}

// TokenETHBalance is a free data retrieval call binding the contract method 0x49fcf4a7.
//
// Solidity: function tokenETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) TokenETHBalance() (*big.Int, error) {
	return _StrategyBase.Contract.TokenETHBalance(&_StrategyBase.CallOpts)
}

// TokenWETHBalance is a free data retrieval call binding the contract method 0xac36440a.
//
// Solidity: function tokenWETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) TokenWETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "tokenWETHBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenWETHBalance is a free data retrieval call binding the contract method 0xac36440a.
//
// Solidity: function tokenWETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseSession) TokenWETHBalance() (*big.Int, error) {
	return _StrategyBase.Contract.TokenWETHBalance(&_StrategyBase.CallOpts)
}

// TokenWETHBalance is a free data retrieval call binding the contract method 0xac36440a.
//
// Solidity: function tokenWETHBalance() view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) TokenWETHBalance() (*big.Int, error) {
	return _StrategyBase.Contract.TokenWETHBalance(&_StrategyBase.CallOpts)
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

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) UnderlyingToShares(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "underlyingToShares", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.UnderlyingToShares(&_StrategyBase.CallOpts, amountUnderlying)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.UnderlyingToShares(&_StrategyBase.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.UnderlyingToSharesView(&_StrategyBase.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBase.Contract.UnderlyingToSharesView(&_StrategyBase.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_StrategyBase *StrategyBaseCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_StrategyBase *StrategyBaseSession) UnderlyingToken() (common.Address, error) {
	return _StrategyBase.Contract.UnderlyingToken(&_StrategyBase.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_StrategyBase *StrategyBaseCallerSession) UnderlyingToken() (common.Address, error) {
	return _StrategyBase.Contract.UnderlyingToken(&_StrategyBase.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBase.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _StrategyBase.Contract.UserUnderlyingView(&_StrategyBase.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_StrategyBase *StrategyBaseCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _StrategyBase.Contract.UserUnderlyingView(&_StrategyBase.CallOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_StrategyBase *StrategyBaseTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_StrategyBase *StrategyBaseSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Deposit(&_StrategyBase.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_StrategyBase *StrategyBaseTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Deposit(&_StrategyBase.TransactOpts, token, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _underlyingToken, address _relayer, address _pauserRegistry, address _strategyManager) returns()
func (_StrategyBase *StrategyBaseTransactor) Initialize(opts *bind.TransactOpts, _underlyingToken common.Address, _relayer common.Address, _pauserRegistry common.Address, _strategyManager common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "initialize", _underlyingToken, _relayer, _pauserRegistry, _strategyManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _underlyingToken, address _relayer, address _pauserRegistry, address _strategyManager) returns()
func (_StrategyBase *StrategyBaseSession) Initialize(_underlyingToken common.Address, _relayer common.Address, _pauserRegistry common.Address, _strategyManager common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.Initialize(&_StrategyBase.TransactOpts, _underlyingToken, _relayer, _pauserRegistry, _strategyManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _underlyingToken, address _relayer, address _pauserRegistry, address _strategyManager) returns()
func (_StrategyBase *StrategyBaseTransactorSession) Initialize(_underlyingToken common.Address, _relayer common.Address, _pauserRegistry common.Address, _strategyManager common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.Initialize(&_StrategyBase.TransactOpts, _underlyingToken, _relayer, _pauserRegistry, _strategyManager)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyBase *StrategyBaseTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyBase *StrategyBaseSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Pause(&_StrategyBase.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyBase *StrategyBaseTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Pause(&_StrategyBase.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyBase *StrategyBaseTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyBase *StrategyBaseSession) PauseAll() (*types.Transaction, error) {
	return _StrategyBase.Contract.PauseAll(&_StrategyBase.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyBase *StrategyBaseTransactorSession) PauseAll() (*types.Transaction, error) {
	return _StrategyBase.Contract.PauseAll(&_StrategyBase.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_StrategyBase *StrategyBaseTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_StrategyBase *StrategyBaseSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetPauserRegistry(&_StrategyBase.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_StrategyBase *StrategyBaseTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.SetPauserRegistry(&_StrategyBase.TransactOpts, newPauserRegistry)
}

// TransferETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0x4fcffb8a.
//
// Solidity: function transferETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, uint256 gasLimit) payable returns(bool)
func (_StrategyBase *StrategyBaseTransactor) TransferETHToL2DappLinkBridge(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "transferETHToL2DappLinkBridge", sourceChainId, destChainId, bridge, gasLimit)
}

// TransferETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0x4fcffb8a.
//
// Solidity: function transferETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, uint256 gasLimit) payable returns(bool)
func (_StrategyBase *StrategyBaseSession) TransferETHToL2DappLinkBridge(sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.TransferETHToL2DappLinkBridge(&_StrategyBase.TransactOpts, sourceChainId, destChainId, bridge, gasLimit)
}

// TransferETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0x4fcffb8a.
//
// Solidity: function transferETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, uint256 gasLimit) payable returns(bool)
func (_StrategyBase *StrategyBaseTransactorSession) TransferETHToL2DappLinkBridge(sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.TransferETHToL2DappLinkBridge(&_StrategyBase.TransactOpts, sourceChainId, destChainId, bridge, gasLimit)
}

// TransferWETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0x4102e0c2.
//
// Solidity: function transferWETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, uint256 gasLimit) payable returns(bool)
func (_StrategyBase *StrategyBaseTransactor) TransferWETHToL2DappLinkBridge(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "transferWETHToL2DappLinkBridge", sourceChainId, destChainId, bridge, gasLimit)
}

// TransferWETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0x4102e0c2.
//
// Solidity: function transferWETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, uint256 gasLimit) payable returns(bool)
func (_StrategyBase *StrategyBaseSession) TransferWETHToL2DappLinkBridge(sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.TransferWETHToL2DappLinkBridge(&_StrategyBase.TransactOpts, sourceChainId, destChainId, bridge, gasLimit)
}

// TransferWETHToL2DappLinkBridge is a paid mutator transaction binding the contract method 0x4102e0c2.
//
// Solidity: function transferWETHToL2DappLinkBridge(uint256 sourceChainId, uint256 destChainId, address bridge, uint256 gasLimit) payable returns(bool)
func (_StrategyBase *StrategyBaseTransactorSession) TransferWETHToL2DappLinkBridge(sourceChainId *big.Int, destChainId *big.Int, bridge common.Address, gasLimit *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.TransferWETHToL2DappLinkBridge(&_StrategyBase.TransactOpts, sourceChainId, destChainId, bridge, gasLimit)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyBase *StrategyBaseTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyBase *StrategyBaseSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Unpause(&_StrategyBase.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyBase *StrategyBaseTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Unpause(&_StrategyBase.TransactOpts, newPausedStatus)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_StrategyBase *StrategyBaseTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_StrategyBase *StrategyBaseSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.UserUnderlying(&_StrategyBase.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_StrategyBase *StrategyBaseTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _StrategyBase.Contract.UserUnderlying(&_StrategyBase.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_StrategyBase *StrategyBaseTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBase.contract.Transact(opts, "withdraw", recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_StrategyBase *StrategyBaseSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Withdraw(&_StrategyBase.TransactOpts, recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_StrategyBase *StrategyBaseTransactorSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBase.Contract.Withdraw(&_StrategyBase.TransactOpts, recipient, token, amountShares)
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

// StrategyBasePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StrategyBase contract.
type StrategyBasePausedIterator struct {
	Event *StrategyBasePaused // Event containing the contract specifics and raw log

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
func (it *StrategyBasePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBasePaused)
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
		it.Event = new(StrategyBasePaused)
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
func (it *StrategyBasePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBasePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBasePaused represents a Paused event raised by the StrategyBase contract.
type StrategyBasePaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyBase *StrategyBaseFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*StrategyBasePausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyBasePausedIterator{contract: _StrategyBase.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyBase *StrategyBaseFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StrategyBasePaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBasePaused)
				if err := _StrategyBase.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyBase *StrategyBaseFilterer) ParsePaused(log types.Log) (*StrategyBasePaused, error) {
	event := new(StrategyBasePaused)
	if err := _StrategyBase.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBasePauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the StrategyBase contract.
type StrategyBasePauserRegistrySetIterator struct {
	Event *StrategyBasePauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *StrategyBasePauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBasePauserRegistrySet)
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
		it.Event = new(StrategyBasePauserRegistrySet)
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
func (it *StrategyBasePauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBasePauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBasePauserRegistrySet represents a PauserRegistrySet event raised by the StrategyBase contract.
type StrategyBasePauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_StrategyBase *StrategyBaseFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*StrategyBasePauserRegistrySetIterator, error) {

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &StrategyBasePauserRegistrySetIterator{contract: _StrategyBase.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_StrategyBase *StrategyBaseFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *StrategyBasePauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBasePauserRegistrySet)
				if err := _StrategyBase.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_StrategyBase *StrategyBaseFilterer) ParsePauserRegistrySet(log types.Log) (*StrategyBasePauserRegistrySet, error) {
	event := new(StrategyBasePauserRegistrySet)
	if err := _StrategyBase.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StrategyBase contract.
type StrategyBaseUnpausedIterator struct {
	Event *StrategyBaseUnpaused // Event containing the contract specifics and raw log

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
func (it *StrategyBaseUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseUnpaused)
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
		it.Event = new(StrategyBaseUnpaused)
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
func (it *StrategyBaseUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseUnpaused represents a Unpaused event raised by the StrategyBase contract.
type StrategyBaseUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyBase *StrategyBaseFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*StrategyBaseUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBase.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseUnpausedIterator{contract: _StrategyBase.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyBase *StrategyBaseFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StrategyBaseUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBase.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseUnpaused)
				if err := _StrategyBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyBase *StrategyBaseFilterer) ParseUnpaused(log types.Log) (*StrategyBaseUnpaused, error) {
	event := new(StrategyBaseUnpaused)
	if err := _StrategyBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
