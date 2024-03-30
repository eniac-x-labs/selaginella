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

// IL1PoolManagerKeyValuePair is an auto generated low-level Go binding around an user-defined struct.
type IL1PoolManagerKeyValuePair struct {
	Key   common.Address
	Value *big.Int
}

// IL1PoolManagerPool is an auto generated low-level Go binding around an user-defined struct.
type IL1PoolManagerPool struct {
	StartTimestamp  uint32
	EndTimestamp    uint32
	Token           common.Address
	TotalAmount     *big.Int
	TotalFee        *big.Int
	TotalFeeClaimed *big.Int
	IsCompleted     bool
}

// IL1PoolManagerUser is an auto generated low-level Go binding around an user-defined struct.
type IL1PoolManagerUser struct {
	IsWithdrawed bool
	Token        common.Address
	StartPoolId  *big.Int
	EndPoolId    *big.Int
	Amount       *big.Int
}

// L1PoolManagerMetaData contains all meta data concerning the L1PoolManager contract.
var L1PoolManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ChainIdIsNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"ChainIdNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorBlockChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"providedAmount\",\"type\":\"uint256\"}],\"name\":\"LessThanMinStakeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"MinTransferAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LessThanMinTransferAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LessThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MantaNotWETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MantleNotWETH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PoolIndex\",\"type\":\"uint256\"}],\"name\":\"NewPoolIsNotCreate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReward\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughETH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"NotEnoughToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"PoolLength\",\"type\":\"uint256\"}],\"name\":\"OutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolIndex\",\"type\":\"uint256\"}],\"name\":\"PoolIsCompleted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"name\":\"TokenIsAlreadySupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"TokenIsNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"sourceChainIdError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endPoolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Reward\",\"type\":\"uint256\"}],\"name\":\"ClaimReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolIndex\",\"type\":\"uint256\"}],\"name\":\"CompletePoolEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetMinStakeAmountEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSupport\",\"type\":\"bool\"}],\"name\":\"SetSupportTokenEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingETHEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingWETHEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StarkingERC20Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Blockchain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferAssertTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Reward\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakingManager\",\"type\":\"address\"}],\"name\":\"BridgeFinalizeETHForStaking\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"BridgeInitiateETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ClaimAllReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"ClaimbyID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"IsCompleted\",\"type\":\"bool\"}],\"internalType\":\"structIL1PoolManager.Pool[]\",\"name\":\"CompletePools\",\"type\":\"tuple[]\"}],\"name\":\"CompletePoolAndNew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DepositAndStaking\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DepositAndStakingERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DepositAndStakingETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositAndStakingWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FeePoolValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FundingPoolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"IsSupportChainId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"IsSupportToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MinStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinTransferAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Pools\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"IsCompleted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"QuickSendAssertToUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ReLayer\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SupportTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"Blockchain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TransferAssertToBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UpdateFundingPoolBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Users\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isWithdrawed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"StartPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WithdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"WithdrawByID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getPool\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"IsCompleted\",\"type\":\"bool\"}],\"internalType\":\"structIL1PoolManager.Pool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getPoolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrincipal\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structIL1PoolManager.KeyValuePair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structIL1PoolManager.KeyValuePair[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isWithdrawed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"StartPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"}],\"internalType\":\"structIL1PoolManager.User[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isWithdrawed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"StartPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"}],\"internalType\":\"structIL1PoolManager.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MultisigWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageManager\",\"outputs\":[{\"internalType\":\"contractIMessageManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodTime\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setMinStakeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_MinTransferAmount\",\"type\":\"uint256\"}],\"name\":\"setMinTransferAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_PerFee\",\"type\":\"uint256\"}],\"name\":\"setPerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setSupportERC20Token\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isSupport\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"startTimes\",\"type\":\"uint32\"}],\"name\":\"setSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setValidChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60808060405234620000bd577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c16620000ae57506001600160401b036002600160401b03198282160162000068575b6040516160c89081620000c38239f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a138808062000058565b63f92ee8a960e01b8152600490fd5b600080fdfe60a080604052600436101561001357600080fd5b60008060805260e0908035821c92836301ffc9a714614ef1575050816309dc480c14614ed15781630cd727df14614e8e57816310875a1314614e6557816313e8544e14614dda5781631ca2f17314614d9e5781631d31fac014614d78578163248a9ca314614d3a5781632c37388b146149c85781632f2ff15d1461497a578163325e5d43146149425781633338562c1461467657816334d34ef5146143d857816336568abe1461439157816338731c471461425f5781633b0230f01461400d5781633d18b91214613db75781633f4ba83a14613d3a5781634663cdc814613c96578163485cc95514613ac65781634dfaef841461394257816354dc027e146139065781635765a4eb146137825781635a648bc5146133225781635b5b9ea2146132085781635c975abb146131d6578163626417b514613194578163650c22761461316f5781636f77926b1461306d57816372fe6a7e1461303a5781638456cb5914612fcd57816391d1485414612f7057816392d8473814612de757816396f984e614612ce75781639ab0b6521461272c5781639b44238014612570578163a217fddf14612552578163a69bdf16146123f2578163ab0f9e1914612386578163b1887e9b14612357578163b92e63961461232f578163bc42493d14612107578163cb314fab14612098578163cb4f04ad1461205c578163cc9c021814611ec357508063cd01c66514611bce578063d547741f14611b79578063d73f14e514611b2d578063dac2956814611aec578063dbb0481f14611ab0578063dd0c346014611a90578063e2070893146118ee578063f62a89cf14611538578063f8fcc2aa146114fc578063f91fa9a814610380578063fa861848146103585763ff2bf64f1461029b57600080fd5b34610352576040366003190112610352576102b4614f4c565b6040516102c08161503d565b60805181526080516020820152608051604082015260805160608201526080805191015260018060a01b031660805152600a60205260a061031161030b602435604060805120614fa9565b50615ac2565b61035060405180926080809180511515845260018060a01b03602082015116602085015260408101516040850152606081015160608501520151910152565bf35b60805180fd5b34610352576080518060031936011261037d5760206103756151b2565b604051908152f35b80fd5b346103525760803660031901126103525760043561039c614f67565b6103a4614f7d565b606435906103b86103b36151b2565b61524c565b60018060a01b03918284169384608051526020906004825260ff6040608051205416156114e3576208275087036107e85760008051602061601383398151915286036105375750604051636bb825d760e11b81526202981060048201819052908281602481730d7e906bd9cafa154b048cfa766cc1e54e39af9b5afa80156104fb5760805190610508575b61044e915084615533565b90737f2b8c31f88b6006c382775eea88297ec1e3e90590813b1561035257606460405180948193636705b1e760e11b83528a8a1660048401528860248401526044830152608051945af180156104fb576104e6575b5091600080516020615ff383398151915293916040935b86608051526005815284608051206104d383825461518f565b905584519788528701521693a360805180f35b6104ef9061500d565b60805180156104a35780fd5b6040513d608051823e3d90fd5b508281813d8311610530575b61051e8183615059565b810103126103525761044e9051610443565b503d610514565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc286036106bb57604051636bb825d760e11b8152614e2060048201528281602481730d7e906bd9cafa154b048cfa766cc1e54e39af9b5afa9081156104fb578391610692575b50506040519063095ea7b360e01b8252737ac440cae8eb6328de4fa621163a792c1ea9d4fe9182600482015284602482015283816044816080518c5af180156104fb57610665575b50813b156103525760405163790cfd3360e11b81526080516001600160a01b039283166004830152918616602482015260448101859052614e2060648201529182908180608481015b0391608051905af180156104fb57610650575b5091600080516020615ff383398151915293916040936104ba565b6106599061500d565b60805180156106355780fd5b61068490843d861161068b575b61067c8183615059565b8101906155c3565b50886105d9565b503d610672565b813d83116106b4575b6106a58183615059565b81010312610352578188610591565b503d61069b565b604051636bb825d760e11b8152614e2060048201528281602481730d7e906bd9cafa154b048cfa766cc1e54e39af9b5afa9081156104fb5783916107bf575b505060405163095ea7b360e01b8152737ac440cae8eb6328de4fa621163a792c1ea9d4fe600482015283602482015282816044816080518b5af180156104fb576107a2575b5073d8a791fe2be73eb6e6cf1eb0cb3f36adc9b3f8f990813b156103525760405163790cfd3360e11b81526080516001600160a01b039283166004830152918616602482015260448101859052614e206064820152918290818060848101610622565b6107b890833d851161068b5761067c8183615059565b508761073f565b813d83116107e1575b6107d28183615059565b810103126103525781886106fa565b503d6107c8565b61044d870361098b5760008051602061601383398151915286036108a15750732a3dd3eb832af982ec71669e178424b10dca2ede803b156103525760405163cd58657960e01b8152608051600160048301526001600160a01b0386166024830152604482018590526000606483018190526084830181905260c060a484015260c48301529091829081868160e481015b03925af180156104fb57610650575091600080516020615ff383398151915293916040936104ba565b6040519063095ea7b360e01b8252732a3dd3eb832af982ec71669e178424b10dca2ede9182600482015284602482015283816044816080518c5af180156104fb5761096e575b50813b156103525760405163cd58657960e01b815260808051600160048401526001600160a01b0388811660248501526044840188905293909316606483015260006084830181905260c060a484015260c4830152519092839160e4918391905af180156104fb57610650575091600080516020615ff383398151915293916040936104ba565b61098490843d861161068b5761067c8183615059565b50886108e7565b600a8703610aad576000805160206160138339815191528603610a0957507399c9fc46f92e8a1c0dec1b1747d010903e884be1803b1561035257604051639a2ac6d560e01b81526080516001600160a01b03861660048301526000602483018190526060604484015260648301529091829081868160848101610878565b610a1281615b0c565b906040519163095ea7b360e01b83527399c9fc46f92e8a1c0dec1b1747d010903e884be19283600482015285602482015284816044816080518d5af180156104fb57610a90575b5063ffffffff5a1692803b15610352578686926106226040519687958694859463041c592960e51b86526080519960048701615a41565b610aa690853d871161068b5761067c8183615059565b5089610a59565b91939092909161a4b18703610c93576000805160206160138339815191528603610b57575060405193634fb1a07b60e01b85526080518580610af4848630600485016159ee565b0381847372ce9c846789fdb6fc1f34ac4ad25dd9ef7031ef5af19384156104fb57600080516020615ff383398151915295604095610b33575b506104ba565b610b50903d80608051833e610b488183615059565b81019061592a565b5088610b2d565b9373c02aaa39b223fe8d0a0e5c4f27ead9083c756cc28603610c28576040519463095ea7b360e01b865273d92023e9d9911199a6711321d1277285e6d4e2db9586600482015282602482015285816044816080518c5af180156104fb57610c0b575b506040518096634fb1a07b60e01b82528180610bdf8688608051973090600486016159a8565b0391608051905af19384156104fb57600080516020615ff383398151915295604095610b3357506104ba565b610c2190863d881161068b5761067c8183615059565b5088610bb9565b6040519463095ea7b360e01b865273a3a7b6f88361f48403514059f1f16c8e78d60eec9586600482015282602482015285816044816080518c5af180156104fb57610c0b57506040518096634fb1a07b60e01b82528180610bdf8688608051973090600486016159a8565b61a4ba8703610e04576000805160206160138339815191528603610d12575060405193634fb1a07b60e01b85526080518580610cd4848630600485016159ee565b03818473c840838bc438d73c16c2f8b22d2ce3669963cd485af19384156104fb57600080516020615ff383398151915295604095610b3357506104ba565b9373c02aaa39b223fe8d0a0e5c4f27ead9083c756cc28603610d99576040519463095ea7b360e01b865273e4e2121b479017955be0b175305b35f312330bae9586600482015282602482015285816044816080518c5af180156104fb57610c0b57506040518096634fb1a07b60e01b82528180610bdf8688608051973090600486016159a8565b6040519463095ea7b360e01b865273b2535b988dce19f9d71dfb22db6da744acac21bf9586600482015282602482015285816044816080518c5af180156104fb57610c0b57506040518096634fb1a07b60e01b82528180610bdf8688608051973090600486016159a8565b919290916101448703610f9f57506000805160206160138339815191528503610ed15760405163e8b99b1b60e01b8152828416600482015260808051602483015260448201869052805160648301525160848201523060a4820152818160c481887357891966931eb4bb6fb81430e6ce0a03aabde0635af180156104fb57610ea2575b5090600080516020615ff383398151915293604093926104ba565b9080939291813d8311610eca575b610eba8183615059565b8101031261035257909186610e87565b503d610eb0565b60405163095ea7b360e01b81527357891966931eb4bb6fb81430e6ce0a03aabde06380600483015285602483015282826044816080518b5af19081156104fb5760c4928492610f82575b506040519283809263e8b99b1b60e01b825288881660048301528a6024830152896044830152608051606483015260805160848301523060a4830152608051905af180156104fb57610ea2575090600080516020615ff383398151915293604093926104ba565b610f9890833d851161068b5761067c8183615059565b5089610f1b565b92939192611388870361112c57600080516020616013833981519152860361104757507395fc37a27a2f68e3a647cdc081f0a89bb47c3012803b1561035257604051639a2ac6d560e01b8152608080516001600160a01b0387166004840152600060248401819052606060448501526064840152905191929091839160849183915af180156104fb57610650575091600080516020615ff383398151915293916040936104ba565b60405163095ea7b360e01b81527395fc37a27a2f68e3a647cdc081f0a89bb47c30129182600483015284602483015283826044816080518c5af19182156104fb576110979261110f575b50615bf0565b90803b1561035257856040519263041c592960e51b84528860048501521660248301528585166044830152836064830152608051608483015260c060a483015260805160c48301528160e48160805180945af180156104fb57610650575091600080516020615ff383398151915293916040936104ba565b61112590853d871161068b5761067c8183615059565b5089611091565b60a987036112525760008051602061601383398151915286036111cc5750733b95bc951ee0f553ba487327278cac44f29715e5803b1561035257604051639a2ac6d560e01b81526080516001600160a01b038616600483015260006024830181905260606044840152606483015290918290608490829087905af180156104fb57610650575091600080516020615ff383398151915293916040936104ba565b6111d581615c61565b906040519163095ea7b360e01b8352733b95bc951ee0f553ba487327278cac44f29715e59283600482015285602482015284816044816080518d5af180156104fb57610a90575063ffffffff5a1692803b15610352578686926106226040519687958694859463041c592960e51b86526080519960048701615a41565b61a70e87036113aa5760008051602061601383398151915286036112e65750739cb4706e20a18e59a48ffa7616d700a3891e1861803b156103525760405163cd58657960e01b8152608051600160048301526001600160a01b0386166024830152604482018590526000606483018190526084830181905260c060a484015260c48301529091829081868160e48101610878565b6040519063095ea7b360e01b8252739cb4706e20a18e59a48ffa7616d700a3891e18619182600482015284602482015283816044816080518c5af180156104fb5761138d575b50813b156103525760405163cd58657960e01b8152608051600160048301526001600160a01b03808816602484015260448301879052909216606482015260006084820181905260c060a483015260c4820152918290818060e48101610622565b6113a390843d861161068b5761067c8183615059565b508861132c565b61210587036114d157600080516020616013833981519152860361144b5750733154cf16ccdb4c6d922629664174b904d80f2c35803b1561035257604051639a2ac6d560e01b81526080516001600160a01b038616600483015260006024830181905260606044840152606483015290918290608490829087905af180156104fb57610650575091600080516020615ff383398151915293916040936104ba565b61145481615b0c565b906040519163095ea7b360e01b8352733154cf16ccdb4c6d922629664174b904d80f2c359283600482015285602482015284816044816080518d5af180156104fb57610a90575063ffffffff5a1692803b15610352578686926106226040519687958694859463041c592960e51b86526080519960048701615a41565b604051639474dee160e01b8152600490fd5b6040516305fd61ad60e01b815260048101879052602490fd5b34610352576020366003190112610352576001600160a01b0361151d614f4c565b16608051526009602052602060406080512054604051908152f35b34610352576020806003193601126103525760043590611556615540565b61155e615470565b3360805152600a9182825260406080512054808210156118d05750336080515282825261159081604060805120614fa9565b5060018060a01b03905460081c1692836080515260099081845260805194604086205493600019968786019586116118a7575033608051528286526115da81604060805120614fa9565b50926003809401549683608051528581526115fa8760406080512061513c565b509761160c6001809a0191825461518f565b9055608051943386528282528061162885604060805120614fa9565b5001549033608051528383528961164486604060805120614fa9565b500154978989116118be57908791895b8b81106117db575050506116729161166d821515615737565b615533565b92336080515282825261168a81604060805120614fa9565b5060ff19908a828254161790553360805152838352896116af83604060805120614fa9565b50918254161790556116c2843387615e95565b5033608051528282526080519960408b208054918b83116116f5575b8b6000805160206160738339815191525560805180f35b82019182116117c4577f0ad30059d4072ae2525ebaaa5b6878f1c50df5f3b921615031c7d0f12dc4a77b9a9b509261175a8897969361175461173f61177398956117b09c98614fa9565b50913360805152858552604060805120614fa9565b9061576f565b33608051525261176e6040608051206157d7565b61518f565b6040805133815260208101969096528501959095526001600160a01b03166060840152608083019390935260a082019290925290819060c0820190565b0390a18082808080808080808080806116de565b634e487b7160e01b8c526011600452608051602490fd5b9091925087608051528185528c60805160408120549182019182116118a75750811161188f57611886906118596118528e8b6080515285895261183160026118288660406080512061513c565b50015489615704565b908c60805152868a526118498560406080512061513c565b50015490615717565b809b615533565b99896080515283875261187f856118758460406080512061513c565b5001918254615533565b90556156f5565b90889291611654565b60249060405190637d58ebb960e01b82526004820152fd5b634e487b7160e01b90526011600452608051602490fd5b60405163374c934360e11b8152600490fd5b604492506040519163abe5c32f60e01b835260048301526024820152fd5b3461035257606036600319011261035257611907614f4c565b61190f6150a4565b6044359063ffffffff908183168084036103525761192b6151f4565b60018060a01b0385169384608051526020936004855260ff6040608051205416611a6f5785608051526004855260805161197485604083209060ff801983541691151516179055565b600986526040812082600854168503918383116118a7575092826119fe611a6096937fc8c34f23fafb34e68119c1d231ef03d0d47225b15e2c4de3efbefa14b0181d869a9b96611a5b96604051926119cb84614fdb565b168252858b8301528c60408301526080516060830152608051608083015260805160a083015260805160c083015261584d565b896080515260098852611a1c60406080512092826008541690615835565b60405193611a2985614fdb565b845216878301528860408301526080516060830152608051608083015260805160a083015260805160c083015261584d565b615a85565b6040519015158152a260805180f35b60405163411befff60e11b8152600481018790528415156024820152604490fd5b34610352576080518060031936011261037d576020600254604051908152f35b34610352576020366003190112610352576001600160a01b03611ad1614f4c565b16608051526005602052602060406080512054604051908152f35b34610352576020366003190112610352576001600160a01b03611b0d614f4c565b16608051526004602052602060ff60406080512054166040519015158152f35b3461035257604036600319011261035257611b466150a4565b611b4e6151f4565b600435608051526003602052611b7660805191604083209060ff801983541691151516179055565b80f35b3461035257604036600319011261035257611bc7600435611b98614f67565b908060805152600080516020616033833981519152602052611bc26001604060805120015461524c565b6153ac565b5060805180f35b3461035257608036600319011261035257600435602435611bed614f7d565b91468103611eb15781608051526020926003845260ff604060805120541615611e98576001600160a01b039081611c22615da0565b166040516370a0823160e01b8082523060048301528782602481865afa9182156104fb5760805192611e69575b506040516323b872dd60e01b81523360048201523060248201526064803560448301526080519192918a9184918290885af19182156104fb578992611e4c575b5060246040518095819382523060048301525afa80156104fb5760805190611e1d575b611cbc925061518f565b600154808210611e00575073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc290816080515260058752604060805120611cf7828254615533565b9055611d14620f4240611d0c60025484615704565b04809261518f565b916080515260068752604060805120611d2e828254615533565b90556080518481541690813b1561037d575060405163305f478560e21b81526080805160048301899052602483018a90526001600160a01b038716604484015260648301869052608483019490945251909291839160a4918391905af180156104fb57611deb575b5090611ddf7fcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a89939260405193849316963396846040919493926060820195825260208201520152565b0390a360405160018152f35b611df49061500d565b6080518015611d965780fd5b604491604051916358f8331360e11b835260048301526024820152fd5b508682813d8311611e45575b611e338183615059565b8101031261035257611cbc9151611cb2565b503d611e29565b611e6290833d851161068b5761067c8183615059565b5089611c8f565b9091508781813d8311611e91575b611e818183615059565b8101031261035257519088611c4f565b503d611e77565b6040516363b5c9db60e01b815260048101849052602490fd5b604051631a26aa4d60e21b8152600490fd5b36600319011261035257602435600435611edb614f7d565b6001600160a01b039060643560c4358381169081900361035257611f006103b36151b2565b468603611eb1578460805152600360205260ff60406080512054161561204357803b15610352576040519063534a7e1d60e11b8252346004830152816024816080519334905af180156104fb57612034575b50600080516020616013833981519152608051526005602052604060805120611f7c82825461518f565b90556080518381541690813b1561037d5750604051631eb65ffb60e01b8152608051909182908180611fba60a435886084358b8f8f600488016158fb565b0391608051905af180156104fb5761201f575b50906120167f61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f32939260405193849316963096846040919493926060820195825260208201520152565b0390a360805180f35b6120289061500d565b6080518015611fcd5780fd5b61203d9061500d565b85611f52565b604051632ef6974160e11b815260048101869052602490fd5b34610352576020366003190112610352576001600160a01b0361207d614f4c565b1660805152600a602052602060406080512054604051908152f35b34610352576040366003190112610352576120b1614f4c565b7ff54d3b756d286b6b08e5d4eda6dfe5b135664abf029e58e637cbf013c442c9506020602435926120e06151f4565b60018060a01b0316928360805152600b82528060406080512055604051908152a260805180f35b346103525760208060031936011261035257600435612124615540565b61212c615470565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2918260805152600b81526080516040812054831061230757506040516323b872dd60e01b815233600482015230602482015260448101839052818180606481010381608051885af180156104fb576122ea575b508260805152600981526080519260408420549360001985019485116118a7575080608051526009825260ff60046121d18660406080512061513c565b500154166122d1577fc138e3bd6d13eefa5cb01c0d35c5794001141efaf4e5ad888cad059935f83383929381608051526009835263ffffffff6122198260406080512061513c565b5054164210612260575b50608051526005815260406080512061223d848254615533565b90556040519283523392a260016000805160206160738339815191525560805180f35b6122bc6001913360805152600a85526122a8604060805120604051906122858261503d565b60805182528688830152836040830152608051606083015289608083015261549b565b83608051526009855260406080512061513c565b50016122c9858254615533565b905584612223565b6040516311cf1b0760e31b815260048101859052602490fd5b61230090823d841161068b5761067c8183615059565b5083612194565b8281604492526040608051205490604051916327500c6d60e21b835260048301526024820152fd5b346103525760203660031901126103525761234b6103b36151b2565b60043560015560805180f35b34610352576080518060031936011261037d576020612374615da0565b6040516001600160a01b039091168152f35b346103525760403660031901126103525761239f614f4c565b6123a76150a4565b6123af6151f4565b608080516001600160a01b038416905260046020525160409020805460ff191660ff831515161790556123e3575b60805180f35b6123ec90615a85565b806123dd565b34610352576080518060031936011261037d5760075461241181615d06565b906080515b81811061242f576040518061242b85826150e9565b0390f35b60805191825b3360805152600a60209080825260406080512080548410156124f7578361245b91614fa9565b50546001600160a01b03929060081c831661247587615158565b949054600395861b1c1614612496575b505050612491906156f5565b612435565b959295336080515281815260ff6124b288604060805120614fa9565b5054166124ea579161249193916124e1933360805152526124d887604060805120614fa9565b50015490615533565b93908680612485565b50505092612491906156f5565b50505050919061254d9161250a82615158565b919054916040519261251b84615021565b60039190911b1c6001600160a01b03168252602082015261253c8286615821565b526125478185615821565b506156f5565b612416565b34610352576080518060031936011261037d57602090604051908152f35b60603660031901126103525760043560243561258a614f7d565b91468103611eb1578160805152600360205260ff604060805120541615612713576001548034106126f55750600080516020616013833981519152806080515260056020526040608051206125e0348254615533565b9055620f42406125f260025434615704565b04936125fe853461518f565b91608051526006602052604060805120612619868254615533565b905560805180546001600160a01b03939190841690813b1561037d575060405163305f478560e21b815260808051466004840152602483018990526001600160a01b03861660448401526064830185905260848301999099525196979596909591869160a4918391905af19081156104fb577f0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc5946126d8926126e6575b5060405193849316963396846040919493926060820195825260208201520152565b0390a3602060405160018152f35b6126ef9061500d565b876126b6565b604490604051906358f8331360e11b82526004820152346024820152fd5b604051632ef6974160e11b815260048101839052602490fd5b604036600319011261035257612740614f4c565b60249081359061274e615470565b341561290157505061275e615540565b612766615470565b6000805160206160138339815191528060805152602091600b8352604060805120548034106128e457508160805152600983526080519060408220546000198101928184116128cf575083608051526009855260406080512054156128b85783608051526009855263ffffffff6127e28460406080512061513c565b5054164210156128a35750506128366001913360805152600a85526122a8604060805120604051906128138261503d565b60805182528688830152836040830152608051606083015234608083015261549b565b5001612843348254615533565b9055608051526005815260406080512061285e348254615533565b90557fe7466ea83435490635fc76a5f33da4815758ab48b1d45858f0452ca646556937604051913483523392a260016000805160206160738339815191525560805180f35b60405190637d58ebb960e01b82526004820152fd5b50604051637d58ebb960e01b815260016004820152fd5b634e487b7160e01b9052506011600452608051fd5b604491604051916327500c6d60e21b835260048301523490820152fd5b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2906001600160a01b0316818103612b035750612930615540565b612938615470565b8060805152602090600b825260805160408120548410612adc57506040516323b872dd60e01b815233600482015230602482015260448101849052828180606481010381608051865af180156104fb57612abf575b508060805152600982526080516040812054906000198201918211612aa9575081608051526009835260ff60046129c98360406080512061513c565b50015416612a92577fc138e3bd6d13eefa5cb01c0d35c5794001141efaf4e5ad888cad059935f8338393945081608051526009835263ffffffff612a128260406080512061513c565b5054164210612a58575b506080515260058152604060805120612a36848254615533565b90556040519283523392a26001600080516020616073833981519152556123dd565b612a7d6001913360805152600a85526122a8604060805120604051906122858261503d565b5001612a8a858254615533565b905584612a1c565b8490604051906311cf1b0760e31b82526004820152fd5b634e487b7160e01b905260116004526080518590fd5b612ad590833d851161068b5761067c8183615059565b508461298d565b848482604493526040608051205491604051926327500c6d60e21b84526004840152820152fd5b80939150608051526020906004825260ff6040608051205416612b2a575b505050506123dd565b612b32615540565b612b3a615470565b83608051526004825260ff604060805120541615612cd1578360805152600b825260406080512054808410612cb55750612b7683303387615570565b836080515260098252608051906040822054916000198301928311612ca1575084608051526009835263ffffffff612bb38360406080512061513c565b505416421015612c8b5750906001612c3c7f3ac7d3823c11677fba9479ed26f696a3f17e16a5a5c39162fa6d183905aa6735933360805152600a8452612c2860406080512060405190612c058261503d565b60805182528987830152836040830152608051606083015288608083015261549b565b86608051526009845260406080512061513c565b5001612c49848254615533565b9055836080515260058152604060805120612c65848254615533565b90556040519283523392a360016000805160206160738339815191525580808080612b21565b9060405190637d58ebb960e01b82526004820152fd5b634e487b7160e01b90526011600452608051fd5b83604492604051926327500c6d60e21b84526004840152820152fd5b6040516305fd61ad60e01b815260048101859052fd5b6080518060031936011261037d57612cfd615540565b612d05615470565b6000805160206160138339815191528060805152602090600b825260406080512054803410612dc9575080608051526009825260805160408120546000198101918183116118a757508260805152600984526040608051205415612db05782608051526009845263ffffffff612d808360406080512061513c565b50541642101561188f57506128366001913360805152600a85526122a8604060805120604051906128138261503d565b604051637d58ebb960e01b815260016004820152602490fd5b604490604051906327500c6d60e21b82526004820152346024820152fd5b346103525736600319011261035257602435600435612e04614f7d565b612e0c614f93565b60843590612e1b6103b36151b2565b468503611eb1578360805152600360205260ff604060805120541615612f575760018060a01b03809116928360805152600460205260ff604060805120541615612f3e57612e6a838286615d67565b83608051526005602052604060805120612e8584825461518f565b9055608051928284541693843b1561037d575060405193631eb65ffb60e01b855284608051918180612ec58c878d8a60c4359360a43592600488016158fb565b0391608051905af19081156104fb577f0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe94612f2192612f2f575b5060405193849316973097846040919493926060820195825260208201520152565b0390a4602060405160018152f35b612f389061500d565b88612eff565b6040516305fd61ad60e01b815260048101859052602490fd5b604051632ef6974160e11b815260048101859052602490fd5b3461035257604036600319011261035257612f89614f67565b60043560805152600080516020616033833981519152602052608051604081209160018060a01b03169052602052602060ff60406080512054166040519015158152f35b34610352576080518060031936011261037d57612fe86151f4565b612ff0615470565b600080516020616053833981519152600160ff198254161790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a160805180f35b3461035257602036600319011261035257600435608051526003602052602060ff60406080512054166040519015158152f35b3461035257602080600319360112610352576001600160a01b0361308f614f4c565b1660805152600a81526040608051209081546130aa8161507b565b906130b86040519283615059565b80825282820180946080515260805184812090915b838310613152578585886040519183830190848452518091526040830191936080515b8281106130fd5784840385f35b9091928260a0826131436001948a516080809180511515845260018060a01b03602082015116602085015260408101516040850152606081015160608501520151910152565b019601910194929190946130f0565b60048660019261316185615ac2565b8152019201920191906130cd565b34610352576020366003190112610352576131886151f4565b60043560025560805180f35b3461035257602036600319011261035257600435600754811015610352576131bd602091615158565b905460405160039290921b1c6001600160a01b03168152f35b34610352576080518060031936011261037d57602060ff60008051602061605383398151915254166040519015158152f35b3461035257604036600319011261035257613221614f4c565b60405161322d81614fdb565b6080518152608051602082015260805160408201526080516060820152608051608082015260805160a082015260c060805191015260018060a01b0380911660805152600960205261328660243560406080512061513c565b506040519061329482614fdb565b805463ffffffff808216948585526020850192828160201c16845281604087019160401c168152600185015492606087019384526002860154946080880195865260c060ff600460038a01549960a08c019a8b520154169801971515885260405198895251166020880152511660408601525160608501525160808401525160a083015251151560c0820152f35b34610352576080518060031936011261037d5761333d615540565b613345615470565b6080515b60075481101561376a5761335c81615158565b90549060018060a01b03828260031b1c166080515260096020526040608051205415613750576080515b3360805152600a602052604060805120805482101561373e57816133a991614fa9565b505460081c6001600160a01b03908116600384901b85901c909116146133f8575b6001600160ff1b0381146133e057600101613386565b634e487b7160e01b6080515260116004526024608051fd5b3360805152600a60205260ff61341382604060805120614fa9565b5054166137395760808051600384901b85901c6001600160a01b03169052600960205251604081205490600019820182106118a757503360805152600a602052600361346483604060805120614fa9565b50015460018060a01b03858560031b1c166080515260096020526134a06001613496600019850160406080512061513c565b500191825461518f565b9055608051338152600a60205260036134be84604060805120614fa9565b500154913360805152600a60205260016134dd85604060805120614fa9565b50015492600019820184116118be5782845b60001984018110613661575061350a9161166d821515615737565b923360805152600a60205261352485604060805120614fa9565b50805460ff1916600117905561354b8433600389901b8a901c6001600160a01b0316615e95565b503360805152600a6020526080516040812080549160018311613575575b505050505050506133ca565b60009793949596971983019283116118a757506135af9161359591614fa9565b503360805152600a60205261175483604060805120614fa9565b3360805152600a6020526135c76040608051206157d7565b600160ff1b81146133e057836136067f0ad30059d4072ae2525ebaaa5b6878f1c50df5f3b921615031c7d0f12dc4a77b9561365293600019019761518f565b60408051338152602081019590955260001990950194840194909452600387901b88901c6001600160a01b03166060840152608083019390935260a082019290925290819060c0820190565b0390a184808080808080613569565b6080805160038a901b8b901c6001600160a01b03169052600960205251604081205491925060001982019182116118a75750811161188f576137329060018060a01b03898960031b1c166080515260096020526137066136ff6136d760026136ce8560406080512061513c565b50015486615704565b60018060a01b038c8c60031b1c1660805152600960205260016118498560406080512061513c565b8096615533565b9460018060a01b038a8a60031b1c1660805152600960205261187f60036118758460406080512061513c565b83906134ef565b6133ca565b5050505061374b906156f5565b613349565b604051637d58ebb960e01b81526080516004820152602490fd5b60016000805160206160738339815191525560805180f35b3461035257613790366150b3565b90949261379e6103b36151b2565b468503611eb15783608051526020956003875260ff604060805120541615612043576001600160a01b0392836137d2615da0565b60405163a9059cbb60e01b81526001600160a01b03861660048201526024810188905260805190928b9284926044928492165af180156104fb576138e9575b5073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2608051526005885260406080512061384086825461518f565b9055608051908482541691823b1561037d5750918691848961387e89966040519a8b9687958695631eb65ffb60e01b87526080519a600488016158fb565b0391608051905af19081156104fb577f7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc1494611ddf926138da575b5060405193849316963096846040919493926060820195825260208201520152565b6138e39061500d565b886138b8565b6138ff90893d8b1161068b5761067c8183615059565b5088613811565b34610352576020366003190112610352576001600160a01b03613927614f4c565b16608051526006602052602060406080512054604051908152f35b346103525760403660031901126103525761395b614f4c565b60243590613967615540565b61396f615470565b60018060a01b031690816080515260206004815260ff604060805120541615613aad578260805152600b815260406080512054808310613a8f57506139b682303386615570565b82608051526009815260805160408120549060001982019182116118a7575083608051526009825263ffffffff6139f28260406080512061513c565b50541642101561188f57906001613a437f3ac7d3823c11677fba9479ed26f696a3f17e16a5a5c39162fa6d183905aa6735933360805152600a8452612c2860406080512060405190612c058261503d565b5001613a50848254615533565b9055836080515260058152604060805120613a6c848254615533565b90556040519283523392a360016000805160206160738339815191525560805180f35b60449083604051916327500c6d60e21b835260048301526024820152fd5b6040516305fd61ad60e01b815260048101849052602490fd5b3461035257604036600319011261035257613adf614f4c565b613ae7614f67565b907ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0091825460ff8160401c16159267ffffffffffffffff821680159081613c8e575b6001149081613c84575b159081613c7b575b50613c695767ffffffffffffffff1982166001178555613bcc9184613c4a575b50613b6461542f565b613b6c61542f565b613b7461542f565b600160008051602061607383398151915255613b8e61542f565b613b9661542f565b600080516020616053833981519152805460ff19169055613bb561542f565b67016345785d8a000060015561271060025561527f565b5060805180546001600160a01b0319166001600160a01b039092169190911790556008805463ffffffff1916621baf80179055613c095760805180f35b68ff00000000000000001981541690557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1806123dd565b68ffffffffffffffffff19166801000000000000000117855585613b5b565b60405163f92ee8a960e01b8152600490fd5b90501586613b3b565b303b159150613b33565b859150613b29565b3461035257604036600319011261035257613caf614f4c565b608080516001600160a01b039283169052600960205251604081208054602435929083101561037d575090613ce39161513c565b5090815491600181015460028201549060ff6004600385015494015416936040519563ffffffff80821688528160201c16602088015260401c1660408601526060850152608084015260a0830152151560c0820152f35b34610352576080518060031936011261037d57613d556151f4565b600080516020616053833981519152805460ff811615613da55760ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a160805180f35b604051638dfc202b60e01b8152600490fd5b34610352576080518060031936011261037d57600754613dd681615d06565b6080515b828110613def576040518061242b84826150e9565b9060805192835b3360805152600a806020526040608051208054831015613ff35782613e1a91614fa9565b505460081c6001600160a01b0316613e3186615158565b905460039190911b1c6001600160a01b031614613e58575b50613e53906156f5565b613df6565b33608051528060205260ff613e7283604060805120614fa9565b505416613e4957613e8285615158565b60018060a01b0391549060031b1c1660805152600960205260805190604082205491826000198101116118a757503360805152806020526003613eca84604060805120614fa9565b5001549033608051526020526001613ee784604060805120614fa9565b50015460001983018111613fe75791905b60001982018310613f0a575050613e49565b909196613f1687615158565b60018060a01b0391549060031b1c1660805152600960205260805160408120549060001982019182116118a757508811613fce57613fc1613fc791613f5a89615158565b60018060a01b0391549060031b1c16608051526009602052613fbb613f8960026136ce8d60406080512061513c565b613f928b615158565b60018060a01b0391549060031b1c1660805152600960205260016118498d60406080512061513c565b90615533565b976156f5565b9190613ef8565b604051637d58ebb960e01b815260048101899052602490fd5b505050613e53906156f5565b50505091614005919361250a82615158565b919091613dda565b346103525760a03660031901126103525760043560243561402c614f7d565b91614035614f93565b92608435468303611eb15783608051526020946003865260ff6040608051205416156120435760018060a01b038091169283608051526004875260ff604060805120541615612f3e576040516370a0823160e01b808252306004830152908881602481895afa9081156104fb5760805191614232575b506140b885303389615570565b6040519182523060048301528882602481895afa80156104fb5760805190614203575b6140e5925061518f565b92846080515260058852614100604060805120918254615533565b905561411d620f424061411560025486615704565b04809461518f565b846080515260068852604060805120614137858254615533565b90556080518381541690813b1561037d575060405163305f478560e21b815260808051600483018a9052602483018b90526001600160a01b038616604484015260648301859052608483019790975251909591869160a4918391905af19081156104fb577fece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a4946141e8926141f4575b5060405193849316973397846040919493926060820195825260208201520152565b0390a460405160018152f35b6141fd9061500d565b896141c6565b508882813d831161422b575b6142198183615059565b81010312610352576140e591516140db565b503d61420f565b90508881813d8311614258575b6142498183615059565b810103126103525751896140ab565b503d61423f565b614268366150b3565b929493906142776103b36151b2565b468203611eb1578460805152600360205260ff604060805120541615612043576080516001600160a01b0387811697909291808981888015614387575b83928392f1156104fb576000805160206160138339815191526080515260056020526040608051206142e786825461518f565b90556080519283541692833b1561037d5750918491879361432360405198899586948594631eb65ffb60e01b86528a6080519a600488016158fb565b0391608051905af19283156104fb577f61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f3293614378575b50604080519485526020850191909152830152309180606081016126d8565b6143819061500d565b85614359565b6108fc91506142b4565b34610352576040366003190112610352576143aa614f67565b336001600160a01b038216036143c657611bc7906004356153ac565b60405163334bd91960e11b8152600490fd5b346103525760208060031936011261035257600435906143f6615540565b6143fe615470565b3360805152600a80825260406080512054808410156146585750336080515280825261442f83604060805120614fa9565b5060018060a01b03905460081c1691826080515260099081815260805193604085205493600019958686019586116118a75750336080515280835261447987604060805120614fa9565b50926003809401549683608051528582526144998760406080512061513c565b50976144ab6001809a0191825461518f565b905560805194338652838352806144c78b604060805120614fa9565b500154913360805152848452896144e38c604060805120614fa9565b500154978989116118be57885b8a81106145cd57505050509661456e8199869594899461453a887f7e6fedb6f824b2112ac0626ced32f6be6324d554cc03efd112aa021ee4d740ac9d61166d6145b49c1515615737565b50336080515281815261455283604060805120614fa9565b508560ff19825416179055336080515252604060805120614fa9565b50015561457c823383615e95565b506040805133815260208101959095528401949094526001600160a01b039093166060830152608082019290925290819060a0820190565b0390a16000805160206160738339815191525560805180f35b87608051528186526080516040812054908482019182116118a75750811161188f57614653906146376118528e8b60805152858a5261461f60026146168660406080512061513c565b5001548a615704565b908c60805152868b526118498560406080512061513c565b99896080515283885261187f866118758460406080512061513c565b6144f0565b836044916040519163abe5c32f60e01b835260048301526024820152fd5b60209081600319360112610352576004359067ffffffffffffffff8211610352573660238301121561035257808260040135926146b28461507b565b936146c06040519586615059565b80855260249302810183019085850190368311610352578401905b8282106148b757505050506146f16103b36151b2565b6080515b82518110156123dd576001600160a01b0360406147128386615821565b51015116806080515260098086526080519060408220546000198101928184116148a157508360805152818852608051906040822090600119019184831161488b57509261483e7fb6f449f07ceaf55392c9899e0797c6529908ae827c2d498c682e90d42c2411679383600461478e8d9661484a9b9a9961513c565b5001906001918260ff198254161790558061484f575b5086608051528185526147bc8460406080512061513c565b509163ffffffff809354871c16928860805152818752604060805120926147e7826008541686615835565b928a6080515288526147fe8760406080512061513c565b500154916040519461480f86614fdb565b855216868401528760408401526060830152608051608083015260805160a083015260805160c083015261584d565b604051908152a26156f5565b6146f5565b876080515260069081875260026148736080519260408085205494888c522061513c565b5001558760805152855260805160408120558c6147a4565b634e487b7160e01b905260116004526080518790fd5b634e487b7160e01b905260116004526080518690fd5b838236031261035257604051906148cd82614fdb565b6148d683615093565b82526148e3888401615093565b8289015260408301356001600160a01b0381168103610352576040830152606080840135908301526080830135608083015260a0808401359083015260c090818401359283151584036103525786938a938201528152019101906146db565b3461035257606036600319011261035257611bc761495e614f4c565b614966614f67565b6149716103b36151b2565b60443591615e95565b3461035257604036600319011261035257611bc7600435614999614f67565b9080608051526000805160206160338339815191526020526149c36001604060805120015461524c565b61531f565b34610352576080518060031936011261037d576149e3615540565b6149eb615470565b6080515b60075481101561376a57614a0281615158565b9190549160018060a01b03838260031b1c166080515260096020526040608051205415613750576080515b3360805152600a6020526040608051208054821015614d285781614a5091614fa9565b505460081c6001600160a01b03908116600384901b86901c90911614614a87575b6001600160ff1b0381146133e057600101614a2d565b9291903360805152600a60205260ff614aa585604060805120614fa9565b505416614d205760808051600383901b85901c6001600160a01b0316905260096020525160408120549490600019860186106118a757503360805152600a6020526003614af782604060805120614fa9565b50015460018060a01b03858460031b1c16608051526009602052614b296001613496600019890160406080512061513c565b905560805194338652600a6020526003614b4883604060805120614fa9565b500154953360805152600a6020526001614b6784604060805120614fa9565b50015490600019830182116118be57815b60001984018110614c4c575080614bbd817f7e6fedb6f824b2112ac0626ced32f6be6324d554cc03efd112aa021ee4d740ac969798999a61166d614c44951515615737565b503360805152600a602052614bd786604060805120614fa9565b50600019850160019190910155614bff813360038a901b8c901c6001600160a01b0316615e95565b5060408051338152602081019490945260001990940193830193909352600386901b88901c6001600160a01b03166060830152608082019290925290819060a0820190565b0390a1614a71565b60808051600388901b8a901c6001600160a01b0316905260096020525160408120546000198101919082116118a75750811161188f57614d1b9060018060a01b03898860031b1c16608051526009602052614cef614ce8614cc06002614cb78560406080512061513c565b5001548d615704565b60018060a01b038c8b60031b1c1660805152600960205260016118498560406080512061513c565b8094615533565b9260018060a01b038a8960031b1c1660805152600960205261187f60036118758460406080512061513c565b614b78565b909192614a71565b505050614d3591506156f5565b6149ef565b346103525760203660031901126103525760043560805152600080516020616033833981519152602052602060016040608051200154604051908152f35b34610352576080518060031936011261037d57602063ffffffff60085416604051908152f35b34610352576020366003190112610352576001600160a01b03614dbf614f4c565b1660805152600b602052602060406080512054604051908152f35b3461035257604036600319011261035257614df3614f4c565b608080516001600160a01b039283169052600a6020525160408120805460243593929084101561037d575060a092614e2a91614fa9565b5090815491600181015460036002830154920154926040519460ff81161515865260081c166020850152604084015260608301526080820152f35b34610352576080518060031936011261037d57546040516001600160a01b039091168152602090f35b3461035257604036600319011261035257614ea7614f4c565b614eb26103b36151b2565b60018060a01b0316608051526005602052608051602435604082205580f35b34610352576080518060031936011261037d576020600154604051908152f35b34614f48576020366003190112614f485760043563ffffffff60e01b8116809103614f445760209250637965db0b60e01b8114908115614f33575b5015158152f35b6301ffc9a760e01b14905083614f2c565b8280fd5b5080fd5b600435906001600160a01b0382168203614f6257565b600080fd5b602435906001600160a01b0382168203614f6257565b604435906001600160a01b0382168203614f6257565b606435906001600160a01b0382168203614f6257565b8054821015614fc55760005260206000209060021b0190600090565b634e487b7160e01b600052603260045260246000fd5b60e0810190811067ffffffffffffffff821117614ff757604052565b634e487b7160e01b600052604160045260246000fd5b67ffffffffffffffff8111614ff757604052565b6040810190811067ffffffffffffffff821117614ff757604052565b60a0810190811067ffffffffffffffff821117614ff757604052565b90601f8019910116810190811067ffffffffffffffff821117614ff757604052565b67ffffffffffffffff8111614ff75760051b60200190565b359063ffffffff82168203614f6257565b602435908115158203614f6257565b60c0906003190112614f625760043590602435906044356001600160a01b0381168103614f625790606435906084359060a43590565b60208082019080835283518092528060408094019401926000905b83821061511357505050505090565b845180516001600160a01b03168752830151868401529485019493820193600190910190615104565b8054821015614fc5576000526005602060002091020190600090565b600754811015614fc55760076000527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880190600090565b9190820391821161519c57565b634e487b7160e01b600052601160045260246000fd5b604051602081017f33fe247d78fee2e7fd135c405eda4bd2911c0a73c0a81b36c3bcc967dd06e5ad8152602082526151e982615021565b9051902060ff191690565b3360009081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604081205460ff161561522e5750565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b8060005260008051602061603383398151915260205260406000203360005260205260ff604060002054161561522e5750565b6001600160a01b031660008181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260408120549091906000805160206160338339815191529060ff1661531a578280526020526040822081835260205260408220600160ff1982541617905533917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b505090565b9060009180835260008051602061603383398151915280602052604084209260018060a01b03169283855260205260ff604085205416156000146153a6578184526020526040832082845260205260408320600160ff198254161790557f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d339380a4600190565b50505090565b9060009180835260008051602061603383398151915280602052604084209260018060a01b03169283855260205260ff6040852054166000146153a657818452602052604083208284526020526040832060ff1981541690557ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b339380a4600190565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c161561545e57565b604051631afcd79f60e31b8152600490fd5b60ff600080516020616053833981519152541661548957565b60405163d93c066560e01b8152600490fd5b8054600160401b811015614ff7576154b891600182018155614fa9565b91909161551d576080816154dd600393511515859060ff801983541691151516179055565b60208101518454610100600160a81b03191660089190911b610100600160a81b031617845560408101516001850155606081015160028501550151910155565b634e487b7160e01b600052600060045260246000fd5b9190820180921161519c57565b600080516020616073833981519152600281541461555e5760029055565b604051633ee5aeb560e01b8152600490fd5b6040516323b872dd60e01b60208201526001600160a01b03928316602482015292909116604483015260648201929092526155c1916155bc82608481015b03601f198101845283615059565b6155db565b565b90816020910312614f6257518015158103614f625790565b6000806156259260018060a01b03169360208151910182865af13d1561566e573d9061560682615676565b916156146040519384615059565b82523d6000602084013e5b83615692565b8051908115159182615653575b505061563b5750565b60249060405190635274afe760e01b82526004820152fd5b61566692506020809183010191016155c3565b153880615632565b60609061561f565b67ffffffffffffffff8111614ff757601f01601f191660200190565b906156b957508051156156a757805190602001fd5b604051630a12f52160e11b8152600490fd5b815115806156ec575b6156ca575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b156156c2565b600019811461519c5760010190565b8181029291811591840414171561519c57565b8115615721570490565b634e487b7160e01b600052601260045260246000fd5b1561573e57565b60405162461bcd60e51b8152602060048201526009602482015268139bc814995dd85c9960ba1b6044820152606490fd5b919061551d57808203615780575050565b60038161579e60ff83945416859060ff801983541691151516179055565b80548454610100600160a81b031916610100600160a81b0390911617845560018101546001850155600281015460028501550154910155565b8054801561580b5760001901906157ee8282614fa9565b61551d576003816000809355826001820155826002820155015555565b634e487b7160e01b600052603160045260246000fd5b8051821015614fc55760209160051b010190565b91909163ffffffff8080941691160191821161519c57565b8054600160401b811015614ff75761586a9160018201815561513c565b61551d578151815460208085015160408087015168010000000000000000600160e01b03911b1667ffffffff000000009190921b1663ffffffff9093166001600160e01b03199092169190911791909117178155606082015160018201556080820151600282015560a0820151600382015560c0909101516004909101805460ff191691151560ff16919091179055565b9260a094919796959260c085019885526020850152600180861b03166040840152606083015260808201520152565b91906020928381830312614f625780519067ffffffffffffffff8211614f6257019281601f85011215614f6257835161596281615676565b926159706040519485615059565b818452828287010111614f625760005b81811061599557508260009394955001015290565b8581018301518482018401528201615980565b9261010094929160018060a01b0392838092168652166020850152166040830152606082015260006080820152600060a082015260e060c0820152600060e08201520190565b60008051602061601383398151915281526001600160a01b03918216602082015291166040820152606081019190915260006080820181905260a0820181905260e060c083018190528201526101000190565b93909163ffffffff9360e0969360018060a01b0392838092168852166020870152166040850152606084015216608082015260c060a0820152600060c08201520190565b600754600160401b811015614ff757806001615aa49201600755615158565b819291549060031b9160018060a01b03809116831b921b1916179055565b90604051615acf8161503d565b825460ff81161515825260081c6001600160a01b031660208201526001830154604082015260028301546060820152600392909201546080830152565b6001600160a01b031673c02aaa39b223fe8d0a0e5c4f27ead9083c756cc28103615b3c57506006602160991b0190565b73dac17f958d2ee523a2206206994597c13d831ec78103615b7057507394b008aa00579c1307b0ef2c499ad98a8ce58e5890565b73a0b86991c6218b36c1d19d4a2e9eb0ce3606eb488103615ba45750730b2c639c533813f4aa9d7837caf62653d097ff8590565b736b175474e89094c44da98b954eedeac495271d0f8103615bd8575073da10009cbd5d07dd0cecc66161fc93d7c9000da190565b602490604051906305fd61ad60e01b82526004820152fd5b6001600160a01b031673dac17f958d2ee523a2206206994597c13d831ec78103615c2d575073201eba5cc46d216ce6dc03f6a759e8e766e956ae90565b73a0b86991c6218b36c1d19d4a2e9eb0ce3606eb488103615bd8575073201eba5cc46d216ce6dc03f6a759e8e766e956ae90565b6001600160a01b031673dac17f958d2ee523a2206206994597c13d831ec78103615c9e575073f417f5a458ec102b90352f697d6e2ac3a3d2851f90565b73a0b86991c6218b36c1d19d4a2e9eb0ce3606eb488103615cd2575073b73603c5d87fa094b7314c74ace2e64d165016fb90565b736b175474e89094c44da98b954eedeac495271d0f8103615bd85750731c466b9371f8aba0d7c458be10a62192fcb8aa7190565b90615d108261507b565b604090615d1f82519182615059565b8381528093615d30601f199161507b565b0191600090815b848110615d45575050505050565b6020908251615d5381615021565b848152828581830152828701015201615d37565b60405163a9059cbb60e01b60208201526001600160a01b03909216602483015260448201929092526155c1916155bc82606481016155ae565b466208275003615db5576004605360981b0190565b4661044d03615dd657734f9a0e7fd2bf6067db6994cf12e4495df938e6e990565b46600a03615de9576006602160991b0190565b4661a4b103615e0a577382af49447d8a07e3bd95bd0d56f35241523fbab190565b4661a4ba03615e2b5773722e8bdd2ce80a4422e880164f2079488e11536590565b4661014403615e4c57735aea5775959fbc2557cc8789bc1bf90a239d9a9190565b4661138803615e675760405163e31d668960e01b8152600490fd5b4660a903615e81576040516322c20c6960e11b8152600490fd5b46612105036114d1576006602160991b0190565b9160018060a01b03809316916000938385526020936004855260409460ff868820541615615fda5781875260058152858720615ed286825461518f565b90556000805160206160138339815191528203615f36575050824710615f25578480938193829383918315615f1b575b1690f115615f11575050600190565b51903d90823e3d90fd5b6108fc9250615f02565b8351632c1d501360e11b8152600490fd5b915091929395948651906370a0823160e01b82523060048301528282602481875afa928315615fce5790869392918193615f9a575b50505010615f8357615f7e939450615d67565b600190565b6024908551906311745b6160e21b82526004820152fd5b919350915082813d8111615fc7575b615fb38183615059565b8101031261037d5750839051388080615f6b565b503d615fa9565b508751903d90823e3d90fd5b85516305fd61ad60e01b815260048101839052602490fdfe08ffba656e4ac665e1aa4bb6a342e1a69d4cc12fd751ac862fa3fe27b0c7524a000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a2646970667358221220bfbb18a8c5406cbf3c8034fb55553cd5ded2f4ed765e7999179b8eb2881ade0364736f6c63430008140033",
}

