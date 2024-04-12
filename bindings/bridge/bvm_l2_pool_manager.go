// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ChainIdIsNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"ChainIdNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorBlockChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"MinTransferAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LessThanMinTransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MantaNotWETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MantleNotWETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughETH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"NotEnoughToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"TokenIsNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"sourceChainIdError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"BridgeInitiateForStakingSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20toL1Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawETHtoL1Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawWETHtoL1Success\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateERC20ForStaking\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"BridgeInitiateETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateETHForStaking\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FeePoolValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FundingPoolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"IsSupportChainId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"IsSupportToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_GAS_Limit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinTransferAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"QuickSendAssertToUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ReLayer\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SupportTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UpdateFundingPoolBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20ToL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawETHtoL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawWETHToL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MultisigWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageManager\",\"outputs\":[{\"internalType\":\"contractIMessageManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_MinTransferAmount\",\"type\":\"uint256\"}],\"name\":\"setMinTransferAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_PerFee\",\"type\":\"uint256\"}],\"name\":\"setPerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setSupportERC20Token\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setValidChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60808060405234620000bd577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c16620000ae57506001600160401b036002600160401b03198282160162000068575b6040516137d09081620000c38239f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a138808062000058565b63f92ee8a960e01b8152600490fd5b600080fdfe608060408181526004918236101561001657600080fd5b600092833560e01c91826301ffc9a714612e7c575081630227047a1461266157816309dc480c146126425781630cd727df1461260357816310875a13146125db578163248a9ca3146125a35781632f2ff15d14612579578163325e5d431461241757816336568abe146123d157816338731c471461225657816339891dd714611bac5781633b0230f0146118f25781633b90b1f9146118305781633f4ba83a146117b8578163485cc9551461161357816354dc027e146115db5781635765a4eb146114105781635c975abb146113df578163626417b51461139e578163650c22761461137c57816372fe6a7e1461134f5781638456cb59146112d557816391d148541461128257816392d84738146110dc5781639b44238014610f20578163a217fddf14610f05578163a525ae2b14610ee0578163ab0f9e1914610e35578163b1887e9b14610e07578163b92e639614610de2578163ba97e76e14610cd1578163cd01c665146109e4578163d28597f414610320578163d547741f146102d3578163d73f14e51461028e578163dac295681461024d57508063dbb0481f14610216578063dd0c3460146101f85763fa861848146101d257600080fd5b346101f457816003193601126101f4576020906101ed61308b565b9051908152f35b5080fd5b50346101f457816003193601126101f4576020906002549051908152f35b50346101f45760203660031901126101f45760209181906001600160a01b0361023d612ecf565b1681526005845220549051908152f35b90503461028a57602036600319011261028a5760209260ff918391906001600160a01b03610279612ecf565b168252855220541690519015158152f35b8280fd5b90503461028a578160031936011261028a576102d0916102ac61300d565b916102b56130de565b358452600360205283209060ff801983541691151516179055565b80f35b9190503461028a578060031936011261028a5761031c913561031760016102f8612ee5565b9383875260008051602061371b83398151915260205286200154613136565b613296565b5080f35b90508160031936011261028a57610335612ecf565b9260249384359261034c61034761308b565b613136565b6001600160a01b0361035c6135f3565b168551966370a0823160e01b8852308389015260209788818381865afa9081156105065786916109b2575b50861161099e578488466208275003610510575050865163095ea7b360e01b815286737003e7b7186f0e6601203b99f7b8decbfa391cf9928386840152820152888160448189875af18015610506576104d9575b5063ffffffff6008541692813b156104d557875163a93a4af960e01b81526001600160a01b0393841691810191825292851660208201526040810187905260608101849052919285928492839185919083906080010393f180156104cb579184939186936104af575b506104a69360008051602061377b833981519152915b84514681524260208201526001600160a01b039190911660408201526060810191909152608090a173c02aaa39b223fe8d0a0e5c4f27ead9083c756cc28152600586522091825461301c565b90555160018152f35b8294506104bd91935061303f565b6101f4578291849138610444565b85513d85823e3d90fd5b8580fd5b6104f890893d8b116104ff575b6104f08183613069565b8101906133e0565b50386103db565b503d6104e6565b88513d88823e3d90fd5b9091939261044d461460001461063357604489518094819363095ea7b360e01b83528b732a3dd3eb832af982ec71669e178424b10dca2ede998a8a8601528401525af180156106295761060c575b50813b1561060857855163cd58657960e01b815260009181018281526001600160a01b038516602082015260408101879052606081018390526080810183905260c060a0820181905281019290925291849183919082908890829060e0015b03925af180156104cb579184939186936105ec575b506104a69360008051602061377b8339815191529161045a565b8294506105fa91935061303f565b6101f45782918491386105d2565b8380fd5b61062290883d8a116104ff576104f08183613069565b503861055e565b87513d87823e3d90fd5b9291505046600a036106c657865163095ea7b360e01b81526010602160991b01828201819052938101879052888160448189875af18015610506576106a9575b5063ffffffff60085416833b156104d5579285928786936105bd82978c5198899788968795631474f2a960e31b875286016133a6565b6106bf90893d8b116104ff576104f08183613069565b5038610673565b8488466121050361073c575050865163095ea7b360e01b81526010602160991b01828201819052938101879052888160448189875af18015610506576106a9575063ffffffff60085416833b156104d5579285928786936105bd82978c5198899788968795631474f2a960e31b875286016133a6565b90919261a4b1461460001461085557604489518094819363095ea7b360e01b83528b736c411ad3e74de3e7bd422b94a27770f5b86c623b998a8a8601528401525af1801561062957928486809489946107e997610838575b508a5196879586948593637b3a3c8b60e01b85528401909160a09273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc28352600180851b03166020830152604082015260806060820152600060808201520190565b03925af180156104cb578486936104a6959360008051602061377b83398151915293610816575b5061045a565b610831903d8087833e6108298183613069565b810190613414565b5038610810565b61084e908d803d106104ff576104f08183613069565b5038610794565b4661a4ba036108fb57604489518094819363095ea7b360e01b83528b737626841cb6113412f9c88d3adc720c9fac88d9ed998a8a8601528401525af1801561062957928486809489946107e99761083857508a5196879586948593637b3a3c8b60e01b85528401909160a09273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc28352600180851b03166020830152604082015260806060820152600060808201520190565b935050504661014403610990577311f943b2c77b743ab90f4a0ae7d5a4e7fca3e10290813b1561098c578651636ce5768960e11b81526001600160a01b038086169282019283529093166020820152604081018690528491839182908890829060600103925af180156104cb579184939186936105ec57506104a69360008051602061377b8339815191529161045a565b8480fd5b8551639474dee160e01b8152fd5b92508551916311745b6160e21b8352820152fd5b90508881813d83116109dd575b6109c98183613069565b810103126109d8575138610387565b600080fd5b503d6109bf565b90503461028a576109f436612f92565b9193468403610cc3578487526020966003885260ff878220541615610cac576001600160a01b039384610a256135f3565b16908851906370a0823160e01b9081835230868401528b83602481875afa928315610c6f57908c918694610c79575b508b516323b872dd60e01b81523388820190815230602082015260408101929092529190829081906060015b038188885af18015610c6f57908c9291610c52575b5060248b5180958193825230898301525afa8015610c48578390610c19575b610abe925061301c565b600154808210610bfe575073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc29283835260058a52888320610af4838254613492565b9055610b11620f4240610b096002548561349f565b04809361301c565b93835260068a52888320610b26838254613492565b90558583541691823b1561060857895163305f478560e21b8152918201888152602081018a90526001600160a01b03871660408201526060810186905260808101919091529091839183919082908490829060a00103925af18015610bf457610bdd575b505090610bd37fcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a899392875193849316963396846040919493926060820195825260208201520152565b0390a35160018152f35b610be7829161303f565b610bf15780610b8a565b80fd5b88513d84823e3d90fd5b88516358f8331360e11b815293840152602483015250604490fd5b508982813d8311610c41575b610c2f8183613069565b810103126109d857610abe9151610ab4565b503d610c25565b89513d85823e3d90fd5b610c6890833d85116104ff576104f08183613069565b5038610a95565b8b513d87823e3d90fd5b8281939295503d8311610ca5575b610c918183613069565b810103126109d85751918b90610a80610a54565b503d610c87565b86516363b5c9db60e01b8152808301879052602490fd5b8551631a26aa4d60e21b8152fd5b90503461028a5760c036600319011261028a578035602435916020610cf4612efb565b610cfc612f11565b86516303b0230f60e41b8152938401858152602081018790526001600160a01b0392831660408201529116606082015260843560808201528290819060a001038188305af18015610dd857610dad7f48917de34e53153e06b5ed7e1c30d2458abfe4ba4ddb223cfce220c13131c5f693602097969593610d9f93610dbb575b508551888101918252602082019490945260a4356040820152929182906060850190565b03601f198101835282613069565b51902091519280a260018152f35b610dd190893d81116104ff576104f08183613069565b5038610d7b565b84513d87823e3d90fd5b8390346101f45760203660031901126101f457610e0061034761308b565b3560015580f35b5050346101f457816003193601126101f457602090610e246135f3565b90516001600160a01b039091168152f35b9190503461028a578060031936011261028a57610e50612ecf565b91610e5961300d565b610e616130de565b6001600160a01b039384168086526020839052928520805460ff191660ff83151516179055610e8e578380f35b6007549068010000000000000000821015610ecd5750806001610eb49201600755612fc0565b909283549160031b92831b921b19161790553880808380f35b634e487b7160e01b855260419052602484fd5b5050346101f457816003193601126101f45760209063ffffffff600854169051908152f35b5050346101f457816003193601126101f45751908152602090f35b82846060366003190112610bf157823560243591610f3c612efb565b904683036110cc57838152600360205260ff8582205416156110b557600154803410611099575073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee908181526005602052858120610f8f348254613492565b9055620f4240610fa16002543461349f565b04610fac813461301c565b9282526006602052868220610fc2828254613492565b905581546001600160a01b039890891691823b1561060857885163305f478560e21b815246928101928352602083018990526001600160a01b038716604084015260608301869052608083019190915291839183919082908490829060a00103925af1801561108f5761107b575b5050947f0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc591610bd3602097875193849316963396846040919493926060820195825260208201520152565b611085829161303f565b610bf15780611030565b87513d84823e3d90fd5b866044918751916358f8331360e11b8352820152346024820152fd5b8451632ef6974160e11b8152808701859052602490fd5b8451631a26aa4d60e21b81528690fd5b90503461028a5760e036600319011261028a578035602435916110fd612efb565b611105612f11565b956084359061111561034761308b565b46860361127257848152600360205260ff87822054161561125b576001600160a01b0397881680825260208590528782205490949060ff16156112445761115d8385876135b0565b848252600560205287822061117384825461301c565b905588825416803b1561028a578851631eb65ffb60e01b8152918201878152602081018990526001600160a01b038616604082015260a43560608201526080810185905260c43560a08201528391839182908490829060c00103925af18015610bf457611230575b5050957f0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe91611226602098885193849316973097846040919493926060820195825260208201520152565b0390a45160018152f35b61123a829161303f565b610bf157806111db565b87516305fd61ad60e01b8152908101859052602490fd5b8651632ef6974160e11b8152808501869052602490fd5b8651631a26aa4d60e21b81528490fd5b90503461028a578160031936011261028a57816020936112a0612ee5565b9235815260008051602061371b8339815191528552209060018060a01b0316600052825260ff81600020541690519015158152f35b90503461028a578260031936011261028a576112ef6130de565b60008051602061375b8339815191529081549060ff8216611341575060ff19166001179055513381527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25890602090a180f35b835163d93c066560e01b8152fd5b90503461028a57602036600319011261028a578160209360ff923581526003855220541690519015158152f35b8390346101f45760203660031901126101f4576113976130de565b3560025580f35b90503461028a57602036600319011261028a573591600754831015610bf157506113c9602092612fc0565b905491519160018060a01b039160031b1c168152f35b5050346101f457816003193601126101f45760209060ff60008051602061375b833981519152541690519015158152f35b828434610bf15761142036612f5c565b93959061142e61034761308b565b4687036115cb5785825260209860038a5260ff8984205416156115b4576001600160a01b03956114958b868c89888c6114656135f3565b16925180968195829463a9059cbb60e01b84528b840160209093929193604081019460018060a01b031681520152565b03925af18015610c6f57611597575b5073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2845260058b528984206114ce86825461301c565b90558684541692833b1561098c578a51631eb65ffb60e01b8152928301898152602081018b90526001600160a01b038816604082015260608101919091526080810186905260a08101919091529091839183919082908490829060c00103925af18015610bf457611583575b505090610bd37f7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc149392875193849316963096846040919493926060820195825260208201520152565b61158d829161303f565b610bf1578061153a565b6115ad908c8d3d106104ff576104f08183613069565b508b6114a4565b8851632ef6974160e11b8152908101879052602490fd5b8751631a26aa4d60e21b81528990fd5b5050346101f45760203660031901126101f45760209181906001600160a01b03611603612ecf565b1681526006845220549051908152f35b90503461028a578160031936011261028a5761162d612ecf565b90611636612ee5565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0092835460ff81871c16159367ffffffffffffffff8216801590816117b0575b60011490816117a6575b15908161179d575b5061178f575067ffffffffffffffff19811660011785556116f9919084611770575b506116b3613319565b6116bb613319565b6116c3613319565b60008051602061375b833981519152805460ff191690556116e2613319565b67016345785d8a0000600155612710600255613169565b5084546001600160a01b0319166001600160a01b03919091161784556008805463ffffffff1916620493e017905561172f578280f35b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a138808280f35b68ffffffffffffffffff191668010000000000000001178555386116aa565b865163f92ee8a960e01b8152fd5b90501538611688565b303b159150611680565b869150611676565b90503461028a578260031936011261028a576117d26130de565b60008051602061375b8339815191529081549060ff821615611822575060ff19169055513381527f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa90602090a180f35b8351638dfc202b60e01b8152fd5b905061183b36612f92565b8551630136884760e71b8152948501849052602485018390526001600160a01b039091166044850152919260208160648189305af180156118e8576020969594927f48917de34e53153e06b5ed7e1c30d2458abfe4ba4ddb223cfce220c13131c5f69492610dad926118cb575b508551888101948552602085019190915260408401929092528160608401610d9f565b6118e190893d81116104ff576104f08183613069565b50386118a8565b85513d88823e3d90fd5b90503461028a5760a036600319011261028a5780356024803592611914612efb565b61191c612f11565b9160843597468603611b9d5786815260209860038a5260ff898320541615611b87576001600160a01b03948516808352838b528983205490969060ff1615611b735789516370a0823160e01b80825230868301528c8284818c5afa918215611b6957908d918693611b38575b508c516323b872dd60e01b8382015233858201908152306020820152604081018790526119c491906119be908290606001610d9f565b8b6134b2565b8c5193849182523088830152818b5afa8015611b2e578490611aff575b6119eb925061301c565b9086835260058b52611a018a8420918254613492565b9055611a1e620f4240611a166002548461349f565b04809261301c565b9286835260068b52898320611a34838254613492565b90558583541691823b15610608578a5163305f478560e21b8152918201898152602081018b90526001600160a01b03871660408201526060810186905260808101919091529091839183919082908490829060a00103925af18015611af557611ae1575b5050906112267fece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a49392885193849316973397846040919493926060820195825260208201520152565b611aeb829161303f565b610bf15780611a98565b89513d84823e3d90fd5b508b82813d8311611b27575b611b158183613069565b810103126109d8576119eb91516119e1565b503d611b0b565b8b513d86823e3d90fd5b8281939294503d8311611b62575b611b508183613069565b810103126109d8578c90519138611988565b503d611b46565b8c513d87823e3d90fd5b89516305fd61ad60e01b8152808501889052fd5b8851632ef6974160e11b81528084018990528690fd5b508651631a26aa4d60e21b8152fd5b90508160031936011261028a57611bc1612ecf565b602491823592611bd261034761308b565b4784116122475785908484466208275003611cd057505063ffffffff6008541692736ea73e05adc79974b931123675ea8f78ffdacdf090813b15610608578660649281968a5197889687956317e614fd60e11b875260018060a01b038c1690870152850152806044850152f18015610dd857611cbc575b50906000805160206136fb83398151915291836020955b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee81526005875220611c8783825461301c565b905583514681524260208201526001600160a01b039091166040820152606081019190915280608081015b0390a15160018152f35b611cc6859161303f565b6106085738611c49565b9092914661044d03611d8f5750509050732a3dd3eb832af982ec71669e178424b10dca2ede91823b156101f457855163cd58657960e01b815260009181018281526001600160a01b038616602082015260408101879052606081018390526080810183905260c060a0820181905281019290925292839182908790829060e0015b03925af18015610dd857611d7b575b50906000805160206136fb8339815191529183602095611c60565b611d85859161303f565b6106085738611d60565b46600a03611dd657505060085463ffffffff16929091506010602160991b01803b1561028a5784868093611d518a5197889687958694631474f2a960e31b8652850161335a565b61a4b194919293944614600014611e9a57508651637b3a3c8b60e01b815273eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee9181019182526001600160a01b0390921660208201526040810193909352608060608401819052600090840152918290819060a001038186735288c571fd7ad117bea99bf60fe0846c4e84f9335af18015610dd8576020956000805160206136fb8339815191529493928692611e80575b50611c60565b611e93903d8084833e6108298183613069565b5038611e7a565b4661a4ba03611f3b57508651637b3a3c8b60e01b815273eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee9181019182526001600160a01b0390921660208201526040810193909352608060608401819052600090840152918290819060a0010381867321903d3f8176b1a0c17e953cd896610be9ffdfa85af18015610dd8576020956000805160206136fb8339815191529493928692611e805750611c60565b915092506101444614600014611fbc577311f943b2c77b743ab90f4a0ae7d5a4e7fca3e102803b1561028a576000938660649289519687958694636ce5768960e11b865260018060a01b038b16908601528401528160448401525af18015610dd857611d7b5750906000805160206136fb8339815191529183602095611c60565b9146611388036120b257855163095ea7b360e01b81526010602160991b0182820181905284820187905273deaddeaddeaddeaddeaddeaddeaddeaddead111192909160208160448188885af180156120a85761208a575b5063ffffffff60085416823b1561098c5760c49285918a51978896879563297e27a960e11b875286015260018060a01b038a1690850152896044850152606484015260a06084840152600060a48401525af18015610dd857611d7b5750906000805160206136fb8339815191529183602095611c60565b6120a19060203d81116104ff576104f08183613069565b5038612013565b89513d87823e3d90fd5b91904660a9036120f95760085463ffffffff16929091506010602160991b01803b1561028a5784868093611d518a5197889687958694631474f2a960e31b8652850161335a565b90914661a70e036121f15785519063095ea7b360e01b8252732c4813276869d93afdab4dd2b01cd670342da1949081818401528685840152734b21b980d0dc7d3c0c6175b0a412694f3a1c7c6b9260208160448188885af180156120a8576121d3575b50813b15610608578360e4926000948a51978896879563cd58657960e01b875286015260018060a01b038a169085015289604485015260648401526000608484015260c060a4840152600060c48401525af18015610dd857611d7b5750906000805160206136fb8339815191529183602095611c60565b6121ea9060203d81116104ff576104f08183613069565b503861215c565b91505046612105036122395760085463ffffffff169085906010602160991b01803b1561028a5784868093611d518a5197889687958694631474f2a960e31b8652850161335a565b8351639474dee160e01b8152fd5b508351632c1d501360e11b8152fd5b91905061226236612f5c565b9792919661227161034761308b565b4685036123c357858452600360205260ff8785205416156123ac57969795966001600160a01b038381169790858080808e8d8282156123a3575bf1156120a85773eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee865260056020528986206122db8c825461301c565b905585541692833b156104d5578951631eb65ffb60e01b8152928301888152602081018890526001600160a01b0390951660408601526060850152608084018a905260a0840152918391839182908490829060c00103925af1801561239957612385575b50508351918252602082810191909152604082019490945230907f61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f32908060608101610bd3565b61238f829161303f565b610bf1578061233f565b86513d84823e3d90fd5b506108fc6122ab565b8651632ef6974160e11b8152908101869052602490fd5b8651631a26aa4d60e21b8152fd5b8383346101f457806003193601126101f4576123eb612ee5565b90336001600160a01b03831603612408575061031c919235613296565b5163334bd91960e11b81528390fd5b90503461028a5761242736612f27565b9261243661034793929361308b565b6001600160a01b0391821680875260208281528688205460ff1615612562578188526005815286882061246a87825461301c565b905573eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee82036124cf5750508347106124c1575084809381938293839183156124b7575b1690f1156124ad575080f35b51903d90823e3d90fd5b6108fc92506124a1565b8451632c1d501360e11b8152fd5b909250859695516370a0823160e01b815230838201528181602481875afa91821561255857908692918892612527575b50501061251257506102d09394506135b0565b6024918651916311745b6160e21b8352820152fd5b8193508092503d8311612551575b61253f8183613069565b810103126109d85784905138806124ff565b503d612535565b88513d89823e3d90fd5b86516305fd61ad60e01b8152808401839052602490fd5b9190503461028a578060031936011261028a5761031c913561259e60016102f8612ee5565b613209565b90503461028a57602036600319011261028a578160209360019235815260008051602061371b83398151915285522001549051908152f35b5050346101f457816003193601126101f457905490516001600160a01b039091168152602090f35b5050346101f457806003193601126101f45761261d612ecf565b61262861034761308b565b6001600160a01b0316825260056020528120602435905580f35b5050346101f457816003193601126101f4576020906001549051908152f35b91905061266d36612f27565b61267b61034793929361308b565b6001600160a01b03821680875260208681528588205490969060ff1615612e67574662082750036127e8578786519163095ea7b360e01b835273d8a791fe2be73eb6e6cf1eb0cb3f36adc9b3f8f9928382820152856024820152898160448186895af18015610c48576127cb575b5063ffffffff60085416833b1561028a57885163a93a4af960e01b81526001600160a01b038089169382019384528916602084015260408301879052606083018290529384928391859183906080010393f180156127c1576127a9575b50908460008051602061373b833981519152959697611cb2935b8152600589522061277282825461301c565b905585514681524260208201526001600160a01b0393841660408201529390921660608401526080830191909152819060a0820190565b6127b3889161303f565b6127bd5738612746565b8680fd5b86513d8a823e3d90fd5b6127e1908a3d8c116104ff576104f08183613069565b50386126e9565b4661044d036128ef578786519163095ea7b360e01b8352732a3dd3eb832af982ec71669e178424b10dca2ede928382820152856024820152898160448186895af18015610c48576128d2575b50823b156101f457875163cd58657960e01b815260009181018281526001600160a01b03808a16602083015260408201889052881660608201526080810183905260c060a0820181905281019290925292839182908490829060e0015b03925af180156127c1576128be575b50908460008051602061373b833981519152959697611cb293612760565b6128c8889161303f565b6127bd57386128a0565b6128e8908a3d8c116104ff576104f08183613069565b5038612834565b46600a0361298057855163095ea7b360e01b81526010602160991b01828201819052602482018590528991898160448186895af18015610c4857612963575b5063ffffffff6008541690803b1561028a57868680956128918b958d5198899788968795631474f2a960e31b875286016133a6565b612979908a3d8c116104ff576104f08183613069565b503861292e565b46612105036129f457855163095ea7b360e01b81526010602160991b01828201819052602482018590528991898160448186895af18015610c4857612963575063ffffffff6008541690803b1561028a57868680956128918b958d5198899788968795631474f2a960e31b875286016133a6565b61a4b1979293949596974614600014612b0c5786519063095ea7b360e01b82527309e9222e96e7b4ae2a407b98d48e330053351eee918282820152856024820152898160448188885af180156120a85792878680948d8c958b98612aee575b5050612a9b8d5197889687958694637b3a3c8b60e01b86528501919060a09391600180861b038092168452166020830152604082015260806060820152600060808201520190565b03925af18015612ae4579260008051602061373b833981519152969594928892611cb295612aca575b50612760565b612add903d8085833e6108298183613069565b5038612ac4565b87513d85823e3d90fd5b81612b0492903d106104ff576104f08183613069565b508d38612a53565b4661a4ba03612ba85786519063095ea7b360e01b825273b2535b988dce19f9d71dfb22db6da744acac21bf918282820152856024820152898160448188885af180156120a85792878680948d8c958b98612aee575050612a9b8d5197889687958694637b3a3c8b60e01b86528501919060a09391600180861b038092168452166020830152604082015260806060820152600060808201520190565b919695949392914661014403612c3957877311f943b2c77b743ab90f4a0ae7d5a4e7fca3e10291823b156101f4578751636ce5768960e11b81526001600160a01b03808916928201928352871660208301526040820186905292839182908790829060600103925af180156127c1576128be5750908460008051602061373b833981519152959697611cb293612760565b4661138803612ccb57855163095ea7b360e01b81526010602160991b01828201819052602482018590528991898160448186895af18015610c4857612cae575b5063ffffffff6008541690803b1561028a57868387956128918b958d519889978896879563297e27a960e11b875286016133a6565b612cc4908a3d8c116104ff576104f08183613069565b5038612c79565b4660a903612d3e57855163095ea7b360e01b81526010602160991b01828201819052602482018590528991898160448186895af18015610c4857612963575063ffffffff6008541690803b1561028a57868680956128918b958d5198899788968795631474f2a960e31b875286016133a6565b61a70e97919293949596974614600014612e595786519063095ea7b360e01b8252732c4813276869d93afdab4dd2b01cd670342da194918282820152856024820152898160448187895af18015612e4f57612e32575b50813b1561028a57875163cd58657960e01b815260009181018281526001600160a01b03808a16602083015260408201889052881660608201526080810183905260c060a0820181905281019290925291839183919082908490829060e00103925af1801561108f579260008051602061373b833981519152969594928892611cb295612e23575b5090612760565b612e2c9061303f565b38612e1c565b612e48908a3d8c116104ff576104f08183613069565b5038612d94565b89513d86823e3d90fd5b8651639474dee160e01b8152fd5b6024918651916305fd61ad60e01b8352820152fd5b84913461028a57602036600319011261028a573563ffffffff60e01b811680910361028a5760209250637965db0b60e01b8114908115612ebe575b5015158152f35b6301ffc9a760e01b14905083612eb7565b600435906001600160a01b03821682036109d857565b602435906001600160a01b03821682036109d857565b604435906001600160a01b03821682036109d857565b606435906001600160a01b03821682036109d857565b60609060031901126109d8576001600160a01b039060043582811681036109d8579160243590811681036109d8579060443590565b60c09060031901126109d85760043590602435906044356001600160a01b03811681036109d85790606435906084359060a43590565b60809060031901126109d85760043590602435906044356001600160a01b03811681036109d8579060643590565b600754811015612ff75760076000527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880190600090565b634e487b7160e01b600052603260045260246000fd5b6024359081151582036109d857565b9190820391821161302957565b634e487b7160e01b600052601160045260246000fd5b67ffffffffffffffff811161305357604052565b634e487b7160e01b600052604160045260246000fd5b90601f8019910116810190811067ffffffffffffffff82111761305357604052565b604051602081017f33fe247d78fee2e7fd135c405eda4bd2911c0a73c0a81b36c3bcc967dd06e5ad8152602082526040820182811067ffffffffffffffff821117613053576040529051902060ff191690565b3360009081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604081205460ff16156131185750565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b8060005260008051602061371b83398151915260205260406000203360005260205260ff60406000205416156131185750565b6001600160a01b031660008181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604081205490919060008051602061371b8339815191529060ff16613204578280526020526040822081835260205260408220600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b505090565b9060009180835260008051602061371b83398151915280602052604084209260018060a01b03169283855260205260ff60408520541615600014613290578184526020526040832082845260205260408320600160ff198254161790557f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d339380a4600190565b50505090565b9060009180835260008051602061371b83398151915280602052604084209260018060a01b03169283855260205260ff60408520541660001461329057818452602052604083208284526020526040832060ff1981541690557ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b339380a4600190565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c161561334857565b604051631afcd79f60e31b8152600490fd5b919260c09363ffffffff9273deaddeaddeaddeaddeaddeaddeaddeaddead0000855260018060a01b03166020850152604084015216606082015260a06080820152600060a08201520190565b9260c0949163ffffffff9360018060a01b038092168652166020850152604084015216606082015260a06080820152600060a08201520190565b908160209103126109d8575180151581036109d85790565b67ffffffffffffffff811161305357601f01601f191660200190565b919060209283818303126109d85780519067ffffffffffffffff82116109d857019281601f850112156109d857835161344c816133f8565b9261345a6040519485613069565b8184528282870101116109d85760005b81811061347f57508260009394955001015290565b858101830151848201840152820161346a565b9190820180921161302957565b8181029291811591840414171561302957565b6000806134fc9260018060a01b03169360208151910182865af13d15613545573d906134dd826133f8565b916134eb6040519384613069565b82523d6000602084013e5b8361354d565b805190811515918261352a575b50506135125750565b60249060405190635274afe760e01b82526004820152fd5b61353d92506020809183010191016133e0565b153880613509565b6060906134f6565b90613574575080511561356257805190602001fd5b604051630a12f52160e11b8152600490fd5b815115806135a7575b613585575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b1561357d565b60405163a9059cbb60e01b60208201526001600160a01b039290921660248301526044808301939093529181526135f1916135ec606483613069565b6134b2565b565b466208275003613608576004605360981b0190565b4661044d0361362957734f9a0e7fd2bf6067db6994cf12e4495df938e6e990565b46600a0361363c576006602160991b0190565b4661a4b10361365d577382af49447d8a07e3bd95bd0d56f35241523fbab190565b4661a4ba0361367e5773722e8bdd2ce80a4422e880164f2079488e11536590565b466101440361369f57735aea5775959fbc2557cc8789bc1bf90a239d9a9190565b46611388036136ba5760405163e31d668960e01b8152600490fd5b4660a9036136d4576040516322c20c6960e11b8152600490fd5b46612105036136e8576006602160991b0190565b604051639474dee160e01b8152600490fdfef1fa6dbbb3e74f1bef249f22545b238726c1cc79c5571a1416575674c7e63f3902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268001bc1b1207e0c319b7baad0317c09f30575218af6f5a59f1f0f15de8d217672cccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300cb1a101498aadad4159168d719f955e559e768912e6f5f5c57b094eb047d000fa2646970667358221220ef4804491b845c280e16b00222337281e5fac259475bb7e50c3744be7d00910864736f6c63430008140033",
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

