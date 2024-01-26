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

// L2PoolMetaData contains all meta data concerning the L2Pool contract.
var L2PoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ChainIdIsNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"ChainIdNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorBlockChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"MINDepositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LessThanMINDepositAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughETH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"NotEnoughToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"StableCoinNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"DepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"DepositETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"DepositWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20toL1Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawETHtoL1Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawWETHtoL1Success\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"IsSupportChainId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"IsSupportStableCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_GAS_Limit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ReLayer\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SendAssertToUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"TransferETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20ToL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawETHtoL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawWETHToL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MultisigWallet\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_MINDepositAmount\",\"type\":\"uint256\"}],\"name\":\"setMINDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_PerFee\",\"type\":\"uint256\"}],\"name\":\"setPerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setSupportStableCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setValidChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100cf565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006d5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cc5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b61240280620000dd5f395ff3fe6080604052600436106101c5575f3560e01c80638456cb59116100f2578063c4d66de811610092578063dd0c346011610062578063dd0c3460146104ea578063ed06da87146104ff578063f2f678ed14610536578063fa86184814610555575f80fd5b8063c4d66de81461047a578063d28597f414610499578063d547741f146104ac578063d73f14e5146104cb575f80fd5b8063a525ae2b116100cd578063a525ae2b146103df578063a835227714610410578063b1887e9b1461042f578063b33280dd1461045b575f80fd5b80638456cb591461039957806391d14854146103ad578063a217fddf146103cc575f80fd5b80633f4ba83a1161016857806372e6950a1161013857806372e6950a1461032657806372fe6a7e1461034557806373c23be914610373578063816ba7e714610386575f80fd5b80633f4ba83a146102bb5780634d043743146102cf5780635c975abb146102e4578063650c227614610307575f80fd5b80632f2ff15d116101a35780632f2ff15d1461023d57806332b806761461025e57806336568abe1461028957806339891dd7146102a8575f80fd5b806301ffc9a7146101c95780630227047a146101fd578063248a9ca314610210575b5f80fd5b3480156101d4575f80fd5b506101e86101e3366004612057565b610569565b60405190151581526020015b60405180910390f35b6101e861020b366004612099565b61059f565b34801561021b575f80fd5b5061022f61022a3660046120d2565b6108da565b6040519081526020016101f4565b348015610248575f80fd5b5061025c6102573660046120e9565b6108fa565b005b348015610269575f80fd5b5061022f610278366004612113565b60036020525f908152604090205481565b348015610294575f80fd5b5061025c6102a33660046120e9565b61091c565b6101e86102b636600461212c565b610954565b3480156102c6575f80fd5b5061025c610b85565b3480156102da575f80fd5b5061022f60045481565b3480156102ef575f80fd5b505f805160206123ad8339815191525460ff166101e8565b348015610312575f80fd5b5061025c6103213660046120d2565b610b9a565b348015610331575f80fd5b506101e8610340366004612154565b610baa565b348015610350575f80fd5b506101e861035f3660046120d2565b5f9081526020819052604090205460ff1690565b6101e8610381366004612099565b610e9a565b6101e86103943660046120e9565b61105b565b3480156103a4575f80fd5b5061025c6111e2565b3480156103b8575f80fd5b506101e86103c73660046120e9565b6111f4565b3480156103d7575f80fd5b5061022f5f81565b3480156103ea575f80fd5b506006546103fb9063ffffffff1681565b60405163ffffffff90911681526020016101f4565b34801561041b575f80fd5b506101e861042a366004612176565b61122a565b34801561043a575f80fd5b506104436114a0565b6040516001600160a01b0390911681526020016101f4565b348015610466575f80fd5b5061025c6104753660046121c4565b6114f5565b348015610485575f80fd5b5061025c610494366004612113565b61152a565b6101e86104a736600461212c565b61166d565b3480156104b7575f80fd5b5061025c6104c63660046120e9565b611a92565b3480156104d6575f80fd5b5061025c6104e53660046121f9565b611aae565b3480156104f5575f80fd5b5061022f60055481565b34801561050a575f80fd5b506101e8610519366004612113565b6001600160a01b03165f9081526001602052604090205460ff1690565b348015610541575f80fd5b5061025c6105503660046120d2565b611ad8565b348015610560575f80fd5b5061022f611b2b565b5f6001600160e01b03198216637965db0b60e01b148061059957506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f60ff196105bb60015f8051602061236d833981519152612230565b6040516020016105cd91815260200190565b60405160208183030381529060405280519060200120166105ed81611b72565b4661060f866001600160a01b03165f9081526001602052604090205460ff1690565b61063c57604051631e859bc960e01b81526001600160a01b03871660048201526024015b60405180910390fd5b8062082750036107595760405163095ea7b360e01b815273d8a791fe2be73eb6e6cf1eb0cb3f36adc9b3f8f96004820152602481018590526001600160a01b0387169063095ea7b3906044016020604051808303815f875af11580156106a4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106c89190612243565b5060065460405163a93a4af960e01b81526001600160a01b038089166004830152871660248201526044810186905263ffffffff9091166064820181905273d8a791fe2be73eb6e6cf1eb0cb3f36adc9b3f8f99163a93a4af991906084015f604051808303815f88803b15801561073d575f80fd5b5087f115801561074f573d5f803e3d5ffd5b5050505050610878565b8061044d036108565760405163095ea7b360e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede6004820152602481018590526001600160a01b0387169063095ea7b3906044016020604051808303815f875af11580156107c0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107e49190612243565b5060405163cd58657960e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede9063cd58657990610824905f90899089908c90849060040161225e565b5f604051808303815f87803b15801561083b575f80fd5b505af115801561084d573d5f803e3d5ffd5b50505050610878565b600a8114610878575b604051639474dee160e01b815260040160405180910390fd5b604080514681524260208201526001600160a01b0388811682840152871660608201526080810186905290517f1bc1b1207e0c319b7baad0317c09f30575218af6f5a59f1f0f15de8d217672cc9181900360a00190a150600195945050505050565b5f9081525f8051602061238d833981519152602052604090206001015490565b610903826108da565b61090c81611b72565b6109168383611b7c565b50505050565b6001600160a01b03811633146109455760405163334bd91960e11b815260040160405180910390fd5b61094f8282611c1d565b505050565b5f60ff1961097060015f8051602061236d833981519152612230565b60405160200161098291815260200190565b60405160208183030381529060405280519060200120166109a281611b72565b46478411156109c457604051632c1d501360e11b815260040160405180910390fd5b806208275003610a5957600654604051636705b1e760e11b81526001600160a01b03871660048201526024810186905263ffffffff90911660448201819052737003e7b7186f0e6601203b99f7b8decbfa391cf99163ce0b63ce919087906064015f604051808303818589803b158015610a3c575f80fd5b5088f1158015610a4e573d5f803e3d5ffd5b505050505050610b2c565b8061044d03610ad65760405163cd58657960e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede9063cd586579908690610aa3905f908a9084908390819060040161225e565b5f604051808303818588803b158015610aba575f80fd5b505af1158015610acc573d5f803e3d5ffd5b5050505050610b2c565b80600a0361085f57600654604051631474f2a960e31b81526010602160991b019163a3a79548918791610aa39173deaddeaddeaddeaddeaddeaddeaddeaddead0000918b91859163ffffffff16906004016122a4565b604080514681524260208201526001600160a01b038716818301526060810186905290517ff1fa6dbbb3e74f1bef249f22545b238726c1cc79c5571a1416575674c7e63f399181900360800190a1506001949350505050565b5f610b8f81611b72565b610b97611c96565b50565b5f610ba481611b72565b50600555565b5f8381526020819052604081205460ff16610bdb576040516363b5c9db60e01b815260048101859052602401610633565b5f610be46114a0565b6040516370a0823160e01b81523060048201529091505f906001600160a01b038316906370a0823190602401602060405180830381865afa158015610c2b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c4f91906122e0565b6040516323b872dd60e01b8152336004820152306024820152604481018690529091506001600160a01b038316906323b872dd906064016020604051808303815f875af1158015610ca2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cc69190612243565b506040516370a0823160e01b81523060048201525f906001600160a01b038416906370a0823190602401602060405180830381865afa158015610d0b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d2f91906122e0565b90505f610d3c8383612230565b9050600454811015610d6c5760048054604051631e50dd1360e31b81529182015260248101829052604401610633565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25f90815260026020527fde032e96a833a2191f55c5766e34554d9b07734477ea9cb69c40158c1078e7938054839290610dbb9084906122f7565b90915550506005545f90620f424090610dd4908461230a565b610dde9190612321565b9050610dea8183612230565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25f90815260036020527f67aa9b7d2b6d14f3837d07b1073399a41e4104b1d98f169f02cc04f44f14f4b080549294508392909190610e3e9084906122f7565b9091555050604080518a8152602081018490526001600160a01b038a169133917f148ba123946a56f22f1acafa21440f6834fa36d19a4ad5d752f3647d45488541910160405180910390a36001955050505050505b9392505050565b5f60ff19610eb660015f8051602061236d833981519152612230565b604051602001610ec891815260200190565b6040516020818303038152906040528051906020012016610ee881611b72565b6001600160a01b0385165f9081526001602052604090205460ff16610f2b57604051631e859bc960e01b81526001600160a01b0386166004820152602401610633565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeed196001600160a01b03861601610faa5782471015610f7157604051632c1d501360e11b815260040160405180910390fd5b6040516001600160a01b0385169084156108fc029085905f818181858888f19350505050158015610fa4573d5f803e3d5ffd5b50611050565b6040516370a0823160e01b815230600482015283906001600160a01b038716906370a0823190602401602060405180830381865afa158015610fee573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061101291906122e0565b101561103c576040516311745b6160e21b81526001600160a01b0386166004820152602401610633565b6110506001600160a01b0386168585611cf5565b506001949350505050565b5f8281526020819052604081205460ff1661108c57604051632ef6974160e11b815260048101849052602401610633565b6004543410156110b95760048054604051631e50dd1360e31b815291820152346024820152604401610633565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f90815260026020527f6a26712a1b2f732f4c1fd85f9d6ed8235573aaa2a79aa2bab72c2423a73a9faf80543492906111089084906122f7565b90915550506005545f90620f424090611121903461230a565b61112b9190612321565b90505f6111388234612230565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f90815260036020527f40c35f83c179e47de306c9fe55fdc60064f11dd52adb51ab61b5643ee626f98d8054929350849290919061118c9084906122f7565b909155505060408051868152602081018390526001600160a01b0386169133917f25dbc63f879c7913a1e148949bf04761d9ff12ca9519a1c0ac9d854e8a90a1e9910160405180910390a3506001949350505050565b5f6111ec81611b72565b610b97611d54565b5f9182525f8051602061238d833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b5f8481526020819052604081205460ff1661125b576040516363b5c9db60e01b815260048101869052602401610633565b6001600160a01b0383165f9081526001602052604090205460ff1661129e57604051631e859bc960e01b81526001600160a01b0384166004820152602401610633565b6040516370a0823160e01b81523060048201525f906001600160a01b038516906370a0823190602401602060405180830381865afa1580156112e2573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061130691906122e0565b905061131d6001600160a01b038516333086611d9c565b6040516370a0823160e01b81523060048201525f906001600160a01b038616906370a0823190602401602060405180830381865afa158015611361573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061138591906122e0565b90505f6113928383612230565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee5f90815260026020527f6a26712a1b2f732f4c1fd85f9d6ed8235573aaa2a79aa2bab72c2423a73a9faf805492935087929091906113e69084906122f7565b90915550506005545f90620f4240906113ff908461230a565b6114099190612321565b90506114158183612230565b6001600160a01b0388165f908152600360205260408120805492945083929091906114419084906122f7565b9091555050604080518a8152602081018490526001600160a01b03808b16923392918b16917fabc6fa8a366ad5068e5fdafc0582c211ead1967abbc53a92bd3ec8770120af37910160405180910390a450600198975050505050505050565b5f46620827508190036114bb576004605360981b0191505090565b8061044d036114df57734f9a0e7fd2bf6067db6994cf12e4495df938e6e991505090565b80600a0361085f576006602160991b0191505090565b5f6114ff81611b72565b506001600160a01b03919091165f908152600160205260409020805460ff1916911515919091179055565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f8115801561156f5750825b90505f8267ffffffffffffffff16600114801561158b5750303b155b905081158015611599575080155b156115b75760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156115e157845460ff60401b1916600160401b1785555b6115e9611dd5565b6115f1611ddf565b6115fb5f876108fa565b67016345785d8a00006004556127106005556006805463ffffffff1916620493e0179055831561166557845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b5f60ff1961168960015f8051602061236d833981519152612230565b60405160200161169b91815260200190565b60405160208183030381529060405280519060200120166116bb81611b72565b465f6116c56114a0565b6040516370a0823160e01b81523060048201529091506001600160a01b038216906370a0823190602401602060405180830381865afa15801561170a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061172e91906122e0565b851115611759576040516311745b6160e21b81526001600160a01b0382166004820152602401610633565b8162082750036118765760405163095ea7b360e01b8152737003e7b7186f0e6601203b99f7b8decbfa391cf96004820152602481018690526001600160a01b0382169063095ea7b3906044016020604051808303815f875af11580156117c1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117e59190612243565b5060065460405163a93a4af960e01b81526001600160a01b038084166004830152881660248201526044810187905263ffffffff90911660648201819052737003e7b7186f0e6601203b99f7b8decbfa391cf99163a93a4af991906084015f604051808303815f88803b15801561185a575f80fd5b5087f115801561186c573d5f803e3d5ffd5b5050505050611a38565b8161044d0361196c5760405163095ea7b360e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede6004820152602481018690526001600160a01b0382169063095ea7b3906044016020604051808303815f875af11580156118dd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119019190612243565b5060405163cd58657960e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede9063cd586579908790611943905f908b9084908390819060040161225e565b5f604051808303818588803b15801561195a575f80fd5b505af115801561186c573d5f803e3d5ffd5b81600a0361085f5760405163095ea7b360e01b81526010602160991b016004820152602481018690526001600160a01b0382169063095ea7b3906044016020604051808303815f875af11580156119c5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119e99190612243565b50600654604051631474f2a960e31b81526010602160991b019163a3a795489188916119439173deaddeaddeaddeaddeaddeaddeaddeaddead0000918c91859163ffffffff16906004016122a4565b604080514681524260208201526001600160a01b038816818301526060810187905290517fcb1a101498aadad4159168d719f955e559e768912e6f5f5c57b094eb047d000f9181900360800190a150600195945050505050565b611a9b826108da565b611aa481611b72565b6109168383611c1d565b5f611ab881611b72565b505f91825260208290526040909120805460ff1916911515919091179055565b60ff19611af360015f8051602061236d833981519152612230565b604051602001611b0591815260200190565b6040516020818303038152906040528051906020012016611b2581611b72565b50600455565b60ff19611b4660015f8051602061236d833981519152612230565b604051602001611b5891815260200190565b604051602081830303815290604052805190602001201681565b610b978133611def565b5f5f8051602061238d833981519152611b9584846111f4565b611c14575f848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055611bca3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610599565b5f915050610599565b5f5f8051602061238d833981519152611c3684846111f4565b15611c14575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610599565b611c9e611e2c565b5f805160206123ad833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b6040516001600160a01b0383811660248301526044820183905261094f91859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050611e5b565b611d5c611ebc565b5f805160206123ad833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833611cd7565b6040516001600160a01b0384811660248301528381166044830152606482018390526109169186918216906323b872dd90608401611d22565b611ddd611eec565b565b611de7611eec565b611ddd611f35565b611df982826111f4565b611e285760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610633565b5050565b5f805160206123ad8339815191525460ff16611ddd57604051638dfc202b60e01b815260040160405180910390fd5b5f611e6f6001600160a01b03841683611f55565b905080515f14158015611e93575080806020019051810190611e919190612243565b155b1561094f57604051635274afe760e01b81526001600160a01b0384166004820152602401610633565b5f805160206123ad8339815191525460ff1615611ddd5760405163d93c066560e01b815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16611ddd57604051631afcd79f60e31b815260040160405180910390fd5b611f3d611eec565b5f805160206123ad833981519152805460ff19169055565b6060610e9383835f845f80856001600160a01b03168486604051611f799190612340565b5f6040518083038185875af1925050503d805f8114611fb3576040519150601f19603f3d011682016040523d82523d5f602084013e611fb8565b606091505b5091509150611fc8868383611fd2565b9695505050505050565b606082611fe757611fe28261202e565b610e93565b8151158015611ffe57506001600160a01b0384163b155b1561202757604051639996b31560e01b81526001600160a01b0385166004820152602401610633565b5080610e93565b80511561203e5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f60208284031215612067575f80fd5b81356001600160e01b031981168114610e93575f80fd5b80356001600160a01b0381168114612094575f80fd5b919050565b5f805f606084860312156120ab575f80fd5b6120b48461207e565b92506120c26020850161207e565b9150604084013590509250925092565b5f602082840312156120e2575f80fd5b5035919050565b5f80604083850312156120fa575f80fd5b8235915061210a6020840161207e565b90509250929050565b5f60208284031215612123575f80fd5b610e938261207e565b5f806040838503121561213d575f80fd5b6121468361207e565b946020939093013593505050565b5f805f60608486031215612166575f80fd5b833592506120c26020850161207e565b5f805f8060808587031215612189575f80fd5b843593506121996020860161207e565b92506121a76040860161207e565b9396929550929360600135925050565b8015158114610b97575f80fd5b5f80604083850312156121d5575f80fd5b6121de8361207e565b915060208301356121ee816121b7565b809150509250929050565b5f806040838503121561220a575f80fd5b8235915060208301356121ee816121b7565b634e487b7160e01b5f52601160045260245ffd5b818103818111156105995761059961221c565b5f60208284031215612253575f80fd5b8151610e93816121b7565b63ffffffff9590951685526001600160a01b039384166020860152604085019290925290911660608301521515608082015260c060a082018190525f9082015260e00190565b6001600160a01b039485168152929093166020830152604082015263ffffffff909116606082015260a0608082018190525f9082015260c00190565b5f602082840312156122f0575f80fd5b5051919050565b808201808211156105995761059961221c565b80820281158282048414176105995761059961221c565b5f8261233b57634e487b7160e01b5f52601260045260245ffd5b500490565b5f82515f5b8181101561235f5760208186018101518583015201612345565b505f92019182525091905056fe33fe247d78fee2e7fd135c405eda4bd2911c0a73c0a81b36c3bcc967dd06e5ae02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300a2646970667358221220c2dfd6da0e7480554e139aa8e881cfc691005e7d387c1cb6b7bb46dda23b08a164736f6c63430008140033",
}