// L1PoolManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use L1PoolManagerMetaData.ABI instead.
var L1PoolManagerABI = L1PoolManagerMetaData.ABI

// L1PoolManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1PoolManagerMetaData.Bin instead.
var L1PoolManagerBin = L1PoolManagerMetaData.Bin

// DeployL1PoolManager deploys a new Ethereum contract, binding an instance of L1PoolManager to it.
func DeployL1PoolManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L1PoolManager, error) {
	parsed, err := L1PoolManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1PoolManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1PoolManager{L1PoolManagerCaller: L1PoolManagerCaller{contract: contract}, L1PoolManagerTransactor: L1PoolManagerTransactor{contract: contract}, L1PoolManagerFilterer: L1PoolManagerFilterer{contract: contract}}, nil
}

// L1PoolManager is an auto generated Go binding around an Ethereum contract.
type L1PoolManager struct {
	L1PoolManagerCaller     // Read-only binding to the contract
	L1PoolManagerTransactor // Write-only binding to the contract
	L1PoolManagerFilterer   // Log filterer for contract events
}

// L1PoolManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1PoolManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1PoolManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1PoolManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1PoolManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1PoolManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1PoolManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1PoolManagerSession struct {
	Contract     *L1PoolManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1PoolManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1PoolManagerCallerSession struct {
	Contract *L1PoolManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L1PoolManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1PoolManagerTransactorSession struct {
	Contract     *L1PoolManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L1PoolManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1PoolManagerRaw struct {
	Contract *L1PoolManager // Generic contract binding to access the raw methods on
}

// L1PoolManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1PoolManagerCallerRaw struct {
	Contract *L1PoolManagerCaller // Generic read-only contract binding to access the raw methods on
}

// L1PoolManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1PoolManagerTransactorRaw struct {
	Contract *L1PoolManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1PoolManager creates a new instance of L1PoolManager, bound to a specific deployed contract.
func NewL1PoolManager(address common.Address, backend bind.ContractBackend) (*L1PoolManager, error) {
	contract, err := bindL1PoolManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1PoolManager{L1PoolManagerCaller: L1PoolManagerCaller{contract: contract}, L1PoolManagerTransactor: L1PoolManagerTransactor{contract: contract}, L1PoolManagerFilterer: L1PoolManagerFilterer{contract: contract}}, nil
}

// NewL1PoolManagerCaller creates a new read-only instance of L1PoolManager, bound to a specific deployed contract.
func NewL1PoolManagerCaller(address common.Address, caller bind.ContractCaller) (*L1PoolManagerCaller, error) {
	contract, err := bindL1PoolManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerCaller{contract: contract}, nil
}

// NewL1PoolManagerTransactor creates a new write-only instance of L1PoolManager, bound to a specific deployed contract.
func NewL1PoolManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*L1PoolManagerTransactor, error) {
	contract, err := bindL1PoolManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerTransactor{contract: contract}, nil
}

