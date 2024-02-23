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

// L2PoolManagerMetaData contains all meta data concerning the L2PoolManager contract.
var L2PoolManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ChainIdIsNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"ChainIdNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorBlockChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"MinTransferAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LessThanMinTransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MantaNotWETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MantleNotWETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughETH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"NotEnoughToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"TokenIsNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"sourceChainIdError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20toL1Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawETHtoL1Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawWETHtoL1Success\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"BridgeInitiateETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FeePoolValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FundingPoolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"IsSupportChainId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"IsSupportToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_GAS_Limit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinTransferAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"QuickSendAssertToUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ReLayer\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SupportTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UpdateFundingPoolBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20ToL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawETHtoL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawWETHToL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MultisigWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageManager\",\"outputs\":[{\"internalType\":\"contractIMessageManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_MinTransferAmount\",\"type\":\"uint256\"}],\"name\":\"setMinTransferAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_PerFee\",\"type\":\"uint256\"}],\"name\":\"setPerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setSupportERC20Token\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setValidChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001c62000022565b620000d6565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000735760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d35780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b613eb180620000e66000396000f3fe60806040526004361061021a5760003560e01c8063650c227611610123578063b1887e9b116100ab578063d73f14e51161006f578063d73f14e514610619578063dac2956814610639578063dbb0481f14610669578063dd0c346014610696578063fa861848146106ac57600080fd5b8063b1887e9b14610591578063b92e6396146105a6578063cd01c665146105c6578063d28597f4146105e6578063d547741f146105f957600080fd5b806392d84738116100f257806392d84738146104f75780639b44238014610517578063a217fddf1461052a578063a525ae2b1461053f578063ab0f9e191461057157600080fd5b8063650c22761461047257806372fe6a7e146104925780638456cb59146104c257806391d14854146104d757600080fd5b806338731c47116101a6578063485cc95511610175578063485cc955146103c057806354dc027e146103e05780635765a4eb1461040d5780635c975abb1461042d578063626417b51461045257600080fd5b806338731c471461036557806339891dd7146103785780633b0230f01461038b5780633f4ba83a146103ab57600080fd5b806310875a13116101ed57806310875a13146102ad578063248a9ca3146102e55780632f2ff15d14610305578063325e5d431461032557806336568abe1461034557600080fd5b806301ffc9a71461021f5780630227047a1461025457806309dc480c146102675780630cd727df1461028b575b600080fd5b34801561022b57600080fd5b5061023f61023a3660046137ca565b6106c1565b60405190151581526020015b60405180910390f35b61023f610262366004613810565b6106f8565b34801561027357600080fd5b5061027d60015481565b60405190815260200161024b565b34801561029757600080fd5b506102ab6102a636600461384c565b610f6b565b005b3480156102b957600080fd5b506000546102cd906001600160a01b031681565b6040516001600160a01b03909116815260200161024b565b3480156102f157600080fd5b5061027d610300366004613876565b610fd6565b34801561031157600080fd5b506102ab61032036600461388f565b610ff8565b34801561033157600080fd5b506102ab610340366004613810565b61101a565b34801561035157600080fd5b506102ab61036036600461388f565b61107a565b61023f6103733660046138bb565b6110b2565b61023f61038636600461384c565b611296565b34801561039757600080fd5b5061023f6103a636600461390c565b6118f7565b3480156103b757600080fd5b506102ab611bfb565b3480156103cc57600080fd5b506102ab6103db36600461395a565b611c11565b3480156103ec57600080fd5b5061027d6103fb366004613984565b60066020526000908152604090205481565b34801561041957600080fd5b5061023f6104283660046138bb565b611d45565b34801561043957600080fd5b50600080516020613e5c8339815191525460ff1661023f565b34801561045e57600080fd5b506102cd61046d366004613876565b611f74565b34801561047e57600080fd5b506102ab61048d366004613876565b611f9e565b34801561049e57600080fd5b5061023f6104ad366004613876565b60009081526003602052604090205460ff1690565b3480156104ce57600080fd5b506102ab611faf565b3480156104e357600080fd5b5061023f6104f236600461388f565b611fc2565b34801561050357600080fd5b5061023f61051236600461399f565b611ffa565b61023f610525366004613a00565b612205565b34801561053657600080fd5b5061027d600081565b34801561054b57600080fd5b5060085461055c9063ffffffff1681565b60405163ffffffff909116815260200161024b565b34801561057d57600080fd5b506102ab61058c366004613a43565b612419565b34801561059d57600080fd5b506102cd6124a3565b3480156105b257600080fd5b506102ab6105c1366004613876565b6125a8565b3480156105d257600080fd5b5061023f6105e1366004613a7a565b6125fc565b61023f6105f436600461384c565b61297b565b34801561060557600080fd5b506102ab61061436600461388f565b61307a565b34801561062557600080fd5b506102ab610634366004613ab7565b613096565b34801561064557600080fd5b5061023f610654366004613984565b60046020526000908152604090205460ff1681565b34801561067557600080fd5b5061027d610684366004613984565b60056020526000908152604090205481565b3480156106a257600080fd5b5061027d60025481565b3480156106b857600080fd5b5061027d6130c2565b60006001600160e01b03198216637965db0b60e01b14806106f257506301ffc9a760e01b6001600160e01b03198316145b92915050565b600060ff196107166001600080516020613e1c833981519152613af2565b60405160200161072891815260200190565b60405160208183030381529060405280519060200120166107488161310a565b6001600160a01b038516600090815260046020526040902054469060ff16610793576040516305fd61ad60e01b81526001600160a01b03871660048201526024015b60405180910390fd5b6001600160a01b038616600090815260056020526040812080548692906107bb908490613af2565b9091555050620827508190036108e65760405163095ea7b360e01b81526001600160a01b0387169063095ea7b39061080d9073d8a791fe2be73eb6e6cf1eb0cb3f36adc9b3f8f9908890600401613b05565b6020604051808303816000875af115801561082c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108509190613b1e565b5060085460405163a93a4af960e01b81526001600160a01b038089166004830152871660248201526044810186905263ffffffff9091166064820181905273d8a791fe2be73eb6e6cf1eb0cb3f36adc9b3f8f99163a93a4af99190608401600060405180830381600088803b1580156108c857600080fd5b5087f11580156108dc573d6000803e3d6000fd5b5050505050610edc565b8061044d036109ec5760405163095ea7b360e01b81526001600160a01b0387169063095ea7b39061093190732a3dd3eb832af982ec71669e178424b10dca2ede908890600401613b05565b6020604051808303816000875af1158015610950573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109749190613b1e565b5060405163cd58657960e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede9063cd586579906109b590600090899089908c908490600401613b3b565b600060405180830381600087803b1580156109cf57600080fd5b505af11580156109e3573d6000803e3d6000fd5b50505050610edc565b80600a03610ad45760405163095ea7b360e01b81526001600160a01b0387169063095ea7b390610a29906010602160991b01908890600401613b05565b6020604051808303816000875af1158015610a48573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a6c9190613b1e565b50600854604051631474f2a960e31b81526010602160991b019163a3a79548918791610aa7918b918b91859163ffffffff1690600401613b82565b6000604051808303818588803b158015610ac057600080fd5b505af11580156108dc573d6000803e3d6000fd5b8061a4b103610beb5760405163095ea7b360e01b81526001600160a01b0387169063095ea7b390610b1f907309e9222e96e7b4ae2a407b98d48e330053351eee908890600401613b05565b6020604051808303816000875af1158015610b3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b629190613b1e565b50604051637b3a3c8b60e01b81527309e9222e96e7b4ae2a407b98d48e330053351eee90637b3a3c8b90610b9e90899089908990600401613bbf565b6000604051808303816000875af1158015610bbd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610be59190810190613c2c565b50610edc565b8061a4ba03610cb55760405163095ea7b360e01b81526001600160a01b0387169063095ea7b390610c369073b2535b988dce19f9d71dfb22db6da744acac21bf908890600401613b05565b6020604051808303816000875af1158015610c55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c799190613b1e565b50604051637b3a3c8b60e01b815273b2535b988dce19f9d71dfb22db6da744acac21bf90637b3a3c8b90610b9e90899089908990600401613bbf565b8061014403610cfb57604051636ce5768960e11b81527311f943b2c77b743ab90f4a0ae7d5a4e7fca3e1029063d9caed12908690610aa79089908b908490600401613cd9565b8061138803610db75760405163095ea7b360e01b81526001600160a01b0387169063095ea7b390610d39906010602160991b01908890600401613b05565b6020604051808303816000875af1158015610d58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d7c9190613b1e565b5060085460405163297e27a960e11b81526010602160991b01916352fc4f52916109b5918a918a918a9163ffffffff90911690600401613b82565b8060a903610df45760405163095ea7b360e01b81526001600160a01b0387169063095ea7b390610a29906010602160991b01908890600401613b05565b8061a70e03610ec35760405163095ea7b360e01b81526001600160a01b0387169063095ea7b390610e3f90732c4813276869d93afdab4dd2b01cd670342da194908890600401613b05565b6020604051808303816000875af1158015610e5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e829190613b1e565b5060405163cd58657960e01b8152732c4813276869d93afdab4dd2b01cd670342da1949063cd586579906109b590600090899089908c908490600401613b3b565b604051639474dee160e01b815260040160405180910390fd5b6001600160a01b03861660009081526005602052604081208054869290610f04908490613af2565b9091555050604080514681524260208201526001600160a01b0388811682840152871660608201526080810186905290517f1bc1b1207e0c319b7baad0317c09f30575218af6f5a59f1f0f15de8d217672cc9181900360a00190a150600195945050505050565b60ff19610f876001600080516020613e1c833981519152613af2565b604051602001610f9991815260200190565b6040516020818303038152906040528051906020012016610fb98161310a565b506001600160a01b03909116600090815260056020526040902055565b6000908152600080516020613e3c833981519152602052604090206001015490565b61100182610fd6565b61100a8161310a565b6110148383613114565b50505050565b60ff196110366001600080516020613e1c833981519152613af2565b60405160200161104891815260200190565b60405160208183030381529060405280519060200120166110688161310a565b6110738484846131b9565b5050505050565b6001600160a01b03811633146110a35760405163334bd91960e11b815260040160405180910390fd5b6110ad828261335e565b505050565b600060ff196110d06001600080516020613e1c833981519152613af2565b6040516020016110e291815260200190565b60405160208183030381529060405280519060200120166111028161310a565b46871461112257604051631a26aa4d60e21b815260040160405180910390fd5b60008881526003602052604090205460ff1661115457604051632ef6974160e11b81526004810189905260240161078a565b6040516001600160a01b0387169086156108fc029087906000818181858888f1935050505015801561118a573d6000803e3d6000fd5b5073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee60009081526005602052600080516020613ddc83398151915280548792906111c9908490613af2565b9091555050600054604051631eb65ffb60e01b81526001600160a01b0390911690631eb65ffb90611208908b908b908b908a908c908b90600401613cfd565b600060405180830381600087803b15801561122257600080fd5b505af1158015611236573d6000803e3d6000fd5b5050604080518b8152602081018b90529081018890526001600160a01b03891692503091507f61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f329060600160405180910390a3506001979650505050505050565b600060ff196112b46001600080516020613e1c833981519152613af2565b6040516020016112c691815260200190565b60405160208183030381529060405280519060200120166112e68161310a565b464784111561130857604051632c1d501360e11b815260040160405180910390fd5b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee60009081526005602052600080516020613ddc8339815191528054869290611346908490613af2565b9091555050620827508190036113e557600854604051636705b1e760e11b81526001600160a01b03871660048201526024810186905263ffffffff90911660448201819052737003e7b7186f0e6601203b99f7b8decbfa391cf99163ce0b63ce919087906064016000604051808303818589803b1580156113c657600080fd5b5088f11580156113da573d6000803e3d6000fd5b50505050505061185b565b8061044d036114675760405163cd58657960e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede9063cd586579908690611430906000908a90849083908190600401613b3b565b6000604051808303818588803b15801561144957600080fd5b505af115801561145d573d6000803e3d6000fd5b505050505061185b565b80600a036114bd57600854604051631474f2a960e31b81526010602160991b019163a3a795489187916114309173deaddeaddeaddeaddeaddeaddeaddeaddead0000918b91859163ffffffff1690600401613b82565b8061a4b10361156457604051637b3a3c8b60e01b8152735288c571fd7ad117bea99bf60fe0846c4e84f93390637b3a3c8b9086906115179073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee908a908490600401613bbf565b60006040518083038185885af1158015611535573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f1916820160405261155e9190810190613c2c565b5061185b565b8061a4ba036115be57604051637b3a3c8b60e01b81527321903d3f8176b1a0c17e953cd896610be9ffdfa890637b3a3c8b9086906115179073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee908a908490600401613bbf565b806101440361160557604051636ce5768960e11b81527311f943b2c77b743ab90f4a0ae7d5a4e7fca3e1029063d9caed129086906114309089906000908490600401613cd9565b80611388036117175760405163095ea7b360e01b815273deaddeaddeaddeaddeaddeaddeaddeaddead11119063095ea7b39061164e906010602160991b01908890600401613b05565b6020604051808303816000875af115801561166d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116919190613b1e565b5060085460405163297e27a960e11b81526010602160991b01916352fc4f52916116e09173deaddeaddeaddeaddeaddeaddeaddeaddead1111918a918a9163ffffffff90911690600401613b82565b600060405180830381600087803b1580156116fa57600080fd5b505af115801561170e573d6000803e3d6000fd5b5050505061185b565b8060a90361176d57600854604051631474f2a960e31b81526010602160991b019163a3a795489187916114309173deaddeaddeaddeaddeaddeaddeaddeaddead0000918b91859163ffffffff1690600401613b82565b8061a70e03610ec35760405163095ea7b360e01b8152734b21b980d0dc7d3c0c6175b0a412694f3a1c7c6b9063095ea7b3906117c390732c4813276869d93afdab4dd2b01cd670342da194908890600401613b05565b6020604051808303816000875af11580156117e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118069190613b1e565b5060405163cd58657960e01b8152732c4813276869d93afdab4dd2b01cd670342da1949063cd586579906116e09060009089908990734b21b980d0dc7d3c0c6175b0a412694f3a1c7c6b908490600401613b3b565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee60009081526005602052600080516020613ddc8339815191528054869290611899908490613af2565b9091555050604080514681524260208201526001600160a01b038716818301526060810186905290517ff1fa6dbbb3e74f1bef249f22545b238726c1cc79c5571a1416575674c7e63f399181900360800190a1506001949350505050565b600046861461191957604051631a26aa4d60e21b815260040160405180910390fd5b60008581526003602052604090205460ff1661194b57604051632ef6974160e11b81526004810186905260240161078a565b6001600160a01b03831660009081526004602052604090205460ff1661198f576040516305fd61ad60e01b81526001600160a01b038416600482015260240161078a565b6040516370a0823160e01b81523060048201526000906001600160a01b038516906370a0823190602401602060405180830381865afa1580156119d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119fa9190613d2e565b9050611a116001600160a01b0385163330866133da565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa158015611a58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a7c9190613d2e565b90506000611a8a8383613af2565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee60009081526005602052600080516020613ddc83398151915280549293508792909190611acd908490613d47565b9091555050600254600090620f424090611ae79084613d5a565b611af19190613d71565b9050611afd8183613af2565b6001600160a01b038816600090815260066020526040812080549294508392909190611b2a908490613d47565b909155505060005460405163305f478560e21b81526001600160a01b039091169063c17d1e1490611b67908d908d908d908c908890600401613d93565b600060405180830381600087803b158015611b8157600080fd5b505af1158015611b95573d6000803e3d6000fd5b5050604080518d8152602081018d90529081018590526001600160a01b03808c1693503392508a16907fece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a49060600160405180910390a45060019998505050505050505050565b6000611c068161310a565b611c0e613434565b50565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff16600081158015611c575750825b905060008267ffffffffffffffff166001148015611c745750303b155b905081158015611c82575080155b15611ca05760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315611cca57845460ff60401b1916600160401b1785555b611cd2613494565b611cda61349e565b611ce487876134ae565b6008805463ffffffff1916620493e01790558315611d3c57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b600060ff19611d636001600080516020613e1c833981519152613af2565b604051602001611d7591815260200190565b6040516020818303038152906040528051906020012016611d958161310a565b468714611db557604051631a26aa4d60e21b815260040160405180910390fd5b60008881526003602052604090205460ff16611de757604051632ef6974160e11b81526004810189905260240161078a565b6000611df16124a3565b6040516323b872dd60e01b81529091506001600160a01b038216906323b872dd90611e249030908b908b90600401613cd9565b6020604051808303816000875af1158015611e43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e679190613b1e565b5073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc260009081526005602052600080516020613dfc8339815191528054889290611ea6908490613af2565b9091555050600054604051631eb65ffb60e01b81526001600160a01b0390911690631eb65ffb90611ee5908c908c908c908b908d908c90600401613cfd565b600060405180830381600087803b158015611eff57600080fd5b505af1158015611f13573d6000803e3d6000fd5b5050604080518c8152602081018c90529081018990526001600160a01b038a1692503091507f7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc149060600160405180910390a350600198975050505050505050565b60078181548110611f8457600080fd5b6000918252602090912001546001600160a01b0316905081565b6000611fa98161310a565b50600255565b6000611fba8161310a565b611c0e6134f7565b6000918252600080516020613e3c833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b600060ff196120186001600080516020613e1c833981519152613af2565b60405160200161202a91815260200190565b604051602081830303815290604052805190602001201661204a8161310a565b46881461206a57604051631a26aa4d60e21b815260040160405180910390fd5b60008981526003602052604090205460ff1661209c57604051632ef6974160e11b8152600481018a905260240161078a565b6001600160a01b03861660009081526004602052604090205460ff166120e0576040516305fd61ad60e01b81526001600160a01b038716600482015260240161078a565b6120f56001600160a01b0387163089886133da565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee60009081526005602052600080516020613ddc8339815191528054879290612133908490613af2565b9091555050600054604051631eb65ffb60e01b81526001600160a01b0390911690631eb65ffb90612172908c908c908c908a908c908b90600401613cfd565b600060405180830381600087803b15801561218c57600080fd5b505af11580156121a0573d6000803e3d6000fd5b5050604080518c8152602081018c90529081018890526001600160a01b03808b1693503092508916907f0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe9060600160405180910390a450600198975050505050505050565b600046841461222757604051631a26aa4d60e21b815260040160405180910390fd5b60008381526003602052604090205460ff1661225957604051632ef6974160e11b81526004810184905260240161078a565b600154341015612289576001546040516358f8331360e11b8152600481019190915234602482015260440161078a565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee60009081526005602052600080516020613ddc83398151915280543492906122c7908490613d47565b9091555050600254600090620f4240906122e19034613d5a565b6122eb9190613d71565b905060006122f98234613af2565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee600090815260066020527fa2e5aefc6e2cbe2917a296f0fd89c5f915c487c803db1d98eccb43f14012d7118054929350849290919061234e908490613d47565b909155505060005460405163305f478560e21b81526001600160a01b039091169063c17d1e149061238b9046908990899087908990600401613d93565b600060405180830381600087803b1580156123a557600080fd5b505af11580156123b9573d6000803e3d6000fd5b505060408051898152602081018990529081018490526001600160a01b03871692503391507f0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc59060600160405180910390a36001925050505b9392505050565b60006124248161310a565b6001600160a01b0383166000908152600460205260409020805460ff191683158015919091179091556110ad57600780546001810182556000919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880180546001600160a01b0385166001600160a01b0319909116179055505050565b600046620827508190036124bf576004605360981b0191505090565b8061044d036124e357734f9a0e7fd2bf6067db6994cf12e4495df938e6e991505090565b80600a036124f9576006602160991b0191505090565b8061a4b10361251d577382af49447d8a07e3bd95bd0d56f35241523fbab191505090565b8061a4ba036125415773722e8bdd2ce80a4422e880164f2079488e11536591505090565b806101440361256557735aea5775959fbc2557cc8789bc1bf90a239d9a9191505090565b80611388036125875760405163e31d668960e01b815260040160405180910390fd5b8060a903610ec3576040516322c20c6960e11b815260040160405180910390fd5b60ff196125c46001600080516020613e1c833981519152613af2565b6040516020016125d691815260200190565b60405160208183030381529060405280519060200120166125f68161310a565b50600155565b600046851461261e57604051631a26aa4d60e21b815260040160405180910390fd5b60008481526003602052604090205460ff16612650576040516363b5c9db60e01b81526004810185905260240161078a565b600061265a6124a3565b6040516370a0823160e01b81523060048201529091506000906001600160a01b038316906370a0823190602401602060405180830381865afa1580156126a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126c89190613d2e565b6040516323b872dd60e01b81529091506001600160a01b038316906323b872dd906126fb90339030908990600401613cd9565b6020604051808303816000875af115801561271a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061273e9190613b1e565b506040516370a0823160e01b81523060048201526000906001600160a01b038416906370a0823190602401602060405180830381865afa158015612786573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127aa9190613d2e565b905060006127b88383613af2565b90506001548110156127eb576001546040516358f8331360e11b815260048101919091526024810182905260440161078a565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc260009081526005602052600080516020613dfc8339815191528054839290612829908490613d47565b9091555050600254600090620f4240906128439084613d5a565b61284d9190613d71565b90506128598183613af2565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2600090815260066020527f5e5777fab7622aff3c042c1ece74307c2e9d699a9da444f416c35f2e1def28a5805492945083929091906128ae908490613d47565b909155505060005460405163305f478560e21b81526001600160a01b039091169063c17d1e14906128eb908d908d908d908d908890600401613d93565b600060405180830381600087803b15801561290557600080fd5b505af1158015612919573d6000803e3d6000fd5b5050604080518d8152602081018d90529081018590526001600160a01b038b1692503391507fcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a899060600160405180910390a35060019998505050505050505050565b600060ff196129996001600080516020613e1c833981519152613af2565b6040516020016129ab91815260200190565b60405160208183030381529060405280519060200120166129cb8161310a565b4660006129d66124a3565b6040516370a0823160e01b81523060048201529091506001600160a01b038216906370a0823190602401602060405180830381865afa158015612a1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a419190613d2e565b851115612a6c576040516311745b6160e21b81526001600160a01b038216600482015260240161078a565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc260009081526005602052600080516020613dfc8339815191528054879290612aaa908490613af2565b909155505062082750829003612bd55760405163095ea7b360e01b81526001600160a01b0382169063095ea7b390612afc90737003e7b7186f0e6601203b99f7b8decbfa391cf9908990600401613b05565b6020604051808303816000875af1158015612b1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b3f9190613b1e565b5060085460405163a93a4af960e01b81526001600160a01b038084166004830152881660248201526044810187905263ffffffff90911660648201819052737003e7b7186f0e6601203b99f7b8decbfa391cf99163a93a4af99190608401600060405180830381600088803b158015612bb757600080fd5b5087f1158015612bcb573d6000803e3d6000fd5b5050505050612fdd565b8161044d03612cd35760405163095ea7b360e01b81526001600160a01b0382169063095ea7b390612c2090732a3dd3eb832af982ec71669e178424b10dca2ede908990600401613b05565b6020604051808303816000875af1158015612c3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c639190613b1e565b5060405163cd58657960e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede9063cd586579908790612ca6906000908b90849083908190600401613b3b565b6000604051808303818588803b158015612cbf57600080fd5b505af1158015612bcb573d6000803e3d6000fd5b81600a03612d8e5760405163095ea7b360e01b81526001600160a01b0382169063095ea7b390612d10906010602160991b01908990600401613b05565b6020604051808303816000875af1158015612d2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d539190613b1e565b50600854604051631474f2a960e31b81526010602160991b019163a3a79548918891612ca69186918c91859163ffffffff1690600401613b82565b8161a4b103612eb95760405163095ea7b360e01b81526001600160a01b0382169063095ea7b390612dd990736c411ad3e74de3e7bd422b94a27770f5b86c623b908990600401613b05565b6020604051808303816000875af1158015612df8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e1c9190613b1e565b50604051637b3a3c8b60e01b8152736c411ad3e74de3e7bd422b94a27770f5b86c623b90637b3a3c8b90612e6c9073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2908a908a90600401613bbf565b6000604051808303816000875af1158015612e8b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612eb39190810190613c2c565b50612fdd565b8161a4ba03612f975760405163095ea7b360e01b81526001600160a01b0382169063095ea7b390612f0490737626841cb6113412f9c88d3adc720c9fac88d9ed908990600401613b05565b6020604051808303816000875af1158015612f23573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f479190613b1e565b50604051637b3a3c8b60e01b8152737626841cb6113412f9c88d3adc720c9fac88d9ed90637b3a3c8b90612e6c9073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2908a908a90600401613bbf565b8161014403610ec357604051636ce5768960e11b81527311f943b2c77b743ab90f4a0ae7d5a4e7fca3e1029063d9caed12908790612ca6908a9086908490600401613cd9565b604080514681524260208201526001600160a01b038816818301526060810187905290517fcb1a101498aadad4159168d719f955e559e768912e6f5f5c57b094eb047d000f9181900360800190a173c02aaa39b223fe8d0a0e5c4f27ead9083c756cc260009081526005602052600080516020613dfc8339815191528054879290613069908490613af2565b909155506001979650505050505050565b61308382610fd6565b61308c8161310a565b611014838361335e565b60006130a18161310a565b50600091825260036020526040909120805460ff1916911515919091179055565b60ff196130de6001600080516020613e1c833981519152613af2565b6040516020016130f091815260200190565b604051602081830303815290604052805190602001201681565b611c0e8133613540565b6000600080516020613e3c83398151915261312f8484611fc2565b6131af576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556131653390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506106f2565b60009150506106f2565b6001600160a01b03831660009081526004602052604081205460ff166131fd576040516305fd61ad60e01b81526001600160a01b038516600482015260240161078a565b6001600160a01b03841660009081526005602052604081208054849290613225908490613af2565b909155505073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeed196001600160a01b038516016132ac578147101561327057604051632c1d501360e11b815260040160405180910390fd5b6040516001600160a01b0384169083156108fc029084906000818181858888f193505050501580156132a6573d6000803e3d6000fd5b50613354565b6040516370a0823160e01b815230600482015282906001600160a01b038616906370a0823190602401602060405180830381865afa1580156132f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133169190613d2e565b1015613340576040516311745b6160e21b81526001600160a01b038516600482015260240161078a565b6133546001600160a01b038516848461356f565b5060019392505050565b6000600080516020613e3c8339815191526133798484611fc2565b156131af576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506106f2565b61101484856001600160a01b03166323b872dd86868660405160240161340293929190613cd9565b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613595565b61343c6135f8565b600080516020613e5c833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b61349c613628565b565b6134a6613628565b61349c613671565b6134b6613628565b67016345785d8a00006001556127106002556134d3600083613114565b50600080546001600160a01b0319166001600160a01b039290921691909117905550565b6134ff613692565b600080516020613e5c833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833613476565b61354a8282611fc2565b61356b57808260405163e2517d3f60e01b815260040161078a929190613b05565b5050565b6110ad83846001600160a01b031663a9059cbb8585604051602401613402929190613b05565b60006135aa6001600160a01b038416836136c3565b905080516000141580156135cf5750808060200190518101906135cd9190613b1e565b155b156110ad57604051635274afe760e01b81526001600160a01b038416600482015260240161078a565b600080516020613e5c8339815191525460ff1661349c57604051638dfc202b60e01b815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661349c57604051631afcd79f60e31b815260040160405180910390fd5b613679613628565b600080516020613e5c833981519152805460ff19169055565b600080516020613e5c8339815191525460ff161561349c5760405163d93c066560e01b815260040160405180910390fd5b60606124128383600084600080856001600160a01b031684866040516136e99190613dbf565b60006040518083038185875af1925050503d8060008114613726576040519150601f19603f3d011682016040523d82523d6000602084013e61372b565b606091505b509150915061373b868383613745565b9695505050505050565b60608261375a57613755826137a1565b612412565b815115801561377157506001600160a01b0384163b155b1561379a57604051639996b31560e01b81526001600160a01b038516600482015260240161078a565b5080612412565b8051156137b15780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b6000602082840312156137dc57600080fd5b81356001600160e01b03198116811461241257600080fd5b80356001600160a01b038116811461380b57600080fd5b919050565b60008060006060848603121561382557600080fd5b61382e846137f4565b925061383c602085016137f4565b9150604084013590509250925092565b6000806040838503121561385f57600080fd5b613868836137f4565b946020939093013593505050565b60006020828403121561388857600080fd5b5035919050565b600080604083850312156138a257600080fd5b823591506138b2602084016137f4565b90509250929050565b60008060008060008060c087890312156138d457600080fd5b86359550602087013594506138eb604088016137f4565b9350606087013592506080870135915060a087013590509295509295509295565b600080600080600060a0868803121561392457600080fd5b853594506020860135935061393b604087016137f4565b9250613949606087016137f4565b949793965091946080013592915050565b6000806040838503121561396d57600080fd5b613976836137f4565b91506138b2602084016137f4565b60006020828403121561399657600080fd5b612412826137f4565b600080600080600080600060e0888a0312156139ba57600080fd5b87359650602088013595506139d1604089016137f4565b94506139df606089016137f4565b9699959850939660808101359560a0820135955060c0909101359350915050565b600080600060608486031215613a1557600080fd5b8335925060208401359150613a2c604085016137f4565b90509250925092565b8015158114611c0e57600080fd5b60008060408385031215613a5657600080fd5b613a5f836137f4565b91506020830135613a6f81613a35565b809150509250929050565b60008060008060808587031215613a9057600080fd5b8435935060208501359250613aa7604086016137f4565b9396929550929360600135925050565b60008060408385031215613aca57600080fd5b823591506020830135613a6f81613a35565b634e487b7160e01b600052601160045260246000fd5b818103818111156106f2576106f2613adc565b6001600160a01b03929092168252602082015260400190565b600060208284031215613b3057600080fd5b815161241281613a35565b63ffffffff9590951685526001600160a01b039384166020860152604085019290925290911660608301521515608082015260c060a0820181905260009082015260e00190565b6001600160a01b039485168152929093166020830152604082015263ffffffff909116606082015260a06080820181905260009082015260c00190565b6001600160a01b039384168152919092166020820152604081019190915260806060820181905260009082015260a00190565b634e487b7160e01b600052604160045260246000fd5b60005b83811015613c23578181015183820152602001613c0b565b50506000910152565b600060208284031215613c3e57600080fd5b815167ffffffffffffffff80821115613c5657600080fd5b818401915084601f830112613c6a57600080fd5b815181811115613c7c57613c7c613bf2565b604051601f8201601f19908116603f01168101908382118183101715613ca457613ca4613bf2565b81604052828152876020848701011115613cbd57600080fd5b613cce836020830160208801613c08565b979650505050505050565b6001600160a01b039384168152919092166020820152604081019190915260600190565b95865260208601949094526001600160a01b039290921660408501526060840152608083015260a082015260c00190565b600060208284031215613d4057600080fd5b5051919050565b808201808211156106f2576106f2613adc565b80820281158282048414176106f2576106f2613adc565b600082613d8e57634e487b7160e01b600052601260045260246000fd5b500490565b94855260208501939093526001600160a01b039190911660408401526060830152608082015260a00190565b60008251613dd1818460208701613c08565b919091019291505056fea1829a9003092132f585b6ccdd167c19fe9774dbdea4260287e8a8e8ca8185d7a550ba85c46b24b567d2e17cd597f2283877afab43603f46d5de7858f1bdb73133fe247d78fee2e7fd135c405eda4bd2911c0a73c0a81b36c3bcc967dd06e5ae02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300a2646970667358221220a9575be29da1e8ce38e3c6c408fdc6182a3ecbbd19f369f2e23b61fec0c5f41a64736f6c63430008140033",
}

