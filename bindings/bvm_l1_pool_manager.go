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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ChainIdIsNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"ChainIdNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorBlockChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"providedAmount\",\"type\":\"uint256\"}],\"name\":\"LessThanMinStakeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"MinTransferAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LessThanMinTransferAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LessThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MantaNotWETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MantleNotWETH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PoolIndex\",\"type\":\"uint256\"}],\"name\":\"NewPoolIsNotCreate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReward\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughETH\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"NotEnoughToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"PoolLength\",\"type\":\"uint256\"}],\"name\":\"OutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolIndex\",\"type\":\"uint256\"}],\"name\":\"PoolIsCompleted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"name\":\"TokenIsAlreadySupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"}],\"name\":\"TokenIsNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"sourceChainIdError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endPoolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Reward\",\"type\":\"uint256\"}],\"name\":\"ClaimReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolIndex\",\"type\":\"uint256\"}],\"name\":\"CompletePoolEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FinalizeWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InitiateWETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetMinStakeAmountEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSupport\",\"type\":\"bool\"}],\"name\":\"SetSupportTokenEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingETHEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingWETHEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StarkingERC20Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Blockchain\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferAssertTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Reward\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"BridgeFinalizeWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"BridgeInitiateETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"BridgeInitiateWETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ClaimAllReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"ClaimbyID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"IsCompleted\",\"type\":\"bool\"}],\"internalType\":\"structIL1PoolManager.Pool[]\",\"name\":\"CompletePools\",\"type\":\"tuple[]\"}],\"name\":\"CompletePoolAndNew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DepositAndStaking\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DepositAndStakingERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DepositAndStakingETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositAndStakingWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FeePoolValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"FundingPoolBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"IsSupportChainId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"IsSupportToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MinStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MinTransferAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PerFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Pools\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"IsCompleted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"QuickSendAssertToUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ReLayer\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"SupportTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"Blockchain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TransferAssertToBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Users\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isWithdrawed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"StartPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WithdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"WithdrawByID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getPool\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTimestamp\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TotalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFeeClaimed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"IsCompleted\",\"type\":\"bool\"}],\"internalType\":\"structIL1PoolManager.Pool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getPoolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isWithdrawed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"StartPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"}],\"internalType\":\"structIL1PoolManager.User[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isWithdrawed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"StartPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"EndPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"}],\"internalType\":\"structIL1PoolManager.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_MultisigWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageManager\",\"outputs\":[{\"internalType\":\"contractIMessageManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodTime\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setMinStakeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_MinTransferAmount\",\"type\":\"uint256\"}],\"name\":\"setMinTransferAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_PerFee\",\"type\":\"uint256\"}],\"name\":\"setPerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ERC20Address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setSupportERC20Token\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isSupport\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"startTimes\",\"type\":\"uint32\"}],\"name\":\"setSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setValidChainId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001c62000022565b620000d6565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000735760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d35780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6169ec80620000e66000396000f3fe60806040526004361061031a5760003560e01c80636f77926b116101ab578063cb314fab116100f7578063dd0c346011610095578063f8fcc2aa1161006f578063f8fcc2aa14610a41578063f91fa9a814610a77578063fa86184814610a97578063ff2bf64f14610aac57600080fd5b8063dd0c3460146109eb578063e207089314610a01578063f62a89cf14610a2157600080fd5b8063d547741f116100d1578063d547741f1461094e578063d73f14e51461096e578063dac295681461098e578063dbb0481f146109be57600080fd5b8063cb314fab146108d8578063cb4f04ad146108f8578063cd01c6651461092e57600080fd5b80639ab0b65211610164578063ab0f9e191161013e578063ab0f9e1914610863578063b1887e9b14610883578063b92e639614610898578063bc42493d146108b857600080fd5b80639ab0b652146108285780639b4423801461083b578063a217fddf1461084e57600080fd5b80636f77926b1461076e57806372fe6a7e1461079b5780638456cb59146107cb57806391d14854146107e057806392d847381461080057806396f984e61461082057600080fd5b806338731c471161026a57806354dc027e116102235780635b5b9ea2116101fd5780635b5b9ea2146106825780635c975abb14610709578063626417b51461072e578063650c22761461074e57600080fd5b806354dc027e146106205780635765a4eb1461064d5780635a648bc51461066d57600080fd5b806338731c471461052b5780633b0230f01461053e5780633f4ba83a1461055e5780634663cdc814610573578063485cc955146105e05780634dfaef841461060057600080fd5b8063248a9ca3116102d7578063325e5d43116102b1578063325e5d43146104b85780633338562c146104d857806334d34ef5146104eb57806336568abe1461050b57600080fd5b8063248a9ca3146104615780632c37388b146104815780632f2ff15d1461049857600080fd5b806301ffc9a71461031f57806309dc480c1461035457806310875a131461037857806313e8544e146103b05780631ca2f173146104025780631d31fac01461042f575b600080fd5b34801561032b57600080fd5b5061033f61033a366004615ea2565b610ad9565b60405190151581526020015b60405180910390f35b34801561036057600080fd5b5061036a60015481565b60405190815260200161034b565b34801561038457600080fd5b50600054610398906001600160a01b031681565b6040516001600160a01b03909116815260200161034b565b3480156103bc57600080fd5b506103d06103cb366004615ee3565b610b10565b6040805195151586526001600160a01b039094166020860152928401919091526060830152608082015260a00161034b565b34801561040e57600080fd5b5061036a61041d366004615f0d565b600b6020526000908152604090205481565b34801561043b57600080fd5b5060085461044c9063ffffffff1681565b60405163ffffffff909116815260200161034b565b34801561046d57600080fd5b5061036a61047c366004615f28565b610b6a565b34801561048d57600080fd5b50610496610b8c565b005b3480156104a457600080fd5b506104966104b3366004615f41565b610c06565b3480156104c457600080fd5b506104966104d3366004615f6d565b610c28565b6104966104e636600461603b565b610c88565b3480156104f757600080fd5b50610496610506366004615f28565b611003565b34801561051757600080fd5b50610496610526366004615f41565b61108d565b61033f610539366004616155565b6110c0565b34801561054a57600080fd5b5061033f6105593660046161a6565b61129e565b34801561056a57600080fd5b5061049661159c565b34801561057f57600080fd5b5061059361058e366004615ee3565b6115af565b6040805163ffffffff98891681529790961660208801526001600160a01b03909416948601949094526060850191909152608084015260a0830191909152151560c082015260e00161034b565b3480156105ec57600080fd5b506104966105fb3660046161f4565b611624565b34801561060c57600080fd5b5061049661061b366004615ee3565b611760565b34801561062c57600080fd5b5061036a61063b366004615f0d565b60066020526000908152604090205481565b34801561065957600080fd5b5061033f610668366004616155565b611a2b565b34801561067957600080fd5b50610496611c6c565b34801561068e57600080fd5b506106a261069d366004615ee3565b611cce565b60405161034b9190815163ffffffff9081168252602080840151909116908201526040808301516001600160a01b031690820152606080830151908201526080808301519082015260a0808301519082015260c09182015115159181019190915260e00190565b34801561071557600080fd5b506000805160206169778339815191525460ff1661033f565b34801561073a57600080fd5b50610398610749366004615f28565b611db2565b34801561075a57600080fd5b50610496610769366004615f28565b611ddc565b34801561077a57600080fd5b5061078e610789366004615f0d565b611ded565b60405161034b9190616256565b3480156107a757600080fd5b5061033f6107b6366004615f28565b60009081526003602052604090205460ff1690565b3480156107d757600080fd5b50610496611e9d565b3480156107ec57600080fd5b5061033f6107fb366004615f41565b611eb0565b34801561080c57600080fd5b5061033f61081b3660046162a4565b611ee8565b6104966120ed565b610496610836366004615ee3565b612428565b61033f610849366004616305565b612497565b34801561085a57600080fd5b5061036a600081565b34801561086f57600080fd5b5061049661087e36600461633a565b61269f565b34801561088f57600080fd5b50610398612729565b3480156108a457600080fd5b506104966108b3366004615f28565b612847565b3480156108c457600080fd5b506104966108d3366004615f28565b61289b565b3480156108e457600080fd5b506104966108f3366004615ee3565b612c7c565b34801561090457600080fd5b5061036a610913366004615f0d565b6001600160a01b03166000908152600a602052604090205490565b34801561093a57600080fd5b5061033f610949366004616371565b612ce1565b34801561095a57600080fd5b50610496610969366004615f41565b613072565b34801561097a57600080fd5b506104966109893660046163ae565b61308e565b34801561099a57600080fd5b5061033f6109a9366004615f0d565b60046020526000908152604090205460ff1681565b3480156109ca57600080fd5b5061036a6109d9366004615f0d565b60056020526000908152604090205481565b3480156109f757600080fd5b5061036a60025481565b348015610a0d57600080fd5b50610496610a1c3660046163d3565b6130ba565b348015610a2d57600080fd5b50610496610a3c366004615f28565b6133b5565b348015610a4d57600080fd5b5061036a610a5c366004615f0d565b6001600160a01b031660009081526009602052604090205490565b348015610a8357600080fd5b50610496610a9236600461640f565b61341c565b348015610aa357600080fd5b5061036a6135e3565b348015610ab857600080fd5b50610acc610ac7366004615ee3565b61362b565b60405161034b9190616443565b60006001600160e01b03198216637965db0b60e01b1480610b0a57506301ffc9a760e01b6001600160e01b03198316145b92915050565b600a6020528160005260406000208181548110610b2c57600080fd5b6000918252602090912060049091020180546001820154600283015460039093015460ff831695506101009092046001600160a01b03169350919085565b6000908152600080516020616957833981519152602052604090206001015490565b610b946136f4565b610b9c61372c565b60005b600754811015610bec57610bda3360078381548110610bc057610bc0616451565b60009182526020822001546001600160a01b03169061375d565b80610be48161647d565b915050610b9f565b50610c04600160008051602061699783398151915255565b565b610c0f82610b6a565b610c1881613e68565b610c228383613e72565b50505050565b60ff19610c4460016000805160206168f7833981519152616496565b604051602001610c5691815260200190565b6040516020818303038152906040528051906020012016610c7681613e68565b610c81848484613f1e565b5050505050565b60ff19610ca460016000805160206168f7833981519152616496565b604051602001610cb691815260200190565b6040516020818303038152906040528051906020012016610cd681613e68565b60005b8251811015610ffe576000838281518110610cf657610cf6616451565b60200260200101516040015190506000600160096000846001600160a01b03166001600160a01b0316815260200190815260200160002080549050610d3b9190616496565b6001600160a01b0383166000908152600960205260409020805491925060019183908110610d6b57610d6b616451565b906000526020600020906005020160040160006101000a81548160ff021916908315150217905550848381518110610da557610da5616451565b60200260200101516080015160096000846001600160a01b03166001600160a01b031681526020019081526020016000208281548110610de757610de7616451565b906000526020600020906005020160020181905550600060096000846001600160a01b03166001600160a01b031681526020019081526020016000208281548110610e3457610e34616451565b600091825260208083206005909202909101546001600160a01b038616835260098252604092839020835160e0810190945263ffffffff6401000000009092048216808552600854909550909392830191610e909116856164a9565b63ffffffff168152602001856001600160a01b0316815260200160096000876001600160a01b03166001600160a01b031681526020019081526020016000208581548110610ee057610ee0616451565b6000918252602080832060059283020160019081015485528482018490526040808601859052606095860185905287548083018955978552938290208651979093029092018054868301518786015163ffffffff998a1667ffffffffffffffff1990931692909217640100000000999091169890980297909717600160401b600160e01b031916600160401b6001600160a01b039889160217815593850151918401919091556080840151600284015560a0840151600384015560c0909301516004909201805460ff19169215159290921790915551848152918516917fb6f449f07ceaf55392c9899e0797c6529908ae827c2d498c682e90d42c241167910160405180910390a25050508080610ff69061647d565b915050610cd9565b505050565b61100b6136f4565b61101361372c565b336000908152600a6020526040902054811061106757336000908152600a60205260409081902054905163abe5c32f60e01b815261105e918391600401918252602082015260400190565b60405180910390fd5b611073338260006140bc565b61108a600160008051602061699783398151915255565b50565b6001600160a01b03811633146110b65760405163334bd91960e11b815260040160405180910390fd5b610ffe828261472a565b600060ff196110de60016000805160206168f7833981519152616496565b6040516020016110f091815260200190565b604051602081830303815290604052805190602001201661111081613e68565b46871461113057604051631a26aa4d60e21b815260040160405180910390fd5b60008881526003602052604090205460ff1661116257604051632ef6974160e11b81526004810189905260240161105e565b6040516001600160a01b0387169086156108fc029087906000818181858888f19350505050158015611198573d6000803e3d6000fd5b506000805160206169178339815191526000908152600560205260008051602061689783398151915280548792906111d1908490616496565b9091555050600054604051631eb65ffb60e01b81526001600160a01b0390911690631eb65ffb90611210908b908b908b908a908c908b906004016164c6565b600060405180830381600087803b15801561122a57600080fd5b505af115801561123e573d6000803e3d6000fd5b5050604080518b8152602081018b90529081018890526001600160a01b03891692503091507f61d99af9dd09bd984d42cc07939d4cba4388e85aa29b3cc88df954e4ca719f329060600160405180910390a3506001979650505050505050565b60004686146112c057604051631a26aa4d60e21b815260040160405180910390fd5b60008581526003602052604090205460ff166112f257604051632ef6974160e11b81526004810186905260240161105e565b6001600160a01b03831660009081526004602052604090205460ff16611336576040516305fd61ad60e01b81526001600160a01b038416600482015260240161105e565b6040516370a0823160e01b81523060048201526000906001600160a01b038516906370a0823190602401602060405180830381865afa15801561137d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113a191906164f7565b90506113b86001600160a01b0385163330866147a6565b6040516370a0823160e01b81523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa1580156113ff573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061142391906164f7565b905060006114318383616496565b600080516020616917833981519152600090815260056020526000805160206168978339815191528054929350879290919061146e908490616510565b9091555050600254600090620f4240906114889084616523565b611492919061653a565b905061149e8183616496565b6001600160a01b0388166000908152600660205260408120805492945083929091906114cb908490616510565b909155505060005460405163305f478560e21b81526001600160a01b039091169063c17d1e1490611508908d908d908d908c90889060040161655c565b600060405180830381600087803b15801561152257600080fd5b505af1158015611536573d6000803e3d6000fd5b5050604080518d8152602081018d90529081018590526001600160a01b03808c1693503392508a16907fece797608c81cc46a09f9859cf8ef650ec541b5da6989f30acb30dba9b8a40a49060600160405180910390a45060019998505050505050505050565b60006115a781613e68565b61108a614800565b600960205281600052604060002081815481106115cb57600080fd5b60009182526020909120600590910201805460018201546002830154600384015460049094015463ffffffff80851697506401000000008504169550600160401b9093046001600160a01b031693919290919060ff1687565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff1660008115801561166a5750825b905060008267ffffffffffffffff1660011480156116875750303b155b905081158015611695575080155b156116b35760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156116dd57845460ff60401b1916600160401b1785555b6116e5614860565b6116ed614868565b6116f5614878565b6116ff8787614888565b6008805463ffffffff1916621baf80179055831561175757845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b6117686136f4565b61177061372c565b6001600160a01b03821660009081526004602052604090205460ff166117b4576040516305fd61ad60e01b81526001600160a01b038316600482015260240161105e565b6001600160a01b0382166000908152600b6020526040902054811015611812576001600160a01b0382166000908152600b6020526040908190205490516327500c6d60e21b815260048101919091526024810182905260440161105e565b6118276001600160a01b0383163330846147a6565b6001600160a01b03821660009081526009602052604081205461184c90600190616496565b6001600160a01b0384166000908152600960205260409020805491925042918390811061187b5761187b616451565b600091825260209091206005909102015463ffffffff16111561198657336000908152600a60209081526040808320815160a0810183528481526001600160a01b0388811682860181815283860189815260608501898152608086018c8152875460018082018a55988c528a8c2097516004909102909701805494516001600160a81b0319909516971515610100600160a81b031916979097176101009490961693909302949094178555519484019490945590516002830155915160039091015583526009909152902080548391908390811061195b5761195b616451565b9060005260206000209060050201600101600082825461197b9190616510565b909155506119a29050565b604051637d58ebb960e01b81526004810182905260240161105e565b6001600160a01b038316600090815260056020526040812080548492906119ca908490616510565b90915550506040518281526001600160a01b0384169033907f3ac7d3823c11677fba9479ed26f696a3f17e16a5a5c39162fa6d183905aa67359060200160405180910390a350611a27600160008051602061699783398151915255565b5050565b600060ff19611a4960016000805160206168f7833981519152616496565b604051602001611a5b91815260200190565b6040516020818303038152906040528051906020012016611a7b81613e68565b468714611a9b57604051631a26aa4d60e21b815260040160405180910390fd5b60008881526003602052604090205460ff16611acd57604051632ef6974160e11b81526004810189905260240161105e565b6000611ad7612729565b6040516323b872dd60e01b81529091506001600160a01b038216906323b872dd90611b0a9030908b908b90600401616588565b6020604051808303816000875af1158015611b29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b4d91906165ac565b5073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2600090815260056020527fa550ba85c46b24b567d2e17cd597f2283877afab43603f46d5de7858f1bdb7318054889290611b9e908490616496565b9091555050600054604051631eb65ffb60e01b81526001600160a01b0390911690631eb65ffb90611bdd908c908c908c908b908d908c906004016164c6565b600060405180830381600087803b158015611bf757600080fd5b505af1158015611c0b573d6000803e3d6000fd5b5050604080518c8152602081018c90529081018990526001600160a01b038a1692503091507f7dc80be13e04000533f3763be4b9359a36a0aeb0daabbf09fd96c3a275c1cc149060600160405180910390a350600198975050505050505050565b611c746136f4565b611c7c61372c565b60005b600754811015610bec57611cbc3360078381548110611ca057611ca0616451565b6000918252602090912001546001600160a01b0316600161375d565b80611cc68161647d565b915050611c7f565b6040805160e08101825260008082526020808301829052828401829052606083018290526080830182905260a0830182905260c083018290526001600160a01b03861682526009905291909120805483908110611d2d57611d2d616451565b60009182526020918290206040805160e081018252600593909302909101805463ffffffff808216855264010000000082041694840194909452600160401b9093046001600160a01b0316908201526001820154606082015260028201546080820152600382015460a082015260049091015460ff16151560c0820152905092915050565b60078181548110611dc257600080fd5b6000918252602090912001546001600160a01b0316905081565b6000611de781613e68565b50600255565b6001600160a01b0381166000908152600a60209081526040808320805482518185028101850190935280835260609492939192909184015b82821015611e925760008481526020908190206040805160a08101825260048602909201805460ff81161515845261010090046001600160a01b03168385015260018082015492840192909252600281015460608401526003015460808301529083529092019101611e25565b505050509050919050565b6000611ea881613e68565b61108a6148d1565b6000918252600080516020616957833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b600060ff19611f0660016000805160206168f7833981519152616496565b604051602001611f1891815260200190565b6040516020818303038152906040528051906020012016611f3881613e68565b468814611f5857604051631a26aa4d60e21b815260040160405180910390fd5b60008981526003602052604090205460ff16611f8a57604051632ef6974160e11b8152600481018a905260240161105e565b6001600160a01b03861660009081526004602052604090205460ff16611fce576040516305fd61ad60e01b81526001600160a01b038716600482015260240161105e565b611fe36001600160a01b0387163089886147a6565b60008051602061691783398151915260009081526005602052600080516020616897833981519152805487929061201b908490616496565b9091555050600054604051631eb65ffb60e01b81526001600160a01b0390911690631eb65ffb9061205a908c908c908c908a908c908b906004016164c6565b600060405180830381600087803b15801561207457600080fd5b505af1158015612088573d6000803e3d6000fd5b5050604080518c8152602081018c90529081018890526001600160a01b03808b1693503092508916907f0f50ec03884adc78b9840a90cf8805a70ffe160f18d6b83198103b6fade443fe9060600160405180910390a450600198975050505050505050565b6120f56136f4565b6120fd61372c565b600080516020616917833981519152600052600b6020527fc81ad962d4d90b45a70da3e5e59847dcc28dc55db029f697d18929adae48413f5434101561219957600080516020616917833981519152600052600b6020527fc81ad962d4d90b45a70da3e5e59847dcc28dc55db029f697d18929adae48413f546040516327500c6d60e21b8152600481019190915234602482015260440161105e565b600080516020616917833981519152600090815260096020526000805160206168b7833981519152546121ce90600190616496565b600080516020616917833981519152600090815260096020526000805160206168b7833981519152549192500361221b57604051637d58ebb960e01b81526001600482015260240161105e565b60008051602061691783398151915260005260096020526000805160206168b783398151915280544291908390811061225657612256616451565b600091825260209091206005909102015463ffffffff16111561237757336000908152600a60209081526040808320815160a081018352848152600080516020616917833981519152818501818152938201878152606083018781523460808501818152865460018082018955978b52898b2096516004909102909601805498516001600160a81b0319909916961515610100600160a81b031916969096176101006001600160a01b03909916989098029790971785559151948401949094559251600283015592516003909101559252600990526000805160206168b783398151915280548390811061234c5761234c616451565b9060005260206000209060050201600101600082825461236c9190616510565b9091555061239e9050565b612382816001616510565b604051637d58ebb960e01b815260040161105e91815260200190565b6000805160206169178339815191526000908152600560205260008051602061689783398151915280543492906123d6908490616510565b909155505060405134815233907fe7466ea83435490635fc76a5f33da4815758ab48b1d45858f0452ca6465569379060200160405180910390a250610c04600160008051602061699783398151915255565b61243061372c565b341561243e57611a276120ed565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc1196001600160a01b0383160161246c57611a278161289b565b6001600160a01b03821660009081526004602052604090205460ff1615611a2757611a278282611760565b60004684146124b957604051631a26aa4d60e21b815260040160405180910390fd5b60008381526003602052604090205460ff166124eb57604051632ef6974160e11b81526004810184905260240161105e565b60015434101561251b576001546040516358f8331360e11b8152600481019190915234602482015260440161105e565b600080516020616917833981519152600090815260056020526000805160206168978339815191528054349290612553908490616510565b9091555050600254600090620f42409061256d9034616523565b612577919061653a565b905060006125858234616496565b600080516020616917833981519152600090815260066020527fa2e5aefc6e2cbe2917a296f0fd89c5f915c487c803db1d98eccb43f14012d711805492935084929091906125d4908490616510565b909155505060005460405163305f478560e21b81526001600160a01b039091169063c17d1e1490612611904690899089908790899060040161655c565b600060405180830381600087803b15801561262b57600080fd5b505af115801561263f573d6000803e3d6000fd5b505060408051898152602081018990529081018490526001600160a01b03871692503391507f0b671089787b6d29f730bc690214e6a7e28064fef5a3cdbb9b930a22fac01dc59060600160405180910390a36001925050505b9392505050565b60006126aa81613e68565b6001600160a01b0383166000908152600460205260409020805460ff19168315801591909117909155610ffe57600780546001810182556000919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880180546001600160a01b0385166001600160a01b0319909116179055505050565b60004662082750819003612745576004605360981b0191505090565b8061044d0361276957734f9a0e7fd2bf6067db6994cf12e4495df938e6e991505090565b80600a0361277f576006602160991b0191505090565b8061a4b1036127a3577382af49447d8a07e3bd95bd0d56f35241523fbab191505090565b8061a4ba036127c75773722e8bdd2ce80a4422e880164f2079488e11536591505090565b80610144036127eb57735aea5775959fbc2557cc8789bc1bf90a239d9a9191505090565b806113880361280d5760405163e31d668960e01b815260040160405180910390fd5b8060a90361282e576040516322c20c6960e11b815260040160405180910390fd5b604051639474dee160e01b815260040160405180910390fd5b60ff1961286360016000805160206168f7833981519152616496565b60405160200161287591815260200190565b604051602081830303815290604052805190602001201661289581613e68565b50600155565b6128a36136f4565b6128ab61372c565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2600052600b6020527ffc8cffbc59e5d552299bcf0a6a1c92c3958cd08bc38d18034b5c765b88fdfdfa548110156129405760008052600b6020527fdf7de25b7f1fd6d0b5205f0e18f1f35bd7b8d84cce336588d184533ce43a6f76546040516327500c6d60e21b815260048101919091526024810182905260440161105e565b6040516323b872dd60e01b815273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2906323b872dd9061297b90339030908690600401616588565b6020604051808303816000875af115801561299a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129be91906165ac565b5073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc260009081526009602052600080516020616937833981519152546129fa90600190616496565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2600052600960205260008051602061693783398151915280549192509082908110612a3c57612a3c616451565b600091825260209091206004600590920201015460ff1615612a74576040516311cf1b0760e31b81526004810182905260240161105e565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26000526009602052600080516020616937833981519152805442919083908110612ab557612ab5616451565b600091825260209091206005909102015463ffffffff161115612bda57336000908152600a60209081526040808320815160a08101835284815273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc281850181815293820187815260608301878152608084018a8152855460018181018855968a52888a2095516004909102909501805497516001600160a01b031661010002610100600160a81b0319961515969096166001600160a81b031990981697909717949094178655905193850193909355915160028401555160039092019190915590915260099052600080516020616937833981519152805483919083908110612bb457612bb4616451565b90600052602060002090600502016001016000828254612bd49190616510565b90915550505b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2600090815260056020527fa550ba85c46b24b567d2e17cd597f2283877afab43603f46d5de7858f1bdb7318054849290612c2a908490616510565b909155505060405182815233907fc138e3bd6d13eefa5cb01c0d35c5794001141efaf4e5ad888cad059935f833839060200160405180910390a25061108a600160008051602061699783398151915255565b6000612c8781613e68565b6001600160a01b0383166000818152600b602052604090819020849055517ff54d3b756d286b6b08e5d4eda6dfe5b135664abf029e58e637cbf013c442c95090612cd49085815260200190565b60405180910390a2505050565b6000468514612d0357604051631a26aa4d60e21b815260040160405180910390fd5b60008481526003602052604090205460ff16612d35576040516363b5c9db60e01b81526004810185905260240161105e565b6000612d3f612729565b6040516370a0823160e01b81523060048201529091506000906001600160a01b038316906370a0823190602401602060405180830381865afa158015612d89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dad91906164f7565b6040516323b872dd60e01b81529091506001600160a01b038316906323b872dd90612de090339030908990600401616588565b6020604051808303816000875af1158015612dff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e2391906165ac565b506040516370a0823160e01b81523060048201526000906001600160a01b038416906370a0823190602401602060405180830381865afa158015612e6b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e8f91906164f7565b90506000612e9d8383616496565b9050600154811015612ed0576001546040516358f8331360e11b815260048101919091526024810182905260440161105e565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2600090815260056020527fa550ba85c46b24b567d2e17cd597f2283877afab43603f46d5de7858f1bdb7318054839290612f20908490616510565b9091555050600254600090620f424090612f3a9084616523565b612f44919061653a565b9050612f508183616496565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2600090815260066020527f5e5777fab7622aff3c042c1ece74307c2e9d699a9da444f416c35f2e1def28a580549294508392909190612fa5908490616510565b909155505060005460405163305f478560e21b81526001600160a01b039091169063c17d1e1490612fe2908d908d908d908d90889060040161655c565b600060405180830381600087803b158015612ffc57600080fd5b505af1158015613010573d6000803e3d6000fd5b5050604080518d8152602081018d90529081018590526001600160a01b038b1692503391507fcb9e641636f35317739cff2833f1831bf3a25b930b5615ada09367c080d64a899060600160405180910390a35060019998505050505050505050565b61307b82610b6a565b61308481613e68565b610c22838361472a565b600061309981613e68565b50600091825260036020526040909120805460ff1916911515919091179055565b60006130c581613e68565b6001600160a01b03841660009081526004602052604090205460ff16156131125760405163411befff60e11b81526001600160a01b0385166004820152831515602482015260440161105e565b6001600160a01b0384166000908152600460209081526040808320805460ff1916871515179055600990915290819020815160e0810190925260085490919081906131639063ffffffff16866165c9565b63ffffffff908116825285811660208084018290526001600160a01b03808b166040808701829052600060608089018290526080808a0183905260a0808b0184905260c09a8b018490528c5460018181018f559d85528885208d5160059092020180548e8b01518f890151909a16600160401b02600160401b600160e01b03199a8e166401000000000267ffffffffffffffff19909216938e16939093171798909816178755918b01519b86019b909b5599890151600285015598880151600384015596909501516004909101805491151560ff199092169190911790559285526009835293839020835160e081019094529383526008549183019161326a9116866164a9565b63ffffffff90811682526001600160a01b038881166020808501829052600060408087018290526060808801839052608080890184905260a09889018490528a5460018082018d559b85528585208b5160059092020180548c8801518d870151938c1667ffffffffffffffff199092169190911764010000000091909b160299909917600160401b600160e01b031916600160401b9190981602969096178755880151868a015593870151600286015594860151600385015560c0909501516004909301805460ff1916931515939093179092556007805495860181559092527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68890930180546001600160a01b031916841790555185151581527fc8c34f23fafb34e68119c1d231ef03d0d47225b15e2c4de3efbefa14b0181d86910160405180910390a250505050565b6133bd6136f4565b6133c561372c565b336000908152600a6020526040902054811061341057336000908152600a60205260409081902054905163abe5c32f60e01b815261105e918391600401918252602082015260400190565b611073338260016140bc565b60ff1961343860016000805160206168f7833981519152616496565b60405160200161344a91815260200190565b604051602081830303815290604052805190602001201661346a81613e68565b6001600160a01b03841660009081526004602052604090205460ff166134ae576040516305fd61ad60e01b81526001600160a01b038516600482015260240161105e565b8462082750036134c8576134c384848461491a565b613566565b8461044d036134dc576134c3848484614d44565b84600a036134ef576134c3848484614eb3565b8461a4b103613503576134c3848484614fe0565b8461a4ba03613517576134c3848484615293565b846101440361352b576134c38484846154b8565b846113880361353f576134c3848484615668565b8460a903613552576134c384848461578e565b8461a70e0361282e576134c38484846158bb565b6001600160a01b0384166000908152600560205260408120805484929061358e908490616496565b909155505060408051868152602081018490526001600160a01b0380861692908716917f08ffba656e4ac665e1aa4bb6a342e1a69d4cc12fd751ac862fa3fe27b0c7524a910160405180910390a35050505050565b60ff196135ff60016000805160206168f7833981519152616496565b60405160200161361191815260200190565b604051602081830303815290604052805190602001201681565b6136686040518060a0016040528060001515815260200160006001600160a01b031681526020016000815260200160008152602001600081525090565b6001600160a01b0383166000908152600a6020526040902080548390811061369257613692616451565b60009182526020918290206040805160a081018252600493909302909101805460ff81161515845261010090046001600160a01b0316938301939093526001830154908201526002820154606082015260039091015460808201529392505050565b60008051602061699783398151915280546001190161372657604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6000805160206169778339815191525460ff1615610c045760405163d93c066560e01b815260040160405180910390fd5b6001600160a01b038216600090815260096020526040812054900361379857604051637d58ebb960e01b81526000600482015260240161105e565b60005b6001600160a01b0384166000908152600a6020526040902054811015610c22576001600160a01b038481166000908152600a602052604090208054839286169190839081106137ec576137ec616451565b600091825260209091206004909102015461010090046001600160a01b031603613e40576001600160a01b0385166000908152600a6020526040902080548290811061383a5761383a616451565b600091825260209091206004909102015460ff16156138595750613e42565b6001600160a01b03841660009081526009602052604081205461387e90600190616496565b6001600160a01b0387166000908152600a60205260409020805491925090839081106138ac576138ac616451565b90600052602060002090600402016003015460096000876001600160a01b03166001600160a01b0316815260200190815260200160002082815481106138f4576138f4616451565b906000526020600020906005020160010160008282546139149190616496565b90915550506001600160a01b0386166000908152600a6020526040812080548291908590811061394657613946616451565b90600052602060002090600402016003015490506000600a60008a6001600160a01b03166001600160a01b03168152602001908152602001600020858154811061399257613992616451565b9060005260206000209060040201600101549050838111156139c75760405163374c934360e11b815260040160405180910390fd5b805b84811015613b27576001600160a01b0389166000908152600960205260409020546139f690600190616496565b811115613a1957604051637d58ebb960e01b81526004810182905260240161105e565b6001600160a01b0389166000908152600960205260408120805483908110613a4357613a43616451565b906000526020600020906005020160010154600960008c6001600160a01b03166001600160a01b031681526020019081526020016000208381548110613a8b57613a8b616451565b90600052602060002090600502016002015485613aa89190616523565b613ab2919061653a565b9050613abe8186616510565b6001600160a01b038b1660009081526009602052604090208054919650829184908110613aed57613aed616451565b90600052602060002090600502016003016000828254613b0d9190616510565b90915550829150613b1f90508161647d565b9150506139c9565b5060008311613b645760405162461bcd60e51b8152602060048201526009602482015268139bc814995dd85c9960ba1b604482015260640161105e565b613b6e8383616510565b91508615613e3b576001600160a01b0389166000908152600a6020526040902080546001919087908110613ba457613ba4616451565b60009182526020909120600490910201805460ff1916911515919091179055613bce888a85613f1e565b506001600160a01b0389166000908152600a602052604090205460011015613d94576001600160a01b0389166000908152600a602052604090208054613c1690600190616496565b81548110613c2657613c26616451565b9060005260206000209060040201600a60008b6001600160a01b03166001600160a01b031681526020019081526020016000208681548110613c6a57613c6a616451565b600091825260208083208454600490930201805460ff909316151560ff1984168117825585546001600160a01b03610100918290048116909102610100600160a81b03199092166001600160a81b03199095169490941717815560018086015490820155600280860154908201556003948501549401939093558b168152600a90915260409020805480613d0057613d006165e6565b60008281526020812060046000199093019283020180546001600160a81b0319168155600181018290556002810182905560030155905585613d41816165fc565b96507f0ad30059d4072ae2525ebaaa5b6878f1c50df5f3b921615031c7d0f12dc4a77b90508982868b613d748888616496565b88604051613d8796959493929190616619565b60405180910390a1613e3b565b6001600160a01b0389166000908152600a60205260409020805485919087908110613dc157613dc1616451565b906000526020600020906004020160010181905550613de1888a84613f1e565b50604080516001600160a01b038b81168252602082018490528183018790528a1660608201526080810185905290517f7e6fedb6f824b2112ac0626ced32f6be6324d554cc03efd112aa021ee4d740ac9181900360a00190a15b505050505b505b80613e4c81616653565b91505061379b565b600160008051602061699783398151915255565b61108a81336159e3565b6000600080516020616957833981519152613e8d8484611eb0565b613f0d576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055613ec33390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610b0a565b6000915050610b0a565b5092915050565b6001600160a01b03831660009081526004602052604081205460ff16613f62576040516305fd61ad60e01b81526001600160a01b038516600482015260240161105e565b6001600160a01b03841660009081526005602052604081208054849290613f8a908490616496565b90915550506000805160206168d78339815191526001600160a01b0385160161400a5781471015613fce57604051632c1d501360e11b815260040160405180910390fd5b6040516001600160a01b0384169083156108fc029084906000818181858888f19350505050158015614004573d6000803e3d6000fd5b506140b2565b6040516370a0823160e01b815230600482015282906001600160a01b038616906370a0823190602401602060405180830381865afa158015614050573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061407491906164f7565b101561409e576040516311745b6160e21b81526001600160a01b038516600482015260240161105e565b6140b26001600160a01b0385168484615a0e565b5060019392505050565b6001600160a01b0383166000908152600a602052604081208054849081106140e6576140e6616451565b600091825260208083206004909202909101546001600160a01b03610100909104168083526009909152604082205490925061412490600190616496565b6001600160a01b0386166000908152600a602052604090208054919250908590811061415257614152616451565b90600052602060002090600402016003015460096000846001600160a01b03166001600160a01b03168152602001908152602001600020828154811061419a5761419a616451565b906000526020600020906005020160010160008282546141ba9190616496565b90915550506001600160a01b0385166000908152600a602052604081208054829190879081106141ec576141ec616451565b90600052602060002090600402016003015490506000600a6000896001600160a01b03166001600160a01b03168152602001908152602001600020878154811061423857614238616451565b90600052602060002090600402016001015490508381111561426d5760405163374c934360e11b815260040160405180910390fd5b805b848110156143cd576001600160a01b03861660009081526009602052604090205461429c90600190616496565b8111156142bf57604051637d58ebb960e01b81526004810182905260240161105e565b6001600160a01b03861660009081526009602052604081208054839081106142e9576142e9616451565b90600052602060002090600502016001015460096000896001600160a01b03166001600160a01b03168152602001908152602001600020838154811061433157614331616451565b9060005260206000209060050201600201548561434e9190616523565b614358919061653a565b90506143648186616510565b6001600160a01b0388166000908152600960205260409020805491965082918490811061439357614393616451565b906000526020600020906005020160030160008282546143b39190616510565b909155508291506143c590508161647d565b91505061426f565b506000831161440a5760405162461bcd60e51b8152602060048201526009602482015268139bc814995dd85c9960ba1b604482015260640161105e565b6144148383616510565b6001600160a01b0389166000908152600a6020526040902080549193506001918990811061444457614444616451565b60009182526020909120600490910201805460ff19169115159190911790558515614720576001600160a01b0388166000908152600a602052604090208054600191908990811061449757614497616451565b60009182526020909120600490910201805460ff19169115159190911790556144c1858985613f1e565b506001600160a01b0388166000908152600a602052604090205460011015614679576001600160a01b0388166000908152600a60205260409020805461450990600190616496565b8154811061451957614519616451565b9060005260206000209060040201600a60008a6001600160a01b03166001600160a01b03168152602001908152602001600020888154811061455d5761455d616451565b600091825260208083208454600490930201805460ff909316151560ff1984168117825585546001600160a01b03610100918290048116909102610100600160a81b03199092166001600160a81b03199095169490941717815560018086015490820155600280860154908201556003948501549401939093558a168152600a909152604090208054806145f3576145f36165e6565b60008281526020812060046000199093019283020180546001600160a81b031916815560018101829055600281018290556003015590557f0ad30059d4072ae2525ebaaa5b6878f1c50df5f3b921615031c7d0f12dc4a77b888286886146598888616496565b8860405161466c96959493929190616619565b60405180910390a1614720565b6001600160a01b0388166000908152600a602052604090208054859190899081106146a6576146a6616451565b9060005260206000209060040201600101819055506146c6858984613f1e565b50604080516001600160a01b038a8116825260208201849052818301879052871660608201526080810185905290517f7e6fedb6f824b2112ac0626ced32f6be6324d554cc03efd112aa021ee4d740ac9181900360a00190a15b5050505050505050565b60006000805160206169578339815191526147458484611eb0565b15613f0d576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610b0a565b610c2284856001600160a01b03166323b872dd8686866040516024016147ce93929190616588565b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050615a34565b614808615a97565b600080516020616977833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b610c04615ac7565b614870615ac7565b610c04615b10565b614880615ac7565b610c04615b18565b614890615ac7565b67016345785d8a00006001556127106002556148ad600083613e72565b50600080546001600160a01b0319166001600160a01b039290921691909117905550565b6148d961372c565b600080516020616977833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833614842565b6000805160206168d78339815191526001600160a01b03841601614a4557604051636bb825d760e11b8152620298106004820152600090730d7e906bd9cafa154b048cfa766cc1e54e39af9b9063d7704bae90602401602060405180830381865afa15801561498d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906149b191906164f7565b9050737f2b8c31f88b6006c382775eea88297ec1e3e90563ce0b63ce6149d78385616510565b6040516001600160e01b031960e084901b1681526001600160a01b0387166004820152602481018690526202981060448201526064016000604051808303818588803b158015614a2657600080fd5b505af1158015614a3a573d6000803e3d6000fd5b505050505050505050565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc1196001600160a01b03841601614bec57604051636bb825d760e11b8152614e206004820152600090730d7e906bd9cafa154b048cfa766cc1e54e39af9b9063d7704bae90602401602060405180830381865afa158015614abe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614ae291906164f7565b60405163095ea7b360e01b81529091506001600160a01b0385169063095ea7b390614b2790737ac440cae8eb6328de4fa621163a792c1ea9d4fe90869060040161666b565b6020604051808303816000875af1158015614b46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614b6a91906165ac565b5060405163790cfd3360e11b81526001600160a01b0380861660048301528416602482015260448101839052614e206064820152737ac440cae8eb6328de4fa621163a792c1ea9d4fe9063f219fa66906084015b600060405180830381600087803b158015614bd857600080fd5b505af1158015614720573d6000803e3d6000fd5b604051636bb825d760e11b8152614e206004820152600090730d7e906bd9cafa154b048cfa766cc1e54e39af9b9063d7704bae90602401602060405180830381865afa158015614c40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614c6491906164f7565b60405163095ea7b360e01b81529091506001600160a01b0385169063095ea7b390614ca990737ac440cae8eb6328de4fa621163a792c1ea9d4fe90869060040161666b565b6020604051808303816000875af1158015614cc8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614cec91906165ac565b5060405163790cfd3360e11b81526001600160a01b0380861660048301528416602482015260448101839052614e20606482015273d8a791fe2be73eb6e6cf1eb0cb3f36adc9b3f8f99063f219fa6690608401614bbe565b6000805160206168d78339815191526001600160a01b03841601614dbe5760405163cd58657960e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede9063cd586579908390614da590600190879084906000908190600401616684565b6000604051808303818588803b158015614bd857600080fd5b60405163095ea7b360e01b81526001600160a01b0384169063095ea7b390614e0090732a3dd3eb832af982ec71669e178424b10dca2ede90859060040161666b565b6020604051808303816000875af1158015614e1f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614e4391906165ac565b5060405163cd58657960e01b8152732a3dd3eb832af982ec71669e178424b10dca2ede9063cd58657990614e8590600190869086908990600090600401616684565b600060405180830381600087803b158015614e9f57600080fd5b505af1158015611757573d6000803e3d6000fd5b6000805160206168d78339815191526001600160a01b03841601614f0d57604051639a2ac6d560e01b81527399c9fc46f92e8a1c0dec1b1747d010903e884be190639a2ac6d5908390614da59086906000906004016166cb565b6000614f1884615b39565b60405163095ea7b360e01b81529091506001600160a01b0385169063095ea7b390614f5d907399c9fc46f92e8a1c0dec1b1747d010903e884be190869060040161666b565b6020604051808303816000875af1158015614f7c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614fa091906165ac565b507399c9fc46f92e8a1c0dec1b1747d010903e884be163838b2520858386865a6040518663ffffffff1660e01b8152600401614bbe9594939291906166f9565b6000805160206168d78339815191526001600160a01b0384160161509757604051634fb1a07b60e01b81527372ce9c846789fdb6fc1f34ac4ad25dd9ef7031ef90634fb1a07b90839061505090600080516020616917833981519152903090889085906000908190600401616741565b60006040518083038185885af115801561506e573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f19168201604052610c2291908101906167ae565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc1196001600160a01b038416016151cb5760405163095ea7b360e01b81526001600160a01b0384169063095ea7b3906150fe9073d92023e9d9911199a6711321d1277285e6d4e2db90859060040161666b565b6020604051808303816000875af115801561511d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061514191906165ac565b50604051634fb1a07b60e01b815273d92023e9d9911199a6711321d1277285e6d4e2db90634fb1a07b906151849086903090879087906000908190600401616741565b6000604051808303816000875af11580156151a3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c2291908101906167ae565b60405163095ea7b360e01b81526001600160a01b0384169063095ea7b39061520d9073a3a7b6f88361f48403514059f1f16c8e78d60eec90859060040161666b565b6020604051808303816000875af115801561522c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061525091906165ac565b50604051634fb1a07b60e01b815273a3a7b6f88361f48403514059f1f16c8e78d60eec90634fb1a07b906151849086903090879087906000908190600401616741565b6000805160206168d78339815191526001600160a01b0384160161530357604051634fb1a07b60e01b815273c840838bc438d73c16c2f8b22d2ce3669963cd4890634fb1a07b90839061505090600080516020616917833981519152903090889085906000908190600401616741565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc1196001600160a01b038416016153f05760405163095ea7b360e01b81526001600160a01b0384169063095ea7b39061536a9073e4e2121b479017955be0b175305b35f312330bae90859060040161666b565b6020604051808303816000875af1158015615389573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906153ad91906165ac565b50604051634fb1a07b60e01b815273e4e2121b479017955be0b175305b35f312330bae90634fb1a07b906151849086903090879087906000908190600401616741565b60405163095ea7b360e01b81526001600160a01b0384169063095ea7b3906154329073b2535b988dce19f9d71dfb22db6da744acac21bf90859060040161666b565b6020604051808303816000875af1158015615451573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061547591906165ac565b50604051634fb1a07b60e01b815273b2535b988dce19f9d71dfb22db6da744acac21bf90634fb1a07b906151849086903090879087906000908190600401616741565b6000805160206168d78339815191526001600160a01b0384160161555d5760405163e8b99b1b60e01b81527357891966931eb4bb6fb81430e6ce0a03aabde0639063e8b99b1b90839061551a9086906000908490829081903090600401616842565b60206040518083038185885af1158015615538573d6000803e3d6000fd5b50505050506040513d601f19601f82011682018060405250810190610c2291906164f7565b60405163095ea7b360e01b81526001600160a01b0384169063095ea7b39061559f907357891966931eb4bb6fb81430e6ce0a03aabde06390859060040161666b565b6020604051808303816000875af11580156155be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906155e291906165ac565b5060405163e8b99b1b60e01b81527357891966931eb4bb6fb81430e6ce0a03aabde0639063e8b99b1b906156259085908790869060009081903090600401616842565b6020604051808303816000875af1158015615644573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c2291906164f7565b6000805160206168d78339815191526001600160a01b038416016156c057604051639a2ac6d560e01b81527395fc37a27a2f68e3a647cdc081f0a89bb47c301290639a2ac6d590614e859085906000906004016166cb565b60405163095ea7b360e01b81526001600160a01b0384169063095ea7b390615702907395fc37a27a2f68e3a647cdc081f0a89bb47c301290859060040161666b565b6020604051808303816000875af1158015615721573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061574591906165ac565b507395fc37a27a2f68e3a647cdc081f0a89bb47c301263838b25208461576a81615c57565b858560006040518663ffffffff1660e01b8152600401614e859594939291906166f9565b6000805160206168d78339815191526001600160a01b038416016157e857604051639a2ac6d560e01b8152733b95bc951ee0f553ba487327278cac44f29715e590639a2ac6d5908390614da59086906000906004016166cb565b60006157f384615cd9565b60405163095ea7b360e01b81529091506001600160a01b0385169063095ea7b39061583890733b95bc951ee0f553ba487327278cac44f29715e590869060040161666b565b6020604051808303816000875af1158015615857573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061587b91906165ac565b50733b95bc951ee0f553ba487327278cac44f29715e563838b2520858386865a6040518663ffffffff1660e01b8152600401614bbe9594939291906166f9565b6000805160206168d78339815191526001600160a01b0384160161591c5760405163cd58657960e01b8152739cb4706e20a18e59a48ffa7616d700a3891e18619063cd586579908390614da590600190879084906000908190600401616684565b60405163095ea7b360e01b81526001600160a01b0384169063095ea7b39061595e90739cb4706e20a18e59a48ffa7616d700a3891e186190859060040161666b565b6020604051808303816000875af115801561597d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906159a191906165ac565b5060405163cd58657960e01b8152739cb4706e20a18e59a48ffa7616d700a3891e18619063cd58657990614e8590600190869086908990600090600401616684565b6159ed8282611eb0565b611a2757808260405163e2517d3f60e01b815260040161105e92919061666b565b610ffe83846001600160a01b031663a9059cbb85856040516024016147ce92919061666b565b6000615a496001600160a01b03841683615d9b565b90508051600014158015615a6e575080806020019051810190615a6c91906165ac565b155b15610ffe57604051635274afe760e01b81526001600160a01b038416600482015260240161105e565b6000805160206169778339815191525460ff16610c0457604051638dfc202b60e01b815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610c0457604051631afcd79f60e31b815260040160405180910390fd5b613e54615ac7565b615b20615ac7565b600080516020616977833981519152805460ff19169055565b600073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc1196001600160a01b03831601615b6e57506006602160991b01919050565b73dac17f958d2ee523a2206206994597c13d831ec6196001600160a01b03831601615bae57507394b008aa00579c1307b0ef2c499ad98a8ce58e58919050565b73a0b86991c6218b36c1d19d4a2e9eb0ce3606eb47196001600160a01b03831601615bee5750730b2c639c533813f4aa9d7837caf62653d097ff85919050565b736b175474e89094c44da98b954eedeac495271d0e196001600160a01b03831601615c2e575073da10009cbd5d07dd0cecc66161fc93d7c9000da1919050565b6040516305fd61ad60e01b81526001600160a01b038316600482015260240161105e565b919050565b600073dac17f958d2ee523a2206206994597c13d831ec6196001600160a01b03831601615c99575073201eba5cc46d216ce6dc03f6a759e8e766e956ae919050565b73a0b86991c6218b36c1d19d4a2e9eb0ce3606eb47196001600160a01b03831601615c2e575073201eba5cc46d216ce6dc03f6a759e8e766e956ae919050565b600073dac17f958d2ee523a2206206994597c13d831ec6196001600160a01b03831601615d1b575073f417f5a458ec102b90352f697d6e2ac3a3d2851f919050565b73a0b86991c6218b36c1d19d4a2e9eb0ce3606eb47196001600160a01b03831601615d5b575073b73603c5d87fa094b7314c74ace2e64d165016fb919050565b736b175474e89094c44da98b954eedeac495271d0e196001600160a01b03831601615c2e5750731c466b9371f8aba0d7c458be10a62192fcb8aa71919050565b60606126988383600084600080856001600160a01b03168486604051615dc1919061687a565b60006040518083038185875af1925050503d8060008114615dfe576040519150601f19603f3d011682016040523d82523d6000602084013e615e03565b606091505b5091509150615e13868383615e1d565b9695505050505050565b606082615e3257615e2d82615e79565b612698565b8151158015615e4957506001600160a01b0384163b155b15615e7257604051639996b31560e01b81526001600160a01b038516600482015260240161105e565b5080612698565b805115615e895780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b600060208284031215615eb457600080fd5b81356001600160e01b03198116811461269857600080fd5b80356001600160a01b0381168114615c5257600080fd5b60008060408385031215615ef657600080fd5b615eff83615ecc565b946020939093013593505050565b600060208284031215615f1f57600080fd5b61269882615ecc565b600060208284031215615f3a57600080fd5b5035919050565b60008060408385031215615f5457600080fd5b82359150615f6460208401615ecc565b90509250929050565b600080600060608486031215615f8257600080fd5b615f8b84615ecc565b9250615f9960208501615ecc565b9150604084013590509250925092565b634e487b7160e01b600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715615fe257615fe2615fa9565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561601157616011615fa9565b604052919050565b803563ffffffff81168114615c5257600080fd5b801515811461108a57600080fd5b6000602080838503121561604e57600080fd5b823567ffffffffffffffff8082111561606657600080fd5b818501915085601f83011261607a57600080fd5b81358181111561608c5761608c615fa9565b61609a848260051b01615fe8565b818152848101925060e09182028401850191888311156160b957600080fd5b938501935b828510156161495780858a0312156160d65760008081fd5b6160de615fbf565b6160e786616019565b81526160f4878701616019565b878201526040616105818801615ecc565b90820152606086810135908201526080808701359082015260a0808701359082015260c0808701356161368161602d565b90820152845293840193928501926160be565b50979650505050505050565b60008060008060008060c0878903121561616e57600080fd5b863595506020870135945061618560408801615ecc565b9350606087013592506080870135915060a087013590509295509295509295565b600080600080600060a086880312156161be57600080fd5b85359450602086013593506161d560408701615ecc565b92506161e360608701615ecc565b949793965091946080013592915050565b6000806040838503121561620757600080fd5b61621083615ecc565b9150615f6460208401615ecc565b8051151582526020808201516001600160a01b0316908301526040808201519083015260608082015190830152608090810151910152565b6020808252825182820181905260009190848201906040850190845b818110156162985761628583855161621e565b9284019260a09290920191600101616272565b50909695505050505050565b600080600080600080600060e0888a0312156162bf57600080fd5b87359650602088013595506162d660408901615ecc565b94506162e460608901615ecc565b9699959850939660808101359560a0820135955060c0909101359350915050565b60008060006060848603121561631a57600080fd5b833592506020840135915061633160408501615ecc565b90509250925092565b6000806040838503121561634d57600080fd5b61635683615ecc565b915060208301356163668161602d565b809150509250929050565b6000806000806080858703121561638757600080fd5b843593506020850135925061639e60408601615ecc565b9396929550929360600135925050565b600080604083850312156163c157600080fd5b8235915060208301356163668161602d565b6000806000606084860312156163e857600080fd5b6163f184615ecc565b925060208401356164018161602d565b915061633160408501616019565b6000806000806080858703121561642557600080fd5b8435935061643560208601615ecc565b925061639e60408601615ecc565b60a08101610b0a828461621e565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161648f5761648f616467565b5060010190565b81810381811115610b0a57610b0a616467565b63ffffffff818116838216019080821115613f1757613f17616467565b95865260208601949094526001600160a01b039290921660408501526060840152608083015260a082015260c00190565b60006020828403121561650957600080fd5b5051919050565b80820180821115610b0a57610b0a616467565b8082028115828204841417610b0a57610b0a616467565b60008261655757634e487b7160e01b600052601260045260246000fd5b500490565b94855260208501939093526001600160a01b039190911660408401526060830152608082015260a00190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6000602082840312156165be57600080fd5b81516126988161602d565b63ffffffff828116828216039080821115613f1757613f17616467565b634e487b7160e01b600052603160045260246000fd5b6000600160ff1b820161661157616611616467565b506000190190565b6001600160a01b0396871681526020810195909552604085019390935293166060830152608082019290925260a081019190915260c00190565b60006001600160ff1b01820161648f5761648f616467565b6001600160a01b03929092168252602082015260400190565b63ffffffff9590951685526001600160a01b039384166020860152604085019290925290911660608301521515608082015260c060a0820181905260009082015260e00190565b6001600160a01b0392909216825263ffffffff16602082015260606040820181905260009082015260800190565b6001600160a01b0395861681529385166020850152919093166040830152606082019290925263ffffffff909116608082015260c060a0820181905260009082015260e00190565b6001600160a01b03968716815294861660208601529290941660408401526060830152608082019290925260a081019190915260e060c082018190526000908201526101000190565b60005b838110156167a557818101518382015260200161678d565b50506000910152565b6000602082840312156167c057600080fd5b815167ffffffffffffffff808211156167d857600080fd5b818401915084601f8301126167ec57600080fd5b8151818111156167fe576167fe615fa9565b616811601f8201601f1916602001615fe8565b915080825285602082850101111561682857600080fd5b61683981602084016020860161678a565b50949350505050565b6001600160a01b039687168152948616602086015260408501939093526060840191909152608083015290911660a082015260c00190565b6000825161688c81846020870161678a565b919091019291505056fea1829a9003092132f585b6ccdd167c19fe9774dbdea4260287e8a8e8ca8185d700a206b70d87cacff7a19f3d98e0957b781c25fb4ae954252c87bcf406261742ffffffffffffffffffffffff111111111111111111111111111111111111111233fe247d78fee2e7fd135c405eda4bd2911c0a73c0a81b36c3bcc967dd06e5ae000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeea4adb76af136dce1635a653828507074afe3128b2c1b9d295c78c04cdae4286e02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a264697066735822122015d0313aae4db6007c52626622b91e719940c45334a038ad4d8b09e8fa350bd864736f6c63430008140033",
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