// NewL1PoolManagerFilterer creates a new log filterer instance of L1PoolManager, bound to a specific deployed contract.
func NewL1PoolManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*L1PoolManagerFilterer, error) {
	contract, err := bindL1PoolManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerFilterer{contract: contract}, nil
}

// bindL1PoolManager binds a generic wrapper to an already deployed contract.
func bindL1PoolManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1PoolManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1PoolManager *L1PoolManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1PoolManager.Contract.L1PoolManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1PoolManager *L1PoolManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PoolManager.Contract.L1PoolManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1PoolManager *L1PoolManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1PoolManager.Contract.L1PoolManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1PoolManager *L1PoolManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1PoolManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1PoolManager *L1PoolManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PoolManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1PoolManager *L1PoolManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1PoolManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L1PoolManager *L1PoolManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L1PoolManager *L1PoolManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _L1PoolManager.Contract.DEFAULTADMINROLE(&_L1PoolManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_L1PoolManager *L1PoolManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _L1PoolManager.Contract.DEFAULTADMINROLE(&_L1PoolManager.CallOpts)
}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) FeePoolValue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "FeePoolValue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) FeePoolValue(arg0 common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.FeePoolValue(&_L1PoolManager.CallOpts, arg0)
}