// L2PoolManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use L2PoolManagerMetaData.ABI instead.
var L2PoolManagerABI = L2PoolManagerMetaData.ABI

// L2PoolManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2PoolManagerMetaData.Bin instead.
var L2PoolManagerBin = L2PoolManagerMetaData.Bin

// DeployL2PoolManager deploys a new Ethereum contract, binding an instance of L2PoolManager to it.
func DeployL2PoolManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L2PoolManager, error) {
	parsed, err := L2PoolManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2PoolManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2PoolManager{L2PoolManagerCaller: L2PoolManagerCaller{contract: contract}, L2PoolManagerTransactor: L2PoolManagerTransactor{contract: contract}, L2PoolManagerFilterer: L2PoolManagerFilterer{contract: contract}}, nil
}

// L2PoolManager is an auto generated Go binding around an Ethereum contract.
type L2PoolManager struct {
	L2PoolManagerCaller     // Read-only binding to the contract
	L2PoolManagerTransactor // Write-only binding to the contract
	L2PoolManagerFilterer   // Log filterer for contract events
}

// L2PoolManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2PoolManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2PoolManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2PoolManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2PoolManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2PoolManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2PoolManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2PoolManagerSession struct {
	Contract     *L2PoolManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2PoolManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2PoolManagerCallerSession struct {
	Contract *L2PoolManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L2PoolManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2PoolManagerTransactorSession struct {
	Contract     *L2PoolManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L2PoolManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2PoolManagerRaw struct {
	Contract *L2PoolManager // Generic contract binding to access the raw methods on
}

// L2PoolManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2PoolManagerCallerRaw struct {
	Contract *L2PoolManagerCaller // Generic read-only contract binding to access the raw methods on
}

// L2PoolManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2PoolManagerTransactorRaw struct {
	Contract *L2PoolManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2PoolManager creates a new instance of L2PoolManager, bound to a specific deployed contract.
func NewL2PoolManager(address common.Address, backend bind.ContractBackend) (*L2PoolManager, error) {
	contract, err := bindL2PoolManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2PoolManager{L2PoolManagerCaller: L2PoolManagerCaller{contract: contract}, L2PoolManagerTransactor: L2PoolManagerTransactor{contract: contract}, L2PoolManagerFilterer: L2PoolManagerFilterer{contract: contract}}, nil
}

// NewL2PoolManagerCaller creates a new read-only instance of L2PoolManager, bound to a specific deployed contract.
func NewL2PoolManagerCaller(address common.Address, caller bind.ContractCaller) (*L2PoolManagerCaller, error) {
	contract, err := bindL2PoolManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerCaller{contract: contract}, nil
}

// NewL2PoolManagerTransactor creates a new write-only instance of L2PoolManager, bound to a specific deployed contract.
func NewL2PoolManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*L2PoolManagerTransactor, error) {
	contract, err := bindL2PoolManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerTransactor{contract: contract}, nil
}