// BridgeInitiateERC20ForStaking is a paid mutator transaction binding the contract method 0xba97e76e.
//
// Solidity: function BridgeInitiateERC20ForStaking(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value, uint256 nonce) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeInitiateERC20ForStaking(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeInitiateERC20ForStaking", sourceChainId, destChainId, to, ERC20Address, value, nonce)
}

// BridgeInitiateERC20ForStaking is a paid mutator transaction binding the contract method 0xba97e76e.
//
// Solidity: function BridgeInitiateERC20ForStaking(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value, uint256 nonce) returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeInitiateERC20ForStaking(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateERC20ForStaking(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, value, nonce)
}

// BridgeInitiateERC20ForStaking is a paid mutator transaction binding the contract method 0xba97e76e.
//
// Solidity: function BridgeInitiateERC20ForStaking(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value, uint256 nonce) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeInitiateERC20ForStaking(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateERC20ForStaking(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, value, nonce)
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

// BridgeInitiateETHForStaking is a paid mutator transaction binding the contract method 0x3b90b1f9.
//
// Solidity: function BridgeInitiateETHForStaking(uint256 sourceChainId, uint256 destChainId, address to, uint256 nonce) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeInitiateETHForStaking(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeInitiateETHForStaking", sourceChainId, destChainId, to, nonce)
}