// FeePoolValue is a free data retrieval call binding the contract method 0x54dc027e.
//
// Solidity: function FeePoolValue(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) FeePoolValue(arg0 common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.FeePoolValue(&_L1PoolManager.CallOpts, arg0)
}

// FundingPoolBalance is a free data retrieval call binding the contract method 0xdbb0481f.
//
// Solidity: function FundingPoolBalance(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) FundingPoolBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "FundingPoolBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingPoolBalance is a free data retrieval call binding the contract method 0xdbb0481f.
//
// Solidity: function FundingPoolBalance(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) FundingPoolBalance(arg0 common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.FundingPoolBalance(&_L1PoolManager.CallOpts, arg0)
}

// FundingPoolBalance is a free data retrieval call binding the contract method 0xdbb0481f.
//
// Solidity: function FundingPoolBalance(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) FundingPoolBalance(arg0 common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.FundingPoolBalance(&_L1PoolManager.CallOpts, arg0)
}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_L1PoolManager *L1PoolManagerCaller) IsSupportChainId(opts *bind.CallOpts, chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "IsSupportChainId", chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_L1PoolManager *L1PoolManagerSession) IsSupportChainId(chainId *big.Int) (bool, error) {
	return _L1PoolManager.Contract.IsSupportChainId(&_L1PoolManager.CallOpts, chainId)
}

// IsSupportChainId is a free data retrieval call binding the contract method 0x72fe6a7e.
//
// Solidity: function IsSupportChainId(uint256 chainId) view returns(bool)
func (_L1PoolManager *L1PoolManagerCallerSession) IsSupportChainId(chainId *big.Int) (bool, error) {
	return _L1PoolManager.Contract.IsSupportChainId(&_L1PoolManager.CallOpts, chainId)
}

// IsSupportToken is a free data retrieval call binding the contract method 0xdac29568.
//
// Solidity: function IsSupportToken(address ) view returns(bool)
func (_L1PoolManager *L1PoolManagerCaller) IsSupportToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "IsSupportToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportToken is a free data retrieval call binding the contract method 0xdac29568.
//
// Solidity: function IsSupportToken(address ) view returns(bool)
func (_L1PoolManager *L1PoolManagerSession) IsSupportToken(arg0 common.Address) (bool, error) {
	return _L1PoolManager.Contract.IsSupportToken(&_L1PoolManager.CallOpts, arg0)
}

// IsSupportToken is a free data retrieval call binding the contract method 0xdac29568.
//
// Solidity: function IsSupportToken(address ) view returns(bool)
func (_L1PoolManager *L1PoolManagerCallerSession) IsSupportToken(arg0 common.Address) (bool, error) {
	return _L1PoolManager.Contract.IsSupportToken(&_L1PoolManager.CallOpts, arg0)
}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_L1PoolManager *L1PoolManagerCaller) L2WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "L2WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_L1PoolManager *L1PoolManagerSession) L2WETH() (common.Address, error) {
	return _L1PoolManager.Contract.L2WETH(&_L1PoolManager.CallOpts)
}

// L2WETH is a free data retrieval call binding the contract method 0xb1887e9b.
//
// Solidity: function L2WETH() view returns(address)
func (_L1PoolManager *L1PoolManagerCallerSession) L2WETH() (common.Address, error) {
	return _L1PoolManager.Contract.L2WETH(&_L1PoolManager.CallOpts)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0x1ca2f173.
//
// Solidity: function MinStakeAmount(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) MinStakeAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "MinStakeAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeAmount is a free data retrieval call binding the contract method 0x1ca2f173.
//
// Solidity: function MinStakeAmount(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) MinStakeAmount(arg0 common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.MinStakeAmount(&_L1PoolManager.CallOpts, arg0)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0x1ca2f173.
//
// Solidity: function MinStakeAmount(address ) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) MinStakeAmount(arg0 common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.MinStakeAmount(&_L1PoolManager.CallOpts, arg0)
}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) MinTransferAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "MinTransferAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) MinTransferAmount() (*big.Int, error) {
	return _L1PoolManager.Contract.MinTransferAmount(&_L1PoolManager.CallOpts)
}

// MinTransferAmount is a free data retrieval call binding the contract method 0x09dc480c.
//
// Solidity: function MinTransferAmount() view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) MinTransferAmount() (*big.Int, error) {
	return _L1PoolManager.Contract.MinTransferAmount(&_L1PoolManager.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) PerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "PerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) PerFee() (*big.Int, error) {
	return _L1PoolManager.Contract.PerFee(&_L1PoolManager.CallOpts)
}

// PerFee is a free data retrieval call binding the contract method 0xdd0c3460.
//
// Solidity: function PerFee() view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) PerFee() (*big.Int, error) {
	return _L1PoolManager.Contract.PerFee(&_L1PoolManager.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0x4663cdc8.
//
// Solidity: function Pools(address , uint256 ) view returns(uint32 startTimestamp, uint32 endTimestamp, address token, uint256 TotalAmount, uint256 TotalFee, uint256 TotalFeeClaimed, bool IsCompleted)
func (_L1PoolManager *L1PoolManagerCaller) Pools(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StartTimestamp  uint32
	EndTimestamp    uint32
	Token           common.Address
	TotalAmount     *big.Int
	TotalFee        *big.Int
	TotalFeeClaimed *big.Int
	IsCompleted     bool
}, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "Pools", arg0, arg1)

	outstruct := new(struct {
		StartTimestamp  uint32
		EndTimestamp    uint32
		Token           common.Address
		TotalAmount     *big.Int
		TotalFee        *big.Int
		TotalFeeClaimed *big.Int
		IsCompleted     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTimestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.EndTimestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Token = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TotalAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalFee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalFeeClaimed = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.IsCompleted = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Pools is a free data retrieval call binding the contract method 0x4663cdc8.
//
// Solidity: function Pools(address , uint256 ) view returns(uint32 startTimestamp, uint32 endTimestamp, address token, uint256 TotalAmount, uint256 TotalFee, uint256 TotalFeeClaimed, bool IsCompleted)
func (_L1PoolManager *L1PoolManagerSession) Pools(arg0 common.Address, arg1 *big.Int) (struct {
	StartTimestamp  uint32
	EndTimestamp    uint32
	Token           common.Address
	TotalAmount     *big.Int
	TotalFee        *big.Int
	TotalFeeClaimed *big.Int
	IsCompleted     bool
}, error) {
	return _L1PoolManager.Contract.Pools(&_L1PoolManager.CallOpts, arg0, arg1)
}

// Pools is a free data retrieval call binding the contract method 0x4663cdc8.
//
// Solidity: function Pools(address , uint256 ) view returns(uint32 startTimestamp, uint32 endTimestamp, address token, uint256 TotalAmount, uint256 TotalFee, uint256 TotalFeeClaimed, bool IsCompleted)
func (_L1PoolManager *L1PoolManagerCallerSession) Pools(arg0 common.Address, arg1 *big.Int) (struct {
	StartTimestamp  uint32
	EndTimestamp    uint32
	Token           common.Address
	TotalAmount     *big.Int
	TotalFee        *big.Int
	TotalFeeClaimed *big.Int
	IsCompleted     bool
}, error) {
	return _L1PoolManager.Contract.Pools(&_L1PoolManager.CallOpts, arg0, arg1)
}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_L1PoolManager *L1PoolManagerCaller) ReLayer(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "ReLayer")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_L1PoolManager *L1PoolManagerSession) ReLayer() ([32]byte, error) {
	return _L1PoolManager.Contract.ReLayer(&_L1PoolManager.CallOpts)
}

// ReLayer is a free data retrieval call binding the contract method 0xfa861848.
//
// Solidity: function ReLayer() view returns(bytes32)
func (_L1PoolManager *L1PoolManagerCallerSession) ReLayer() ([32]byte, error) {
	return _L1PoolManager.Contract.ReLayer(&_L1PoolManager.CallOpts)
}

// SupportTokens is a free data retrieval call binding the contract method 0x626417b5.
//
// Solidity: function SupportTokens(uint256 ) view returns(address)
func (_L1PoolManager *L1PoolManagerCaller) SupportTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "SupportTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportTokens is a free data retrieval call binding the contract method 0x626417b5.
//
// Solidity: function SupportTokens(uint256 ) view returns(address)
func (_L1PoolManager *L1PoolManagerSession) SupportTokens(arg0 *big.Int) (common.Address, error) {
	return _L1PoolManager.Contract.SupportTokens(&_L1PoolManager.CallOpts, arg0)
}

