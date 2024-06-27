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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ChainIdIsNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"ChainIdNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorBlockChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"MinTransferAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LessThanMinTransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MantaNotWETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MantleNotWETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughETH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"NotEnoughToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"TokenIsNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"sourceChainIdError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeMessageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stakeMessageHash\",\"type\":\"bytes32\"}],\"name\":\"FinalizeStakingMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeMessageNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"stakeMessageHash\",\"type\":\"bytes32\"}],\"name\":\"InitiateStakingMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20toL1Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawETHtoL1Success\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawWETHtoL1Success\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"shareAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeMessageNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeStakingMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"BridgeInitiateETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateStakingMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FeePoolValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FundingPoolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"IsSupportChainId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"IsSupportToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_GAS_Limit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinTransferAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"QuickSendAssertToUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ReLayer\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SupportTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UpdateFundingPoolBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawERC20ToL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawETHtoL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawWETHToL1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MultisigWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageManager\",\"outputs\":[{\"internalType\":\"contractIMessageManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_MinTransferAmount\",\"type\":\"uint256\"}],\"name\":\"setMinTransferAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_PerFee\",\"type\":\"uint256\"}],\"name\":\"setPerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setSupportERC20Token\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setValidChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingMessageNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60808060405234620000bd577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c16620000ae57506001600160401b036002600160401b03198282160162000068575b604051613d7a9081620000c38239f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a138808062000058565b63f92ee8a960e01b8152600490fd5b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a7146102975780630227047a1461029257806309dc480c1461028d5780630cd727df1461028857806310875a1314610283578063248a9ca31461027e5780632f2ff15d1461027957806331d98b3b14610274578063325e5d431461026f57806336568abe1461026a57806338731c471461026557806339891dd7146102605780633b0230f01461025b5780633f4ba83a14610256578063415ade1f14610251578063485cc9551461024c57806354dc027e146102475780635765a4eb146102425780635c975abb1461023d578063626417b514610238578063650c22761461023357806372fe6a7e1461022e5780638456cb591461022957806391d148541461022457806392d847381461021f5780639b4423801461021a578063a217fddf14610215578063a525ae2b14610210578063ab0f9e191461020b578063b1887e9b14610206578063b92e639614610201578063cd01c665146101fc578063d28597f4146101f7578063d547741f146101f2578063d73f14e5146101ed578063dac29568146101e8578063dbb0481f146101e3578063dd0c3460146101de578063e1c9e844146101d95763fa861848146101d457600080fd5b6134a6565b61332e565b613310565b6132d6565b613297565b61324e565b613202565b612a94565b61277e565b61275a565b61272d565b612696565b612668565b61264c565b612449565b612290565b612231565b6121b7565b612186565b612165565b61210a565b61208d565b611ebb565b611e81565b611d4b565b611be1565b611b6a565b61187c565b6111c7565b611003565b610f86565b610e06565b610de8565b610d9a565b610d5e565b610d35565b610cf3565b610cd5565b61037f565b346102ed5760203660031901126102ed5760043563ffffffff60e01b81168091036102ed57602090637965db0b60e01b81149081156102dc575b506040519015158152f35b6301ffc9a760e01b149050386102d1565b600080fd5b600435906001600160a01b03821682036102ed57565b602435906001600160a01b03821682036102ed57565b604435906001600160a01b03821682036102ed57565b606435906001600160a01b03821682036102ed57565b60609060031901126102ed576001600160a01b039060043582811681036102ed579160243590811681036102ed579060443590565b6103883661034a565b91610399610394613453565b613521565b6001600160a01b03811660009081526004602052604090206103c5906103c1905b5460ff1690565b1590565b610ca7574662082750036105755760405163095ea7b360e01b815273d8a791fe2be73eb6e6cf1eb0cb3f36adc9b3f8f960048201526024810184905260208160448160006001600160a01b0387165af1801561054257610547575b5061043c61043360095463ffffffff1690565b63ffffffff1690565b9273d8a791fe2be73eb6e6cf1eb0cb3f36adc9b3f8f993843b156102ed5760405163a93a4af960e01b81526001600160a01b03848116600483015285166024820152604481018390526064810182905294600091869160849183918591f190811561054257600080516020613cc58339815191529461051a92610529575b505b6001600160a01b03831660009081526005602052604090206104df8282546133fa565b9055604080514681524260208201526001600160a01b03948516918101919091529390921660608401526080830191909152819060a0820190565b0390a160405160018152602090f35b8061053661053c9261341d565b80610cca565b386104ba565b6138a0565b6105679060203d811161056e575b61055f8183613431565b8101906138ac565b5038610420565b503d610555565b4661044d036106995760405163095ea7b360e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede60048201526024810184905260208160448160006001600160a01b0387165af180156105425761067b575b50732a3dd3eb832af982ec71669e178424b10dca2ede92833b156102ed5760405163cd58657960e01b81526000600482018190526001600160a01b03808616602484015260448301849052841660648301526084820181905260c060a483015260c482018190529094859081838160e481015b03925af190811561054257600080516020613cc58339815191529461051a92610668575b506104bc565b806105366106759261341d565b38610662565b6106929060203d811161056e5761055f8183613431565b50386105cb565b46600a0361073f5760405163095ea7b360e01b81526010602160991b0160048201526024810184905260208160448160006001600160a01b0387165af1801561054257610721575b5060095463ffffffff166010602160991b01803b156102ed5761063e6000918693604051978880948193631474f2a960e31b8352838b8b60048601613866565b6107389060203d811161056e5761055f8183613431565b50386106e1565b46612105036107c75760405163095ea7b360e01b81526010602160991b0160048201526024810184905260208160448160006001600160a01b0387165af1801561054257610721575060095463ffffffff166010602160991b01803b156102ed5761063e6000918693604051978880948193631474f2a960e31b8352838b8b60048601613866565b4661a4b1036108de5760405163095ea7b360e01b81527309e9222e96e7b4ae2a407b98d48e330053351eee60048201526024810184905260208160448160006001600160a01b0387165af18015610542576108c0575b50604051637b3a3c8b60e01b81526001600160a01b0382811660048301528316602482015260448101849052608060648201526000608482018190529093908460a481837309e9222e96e7b4ae2a407b98d48e330053351eee5af190811561054257600080516020613cc58339815191529461051a9261089d57506104bc565b6108b9903d806000833e6108b18183613431565b8101906138e0565b5038610662565b6108d79060203d811161056e5761055f8183613431565b503861081d565b4661a4ba036109d25760405163095ea7b360e01b815273b2535b988dce19f9d71dfb22db6da744acac21bf60048201526024810184905260208160448160006001600160a01b0387165af18015610542576109b4575b50604051637b3a3c8b60e01b81526001600160a01b0382811660048301528316602482015260448101849052608060648201526000608482018190529093908460a4818373b2535b988dce19f9d71dfb22db6da744acac21bf5af190811561054257600080516020613cc58339815191529461051a9261089d57506104bc565b6109cb9060203d811161056e5761055f8183613431565b5038610934565b4661014403610a57577311f943b2c77b743ab90f4a0ae7d5a4e7fca3e10292833b156102ed57604051636ce5768960e11b81526001600160a01b0384811660048301528316602482015260448101829052936000908590606490829085905af190811561054257600080516020613cc58339815191529461051a9261066857506104bc565b4661138803610b015760405163095ea7b360e01b81526010602160991b0160048201526024810184905260208160448160006001600160a01b0387165af1801561054257610ae3575b5060095463ffffffff16926010602160991b01803b156102ed5760405163297e27a960e11b815294600091869182908490829061063e90888b8b60048601613866565b610afa9060203d811161056e5761055f8183613431565b5038610aa0565b4660a903610b885760405163095ea7b360e01b81526010602160991b0160048201526024810184905260208160448160006001600160a01b0387165af1801561054257610721575060095463ffffffff166010602160991b01803b156102ed5761063e6000918693604051978880948193631474f2a960e31b8352838b8b60048601613866565b4661a70e03610c955760405163095ea7b360e01b8152732c4813276869d93afdab4dd2b01cd670342da19460048201526024810184905260208160448160006001600160a01b0387165af1801561054257610c77575b50732c4813276869d93afdab4dd2b01cd670342da19492833b156102ed5760405163cd58657960e01b81526000600482018190526001600160a01b03858116602484015260448301849052841660648301526084820181905260c060a483015260c482018190529094859060e490829084905af190811561054257600080516020613cc58339815191529461051a9261066857506104bc565b610c8e9060203d811161056e5761055f8183613431565b5038610bde565b604051639474dee160e01b8152600490fd5b6040516305fd61ad60e01b81526001600160a01b03919091166004820152602490fd5b60009103126102ed57565b346102ed5760003660031901126102ed576020600154604051908152f35b346102ed5760403660031901126102ed57610d0c6102f2565b610d17610394613453565b6001600160a01b031660009081526005602052604090206024359055005b346102ed5760003660031901126102ed576000546040516001600160a01b039091168152602090f35b346102ed5760203660031901126102ed57600435600052600080516020613ca58339815191526020526020600160406000200154604051908152f35b346102ed5760403660031901126102ed57610de6600435610db9610308565b9080600052600080516020613ca5833981519152602052610de1600160406000200154613521565b613626565b005b346102ed5760003660031901126102ed576020600854604051908152f35b346102ed57610e143661034a565b90610e20610394613453565b6001600160a01b0383166000908152600460205260409020610e45906103c1906103ba565b610f65576001600160a01b0383166000908152600560205260409020610e6c8382546133fa565b90556001600160a01b039280841673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee8103610ed6575050814710610ec45760009283928392839283918315610eba575b1690f11561054257005b6108fc9250610eb0565b604051632c1d501360e11b8152600490fd5b6040516370a0823160e01b815230600482015291945090602081602481855afa8015610542578491600091610f37575b5010610f1657610de69350613aec565b6040516311745b6160e21b81526001600160a01b0385166004820152602490fd5b610f58915060203d8111610f5e575b610f508183613431565b81019061395e565b38610f06565b503d610f46565b6040516305fd61ad60e01b81526001600160a01b0384166004820152602490fd5b346102ed5760403660031901126102ed57610f9f610308565b336001600160a01b03821603610fbb57610de6906004356136c7565b60405163334bd91960e11b8152600490fd5b60c09060031901126102ed5760043590602435906044356001600160a01b03811681036102ed5790606435906084359060a43590565b61100c36610fcd565b61101c6103949695929496613453565b4682036111b55761103d6103c16103ba876000526003602052604060002090565b61119c576001600160a01b03861695600084888115611192575b600092839283928392f1156105425773eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee60005260056020527fa1829a9003092132f585b6ccdd167c19fe9774dbdea4260287e8a8e8ca8185d76110af8582546133fa565b90556000546110ce906001600160a01b03165b6001600160a01b031690565b803b156102ed57604051631eb65ffb60e01b815260048101889052602481018590526001600160a01b0392909216604483015260648201959095526084810184905260a481019190915292600090849060c490829084905af1928315610542577f61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f329361117f575b50604080519485526020850191909152830152309180606081015b0390a360405160018152602090f35b8061053661118c9261341d565b38611155565b6108fc9250611057565b604051632ef6974160e11b815260048101869052602490fd5b604051631a26aa4d60e21b8152600490fd5b6040806003193601126102ed576111dc6102f2565b602435906111eb610394613453565b47821161186b574662082750036113125761120e61043360095463ffffffff1690565b90736ea73e05adc79974b931123675ea8f78ffdacdf091823b156102ed5784516317e614fd60e11b81526001600160a01b038316600482015260248101859052604481018290529392600091859160649183918791f192831561054257600080516020613c85833981519152936112ff575b505b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee60005260056020527fa1829a9003092132f585b6ccdd167c19fe9774dbdea4260287e8a8e8ca8185d76112cb8382546133fa565b905583514681524260208201526001600160a01b039190911660408201526060810191909152608090a15160018152602090f35b8061053661130c9261341d565b38611280565b904661044d036113c557732a3dd3eb832af982ec71669e178424b10dca2ede91823b156102ed57835163cd58657960e01b81526000600482018190526001600160a01b038316602483015260448201849052606482018190526084820181905260c060a483015260c482018190529093849081858160e481015b03925af192831561054257600080516020613c85833981519152936113b2575b50611282565b806105366113bf9261341d565b386113ac565b9046600a0361140c5760095463ffffffff16906010602160991b01803b156102ed5761138c60009185948751968780948193631474f2a960e31b835283896004850161381a565b904661a4b1036114bd578251637b3a3c8b60e01b815273eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee60048201526001600160a01b038316602482015260448101829052608060648201526000608482018190529092908360a48185735288c571fd7ad117bea99bf60fe0846c4e84f9335af192831561054257600080516020613c85833981519152936114a25750611282565b6114b6903d806000833e6108b18183613431565b50386113ac565b4661a4ba03611552578251637b3a3c8b60e01b815273eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee60048201526001600160a01b038316602482015260448101829052608060648201526000608482018190529092908360a481857321903d3f8176b1a0c17e953cd896610be9ffdfa85af192831561054257600080516020613c85833981519152936114a25750611282565b46610144036115d1577311f943b2c77b743ab90f4a0ae7d5a4e7fca3e10291823b156102ed578351636ce5768960e11b81526001600160a01b03821660048201526000602482018190526044820184905290938490606490829086905af192831561054257600080516020613c85833981519152936113b25750611282565b46611388036116c457825163095ea7b360e01b81526010602160991b01600482015260248101829052602081604481600073deaddeaddeaddeaddeaddeaddeaddeaddead11115af18015610542576116a6575b5060095463ffffffff16916010602160991b01803b156102ed57845163297e27a960e11b815273deaddeaddeaddeaddeaddeaddeaddeaddead111160048201526001600160a01b03831660248201526044810184905263ffffffff909416606485015260a06084850152600060a4850181905290849081838160c4810161138c565b6116bd9060203d811161056e5761055f8183613431565b5038611624565b904660a90361170b5760095463ffffffff16906010602160991b01803b156102ed5761138c60009185948751968780948193631474f2a960e31b835283896004850161381a565b904661a70e0361181257825163095ea7b360e01b8152732c4813276869d93afdab4dd2b01cd670342da1946004820152602481018290526020816044816000734b21b980d0dc7d3c0c6175b0a412694f3a1c7c6b5af18015610542576117f4575b50732c4813276869d93afdab4dd2b01cd670342da19491823b156102ed57835163cd58657960e01b81526000600482018190526001600160a01b038316602483015260448201849052734b21b980d0dc7d3c0c6175b0a412694f3a1c7c6b60648301526084820181905260c060a483015260c482018190529093849081838160e4810161138c565b61180b9060203d811161056e5761055f8183613431565b503861176c565b90466121050361185a5760095463ffffffff16906010602160991b01803b156102ed5761138c60009185948751968780948193631474f2a960e31b835283896004850161381a565b8251639474dee160e01b8152600490fd5b8251632c1d501360e11b8152600490fd5b346102ed5760a03660031901126102ed57600480356024359161189d61031e565b6118a5610334565b608435468503611b59576118c96103c16103ba886000526003602052604060002090565b611b3d576001600160a01b03821660009081526004602052604090206118f2906103c1906103ba565b611b15576040516370a0823160e01b808252308287019081526001600160a01b03858116979095939092602092839083908190830103818c5afa91821561054257600092611af6575b506119488630338c61398d565b6040519081523084820190815283908290819060200103818c5afa9081156105425761197d93600092611ad9575b50506133fa565b926119a561199d8460018060a01b03166000526005602052604060002090565b91825461396d565b90556119e86119cc6119c56119bc6002548761397a565b620f4240900490565b80956133fa565b6001600160a01b03909316600090815260066020526040902090565b6119f384825461396d565b9055600054611a0a906001600160a01b03166110c2565b92833b156102ed576040805163305f478560e21b8152928301898152602081018b90526001600160a01b0388169181019190915260608101849052608081019190915290949260009186919082908490829060a00103925af1908115610542577fece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a494611ab792611ac6575b5060405193849316973397846040919493926060820195825260208201520152565b0390a460405160018152602090f35b80610536611ad39261341d565b38611a95565b611aef9250803d10610f5e57610f508183613431565b3880611976565b611b0e919250833d8511610f5e57610f508183613431565b903861193b565b506040516305fd61ad60e01b81526001600160a01b0390911681840190815281906020010390fd5b604051632ef6974160e11b815280850187815281906020010390fd5b604051631a26aa4d60e21b81528490fd5b346102ed5760003660031901126102ed57611b836134c9565b600080516020613ce5833981519152805460ff811615611bcf5760ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b604051638dfc202b60e01b8152600490fd5b346102ed5760c03660031901126102ed57611bfa6102f2565b611c02610308565b9081611c0c61031e565b604080516001600160a01b03848116602083019081529084169282019290925260643560608201819052608435608080840182905283529095909392601f1992611c5760a082613431565b519020604051635bf55fe360e01b602082019081526001600160a01b0397881660248301529685166044820152606481018890526084810186905260a4938401815292909190611ca79084613431565b622673c060a43560061b01603f5a0210611d1c57600080611cf2927f7f5489fdde5b12d2c5cb032ef5ac630450aac2bc94d3b3b363c1efacea7a3a6d98826111709751925af1613b25565b60405193849360018060a01b03809116981696846040919493926060820195825260208201520152565b6308c379a06000526020805278185361666543616c6c3a204e6f7420656e6f756768206761736058526064601cfd5b346102ed5760403660031901126102ed57611d646102f2565b611d6c610308565b90600080516020613d05833981519152549167ffffffffffffffff60ff8460401c1615931680159081611e79575b6001149081611e6f575b159081611e66575b50611e5457600080516020613d05833981519152805467ffffffffffffffff19166001179055611de09183611e2f5761375f565b611de657005b600080516020613d05833981519152805460ff60401b19169055604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1005b600080516020613d05833981519152805460ff60401b1916600160401b17905561375f565b60405163f92ee8a960e01b8152600490fd5b90501538611dac565b303b159150611da4565b849150611d9a565b346102ed5760203660031901126102ed576001600160a01b03611ea26102f2565b1660005260066020526020604060002054604051908152f35b346102ed57611ec936610fcd565b611eda610394969596949294613453565b4686036111b557611efb6103c16103ba876000526003602052604060002090565b61119c57611f0d6110c26110c2613b8f565b60405163a9059cbb60e01b81526001600160a01b038416600482015260248101859052600092916020908290604490829087905af180156105425761206f575b5073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc260005260056020527fa550ba85c46b24b567d2e17cd597f2283877afab43603f46d5de7858f1bdb731611f978582546133fa565b90558154611fad906001600160a01b03166110c2565b90813b1561206b57604051631eb65ffb60e01b815260048101889052602481018990526001600160a01b038516604482015260648101969096526084860185905260a4860152849060c490829084905af1928315610542577f7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc1493612058575b50604080519485526020850195909552938301526001600160a01b039092169130918060608101611170565b806105366120659261341d565b3861202c565b8280fd5b6120869060203d811161056e5761055f8183613431565b5038611f4d565b346102ed5760003660031901126102ed57602060ff600080516020613ce583398151915254166040519015158152f35b6007548110156120f45760076000527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880190600090565b634e487b7160e01b600052603260045260246000fd5b346102ed5760203660031901126102ed576004356007548110156102ed5760076000527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68801546040516001600160a01b039091168152602090f35b346102ed5760203660031901126102ed5761217e6134c9565b600435600255005b346102ed5760203660031901126102ed576004356000526003602052602060ff604060002054166040519015158152f35b346102ed5760003660031901126102ed576121d06134c9565b600080516020613ce5833981519152805460ff811661221f5760019060ff19161790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b60405163d93c066560e01b8152600490fd5b346102ed5760403660031901126102ed57602060ff612284612251610308565b600435600052600080516020613ca5833981519152845260406000209060018060a01b0316600052602052604060002090565b54166040519015158152f35b346102ed5760e03660031901126102ed576024356004356122af61031e565b6122b7610334565b90608435906122c7610394613453565b4685036111b5576122e86103c16103ba866000526003602052604060002090565b612430576001600160a01b0383166000908152600460205260409020612311906103c1906103ba565b610f65576001600160a01b038381169390919061234d90612333858488613aec565b6001600160a01b0316600090815260056020526040902090565b6123588482546133fa565b905560005461236f906001600160a01b03166110c2565b92833b156102ed57604051631eb65ffb60e01b815260048101879052602481018890526001600160a01b038316604482015260a4803560648301526084820183905260c480359183019190915290946000918691829084905af1908115610542577f0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe94611ab79261241d575b5060405193849316973097846040919493926060820195825260208201520152565b8061053661242a9261341d565b386123fb565b604051632ef6974160e11b815260048101859052602490fd5b60603660031901126102ed5760243560043561246361031e565b4682036111b5576124846103c16103ba856000526003602052604060002090565b61263357600154803410612613575073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee60005260056020527fa1829a9003092132f585b6ccdd167c19fe9774dbdea4260287e8a8e8ca8185d76124dc34825461396d565b90556124ed6119bc6002543461397a565b906124f882346133fa565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee6000526006602052907fa2e5aefc6e2cbe2917a296f0fd89c5f915c487c803db1d98eccb43f14012d71161254284825461396d565b9055600054612559906001600160a01b03166110c2565b803b156102ed5760405163305f478560e21b8152466004820152602481018790526001600160a01b0383166044820152606481018490526084810194909452600090849060a490829084905af1928315610542577f0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc593612600575b50604080519485526020850195909552938301526001600160a01b039092169133918060608101611170565b8061053661260d9261341d565b386125d4565b6040516358f8331360e11b81526004810191909152346024820152604490fd5b604051632ef6974160e11b815260048101849052602490fd5b346102ed5760003660031901126102ed57602060405160008152f35b346102ed5760003660031901126102ed57602063ffffffff60095416604051908152f35b801515036102ed57565b346102ed5760403660031901126102ed576126af6102f2565b6024356126bb8161268c565b6126c36134c9565b6001600160a01b039182166000818152600460205260409020805460ff191660ff84151516179055906126f257005b600754600160401b8110156127285780600161271192016007556120bd565b909283549160031b92831b921b1916179055600080f35b613407565b346102ed5760003660031901126102ed576020612748613b8f565b6040516001600160a01b039091168152f35b346102ed5760203660031901126102ed57612776610394613453565b600435600155005b346102ed5760803660031901126102ed576004602435813561279e61031e565b468203611b59576127bf6103c16103ba856000526003602052604060002090565b612a78576127d16110c26110c2613b8f565b604080516370a0823160e01b8082523082890190815292979360209390919084908490819083010381855afa92831561054257600093612a59575b5088516323b872dd60e01b815233868201908152306020820152606435604082015290919085908390819060600103816000875af1918215610542578592612a3c575b50895190815230868201908152909283918290819060200103915afa9081156105425761288493600092611ad95750506133fa565b600154808210612a19575073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc260005260056020527fa550ba85c46b24b567d2e17cd597f2283877afab43603f46d5de7858f1bdb7316128d882825461396d565b90556128f36128ec6119bc6002548461397a565b80926133fa565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26000526006602052917f5e5777fab7622aff3c042c1ece74307c2e9d699a9da444f416c35f2e1def28a561293d83825461396d565b9055600054612954906001600160a01b03166110c2565b91823b156102ed57875163305f478560e21b8152918201868152602081018890526001600160a01b038616604082015260608101859052608081019190915290939160009185919082908490829060a00103925af1928315610542577fcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a8993612a06575b508551938452602084019490945260408301526001600160a01b03909216913391606090a35160018152602090f35b80610536612a139261341d565b386129d7565b86516358f8331360e11b8152808401918252602082018390529081906040010390fd5b612a5290833d851161056e5761055f8183613431565b503861284f565b612a71919350843d8611610f5e57610f508183613431565b913861280c565b6040516363b5c9db60e01b815280850184815281906020010390fd5b6040806003193601126102ed57612aa96102f2565b60243590612ab8610394613453565b612ac66110c26110c2613b8f565b83516370a0823160e01b815230600480830191909152919291906020908181602481885afa908115610542576000916131e5575b5085116131c057466208275003612c9b57855163095ea7b360e01b8152818180612b4589878301919060206040840193737003e7b7186f0e6601203b99f7b8decbfa391cf981520152565b03816000895af1801561054257612c7d575b5050612b6b61043360095463ffffffff1690565b737003e7b7186f0e6601203b99f7b8decbfa391cf990813b156102ed57865163a93a4af960e01b81526001600160a01b03958616938101938452948416602084015260408301869052606083018190526000928592839185919083906080010393f190811561054257600080516020613d25833981519152928492612c6a575b505b84514681524260208201526001600160a01b039190911660408201526060810191909152608090a173c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26000526005602052612c5e7fa550ba85c46b24b567d2e17cd597f2283877afab43603f46d5de7858f1bdb7319182546133fa565b90555160018152602090f35b80610536612c779261341d565b38612beb565b81612c9392903d1061056e5761055f8183613431565b503880612b57565b90924661044d03612dbf57818651809263095ea7b360e01b825281600081612ce48b8b8301919060206040840193732a3dd3eb832af982ec71669e178424b10dca2ede81520152565b03925af1801561054257612da1575b5050732a3dd3eb832af982ec71669e178424b10dca2ede803b156102ed57845163cd58657960e01b815260009381018481526001600160a01b038416602082015260408101869052606081018590526080810185905260c060a082018190528101859052909391849182908790829060e0015b03925af190811561054257600080516020613d25833981519152928492612d8e575b50612bed565b80610536612d9b9261341d565b38612d88565b81612db792903d1061056e5761055f8183613431565b503880612cf3565b929046600a03612e6357855163095ea7b360e01b81526010602160991b018382019081526020810187905282908290819060400103816000895af1801561054257612e45575b505060095463ffffffff16906010602160991b01803b156102ed578594858593612d666000968b51998a9788968795631474f2a960e31b87528601613866565b81612e5b92903d1061056e5761055f8183613431565b503880612e05565b4661210503612ee757855163095ea7b360e01b81526010602160991b018382019081526020810187905282908290819060400103816000895af1801561054257612e4557505060095463ffffffff16906010602160991b01803b156102ed578594858593612d666000968b51998a9788968795631474f2a960e31b87528601613866565b92904661a4b10361301557838651809363095ea7b360e01b825281600081612f308b888301919060206040840193736c411ad3e74de3e7bd422b94a27770f5b86c623b81520152565b03925af19081156105425785946000938593612ff7575b50508651637b3a3c8b60e01b815273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc29181019182526001600160a01b0390921660208201526040810194909452608060608501819052600090850152928390819060a001038183736c411ad3e74de3e7bd422b94a27770f5b86c623b5af190811561054257600080516020613d25833981519152928492612fdc5750612bed565b612ff0903d806000833e6108b18183613431565b5038612d88565b8161300d92903d1061056e5761055f8183613431565b503880612f47565b4661a4ba0361312657838651809363095ea7b360e01b82528160008161305c8b888301919060206040840193737626841cb6113412f9c88d3adc720c9fac88d9ed81520152565b03925af19081156105425785946000938593613108575b50508651637b3a3c8b60e01b815273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc29181019182526001600160a01b0390921660208201526040810194909452608060608501819052600090850152928390819060a001038183737626841cb6113412f9c88d3adc720c9fac88d9ed5af190811561054257600080516020613d25833981519152928492612fdc5750612bed565b8161311e92903d1061056e5761055f8183613431565b503880613073565b90925046610144036131b2577311f943b2c77b743ab90f4a0ae7d5a4e7fca3e10292833b156102ed578551636ce5768960e11b81526001600160a01b038085169382019384529091166020830152604082018590529260009184919082908790829060600103925af190811561054257600080516020613d25833981519152928492612d8e5750612bed565b8451639474dee160e01b8152fd5b85516311745b6160e21b81526001600160a01b03851681840190815281906020010390fd5b6131fc9150823d8411610f5e57610f508183613431565b38612afa565b346102ed5760403660031901126102ed57610de6600435613221610308565b9080600052600080516020613ca5833981519152602052613249600160406000200154613521565b6136c7565b346102ed5760403660031901126102ed57610de660243561326e8161268c565b6132766134c9565b600435600052600360205260406000209060ff801983541691151516179055565b346102ed5760203660031901126102ed576001600160a01b036132b86102f2565b166000526004602052602060ff604060002054166040519015158152f35b346102ed5760203660031901126102ed576001600160a01b036132f76102f2565b1660005260056020526020604060002054604051908152f35b346102ed5760003660031901126102ed576020600254604051908152f35b346102ed5761333c3661034a565b600854604080516001600160a01b038681166020830190815290861692820192909252606081018490526080808201849052815293949361337e60a082613431565b5190209360001982146133e45760019190910160088190556040805193845260208401919091526001600160a01b0391821693909116917f5e2d465404712c588a7dffc8622a4fad4f4522708eaf138e2b70910cd36992a09190a4602060405160018152f35b634e487b7160e01b600052601160045260246000fd5b919082039182116133e457565b634e487b7160e01b600052604160045260246000fd5b67ffffffffffffffff811161272857604052565b90601f8019910116810190811067ffffffffffffffff82111761272857604052565b604051602081017f33fe247d78fee2e7fd135c405eda4bd2911c0a73c0a81b36c3bcc967dd06e5ad8152602082526040820182811067ffffffffffffffff821117612728576040529051902060ff191690565b346102ed5760003660031901126102ed5760206134c1613453565b604051908152f35b3360009081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604090205460ff161561350257565b60405163e2517d3f60e01b815233600482015260006024820152604490fd5b6000818152600080516020613ca58339815191526020908152604080832033845290915290205460ff16156135535750565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b6001600160a01b03811660009081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d6020526040812054600080516020613ca58339815191529060ff1661362057818052602090815260408083206001600160a01b038516600090815292529020805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50905090565b6000818152600080516020613ca5833981519152602081815260408084206001600160a01b038716855290915282205491929160ff166136c057818352602090815260408084206001600160a01b038616600090815292529020805460ff1916600117905533926001600160a01b0316917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b5050905090565b6000818152600080516020613ca5833981519152602081815260408084206001600160a01b038716855290915282205491929160ff16156136c057818352602090815260408084206001600160a01b038616600090815292529020805460ff1916905533926001600160a01b0316917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b6137b19061376b6137eb565b6137736137eb565b61377b6137eb565b600080516020613ce5833981519152805460ff1916905561379a6137eb565b67016345785d8a0000600155612710600255613571565b5060018060a01b03166bffffffffffffffffffffffff60a01b60005416176000556001600855620493e063ffffffff196009541617600955565b60ff600080516020613d058339815191525460401c161561380857565b604051631afcd79f60e31b8152600490fd5b919260c09363ffffffff9273deaddeaddeaddeaddeaddeaddeaddeaddead0000855260018060a01b03166020850152604084015216606082015260a06080820152600060a08201520190565b9260c0949163ffffffff9360018060a01b038092168652166020850152604084015216606082015260a06080820152600060a08201520190565b6040513d6000823e3d90fd5b908160209103126102ed57516138c18161268c565b90565b67ffffffffffffffff811161272857601f01601f191660200190565b919060209283818303126102ed5780519067ffffffffffffffff82116102ed57019281601f850112156102ed578351613918816138c4565b926139266040519485613431565b8184528282870101116102ed5760005b81811061394b57508260009394955001015290565b8581018301518482018401528201613936565b908160209103126102ed575190565b919082018092116133e457565b818102929181159184041417156133e457565b6040516323b872dd60e01b60208201526001600160a01b03928316602482015292909116604483015260648201929092526139de916139d982608481015b03601f198101845283613431565b6139e0565b565b600080613a2a9260018060a01b03169360208151910182865af13d15613a81573d90613a0b826138c4565b91613a196040519384613431565b82523d6000602084013e5b83613a89565b8051908115159182613a63575b5050613a405750565b604051635274afe760e01b81526001600160a01b03919091166004820152602490fd5b613a7a9250906020806103c19383010191016138ac565b3880613a37565b606090613a24565b90613ab05750805115613a9e57805190602001fd5b604051630a12f52160e11b8152600490fd5b81511580613ae3575b613ac1575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b15613ab9565b60405163a9059cbb60e01b60208201526001600160a01b03909216602483015260448201929092526139de916139d982606481016139cb565b15613b2c57565b60405162461bcd60e51b815260206004820152603560248201527f546f6b656e4272696467652e42726964676546696e616c697a655374616b696e60448201527419d3595cdcd859d94e8818d85b1b0819985a5b1959605a1b6064820152608490fd5b466208275003613ba4576004605360981b0190565b4661044d03613bc557734f9a0e7fd2bf6067db6994cf12e4495df938e6e990565b46600a03613bd8576006602160991b0190565b4661a4b103613bf9577382af49447d8a07e3bd95bd0d56f35241523fbab190565b4661a4ba03613c1a5773722e8bdd2ce80a4422e880164f2079488e11536590565b4661014403613c3b57735aea5775959fbc2557cc8789bc1bf90a239d9a9190565b4661138803613c565760405163e31d668960e01b8152600490fd5b4660a903613c70576040516322c20c6960e11b8152600490fd5b4661210503610c95576006602160991b019056fef1fa6dbbb3e74f1bef249f22545b238726c1cc79c5571a1416575674c7e63f3902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268001bc1b1207e0c319b7baad0317c09f30575218af6f5a59f1f0f15de8d217672cccd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00cb1a101498aadad4159168d719f955e559e768912e6f5f5c57b094eb047d000fa26469706673582212203f4f7b05cf31b4ad83024ac21bfff6161bd1b455ccab68ce1820b6c02fb3ebfb64736f6c63430008140033",
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