// NewL2PoolManagerFilterer creates a new log filterer instance of L2PoolManager, bound to a specific deployed contract.
func NewL2PoolManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*L2PoolManagerFilterer, error) {
	contract, err := bindL2PoolManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerFilterer{contract: contract}, nil
}

// bindL2PoolManager binds a generic wrapper to an already deployed contract.
func bindL2PoolManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2PoolManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2PoolManager *L2PoolManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2PoolManager.Contract.L2PoolManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2PoolManager *L2PoolManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2PoolManager.Contract.L2PoolManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2PoolManager *L2PoolManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2PoolManager.Contract.L2PoolManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2PoolManager *L2PoolManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2PoolManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2PoolManager *L2PoolManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2PoolManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2PoolManager *L2PoolManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2PoolManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L2PoolManager *L2PoolManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L2PoolManager *L2PoolManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _L2PoolManager.Contract.DEFAULTADMINROLE(&_L2PoolManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L2PoolManager *L2PoolManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _L2PoolManager.Contract.DEFAULTADMINROLE(&_L2PoolManager.CallOpts)
}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_L2PoolManager *L2PoolManagerCaller) FeePoolValue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "FeePoolValue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_L2PoolManager *L2PoolManagerSession) FeePoolValue(arg0 common.Address) (*big.Int, error) {
	return _L2PoolManager.Contract.FeePoolValue(&_L2PoolManager.CallOpts, arg0)
}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_L2PoolManager *L2PoolManagerCallerSession) FeePoolValue(arg0 common.Address) (*big.Int, error) {
	return _L2PoolManager.Contract.FeePoolValue(&_L2PoolManager.CallOpts, arg0)
}