// L2PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use L2PoolMetaData.ABI instead.
var L2PoolABI = L2PoolMetaData.ABI

// L2PoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2PoolMetaData.Bin instead.
var L2PoolBin = L2PoolMetaData.Bin

// DeployL2Pool deploys a new Ethereum contract, binding an instance of L2Pool to it.
func DeployL2Pool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2Pool, error) {
	parsed, err := L2PoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2PoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2Pool{L2PoolCaller: L2PoolCaller{contract: contract}, L2PoolTransactor: L2PoolTransactor{contract: contract}, L2PoolFilterer: L2PoolFilterer{contract: contract}}, nil
}

// L2Pool is an auto generated Go binding around an Ethereum contract.
type L2Pool struct {
	L2PoolCaller     // Read-only binding to the contract
	L2PoolTransactor // Write-only binding to the contract
	L2PoolFilterer   // Log filterer for contract events
}

// L2PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2PoolSession struct {
	Contract     *L2Pool           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2PoolCallerSession struct {
	Contract *L2PoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// L2PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2PoolTransactorSession struct {
	Contract     *L2PoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2PoolRaw struct {
	Contract *L2Pool // Generic contract binding to access the raw methods on
}

// L2PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2PoolCallerRaw struct {
	Contract *L2PoolCaller // Generic read-only contract binding to access the raw methods on
}