// StakingMessageNumber is a free data retrieval call binding the contract method 0x31d98b3b.
//
// Solidity: function stakingMessageNumber() view returns(uint256)
func (_L2PoolManager *L2PoolManagerCaller) StakingMessageNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2PoolManager.contract.Call(opts, &out, "stakingMessageNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingMessageNumber is a free data retrieval call binding the contract method 0x31d98b3b.
//
// Solidity: function stakingMessageNumber() view returns(uint256)
func (_L2PoolManager *L2PoolManagerSession) StakingMessageNumber() (*big.Int, error) {
	return _L2PoolManager.Contract.StakingMessageNumber(&_L2PoolManager.CallOpts)
}

// StakingMessageNumber is a free data retrieval call binding the contract method 0x31d98b3b.
//
// Solidity: function stakingMessageNumber() view returns(uint256)
func (_L2PoolManager *L2PoolManagerCallerSession) StakingMessageNumber() (*big.Int, error) {
	return _L2PoolManager.Contract.StakingMessageNumber(&_L2PoolManager.CallOpts)
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

// BridgeFinalizeStakingMessage is a paid mutator transaction binding the contract method 0x415ade1f.
//
// Solidity: function BridgeFinalizeStakingMessage(address shareAddress, address from, address to, uint256 shares, uint256 stakeMessageNonce, uint256 gasLimit) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeFinalizeStakingMessage(opts *bind.TransactOpts, shareAddress common.Address, from common.Address, to common.Address, shares *big.Int, stakeMessageNonce *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeFinalizeStakingMessage", shareAddress, from, to, shares, stakeMessageNonce, gasLimit)
}

// BridgeFinalizeStakingMessage is a paid mutator transaction binding the contract method 0x415ade1f.
//
// Solidity: function BridgeFinalizeStakingMessage(address shareAddress, address from, address to, uint256 shares, uint256 stakeMessageNonce, uint256 gasLimit) returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeFinalizeStakingMessage(shareAddress common.Address, from common.Address, to common.Address, shares *big.Int, stakeMessageNonce *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeFinalizeStakingMessage(&_L2PoolManager.TransactOpts, shareAddress, from, to, shares, stakeMessageNonce, gasLimit)
}