// FundingPoolBalance is a free data retrieval call binding the contract method 0xdbb0481f.
//
// Solidity: function FundingPoolBalance(address ) view returns(uint256)
func (_L2PoolManager *L2PoolManagerCaller) FundingPoolBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "FundingPoolBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingPoolBalance is a free data retrieval call binding the contract method 0xdbb0481f.
//
// Solidity: function FundingPoolBalance(address ) view returns(uint256)
func (_L2PoolManager *L2PoolManagerSession) FundingPoolBalance(arg0 common.Address) (*big.Int, error) {
	return _L2PoolManager.Contract.FundingPoolBalance(&_L2PoolManager.CallOpts, arg0)
}

// FundingPoolBalance is a free data retrieval call binding the contract method 0xdbb0481f.
//
// Solidity: function FundingPoolBalance(address ) view returns(uint256)
func (_L2PoolManager *L2PoolManagerCallerSession) FundingPoolBalance(arg0 common.Address) (*big.Int, error) {
	return _L2PoolManager.Contract.FundingPoolBalance(&_L2PoolManager.CallOpts, arg0)
}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_L2PoolManager *L2PoolManagerCaller) IsSupportChainId(opts *bind.CallOpts, chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "IsSupportChainId", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_L2PoolManager *L2PoolManagerSession) IsSupportChainId(chainId *big.Int) (bool, error) {
	return _L2PoolManager.Contract.IsSupportChainId(&_L2PoolManager.CallOpts, chainId)
}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_L2PoolManager *L2PoolManagerCallerSession) IsSupportChainId(chainId *big.Int) (bool, error) {
	return _L2PoolManager.Contract.IsSupportChainId(&_L2PoolManager.CallOpts, chainId)
}

// IsSupportToken is a free data retrieval call binding the contract method 0xdac29568.
//
// Solidity: function IsSupportToken(address ) view returns(bool)
func (_L2PoolManager *L2PoolManagerCaller) IsSupportToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "IsSupportToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportToken is a free data retrieval call binding the contract method 0xdac29568.
//
// Solidity: function IsSupportToken(address ) view returns(bool)
func (_L2PoolManager *L2PoolManagerSession) IsSupportToken(arg0 common.Address) (bool, error) {
	return _L2PoolManager.Contract.IsSupportToken(&_L2PoolManager.CallOpts, arg0)
}