// SupportTokens is a free data retrieval call binding the contract method 0x626417b5.
//
// Solidity: function SupportTokens(uint256 ) view returns(address)
func (_L1PoolManager *L1PoolManagerCallerSession) SupportTokens(arg0 *big.Int) (common.Address, error) {
	return _L1PoolManager.Contract.SupportTokens(&_L1PoolManager.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x13e8544e.
//
// Solidity: function Users(address , uint256 ) view returns(bool isWithdrawed, address token, uint256 StartPoolId, uint256 EndPoolId, uint256 Amount)
func (_L1PoolManager *L1PoolManagerCaller) Users(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	IsWithdrawed bool
	Token        common.Address
	StartPoolId  *big.Int
	EndPoolId    *big.Int
	Amount       *big.Int
}, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "Users", arg0, arg1)

	outstruct := new(struct {
		IsWithdrawed bool
		Token        common.Address
		StartPoolId  *big.Int
		EndPoolId    *big.Int
		Amount       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsWithdrawed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StartPoolId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndPoolId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0x13e8544e.
//
// Solidity: function Users(address , uint256 ) view returns(bool isWithdrawed, address token, uint256 StartPoolId, uint256 EndPoolId, uint256 Amount)
func (_L1PoolManager *L1PoolManagerSession) Users(arg0 common.Address, arg1 *big.Int) (struct {
	IsWithdrawed bool
	Token        common.Address
	StartPoolId  *big.Int
	EndPoolId    *big.Int
	Amount       *big.Int
}, error) {
	return _L1PoolManager.Contract.Users(&_L1PoolManager.CallOpts, arg0, arg1)
}

// Users is a free data retrieval call binding the contract method 0x13e8544e.
//
// Solidity: function Users(address , uint256 ) view returns(bool isWithdrawed, address token, uint256 StartPoolId, uint256 EndPoolId, uint256 Amount)
func (_L1PoolManager *L1PoolManagerCallerSession) Users(arg0 common.Address, arg1 *big.Int) (struct {
	IsWithdrawed bool
	Token        common.Address
	StartPoolId  *big.Int
	EndPoolId    *big.Int
	Amount       *big.Int
}, error) {
	return _L1PoolManager.Contract.Users(&_L1PoolManager.CallOpts, arg0, arg1)
}

// GetPool is a free data retrieval call binding the contract method 0x5b5b9ea2.
//
// Solidity: function getPool(address _token, uint256 _index) view returns((uint32,uint32,address,uint256,uint256,uint256,bool))
func (_L1PoolManager *L1PoolManagerCaller) GetPool(opts *bind.CallOpts, _token common.Address, _index *big.Int) (IL1PoolManagerPool, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getPool", _token, _index)

	if err != nil {
		return *new(IL1PoolManagerPool), err
	}

	out0 := *abi.ConvertType(out[0], new(IL1PoolManagerPool)).(*IL1PoolManagerPool)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x5b5b9ea2.
//
// Solidity: function getPool(address _token, uint256 _index) view returns((uint32,uint32,address,uint256,uint256,uint256,bool))
func (_L1PoolManager *L1PoolManagerSession) GetPool(_token common.Address, _index *big.Int) (IL1PoolManagerPool, error) {
	return _L1PoolManager.Contract.GetPool(&_L1PoolManager.CallOpts, _token, _index)
}

// GetPool is a free data retrieval call binding the contract method 0x5b5b9ea2.
//
// Solidity: function getPool(address _token, uint256 _index) view returns((uint32,uint32,address,uint256,uint256,uint256,bool))
func (_L1PoolManager *L1PoolManagerCallerSession) GetPool(_token common.Address, _index *big.Int) (IL1PoolManagerPool, error) {
	return _L1PoolManager.Contract.GetPool(&_L1PoolManager.CallOpts, _token, _index)
}

// GetPoolLength is a free data retrieval call binding the contract method 0xf8fcc2aa.
//
// Solidity: function getPoolLength(address _token) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) GetPoolLength(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getPoolLength", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolLength is a free data retrieval call binding the contract method 0xf8fcc2aa.
//
// Solidity: function getPoolLength(address _token) view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) GetPoolLength(_token common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.GetPoolLength(&_L1PoolManager.CallOpts, _token)
}

// GetPoolLength is a free data retrieval call binding the contract method 0xf8fcc2aa.
//
// Solidity: function getPoolLength(address _token) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) GetPoolLength(_token common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.GetPoolLength(&_L1PoolManager.CallOpts, _token)
}

// GetPrincipal is a free data retrieval call binding the contract method 0xa69bdf16.
//
// Solidity: function getPrincipal() view returns((address,uint256)[])
func (_L1PoolManager *L1PoolManagerCaller) GetPrincipal(opts *bind.CallOpts) ([]IL1PoolManagerKeyValuePair, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getPrincipal")

	if err != nil {
		return *new([]IL1PoolManagerKeyValuePair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IL1PoolManagerKeyValuePair)).(*[]IL1PoolManagerKeyValuePair)

	return out0, err

}

// GetPrincipal is a free data retrieval call binding the contract method 0xa69bdf16.
//
// Solidity: function getPrincipal() view returns((address,uint256)[])
func (_L1PoolManager *L1PoolManagerSession) GetPrincipal() ([]IL1PoolManagerKeyValuePair, error) {
	return _L1PoolManager.Contract.GetPrincipal(&_L1PoolManager.CallOpts)
}

// GetPrincipal is a free data retrieval call binding the contract method 0xa69bdf16.
//
// Solidity: function getPrincipal() view returns((address,uint256)[])
func (_L1PoolManager *L1PoolManagerCallerSession) GetPrincipal() ([]IL1PoolManagerKeyValuePair, error) {
	return _L1PoolManager.Contract.GetPrincipal(&_L1PoolManager.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0x3d18b912.
//
// Solidity: function getReward() view returns((address,uint256)[])
func (_L1PoolManager *L1PoolManagerCaller) GetReward(opts *bind.CallOpts) ([]IL1PoolManagerKeyValuePair, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getReward")

	if err != nil {
		return *new([]IL1PoolManagerKeyValuePair), err
	}

	out0 := *abi.ConvertType(out[0], new([]IL1PoolManagerKeyValuePair)).(*[]IL1PoolManagerKeyValuePair)

	return out0, err

}

// GetReward is a free data retrieval call binding the contract method 0x3d18b912.
//
// Solidity: function getReward() view returns((address,uint256)[])
func (_L1PoolManager *L1PoolManagerSession) GetReward() ([]IL1PoolManagerKeyValuePair, error) {
	return _L1PoolManager.Contract.GetReward(&_L1PoolManager.CallOpts)
}

// GetReward is a free data retrieval call binding the contract method 0x3d18b912.
//
// Solidity: function getReward() view returns((address,uint256)[])
func (_L1PoolManager *L1PoolManagerCallerSession) GetReward() ([]IL1PoolManagerKeyValuePair, error) {
	return _L1PoolManager.Contract.GetReward(&_L1PoolManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L1PoolManager *L1PoolManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L1PoolManager *L1PoolManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _L1PoolManager.Contract.GetRoleAdmin(&_L1PoolManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_L1PoolManager *L1PoolManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _L1PoolManager.Contract.GetRoleAdmin(&_L1PoolManager.CallOpts, role)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _user) view returns((bool,address,uint256,uint256,uint256)[])
func (_L1PoolManager *L1PoolManagerCaller) GetUser(opts *bind.CallOpts, _user common.Address) ([]IL1PoolManagerUser, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getUser", _user)

	if err != nil {
		return *new([]IL1PoolManagerUser), err
	}

	out0 := *abi.ConvertType(out[0], new([]IL1PoolManagerUser)).(*[]IL1PoolManagerUser)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _user) view returns((bool,address,uint256,uint256,uint256)[])
func (_L1PoolManager *L1PoolManagerSession) GetUser(_user common.Address) ([]IL1PoolManagerUser, error) {
	return _L1PoolManager.Contract.GetUser(&_L1PoolManager.CallOpts, _user)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _user) view returns((bool,address,uint256,uint256,uint256)[])
func (_L1PoolManager *L1PoolManagerCallerSession) GetUser(_user common.Address) ([]IL1PoolManagerUser, error) {
	return _L1PoolManager.Contract.GetUser(&_L1PoolManager.CallOpts, _user)
}

// GetUser0 is a free data retrieval call binding the contract method 0xff2bf64f.
//
// Solidity: function getUser(address _user, uint256 _index) view returns((bool,address,uint256,uint256,uint256))
func (_L1PoolManager *L1PoolManagerCaller) GetUser0(opts *bind.CallOpts, _user common.Address, _index *big.Int) (IL1PoolManagerUser, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getUser0", _user, _index)

	if err != nil {
		return *new(IL1PoolManagerUser), err
	}

	out0 := *abi.ConvertType(out[0], new(IL1PoolManagerUser)).(*IL1PoolManagerUser)

	return out0, err

}

// GetUser0 is a free data retrieval call binding the contract method 0xff2bf64f.
//
// Solidity: function getUser(address _user, uint256 _index) view returns((bool,address,uint256,uint256,uint256))
func (_L1PoolManager *L1PoolManagerSession) GetUser0(_user common.Address, _index *big.Int) (IL1PoolManagerUser, error) {
	return _L1PoolManager.Contract.GetUser0(&_L1PoolManager.CallOpts, _user, _index)
}

// GetUser0 is a free data retrieval call binding the contract method 0xff2bf64f.
//
// Solidity: function getUser(address _user, uint256 _index) view returns((bool,address,uint256,uint256,uint256))
func (_L1PoolManager *L1PoolManagerCallerSession) GetUser0(_user common.Address, _index *big.Int) (IL1PoolManagerUser, error) {
	return _L1PoolManager.Contract.GetUser0(&_L1PoolManager.CallOpts, _user, _index)
}

// GetUserLength is a free data retrieval call binding the contract method 0xcb4f04ad.
//
// Solidity: function getUserLength(address _user) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCaller) GetUserLength(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "getUserLength", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLength is a free data retrieval call binding the contract method 0xcb4f04ad.
//
// Solidity: function getUserLength(address _user) view returns(uint256)
func (_L1PoolManager *L1PoolManagerSession) GetUserLength(_user common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.GetUserLength(&_L1PoolManager.CallOpts, _user)
}

// GetUserLength is a free data retrieval call binding the contract method 0xcb4f04ad.
//
// Solidity: function getUserLength(address _user) view returns(uint256)
func (_L1PoolManager *L1PoolManagerCallerSession) GetUserLength(_user common.Address) (*big.Int, error) {
	return _L1PoolManager.Contract.GetUserLength(&_L1PoolManager.CallOpts, _user)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L1PoolManager *L1PoolManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L1PoolManager *L1PoolManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _L1PoolManager.Contract.HasRole(&_L1PoolManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_L1PoolManager *L1PoolManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _L1PoolManager.Contract.HasRole(&_L1PoolManager.CallOpts, role, account)
}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_L1PoolManager *L1PoolManagerCaller) MessageManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "messageManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_L1PoolManager *L1PoolManagerSession) MessageManager() (common.Address, error) {
	return _L1PoolManager.Contract.MessageManager(&_L1PoolManager.CallOpts)
}

// MessageManager is a free data retrieval call binding the contract method 0x10875a13.
//
// Solidity: function messageManager() view returns(address)
func (_L1PoolManager *L1PoolManagerCallerSession) MessageManager() (common.Address, error) {
	return _L1PoolManager.Contract.MessageManager(&_L1PoolManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1PoolManager *L1PoolManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1PoolManager *L1PoolManagerSession) Paused() (bool, error) {
	return _L1PoolManager.Contract.Paused(&_L1PoolManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1PoolManager *L1PoolManagerCallerSession) Paused() (bool, error) {
	return _L1PoolManager.Contract.Paused(&_L1PoolManager.CallOpts)
}

// PeriodTime is a free data retrieval call binding the contract method 0x1d31fac0.
//
// Solidity: function periodTime() view returns(uint32)
func (_L1PoolManager *L1PoolManagerCaller) PeriodTime(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "periodTime")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PeriodTime is a free data retrieval call binding the contract method 0x1d31fac0.
//
// Solidity: function periodTime() view returns(uint32)
func (_L1PoolManager *L1PoolManagerSession) PeriodTime() (uint32, error) {
	return _L1PoolManager.Contract.PeriodTime(&_L1PoolManager.CallOpts)
}

// PeriodTime is a free data retrieval call binding the contract method 0x1d31fac0.
//
// Solidity: function periodTime() view returns(uint32)
func (_L1PoolManager *L1PoolManagerCallerSession) PeriodTime() (uint32, error) {
	return _L1PoolManager.Contract.PeriodTime(&_L1PoolManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1PoolManager *L1PoolManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L1PoolManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1PoolManager *L1PoolManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L1PoolManager.Contract.SupportsInterface(&_L1PoolManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1PoolManager *L1PoolManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L1PoolManager.Contract.SupportsInterface(&_L1PoolManager.CallOpts, interfaceId)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x92d84738.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactor) BridgeFinalizeERC20(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeFinalizeERC20", sourceChainId, destChainId, to, ERC20Address, amount, _fee, _nonce)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x92d84738.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L1PoolManager *L1PoolManagerSession) BridgeFinalizeERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeERC20(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, amount, _fee, _nonce)
}

// BridgeFinalizeERC20 is a paid mutator transaction binding the contract method 0x92d84738.
//
// Solidity: function BridgeFinalizeERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeFinalizeERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeERC20(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x38731c47.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_L1PoolManager *L1PoolManagerTransactor) BridgeFinalizeETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeFinalizeETH", sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x38731c47.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_L1PoolManager *L1PoolManagerSession) BridgeFinalizeETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeETH is a paid mutator transaction binding the contract method 0x38731c47.
//
// Solidity: function BridgeFinalizeETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) payable returns(bool)
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeFinalizeETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeETHForStaking is a paid mutator transaction binding the contract method 0xcc9c0218.
//
// Solidity: function BridgeFinalizeETHForStaking(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce, address stakingManager) payable returns()
func (_L1PoolManager *L1PoolManagerTransactor) BridgeFinalizeETHForStaking(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int, stakingManager common.Address) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeFinalizeETHForStaking", sourceChainId, destChainId, to, amount, _fee, _nonce, stakingManager)
}

// BridgeFinalizeETHForStaking is a paid mutator transaction binding the contract method 0xcc9c0218.
//
// Solidity: function BridgeFinalizeETHForStaking(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce, address stakingManager) payable returns()
func (_L1PoolManager *L1PoolManagerSession) BridgeFinalizeETHForStaking(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int, stakingManager common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeETHForStaking(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce, stakingManager)
}

// BridgeFinalizeETHForStaking is a paid mutator transaction binding the contract method 0xcc9c0218.
//
// Solidity: function BridgeFinalizeETHForStaking(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce, address stakingManager) payable returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeFinalizeETHForStaking(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int, stakingManager common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeETHForStaking(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce, stakingManager)
}

// BridgeFinalizeWETH is a paid mutator transaction binding the contract method 0x5765a4eb.
//
// Solidity: function BridgeFinalizeWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactor) BridgeFinalizeWETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeFinalizeWETH", sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeWETH is a paid mutator transaction binding the contract method 0x5765a4eb.
//
// Solidity: function BridgeFinalizeWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L1PoolManager *L1PoolManagerSession) BridgeFinalizeWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeWETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeFinalizeWETH is a paid mutator transaction binding the contract method 0x5765a4eb.
//
// Solidity: function BridgeFinalizeWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 amount, uint256 _fee, uint256 _nonce) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeFinalizeWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, amount *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeFinalizeWETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, amount, _fee, _nonce)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x3b0230f0.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactor) BridgeInitiateERC20(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeInitiateERC20", sourceChainId, destChainId, to, ERC20Address, value)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x3b0230f0.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_L1PoolManager *L1PoolManagerSession) BridgeInitiateERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeInitiateERC20(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, value)
}

// BridgeInitiateERC20 is a paid mutator transaction binding the contract method 0x3b0230f0.
//
// Solidity: function BridgeInitiateERC20(uint256 sourceChainId, uint256 destChainId, address to, address ERC20Address, uint256 value) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeInitiateERC20(sourceChainId *big.Int, destChainId *big.Int, to common.Address, ERC20Address common.Address, value *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeInitiateERC20(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, ERC20Address, value)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x9b442380.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address to) payable returns(bool)
func (_L1PoolManager *L1PoolManagerTransactor) BridgeInitiateETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeInitiateETH", sourceChainId, destChainId, to)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x9b442380.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address to) payable returns(bool)
func (_L1PoolManager *L1PoolManagerSession) BridgeInitiateETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeInitiateETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to)
}

// BridgeInitiateETH is a paid mutator transaction binding the contract method 0x9b442380.
//
// Solidity: function BridgeInitiateETH(uint256 sourceChainId, uint256 destChainId, address to) payable returns(bool)
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeInitiateETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeInitiateETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to)
}

// BridgeInitiateWETH is a paid mutator transaction binding the contract method 0xcd01c665.
//
// Solidity: function BridgeInitiateWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 value) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactor) BridgeInitiateWETH(opts *bind.TransactOpts, sourceChainId *big.Int, destChainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "BridgeInitiateWETH", sourceChainId, destChainId, to, value)
}

// BridgeInitiateWETH is a paid mutator transaction binding the contract method 0xcd01c665.
//
// Solidity: function BridgeInitiateWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 value) returns(bool)
func (_L1PoolManager *L1PoolManagerSession) BridgeInitiateWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeInitiateWETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, value)
}

// BridgeInitiateWETH is a paid mutator transaction binding the contract method 0xcd01c665.
//
// Solidity: function BridgeInitiateWETH(uint256 sourceChainId, uint256 destChainId, address to, uint256 value) returns(bool)
func (_L1PoolManager *L1PoolManagerTransactorSession) BridgeInitiateWETH(sourceChainId *big.Int, destChainId *big.Int, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.BridgeInitiateWETH(&_L1PoolManager.TransactOpts, sourceChainId, destChainId, to, value)
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x2c37388b.
//
// Solidity: function ClaimAllReward() returns()
func (_L1PoolManager *L1PoolManagerTransactor) ClaimAllReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "ClaimAllReward")
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x2c37388b.
//
// Solidity: function ClaimAllReward() returns()
func (_L1PoolManager *L1PoolManagerSession) ClaimAllReward() (*types.Transaction, error) {
	return _L1PoolManager.Contract.ClaimAllReward(&_L1PoolManager.TransactOpts)
}

// ClaimAllReward is a paid mutator transaction binding the contract method 0x2c37388b.
//
// Solidity: function ClaimAllReward() returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) ClaimAllReward() (*types.Transaction, error) {
	return _L1PoolManager.Contract.ClaimAllReward(&_L1PoolManager.TransactOpts)
}

// ClaimbyID is a paid mutator transaction binding the contract method 0x34d34ef5.
//
// Solidity: function ClaimbyID(uint256 i) returns()
func (_L1PoolManager *L1PoolManagerTransactor) ClaimbyID(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "ClaimbyID", i)
}

// ClaimbyID is a paid mutator transaction binding the contract method 0x34d34ef5.
//
// Solidity: function ClaimbyID(uint256 i) returns()
func (_L1PoolManager *L1PoolManagerSession) ClaimbyID(i *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.ClaimbyID(&_L1PoolManager.TransactOpts, i)
}

// ClaimbyID is a paid mutator transaction binding the contract method 0x34d34ef5.
//
// Solidity: function ClaimbyID(uint256 i) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) ClaimbyID(i *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.ClaimbyID(&_L1PoolManager.TransactOpts, i)
}

// CompletePoolAndNew is a paid mutator transaction binding the contract method 0x3338562c.
//
// Solidity: function CompletePoolAndNew((uint32,uint32,address,uint256,uint256,uint256,bool)[] CompletePools) payable returns()
func (_L1PoolManager *L1PoolManagerTransactor) CompletePoolAndNew(opts *bind.TransactOpts, CompletePools []IL1PoolManagerPool) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "CompletePoolAndNew", CompletePools)
}