// L2PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2PoolTransactorRaw struct {
	Contract *L2PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2Pool creates a new instance of L2Pool, bound to a specific deployed contract.
func NewL2Pool(address common.Address, backend bind.ContractBackend) (*L2Pool, error) {
	contract, err := bindL2Pool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2Pool{L2PoolCaller: L2PoolCaller{contract: contract}, L2PoolTransactor: L2PoolTransactor{contract: contract}, L2PoolFilterer: L2PoolFilterer{contract: contract}}, nil
}

// NewL2PoolCaller creates a new read-only instance of L2Pool, bound to a specific deployed contract.
func NewL2PoolCaller(address common.Address, caller bind.ContractCaller) (*L2PoolCaller, error) {
	contract, err := bindL2Pool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2PoolCaller{contract: contract}, nil
}

// NewL2PoolTransactor creates a new write-only instance of L2Pool, bound to a specific deployed contract.
func NewL2PoolTransactor(address common.Address, transactor bind.ContractTransactor) (*L2PoolTransactor, error) {
	contract, err := bindL2Pool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2PoolTransactor{contract: contract}, nil
}

// NewL2PoolFilterer creates a new log filterer instance of L2Pool, bound to a specific deployed contract.
func NewL2PoolFilterer(address common.Address, filterer bind.ContractFilterer) (*L2PoolFilterer, error) {
	contract, err := bindL2Pool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2PoolFilterer{contract: contract}, nil
}