// IsSupportToken is a free data retrieval call binding the contract method 0xdac29568.
//
// Solidity: function IsSupportToken(address ) view returns(bool)
func (_L2PoolManager *L2PoolManagerCallerSession) IsSupportToken(arg0 common.Address) (bool, error) {
	return _L2PoolManager.Contract.IsSupportToken(&_L2PoolManager.CallOpts, arg0)
}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_L2PoolManager *L2PoolManagerCaller) L2WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "L2WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_L2PoolManager *L2PoolManagerSession) L2WETH() (common.Address, error) {
	return _L2PoolManager.Contract.L2WETH(&_L2PoolManager.CallOpts)
}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_L2PoolManager *L2PoolManagerCallerSession) L2WETH() (common.Address, error) {
	return _L2PoolManager.Contract.L2WETH(&_L2PoolManager.CallOpts)
}

// MAXGASLimit is a free data retrieval call binding the contract method 0xa525ae2b.
//
// Solidity: function MAX_GAS_Limit() view returns(uint32)
func (_L2PoolManager *L2PoolManagerCaller) MAXGASLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "MAX_GAS_Limit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXGASLimit is a free data retrieval call binding the contract method 0xa525ae2b.
//
// Solidity: function MAX_GAS_Limit() view returns(uint32)
func (_L2PoolManager *L2PoolManagerSession) MAXGASLimit() (uint32, error) {
	return _L2PoolManager.Contract.MAXGASLimit(&_L2PoolManager.CallOpts)
}

// MAXGASLimit is a free data retrieval call binding the contract method 0xa525ae2b.
//
// Solidity: function MAX_GAS_Limit() view returns(uint32)
func (_L2PoolManager *L2PoolManagerCallerSession) MAXGASLimit() (uint32, error) {
	return _L2PoolManager.Contract.MAXGASLimit(&_L2PoolManager.CallOpts)
}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_L2PoolManager *L2PoolManagerCaller) MinTransferAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "MinTransferAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_L2PoolManager *L2PoolManagerSession) MinTransferAmount() (*big.Int, error) {
	return _L2PoolManager.Contract.MinTransferAmount(&_L2PoolManager.CallOpts)
}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_L2PoolManager *L2PoolManagerCallerSession) MinTransferAmount() (*big.Int, error) {
	return _L2PoolManager.Contract.MinTransferAmount(&_L2PoolManager.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L2PoolManager *L2PoolManagerCaller) PerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "PerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L2PoolManager *L2PoolManagerSession) PerFee() (*big.Int, error) {
	return _L2PoolManager.Contract.PerFee(&_L2PoolManager.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L2PoolManager *L2PoolManagerCallerSession) PerFee() (*big.Int, error) {
	return _L2PoolManager.Contract.PerFee(&_L2PoolManager.CallOpts)
}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_L2PoolManager *L2PoolManagerCaller) ReLayer(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "ReLayer")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_L2PoolManager *L2PoolManagerSession) ReLayer() ([32]byte, error) {
	return _L2PoolManager.Contract.ReLayer(&_L2PoolManager.CallOpts)
}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_L2PoolManager *L2PoolManagerCallerSession) ReLayer() ([32]byte, error) {
	return _L2PoolManager.Contract.ReLayer(&_L2PoolManager.CallOpts)
}

// SupportTokens is a free data retrieval call binding the contract method 0x626417b5.
//
// Solidity: function SupportTokens(uint256 ) view returns(address)
func (_L2PoolManager *L2PoolManagerCaller) SupportTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "SupportTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportTokens is a free data retrieval call binding the contract method 0x626417b5.
//
// Solidity: function SupportTokens(uint256 ) view returns(address)
func (_L2PoolManager *L2PoolManagerSession) SupportTokens(arg0 *big.Int) (common.Address, error) {
	return _L2PoolManager.Contract.SupportTokens(&_L2PoolManager.CallOpts, arg0)
}

// SupportTokens is a free data retrieval call binding the contract method 0x626417b5.
//
// Solidity: function SupportTokens(uint256 ) view returns(address)
func (_L2PoolManager *L2PoolManagerCallerSession) SupportTokens(arg0 *big.Int) (common.Address, error) {
	return _L2PoolManager.Contract.SupportTokens(&_L2PoolManager.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L2PoolManager *L2PoolManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L2PoolManager *L2PoolManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _L2PoolManager.Contract.GetRoleAdmin(&_L2PoolManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L2PoolManager *L2PoolManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _L2PoolManager.Contract.GetRoleAdmin(&_L2PoolManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L2PoolManager *L2PoolManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L2PoolManager *L2PoolManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _L2PoolManager.Contract.HasRole(&_L2PoolManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L2PoolManager *L2PoolManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _L2PoolManager.Contract.HasRole(&_L2PoolManager.CallOpts, role, account)
}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_L2PoolManager *L2PoolManagerCaller) MessageManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "messageManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_L2PoolManager *L2PoolManagerSession) MessageManager() (common.Address, error) {
	return _L2PoolManager.Contract.MessageManager(&_L2PoolManager.CallOpts)
}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_L2PoolManager *L2PoolManagerCallerSession) MessageManager() (common.Address, error) {
	return _L2PoolManager.Contract.MessageManager(&_L2PoolManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2PoolManager *L2PoolManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2PoolManager *L2PoolManagerSession) Paused() (bool, error) {
	return _L2PoolManager.Contract.Paused(&_L2PoolManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2PoolManager *L2PoolManagerCallerSession) Paused() (bool, error) {
	return _L2PoolManager.Contract.Paused(&_L2PoolManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2PoolManager *L2PoolManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2PoolManager *L2PoolManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2PoolManager.Contract.SupportsInterface(&_L2PoolManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L2PoolManager *L2PoolManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L2PoolManager.Contract.SupportsInterface(&_L2PoolManager.CallOpts, interfaceId)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x92d84738.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeFinalizeERC20(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeFinalizeERC20", sourceChainId, destChainId, to, ERC20Address, amount, _fee, _nonce)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x92d84738.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeFinalizeERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeFinalizeERC20(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, amount, _fee, _nonce)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x92d84738.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeFinalizeERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeFinalizeERC20(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x38731c47.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeFinalizeETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeFinalizeETH", sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x38731c47.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeFinalizeETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeFinalizeETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x38731c47.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeFinalizeETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeFinalizeETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeWETH is a paid mutator transaction binding the contract method 0x5765a4eb.
//
// Solidity: function BridgeFinalizeWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeFinalizeWETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeFinalizeWETH", sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeWETH is a paid mutator transaction binding the contract method 0x5765a4eb.
//
// Solidity: function BridgeFinalizeWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeFinalizeWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeFinalizeWETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeWETH is a paid mutator transaction binding the contract method 0x5765a4eb.
//
// Solidity: function BridgeFinalizeWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeFinalizeWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeFinalizeWETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x3b0230f0.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeInitiateERC20(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeInitiateERC20", sourceChainId, destChainId, to, ERC20Address, value)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x3b0230f0.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeInitiateERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateERC20(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, value)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x3b0230f0.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeInitiateERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateERC20(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, value)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x9b442380.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address to) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeInitiateETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeInitiateETH", sourceChainId, destChainId, to)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x9b442380.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address to) payable returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeInitiateETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x9b442380.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address to) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeInitiateETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to)
}

// BridgeInitiateWETH is a paid mutator transaction binding the contract method 0xcd01c665.
//
// Solidity: function BridgeInitiateWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 value) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeInitiateWETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeInitiateWETH", sourceChainId, destChainId, to, value)
}

// BridgeInitiateWETH is a paid mutator transaction binding the contract method 0xcd01c665.
//
// Solidity: function BridgeInitiateWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 value) returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeInitiateWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateWETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, value)
}

// BridgeInitiateWETH is a paid mutator transaction binding the contract method 0xcd01c665.
//
// Solidity: function BridgeInitiateWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 value) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeInitiateWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateWETH(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, value)
}

// QuickSendAssertToUser is a paid mutator transaction binding the contract method 0x325e5d43.
//
// Solidity: function QuickSendAssertToUser(address _token, address to, uint256 _amount) returns()
func (_L2PoolManager *L2PoolManagerTransactor) QuickSendAssertToUser(opts *bind.TransactOpts, _token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "QuickSendAssertToUser", _token, to, _amount)
}

// QuickSendAssertToUser is a paid mutator transaction binding the contract method 0x325e5d43.
//
// Solidity: function QuickSendAssertToUser(address _token, address to, uint256 _amount) returns()
func (_L2PoolManager *L2PoolManagerSession) QuickSendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.QuickSendAssertToUser(&_L2PoolManager.TransactOpts, _token, to, _amount)
}

// QuickSendAssertToUser is a paid mutator transaction binding the contract method 0x325e5d43.
//
// Solidity: function QuickSendAssertToUser(address _token, address to, uint256 _amount) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) QuickSendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.QuickSendAssertToUser(&_L2PoolManager.TransactOpts, _token, to, _amount)
}

// UpdateFundingPoolBalance is a paid mutator transaction binding the contract method 0x0cd727df.
//
// Solidity: function UpdateFundingPoolBalance(address token, uint256 amount) returns()
func (_L2PoolManager *L2PoolManagerTransactor) UpdateFundingPoolBalance(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "UpdateFundingPoolBalance", token, amount)
}

// UpdateFundingPoolBalance is a paid mutator transaction binding the contract method 0x0cd727df.
//
// Solidity: function UpdateFundingPoolBalance(address token, uint256 amount) returns()
func (_L2PoolManager *L2PoolManagerSession) UpdateFundingPoolBalance(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.UpdateFundingPoolBalance(&_L2PoolManager.TransactOpts, token, amount)
}

// UpdateFundingPoolBalance is a paid mutator transaction binding the contract method 0x0cd727df.
//
// Solidity: function UpdateFundingPoolBalance(address token, uint256 amount) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) UpdateFundingPoolBalance(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.UpdateFundingPoolBalance(&_L2PoolManager.TransactOpts, token, amount)
}