// CompletePoolAndNew is a paid mutator transaction binding the contract method 0x3338562c.
//
// Solidity: function CompletePoolAndNew((uint32,uint32,address,uint256,uint256,uint256,bool)[] CompletePools) payable returns()
func (_L1PoolManager *L1PoolManagerSession) CompletePoolAndNew(CompletePools []IL1PoolManagerPool) (*types.Transaction, error) {
	return _L1PoolManager.Contract.CompletePoolAndNew(&_L1PoolManager.TransactOpts, CompletePools)
}

// CompletePoolAndNew is a paid mutator transaction binding the contract method 0x3338562c.
//
// Solidity: function CompletePoolAndNew((uint32,uint32,address,uint256,uint256,uint256,bool)[] CompletePools) payable returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) CompletePoolAndNew(CompletePools []IL1PoolManagerPool) (*types.Transaction, error) {
	return _L1PoolManager.Contract.CompletePoolAndNew(&_L1PoolManager.TransactOpts, CompletePools)
}

// DepositAndStaking is a paid mutator transaction binding the contract method 0x9ab0b652.
//
// Solidity: function DepositAndStaking(address _token, uint256 _amount) payable returns()
func (_L1PoolManager *L1PoolManagerTransactor) DepositAndStaking(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "DepositAndStaking", _token, _amount)
}

// DepositAndStaking is a paid mutator transaction binding the contract method 0x9ab0b652.
//
// Solidity: function DepositAndStaking(address _token, uint256 _amount) payable returns()
func (_L1PoolManager *L1PoolManagerSession) DepositAndStaking(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStaking(&_L1PoolManager.TransactOpts, _token, _amount)
}

// DepositAndStaking is a paid mutator transaction binding the contract method 0x9ab0b652.
//
// Solidity: function DepositAndStaking(address _token, uint256 _amount) payable returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) DepositAndStaking(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStaking(&_L1PoolManager.TransactOpts, _token, _amount)
}

// DepositAndStakingERC20 is a paid mutator transaction binding the contract method 0x4dfaef84.
//
// Solidity: function DepositAndStakingERC20(address _token, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactor) DepositAndStakingERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "DepositAndStakingERC20", _token, _amount)
}

// DepositAndStakingERC20 is a paid mutator transaction binding the contract method 0x4dfaef84.
//
// Solidity: function DepositAndStakingERC20(address _token, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerSession) DepositAndStakingERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStakingERC20(&_L1PoolManager.TransactOpts, _token, _amount)
}

// DepositAndStakingERC20 is a paid mutator transaction binding the contract method 0x4dfaef84.
//
// Solidity: function DepositAndStakingERC20(address _token, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) DepositAndStakingERC20(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStakingERC20(&_L1PoolManager.TransactOpts, _token, _amount)
}

// DepositAndStakingETH is a paid mutator transaction binding the contract method 0x96f984e6.
//
// Solidity: function DepositAndStakingETH() payable returns()
func (_L1PoolManager *L1PoolManagerTransactor) DepositAndStakingETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "DepositAndStakingETH")
}

// DepositAndStakingETH is a paid mutator transaction binding the contract method 0x96f984e6.
//
// Solidity: function DepositAndStakingETH() payable returns()
func (_L1PoolManager *L1PoolManagerSession) DepositAndStakingETH() (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStakingETH(&_L1PoolManager.TransactOpts)
}

// DepositAndStakingETH is a paid mutator transaction binding the contract method 0x96f984e6.
//
// Solidity: function DepositAndStakingETH() payable returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) DepositAndStakingETH() (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStakingETH(&_L1PoolManager.TransactOpts)
}

// DepositAndStakingWETH is a paid mutator transaction binding the contract method 0xbc42493d.
//
// Solidity: function DepositAndStakingWETH(uint256 amount) returns()
func (_L1PoolManager *L1PoolManagerTransactor) DepositAndStakingWETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "DepositAndStakingWETH", amount)
}

// DepositAndStakingWETH is a paid mutator transaction binding the contract method 0xbc42493d.
//
// Solidity: function DepositAndStakingWETH(uint256 amount) returns()
func (_L1PoolManager *L1PoolManagerSession) DepositAndStakingWETH(amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStakingWETH(&_L1PoolManager.TransactOpts, amount)
}

// DepositAndStakingWETH is a paid mutator transaction binding the contract method 0xbc42493d.
//
// Solidity: function DepositAndStakingWETH(uint256 amount) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) DepositAndStakingWETH(amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.DepositAndStakingWETH(&_L1PoolManager.TransactOpts, amount)
}

// QuickSendAssertToUser is a paid mutator transaction binding the contract method 0x325e5d43.
//
// Solidity: function QuickSendAssertToUser(address _token, address to, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactor) QuickSendAssertToUser(opts *bind.TransactOpts, _token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "QuickSendAssertToUser", _token, to, _amount)
}

// QuickSendAssertToUser is a paid mutator transaction binding the contract method 0x325e5d43.
//
// Solidity: function QuickSendAssertToUser(address _token, address to, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerSession) QuickSendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.QuickSendAssertToUser(&_L1PoolManager.TransactOpts, _token, to, _amount)
}

// QuickSendAssertToUser is a paid mutator transaction binding the contract method 0x325e5d43.
//
// Solidity: function QuickSendAssertToUser(address _token, address to, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) QuickSendAssertToUser(_token common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.QuickSendAssertToUser(&_L1PoolManager.TransactOpts, _token, to, _amount)
}

// TransferAssertToBridge is a paid mutator transaction binding the contract method 0xf91fa9a8.
//
// Solidity: function TransferAssertToBridge(uint256 Blockchain, address _token, address _to, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactor) TransferAssertToBridge(opts *bind.TransactOpts, Blockchain *big.Int, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "TransferAssertToBridge", Blockchain, _token, _to, _amount)
}

// TransferAssertToBridge is a paid mutator transaction binding the contract method 0xf91fa9a8.
//
// Solidity: function TransferAssertToBridge(uint256 Blockchain, address _token, address _to, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerSession) TransferAssertToBridge(Blockchain *big.Int, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.TransferAssertToBridge(&_L1PoolManager.TransactOpts, Blockchain, _token, _to, _amount)
}

// TransferAssertToBridge is a paid mutator transaction binding the contract method 0xf91fa9a8.
//
// Solidity: function TransferAssertToBridge(uint256 Blockchain, address _token, address _to, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) TransferAssertToBridge(Blockchain *big.Int, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.TransferAssertToBridge(&_L1PoolManager.TransactOpts, Blockchain, _token, _to, _amount)
}

// UpdateFundingPoolBalance is a paid mutator transaction binding the contract method 0x0cd727df.
//
// Solidity: function UpdateFundingPoolBalance(address token, uint256 amount) returns()
func (_L1PoolManager *L1PoolManagerTransactor) UpdateFundingPoolBalance(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "UpdateFundingPoolBalance", token, amount)
}

// UpdateFundingPoolBalance is a paid mutator transaction binding the contract method 0x0cd727df.
//
// Solidity: function UpdateFundingPoolBalance(address token, uint256 amount) returns()
func (_L1PoolManager *L1PoolManagerSession) UpdateFundingPoolBalance(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.UpdateFundingPoolBalance(&_L1PoolManager.TransactOpts, token, amount)
}

// UpdateFundingPoolBalance is a paid mutator transaction binding the contract method 0x0cd727df.
//
// Solidity: function UpdateFundingPoolBalance(address token, uint256 amount) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) UpdateFundingPoolBalance(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.UpdateFundingPoolBalance(&_L1PoolManager.TransactOpts, token, amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x5a648bc5.
//
// Solidity: function WithdrawAll() returns()
func (_L1PoolManager *L1PoolManagerTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "WithdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x5a648bc5.
//
// Solidity: function WithdrawAll() returns()
func (_L1PoolManager *L1PoolManagerSession) WithdrawAll() (*types.Transaction, error) {
	return _L1PoolManager.Contract.WithdrawAll(&_L1PoolManager.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x5a648bc5.
//
// Solidity: function WithdrawAll() returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _L1PoolManager.Contract.WithdrawAll(&_L1PoolManager.TransactOpts)
}

// WithdrawByID is a paid mutator transaction binding the contract method 0xf62a89cf.
//
// Solidity: function WithdrawByID(uint256 i) returns()
func (_L1PoolManager *L1PoolManagerTransactor) WithdrawByID(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "WithdrawByID", i)
}

// WithdrawByID is a paid mutator transaction binding the contract method 0xf62a89cf.
//
// Solidity: function WithdrawByID(uint256 i) returns()
func (_L1PoolManager *L1PoolManagerSession) WithdrawByID(i *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.WithdrawByID(&_L1PoolManager.TransactOpts, i)
}

// WithdrawByID is a paid mutator transaction binding the contract method 0xf62a89cf.
//
// Solidity: function WithdrawByID(uint256 i) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) WithdrawByID(i *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.WithdrawByID(&_L1PoolManager.TransactOpts, i)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L1PoolManager *L1PoolManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L1PoolManager *L1PoolManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.GrantRole(&_L1PoolManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.GrantRole(&_L1PoolManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _MultisigWallet, address _messageManager) returns()
func (_L1PoolManager *L1PoolManagerTransactor) Initialize(opts *bind.TransactOpts, _MultisigWallet common.Address, _messageManager common.Address) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "initialize", _MultisigWallet, _messageManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _MultisigWallet, address _messageManager) returns()
func (_L1PoolManager *L1PoolManagerSession) Initialize(_MultisigWallet common.Address, _messageManager common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.Initialize(&_L1PoolManager.TransactOpts, _MultisigWallet, _messageManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _MultisigWallet, address _messageManager) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) Initialize(_MultisigWallet common.Address, _messageManager common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.Initialize(&_L1PoolManager.TransactOpts, _MultisigWallet, _messageManager)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1PoolManager *L1PoolManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1PoolManager *L1PoolManagerSession) Pause() (*types.Transaction, error) {
	return _L1PoolManager.Contract.Pause(&_L1PoolManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _L1PoolManager.Contract.Pause(&_L1PoolManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L1PoolManager *L1PoolManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L1PoolManager *L1PoolManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.RenounceRole(&_L1PoolManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.RenounceRole(&_L1PoolManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L1PoolManager *L1PoolManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L1PoolManager *L1PoolManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.RevokeRole(&_L1PoolManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _L1PoolManager.Contract.RevokeRole(&_L1PoolManager.TransactOpts, role, account)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xcb314fab.
//
// Solidity: function setMinStakeAmount(address _token, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactor) SetMinStakeAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "setMinStakeAmount", _token, _amount)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xcb314fab.
//
// Solidity: function setMinStakeAmount(address _token, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerSession) SetMinStakeAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetMinStakeAmount(&_L1PoolManager.TransactOpts, _token, _amount)
}

// SetMinStakeAmount is a paid mutator transaction binding the contract method 0xcb314fab.
//
// Solidity: function setMinStakeAmount(address _token, uint256 _amount) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) SetMinStakeAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetMinStakeAmount(&_L1PoolManager.TransactOpts, _token, _amount)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_L1PoolManager *L1PoolManagerTransactor) SetMinTransferAmount(opts *bind.TransactOpts, _MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "setMinTransferAmount", _MinTransferAmount)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_L1PoolManager *L1PoolManagerSession) SetMinTransferAmount(_MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetMinTransferAmount(&_L1PoolManager.TransactOpts, _MinTransferAmount)
}

// SetMinTransferAmount is a paid mutator transaction binding the contract method 0xb92e6396.
//
// Solidity: function setMinTransferAmount(uint256 _MinTransferAmount) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) SetMinTransferAmount(_MinTransferAmount *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetMinTransferAmount(&_L1PoolManager.TransactOpts, _MinTransferAmount)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_L1PoolManager *L1PoolManagerTransactor) SetPerFee(opts *bind.TransactOpts, _PerFee *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "setPerFee", _PerFee)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_L1PoolManager *L1PoolManagerSession) SetPerFee(_PerFee *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetPerFee(&_L1PoolManager.TransactOpts, _PerFee)
}

// SetPerFee is a paid mutator transaction binding the contract method 0x650c2276.
//
// Solidity: function setPerFee(uint256 _PerFee) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) SetPerFee(_PerFee *big.Int) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetPerFee(&_L1PoolManager.TransactOpts, _PerFee)
}

// SetSupportERC20Token is a paid mutator transaction binding the contract method 0xab0f9e19.
//
// Solidity: function setSupportERC20Token(address ERC20Address, bool isValid) returns()
func (_L1PoolManager *L1PoolManagerTransactor) SetSupportERC20Token(opts *bind.TransactOpts, ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "setSupportERC20Token", ERC20Address, isValid)
}

// SetSupportERC20Token is a paid mutator transaction binding the contract method 0xab0f9e19.
//
// Solidity: function setSupportERC20Token(address ERC20Address, bool isValid) returns()
func (_L1PoolManager *L1PoolManagerSession) SetSupportERC20Token(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetSupportERC20Token(&_L1PoolManager.TransactOpts, ERC20Address, isValid)
}

// SetSupportERC20Token is a paid mutator transaction binding the contract method 0xab0f9e19.
//
// Solidity: function setSupportERC20Token(address ERC20Address, bool isValid) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) SetSupportERC20Token(ERC20Address common.Address, isValid bool) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetSupportERC20Token(&_L1PoolManager.TransactOpts, ERC20Address, isValid)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0xe2070893.
//
// Solidity: function setSupportToken(address _token, bool _isSupport, uint32 startTimes) returns()
func (_L1PoolManager *L1PoolManagerTransactor) SetSupportToken(opts *bind.TransactOpts, _token common.Address, _isSupport bool, startTimes uint32) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "setSupportToken", _token, _isSupport, startTimes)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0xe2070893.
//
// Solidity: function setSupportToken(address _token, bool _isSupport, uint32 startTimes) returns()
func (_L1PoolManager *L1PoolManagerSession) SetSupportToken(_token common.Address, _isSupport bool, startTimes uint32) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetSupportToken(&_L1PoolManager.TransactOpts, _token, _isSupport, startTimes)
}

// SetSupportToken is a paid mutator transaction binding the contract method 0xe2070893.
//
// Solidity: function setSupportToken(address _token, bool _isSupport, uint32 startTimes) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) SetSupportToken(_token common.Address, _isSupport bool, startTimes uint32) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetSupportToken(&_L1PoolManager.TransactOpts, _token, _isSupport, startTimes)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_L1PoolManager *L1PoolManagerTransactor) SetValidChainId(opts *bind.TransactOpts, chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "setValidChainId", chainId, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_L1PoolManager *L1PoolManagerSession) SetValidChainId(chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetValidChainId(&_L1PoolManager.TransactOpts, chainId, isValid)
}