// BridgeFinalizeStakingMessage is a paid mutator transaction binding the contract method 0x415ade1f.
//
// Solidity: function BridgeFinalizeStakingMessage(address shareAddress, address from, address to, uint256 shares, uint256 stakeMessageNonce, uint256 gasLimit) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeFinalizeStakingMessage(shareAddress common.Address, from common.Address, to common.Address, shares *big.Int, stakeMessageNonce *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeFinalizeStakingMessage(&_L2PoolManager.TransactOpts, shareAddress, from, to, shares, stakeMessageNonce, gasLimit)
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

// BridgeInitiateStakingMessage is a paid mutator transaction binding the contract method 0xe1c9e844.
//
// Solidity: function BridgeInitiateStakingMessage(address from, address to, uint256 shares) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactor) BridgeInitiateStakingMessage(opts *bind.TransactOpts, from common.Address, to common.Address, shares *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.contract.Transact(opts, "BridgeInitiateStakingMessage", from, to, shares)
}

// BridgeInitiateStakingMessage is a paid mutator transaction binding the contract method 0xe1c9e844.
//
// Solidity: function BridgeInitiateStakingMessage(address from, address to, uint256 shares) returns(bool)
func (_L2PoolManager *L2PoolManagerSession) BridgeInitiateStakingMessage(from common.Address, to common.Address, shares *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateStakingMessage(&_L2PoolManager.TransactOpts, from, to, shares)
}