// WithdrawERC20ToL1 is a paid mutator transaction binding the contract method 0x0227047a.
//
// Solidity: function WithdrawERC20ToL1(address _token, address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) WithdrawERC20ToL1(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "WithdrawERC20ToL1", _token, _to, _amount)
}

// WithdrawERC20ToL1 is a paid mutator transaction binding the contract method 0x0227047a.
//
// Solidity: function WithdrawERC20ToL1(address _token, address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerSession) WithdrawERC20ToL1(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.WithdrawERC20ToL1(&_L2PoolManager.TransactOpts, _token, _to, _amount)
}

// WithdrawERC20ToL1 is a paid mutator transaction binding the contract method 0x0227047a.
//
// Solidity: function WithdrawERC20ToL1(address _token, address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) WithdrawERC20ToL1(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.WithdrawERC20ToL1(&_L2PoolManager.TransactOpts, _token, _to, _amount)
}

// WithdrawETHtoL1 is a paid mutator transaction binding the contract method 0x39891dd7.
//
// Solidity: function WithdrawETHtoL1(address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) WithdrawETHtoL1(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "WithdrawETHtoL1", _to, _amount)
}

// WithdrawETHtoL1 is a paid mutator transaction binding the contract method 0x39891dd7.
//
// Solidity: function WithdrawETHtoL1(address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerSession) WithdrawETHtoL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.WithdrawETHtoL1(&_L2PoolManager.TransactOpts, _to, _amount)
}

// WithdrawETHtoL1 is a paid mutator transaction binding the contract method 0x39891dd7.
//
// Solidity: function WithdrawETHtoL1(address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) WithdrawETHtoL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.WithdrawETHtoL1(&_L2PoolManager.TransactOpts, _to, _amount)
}

// WithdrawWETHToL1 is a paid mutator transaction binding the contract method 0xd28597f4.
//
// Solidity: function WithdrawWETHToL1(address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) WithdrawWETHToL1(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "WithdrawWETHToL1", _to, _amount)
}

// WithdrawWETHToL1 is a paid mutator transaction binding the contract method 0xd28597f4.
//
// Solidity: function WithdrawWETHToL1(address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerSession) WithdrawWETHToL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.WithdrawWETHToL1(&_L2PoolManager.TransactOpts, _to, _amount)
}