// bindL2Pool binds a generic wrapper to an already deployed contract.
func bindL2Pool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2Pool *L2PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2Pool.Contract.L2PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2Pool *L2PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2Pool.Contract.L2PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2Pool *L2PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2Pool.Contract.L2PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2Pool *L2PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2Pool *L2PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2Pool *L2PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2Pool.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L2Pool *L2PoolCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2Pool.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L2Pool *L2PoolSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _L2Pool.Contract.DEFAULTADMINROLE(&_L2Pool.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L2Pool *L2PoolCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _L2Pool.Contract.DEFAULTADMINROLE(&_L2Pool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0x32b80676.
//
// Solidity: function Fee(address ) view returns(uint256)
func (_L2Pool *L2PoolCaller) Fee(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2Pool.contract.Call(opts, &out, "Fee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0x32b80676.
//
// Solidity: function Fee(address ) view returns(uint256)
func (_L2Pool *L2PoolSession) Fee(arg0 common.Address) (*big.Int, error) {
	return _L2Pool.Contract.Fee(&_L2Pool.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0x32b80676.
//
// Solidity: function Fee(address ) view returns(uint256)
func (_L2Pool *L2PoolCallerSession) Fee(arg0 common.Address) (*big.Int, error) {
	return _L2Pool.Contract.Fee(&_L2Pool.CallOpts, arg0)
}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_L2Pool *L2PoolCaller) IsSupportChainId(opts *bind.CallOpts, chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _L2Pool.contract.Call(opts, &out, "IsSupportChainId", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_L2Pool *L2PoolSession) IsSupportChainId(chainId *big.Int) (bool, error) {
	return _L2Pool.Contract.IsSupportChainId(&_L2Pool.CallOpts, chainId)
}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_L2Pool *L2PoolCallerSession) IsSupportChainId(chainId *big.Int) (bool, error) {
	return _L2Pool.Contract.IsSupportChainId(&_L2Pool.CallOpts, chainId)
}

// IsSupportStableCoin is a free data retrieval call binding the contract method 0xed06da87.
//
// Solidity: function IsSupportStableCoin(address ERC20Address) view returns(bool)
func (_L2Pool *L2PoolCaller) IsSupportStableCoin(opts *bind.CallOpts, ERC20Address common.Address) (bool, error) {
	var out []interface{}
	err := _L2Pool.contract.Call(opts, &out, "IsSupportStableCoin", ERC20Address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportStableCoin is a free data retrieval call binding the contract method 0xed06da87.
//
// Solidity: function IsSupportStableCoin(address ERC20Address) view returns(bool)
func (_L2Pool *L2PoolSession) IsSupportStableCoin(ERC20Address common.Address) (bool, error) {
	return _L2Pool.Contract.IsSupportStableCoin(&_L2Pool.CallOpts, ERC20Address)
}

// IsSupportStableCoin is a free data retrieval call binding the contract method 0xed06da87.
//
// Solidity: function IsSupportStableCoin(address ERC20Address) view returns(bool)
func (_L2Pool *L2PoolCallerSession) IsSupportStableCoin(ERC20Address common.Address) (bool, error) {
	return _L2Pool.Contract.IsSupportStableCoin(&_L2Pool.CallOpts, ERC20Address)
}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_L2Pool *L2PoolCaller) L2WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Pool.contract.Call(opts, &out, "L2WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_L2Pool *L2PoolSession) L2WETH() (common.Address, error) {
	return _L2Pool.Contract.L2WETH(&_L2Pool.CallOpts)
}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_L2Pool *L2PoolCallerSession) L2WETH() (common.Address, error) {
	return _L2Pool.Contract.L2WETH(&_L2Pool.CallOpts)
}

// MAXGASLimit is a free data retrieval call binding the contract method 0xa525ae2b.
//
// Solidity: function MAX_GAS_Limit() view returns(uint32)
func (_L2Pool *L2PoolCaller) MAXGASLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _L2Pool.contract.Call(opts, &out, "MAX_GAS_Limit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXGASLimit is a free data retrieval call binding the contract method 0xa525ae2b.
//
// Solidity: function MAX_GAS_Limit() view returns(uint32)
func (_L2Pool *L2PoolSession) MAXGASLimit() (uint32, error) {
	return _L2Pool.Contract.MAXGASLimit(&_L2Pool.CallOpts)
}

// MAXGASLimit is a free data retrieval call binding the contract method 0xa525ae2b.
//
// Solidity: function MAX_GAS_Limit() view returns(uint32)
func (_L2Pool *L2PoolCallerSession) MAXGASLimit() (uint32, error) {
	return _L2Pool.Contract.MAXGASLimit(&_L2Pool.CallOpts)
}

// MINDepositAmount is a free data retrieval call binding the contract method 0x4d043743.
//
// Solidity: function MINDepositAmount() view returns(uint256)
func (_L2Pool *L2PoolCaller) MINDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Pool.contract.Call(opts, &out, "MINDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINDepositAmount is a free data retrieval call binding the contract method 0x4d043743.
//
// Solidity: function MINDepositAmount() view returns(uint256)
func (_L2Pool *L2PoolSession) MINDepositAmount() (*big.Int, error) {
	return _L2Pool.Contract.MINDepositAmount(&_L2Pool.CallOpts)
}

// MINDepositAmount is a free data retrieval call binding the contract method 0x4d043743.
//
// Solidity: function MINDepositAmount() view returns(uint256)
func (_L2Pool *L2PoolCallerSession) MINDepositAmount() (*big.Int, error) {
	return _L2Pool.Contract.MINDepositAmount(&_L2Pool.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L2Pool *L2PoolCaller) PerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Pool.contract.Call(opts, &out, "PerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L2Pool *L2PoolSession) PerFee() (*big.Int, error) {
	return _L2Pool.Contract.PerFee(&_L2Pool.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L2Pool *L2PoolCallerSession) PerFee() (*big.Int, error) {
	return _L2Pool.Contract.PerFee(&_L2Pool.CallOpts)
}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_L2Pool *L2PoolCaller) ReLayer(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2Pool.contract.Call(opts, &out, "ReLayer")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_L2Pool *L2PoolSession) ReLayer() ([32]byte, error) {
	return _L2Pool.Contract.ReLayer(&_L2Pool.CallOpts)
}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_L2Pool *L2PoolCallerSession) ReLayer() ([32]byte, error) {
	return _L2Pool.Contract.ReLayer(&_L2Pool.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L2Pool *L2PoolCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _L2Pool.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L2Pool *L2PoolSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _L2Pool.Contract.GetRoleAdmin(&_L2Pool.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L2Pool *L2PoolCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _L2Pool.Contract.GetRoleAdmin(&_L2Pool.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L2Pool *L2PoolCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _L2Pool.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L2Pool *L2PoolSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _L2Pool.Contract.HasRole(&_L2Pool.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L2Pool *L2PoolCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _L2Pool.Contract.HasRole(&_L2Pool.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2Pool *L2PoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L2Pool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2Pool *L2PoolSession) Paused() (bool, error) {
	return _L2Pool.Contract.Paused(&_L2Pool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2Pool *L2PoolCallerSession) Paused() (bool, error) {
	return _L2Pool.Contract.Paused(&_L2Pool.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2Pool *L2PoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L2Pool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2Pool *L2PoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2Pool.Contract.SupportsInterface(&_L2Pool.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2Pool *L2PoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2Pool.Contract.SupportsInterface(&_L2Pool.CallOpts, interfaceId)
}

// SendAssertToUser is a paid mutator transaction binding the contract method 0x73c23be9.
//
// Solidity: function SendAssertToUser(address _token, address to, uint256 _amount) payable returns(bool)
func (_L2Pool *L2PoolTransactor) SendAssertToUser(opts *bind.TransactOpts, _token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "SendAssertToUser", _token, to, _amount)
}

// SendAssertToUser is a paid mutator transaction binding the contract method 0x73c23be9.
//
// Solidity: function SendAssertToUser(address _token, address to, uint256 _amount) payable returns(bool)
func (_L2Pool *L2PoolSession) SendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.SendAssertToUser(&_L2Pool.TransactOpts, _token, to, _amount)
}

// SendAssertToUser is a paid mutator transaction binding the contract method 0x73c23be9.
//
// Solidity: function SendAssertToUser(address _token, address to, uint256 _amount) payable returns(bool)
func (_L2Pool *L2PoolTransactorSession) SendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.SendAssertToUser(&_L2Pool.TransactOpts, _token, to, _amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0xa8352277.
//
// Solidity: function TransferERC20(uint256 chainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_L2Pool *L2PoolTransactor) TransferERC20(opts *bind.TransactOpts, chainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "TransferERC20", chainId, to, ERC20Address, value)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0xa8352277.
//
// Solidity: function TransferERC20(uint256 chainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_L2Pool *L2PoolSession) TransferERC20(chainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.TransferERC20(&_L2Pool.TransactOpts, chainId, to, ERC20Address, value)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0xa8352277.
//
// Solidity: function TransferERC20(uint256 chainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_L2Pool *L2PoolTransactorSession) TransferERC20(chainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.TransferERC20(&_L2Pool.TransactOpts, chainId, to, ERC20Address, value)
}

// TransferETH is a paid mutator transaction binding the contract method 0x816ba7e7.
//
// Solidity: function TransferETH(uint256 chainId, address to) payable returns(bool)
func (_L2Pool *L2PoolTransactor) TransferETH(opts *bind.TransactOpts, chainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "TransferETH", chainId, to)
}

// TransferETH is a paid mutator transaction binding the contract method 0x816ba7e7.
//
// Solidity: function TransferETH(uint256 chainId, address to) payable returns(bool)
func (_L2Pool *L2PoolSession) TransferETH(chainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _L2Pool.Contract.TransferETH(&_L2Pool.TransactOpts, chainId, to)
}

// TransferETH is a paid mutator transaction binding the contract method 0x816ba7e7.
//
// Solidity: function TransferETH(uint256 chainId, address to) payable returns(bool)
func (_L2Pool *L2PoolTransactorSession) TransferETH(chainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _L2Pool.Contract.TransferETH(&_L2Pool.TransactOpts, chainId, to)
}

// TransferWETH is a paid mutator transaction binding the contract method 0x72e6950a.
//
// Solidity: function TransferWETH(uint256 chainId, address to, uint256 value) returns(bool)
func (_L2Pool *L2PoolTransactor) TransferWETH(opts *bind.TransactOpts, chainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "TransferWETH", chainId, to, value)
}

// TransferWETH is a paid mutator transaction binding the contract method 0x72e6950a.
//
// Solidity: function TransferWETH(uint256 chainId, address to, uint256 value) returns(bool)
func (_L2Pool *L2PoolSession) TransferWETH(chainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.TransferWETH(&_L2Pool.TransactOpts, chainId, to, value)
}

// TransferWETH is a paid mutator transaction binding the contract method 0x72e6950a.
//
// Solidity: function TransferWETH(uint256 chainId, address to, uint256 value) returns(bool)
func (_L2Pool *L2PoolTransactorSession) TransferWETH(chainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.TransferWETH(&_L2Pool.TransactOpts, chainId, to, value)
}

// WithdrawERC20ToL1 is a paid mutator transaction binding the contract method 0x0227047a.
//
// Solidity: function WithdrawERC20ToL1(address _token, address _to, uint256 _amount) payable returns(bool)
func (_L2Pool *L2PoolTransactor) WithdrawERC20ToL1(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "WithdrawERC20ToL1", _token, _to, _amount)
}

// WithdrawERC20ToL1 is a paid mutator transaction binding the contract method 0x0227047a.
//
// Solidity: function WithdrawERC20ToL1(address _token, address _to, uint256 _amount) payable returns(bool)
func (_L2Pool *L2PoolSession) WithdrawERC20ToL1(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.WithdrawERC20ToL1(&_L2Pool.TransactOpts, _token, _to, _amount)
}

// WithdrawERC20ToL1 is a paid mutator transaction binding the contract method 0x0227047a.
//
// Solidity: function WithdrawERC20ToL1(address _token, address _to, uint256 _amount) payable returns(bool)
func (_L2Pool *L2PoolTransactorSession) WithdrawERC20ToL1(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.WithdrawERC20ToL1(&_L2Pool.TransactOpts, _token, _to, _amount)
}

// WithdrawETHtoL1 is a paid mutator transaction binding the contract method 0x39891dd7.
//
// Solidity: function WithdrawETHtoL1(address _to, uint256 _amount) payable returns(bool)
func (_L2Pool *L2PoolTransactor) WithdrawETHtoL1(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "WithdrawETHtoL1", _to, _amount)
}

// WithdrawETHtoL1 is a paid mutator transaction binding the contract method 0x39891dd7.
//
// Solidity: function WithdrawETHtoL1(address _to, uint256 _amount) payable returns(bool)
func (_L2Pool *L2PoolSession) WithdrawETHtoL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.WithdrawETHtoL1(&_L2Pool.TransactOpts, _to, _amount)
}

// WithdrawETHtoL1 is a paid mutator transaction binding the contract method 0x39891dd7.
//
// Solidity: function WithdrawETHtoL1(address _to, uint256 _amount) payable returns(bool)
func (_L2Pool *L2PoolTransactorSession) WithdrawETHtoL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.WithdrawETHtoL1(&_L2Pool.TransactOpts, _to, _amount)
}

// WithdrawWETHToL1 is a paid mutator transaction binding the contract method 0xd28597f4.
//
// Solidity: function WithdrawWETHToL1(address _to, uint256 _amount) payable returns(bool)
func (_L2Pool *L2PoolTransactor) WithdrawWETHToL1(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "WithdrawWETHToL1", _to, _amount)
}

// WithdrawWETHToL1 is a paid mutator transaction binding the contract method 0xd28597f4.
//
// Solidity: function WithdrawWETHToL1(address _to, uint256 _amount) payable returns(bool)
func (_L2Pool *L2PoolSession) WithdrawWETHToL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.WithdrawWETHToL1(&_L2Pool.TransactOpts, _to, _amount)
}

// WithdrawWETHToL1 is a paid mutator transaction binding the contract method 0xd28597f4.
//
// Solidity: function WithdrawWETHToL1(address _to, uint256 _amount) payable returns(bool)
func (_L2Pool *L2PoolTransactorSession) WithdrawWETHToL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.WithdrawWETHToL1(&_L2Pool.TransactOpts, _to, _amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L2Pool *L2PoolTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L2Pool *L2PoolSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2Pool.Contract.GrantRole(&_L2Pool.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L2Pool *L2PoolTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2Pool.Contract.GrantRole(&_L2Pool.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _MultisigWallet) returns()
func (_L2Pool *L2PoolTransactor) Initialize(opts *bind.TransactOpts, _MultisigWallet common.Address) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "initialize", _MultisigWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _MultisigWallet) returns()
func (_L2Pool *L2PoolSession) Initialize(_MultisigWallet common.Address) (*types.Transaction, error) {
	return _L2Pool.Contract.Initialize(&_L2Pool.TransactOpts, _MultisigWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _MultisigWallet) returns()
func (_L2Pool *L2PoolTransactorSession) Initialize(_MultisigWallet common.Address) (*types.Transaction, error) {
	return _L2Pool.Contract.Initialize(&_L2Pool.TransactOpts, _MultisigWallet)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L2Pool *L2PoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L2Pool *L2PoolSession) Pause() (*types.Transaction, error) {
	return _L2Pool.Contract.Pause(&_L2Pool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L2Pool *L2PoolTransactorSession) Pause() (*types.Transaction, error) {
	return _L2Pool.Contract.Pause(&_L2Pool.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L2Pool *L2PoolTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L2Pool *L2PoolSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L2Pool.Contract.RenounceRole(&_L2Pool.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L2Pool *L2PoolTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L2Pool.Contract.RenounceRole(&_L2Pool.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L2Pool *L2PoolTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L2Pool *L2PoolSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2Pool.Contract.RevokeRole(&_L2Pool.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L2Pool *L2PoolTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2Pool.Contract.RevokeRole(&_L2Pool.TransactOpts, role, account)
}

// SetMINDepositAmount is a paid mutator transaction binding the contract method 0xf2f678ed.
//
// Solidity: function setMINDepositAmount(uint256 _MINDepositAmount) returns()
func (_L2Pool *L2PoolTransactor) SetMINDepositAmount(opts *bind.TransactOpts, _MINDepositAmount *big.Int) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "setMINDepositAmount", _MINDepositAmount)
}

// SetMINDepositAmount is a paid mutator transaction binding the contract method 0xf2f678ed.
//
// Solidity: function setMINDepositAmount(uint256 _MINDepositAmount) returns()
func (_L2Pool *L2PoolSession) SetMINDepositAmount(_MINDepositAmount *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.SetMINDepositAmount(&_L2Pool.TransactOpts, _MINDepositAmount)
}

// SetMINDepositAmount is a paid mutator transaction binding the contract method 0xf2f678ed.
//
// Solidity: function setMINDepositAmount(uint256 _MINDepositAmount) returns()
func (_L2Pool *L2PoolTransactorSession) SetMINDepositAmount(_MINDepositAmount *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.SetMINDepositAmount(&_L2Pool.TransactOpts, _MINDepositAmount)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_L2Pool *L2PoolTransactor) SetPerFee(opts *bind.TransactOpts, _PerFee *big.Int) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "setPerFee", _PerFee)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_L2Pool *L2PoolSession) SetPerFee(_PerFee *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.SetPerFee(&_L2Pool.TransactOpts, _PerFee)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_L2Pool *L2PoolTransactorSession) SetPerFee(_PerFee *big.Int) (*types.Transaction, error) {
	return _L2Pool.Contract.SetPerFee(&_L2Pool.TransactOpts, _PerFee)
}

// SetSupportStableCoin is a paid mutator transaction binding the contract method 0xb33280dd.
//
// Solidity: function setSupportStableCoin(address ERC20Address, bool isValid) returns()
func (_L2Pool *L2PoolTransactor) SetSupportStableCoin(opts *bind.TransactOpts, ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "setSupportStableCoin", ERC20Address, isValid)
}

// SetSupportStableCoin is a paid mutator transaction binding the contract method 0xb33280dd.
//
// Solidity: function setSupportStableCoin(address ERC20Address, bool isValid) returns()
func (_L2Pool *L2PoolSession) SetSupportStableCoin(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L2Pool.Contract.SetSupportStableCoin(&_L2Pool.TransactOpts, ERC20Address, isValid)
}

// SetSupportStableCoin is a paid mutator transaction binding the contract method 0xb33280dd.
//
// Solidity: function setSupportStableCoin(address ERC20Address, bool isValid) returns()
func (_L2Pool *L2PoolTransactorSession) SetSupportStableCoin(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L2Pool.Contract.SetSupportStableCoin(&_L2Pool.TransactOpts, ERC20Address, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_L2Pool *L2PoolTransactor) SetValidChainId(opts *bind.TransactOpts, chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "setValidChainId", chainId, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_L2Pool *L2PoolSession) SetValidChainId(chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _L2Pool.Contract.SetValidChainId(&_L2Pool.TransactOpts, chainId, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_L2Pool *L2PoolTransactorSession) SetValidChainId(chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _L2Pool.Contract.SetValidChainId(&_L2Pool.TransactOpts, chainId, isValid)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L2Pool *L2PoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2Pool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L2Pool *L2PoolSession) Unpause() (*types.Transaction, error) {
	return _L2Pool.Contract.Unpause(&_L2Pool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L2Pool *L2PoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _L2Pool.Contract.Unpause(&_L2Pool.TransactOpts)
}

// L2PoolDepositERC20Iterator is returned from FilterDepositERC20 and is used to iterate over the raw logs and unpacked data for DepositERC20 events raised by the L2Pool contract.
type L2PoolDepositERC20Iterator struct {
	Event *L2PoolDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L2PoolDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolDepositERC20)
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
		it.Event = new(L2PoolDepositERC20)
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
func (it *L2PoolDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolDepositERC20 represents a DepositERC20 event raised by the L2Pool contract.
type L2PoolDepositERC20 struct {
	ChainId      *big.Int
	ERC20Address common.Address
	From         common.Address
	To           common.Address
	Value        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositERC20 is a free log retrieval operation binding the contract event 0xabc6fa8a366ad5068e5fdafc0582c211ead1967abbc53a92bd3ec8770120af37.
//
// Solidity: event DepositERC20(uint256 chainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L2Pool *L2PoolFilterer) FilterDepositERC20(opts *bind.FilterOpts, ERC20Address []common.Address, from []common.Address, to []common.Address) (*L2PoolDepositERC20Iterator, error) {

	var ERC20AddressRule []interface{}
	for _, ERC20AddressItem := range ERC20Address {
		ERC20AddressRule = append(ERC20AddressRule, ERC20AddressItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2Pool.contract.FilterLogs(opts, "DepositERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolDepositERC20Iterator{contract: _L2Pool.contract, event: "DepositERC20", logs: logs, sub: sub}, nil
}

// WatchDepositERC20 is a free log subscription operation binding the contract event 0xabc6fa8a366ad5068e5fdafc0582c211ead1967abbc53a92bd3ec8770120af37.
//
// Solidity: event DepositERC20(uint256 chainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L2Pool *L2PoolFilterer) WatchDepositERC20(opts *bind.WatchOpts, sink chan<- *L2PoolDepositERC20, ERC20Address []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var ERC20AddressRule []interface{}
	for _, ERC20AddressItem := range ERC20Address {
		ERC20AddressRule = append(ERC20AddressRule, ERC20AddressItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2Pool.contract.WatchLogs(opts, "DepositERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolDepositERC20)
				if err := _L2Pool.contract.UnpackLog(event, "DepositERC20", log); err != nil {
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

// ParseDepositERC20 is a log parse operation binding the contract event 0xabc6fa8a366ad5068e5fdafc0582c211ead1967abbc53a92bd3ec8770120af37.
//
// Solidity: event DepositERC20(uint256 chainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L2Pool *L2PoolFilterer) ParseDepositERC20(log types.Log) (*L2PoolDepositERC20, error) {
	event := new(L2PoolDepositERC20)
	if err := _L2Pool.contract.UnpackLog(event, "DepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolDepositETHIterator is returned from FilterDepositETH and is used to iterate over the raw logs and unpacked data for DepositETH events raised by the L2Pool contract.
type L2PoolDepositETHIterator struct {
	Event *L2PoolDepositETH // Event containing the contract specifics and raw log

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
func (it *L2PoolDepositETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolDepositETH)
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
		it.Event = new(L2PoolDepositETH)
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
func (it *L2PoolDepositETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolDepositETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolDepositETH represents a DepositETH event raised by the L2Pool contract.
type L2PoolDepositETH struct {
	ChainId *big.Int
	From    common.Address
	To      common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositETH is a free log retrieval operation binding the contract event 0x25dbc63f879c7913a1e148949bf04761d9ff12ca9519a1c0ac9d854e8a90a1e9.
//
// Solidity: event DepositETH(uint256 chainId, address indexed from, address indexed to, uint256 value)
func (_L2Pool *L2PoolFilterer) FilterDepositETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2PoolDepositETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2Pool.contract.FilterLogs(opts, "DepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolDepositETHIterator{contract: _L2Pool.contract, event: "DepositETH", logs: logs, sub: sub}, nil
}

// WatchDepositETH is a free log subscription operation binding the contract event 0x25dbc63f879c7913a1e148949bf04761d9ff12ca9519a1c0ac9d854e8a90a1e9.
//
// Solidity: event DepositETH(uint256 chainId, address indexed from, address indexed to, uint256 value)
func (_L2Pool *L2PoolFilterer) WatchDepositETH(opts *bind.WatchOpts, sink chan<- *L2PoolDepositETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2Pool.contract.WatchLogs(opts, "DepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolDepositETH)
				if err := _L2Pool.contract.UnpackLog(event, "DepositETH", log); err != nil {
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

// ParseDepositETH is a log parse operation binding the contract event 0x25dbc63f879c7913a1e148949bf04761d9ff12ca9519a1c0ac9d854e8a90a1e9.
//
// Solidity: event DepositETH(uint256 chainId, address indexed from, address indexed to, uint256 value)
func (_L2Pool *L2PoolFilterer) ParseDepositETH(log types.Log) (*L2PoolDepositETH, error) {
	event := new(L2PoolDepositETH)
	if err := _L2Pool.contract.UnpackLog(event, "DepositETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolDepositWETHIterator is returned from FilterDepositWETH and is used to iterate over the raw logs and unpacked data for DepositWETH events raised by the L2Pool contract.
type L2PoolDepositWETHIterator struct {
	Event *L2PoolDepositWETH // Event containing the contract specifics and raw log

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
func (it *L2PoolDepositWETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolDepositWETH)
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
		it.Event = new(L2PoolDepositWETH)
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
func (it *L2PoolDepositWETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolDepositWETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolDepositWETH represents a DepositWETH event raised by the L2Pool contract.
type L2PoolDepositWETH struct {
	ChainId *big.Int
	From    common.Address
	To      common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositWETH is a free log retrieval operation binding the contract event 0x148ba123946a56f22f1acafa21440f6834fa36d19a4ad5d752f3647d45488541.
//
// Solidity: event DepositWETH(uint256 chainId, address indexed from, address indexed to, uint256 value)
func (_L2Pool *L2PoolFilterer) FilterDepositWETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2PoolDepositWETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2Pool.contract.FilterLogs(opts, "DepositWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolDepositWETHIterator{contract: _L2Pool.contract, event: "DepositWETH", logs: logs, sub: sub}, nil
}

// WatchDepositWETH is a free log subscription operation binding the contract event 0x148ba123946a56f22f1acafa21440f6834fa36d19a4ad5d752f3647d45488541.
//
// Solidity: event DepositWETH(uint256 chainId, address indexed from, address indexed to, uint256 value)
func (_L2Pool *L2PoolFilterer) WatchDepositWETH(opts *bind.WatchOpts, sink chan<- *L2PoolDepositWETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2Pool.contract.WatchLogs(opts, "DepositWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolDepositWETH)
				if err := _L2Pool.contract.UnpackLog(event, "DepositWETH", log); err != nil {
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

// ParseDepositWETH is a log parse operation binding the contract event 0x148ba123946a56f22f1acafa21440f6834fa36d19a4ad5d752f3647d45488541.
//
// Solidity: event DepositWETH(uint256 chainId, address indexed from, address indexed to, uint256 value)
func (_L2Pool *L2PoolFilterer) ParseDepositWETH(log types.Log) (*L2PoolDepositWETH, error) {
	event := new(L2PoolDepositWETH)
	if err := _L2Pool.contract.UnpackLog(event, "DepositWETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2Pool contract.
type L2PoolInitializedIterator struct {
	Event *L2PoolInitialized // Event containing the contract specifics and raw log

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
func (it *L2PoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolInitialized)
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
		it.Event = new(L2PoolInitialized)
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
func (it *L2PoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolInitialized represents a Initialized event raised by the L2Pool contract.
type L2PoolInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L2Pool *L2PoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2PoolInitializedIterator, error) {

	logs, sub, err := _L2Pool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2PoolInitializedIterator{contract: _L2Pool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L2Pool *L2PoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2PoolInitialized) (event.Subscription, error) {

	logs, sub, err := _L2Pool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolInitialized)
				if err := _L2Pool.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2Pool *L2PoolFilterer) ParseInitialized(log types.Log) (*L2PoolInitialized, error) {
	event := new(L2PoolInitialized)
	if err := _L2Pool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the L2Pool contract.
type L2PoolPausedIterator struct {
	Event *L2PoolPaused // Event containing the contract specifics and raw log

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
func (it *L2PoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolPaused)
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
		it.Event = new(L2PoolPaused)
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
func (it *L2PoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolPaused represents a Paused event raised by the L2Pool contract.
type L2PoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2Pool *L2PoolFilterer) FilterPaused(opts *bind.FilterOpts) (*L2PoolPausedIterator, error) {

	logs, sub, err := _L2Pool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &L2PoolPausedIterator{contract: _L2Pool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2Pool *L2PoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *L2PoolPaused) (event.Subscription, error) {

	logs, sub, err := _L2Pool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolPaused)
				if err := _L2Pool.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2Pool *L2PoolFilterer) ParsePaused(log types.Log) (*L2PoolPaused, error) {
	event := new(L2PoolPaused)
	if err := _L2Pool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the L2Pool contract.
type L2PoolRoleAdminChangedIterator struct {
	Event *L2PoolRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *L2PoolRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolRoleAdminChanged)
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
		it.Event = new(L2PoolRoleAdminChanged)
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
func (it *L2PoolRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolRoleAdminChanged represents a RoleAdminChanged event raised by the L2Pool contract.
type L2PoolRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_L2Pool *L2PoolFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*L2PoolRoleAdminChangedIterator, error) {

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

	logs, sub, err := _L2Pool.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolRoleAdminChangedIterator{contract: _L2Pool.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_L2Pool *L2PoolFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *L2PoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _L2Pool.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolRoleAdminChanged)
				if err := _L2Pool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_L2Pool *L2PoolFilterer) ParseRoleAdminChanged(log types.Log) (*L2PoolRoleAdminChanged, error) {
	event := new(L2PoolRoleAdminChanged)
	if err := _L2Pool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the L2Pool contract.
type L2PoolRoleGrantedIterator struct {
	Event *L2PoolRoleGranted // Event containing the contract specifics and raw log

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
func (it *L2PoolRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolRoleGranted)
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
		it.Event = new(L2PoolRoleGranted)
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
func (it *L2PoolRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolRoleGranted represents a RoleGranted event raised by the L2Pool contract.
type L2PoolRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2Pool *L2PoolFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L2PoolRoleGrantedIterator, error) {

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

	logs, sub, err := _L2Pool.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolRoleGrantedIterator{contract: _L2Pool.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2Pool *L2PoolFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *L2PoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2Pool.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolRoleGranted)
				if err := _L2Pool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_L2Pool *L2PoolFilterer) ParseRoleGranted(log types.Log) (*L2PoolRoleGranted, error) {
	event := new(L2PoolRoleGranted)
	if err := _L2Pool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the L2Pool contract.
type L2PoolRoleRevokedIterator struct {
	Event *L2PoolRoleRevoked // Event containing the contract specifics and raw log

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
func (it *L2PoolRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolRoleRevoked)
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
		it.Event = new(L2PoolRoleRevoked)
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
func (it *L2PoolRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolRoleRevoked represents a RoleRevoked event raised by the L2Pool contract.
type L2PoolRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2Pool *L2PoolFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L2PoolRoleRevokedIterator, error) {

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

	logs, sub, err := _L2Pool.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolRoleRevokedIterator{contract: _L2Pool.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2Pool *L2PoolFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *L2PoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2Pool.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolRoleRevoked)
				if err := _L2Pool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_L2Pool *L2PoolFilterer) ParseRoleRevoked(log types.Log) (*L2PoolRoleRevoked, error) {
	event := new(L2PoolRoleRevoked)
	if err := _L2Pool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the L2Pool contract.
type L2PoolUnpausedIterator struct {
	Event *L2PoolUnpaused // Event containing the contract specifics and raw log

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
func (it *L2PoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolUnpaused)
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
		it.Event = new(L2PoolUnpaused)
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
func (it *L2PoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolUnpaused represents a Unpaused event raised by the L2Pool contract.
type L2PoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2Pool *L2PoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*L2PoolUnpausedIterator, error) {

	logs, sub, err := _L2Pool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &L2PoolUnpausedIterator{contract: _L2Pool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2Pool *L2PoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *L2PoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _L2Pool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolUnpaused)
				if err := _L2Pool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2Pool *L2PoolFilterer) ParseUnpaused(log types.Log) (*L2PoolUnpaused, error) {
	event := new(L2PoolUnpaused)
	if err := _L2Pool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolWithdrawERC20toL1SuccessIterator is returned from FilterWithdrawERC20toL1Success and is used to iterate over the raw logs and unpacked data for WithdrawERC20toL1Success events raised by the L2Pool contract.
type L2PoolWithdrawERC20toL1SuccessIterator struct {
	Event *L2PoolWithdrawERC20toL1Success // Event containing the contract specifics and raw log

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
func (it *L2PoolWithdrawERC20toL1SuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolWithdrawERC20toL1Success)
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
		it.Event = new(L2PoolWithdrawERC20toL1Success)
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
func (it *L2PoolWithdrawERC20toL1SuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolWithdrawERC20toL1SuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolWithdrawERC20toL1Success represents a WithdrawERC20toL1Success event raised by the L2Pool contract.
type L2PoolWithdrawERC20toL1Success struct {
	ChainId   *big.Int
	Timestamp *big.Int
	Token     common.Address
	To        common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC20toL1Success is a free log retrieval operation binding the contract event 0x1bc1b1207e0c319b7baad0317c09f30575218af6f5a59f1f0f15de8d217672cc.
//
// Solidity: event WithdrawERC20toL1Success(uint256 chainId, uint256 timestamp, address token, address to, uint256 value)
func (_L2Pool *L2PoolFilterer) FilterWithdrawERC20toL1Success(opts *bind.FilterOpts) (*L2PoolWithdrawERC20toL1SuccessIterator, error) {

	logs, sub, err := _L2Pool.contract.FilterLogs(opts, "WithdrawERC20toL1Success")
	if err != nil {
		return nil, err
	}
	return &L2PoolWithdrawERC20toL1SuccessIterator{contract: _L2Pool.contract, event: "WithdrawERC20toL1Success", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20toL1Success is a free log subscription operation binding the contract event 0x1bc1b1207e0c319b7baad0317c09f30575218af6f5a59f1f0f15de8d217672cc.
//
// Solidity: event WithdrawERC20toL1Success(uint256 chainId, uint256 timestamp, address token, address to, uint256 value)
func (_L2Pool *L2PoolFilterer) WatchWithdrawERC20toL1Success(opts *bind.WatchOpts, sink chan<- *L2PoolWithdrawERC20toL1Success) (event.Subscription, error) {

	logs, sub, err := _L2Pool.contract.WatchLogs(opts, "WithdrawERC20toL1Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolWithdrawERC20toL1Success)
				if err := _L2Pool.contract.UnpackLog(event, "WithdrawERC20toL1Success", log); err != nil {
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

// ParseWithdrawERC20toL1Success is a log parse operation binding the contract event 0x1bc1b1207e0c319b7baad0317c09f30575218af6f5a59f1f0f15de8d217672cc.
//
// Solidity: event WithdrawERC20toL1Success(uint256 chainId, uint256 timestamp, address token, address to, uint256 value)
func (_L2Pool *L2PoolFilterer) ParseWithdrawERC20toL1Success(log types.Log) (*L2PoolWithdrawERC20toL1Success, error) {
	event := new(L2PoolWithdrawERC20toL1Success)
	if err := _L2Pool.contract.UnpackLog(event, "WithdrawERC20toL1Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolWithdrawETHtoL1SuccessIterator is returned from FilterWithdrawETHtoL1Success and is used to iterate over the raw logs and unpacked data for WithdrawETHtoL1Success events raised by the L2Pool contract.
type L2PoolWithdrawETHtoL1SuccessIterator struct {
	Event *L2PoolWithdrawETHtoL1Success // Event containing the contract specifics and raw log

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
func (it *L2PoolWithdrawETHtoL1SuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolWithdrawETHtoL1Success)
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
		it.Event = new(L2PoolWithdrawETHtoL1Success)
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
func (it *L2PoolWithdrawETHtoL1SuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolWithdrawETHtoL1SuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolWithdrawETHtoL1Success represents a WithdrawETHtoL1Success event raised by the L2Pool contract.
type L2PoolWithdrawETHtoL1Success struct {
	ChainId   *big.Int
	Timestamp *big.Int
	To        common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETHtoL1Success is a free log retrieval operation binding the contract event 0xf1fa6dbbb3e74f1bef249f22545b238726c1cc79c5571a1416575674c7e63f39.
//
// Solidity: event WithdrawETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_L2Pool *L2PoolFilterer) FilterWithdrawETHtoL1Success(opts *bind.FilterOpts) (*L2PoolWithdrawETHtoL1SuccessIterator, error) {

	logs, sub, err := _L2Pool.contract.FilterLogs(opts, "WithdrawETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return &L2PoolWithdrawETHtoL1SuccessIterator{contract: _L2Pool.contract, event: "WithdrawETHtoL1Success", logs: logs, sub: sub}, nil
}

// WatchWithdrawETHtoL1Success is a free log subscription operation binding the contract event 0xf1fa6dbbb3e74f1bef249f22545b238726c1cc79c5571a1416575674c7e63f39.
//
// Solidity: event WithdrawETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_L2Pool *L2PoolFilterer) WatchWithdrawETHtoL1Success(opts *bind.WatchOpts, sink chan<- *L2PoolWithdrawETHtoL1Success) (event.Subscription, error) {

	logs, sub, err := _L2Pool.contract.WatchLogs(opts, "WithdrawETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolWithdrawETHtoL1Success)
				if err := _L2Pool.contract.UnpackLog(event, "WithdrawETHtoL1Success", log); err != nil {
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

// ParseWithdrawETHtoL1Success is a log parse operation binding the contract event 0xf1fa6dbbb3e74f1bef249f22545b238726c1cc79c5571a1416575674c7e63f39.
//
// Solidity: event WithdrawETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_L2Pool *L2PoolFilterer) ParseWithdrawETHtoL1Success(log types.Log) (*L2PoolWithdrawETHtoL1Success, error) {
	event := new(L2PoolWithdrawETHtoL1Success)
	if err := _L2Pool.contract.UnpackLog(event, "WithdrawETHtoL1Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolWithdrawWETHtoL1SuccessIterator is returned from FilterWithdrawWETHtoL1Success and is used to iterate over the raw logs and unpacked data for WithdrawWETHtoL1Success events raised by the L2Pool contract.
type L2PoolWithdrawWETHtoL1SuccessIterator struct {
	Event *L2PoolWithdrawWETHtoL1Success // Event containing the contract specifics and raw log

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
func (it *L2PoolWithdrawWETHtoL1SuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolWithdrawWETHtoL1Success)
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
		it.Event = new(L2PoolWithdrawWETHtoL1Success)
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
func (it *L2PoolWithdrawWETHtoL1SuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolWithdrawWETHtoL1SuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolWithdrawWETHtoL1Success represents a WithdrawWETHtoL1Success event raised by the L2Pool contract.
type L2PoolWithdrawWETHtoL1Success struct {
	ChainId   *big.Int
	Timestamp *big.Int
	To        common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawWETHtoL1Success is a free log retrieval operation binding the contract event 0xcb1a101498aadad4159168d719f955e559e768912e6f5f5c57b094eb047d000f.
//
// Solidity: event WithdrawWETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_L2Pool *L2PoolFilterer) FilterWithdrawWETHtoL1Success(opts *bind.FilterOpts) (*L2PoolWithdrawWETHtoL1SuccessIterator, error) {

	logs, sub, err := _L2Pool.contract.FilterLogs(opts, "WithdrawWETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return &L2PoolWithdrawWETHtoL1SuccessIterator{contract: _L2Pool.contract, event: "WithdrawWETHtoL1Success", logs: logs, sub: sub}, nil
}

// WatchWithdrawWETHtoL1Success is a free log subscription operation binding the contract event 0xcb1a101498aadad4159168d719f955e559e768912e6f5f5c57b094eb047d000f.
//
// Solidity: event WithdrawWETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_L2Pool *L2PoolFilterer) WatchWithdrawWETHtoL1Success(opts *bind.WatchOpts, sink chan<- *L2PoolWithdrawWETHtoL1Success) (event.Subscription, error) {

	logs, sub, err := _L2Pool.contract.WatchLogs(opts, "WithdrawWETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolWithdrawWETHtoL1Success)
				if err := _L2Pool.contract.UnpackLog(event, "WithdrawWETHtoL1Success", log); err != nil {
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

// ParseWithdrawWETHtoL1Success is a log parse operation binding the contract event 0xcb1a101498aadad4159168d719f955e559e768912e6f5f5c57b094eb047d000f.
//
// Solidity: event WithdrawWETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_L2Pool *L2PoolFilterer) ParseWithdrawWETHtoL1Success(log types.Log) (*L2PoolWithdrawWETHtoL1Success, error) {
	event := new(L2PoolWithdrawWETHtoL1Success)
	if err := _L2Pool.contract.UnpackLog(event, "WithdrawWETHtoL1Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