// BridgeInitiateETHForStaking is a paid mutator transaction binding the contract method 0x3b90b1f9.
//
// Solidity: function BridgeInitiateETHForStaking(uint256 sourceChainId, uint256 destChainId, address to, uint256 nonce) payable returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeInitiateETHForStaking(sourceChainId *big.Int, destChainId *big.Int, to common.Address, nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateETHForStaking(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, nonce)
}

// BridgeInitiateETHForStaking is a paid mutator transaction binding the contract method 0x3b90b1f9.
//
// Solidity: function BridgeInitiateETHForStaking(uint256 sourceChainId, uint256 destChainId, address to, uint256 nonce) payable returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeInitiateETHForStaking(sourceChainId *big.Int, destChainId *big.Int, to common.Address, nonce *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateETHForStaking(&_L2PoolManager.TransactOpts, sourceChainId, destChainId, to, nonce)
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

// L2PoolManagerBridgeInitiateForStakingSuccessIterator is returned from FilterBridgeInitiateForStakingSuccess and is used to iterate over the raw logs and unpacked data for BridgeInitiateForStakingSuccess events raised by the L2PoolManager contract.
type L2PoolManagerBridgeInitiateForStakingSuccessIterator struct {
	Event *L2PoolManagerBridgeInitiateForStakingSuccess // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerBridgeInitiateForStakingSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerBridgeInitiateForStakingSuccess)
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
		it.Event = new(L2PoolManagerBridgeInitiateForStakingSuccess)
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
func (it *L2PoolManagerBridgeInitiateForStakingSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerBridgeInitiateForStakingSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerBridgeInitiateForStakingSuccess represents a BridgeInitiateForStakingSuccess event raised by the L2PoolManager contract.
type L2PoolManagerBridgeInitiateForStakingSuccess struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgeInitiateForStakingSuccess is a free log retrieval operation binding the contract event 0x48917de34e53153e06b5ed7e1c30d2458abfe4ba4ddb223cfce220c13131c5f6.
//
// Solidity: event BridgeInitiateForStakingSuccess(bytes32 indexed messageHash)
func (_L2PoolManager *L2PoolManagerFilterer) FilterBridgeInitiateForStakingSuccess(opts *bind.FilterOpts, messageHash [][32]byte) (*L2PoolManagerBridgeInitiateForStakingSuccessIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "BridgeInitiateForStakingSuccess", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerBridgeInitiateForStakingSuccessIterator{contract: _L2PoolManager.contract, event: "BridgeInitiateForStakingSuccess", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiateForStakingSuccess is a free log subscription operation binding the contract event 0x48917de34e53153e06b5ed7e1c30d2458abfe4ba4ddb223cfce220c13131c5f6.
//
// Solidity: event BridgeInitiateForStakingSuccess(bytes32 indexed messageHash)
func (_L2PoolManager *L2PoolManagerFilterer) WatchBridgeInitiateForStakingSuccess(opts *bind.WatchOpts, sink chan<- *L2PoolManagerBridgeInitiateForStakingSuccess, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "BridgeInitiateForStakingSuccess", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerBridgeInitiateForStakingSuccess)
				if err := _L2PoolManager.contract.UnpackLog(event, "BridgeInitiateForStakingSuccess", log); err != nil {
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

// ParseBridgeInitiateForStakingSuccess is a log parse operation binding the contract event 0x48917de34e53153e06b5ed7e1c30d2458abfe4ba4ddb223cfce220c13131c5f6.
//
// Solidity: event BridgeInitiateForStakingSuccess(bytes32 indexed messageHash)
func (_L2PoolManager *L2PoolManagerFilterer) ParseBridgeInitiateForStakingSuccess(log types.Log) (*L2PoolManagerBridgeInitiateForStakingSuccess, error) {
	event := new(L2PoolManagerBridgeInitiateForStakingSuccess)
	if err := _L2PoolManager.contract.UnpackLog(event, "BridgeInitiateForStakingSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