// SetValidChainId is a paid mutator transaction binding the contract method 0xd73f14e5.
//
// Solidity: function setValidChainId(uint256 chainId, bool isValid) returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) SetValidChainId(chainId *big.Int, isValid bool) (*types.Transaction, error) {
	return _L1PoolManager.Contract.SetValidChainId(&_L1PoolManager.TransactOpts, chainId, isValid)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1PoolManager *L1PoolManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1PoolManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1PoolManager *L1PoolManagerSession) Unpause() (*types.Transaction, error) {
	return _L1PoolManager.Contract.Unpause(&_L1PoolManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1PoolManager *L1PoolManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _L1PoolManager.Contract.Unpause(&_L1PoolManager.TransactOpts)
}

// L1PoolManagerClaimEventIterator is returned from FilterClaimEvent and is used to iterate over the raw logs and unpacked data for ClaimEvent events raised by the L1PoolManager contract.
type L1PoolManagerClaimEventIterator struct {
	Event *L1PoolManagerClaimEvent // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerClaimEvent)
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
		it.Event = new(L1PoolManagerClaimEvent)
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
func (it *L1PoolManagerClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerClaimEvent represents a ClaimEvent event raised by the L1PoolManager contract.
type L1PoolManagerClaimEvent struct {
	User        common.Address
	StartPoolId *big.Int
	EndPoolId   *big.Int
	Token       common.Address
	Amount      *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaimEvent is a free log retrieval operation binding the contract event 0x991c329471ab5230e2301ee30d0d6fba5906e32411d5d154d5a8c278d021a2ab.
//
// Solidity: event ClaimEvent(address indexed user, uint256 startPoolId, uint256 endPoolId, address indexed token, uint256 amount, uint256 fee)
func (_L1PoolManager *L1PoolManagerFilterer) FilterClaimEvent(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*L1PoolManagerClaimEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "ClaimEvent", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerClaimEventIterator{contract: _L1PoolManager.contract, event: "ClaimEvent", logs: logs, sub: sub}, nil
}

// WatchClaimEvent is a free log subscription operation binding the contract event 0x991c329471ab5230e2301ee30d0d6fba5906e32411d5d154d5a8c278d021a2ab.
//
// Solidity: event ClaimEvent(address indexed user, uint256 startPoolId, uint256 endPoolId, address indexed token, uint256 amount, uint256 fee)
func (_L1PoolManager *L1PoolManagerFilterer) WatchClaimEvent(opts *bind.WatchOpts, sink chan<- *L1PoolManagerClaimEvent, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "ClaimEvent", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerClaimEvent)
				if err := _L1PoolManager.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
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

// ParseClaimEvent is a log parse operation binding the contract event 0x991c329471ab5230e2301ee30d0d6fba5906e32411d5d154d5a8c278d021a2ab.
//
// Solidity: event ClaimEvent(address indexed user, uint256 startPoolId, uint256 endPoolId, address indexed token, uint256 amount, uint256 fee)
func (_L1PoolManager *L1PoolManagerFilterer) ParseClaimEvent(log types.Log) (*L1PoolManagerClaimEvent, error) {
	event := new(L1PoolManagerClaimEvent)
	if err := _L1PoolManager.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerClaimRewardIterator is returned from FilterClaimReward and is used to iterate over the raw logs and unpacked data for ClaimReward events raised by the L1PoolManager contract.
type L1PoolManagerClaimRewardIterator struct {
	Event *L1PoolManagerClaimReward // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerClaimRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerClaimReward)
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
		it.Event = new(L1PoolManagerClaimReward)
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
func (it *L1PoolManagerClaimRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerClaimRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerClaimReward represents a ClaimReward event raised by the L1PoolManager contract.
type L1PoolManagerClaimReward struct {
	User        common.Address
	StartPoolId *big.Int
	EndPoolId   *big.Int
	Token       common.Address
	Reward      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaimReward is a free log retrieval operation binding the contract event 0x7e6fedb6f824b2112ac0626ced32f6be6324d554cc03efd112aa021ee4d740ac.
//
// Solidity: event ClaimReward(address _user, uint256 startPoolId, uint256 EndPoolId, address _token, uint256 Reward)
func (_L1PoolManager *L1PoolManagerFilterer) FilterClaimReward(opts *bind.FilterOpts) (*L1PoolManagerClaimRewardIterator, error) {

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "ClaimReward")
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerClaimRewardIterator{contract: _L1PoolManager.contract, event: "ClaimReward", logs: logs, sub: sub}, nil
}

// WatchClaimReward is a free log subscription operation binding the contract event 0x7e6fedb6f824b2112ac0626ced32f6be6324d554cc03efd112aa021ee4d740ac.
//
// Solidity: event ClaimReward(address _user, uint256 startPoolId, uint256 EndPoolId, address _token, uint256 Reward)
func (_L1PoolManager *L1PoolManagerFilterer) WatchClaimReward(opts *bind.WatchOpts, sink chan<- *L1PoolManagerClaimReward) (event.Subscription, error) {

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "ClaimReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerClaimReward)
				if err := _L1PoolManager.contract.UnpackLog(event, "ClaimReward", log); err != nil {
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

// ParseClaimReward is a log parse operation binding the contract event 0x7e6fedb6f824b2112ac0626ced32f6be6324d554cc03efd112aa021ee4d740ac.
//
// Solidity: event ClaimReward(address _user, uint256 startPoolId, uint256 EndPoolId, address _token, uint256 Reward)
func (_L1PoolManager *L1PoolManagerFilterer) ParseClaimReward(log types.Log) (*L1PoolManagerClaimReward, error) {
	event := new(L1PoolManagerClaimReward)
	if err := _L1PoolManager.contract.UnpackLog(event, "ClaimReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerCompletePoolEventIterator is returned from FilterCompletePoolEvent and is used to iterate over the raw logs and unpacked data for CompletePoolEvent events raised by the L1PoolManager contract.
type L1PoolManagerCompletePoolEventIterator struct {
	Event *L1PoolManagerCompletePoolEvent // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerCompletePoolEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerCompletePoolEvent)
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
		it.Event = new(L1PoolManagerCompletePoolEvent)
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
func (it *L1PoolManagerCompletePoolEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerCompletePoolEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerCompletePoolEvent represents a CompletePoolEvent event raised by the L1PoolManager contract.
type L1PoolManagerCompletePoolEvent struct {
	Token     common.Address
	PoolIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCompletePoolEvent is a free log retrieval operation binding the contract event 0xb6f449f07ceaf55392c9899e0797c6529908ae827c2d498c682e90d42c241167.
//
// Solidity: event CompletePoolEvent(address indexed token, uint256 poolIndex)
func (_L1PoolManager *L1PoolManagerFilterer) FilterCompletePoolEvent(opts *bind.FilterOpts, token []common.Address) (*L1PoolManagerCompletePoolEventIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "CompletePoolEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerCompletePoolEventIterator{contract: _L1PoolManager.contract, event: "CompletePoolEvent", logs: logs, sub: sub}, nil
}

// WatchCompletePoolEvent is a free log subscription operation binding the contract event 0xb6f449f07ceaf55392c9899e0797c6529908ae827c2d498c682e90d42c241167.
//
// Solidity: event CompletePoolEvent(address indexed token, uint256 poolIndex)
func (_L1PoolManager *L1PoolManagerFilterer) WatchCompletePoolEvent(opts *bind.WatchOpts, sink chan<- *L1PoolManagerCompletePoolEvent, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "CompletePoolEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerCompletePoolEvent)
				if err := _L1PoolManager.contract.UnpackLog(event, "CompletePoolEvent", log); err != nil {
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

// ParseCompletePoolEvent is a log parse operation binding the contract event 0xb6f449f07ceaf55392c9899e0797c6529908ae827c2d498c682e90d42c241167.
//
// Solidity: event CompletePoolEvent(address indexed token, uint256 poolIndex)
func (_L1PoolManager *L1PoolManagerFilterer) ParseCompletePoolEvent(log types.Log) (*L1PoolManagerCompletePoolEvent, error) {
	event := new(L1PoolManagerCompletePoolEvent)
	if err := _L1PoolManager.contract.UnpackLog(event, "CompletePoolEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerFinalizeERC20Iterator is returned from FilterFinalizeERC20 and is used to iterate over the raw logs and unpacked data for FinalizeERC20 events raised by the L1PoolManager contract.
type L1PoolManagerFinalizeERC20Iterator struct {
	Event *L1PoolManagerFinalizeERC20 // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerFinalizeERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerFinalizeERC20)
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
		it.Event = new(L1PoolManagerFinalizeERC20)
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
func (it *L1PoolManagerFinalizeERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerFinalizeERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerFinalizeERC20 represents a FinalizeERC20 event raised by the L1PoolManager contract.
type L1PoolManagerFinalizeERC20 struct {
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
func (_L1PoolManager *L1PoolManagerFilterer) FilterFinalizeERC20(opts *bind.FilterOpts, ERC20Address []common.Address, from []common.Address, to []common.Address) (*L1PoolManagerFinalizeERC20Iterator, error) {

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

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "FinalizeERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerFinalizeERC20Iterator{contract: _L1PoolManager.contract, event: "FinalizeERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeERC20 is a free log subscription operation binding the contract event 0x0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe.
//
// Solidity: event FinalizeERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) WatchFinalizeERC20(opts *bind.WatchOpts, sink chan<- *L1PoolManagerFinalizeERC20, ERC20Address []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "FinalizeERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerFinalizeERC20)
				if err := _L1PoolManager.contract.UnpackLog(event, "FinalizeERC20", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseFinalizeERC20(log types.Log) (*L1PoolManagerFinalizeERC20, error) {
	event := new(L1PoolManagerFinalizeERC20)
	if err := _L1PoolManager.contract.UnpackLog(event, "FinalizeERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerFinalizeETHIterator is returned from FilterFinalizeETH and is used to iterate over the raw logs and unpacked data for FinalizeETH events raised by the L1PoolManager contract.
type L1PoolManagerFinalizeETHIterator struct {
	Event *L1PoolManagerFinalizeETH // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerFinalizeETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerFinalizeETH)
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
		it.Event = new(L1PoolManagerFinalizeETH)
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
func (it *L1PoolManagerFinalizeETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerFinalizeETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerFinalizeETH represents a FinalizeETH event raised by the L1PoolManager contract.
type L1PoolManagerFinalizeETH struct {
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
func (_L1PoolManager *L1PoolManagerFilterer) FilterFinalizeETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1PoolManagerFinalizeETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "FinalizeETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerFinalizeETHIterator{contract: _L1PoolManager.contract, event: "FinalizeETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeETH is a free log subscription operation binding the contract event 0x61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f32.
//
// Solidity: event FinalizeETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) WatchFinalizeETH(opts *bind.WatchOpts, sink chan<- *L1PoolManagerFinalizeETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "FinalizeETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerFinalizeETH)
				if err := _L1PoolManager.contract.UnpackLog(event, "FinalizeETH", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseFinalizeETH(log types.Log) (*L1PoolManagerFinalizeETH, error) {
	event := new(L1PoolManagerFinalizeETH)
	if err := _L1PoolManager.contract.UnpackLog(event, "FinalizeETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerFinalizeWETHIterator is returned from FilterFinalizeWETH and is used to iterate over the raw logs and unpacked data for FinalizeWETH events raised by the L1PoolManager contract.
type L1PoolManagerFinalizeWETHIterator struct {
	Event *L1PoolManagerFinalizeWETH // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerFinalizeWETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerFinalizeWETH)
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
		it.Event = new(L1PoolManagerFinalizeWETH)
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
func (it *L1PoolManagerFinalizeWETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerFinalizeWETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerFinalizeWETH represents a FinalizeWETH event raised by the L1PoolManager contract.
type L1PoolManagerFinalizeWETH struct {
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
func (_L1PoolManager *L1PoolManagerFilterer) FilterFinalizeWETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1PoolManagerFinalizeWETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "FinalizeWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerFinalizeWETHIterator{contract: _L1PoolManager.contract, event: "FinalizeWETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeWETH is a free log subscription operation binding the contract event 0x7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc14.
//
// Solidity: event FinalizeWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) WatchFinalizeWETH(opts *bind.WatchOpts, sink chan<- *L1PoolManagerFinalizeWETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "FinalizeWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerFinalizeWETH)
				if err := _L1PoolManager.contract.UnpackLog(event, "FinalizeWETH", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseFinalizeWETH(log types.Log) (*L1PoolManagerFinalizeWETH, error) {
	event := new(L1PoolManagerFinalizeWETH)
	if err := _L1PoolManager.contract.UnpackLog(event, "FinalizeWETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1PoolManager contract.
type L1PoolManagerInitializedIterator struct {
	Event *L1PoolManagerInitialized // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerInitialized)
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
		it.Event = new(L1PoolManagerInitialized)
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
func (it *L1PoolManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerInitialized represents a Initialized event raised by the L1PoolManager contract.
type L1PoolManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L1PoolManager *L1PoolManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1PoolManagerInitializedIterator, error) {

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerInitializedIterator{contract: _L1PoolManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_L1PoolManager *L1PoolManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1PoolManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerInitialized)
				if err := _L1PoolManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseInitialized(log types.Log) (*L1PoolManagerInitialized, error) {
	event := new(L1PoolManagerInitialized)
	if err := _L1PoolManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerInitiateERC20Iterator is returned from FilterInitiateERC20 and is used to iterate over the raw logs and unpacked data for InitiateERC20 events raised by the L1PoolManager contract.
type L1PoolManagerInitiateERC20Iterator struct {
	Event *L1PoolManagerInitiateERC20 // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerInitiateERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerInitiateERC20)
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
		it.Event = new(L1PoolManagerInitiateERC20)
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
func (it *L1PoolManagerInitiateERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerInitiateERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerInitiateERC20 represents a InitiateERC20 event raised by the L1PoolManager contract.
type L1PoolManagerInitiateERC20 struct {
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
func (_L1PoolManager *L1PoolManagerFilterer) FilterInitiateERC20(opts *bind.FilterOpts, ERC20Address []common.Address, from []common.Address, to []common.Address) (*L1PoolManagerInitiateERC20Iterator, error) {

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

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "InitiateERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerInitiateERC20Iterator{contract: _L1PoolManager.contract, event: "InitiateERC20", logs: logs, sub: sub}, nil
}

// WatchInitiateERC20 is a free log subscription operation binding the contract event 0xece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a4.
//
// Solidity: event InitiateERC20(uint256 sourceChainId, uint256 destChainId, address indexed ERC20Address, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) WatchInitiateERC20(opts *bind.WatchOpts, sink chan<- *L1PoolManagerInitiateERC20, ERC20Address []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "InitiateERC20", ERC20AddressRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerInitiateERC20)
				if err := _L1PoolManager.contract.UnpackLog(event, "InitiateERC20", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseInitiateERC20(log types.Log) (*L1PoolManagerInitiateERC20, error) {
	event := new(L1PoolManagerInitiateERC20)
	if err := _L1PoolManager.contract.UnpackLog(event, "InitiateERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerInitiateETHIterator is returned from FilterInitiateETH and is used to iterate over the raw logs and unpacked data for InitiateETH events raised by the L1PoolManager contract.
type L1PoolManagerInitiateETHIterator struct {
	Event *L1PoolManagerInitiateETH // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerInitiateETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerInitiateETH)
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
		it.Event = new(L1PoolManagerInitiateETH)
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
func (it *L1PoolManagerInitiateETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerInitiateETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerInitiateETH represents a InitiateETH event raised by the L1PoolManager contract.
type L1PoolManagerInitiateETH struct {
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
func (_L1PoolManager *L1PoolManagerFilterer) FilterInitiateETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1PoolManagerInitiateETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "InitiateETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerInitiateETHIterator{contract: _L1PoolManager.contract, event: "InitiateETH", logs: logs, sub: sub}, nil
}

// WatchInitiateETH is a free log subscription operation binding the contract event 0x0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc5.
//
// Solidity: event InitiateETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) WatchInitiateETH(opts *bind.WatchOpts, sink chan<- *L1PoolManagerInitiateETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "InitiateETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerInitiateETH)
				if err := _L1PoolManager.contract.UnpackLog(event, "InitiateETH", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseInitiateETH(log types.Log) (*L1PoolManagerInitiateETH, error) {
	event := new(L1PoolManagerInitiateETH)
	if err := _L1PoolManager.contract.UnpackLog(event, "InitiateETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerInitiateWETHIterator is returned from FilterInitiateWETH and is used to iterate over the raw logs and unpacked data for InitiateWETH events raised by the L1PoolManager contract.
type L1PoolManagerInitiateWETHIterator struct {
	Event *L1PoolManagerInitiateWETH // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerInitiateWETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerInitiateWETH)
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
		it.Event = new(L1PoolManagerInitiateWETH)
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
func (it *L1PoolManagerInitiateWETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerInitiateWETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerInitiateWETH represents a InitiateWETH event raised by the L1PoolManager contract.
type L1PoolManagerInitiateWETH struct {
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
func (_L1PoolManager *L1PoolManagerFilterer) FilterInitiateWETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1PoolManagerInitiateWETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "InitiateWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerInitiateWETHIterator{contract: _L1PoolManager.contract, event: "InitiateWETH", logs: logs, sub: sub}, nil
}

// WatchInitiateWETH is a free log subscription operation binding the contract event 0xcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a89.
//
// Solidity: event InitiateWETH(uint256 sourceChainId, uint256 destChainId, address indexed from, address indexed to, uint256 value)
func (_L1PoolManager *L1PoolManagerFilterer) WatchInitiateWETH(opts *bind.WatchOpts, sink chan<- *L1PoolManagerInitiateWETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "InitiateWETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerInitiateWETH)
				if err := _L1PoolManager.contract.UnpackLog(event, "InitiateWETH", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseInitiateWETH(log types.Log) (*L1PoolManagerInitiateWETH, error) {
	event := new(L1PoolManagerInitiateWETH)
	if err := _L1PoolManager.contract.UnpackLog(event, "InitiateWETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the L1PoolManager contract.
type L1PoolManagerPausedIterator struct {
	Event *L1PoolManagerPaused // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerPaused)
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
		it.Event = new(L1PoolManagerPaused)
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
func (it *L1PoolManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerPaused represents a Paused event raised by the L1PoolManager contract.
type L1PoolManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1PoolManager *L1PoolManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*L1PoolManagerPausedIterator, error) {

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerPausedIterator{contract: _L1PoolManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1PoolManager *L1PoolManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *L1PoolManagerPaused) (event.Subscription, error) {

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerPaused)
				if err := _L1PoolManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParsePaused(log types.Log) (*L1PoolManagerPaused, error) {
	event := new(L1PoolManagerPaused)
	if err := _L1PoolManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the L1PoolManager contract.
type L1PoolManagerRoleAdminChangedIterator struct {
	Event *L1PoolManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerRoleAdminChanged)
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
		it.Event = new(L1PoolManagerRoleAdminChanged)
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
func (it *L1PoolManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerRoleAdminChanged represents a RoleAdminChanged event raised by the L1PoolManager contract.
type L1PoolManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_L1PoolManager *L1PoolManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*L1PoolManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerRoleAdminChangedIterator{contract: _L1PoolManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_L1PoolManager *L1PoolManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *L1PoolManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerRoleAdminChanged)
				if err := _L1PoolManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseRoleAdminChanged(log types.Log) (*L1PoolManagerRoleAdminChanged, error) {
	event := new(L1PoolManagerRoleAdminChanged)
	if err := _L1PoolManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the L1PoolManager contract.
type L1PoolManagerRoleGrantedIterator struct {
	Event *L1PoolManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerRoleGranted)
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
		it.Event = new(L1PoolManagerRoleGranted)
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
func (it *L1PoolManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerRoleGranted represents a RoleGranted event raised by the L1PoolManager contract.
type L1PoolManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L1PoolManager *L1PoolManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L1PoolManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerRoleGrantedIterator{contract: _L1PoolManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_L1PoolManager *L1PoolManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *L1PoolManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerRoleGranted)
				if err := _L1PoolManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseRoleGranted(log types.Log) (*L1PoolManagerRoleGranted, error) {
	event := new(L1PoolManagerRoleGranted)
	if err := _L1PoolManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the L1PoolManager contract.
type L1PoolManagerRoleRevokedIterator struct {
	Event *L1PoolManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerRoleRevoked)
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
		it.Event = new(L1PoolManagerRoleRevoked)
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
func (it *L1PoolManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerRoleRevoked represents a RoleRevoked event raised by the L1PoolManager contract.
type L1PoolManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L1PoolManager *L1PoolManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*L1PoolManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerRoleRevokedIterator{contract: _L1PoolManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_L1PoolManager *L1PoolManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *L1PoolManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerRoleRevoked)
				if err := _L1PoolManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseRoleRevoked(log types.Log) (*L1PoolManagerRoleRevoked, error) {
	event := new(L1PoolManagerRoleRevoked)
	if err := _L1PoolManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerSetMinStakeAmountEventIterator is returned from FilterSetMinStakeAmountEvent and is used to iterate over the raw logs and unpacked data for SetMinStakeAmountEvent events raised by the L1PoolManager contract.
type L1PoolManagerSetMinStakeAmountEventIterator struct {
	Event *L1PoolManagerSetMinStakeAmountEvent // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerSetMinStakeAmountEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerSetMinStakeAmountEvent)
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
		it.Event = new(L1PoolManagerSetMinStakeAmountEvent)
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
func (it *L1PoolManagerSetMinStakeAmountEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerSetMinStakeAmountEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerSetMinStakeAmountEvent represents a SetMinStakeAmountEvent event raised by the L1PoolManager contract.
type L1PoolManagerSetMinStakeAmountEvent struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMinStakeAmountEvent is a free log retrieval operation binding the contract event 0xf54d3b756d286b6b08e5d4eda6dfe5b135664abf029e58e637cbf013c442c950.
//
// Solidity: event SetMinStakeAmountEvent(address indexed token, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) FilterSetMinStakeAmountEvent(opts *bind.FilterOpts, token []common.Address) (*L1PoolManagerSetMinStakeAmountEventIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "SetMinStakeAmountEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerSetMinStakeAmountEventIterator{contract: _L1PoolManager.contract, event: "SetMinStakeAmountEvent", logs: logs, sub: sub}, nil
}

// WatchSetMinStakeAmountEvent is a free log subscription operation binding the contract event 0xf54d3b756d286b6b08e5d4eda6dfe5b135664abf029e58e637cbf013c442c950.
//
// Solidity: event SetMinStakeAmountEvent(address indexed token, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) WatchSetMinStakeAmountEvent(opts *bind.WatchOpts, sink chan<- *L1PoolManagerSetMinStakeAmountEvent, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "SetMinStakeAmountEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerSetMinStakeAmountEvent)
				if err := _L1PoolManager.contract.UnpackLog(event, "SetMinStakeAmountEvent", log); err != nil {
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

// ParseSetMinStakeAmountEvent is a log parse operation binding the contract event 0xf54d3b756d286b6b08e5d4eda6dfe5b135664abf029e58e637cbf013c442c950.
//
// Solidity: event SetMinStakeAmountEvent(address indexed token, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) ParseSetMinStakeAmountEvent(log types.Log) (*L1PoolManagerSetMinStakeAmountEvent, error) {
	event := new(L1PoolManagerSetMinStakeAmountEvent)
	if err := _L1PoolManager.contract.UnpackLog(event, "SetMinStakeAmountEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerSetSupportTokenEventIterator is returned from FilterSetSupportTokenEvent and is used to iterate over the raw logs and unpacked data for SetSupportTokenEvent events raised by the L1PoolManager contract.
type L1PoolManagerSetSupportTokenEventIterator struct {
	Event *L1PoolManagerSetSupportTokenEvent // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerSetSupportTokenEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerSetSupportTokenEvent)
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
		it.Event = new(L1PoolManagerSetSupportTokenEvent)
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
func (it *L1PoolManagerSetSupportTokenEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerSetSupportTokenEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerSetSupportTokenEvent represents a SetSupportTokenEvent event raised by the L1PoolManager contract.
type L1PoolManagerSetSupportTokenEvent struct {
	Token     common.Address
	IsSupport bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetSupportTokenEvent is a free log retrieval operation binding the contract event 0xc8c34f23fafb34e68119c1d231ef03d0d47225b15e2c4de3efbefa14b0181d86.
//
// Solidity: event SetSupportTokenEvent(address indexed token, bool isSupport)
func (_L1PoolManager *L1PoolManagerFilterer) FilterSetSupportTokenEvent(opts *bind.FilterOpts, token []common.Address) (*L1PoolManagerSetSupportTokenEventIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "SetSupportTokenEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerSetSupportTokenEventIterator{contract: _L1PoolManager.contract, event: "SetSupportTokenEvent", logs: logs, sub: sub}, nil
}

// WatchSetSupportTokenEvent is a free log subscription operation binding the contract event 0xc8c34f23fafb34e68119c1d231ef03d0d47225b15e2c4de3efbefa14b0181d86.
//
// Solidity: event SetSupportTokenEvent(address indexed token, bool isSupport)
func (_L1PoolManager *L1PoolManagerFilterer) WatchSetSupportTokenEvent(opts *bind.WatchOpts, sink chan<- *L1PoolManagerSetSupportTokenEvent, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "SetSupportTokenEvent", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerSetSupportTokenEvent)
				if err := _L1PoolManager.contract.UnpackLog(event, "SetSupportTokenEvent", log); err != nil {
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

// ParseSetSupportTokenEvent is a log parse operation binding the contract event 0xc8c34f23fafb34e68119c1d231ef03d0d47225b15e2c4de3efbefa14b0181d86.
//
// Solidity: event SetSupportTokenEvent(address indexed token, bool isSupport)
func (_L1PoolManager *L1PoolManagerFilterer) ParseSetSupportTokenEvent(log types.Log) (*L1PoolManagerSetSupportTokenEvent, error) {
	event := new(L1PoolManagerSetSupportTokenEvent)
	if err := _L1PoolManager.contract.UnpackLog(event, "SetSupportTokenEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerStakingETHEventIterator is returned from FilterStakingETHEvent and is used to iterate over the raw logs and unpacked data for StakingETHEvent events raised by the L1PoolManager contract.
type L1PoolManagerStakingETHEventIterator struct {
	Event *L1PoolManagerStakingETHEvent // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerStakingETHEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerStakingETHEvent)
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
		it.Event = new(L1PoolManagerStakingETHEvent)
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
func (it *L1PoolManagerStakingETHEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerStakingETHEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerStakingETHEvent represents a StakingETHEvent event raised by the L1PoolManager contract.
type L1PoolManagerStakingETHEvent struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakingETHEvent is a free log retrieval operation binding the contract event 0xe7466ea83435490635fc76a5f33da4815758ab48b1d45858f0452ca646556937.
//
// Solidity: event StakingETHEvent(address indexed user, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) FilterStakingETHEvent(opts *bind.FilterOpts, user []common.Address) (*L1PoolManagerStakingETHEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "StakingETHEvent", userRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerStakingETHEventIterator{contract: _L1PoolManager.contract, event: "StakingETHEvent", logs: logs, sub: sub}, nil
}

// WatchStakingETHEvent is a free log subscription operation binding the contract event 0xe7466ea83435490635fc76a5f33da4815758ab48b1d45858f0452ca646556937.
//
// Solidity: event StakingETHEvent(address indexed user, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) WatchStakingETHEvent(opts *bind.WatchOpts, sink chan<- *L1PoolManagerStakingETHEvent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "StakingETHEvent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerStakingETHEvent)
				if err := _L1PoolManager.contract.UnpackLog(event, "StakingETHEvent", log); err != nil {
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

// ParseStakingETHEvent is a log parse operation binding the contract event 0xe7466ea83435490635fc76a5f33da4815758ab48b1d45858f0452ca646556937.
//
// Solidity: event StakingETHEvent(address indexed user, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) ParseStakingETHEvent(log types.Log) (*L1PoolManagerStakingETHEvent, error) {
	event := new(L1PoolManagerStakingETHEvent)
	if err := _L1PoolManager.contract.UnpackLog(event, "StakingETHEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerStakingWETHEventIterator is returned from FilterStakingWETHEvent and is used to iterate over the raw logs and unpacked data for StakingWETHEvent events raised by the L1PoolManager contract.
type L1PoolManagerStakingWETHEventIterator struct {
	Event *L1PoolManagerStakingWETHEvent // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerStakingWETHEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerStakingWETHEvent)
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
		it.Event = new(L1PoolManagerStakingWETHEvent)
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
func (it *L1PoolManagerStakingWETHEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerStakingWETHEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerStakingWETHEvent represents a StakingWETHEvent event raised by the L1PoolManager contract.
type L1PoolManagerStakingWETHEvent struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakingWETHEvent is a free log retrieval operation binding the contract event 0xc138e3bd6d13eefa5cb01c0d35c5794001141efaf4e5ad888cad059935f83383.
//
// Solidity: event StakingWETHEvent(address indexed user, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) FilterStakingWETHEvent(opts *bind.FilterOpts, user []common.Address) (*L1PoolManagerStakingWETHEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "StakingWETHEvent", userRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerStakingWETHEventIterator{contract: _L1PoolManager.contract, event: "StakingWETHEvent", logs: logs, sub: sub}, nil
}

// WatchStakingWETHEvent is a free log subscription operation binding the contract event 0xc138e3bd6d13eefa5cb01c0d35c5794001141efaf4e5ad888cad059935f83383.
//
// Solidity: event StakingWETHEvent(address indexed user, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) WatchStakingWETHEvent(opts *bind.WatchOpts, sink chan<- *L1PoolManagerStakingWETHEvent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "StakingWETHEvent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerStakingWETHEvent)
				if err := _L1PoolManager.contract.UnpackLog(event, "StakingWETHEvent", log); err != nil {
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

// ParseStakingWETHEvent is a log parse operation binding the contract event 0xc138e3bd6d13eefa5cb01c0d35c5794001141efaf4e5ad888cad059935f83383.
//
// Solidity: event StakingWETHEvent(address indexed user, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) ParseStakingWETHEvent(log types.Log) (*L1PoolManagerStakingWETHEvent, error) {
	event := new(L1PoolManagerStakingWETHEvent)
	if err := _L1PoolManager.contract.UnpackLog(event, "StakingWETHEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerStarkingERC20EventIterator is returned from FilterStarkingERC20Event and is used to iterate over the raw logs and unpacked data for StarkingERC20Event events raised by the L1PoolManager contract.
type L1PoolManagerStarkingERC20EventIterator struct {
	Event *L1PoolManagerStarkingERC20Event // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerStarkingERC20EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerStarkingERC20Event)
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
		it.Event = new(L1PoolManagerStarkingERC20Event)
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
func (it *L1PoolManagerStarkingERC20EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerStarkingERC20EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerStarkingERC20Event represents a StarkingERC20Event event raised by the L1PoolManager contract.
type L1PoolManagerStarkingERC20Event struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStarkingERC20Event is a free log retrieval operation binding the contract event 0x3ac7d3823c11677fba9479ed26f696a3f17e16a5a5c39162fa6d183905aa6735.
//
// Solidity: event StarkingERC20Event(address indexed user, address indexed token, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) FilterStarkingERC20Event(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*L1PoolManagerStarkingERC20EventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "StarkingERC20Event", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerStarkingERC20EventIterator{contract: _L1PoolManager.contract, event: "StarkingERC20Event", logs: logs, sub: sub}, nil
}

// WatchStarkingERC20Event is a free log subscription operation binding the contract event 0x3ac7d3823c11677fba9479ed26f696a3f17e16a5a5c39162fa6d183905aa6735.
//
// Solidity: event StarkingERC20Event(address indexed user, address indexed token, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) WatchStarkingERC20Event(opts *bind.WatchOpts, sink chan<- *L1PoolManagerStarkingERC20Event, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "StarkingERC20Event", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerStarkingERC20Event)
				if err := _L1PoolManager.contract.UnpackLog(event, "StarkingERC20Event", log); err != nil {
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

// ParseStarkingERC20Event is a log parse operation binding the contract event 0x3ac7d3823c11677fba9479ed26f696a3f17e16a5a5c39162fa6d183905aa6735.
//
// Solidity: event StarkingERC20Event(address indexed user, address indexed token, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) ParseStarkingERC20Event(log types.Log) (*L1PoolManagerStarkingERC20Event, error) {
	event := new(L1PoolManagerStarkingERC20Event)
	if err := _L1PoolManager.contract.UnpackLog(event, "StarkingERC20Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerTransferAssertToIterator is returned from FilterTransferAssertTo and is used to iterate over the raw logs and unpacked data for TransferAssertTo events raised by the L1PoolManager contract.
type L1PoolManagerTransferAssertToIterator struct {
	Event *L1PoolManagerTransferAssertTo // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerTransferAssertToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerTransferAssertTo)
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
		it.Event = new(L1PoolManagerTransferAssertTo)
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
func (it *L1PoolManagerTransferAssertToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerTransferAssertToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerTransferAssertTo represents a TransferAssertTo event raised by the L1PoolManager contract.
type L1PoolManagerTransferAssertTo struct {
	Blockchain *big.Int
	Token      common.Address
	To         common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferAssertTo is a free log retrieval operation binding the contract event 0x08ffba656e4ac665e1aa4bb6a342e1a69d4cc12fd751ac862fa3fe27b0c7524a.
//
// Solidity: event TransferAssertTo(uint256 Blockchain, address indexed token, address indexed to, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) FilterTransferAssertTo(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*L1PoolManagerTransferAssertToIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "TransferAssertTo", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerTransferAssertToIterator{contract: _L1PoolManager.contract, event: "TransferAssertTo", logs: logs, sub: sub}, nil
}

// WatchTransferAssertTo is a free log subscription operation binding the contract event 0x08ffba656e4ac665e1aa4bb6a342e1a69d4cc12fd751ac862fa3fe27b0c7524a.
//
// Solidity: event TransferAssertTo(uint256 Blockchain, address indexed token, address indexed to, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) WatchTransferAssertTo(opts *bind.WatchOpts, sink chan<- *L1PoolManagerTransferAssertTo, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "TransferAssertTo", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerTransferAssertTo)
				if err := _L1PoolManager.contract.UnpackLog(event, "TransferAssertTo", log); err != nil {
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

// ParseTransferAssertTo is a log parse operation binding the contract event 0x08ffba656e4ac665e1aa4bb6a342e1a69d4cc12fd751ac862fa3fe27b0c7524a.
//
// Solidity: event TransferAssertTo(uint256 Blockchain, address indexed token, address indexed to, uint256 amount)
func (_L1PoolManager *L1PoolManagerFilterer) ParseTransferAssertTo(log types.Log) (*L1PoolManagerTransferAssertTo, error) {
	event := new(L1PoolManagerTransferAssertTo)
	if err := _L1PoolManager.contract.UnpackLog(event, "TransferAssertTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the L1PoolManager contract.
type L1PoolManagerUnpausedIterator struct {
	Event *L1PoolManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerUnpaused)
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
		it.Event = new(L1PoolManagerUnpaused)
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
func (it *L1PoolManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerUnpaused represents a Unpaused event raised by the L1PoolManager contract.
type L1PoolManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1PoolManager *L1PoolManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*L1PoolManagerUnpausedIterator, error) {

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerUnpausedIterator{contract: _L1PoolManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1PoolManager *L1PoolManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *L1PoolManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerUnpaused)
				if err := _L1PoolManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_L1PoolManager *L1PoolManagerFilterer) ParseUnpaused(log types.Log) (*L1PoolManagerUnpaused, error) {
	event := new(L1PoolManagerUnpaused)
	if err := _L1PoolManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1PoolManagerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the L1PoolManager contract.
type L1PoolManagerWithdrawIterator struct {
	Event *L1PoolManagerWithdraw // Event containing the contract specifics and raw log

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
func (it *L1PoolManagerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1PoolManagerWithdraw)
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
		it.Event = new(L1PoolManagerWithdraw)
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
func (it *L1PoolManagerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1PoolManagerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1PoolManagerWithdraw represents a Withdraw event raised by the L1PoolManager contract.
type L1PoolManagerWithdraw struct {
	User        common.Address
	StartPoolId *big.Int
	EndPoolId   *big.Int
	Token       common.Address
	Amount      *big.Int
	Reward      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x0ad30059d4072ae2525ebaaa5b6878f1c50df5f3b921615031c7d0f12dc4a77b.
//
// Solidity: event Withdraw(address _user, uint256 startPoolId, uint256 EndPoolId, address _token, uint256 Amount, uint256 Reward)
func (_L1PoolManager *L1PoolManagerFilterer) FilterWithdraw(opts *bind.FilterOpts) (*L1PoolManagerWithdrawIterator, error) {

	logs, sub, err := _L1PoolManager.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &L1PoolManagerWithdrawIterator{contract: _L1PoolManager.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x0ad30059d4072ae2525ebaaa5b6878f1c50df5f3b921615031c7d0f12dc4a77b.
//
// Solidity: event Withdraw(address _user, uint256 startPoolId, uint256 EndPoolId, address _token, uint256 Amount, uint256 Reward)
func (_L1PoolManager *L1PoolManagerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *L1PoolManagerWithdraw) (event.Subscription, error) {

	logs, sub, err := _L1PoolManager.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1PoolManagerWithdraw)
				if err := _L1PoolManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x0ad30059d4072ae2525ebaaa5b6878f1c50df5f3b921615031c7d0f12dc4a77b.
//
// Solidity: event Withdraw(address _user, uint256 startPoolId, uint256 EndPoolId, address _token, uint256 Amount, uint256 Reward)
func (_L1PoolManager *L1PoolManagerFilterer) ParseWithdraw(log types.Log) (*L1PoolManagerWithdraw, error) {
	event := new(L1PoolManagerWithdraw)
	if err := _L1PoolManager.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