// WithdrawWETHToL1 is a paid mutator transaction binding the contract method 0xd28597f4.
//
// Solidity: function WithdrawWETHToL1(address _to, uint256 _amount) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) WithdrawWETHToL1(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.WithdrawWETHToL1(&_L2PoolManager.TransactOpts, _to, _amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L2PoolManager *L2PoolManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L2PoolManager *L2PoolManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.GrantRole(&_L2PoolManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.GrantRole(&_L2PoolManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _MultisigWallet, address _messageManager) returns()
func (_L2PoolManager *L2PoolManagerTransactor) Initialize(opts *bind.TransactOpts, _MultisigWallet common.Address, _messageManager common.Address) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "initialize", _MultisigWallet, _messageManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _MultisigWallet, address _messageManager) returns()
func (_L2PoolManager *L2PoolManagerSession) Initialize(_MultisigWallet common.Address, _messageManager common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.Initialize(&_L2PoolManager.TransactOpts, _MultisigWallet, _messageManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _MultisigWallet, address _messageManager) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) Initialize(_MultisigWallet common.Address, _messageManager common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.Initialize(&_L2PoolManager.TransactOpts, _MultisigWallet, _messageManager)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L2PoolManager *L2PoolManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L2PoolManager *L2PoolManagerSession) Pause() (*types.Transaction, error) {
	return _L2PoolManager.Contract.Pause(&_L2PoolManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _L2PoolManager.Contract.Pause(&_L2PoolManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L2PoolManager *L2PoolManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L2PoolManager *L2PoolManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.RenounceRole(&_L2PoolManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.RenounceRole(&_L2PoolManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L2PoolManager *L2PoolManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L2PoolManager *L2PoolManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.RevokeRole(&_L2PoolManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L2PoolManager.Contract.RevokeRole(&_L2PoolManager.TransactOpts, role, account)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_L2PoolManager *L2PoolManagerTransactor) SetMinTransferAmount(opts *bind.TransactOpts, _MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "setMinTransferAmount", _MinTransferAmount)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_L2PoolManager *L2PoolManagerSession) SetMinTransferAmount(_MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetMinTransferAmount(&_L2PoolManager.TransactOpts, _MinTransferAmount)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) SetMinTransferAmount(_MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetMinTransferAmount(&_L2PoolManager.TransactOpts, _MinTransferAmount)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_L2PoolManager *L2PoolManagerTransactor) SetPerFee(opts *bind.TransactOpts, _PerFee *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "setPerFee", _PerFee)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_L2PoolManager *L2PoolManagerSession) SetPerFee(_PerFee *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetPerFee(&_L2PoolManager.TransactOpts, _PerFee)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) SetPerFee(_PerFee *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetPerFee(&_L2PoolManager.TransactOpts, _PerFee)
}

// SetSupportERC20Token is a paid mutator transaction binding the contract method 0xab0f9e19.
//
// Solidity: function setSupportERC20Token(address ERC20Address, bool isValid) returns()
func (_L2PoolManager *L2PoolManagerTransactor) SetSupportERC20Token(opts *bind.TransactOpts, ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "setSupportERC20Token", ERC20Address, isValid)
}

// SetSupportERC20Token is a paid mutator transaction binding the contract method 0xab0f9e19.
//
// Solidity: function setSupportERC20Token(address ERC20Address, bool isValid) returns()
func (_L2PoolManager *L2PoolManagerSession) SetSupportERC20Token(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetSupportERC20Token(&_L2PoolManager.TransactOpts, ERC20Address, isValid)
}

// SetSupportERC20Token is a paid mutator transaction binding the contract method 0xab0f9e19.
//
// Solidity: function setSupportERC20Token(address ERC20Address, bool isValid) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) SetSupportERC20Token(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetSupportERC20Token(&_L2PoolManager.TransactOpts, ERC20Address, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_L2PoolManager *L2PoolManagerTransactor) SetValidChainId(opts *bind.TransactOpts, chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "setValidChainId", chainId, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_L2PoolManager *L2PoolManagerSession) SetValidChainId(chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetValidChainId(&_L2PoolManager.TransactOpts, chainId, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) SetValidChainId(chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _L2PoolManager.Contract.SetValidChainId(&_L2PoolManager.TransactOpts, chainId, isValid)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L2PoolManager *L2PoolManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L2PoolManager *L2PoolManagerSession) Unpause() (*types.Transaction, error) {
	return _L2PoolManager.Contract.Unpause(&_L2PoolManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L2PoolManager *L2PoolManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _L2PoolManager.Contract.Unpause(&_L2PoolManager.TransactOpts)
}

// L2PoolManagerFinalizeERC20Iterator is returned from FilterFinalizeERC20 and is used to iterate over the raw logs and unpacked data for FinalizeERC20 events raised by the L2PoolManager contract.
type L2PoolManagerFinalizeERC20Iterator struct {
	Event *L2PoolManagerFinalizeERC20 // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerFinalizeERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerFinalizeERC20)
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
		it.Event = new(L2PoolManagerFinalizeERC20)
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
func (it *L2PoolManagerFinalizeERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerFinalizeERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerFinalizeERC20 represents a FinalizeERC20 event raised by the L2PoolManager contract.
type L2PoolManagerFinalizeERC20 struct {
	SourceChainId *big.Int
	DestChainId   *big.Int
	ERC20Address  common.Address
	From          common.Address
	To            common.Address
	Value         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFinalizeERC20 is a free log retrieval operation binding the contract event 0x0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe.
//
// Solidity: event FinalizeERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) FilterFinalizeERC20(opts *bind.FilterOpts, ERC20Address []common.Address, from []common.Address, to []common.Address) (*L2PoolManagerFinalizeERC20Iterator, error) {

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

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "FinalizeERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerFinalizeERC20Iterator{contract: _L2PoolManager.contract, event: "FinalizeERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeERC20 is a free log subscription operation binding the contract event 0x0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe.
//
// Solidity: event FinalizeERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchFinalizeERC20(opts *bind.WatchOpts, sink chan<- *L2PoolManagerFinalizeERC20, ERC20Address []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "FinalizeERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerFinalizeERC20)
				if err := _L2PoolManager.contract.UnpackLog(event, "FinalizeERC20", log); err != nil {
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

// ParseFinalizeERC20 is a log parse operation binding the contract event 0x0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe.
//
// Solidity: event FinalizeERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) ParseFinalizeERC20(log types.Log) (*L2PoolManagerFinalizeERC20, error) {
	event := new(L2PoolManagerFinalizeERC20)
	if err := _L2PoolManager.contract.UnpackLog(event, "FinalizeERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerFinalizeETHIterator is returned from FilterFinalizeETH and is used to iterate over the raw logs and unpacked data for FinalizeETH events raised by the L2PoolManager contract.
type L2PoolManagerFinalizeETHIterator struct {
	Event *L2PoolManagerFinalizeETH // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerFinalizeETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerFinalizeETH)
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
		it.Event = new(L2PoolManagerFinalizeETH)
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
func (it *L2PoolManagerFinalizeETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerFinalizeETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerFinalizeETH represents a FinalizeETH event raised by the L2PoolManager contract.
type L2PoolManagerFinalizeETH struct {
	SourceChainId *big.Int
	DestChainId   *big.Int
	From          common.Address
	To            common.Address
	Value         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFinalizeETH is a free log retrieval operation binding the contract event 0x61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f32.
//
// Solidity: event FinalizeETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) FilterFinalizeETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2PoolManagerFinalizeETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "FinalizeETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerFinalizeETHIterator{contract: _L2PoolManager.contract, event: "FinalizeETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeETH is a free log subscription operation binding the contract event 0x61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f32.
//
// Solidity: event FinalizeETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchFinalizeETH(opts *bind.WatchOpts, sink chan<- *L2PoolManagerFinalizeETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "FinalizeETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerFinalizeETH)
				if err := _L2PoolManager.contract.UnpackLog(event, "FinalizeETH", log); err != nil {
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

// ParseFinalizeETH is a log parse operation binding the contract event 0x61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f32.
//
// Solidity: event FinalizeETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) ParseFinalizeETH(log types.Log) (*L2PoolManagerFinalizeETH, error) {
	event := new(L2PoolManagerFinalizeETH)
	if err := _L2PoolManager.contract.UnpackLog(event, "FinalizeETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerFinalizeWETHIterator is returned from FilterFinalizeWETH and is used to iterate over the raw logs and unpacked data for FinalizeWETH events raised by the L2PoolManager contract.
type L2PoolManagerFinalizeWETHIterator struct {
	Event *L2PoolManagerFinalizeWETH // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerFinalizeWETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerFinalizeWETH)
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
		it.Event = new(L2PoolManagerFinalizeWETH)
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
func (it *L2PoolManagerFinalizeWETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerFinalizeWETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerFinalizeWETH represents a FinalizeWETH event raised by the L2PoolManager contract.
type L2PoolManagerFinalizeWETH struct {
	SourceChainId *big.Int
	DestChainId   *big.Int
	From          common.Address
	To            common.Address
	Value         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFinalizeWETH is a free log retrieval operation binding the contract event 0x7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc14.
//
// Solidity: event FinalizeWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) FilterFinalizeWETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2PoolManagerFinalizeWETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "FinalizeWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerFinalizeWETHIterator{contract: _L2PoolManager.contract, event: "FinalizeWETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeWETH is a free log subscription operation binding the contract event 0x7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc14.
//
// Solidity: event FinalizeWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchFinalizeWETH(opts *bind.WatchOpts, sink chan<- *L2PoolManagerFinalizeWETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "FinalizeWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerFinalizeWETH)
				if err := _L2PoolManager.contract.UnpackLog(event, "FinalizeWETH", log); err != nil {
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

// ParseFinalizeWETH is a log parse operation binding the contract event 0x7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc14.
//
// Solidity: event FinalizeWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) ParseFinalizeWETH(log types.Log) (*L2PoolManagerFinalizeWETH, error) {
	event := new(L2PoolManagerFinalizeWETH)
	if err := _L2PoolManager.contract.UnpackLog(event, "FinalizeWETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2PoolManager contract.
type L2PoolManagerInitializedIterator struct {
	Event *L2PoolManagerInitialized // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerInitialized)
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
		it.Event = new(L2PoolManagerInitialized)
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
func (it *L2PoolManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerInitialized represents a Initialized event raised by the L2PoolManager contract.
type L2PoolManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L2PoolManager *L2PoolManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2PoolManagerInitializedIterator, error) {

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerInitializedIterator{contract: _L2PoolManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L2PoolManager *L2PoolManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2PoolManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerInitialized)
				if err := _L2PoolManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseInitialized(log types.Log) (*L2PoolManagerInitialized, error) {
	event := new(L2PoolManagerInitialized)
	if err := _L2PoolManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerInitiateERC20Iterator is returned from FilterInitiateERC20 and is used to iterate over the raw logs and unpacked data for InitiateERC20 events raised by the L2PoolManager contract.
type L2PoolManagerInitiateERC20Iterator struct {
	Event *L2PoolManagerInitiateERC20 // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerInitiateERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerInitiateERC20)
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
		it.Event = new(L2PoolManagerInitiateERC20)
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
func (it *L2PoolManagerInitiateERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerInitiateERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerInitiateERC20 represents a InitiateERC20 event raised by the L2PoolManager contract.
type L2PoolManagerInitiateERC20 struct {
	SourceChainId *big.Int
	DestChainId   *big.Int
	ERC20Address  common.Address
	From          common.Address
	To            common.Address
	Value         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiateERC20 is a free log retrieval operation binding the contract event 0xece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a4.
//
// Solidity: event InitiateERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) FilterInitiateERC20(opts *bind.FilterOpts, ERC20Address []common.Address, from []common.Address, to []common.Address) (*L2PoolManagerInitiateERC20Iterator, error) {

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

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "InitiateERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerInitiateERC20Iterator{contract: _L2PoolManager.contract, event: "InitiateERC20", logs: logs, sub: sub}, nil
}

// WatchInitiateERC20 is a free log subscription operation binding the contract event 0xece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a4.
//
// Solidity: event InitiateERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchInitiateERC20(opts *bind.WatchOpts, sink chan<- *L2PoolManagerInitiateERC20, ERC20Address []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "InitiateERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerInitiateERC20)
				if err := _L2PoolManager.contract.UnpackLog(event, "InitiateERC20", log); err != nil {
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

// ParseInitiateERC20 is a log parse operation binding the contract event 0xece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a4.
//
// Solidity: event InitiateERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) ParseInitiateERC20(log types.Log) (*L2PoolManagerInitiateERC20, error) {
	event := new(L2PoolManagerInitiateERC20)
	if err := _L2PoolManager.contract.UnpackLog(event, "InitiateERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerInitiateETHIterator is returned from FilterInitiateETH and is used to iterate over the raw logs and unpacked data for InitiateETH events raised by the L2PoolManager contract.
type L2PoolManagerInitiateETHIterator struct {
	Event *L2PoolManagerInitiateETH // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerInitiateETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerInitiateETH)
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
		it.Event = new(L2PoolManagerInitiateETH)
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
func (it *L2PoolManagerInitiateETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerInitiateETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerInitiateETH represents a InitiateETH event raised by the L2PoolManager contract.
type L2PoolManagerInitiateETH struct {
	SourceChainId *big.Int
	DestChainId   *big.Int
	From          common.Address
	To            common.Address
	Value         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiateETH is a free log retrieval operation binding the contract event 0x0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc5.
//
// Solidity: event InitiateETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) FilterInitiateETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2PoolManagerInitiateETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "InitiateETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerInitiateETHIterator{contract: _L2PoolManager.contract, event: "InitiateETH", logs: logs, sub: sub}, nil
}

// WatchInitiateETH is a free log subscription operation binding the contract event 0x0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc5.
//
// Solidity: event InitiateETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchInitiateETH(opts *bind.WatchOpts, sink chan<- *L2PoolManagerInitiateETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "InitiateETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerInitiateETH)
				if err := _L2PoolManager.contract.UnpackLog(event, "InitiateETH", log); err != nil {
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

// ParseInitiateETH is a log parse operation binding the contract event 0x0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc5.
//
// Solidity: event InitiateETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) ParseInitiateETH(log types.Log) (*L2PoolManagerInitiateETH, error) {
	event := new(L2PoolManagerInitiateETH)
	if err := _L2PoolManager.contract.UnpackLog(event, "InitiateETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerInitiateWETHIterator is returned from FilterInitiateWETH and is used to iterate over the raw logs and unpacked data for InitiateWETH events raised by the L2PoolManager contract.
type L2PoolManagerInitiateWETHIterator struct {
	Event *L2PoolManagerInitiateWETH // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerInitiateWETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerInitiateWETH)
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
		it.Event = new(L2PoolManagerInitiateWETH)
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
func (it *L2PoolManagerInitiateWETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerInitiateWETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerInitiateWETH represents a InitiateWETH event raised by the L2PoolManager contract.
type L2PoolManagerInitiateWETH struct {
	SourceChainId *big.Int
	DestChainId   *big.Int
	From          common.Address
	To            common.Address
	Value         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInitiateWETH is a free log retrieval operation binding the contract event 0xcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a89.
//
// Solidity: event InitiateWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) FilterInitiateWETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2PoolManagerInitiateWETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "InitiateWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerInitiateWETHIterator{contract: _L2PoolManager.contract, event: "InitiateWETH", logs: logs, sub: sub}, nil
}

// WatchInitiateWETH is a free log subscription operation binding the contract event 0xcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a89.
//
// Solidity: event InitiateWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchInitiateWETH(opts *bind.WatchOpts, sink chan<- *L2PoolManagerInitiateWETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "InitiateWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerInitiateWETH)
				if err := _L2PoolManager.contract.UnpackLog(event, "InitiateWETH", log); err != nil {
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

// ParseInitiateWETH is a log parse operation binding the contract event 0xcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a89.
//
// Solidity: event InitiateWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) ParseInitiateWETH(log types.Log) (*L2PoolManagerInitiateWETH, error) {
	event := new(L2PoolManagerInitiateWETH)
	if err := _L2PoolManager.contract.UnpackLog(event, "InitiateWETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the L2PoolManager contract.
type L2PoolManagerPausedIterator struct {
	Event *L2PoolManagerPaused // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerPaused)
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
		it.Event = new(L2PoolManagerPaused)
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
func (it *L2PoolManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerPaused represents a Paused event raised by the L2PoolManager contract.
type L2PoolManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2PoolManager *L2PoolManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*L2PoolManagerPausedIterator, error) {

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerPausedIterator{contract: _L2PoolManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2PoolManager *L2PoolManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *L2PoolManagerPaused) (event.Subscription, error) {

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerPaused)
				if err := _L2PoolManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParsePaused(log types.Log) (*L2PoolManagerPaused, error) {
	event := new(L2PoolManagerPaused)
	if err := _L2PoolManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the L2PoolManager contract.
type L2PoolManagerRoleAdminChangedIterator struct {
	Event *L2PoolManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerRoleAdminChanged)
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
		it.Event = new(L2PoolManagerRoleAdminChanged)
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
func (it *L2PoolManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerRoleAdminChanged represents a RoleAdminChanged event raised by the L2PoolManager contract.
type L2PoolManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_L2PoolManager *L2PoolManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*L2PoolManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerRoleAdminChangedIterator{contract: _L2PoolManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_L2PoolManager *L2PoolManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *L2PoolManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerRoleAdminChanged)
				if err := _L2PoolManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseRoleAdminChanged(log types.Log) (*L2PoolManagerRoleAdminChanged, error) {
	event := new(L2PoolManagerRoleAdminChanged)
	if err := _L2PoolManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the L2PoolManager contract.
type L2PoolManagerRoleGrantedIterator struct {
	Event *L2PoolManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerRoleGranted)
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
		it.Event = new(L2PoolManagerRoleGranted)
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
func (it *L2PoolManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerRoleGranted represents a RoleGranted event raised by the L2PoolManager contract.
type L2PoolManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2PoolManager *L2PoolManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L2PoolManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerRoleGrantedIterator{contract: _L2PoolManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2PoolManager *L2PoolManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *L2PoolManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerRoleGranted)
				if err := _L2PoolManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseRoleGranted(log types.Log) (*L2PoolManagerRoleGranted, error) {
	event := new(L2PoolManagerRoleGranted)
	if err := _L2PoolManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the L2PoolManager contract.
type L2PoolManagerRoleRevokedIterator struct {
	Event *L2PoolManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerRoleRevoked)
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
		it.Event = new(L2PoolManagerRoleRevoked)
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
func (it *L2PoolManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerRoleRevoked represents a RoleRevoked event raised by the L2PoolManager contract.
type L2PoolManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2PoolManager *L2PoolManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L2PoolManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerRoleRevokedIterator{contract: _L2PoolManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L2PoolManager *L2PoolManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *L2PoolManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerRoleRevoked)
				if err := _L2PoolManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseRoleRevoked(log types.Log) (*L2PoolManagerRoleRevoked, error) {
	event := new(L2PoolManagerRoleRevoked)
	if err := _L2PoolManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the L2PoolManager contract.
type L2PoolManagerUnpausedIterator struct {
	Event *L2PoolManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerUnpaused)
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
		it.Event = new(L2PoolManagerUnpaused)
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
func (it *L2PoolManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerUnpaused represents a Unpaused event raised by the L2PoolManager contract.
type L2PoolManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2PoolManager *L2PoolManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*L2PoolManagerUnpausedIterator, error) {

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerUnpausedIterator{contract: _L2PoolManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2PoolManager *L2PoolManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *L2PoolManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerUnpaused)
				if err := _L2PoolManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseUnpaused(log types.Log) (*L2PoolManagerUnpaused, error) {
	event := new(L2PoolManagerUnpaused)
	if err := _L2PoolManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerWithdrawERC20toL1SuccessIterator is returned from FilterWithdrawERC20toL1Success and is used to iterate over the raw logs and unpacked data for WithdrawERC20toL1Success events raised by the L2PoolManager contract.
type L2PoolManagerWithdrawERC20toL1SuccessIterator struct {
	Event *L2PoolManagerWithdrawERC20toL1Success // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerWithdrawERC20toL1SuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerWithdrawERC20toL1Success)
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
		it.Event = new(L2PoolManagerWithdrawERC20toL1Success)
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
func (it *L2PoolManagerWithdrawERC20toL1SuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerWithdrawERC20toL1SuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerWithdrawERC20toL1Success represents a WithdrawERC20toL1Success event raised by the L2PoolManager contract.
type L2PoolManagerWithdrawERC20toL1Success struct {
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
func (_L2PoolManager *L2PoolManagerFilterer) FilterWithdrawERC20toL1Success(opts *bind.FilterOpts) (*L2PoolManagerWithdrawERC20toL1SuccessIterator, error) {

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "WithdrawERC20toL1Success")
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerWithdrawERC20toL1SuccessIterator{contract: _L2PoolManager.contract, event: "WithdrawERC20toL1Success", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20toL1Success is a free log subscription operation binding the contract event 0x1bc1b1207e0c319b7baad0317c09f30575218af6f5a59f1f0f15de8d217672cc.
//
// Solidity: event WithdrawERC20toL1Success(uint256 chainId, uint256 timestamp, address token, address to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchWithdrawERC20toL1Success(opts *bind.WatchOpts, sink chan<- *L2PoolManagerWithdrawERC20toL1Success) (event.Subscription, error) {

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "WithdrawERC20toL1Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerWithdrawERC20toL1Success)
				if err := _L2PoolManager.contract.UnpackLog(event, "WithdrawERC20toL1Success", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseWithdrawERC20toL1Success(log types.Log) (*L2PoolManagerWithdrawERC20toL1Success, error) {
	event := new(L2PoolManagerWithdrawERC20toL1Success)
	if err := _L2PoolManager.contract.UnpackLog(event, "WithdrawERC20toL1Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerWithdrawETHtoL1SuccessIterator is returned from FilterWithdrawETHtoL1Success and is used to iterate over the raw logs and unpacked data for WithdrawETHtoL1Success events raised by the L2PoolManager contract.
type L2PoolManagerWithdrawETHtoL1SuccessIterator struct {
	Event *L2PoolManagerWithdrawETHtoL1Success // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerWithdrawETHtoL1SuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerWithdrawETHtoL1Success)
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
		it.Event = new(L2PoolManagerWithdrawETHtoL1Success)
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
func (it *L2PoolManagerWithdrawETHtoL1SuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerWithdrawETHtoL1SuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerWithdrawETHtoL1Success represents a WithdrawETHtoL1Success event raised by the L2PoolManager contract.
type L2PoolManagerWithdrawETHtoL1Success struct {
	ChainId   *big.Int
	Timestamp *big.Int
	To        common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETHtoL1Success is a free log retrieval operation binding the contract event 0xf1fa6dbbb3e74f1bef249f22545b238726c1cc79c5571a1416575674c7e63f39.
//
// Solidity: event WithdrawETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) FilterWithdrawETHtoL1Success(opts *bind.FilterOpts) (*L2PoolManagerWithdrawETHtoL1SuccessIterator, error) {

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "WithdrawETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerWithdrawETHtoL1SuccessIterator{contract: _L2PoolManager.contract, event: "WithdrawETHtoL1Success", logs: logs, sub: sub}, nil
}

// WatchWithdrawETHtoL1Success is a free log subscription operation binding the contract event 0xf1fa6dbbb3e74f1bef249f22545b238726c1cc79c5571a1416575674c7e63f39.
//
// Solidity: event WithdrawETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchWithdrawETHtoL1Success(opts *bind.WatchOpts, sink chan<- *L2PoolManagerWithdrawETHtoL1Success) (event.Subscription, error) {

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "WithdrawETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerWithdrawETHtoL1Success)
				if err := _L2PoolManager.contract.UnpackLog(event, "WithdrawETHtoL1Success", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseWithdrawETHtoL1Success(log types.Log) (*L2PoolManagerWithdrawETHtoL1Success, error) {
	event := new(L2PoolManagerWithdrawETHtoL1Success)
	if err := _L2PoolManager.contract.UnpackLog(event, "WithdrawETHtoL1Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PoolManagerWithdrawWETHtoL1SuccessIterator is returned from FilterWithdrawWETHtoL1Success and is used to iterate over the raw logs and unpacked data for WithdrawWETHtoL1Success events raised by the L2PoolManager contract.
type L2PoolManagerWithdrawWETHtoL1SuccessIterator struct {
	Event *L2PoolManagerWithdrawWETHtoL1Success // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerWithdrawWETHtoL1SuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerWithdrawWETHtoL1Success)
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
		it.Event = new(L2PoolManagerWithdrawWETHtoL1Success)
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
func (it *L2PoolManagerWithdrawWETHtoL1SuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerWithdrawWETHtoL1SuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerWithdrawWETHtoL1Success represents a WithdrawWETHtoL1Success event raised by the L2PoolManager contract.
type L2PoolManagerWithdrawWETHtoL1Success struct {
	ChainId   *big.Int
	Timestamp *big.Int
	To        common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawWETHtoL1Success is a free log retrieval operation binding the contract event 0xcb1a101498aadad4159168d719f955e559e768912e6f5f5c57b094eb047d000f.
//
// Solidity: event WithdrawWETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) FilterWithdrawWETHtoL1Success(opts *bind.FilterOpts) (*L2PoolManagerWithdrawWETHtoL1SuccessIterator, error) {

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "WithdrawWETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerWithdrawWETHtoL1SuccessIterator{contract: _L2PoolManager.contract, event: "WithdrawWETHtoL1Success", logs: logs, sub: sub}, nil
}

// WatchWithdrawWETHtoL1Success is a free log subscription operation binding the contract event 0xcb1a101498aadad4159168d719f955e559e768912e6f5f5c57b094eb047d000f.
//
// Solidity: event WithdrawWETHtoL1Success(uint256 chainId, uint256 timestamp, address to, uint256 value)
func (_L2PoolManager *L2PoolManagerFilterer) WatchWithdrawWETHtoL1Success(opts *bind.WatchOpts, sink chan<- *L2PoolManagerWithdrawWETHtoL1Success) (event.Subscription, error) {

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "WithdrawWETHtoL1Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerWithdrawWETHtoL1Success)
				if err := _L2PoolManager.contract.UnpackLog(event, "WithdrawWETHtoL1Success", log); err != nil {
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
func (_L2PoolManager *L2PoolManagerFilterer) ParseWithdrawWETHtoL1Success(log types.Log) (*L2PoolManagerWithdrawWETHtoL1Success, error) {
	event := new(L2PoolManagerWithdrawWETHtoL1Success)
	if err := _L2PoolManager.contract.UnpackLog(event, "WithdrawWETHtoL1Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