// BridgeInitiateStakingMessage is a paid mutator transaction binding the contract method 0xe1c9e844.
//
// Solidity: function BridgeInitiateStakingMessage(address from, address to, uint256 shares) returns(bool)
func (_L2PoolManager *L2PoolManagerTransactorSession) BridgeInitiateStakingMessage(from common.Address, to common.Address, shares *big.Int) (*types.Transaction, error) {
	return _L2PoolManager.Contract.BridgeInitiateStakingMessage(&_L2PoolManager.TransactOpts, from, to, shares)
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

// L2PoolManagerFinalizeStakingMessageIterator is returned from FilterFinalizeStakingMessage and is used to iterate over the raw logs and unpacked data for FinalizeStakingMessage events raised by the L2PoolManager contract.
type L2PoolManagerFinalizeStakingMessageIterator struct {
	Event *L2PoolManagerFinalizeStakingMessage // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerFinalizeStakingMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerFinalizeStakingMessage)
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
		it.Event = new(L2PoolManagerFinalizeStakingMessage)
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
func (it *L2PoolManagerFinalizeStakingMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerFinalizeStakingMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerFinalizeStakingMessage represents a FinalizeStakingMessage event raised by the L2PoolManager contract.
type L2PoolManagerFinalizeStakingMessage struct {
	From              common.Address
	To                common.Address
	Shares            *big.Int
	StakeMessageNonce *big.Int
	StakeMessageHash  [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFinalizeStakingMessage is a free log retrieval operation binding the contract event 0x7f5489fdde5b12d2c5cb032ef5ac630450aac2bc94d3b3b363c1efacea7a3a6d.
//
// Solidity: event FinalizeStakingMessage(address indexed from, address indexed to, uint256 shares, uint256 stakeMessageNonce, bytes32 stakeMessageHash)
func (_L2PoolManager *L2PoolManagerFilterer) FilterFinalizeStakingMessage(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2PoolManagerFinalizeStakingMessageIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "FinalizeStakingMessage", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerFinalizeStakingMessageIterator{contract: _L2PoolManager.contract, event: "FinalizeStakingMessage", logs: logs, sub: sub}, nil
}

// WatchFinalizeStakingMessage is a free log subscription operation binding the contract event 0x7f5489fdde5b12d2c5cb032ef5ac630450aac2bc94d3b3b363c1efacea7a3a6d.
//
// Solidity: event FinalizeStakingMessage(address indexed from, address indexed to, uint256 shares, uint256 stakeMessageNonce, bytes32 stakeMessageHash)
func (_L2PoolManager *L2PoolManagerFilterer) WatchFinalizeStakingMessage(opts *bind.WatchOpts, sink chan<- *L2PoolManagerFinalizeStakingMessage, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "FinalizeStakingMessage", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerFinalizeStakingMessage)
				if err := _L2PoolManager.contract.UnpackLog(event, "FinalizeStakingMessage", log); err != nil {
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

// ParseFinalizeStakingMessage is a log parse operation binding the contract event 0x7f5489fdde5b12d2c5cb032ef5ac630450aac2bc94d3b3b363c1efacea7a3a6d.
//
// Solidity: event FinalizeStakingMessage(address indexed from, address indexed to, uint256 shares, uint256 stakeMessageNonce, bytes32 stakeMessageHash)
func (_L2PoolManager *L2PoolManagerFilterer) ParseFinalizeStakingMessage(log types.Log) (*L2PoolManagerFinalizeStakingMessage, error) {
	event := new(L2PoolManagerFinalizeStakingMessage)
	if err := _L2PoolManager.contract.UnpackLog(event, "FinalizeStakingMessage", log); err != nil {
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

// L2PoolManagerInitiateStakingMessageIterator is returned from FilterInitiateStakingMessage and is used to iterate over the raw logs and unpacked data for InitiateStakingMessage events raised by the L2PoolManager contract.
type L2PoolManagerInitiateStakingMessageIterator struct {
	Event *L2PoolManagerInitiateStakingMessage // Event containing the contract specifics and raw log

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
func (it *L2PoolManagerInitiateStakingMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PoolManagerInitiateStakingMessage)
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
		it.Event = new(L2PoolManagerInitiateStakingMessage)
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
func (it *L2PoolManagerInitiateStakingMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PoolManagerInitiateStakingMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PoolManagerInitiateStakingMessage represents a InitiateStakingMessage event raised by the L2PoolManager contract.
type L2PoolManagerInitiateStakingMessage struct {
	From              common.Address
	To                common.Address
	Shares            *big.Int
	StakeMessageNonce *big.Int
	StakeMessageHash  [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInitiateStakingMessage is a free log retrieval operation binding the contract event 0x5e2d465404712c588a7dffc8622a4fad4f4522708eaf138e2b70910cd36992a0.
//
// Solidity: event InitiateStakingMessage(address indexed from, address indexed to, uint256 shares, uint256 stakeMessageNonce, bytes32 indexed stakeMessageHash)
func (_L2PoolManager *L2PoolManagerFilterer) FilterInitiateStakingMessage(opts *bind.FilterOpts, from []common.Address, to []common.Address, stakeMessageHash [][32]byte) (*L2PoolManagerInitiateStakingMessageIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var stakeMessageHashRule []interface{}
	for _, stakeMessageHashItem := range stakeMessageHash {
		stakeMessageHashRule = append(stakeMessageHashRule, stakeMessageHashItem)
	}

	logs, sub, err := _L2PoolManager.contract.FilterLogs(opts, "InitiateStakingMessage", fromRule, toRule, stakeMessageHashRule)
	if err != nil {
		return nil, err
	}
	return &L2PoolManagerInitiateStakingMessageIterator{contract: _L2PoolManager.contract, event: "InitiateStakingMessage", logs: logs, sub: sub}, nil
}

// WatchInitiateStakingMessage is a free log subscription operation binding the contract event 0x5e2d465404712c588a7dffc8622a4fad4f4522708eaf138e2b70910cd36992a0.
//
// Solidity: event InitiateStakingMessage(address indexed from, address indexed to, uint256 shares, uint256 stakeMessageNonce, bytes32 indexed stakeMessageHash)
func (_L2PoolManager *L2PoolManagerFilterer) WatchInitiateStakingMessage(opts *bind.WatchOpts, sink chan<- *L2PoolManagerInitiateStakingMessage, from []common.Address, to []common.Address, stakeMessageHash [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var stakeMessageHashRule []interface{}
	for _, stakeMessageHashItem := range stakeMessageHash {
		stakeMessageHashRule = append(stakeMessageHashRule, stakeMessageHashItem)
	}

	logs, sub, err := _L2PoolManager.contract.WatchLogs(opts, "InitiateStakingMessage", fromRule, toRule, stakeMessageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PoolManagerInitiateStakingMessage)
				if err := _L2PoolManager.contract.UnpackLog(event, "InitiateStakingMessage", log); err != nil {
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

// ParseInitiateStakingMessage is a log parse operation binding the contract event 0x5e2d465404712c588a7dffc8622a4fad4f4522708eaf138e2b70910cd36992a0.
//
// Solidity: event InitiateStakingMessage(address indexed from, address indexed to, uint256 shares, uint256 stakeMessageNonce, bytes32 indexed stakeMessageHash)
func (_L2PoolManager *L2PoolManagerFilterer) ParseInitiateStakingMessage(log types.Log) (*L2PoolManagerInitiateStakingMessage, error) {
	event := new(L2PoolManagerInitiateStakingMessage)
	if err := _L2PoolManager.contract.UnpackLog(event, "InitiateStakingMessage", log); err != nil {
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
