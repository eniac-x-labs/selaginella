// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// IDETHBatchMint is an auto generated low-level Go binding around an user-defined struct.
type IDETHBatchMint struct {
	Staker common.Address
	Amount *big.Int
}

// StakingManagerInit is an auto generated low-level Go binding around an user-defined struct.
type StakingManagerInit struct {
	Admin            common.Address
	Manager          common.Address
	AllocatorService common.Address
	InitiatorService common.Address
	WithdrawalWallet common.Address
}

// StakingManagerStorageValidatorParams is an auto generated low-level Go binding around an user-defined struct.
type StakingManagerStorageValidatorParams struct {
	OperatorID            *big.Int
	DepositAmount         *big.Int
	Pubkey                []byte
	WithdrawalCredentials []byte
	Signature             []byte
	DepositDataRoot       [32]byte
}

// StakingManagerMetaData contains all meta data concerning the StakingManager contract.
var StakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DoesNotReceiveETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfiguration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"InvalidDepositRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes12\",\"name\":\"\",\"type\":\"bytes12\"}],\"name\":\"InvalidWithdrawalCredentialsNotETH1\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"InvalidWithdrawalCredentialsWrongAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"InvalidWithdrawalCredentialsWrongLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MathOverflowedMulDiv\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaximumDETHSupplyExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaximumValidatorDepositExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinimumDepositAmountNotSatisfied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinimumStakeBoundNotSatisfied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinimumUnstakeBoundNotSatisfied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinimumValidatorDepositNotSatisfied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotDappLinkBridge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughDepositETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughUnallocatedETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotReturnsAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotUnstakeRequestsManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreviouslyUsedValidator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedMinimum\",\"type\":\"uint256\"}],\"name\":\"UnstakeBelowMinimudETHAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AllocatedETHToDeposits\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AllocatedETHToUnstakeRequestsManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"setterSelector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"setterSignature\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"ProtocolConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReturnsReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dETHAmount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dETHLocked\",\"type\":\"uint256\"}],\"name\":\"UnstakeLaveAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Strategy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"}],\"name\":\"UnstakeRequestClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dETHLocked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"}],\"name\":\"UnstakeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"operatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountDeposited\",\"type\":\"uint256\"}],\"name\":\"ValidatorInitiated\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"ALLOCATOR_SERVICE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIATOR_SERVICE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_ALLOWLIST_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_ALLOWLIST_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOP_UP_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"allocateToUnstakeRequestsManager\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocateToDeposits\",\"type\":\"uint256\"}],\"name\":\"allocateETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocatedETHForDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"claimUnstakeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dETHAmount\",\"type\":\"uint256\"}],\"name\":\"dETHToETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"}],\"name\":\"ethToDETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeAdjustmentRate\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializationBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allocatorService\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initiatorService\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawalWallet\",\"type\":\"address\"}],\"internalType\":\"structStakingManager.Init\",\"name\":\"init\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"operatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"withdrawalCredentials\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"depositDataRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structStakingManagerStorage.ValidatorParams[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"expectedDepositRoot\",\"type\":\"bytes32\"}],\"name\":\"initiateValidatorsWithDeposits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStakingAllowlist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"locator\",\"outputs\":[{\"internalType\":\"contractL1ILocator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maximumDETHSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maximumDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumUnstakeBound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numInitiatedValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveFromUnstakeRequestsManager\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveReturns\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reclaimAllocatedETHSurplus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"exchangeAdjustmentRate_\",\"type\":\"uint16\"}],\"name\":\"setExchangeAdjustmentRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_locator\",\"type\":\"address\"}],\"name\":\"setLocator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maximumDETHSupply_\",\"type\":\"uint256\"}],\"name\":\"setMaximumDETHSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumDepositAmount_\",\"type\":\"uint256\"}],\"name\":\"setMinimumDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumUnstakeBound_\",\"type\":\"uint256\"}],\"name\":\"setMinimumUnstakeBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isStakingAllowlist_\",\"type\":\"bool\"}],\"name\":\"setStakingAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"withdrawalWallet_\",\"type\":\"address\"}],\"name\":\"setWithdrawalWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIDETH.BatchMint[]\",\"name\":\"batchMints\",\"type\":\"tuple[]\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"topUp\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalControlled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDepositedInValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unallocatedETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"dethAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"minETHAmount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"l2Strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"}],\"name\":\"unstakeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2strategy\",\"type\":\"address\"}],\"name\":\"unstakeRequestInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"usedValidators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608080604052346100b9577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff8260401c166100aa57506001600160401b036002600160401b031982821601610065575b6040516132239081620000bf8239f35b6001600160401b031990911681179091556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a1388080610055565b63f92ee8a960e01b8152600490fd5b600080fdfe6080604081815260049182361015610033575b50361561002657516334352c7360e01b8152fd5b516334352c7360e01b8152fd5b600090813560e01c90816301ffc9a714611dcb575080630208e4b514611d1757806304f36cc214611c615780630633af7614611c3e578063080c279a14611c1f57806312e9ead614611942578063147d36d5146119235780631943190d146118af57806319efd5c7146118865780631d2d35ce146117f1578063248a9ca3146117b957806329d48704146116cb5780632f2ff15d146116435780633101d9101461161a57806335ead2a4146115fb57806336568abe146115b557806337a6c881146112845780633937c0b31461125b5780633ccfd60b1461123957806342d3915d1461121257806349336f0f146111f35780634a7d80b3146111ca5780634bd6fc0814610fe75780635915ded114610f765780635940d90b14610f5957806360a0f62814610f3a5780636daa01a214610d295780636fce8ab214610cee57806375796f7614610c1e57806378abb49b14610bff5780637c957fc814610bd75780637dfcdd2914610bb8578063808d663f14610b0557806389e80ed314610aca5780638f656d22146107b85780639010d07c1461076a57806391d148541461071757806399dd1deb14610682578063a217fddf14610667578063a5e8456214610631578063aab483d614610596578063ac1e2257146104df578063b91590b2146104c0578063bb635c65146104a1578063c151aa7214610406578063ca15c873146103d1578063d547741f14610384578063dc29f1de146102f3578063e55d6cc0146102b8578063ea452b6d146102955763ed9daafb036100125734610292576020366003190112610292575061028b60209235612810565b9051908152f35b80fd5b5090346102b457816003193601126102b4576020906037549051908152f35b5080fd5b5090346102b457816003193601126102b457602090517f8ea5b4dbd68db0bf23bf4cda958b61a749f8c5aec6f2912d75a03246753ddd168152f35b50919082600319360112610380577f5e4bd437d29fad01c10cdcfff414f0d6b0e84b96d2dade88d780d45b5630696b908160005260008051602061316e833981519152602052806000203360005260205260ff81600020541615610364578361035e34603654611fcc565b60365580f35b60449350519163e2517d3f60e01b835233908301526024820152fd5b8280fd5b509190346103805780600319360112610380576103cd91356103c860016103a9611e4f565b9383875260008051602061316e8339815191526020528620015461295d565b612b41565b5080f35b50913461038057602036600319011261038057602092829135815260008051602061312e833981519152845220549051908152f35b50919082600319360112610380578254815163352a8adf60e21b81526001600160a01b039160209082908690829086165afa908115610497578591610468575b5016330361045b578261035e34603654611fcc565b51637154fc4360e01b8152fd5b61048a915060203d602011610490575b6104828183611ebf565b810190612564565b38610446565b503d610478565b83513d87823e3d90fd5b5090346102b457816003193601126102b4576020906035549051908152f35b5090346102b457816003193601126102b457602090603d549051908152f35b5082346102b457826003193601126102b457826104fa611e4f565b6001600160a01b03926044908461050f612b07565b168451958694859363d4be074f60e01b85528035908501521660248301525afa90811561058c5782809261054e575b5050825191151582526020820152f35b915091508282813d8311610585575b6105678183611ebf565b810103126102925750602061057b82611f55565b910151838061053e565b503d61055d565b83513d84823e3d90fd5b509134610380576020366003190112610380576000805160206131ae8339815191529035916105c36128f5565b82603a558051926020840152602083526105dc83611ea4565b60208151918083528201527f7365744d696e696d756d4465706f736974416d6f756e742875696e74323536296060820152608060208201528061062b63555a41eb60e11b94608083019061241b565b0390a280f35b5034610292576020366003190112610292576001600160a01b03610653611e65565b82546001600160a01b031916911617815580f35b5090346102b457816003193601126102b45751908152602090f35b509134610380576020366003190112610380576000805160206131ae8339815191529035916106af6128f5565b826038558051926020840152602083526106c883611ea4565b601f8151918083528201527f7365744d696e696d756d556e7374616b65426f756e642875696e7432353629006060820152608060208201528061062b6399dd1deb60e01b94608083019061241b565b50913461038057816003193601126103805781602093610735611e4f565b9235815260008051602061316e8339815191528552209060018060a01b0316600052825260ff81600020541690519015158152f35b5091346103805781600319360112610380576020926107a29135815260008051602061312e8339815191528452826024359120612f18565b905491519160018060a01b039160031b1c168152f35b5091346103805760a03660031901126103805781516001600160401b039060a0810182811182821017610ab55784526107ef611e65565b81526107f9611e4f565b9060208101918252610809611e39565b8186019081526001600160a01b039260643591908483168303610ab05760608401928352608435938585168503610ab057608081019485527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009788549060ff828c1c161598821680159081610aa8575b6001149081610a9e575b159081610a95575b50610a87575086949391928b86809581948c8e60016001600160401b03198416179055610a68575b5051166108be612ed7565b6108c6612ed7565b8c6108d082612b89565b610a40575b50505051166108e381612c17565b610a0a575b5051166108f481612caf565b6109d4575b50511661090581612d47565b61099e575b505116603c5490662386f26fc100006038556801bc16d674ec800000603a55600160a01b916affffffffffffffffffffff60a81b161717603c5543603d55683782dace9d90000000603e5561095d578280f35b805468ff00000000000000001916905551600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a138808280f35b6109cd9060008051602061314e833981519152885260008051602061312e833981519152602052868820612f30565b503861090a565b610a039060008051602061318e8339815191528a5260008051602061312e833981519152602052888a20612f30565b50386108f9565b610a39906000805160206131ce8339815191528c5260008051602061312e8339815191526020528a8c20612f30565b50386108e8565b8280610a5f945260008051602061312e83398151915260205220612f30565b508b388c6108d5565b68ffffffffffffffffff191668010000000000000001178d55386108b3565b8a5163f92ee8a960e01b8152fd5b9050153861088b565b303b159150610883565b8a9150610879565b600080fd5b604184634e487b7160e01b6000525260246000fd5b5090346102b457816003193601126102b457602090517fdec9d30de0821ad67aa5b141b13a539f584a19f99319e6041698a892b0e795598152f35b5091826003193601126103805782548251631faa859d60e11b81526001600160a01b039160209082908590829086165afa908115610bae578591610b8f575b50163303610b82575060207f4cbb9d73b003a252cee3f2ee51d8d65a562af35eebb23730dd4a76d68127b3709151348152a161035e34603654611fcc565b9051626310df60e31b8152fd5b610ba8915060203d602011610490576104828183611ebf565b38610b44565b84513d87823e3d90fd5b5090346102b457816003193601126102b4576020906036549051908152f35b5090346102b457816003193601126102b457905490516001600160a01b039091168152602090f35b5090346102b457816003193601126102b457602090603b549051908152f35b5091903461038057602036600319011261038057610c3a611e65565b610c426128f5565b6001600160a01b0316918215610ce057506000805160206131ae83398151915290826bffffffffffffffffffffffff60a01b603c541617603c55805192602084015260208352610c9183611ea4565b601c8151918083528201527f7365745769746864726177616c57616c6c6574286164647265737329000000006060820152608060208201528061062b633abcb7bb60e11b94608083019061241b565b905163d92e233d60e01b8152fd5b5090346102b457816003193601126102b457602090517f5e4bd437d29fad01c10cdcfff414f0d6b0e84b96d2dade88d780d45b5630696b8152f35b5090346102b457806003193601126102b4578235906024359360008051602061318e8339815191528060005260209060008051602061316e83398151915282528360002033600052825260ff84600020541615610f1d57506001600160a01b0390828183610d95612990565b16865192838092637ee56d2f60e11b82525afa908115610f13578791610eda575b50610eca57610dc58786611fcc565b96603654809811610eba57610de58798610ddf8389611fcc565b90612010565b6036558181610e7c575b505084610dfa578580f35b7ffe89805cf5299ef9fbd1d1ddefb8dcc3fa9408064d2ea31e3fca6565768f5217908451868152a1610e2a612b07565b16803b15610e78578491835180958193632689dfd360e11b83525af1908115610e6f5750610e5b575b808080808580f35b610e6490611e7b565b610292578038610e53565b513d84823e3d90fd5b8480fd5b81610eaa7f9d04ecb465d2c8754acb798a22293dd26215a1c2f7a2a56607afa215c1aadc7793603754611fcc565b6037558651908152a13881610def565b84516396b0c75160e01b81528490fd5b83516313d0ff5960e31b81528390fd5b90508181813d8311610f0c575b610ef18183611ebf565b81010312610f0857610f0290611f55565b38610db6565b8680fd5b503d610ee7565b85513d89823e3d90fd5b9050604492519163e2517d3f60e01b835233908301526024820152fd5b5090346102b457816003193601126102b4576020906034549051908152f35b5090346102b457816003193601126102b45760209061028b6125db565b5034610292576020366003190112610292578235906001600160401b03821161029257366023830112156102925750610fd26020610fc0819584602460ff96369301359101611efb565b81855193828580945193849201611f32565b81016033815203019020541690519015158152f35b509190346103805760a036600319011261038057611003611e65565b9061100c611e4f565b84548251636c9c724960e11b81526001600160a01b0395602095909260443592606435929188908290879082908d165afa908115611193579089918b916111ad575b5016330361119d57838789611061612990565b168851928380926345b09c8d60e11b82525afa908115611193578a9161115a575b5061114a5760a4928795949289808c94169616998a878a51858152848b8201527fdefb23229f98ca9823b290c7a4ca8908d69e2a25a6fa9e0b65b3ffad35798eb38c3392a46110cf612b07565b169188519a8b978896633c82491960e01b885287015260248601526044850152606484015260843560848401525af1908115611141575061110e578280f35b81813d831161113a575b6111228183611ebf565b810103126102b45761113390611f55565b5038808280f35b503d611118565b513d85823e3d90fd5b85516313d0ff5960e31b81528490fd5b90508781813d831161118c575b6111718183611ebf565b810103126111885761118290611f55565b38611082565b8980fd5b503d611167565b87513d8c823e3d90fd5b855163c68e02cf60e01b81528490fd5b6111c49150893d8b11610490576104828183611ebf565b3861104e565b5090346102b457816003193601126102b457603c5490516001600160a01b039091168152602090f35b5090346102b457816003193601126102b457602090603e549051908152f35b5090346102b457816003193601126102b45760209060ff603c5460a01c1690519015158152f35b50346102925780600319360112610292578080808047335af1506103cd612583565b5090346102b457816003193601126102b457602090516000805160206131ce8339815191528152f35b50919080600319360112610380578135906024356001600160401b038082116115b157366023830112156115b157818501359081116115b1573660248260061b840101116115b15760018060a01b03918287541684519283636c9c724960e11b92838252818a60209788935afa908115611193579086918b91611594575b5016330361158457878486611315612990565b16885192838092631ea7ca8960e01b82525afa908115611193578a9161154f575b5061153f57603a54803410908115611535575b506115255761135787612440565b92888587611363612a21565b168951928380926318160ddd60e01b82525afa801561151b578b906114ec575b61138e915085611fcc565b603e54106114dc576113a288603654611fcc565b603655856113ae612a21565b1690813b156114d857908a9291885192639386e19760e01b8452808c89602487019187015252602460448501920190855b8181106114a4575050508391838381809403925af1801561149a57908593929161147d575b50859054169786518099819382525afa918215611473577f1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee9095968893611454575b5084519687528601521692a280f35b8161146c9294503d8511610490576104828183611ebf565b9138611445565b84513d89823e3d90fd5b61148991929350611e7b565b6114965782908838611404565b8780fd5b87513d84823e3d90fd5b91809495965092909235908a8216809203610ab05790815283890135898201528d959493908b0192908b01916001016113df565b8a80fd5b8651630935f98160e31b81528990fd5b508581813d8311611514575b6115028183611ebf565b81010312610ab05761138e9051611383565b503d6114f8565b88513d8d823e3d90fd5b8551630a3287a960e21b81528890fd5b9050871038611349565b85516313d0ff5960e31b81528890fd5b90508481813d831161157d575b6115668183611ebf565b810103126111885761157790611f55565b38611336565b503d61155c565b855163c68e02cf60e01b81528890fd5b6115ab9150863d8811610490576104828183611ebf565b38611302565b8580fd5b5090346102b457806003193601126102b4576115cf611e4f565b90336001600160a01b038316036115ec57506103cd919235612b41565b5163334bd91960e11b81528390fd5b5090346102b457816003193601126102b4576020906038549051908152f35b5090346102b457816003193601126102b4576020905160008051602061318e8339815191528152f35b509134610380578160031936011261038057359061165f611e4f565b9082845260008051602061316e8339815191526020526116846001828620015461295d565b61168e8284612ddf565b611696578380f35b600092835260008051602061312e8339815191526020529091206116c3916001600160a01b031690612f30565b503880808380f35b509190346103805760203660031901126103805781359161ffff83168093036117b5576116f66128f5565b6103e883116117a757612710831161179457506000805160206131ae833981519152908261ffff19603954161760395580519260208401526020835261173b83611ea4565b60218151918083528201527f73657445786368616e676541646a7573746d656e74526174652875696e7431366060820152602960f81b608082015260a060208201528061062b630a7521c160e21b9460a083019061241b565b634e487b7160e01b845260019052602483fd5b905163c52a9bd360e01b8152fd5b8380fd5b509134610380576020366003190112610380578160209360019235815260008051602061316e83398151915285522001549051908152f35b509134610380576020366003190112610380576000805160206131ae83398151915290359161181e6128f5565b82603e5580519260208401526020835261183783611ea4565b601d8151918083528201527f7365744d6178696d756d44455448537570706c792875696e74323536290000006060820152608060208201528061062b630e969ae760e11b94608083019061241b565b5090346102b457816003193601126102b4576020905160008051602061314e8339815191528152f35b509182913461191f578260031936011261191f576118cb6128f5565b6001600160a01b036118db612b07565b1691823b1561191a57815163596a15a360e11b81529284918491829084905af1908115610e6f575061190a5750f35b61191390611e7b565b6102925780f35b505050fd5b5050fd5b5034610292576020366003190112610292575061028b60209235612440565b5091346103805760803660031901126103805780356001600160801b0392838216809203610e78576024938435818116809103610f0857611981611e39565b6001600160a01b039160643583611996612990565b1694865180966345b09c8d60e11b8252818b6020998a935afa908115611c15578c91611bdc575b50611bcc576038548810611bbd576119d488612810565b1691808310611ba25750836119e7612b07565b168a813b15610292578a9460a48b8389958c51968795869462448d8d60e91b8652339086015216809a8401528d60448401528860648401528760848401525af18015611b9857611b7f575b50908793929186519182528786830152868201527f7fd59a2c98b4d595be25d7976e588f7b6189ace6d41522f11868d951c293436760603392a3611a74612a21565b83828a541686519485809263352a8adf60e21b82525afa928315611b75578993611b54575b50811694845191848301936323b872dd60e01b8552338a85015216604483015260648201526064815260a081018181106001600160401b03821117611b4057845251611af6918891829182885af1611aef612583565b90856130aa565b8051918215159283611b1f575b505050611b0e578480f35b51635274afe760e01b815291820152fd5b82935091819281010312610f0857611b379101611f55565b15388080611b03565b87604188634e487b7160e01b600052526000fd5b82919350611b6e90853d8711610490576104828183611ebf565b9290611a99565b85513d8b823e3d90fd5b99611b8e89959493929b611e7b565b9990919293611a32565b87513d8d823e3d90fd5b886044918b858a51936347f961e960e11b8552840152820152fd5b865162771ba760e71b81528990fd5b86516313d0ff5960e31b81528990fd5b90508681813d8311611c0e575b611bf38183611ebf565b81010312611c0a57611c0490611f55565b386119bd565b8b80fd5b503d611be9565b88513d8e823e3d90fd5b5090346102b457816003193601126102b457602090603a549051908152f35b5090346102b457816003193601126102b45760209061ffff603954169051908152f35b509134610380576020366003190112610380573590811515809203610380576000805160206131ae83398151915290611c986128f5565b603c805460ff60a01b191660a085901b60ff60a01b161790558051602080820194909452928352611cc883611ea4565b60198151918083528201527f7365745374616b696e67416c6c6f776c69737428626f6f6c29000000000000006060820152608060208201528061062b630279b66160e11b94608083019061241b565b5082346102b457826003193601126102b4578035906001600160401b03908183116117b557366023840112156117b557828101359182116117b5573660248360051b850101116117b55760008051602061314e83398151915280855260008051602061316e83398151915260205285852033865260205260ff868620541615611dae5784611dab858560248035920161201d565b80f35b6044925085519163e2517d3f60e01b835233908301526024820152fd5b90508334610380576020366003190112610380573563ffffffff60e01b81168091036103805760209250635a05180f60e01b8114908115611e0e575b5015158152f35b637965db0b60e01b811491508115611e28575b5083611e07565b6301ffc9a760e01b14905083611e21565b604435906001600160a01b0382168203610ab057565b602435906001600160a01b0382168203610ab057565b600435906001600160a01b0382168203610ab057565b6001600160401b038111611e8e57604052565b634e487b7160e01b600052604160045260246000fd5b604081019081106001600160401b03821117611e8e57604052565b90601f801991011681019081106001600160401b03821117611e8e57604052565b6001600160401b038111611e8e57601f01601f191660200190565b929192611f0782611ee0565b91611f156040519384611ebf565b829481845281830111610ab0578281602093846000960137010152565b60005b838110611f455750506000910152565b8181015183820152602001611f35565b51908115158203610ab057565b9190811015611f845760051b8101359060be1981360301821215610ab0570190565b634e487b7160e01b600052603260045260246000fd5b903590601e1981360301821215610ab057018035906001600160401b038211610ab057602001918136038313610ab057565b91908201809211611fd957565b634e487b7160e01b600052601160045260246000fd5b908060209392818452848401376000828201840152601f01601f1916010190565b91908203918211611fd957565b90916001600160a01b039081612031612990565b16604094855192838093637750955b60e11b8252602094859160049788915afa9081156121cb576000916123e6575b506123d65781156123cd578383866120766129e7565b1689519283809263c5f2892f60e01b82525afa9081156121cb576000916123a0575b5080910361238a57506000805b8281106121e657506037548082116121d657906120c5816120d093612010565b603755603454611fcc565b6034556120df81603554611fcc565b60355560005b8181106120f55750505050505050565b612100818388611f62565b856121096129e7565b1661211689830183611f9a565b90926121256060820182611f9a565b909360809161213683850185611f9a565b9091833b15610ab0578f978d928d946121856121949461217160009d519e8f9d8e9c8d9a6304512a2360e31b8c528b015260848a0191611fef565b9060031995868984030160248a0152611fef565b92858403016044860152611fef565b60a0860135606483015203930135905af180156121cb5790600192916121bc575b50016120e5565b6121c590611e7b565b386121b5565b88513d6000823e3d90fd5b87516307bb2bd760e21b81528590fd5b60ff6121f382858a611f62565b92898401906122028286611f9a565b8c51949181908637840193888160339687815203019020541661237a578685013592603a54840361236a57606061223b81880188611f9a565b908a820361235357600c91808311610ab05781356001600160a01b0319166001600160f81b0319810161233857508b11610ab0570135901c8a603c5416810361232257506123128c6122db8686957f15f16c2e13e50235799a97b981bf4a66c8cd86051f06aca745c5ff26f39b330e958d6122ba60019c9b9a8e611f9a565b928388519485938437820190815203019020805460ff19168b179055611fcc565b976123006122f36122ec8784611f9a565b3691611efb565b8c81519101209582611f9a565b93908351948486958652850191611fef565b958b83015235940390a3016120a5565b896024918e5191631b4d561960e01b8352820152fd5b90508f602493508d925051916308ebf56560e01b8352820152fd5b508a6024918f5191639b0ec52760e01b8352820152fd5b8b516305cacc5560e21b81528990fd5b8a5163932c5b0d60e01b81528890fd5b83602491885191631497ae9360e01b8352820152fd5b90508381813d83116123c6575b6123b78183611ebf565b81010312610ab0575138612098565b503d6123ad565b50505050505050565b86516313d0ff5960e31b81528490fd5b90508381813d8311612414575b6123fd8183611ebf565b81010312610ab05761240e90611f55565b38612060565b503d6123f3565b9060209161243481518092818552858086019101611f32565b601f01601f1916010190565b6001600160a01b0380612451612a21565b169060405190816318160ddd60e01b9384825281600460209586935afa90811561252557600091612537575b501561253157819061248d612a21565b169260046040518095819382525afa908115612525576000916124f9575b50603954612710925061ffff9081168303818111611fd9571690818102918183041490151715611fd9576124dd6125db565b828102928184041490151715611fd9576124f692612a5b565b90565b82813d831161251e575b61250d8183611ebf565b8101031261029257505180386124ab565b503d612503565b6040513d6000823e3d90fd5b50505090565b90508281813d831161255d575b61254e8183611ebf565b81010312610ab057513861247d565b503d612544565b90816020910312610ab057516001600160a01b0381168103610ab05790565b3d156125ae573d9061259482611ee0565b916125a26040519384611ebf565b82523d6000602084013e565b606090565b51906001600160401b0382168203610ab057565b51906001600160801b0382168203610ab057565b600080546040805163159761e360e21b815292939290916001600160a01b039160209182908290600490829087165afa9081156128065786916127e9575b508284519163079d004d60e51b835282600481610100948594165afa918215610f13578792612706575b5050908061268e60049361265c60365460375490611fcc565b60c06126846034549261267e6001600160801b03948560e08801511690612010565b90611fcc565b9201511690611fcc565b93612697612b07565b168551938480926316d3df1560e31b82525afa9384156126fd575085936126c6575b50506124f6929350611fcc565b9080929350813d83116126f6575b6126de8183611ebf565b810103126117b5576124f692935051908392386126b9565b503d6126d4565b513d87823e3d90fd5b9080925081813d83116127e2575b61271e8183611ebf565b81010312610f085784519182018281106001600160401b038211176127ce5783926127c160e06004969461268e948a52612757816125b3565b84526127648682016125b3565b868501526127738a82016125b3565b8a850152612783606082016125b3565b6060850152612794608082016125c7565b60808501526127a560a082016125c7565b60a08501526127b660c082016125c7565b60c0850152016125c7565b60e0820152929350612643565b634e487b7160e01b88526041600452602488fd5b503d612714565b6128009150823d8411610490576104828183611ebf565b38612619565b84513d88823e3d90fd5b6001600160a01b039081612822612a21565b169160405191826318160ddd60e01b9485825281600460209687935afa908115612525576000916128c8575b50156128c1578261285d6125db565b92612866612a21565b169460046040518097819382525afa9283156125255760009361288f575b506124f69350612a5b565b90925083813d83116128ba575b6128a68183611ebf565b81010312610ab0576124f692519138612884565b503d61289c565b9250505090565b90508381813d83116128ee575b6128df8183611ebf565b81010312610ab057513861284e565b503d6128d5565b3360009081527fc5abc2d863e77a76627fe1702320d3afc3f93a9ad30aebcafab5d12854da5c0f60205260409020546000805160206131ce8339815191529060ff161561293f5750565b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b8060005260008051602061316e83398151915260205260406000203360005260205260ff604060002054161561293f5750565b600054604051639fd0506d60e01b81526001600160a01b03916020908290600490829086165afa908115612525576000916129ca57501690565b6129e3915060203d602011610490576104828183611ebf565b1690565b60005460405163e94ad65b60e01b81526001600160a01b03916020908290600490829086165afa908115612525576000916129ca57501690565b6000546040516304efe2eb60e31b81526001600160a01b03916020908290600490829086165afa908115612525576000916129ca57501690565b90918282029160001984820993838086109503948086039514612ae35784831115612ad15782910981600003821680920460028082600302188083028203028083028203028083028203028083028203028083028203028092029003029360018380600003040190848311900302920304170290565b60405163227bc15360e01b8152600490fd5b505080925015612af1570490565b634e487b7160e01b600052601260045260246000fd5b60005460405163352a8adf60e21b81526001600160a01b03916020908290600490829086165afa908115612525576000916129ca57501690565b612b4b8282612e54565b9182612b5657505090565b600091825260008051602061312e8339815191526020526040909120612b85916001600160a01b031690612fba565b5090565b6001600160a01b031660008181527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d602052604081205490919060008051602061316e8339815191529060ff16612c12578280526020526040822081835260205260408220600160ff19825416179055339160008051602061310e8339815191528180a4600190565b505090565b6001600160a01b031660008181527fc5abc2d863e77a76627fe1702320d3afc3f93a9ad30aebcafab5d12854da5c0f60205260408120549091906000805160206131ce8339815191529060008051602061316e8339815191529060ff16612531578184526020526040832082845260205260408320600160ff1982541617905560008051602061310e833981519152339380a4600190565b6001600160a01b031660008181527ff79084f49a5c4fa4c48f70bba1e67b61c2b9ca8b3d302dc944c028fdea010b82602052604081205490919060008051602061318e8339815191529060008051602061316e8339815191529060ff16612531578184526020526040832082845260205260408320600160ff1982541617905560008051602061310e833981519152339380a4600190565b6001600160a01b031660008181527ff1e23661530d14d05c9291333c54312223931d3f1ab2285de8cf548f5a18240d602052604081205490919060008051602061314e8339815191529060008051602061316e8339815191529060ff16612531578184526020526040832082845260205260408320600160ff1982541617905560008051602061310e833981519152339380a4600190565b9060009180835260008051602061316e83398151915280602052604084209260018060a01b03169283855260205260ff60408520541615600014612531578184526020526040832082845260205260408320600160ff1982541617905560008051602061310e833981519152339380a4600190565b9060009180835260008051602061316e83398151915280602052604084209260018060a01b03169283855260205260ff60408520541660001461253157818452602052604083208284526020526040832060ff1981541690557ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b339380a4600190565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c1615612f0657565b604051631afcd79f60e31b8152600490fd5b8054821015611f845760005260206000200190600090565b91906001830160009082825280602052604082205415600014612fb45784549468010000000000000000861015612fa05783612f90612f79886001604098999a01855584612f18565b819391549060031b91821b91600019901b19161790565b9055549382526020522055600190565b634e487b7160e01b83526041600452602483fd5b50925050565b906001820190600092818452826020526040842054908115156000146130a3576000199180830181811161308f5782549084820191821161307b57818103613046575b50505080548015613032578201916130158383612f18565b909182549160031b1b191690555582526020526040812055600190565b634e487b7160e01b86526031600452602486fd5b613066613056612f799386612f18565b90549060031b1c92839286612f18565b90558652846020526040862055388080612ffd565b634e487b7160e01b88526011600452602488fd5b634e487b7160e01b87526011600452602487fd5b5050505090565b906130d157508051156130bf57805190602001fd5b604051630a12f52160e11b8152600490fd5b81511580613104575b6130e2575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b156130da56fe2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0dc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000e30bb2df90b65284acd0e8b5ebe3483bb2bbe65a08e43f0f9e8300fd8607ee1102dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800e6ef7125bfa79685f3bd2e4c4cea243c1e988ebbc0801ab7641ae36b9e2c529101d854e8dde9402801a4c6f2840193465752abfad61e0bb7c4258d526ae42e74a6b5d83d32632203555cb9b2c2f68a8d94da48cadd9266ac0d17babedb52ea5ba2646970667358221220637c14c7f7b9ae4d89da88a6f49e17e19a34c5cecccbb8d16f950a85c7c66c5564736f6c63430008180033",
}

// StakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingManagerMetaData.ABI instead.
var StakingManagerABI = StakingManagerMetaData.ABI

// StakingManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingManagerMetaData.Bin instead.
var StakingManagerBin = StakingManagerMetaData.Bin

// DeployStakingManager deploys a new Ethereum contract, binding an instance of StakingManager to it.
func DeployStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StakingManager, error) {
	parsed, err := StakingManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// StakingManager is an auto generated Go binding around an Ethereum contract.
type StakingManager struct {
	StakingManagerCaller     // Read-only binding to the contract
	StakingManagerTransactor // Write-only binding to the contract
	StakingManagerFilterer   // Log filterer for contract events
}

// StakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingManagerSession struct {
	Contract     *StakingManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingManagerCallerSession struct {
	Contract *StakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingManagerTransactorSession struct {
	Contract     *StakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingManagerRaw struct {
	Contract *StakingManager // Generic contract binding to access the raw methods on
}

// StakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingManagerCallerRaw struct {
	Contract *StakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingManagerTransactorRaw struct {
	Contract *StakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingManager creates a new instance of StakingManager, bound to a specific deployed contract.
func NewStakingManager(address common.Address, backend bind.ContractBackend) (*StakingManager, error) {
	contract, err := bindStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// NewStakingManagerCaller creates a new read-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*StakingManagerCaller, error) {
	contract, err := bindStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerCaller{contract: contract}, nil
}

// NewStakingManagerTransactor creates a new write-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingManagerTransactor, error) {
	contract, err := bindStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerTransactor{contract: contract}, nil
}

// NewStakingManagerFilterer creates a new log filterer instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingManagerFilterer, error) {
	contract, err := bindStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingManagerFilterer{contract: contract}, nil
}

// bindStakingManager binds a generic wrapper to an already deployed contract.
func bindStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.StakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transact(opts, method, params...)
}

// ALLOCATORSERVICEROLE is a free data retrieval call binding the contract method 0x3101d910.
//
// Solidity: function ALLOCATOR_SERVICE_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) ALLOCATORSERVICEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "ALLOCATOR_SERVICE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ALLOCATORSERVICEROLE is a free data retrieval call binding the contract method 0x3101d910.
//
// Solidity: function ALLOCATOR_SERVICE_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) ALLOCATORSERVICEROLE() ([32]byte, error) {
	return _StakingManager.Contract.ALLOCATORSERVICEROLE(&_StakingManager.CallOpts)
}

// ALLOCATORSERVICEROLE is a free data retrieval call binding the contract method 0x3101d910.
//
// Solidity: function ALLOCATOR_SERVICE_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) ALLOCATORSERVICEROLE() ([32]byte, error) {
	return _StakingManager.Contract.ALLOCATORSERVICEROLE(&_StakingManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingManager.Contract.DEFAULTADMINROLE(&_StakingManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingManager.Contract.DEFAULTADMINROLE(&_StakingManager.CallOpts)
}

// INITIATORSERVICEROLE is a free data retrieval call binding the contract method 0x19efd5c7.
//
// Solidity: function INITIATOR_SERVICE_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) INITIATORSERVICEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "INITIATOR_SERVICE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITIATORSERVICEROLE is a free data retrieval call binding the contract method 0x19efd5c7.
//
// Solidity: function INITIATOR_SERVICE_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) INITIATORSERVICEROLE() ([32]byte, error) {
	return _StakingManager.Contract.INITIATORSERVICEROLE(&_StakingManager.CallOpts)
}

// INITIATORSERVICEROLE is a free data retrieval call binding the contract method 0x19efd5c7.
//
// Solidity: function INITIATOR_SERVICE_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) INITIATORSERVICEROLE() ([32]byte, error) {
	return _StakingManager.Contract.INITIATORSERVICEROLE(&_StakingManager.CallOpts)
}

// STAKINGALLOWLISTMANAGERROLE is a free data retrieval call binding the contract method 0xe55d6cc0.
//
// Solidity: function STAKING_ALLOWLIST_MANAGER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) STAKINGALLOWLISTMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "STAKING_ALLOWLIST_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGALLOWLISTMANAGERROLE is a free data retrieval call binding the contract method 0xe55d6cc0.
//
// Solidity: function STAKING_ALLOWLIST_MANAGER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) STAKINGALLOWLISTMANAGERROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGALLOWLISTMANAGERROLE(&_StakingManager.CallOpts)
}

// STAKINGALLOWLISTMANAGERROLE is a free data retrieval call binding the contract method 0xe55d6cc0.
//
// Solidity: function STAKING_ALLOWLIST_MANAGER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) STAKINGALLOWLISTMANAGERROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGALLOWLISTMANAGERROLE(&_StakingManager.CallOpts)
}

// STAKINGALLOWLISTROLE is a free data retrieval call binding the contract method 0x89e80ed3.
//
// Solidity: function STAKING_ALLOWLIST_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) STAKINGALLOWLISTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "STAKING_ALLOWLIST_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGALLOWLISTROLE is a free data retrieval call binding the contract method 0x89e80ed3.
//
// Solidity: function STAKING_ALLOWLIST_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) STAKINGALLOWLISTROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGALLOWLISTROLE(&_StakingManager.CallOpts)
}

// STAKINGALLOWLISTROLE is a free data retrieval call binding the contract method 0x89e80ed3.
//
// Solidity: function STAKING_ALLOWLIST_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) STAKINGALLOWLISTROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGALLOWLISTROLE(&_StakingManager.CallOpts)
}

// STAKINGMANAGERROLE is a free data retrieval call binding the contract method 0x3937c0b3.
//
// Solidity: function STAKING_MANAGER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) STAKINGMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "STAKING_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMANAGERROLE is a free data retrieval call binding the contract method 0x3937c0b3.
//
// Solidity: function STAKING_MANAGER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) STAKINGMANAGERROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGMANAGERROLE(&_StakingManager.CallOpts)
}

// STAKINGMANAGERROLE is a free data retrieval call binding the contract method 0x3937c0b3.
//
// Solidity: function STAKING_MANAGER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) STAKINGMANAGERROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGMANAGERROLE(&_StakingManager.CallOpts)
}

// TOPUPROLE is a free data retrieval call binding the contract method 0x6fce8ab2.
//
// Solidity: function TOP_UP_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) TOPUPROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "TOP_UP_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOPUPROLE is a free data retrieval call binding the contract method 0x6fce8ab2.
//
// Solidity: function TOP_UP_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) TOPUPROLE() ([32]byte, error) {
	return _StakingManager.Contract.TOPUPROLE(&_StakingManager.CallOpts)
}

// TOPUPROLE is a free data retrieval call binding the contract method 0x6fce8ab2.
//
// Solidity: function TOP_UP_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) TOPUPROLE() ([32]byte, error) {
	return _StakingManager.Contract.TOPUPROLE(&_StakingManager.CallOpts)
}

// AllocatedETHForDeposits is a free data retrieval call binding the contract method 0xea452b6d.
//
// Solidity: function allocatedETHForDeposits() view returns(uint256)
func (_StakingManager *StakingManagerCaller) AllocatedETHForDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "allocatedETHForDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllocatedETHForDeposits is a free data retrieval call binding the contract method 0xea452b6d.
//
// Solidity: function allocatedETHForDeposits() view returns(uint256)
func (_StakingManager *StakingManagerSession) AllocatedETHForDeposits() (*big.Int, error) {
	return _StakingManager.Contract.AllocatedETHForDeposits(&_StakingManager.CallOpts)
}

// AllocatedETHForDeposits is a free data retrieval call binding the contract method 0xea452b6d.
//
// Solidity: function allocatedETHForDeposits() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) AllocatedETHForDeposits() (*big.Int, error) {
	return _StakingManager.Contract.AllocatedETHForDeposits(&_StakingManager.CallOpts)
}

// ExchangeAdjustmentRate is a free data retrieval call binding the contract method 0x0633af76.
//
// Solidity: function exchangeAdjustmentRate() view returns(uint16)
func (_StakingManager *StakingManagerCaller) ExchangeAdjustmentRate(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "exchangeAdjustmentRate")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ExchangeAdjustmentRate is a free data retrieval call binding the contract method 0x0633af76.
//
// Solidity: function exchangeAdjustmentRate() view returns(uint16)
func (_StakingManager *StakingManagerSession) ExchangeAdjustmentRate() (uint16, error) {
	return _StakingManager.Contract.ExchangeAdjustmentRate(&_StakingManager.CallOpts)
}

// ExchangeAdjustmentRate is a free data retrieval call binding the contract method 0x0633af76.
//
// Solidity: function exchangeAdjustmentRate() view returns(uint16)
func (_StakingManager *StakingManagerCallerSession) ExchangeAdjustmentRate() (uint16, error) {
	return _StakingManager.Contract.ExchangeAdjustmentRate(&_StakingManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingManager *StakingManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingManager *StakingManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingManager.Contract.GetRoleAdmin(&_StakingManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingManager.Contract.GetRoleAdmin(&_StakingManager.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StakingManager *StakingManagerCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StakingManager *StakingManagerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _StakingManager.Contract.GetRoleMember(&_StakingManager.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StakingManager *StakingManagerCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _StakingManager.Contract.GetRoleMember(&_StakingManager.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _StakingManager.Contract.GetRoleMemberCount(&_StakingManager.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _StakingManager.Contract.GetRoleMemberCount(&_StakingManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingManager *StakingManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingManager *StakingManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingManager.Contract.HasRole(&_StakingManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingManager.Contract.HasRole(&_StakingManager.CallOpts, role, account)
}

// InitializationBlockNumber is a free data retrieval call binding the contract method 0xb91590b2.
//
// Solidity: function initializationBlockNumber() view returns(uint256)
func (_StakingManager *StakingManagerCaller) InitializationBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "initializationBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitializationBlockNumber is a free data retrieval call binding the contract method 0xb91590b2.
//
// Solidity: function initializationBlockNumber() view returns(uint256)
func (_StakingManager *StakingManagerSession) InitializationBlockNumber() (*big.Int, error) {
	return _StakingManager.Contract.InitializationBlockNumber(&_StakingManager.CallOpts)
}

// InitializationBlockNumber is a free data retrieval call binding the contract method 0xb91590b2.
//
// Solidity: function initializationBlockNumber() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) InitializationBlockNumber() (*big.Int, error) {
	return _StakingManager.Contract.InitializationBlockNumber(&_StakingManager.CallOpts)
}

// IsStakingAllowlist is a free data retrieval call binding the contract method 0x42d3915d.
//
// Solidity: function isStakingAllowlist() view returns(bool)
func (_StakingManager *StakingManagerCaller) IsStakingAllowlist(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "isStakingAllowlist")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStakingAllowlist is a free data retrieval call binding the contract method 0x42d3915d.
//
// Solidity: function isStakingAllowlist() view returns(bool)
func (_StakingManager *StakingManagerSession) IsStakingAllowlist() (bool, error) {
	return _StakingManager.Contract.IsStakingAllowlist(&_StakingManager.CallOpts)
}

// IsStakingAllowlist is a free data retrieval call binding the contract method 0x42d3915d.
//
// Solidity: function isStakingAllowlist() view returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsStakingAllowlist() (bool, error) {
	return _StakingManager.Contract.IsStakingAllowlist(&_StakingManager.CallOpts)
}

// Locator is a free data retrieval call binding the contract method 0x7c957fc8.
//
// Solidity: function locator() view returns(address)
func (_StakingManager *StakingManagerCaller) Locator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "locator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Locator is a free data retrieval call binding the contract method 0x7c957fc8.
//
// Solidity: function locator() view returns(address)
func (_StakingManager *StakingManagerSession) Locator() (common.Address, error) {
	return _StakingManager.Contract.Locator(&_StakingManager.CallOpts)
}

// Locator is a free data retrieval call binding the contract method 0x7c957fc8.
//
// Solidity: function locator() view returns(address)
func (_StakingManager *StakingManagerCallerSession) Locator() (common.Address, error) {
	return _StakingManager.Contract.Locator(&_StakingManager.CallOpts)
}

// MaximumDETHSupply is a free data retrieval call binding the contract method 0x49336f0f.
//
// Solidity: function maximumDETHSupply() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MaximumDETHSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "maximumDETHSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumDETHSupply is a free data retrieval call binding the contract method 0x49336f0f.
//
// Solidity: function maximumDETHSupply() view returns(uint256)
func (_StakingManager *StakingManagerSession) MaximumDETHSupply() (*big.Int, error) {
	return _StakingManager.Contract.MaximumDETHSupply(&_StakingManager.CallOpts)
}

// MaximumDETHSupply is a free data retrieval call binding the contract method 0x49336f0f.
//
// Solidity: function maximumDETHSupply() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MaximumDETHSupply() (*big.Int, error) {
	return _StakingManager.Contract.MaximumDETHSupply(&_StakingManager.CallOpts)
}

// MaximumDepositAmount is a free data retrieval call binding the contract method 0x78abb49b.
//
// Solidity: function maximumDepositAmount() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MaximumDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "maximumDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumDepositAmount is a free data retrieval call binding the contract method 0x78abb49b.
//
// Solidity: function maximumDepositAmount() view returns(uint256)
func (_StakingManager *StakingManagerSession) MaximumDepositAmount() (*big.Int, error) {
	return _StakingManager.Contract.MaximumDepositAmount(&_StakingManager.CallOpts)
}

// MaximumDepositAmount is a free data retrieval call binding the contract method 0x78abb49b.
//
// Solidity: function maximumDepositAmount() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MaximumDepositAmount() (*big.Int, error) {
	return _StakingManager.Contract.MaximumDepositAmount(&_StakingManager.CallOpts)
}

// MinimumDepositAmount is a free data retrieval call binding the contract method 0x080c279a.
//
// Solidity: function minimumDepositAmount() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MinimumDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "minimumDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumDepositAmount is a free data retrieval call binding the contract method 0x080c279a.
//
// Solidity: function minimumDepositAmount() view returns(uint256)
func (_StakingManager *StakingManagerSession) MinimumDepositAmount() (*big.Int, error) {
	return _StakingManager.Contract.MinimumDepositAmount(&_StakingManager.CallOpts)
}

// MinimumDepositAmount is a free data retrieval call binding the contract method 0x080c279a.
//
// Solidity: function minimumDepositAmount() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MinimumDepositAmount() (*big.Int, error) {
	return _StakingManager.Contract.MinimumDepositAmount(&_StakingManager.CallOpts)
}

// MinimumUnstakeBound is a free data retrieval call binding the contract method 0x35ead2a4.
//
// Solidity: function minimumUnstakeBound() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MinimumUnstakeBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "minimumUnstakeBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumUnstakeBound is a free data retrieval call binding the contract method 0x35ead2a4.
//
// Solidity: function minimumUnstakeBound() view returns(uint256)
func (_StakingManager *StakingManagerSession) MinimumUnstakeBound() (*big.Int, error) {
	return _StakingManager.Contract.MinimumUnstakeBound(&_StakingManager.CallOpts)
}

// MinimumUnstakeBound is a free data retrieval call binding the contract method 0x35ead2a4.
//
// Solidity: function minimumUnstakeBound() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MinimumUnstakeBound() (*big.Int, error) {
	return _StakingManager.Contract.MinimumUnstakeBound(&_StakingManager.CallOpts)
}

// NumInitiatedValidators is a free data retrieval call binding the contract method 0xbb635c65.
//
// Solidity: function numInitiatedValidators() view returns(uint256)
func (_StakingManager *StakingManagerCaller) NumInitiatedValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "numInitiatedValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumInitiatedValidators is a free data retrieval call binding the contract method 0xbb635c65.
//
// Solidity: function numInitiatedValidators() view returns(uint256)
func (_StakingManager *StakingManagerSession) NumInitiatedValidators() (*big.Int, error) {
	return _StakingManager.Contract.NumInitiatedValidators(&_StakingManager.CallOpts)
}

// NumInitiatedValidators is a free data retrieval call binding the contract method 0xbb635c65.
//
// Solidity: function numInitiatedValidators() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) NumInitiatedValidators() (*big.Int, error) {
	return _StakingManager.Contract.NumInitiatedValidators(&_StakingManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingManager *StakingManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingManager *StakingManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingManager.Contract.SupportsInterface(&_StakingManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingManager.Contract.SupportsInterface(&_StakingManager.CallOpts, interfaceId)
}

// TotalDepositedInValidators is a free data retrieval call binding the contract method 0x60a0f628.
//
// Solidity: function totalDepositedInValidators() view returns(uint256)
func (_StakingManager *StakingManagerCaller) TotalDepositedInValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "totalDepositedInValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDepositedInValidators is a free data retrieval call binding the contract method 0x60a0f628.
//
// Solidity: function totalDepositedInValidators() view returns(uint256)
func (_StakingManager *StakingManagerSession) TotalDepositedInValidators() (*big.Int, error) {
	return _StakingManager.Contract.TotalDepositedInValidators(&_StakingManager.CallOpts)
}

// TotalDepositedInValidators is a free data retrieval call binding the contract method 0x60a0f628.
//
// Solidity: function totalDepositedInValidators() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) TotalDepositedInValidators() (*big.Int, error) {
	return _StakingManager.Contract.TotalDepositedInValidators(&_StakingManager.CallOpts)
}

// UnallocatedETH is a free data retrieval call binding the contract method 0x7dfcdd29.
//
// Solidity: function unallocatedETH() view returns(uint256)
func (_StakingManager *StakingManagerCaller) UnallocatedETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "unallocatedETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnallocatedETH is a free data retrieval call binding the contract method 0x7dfcdd29.
//
// Solidity: function unallocatedETH() view returns(uint256)
func (_StakingManager *StakingManagerSession) UnallocatedETH() (*big.Int, error) {
	return _StakingManager.Contract.UnallocatedETH(&_StakingManager.CallOpts)
}

// UnallocatedETH is a free data retrieval call binding the contract method 0x7dfcdd29.
//
// Solidity: function unallocatedETH() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) UnallocatedETH() (*big.Int, error) {
	return _StakingManager.Contract.UnallocatedETH(&_StakingManager.CallOpts)
}

// UnstakeRequestInfo is a free data retrieval call binding the contract method 0xac1e2257.
//
// Solidity: function unstakeRequestInfo(uint256 destChainId, address l2strategy) view returns(bool, uint256)
func (_StakingManager *StakingManagerCaller) UnstakeRequestInfo(opts *bind.CallOpts, destChainId *big.Int, l2strategy common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "unstakeRequestInfo", destChainId, l2strategy)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// UnstakeRequestInfo is a free data retrieval call binding the contract method 0xac1e2257.
//
// Solidity: function unstakeRequestInfo(uint256 destChainId, address l2strategy) view returns(bool, uint256)
func (_StakingManager *StakingManagerSession) UnstakeRequestInfo(destChainId *big.Int, l2strategy common.Address) (bool, *big.Int, error) {
	return _StakingManager.Contract.UnstakeRequestInfo(&_StakingManager.CallOpts, destChainId, l2strategy)
}

// UnstakeRequestInfo is a free data retrieval call binding the contract method 0xac1e2257.
//
// Solidity: function unstakeRequestInfo(uint256 destChainId, address l2strategy) view returns(bool, uint256)
func (_StakingManager *StakingManagerCallerSession) UnstakeRequestInfo(destChainId *big.Int, l2strategy common.Address) (bool, *big.Int, error) {
	return _StakingManager.Contract.UnstakeRequestInfo(&_StakingManager.CallOpts, destChainId, l2strategy)
}

// UsedValidators is a free data retrieval call binding the contract method 0x5915ded1.
//
// Solidity: function usedValidators(bytes pubkey) view returns(bool exists)
func (_StakingManager *StakingManagerCaller) UsedValidators(opts *bind.CallOpts, pubkey []byte) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "usedValidators", pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedValidators is a free data retrieval call binding the contract method 0x5915ded1.
//
// Solidity: function usedValidators(bytes pubkey) view returns(bool exists)
func (_StakingManager *StakingManagerSession) UsedValidators(pubkey []byte) (bool, error) {
	return _StakingManager.Contract.UsedValidators(&_StakingManager.CallOpts, pubkey)
}

// UsedValidators is a free data retrieval call binding the contract method 0x5915ded1.
//
// Solidity: function usedValidators(bytes pubkey) view returns(bool exists)
func (_StakingManager *StakingManagerCallerSession) UsedValidators(pubkey []byte) (bool, error) {
	return _StakingManager.Contract.UsedValidators(&_StakingManager.CallOpts, pubkey)
}

// WithdrawalWallet is a free data retrieval call binding the contract method 0x4a7d80b3.
//
// Solidity: function withdrawalWallet() view returns(address)
func (_StakingManager *StakingManagerCaller) WithdrawalWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "withdrawalWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawalWallet is a free data retrieval call binding the contract method 0x4a7d80b3.
//
// Solidity: function withdrawalWallet() view returns(address)
func (_StakingManager *StakingManagerSession) WithdrawalWallet() (common.Address, error) {
	return _StakingManager.Contract.WithdrawalWallet(&_StakingManager.CallOpts)
}

// WithdrawalWallet is a free data retrieval call binding the contract method 0x4a7d80b3.
//
// Solidity: function withdrawalWallet() view returns(address)
func (_StakingManager *StakingManagerCallerSession) WithdrawalWallet() (common.Address, error) {
	return _StakingManager.Contract.WithdrawalWallet(&_StakingManager.CallOpts)
}

// AllocateETH is a paid mutator transaction binding the contract method 0x6daa01a2.
//
// Solidity: function allocateETH(uint256 allocateToUnstakeRequestsManager, uint256 allocateToDeposits) returns()
func (_StakingManager *StakingManagerTransactor) AllocateETH(opts *bind.TransactOpts, allocateToUnstakeRequestsManager *big.Int, allocateToDeposits *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "allocateETH", allocateToUnstakeRequestsManager, allocateToDeposits)
}

// AllocateETH is a paid mutator transaction binding the contract method 0x6daa01a2.
//
// Solidity: function allocateETH(uint256 allocateToUnstakeRequestsManager, uint256 allocateToDeposits) returns()
func (_StakingManager *StakingManagerSession) AllocateETH(allocateToUnstakeRequestsManager *big.Int, allocateToDeposits *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.AllocateETH(&_StakingManager.TransactOpts, allocateToUnstakeRequestsManager, allocateToDeposits)
}

// AllocateETH is a paid mutator transaction binding the contract method 0x6daa01a2.
//
// Solidity: function allocateETH(uint256 allocateToUnstakeRequestsManager, uint256 allocateToDeposits) returns()
func (_StakingManager *StakingManagerTransactorSession) AllocateETH(allocateToUnstakeRequestsManager *big.Int, allocateToDeposits *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.AllocateETH(&_StakingManager.TransactOpts, allocateToUnstakeRequestsManager, allocateToDeposits)
}

// ClaimUnstakeRequest is a paid mutator transaction binding the contract method 0x4bd6fc08.
//
// Solidity: function claimUnstakeRequest(address l2Strategy, address bridge, uint256 sourceChainId, uint256 destChainId, uint256 gasLimit) returns()
func (_StakingManager *StakingManagerTransactor) ClaimUnstakeRequest(opts *bind.TransactOpts, l2Strategy common.Address, bridge common.Address, sourceChainId *big.Int, destChainId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "claimUnstakeRequest", l2Strategy, bridge, sourceChainId, destChainId, gasLimit)
}

// ClaimUnstakeRequest is a paid mutator transaction binding the contract method 0x4bd6fc08.
//
// Solidity: function claimUnstakeRequest(address l2Strategy, address bridge, uint256 sourceChainId, uint256 destChainId, uint256 gasLimit) returns()
func (_StakingManager *StakingManagerSession) ClaimUnstakeRequest(l2Strategy common.Address, bridge common.Address, sourceChainId *big.Int, destChainId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.ClaimUnstakeRequest(&_StakingManager.TransactOpts, l2Strategy, bridge, sourceChainId, destChainId, gasLimit)
}

// ClaimUnstakeRequest is a paid mutator transaction binding the contract method 0x4bd6fc08.
//
// Solidity: function claimUnstakeRequest(address l2Strategy, address bridge, uint256 sourceChainId, uint256 destChainId, uint256 gasLimit) returns()
func (_StakingManager *StakingManagerTransactorSession) ClaimUnstakeRequest(l2Strategy common.Address, bridge common.Address, sourceChainId *big.Int, destChainId *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.ClaimUnstakeRequest(&_StakingManager.TransactOpts, l2Strategy, bridge, sourceChainId, destChainId, gasLimit)
}

// DETHToETH is a paid mutator transaction binding the contract method 0xed9daafb.
//
// Solidity: function dETHToETH(uint256 dETHAmount) returns(uint256)
func (_StakingManager *StakingManagerTransactor) DETHToETH(opts *bind.TransactOpts, dETHAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "dETHToETH", dETHAmount)
}

// DETHToETH is a paid mutator transaction binding the contract method 0xed9daafb.
//
// Solidity: function dETHToETH(uint256 dETHAmount) returns(uint256)
func (_StakingManager *StakingManagerSession) DETHToETH(dETHAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.DETHToETH(&_StakingManager.TransactOpts, dETHAmount)
}

// DETHToETH is a paid mutator transaction binding the contract method 0xed9daafb.
//
// Solidity: function dETHToETH(uint256 dETHAmount) returns(uint256)
func (_StakingManager *StakingManagerTransactorSession) DETHToETH(dETHAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.DETHToETH(&_StakingManager.TransactOpts, dETHAmount)
}

// EthToDETH is a paid mutator transaction binding the contract method 0x147d36d5.
//
// Solidity: function ethToDETH(uint256 ethAmount) returns(uint256)
func (_StakingManager *StakingManagerTransactor) EthToDETH(opts *bind.TransactOpts, ethAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "ethToDETH", ethAmount)
}

// EthToDETH is a paid mutator transaction binding the contract method 0x147d36d5.
//
// Solidity: function ethToDETH(uint256 ethAmount) returns(uint256)
func (_StakingManager *StakingManagerSession) EthToDETH(ethAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.EthToDETH(&_StakingManager.TransactOpts, ethAmount)
}

// EthToDETH is a paid mutator transaction binding the contract method 0x147d36d5.
//
// Solidity: function ethToDETH(uint256 ethAmount) returns(uint256)
func (_StakingManager *StakingManagerTransactorSession) EthToDETH(ethAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.EthToDETH(&_StakingManager.TransactOpts, ethAmount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.GrantRole(&_StakingManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.GrantRole(&_StakingManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f656d22.
//
// Solidity: function initialize((address,address,address,address,address) init) returns()
func (_StakingManager *StakingManagerTransactor) Initialize(opts *bind.TransactOpts, init StakingManagerInit) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "initialize", init)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f656d22.
//
// Solidity: function initialize((address,address,address,address,address) init) returns()
func (_StakingManager *StakingManagerSession) Initialize(init StakingManagerInit) (*types.Transaction, error) {
	return _StakingManager.Contract.Initialize(&_StakingManager.TransactOpts, init)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f656d22.
//
// Solidity: function initialize((address,address,address,address,address) init) returns()
func (_StakingManager *StakingManagerTransactorSession) Initialize(init StakingManagerInit) (*types.Transaction, error) {
	return _StakingManager.Contract.Initialize(&_StakingManager.TransactOpts, init)
}

// InitiateValidatorsWithDeposits is a paid mutator transaction binding the contract method 0x0208e4b5.
//
// Solidity: function initiateValidatorsWithDeposits((uint256,uint256,bytes,bytes,bytes,bytes32)[] validators, bytes32 expectedDepositRoot) returns()
func (_StakingManager *StakingManagerTransactor) InitiateValidatorsWithDeposits(opts *bind.TransactOpts, validators []StakingManagerStorageValidatorParams, expectedDepositRoot [32]byte) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "initiateValidatorsWithDeposits", validators, expectedDepositRoot)
}

// InitiateValidatorsWithDeposits is a paid mutator transaction binding the contract method 0x0208e4b5.
//
// Solidity: function initiateValidatorsWithDeposits((uint256,uint256,bytes,bytes,bytes,bytes32)[] validators, bytes32 expectedDepositRoot) returns()
func (_StakingManager *StakingManagerSession) InitiateValidatorsWithDeposits(validators []StakingManagerStorageValidatorParams, expectedDepositRoot [32]byte) (*types.Transaction, error) {
	return _StakingManager.Contract.InitiateValidatorsWithDeposits(&_StakingManager.TransactOpts, validators, expectedDepositRoot)
}

// InitiateValidatorsWithDeposits is a paid mutator transaction binding the contract method 0x0208e4b5.
//
// Solidity: function initiateValidatorsWithDeposits((uint256,uint256,bytes,bytes,bytes,bytes32)[] validators, bytes32 expectedDepositRoot) returns()
func (_StakingManager *StakingManagerTransactorSession) InitiateValidatorsWithDeposits(validators []StakingManagerStorageValidatorParams, expectedDepositRoot [32]byte) (*types.Transaction, error) {
	return _StakingManager.Contract.InitiateValidatorsWithDeposits(&_StakingManager.TransactOpts, validators, expectedDepositRoot)
}

// ReceiveFromUnstakeRequestsManager is a paid mutator transaction binding the contract method 0xc151aa72.
//
// Solidity: function receiveFromUnstakeRequestsManager() payable returns()
func (_StakingManager *StakingManagerTransactor) ReceiveFromUnstakeRequestsManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "receiveFromUnstakeRequestsManager")
}

// ReceiveFromUnstakeRequestsManager is a paid mutator transaction binding the contract method 0xc151aa72.
//
// Solidity: function receiveFromUnstakeRequestsManager() payable returns()
func (_StakingManager *StakingManagerSession) ReceiveFromUnstakeRequestsManager() (*types.Transaction, error) {
	return _StakingManager.Contract.ReceiveFromUnstakeRequestsManager(&_StakingManager.TransactOpts)
}

// ReceiveFromUnstakeRequestsManager is a paid mutator transaction binding the contract method 0xc151aa72.
//
// Solidity: function receiveFromUnstakeRequestsManager() payable returns()
func (_StakingManager *StakingManagerTransactorSession) ReceiveFromUnstakeRequestsManager() (*types.Transaction, error) {
	return _StakingManager.Contract.ReceiveFromUnstakeRequestsManager(&_StakingManager.TransactOpts)
}

// ReceiveReturns is a paid mutator transaction binding the contract method 0x808d663f.
//
// Solidity: function receiveReturns() payable returns()
func (_StakingManager *StakingManagerTransactor) ReceiveReturns(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "receiveReturns")
}

// ReceiveReturns is a paid mutator transaction binding the contract method 0x808d663f.
//
// Solidity: function receiveReturns() payable returns()
func (_StakingManager *StakingManagerSession) ReceiveReturns() (*types.Transaction, error) {
	return _StakingManager.Contract.ReceiveReturns(&_StakingManager.TransactOpts)
}

// ReceiveReturns is a paid mutator transaction binding the contract method 0x808d663f.
//
// Solidity: function receiveReturns() payable returns()
func (_StakingManager *StakingManagerTransactorSession) ReceiveReturns() (*types.Transaction, error) {
	return _StakingManager.Contract.ReceiveReturns(&_StakingManager.TransactOpts)
}

// ReclaimAllocatedETHSurplus is a paid mutator transaction binding the contract method 0x1943190d.
//
// Solidity: function reclaimAllocatedETHSurplus() returns()
func (_StakingManager *StakingManagerTransactor) ReclaimAllocatedETHSurplus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "reclaimAllocatedETHSurplus")
}

// ReclaimAllocatedETHSurplus is a paid mutator transaction binding the contract method 0x1943190d.
//
// Solidity: function reclaimAllocatedETHSurplus() returns()
func (_StakingManager *StakingManagerSession) ReclaimAllocatedETHSurplus() (*types.Transaction, error) {
	return _StakingManager.Contract.ReclaimAllocatedETHSurplus(&_StakingManager.TransactOpts)
}

// ReclaimAllocatedETHSurplus is a paid mutator transaction binding the contract method 0x1943190d.
//
// Solidity: function reclaimAllocatedETHSurplus() returns()
func (_StakingManager *StakingManagerTransactorSession) ReclaimAllocatedETHSurplus() (*types.Transaction, error) {
	return _StakingManager.Contract.ReclaimAllocatedETHSurplus(&_StakingManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingManager *StakingManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingManager *StakingManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceRole(&_StakingManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingManager *StakingManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceRole(&_StakingManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RevokeRole(&_StakingManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RevokeRole(&_StakingManager.TransactOpts, role, account)
}

// SetExchangeAdjustmentRate is a paid mutator transaction binding the contract method 0x29d48704.
//
// Solidity: function setExchangeAdjustmentRate(uint16 exchangeAdjustmentRate_) returns()
func (_StakingManager *StakingManagerTransactor) SetExchangeAdjustmentRate(opts *bind.TransactOpts, exchangeAdjustmentRate_ uint16) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setExchangeAdjustmentRate", exchangeAdjustmentRate_)
}

// SetExchangeAdjustmentRate is a paid mutator transaction binding the contract method 0x29d48704.
//
// Solidity: function setExchangeAdjustmentRate(uint16 exchangeAdjustmentRate_) returns()
func (_StakingManager *StakingManagerSession) SetExchangeAdjustmentRate(exchangeAdjustmentRate_ uint16) (*types.Transaction, error) {
	return _StakingManager.Contract.SetExchangeAdjustmentRate(&_StakingManager.TransactOpts, exchangeAdjustmentRate_)
}

// SetExchangeAdjustmentRate is a paid mutator transaction binding the contract method 0x29d48704.
//
// Solidity: function setExchangeAdjustmentRate(uint16 exchangeAdjustmentRate_) returns()
func (_StakingManager *StakingManagerTransactorSession) SetExchangeAdjustmentRate(exchangeAdjustmentRate_ uint16) (*types.Transaction, error) {
	return _StakingManager.Contract.SetExchangeAdjustmentRate(&_StakingManager.TransactOpts, exchangeAdjustmentRate_)
}

// SetLocator is a paid mutator transaction binding the contract method 0xa5e84562.
//
// Solidity: function setLocator(address _locator) returns()
func (_StakingManager *StakingManagerTransactor) SetLocator(opts *bind.TransactOpts, _locator common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setLocator", _locator)
}

// SetLocator is a paid mutator transaction binding the contract method 0xa5e84562.
//
// Solidity: function setLocator(address _locator) returns()
func (_StakingManager *StakingManagerSession) SetLocator(_locator common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetLocator(&_StakingManager.TransactOpts, _locator)
}

// SetLocator is a paid mutator transaction binding the contract method 0xa5e84562.
//
// Solidity: function setLocator(address _locator) returns()
func (_StakingManager *StakingManagerTransactorSession) SetLocator(_locator common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetLocator(&_StakingManager.TransactOpts, _locator)
}

// SetMaximumDETHSupply is a paid mutator transaction binding the contract method 0x1d2d35ce.
//
// Solidity: function setMaximumDETHSupply(uint256 maximumDETHSupply_) returns()
func (_StakingManager *StakingManagerTransactor) SetMaximumDETHSupply(opts *bind.TransactOpts, maximumDETHSupply_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setMaximumDETHSupply", maximumDETHSupply_)
}

// SetMaximumDETHSupply is a paid mutator transaction binding the contract method 0x1d2d35ce.
//
// Solidity: function setMaximumDETHSupply(uint256 maximumDETHSupply_) returns()
func (_StakingManager *StakingManagerSession) SetMaximumDETHSupply(maximumDETHSupply_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMaximumDETHSupply(&_StakingManager.TransactOpts, maximumDETHSupply_)
}

// SetMaximumDETHSupply is a paid mutator transaction binding the contract method 0x1d2d35ce.
//
// Solidity: function setMaximumDETHSupply(uint256 maximumDETHSupply_) returns()
func (_StakingManager *StakingManagerTransactorSession) SetMaximumDETHSupply(maximumDETHSupply_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMaximumDETHSupply(&_StakingManager.TransactOpts, maximumDETHSupply_)
}

// SetMinimumDepositAmount is a paid mutator transaction binding the contract method 0xaab483d6.
//
// Solidity: function setMinimumDepositAmount(uint256 minimumDepositAmount_) returns()
func (_StakingManager *StakingManagerTransactor) SetMinimumDepositAmount(opts *bind.TransactOpts, minimumDepositAmount_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setMinimumDepositAmount", minimumDepositAmount_)
}

// SetMinimumDepositAmount is a paid mutator transaction binding the contract method 0xaab483d6.
//
// Solidity: function setMinimumDepositAmount(uint256 minimumDepositAmount_) returns()
func (_StakingManager *StakingManagerSession) SetMinimumDepositAmount(minimumDepositAmount_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMinimumDepositAmount(&_StakingManager.TransactOpts, minimumDepositAmount_)
}

// SetMinimumDepositAmount is a paid mutator transaction binding the contract method 0xaab483d6.
//
// Solidity: function setMinimumDepositAmount(uint256 minimumDepositAmount_) returns()
func (_StakingManager *StakingManagerTransactorSession) SetMinimumDepositAmount(minimumDepositAmount_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMinimumDepositAmount(&_StakingManager.TransactOpts, minimumDepositAmount_)
}

// SetMinimumUnstakeBound is a paid mutator transaction binding the contract method 0x99dd1deb.
//
// Solidity: function setMinimumUnstakeBound(uint256 minimumUnstakeBound_) returns()
func (_StakingManager *StakingManagerTransactor) SetMinimumUnstakeBound(opts *bind.TransactOpts, minimumUnstakeBound_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setMinimumUnstakeBound", minimumUnstakeBound_)
}

// SetMinimumUnstakeBound is a paid mutator transaction binding the contract method 0x99dd1deb.
//
// Solidity: function setMinimumUnstakeBound(uint256 minimumUnstakeBound_) returns()
func (_StakingManager *StakingManagerSession) SetMinimumUnstakeBound(minimumUnstakeBound_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMinimumUnstakeBound(&_StakingManager.TransactOpts, minimumUnstakeBound_)
}

// SetMinimumUnstakeBound is a paid mutator transaction binding the contract method 0x99dd1deb.
//
// Solidity: function setMinimumUnstakeBound(uint256 minimumUnstakeBound_) returns()
func (_StakingManager *StakingManagerTransactorSession) SetMinimumUnstakeBound(minimumUnstakeBound_ *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetMinimumUnstakeBound(&_StakingManager.TransactOpts, minimumUnstakeBound_)
}

// SetStakingAllowlist is a paid mutator transaction binding the contract method 0x04f36cc2.
//
// Solidity: function setStakingAllowlist(bool isStakingAllowlist_) returns()
func (_StakingManager *StakingManagerTransactor) SetStakingAllowlist(opts *bind.TransactOpts, isStakingAllowlist_ bool) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setStakingAllowlist", isStakingAllowlist_)
}

// SetStakingAllowlist is a paid mutator transaction binding the contract method 0x04f36cc2.
//
// Solidity: function setStakingAllowlist(bool isStakingAllowlist_) returns()
func (_StakingManager *StakingManagerSession) SetStakingAllowlist(isStakingAllowlist_ bool) (*types.Transaction, error) {
	return _StakingManager.Contract.SetStakingAllowlist(&_StakingManager.TransactOpts, isStakingAllowlist_)
}

// SetStakingAllowlist is a paid mutator transaction binding the contract method 0x04f36cc2.
//
// Solidity: function setStakingAllowlist(bool isStakingAllowlist_) returns()
func (_StakingManager *StakingManagerTransactorSession) SetStakingAllowlist(isStakingAllowlist_ bool) (*types.Transaction, error) {
	return _StakingManager.Contract.SetStakingAllowlist(&_StakingManager.TransactOpts, isStakingAllowlist_)
}

// SetWithdrawalWallet is a paid mutator transaction binding the contract method 0x75796f76.
//
// Solidity: function setWithdrawalWallet(address withdrawalWallet_) returns()
func (_StakingManager *StakingManagerTransactor) SetWithdrawalWallet(opts *bind.TransactOpts, withdrawalWallet_ common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setWithdrawalWallet", withdrawalWallet_)
}

// SetWithdrawalWallet is a paid mutator transaction binding the contract method 0x75796f76.
//
// Solidity: function setWithdrawalWallet(address withdrawalWallet_) returns()
func (_StakingManager *StakingManagerSession) SetWithdrawalWallet(withdrawalWallet_ common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetWithdrawalWallet(&_StakingManager.TransactOpts, withdrawalWallet_)
}

// SetWithdrawalWallet is a paid mutator transaction binding the contract method 0x75796f76.
//
// Solidity: function setWithdrawalWallet(address withdrawalWallet_) returns()
func (_StakingManager *StakingManagerTransactorSession) SetWithdrawalWallet(withdrawalWallet_ common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetWithdrawalWallet(&_StakingManager.TransactOpts, withdrawalWallet_)
}

// Stake is a paid mutator transaction binding the contract method 0x37a6c881.
//
// Solidity: function stake(uint256 stakeAmount, (address,uint256)[] batchMints) payable returns()
func (_StakingManager *StakingManagerTransactor) Stake(opts *bind.TransactOpts, stakeAmount *big.Int, batchMints []IDETHBatchMint) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "stake", stakeAmount, batchMints)
}

// Stake is a paid mutator transaction binding the contract method 0x37a6c881.
//
// Solidity: function stake(uint256 stakeAmount, (address,uint256)[] batchMints) payable returns()
func (_StakingManager *StakingManagerSession) Stake(stakeAmount *big.Int, batchMints []IDETHBatchMint) (*types.Transaction, error) {
	return _StakingManager.Contract.Stake(&_StakingManager.TransactOpts, stakeAmount, batchMints)
}

// Stake is a paid mutator transaction binding the contract method 0x37a6c881.
//
// Solidity: function stake(uint256 stakeAmount, (address,uint256)[] batchMints) payable returns()
func (_StakingManager *StakingManagerTransactorSession) Stake(stakeAmount *big.Int, batchMints []IDETHBatchMint) (*types.Transaction, error) {
	return _StakingManager.Contract.Stake(&_StakingManager.TransactOpts, stakeAmount, batchMints)
}

// TopUp is a paid mutator transaction binding the contract method 0xdc29f1de.
//
// Solidity: function topUp() payable returns()
func (_StakingManager *StakingManagerTransactor) TopUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "topUp")
}

// TopUp is a paid mutator transaction binding the contract method 0xdc29f1de.
//
// Solidity: function topUp() payable returns()
func (_StakingManager *StakingManagerSession) TopUp() (*types.Transaction, error) {
	return _StakingManager.Contract.TopUp(&_StakingManager.TransactOpts)
}

// TopUp is a paid mutator transaction binding the contract method 0xdc29f1de.
//
// Solidity: function topUp() payable returns()
func (_StakingManager *StakingManagerTransactorSession) TopUp() (*types.Transaction, error) {
	return _StakingManager.Contract.TopUp(&_StakingManager.TransactOpts)
}

// TotalControlled is a paid mutator transaction binding the contract method 0x5940d90b.
//
// Solidity: function totalControlled() returns(uint256)
func (_StakingManager *StakingManagerTransactor) TotalControlled(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "totalControlled")
}

// TotalControlled is a paid mutator transaction binding the contract method 0x5940d90b.
//
// Solidity: function totalControlled() returns(uint256)
func (_StakingManager *StakingManagerSession) TotalControlled() (*types.Transaction, error) {
	return _StakingManager.Contract.TotalControlled(&_StakingManager.TransactOpts)
}

// TotalControlled is a paid mutator transaction binding the contract method 0x5940d90b.
//
// Solidity: function totalControlled() returns(uint256)
func (_StakingManager *StakingManagerTransactorSession) TotalControlled() (*types.Transaction, error) {
	return _StakingManager.Contract.TotalControlled(&_StakingManager.TransactOpts)
}

// UnstakeRequest is a paid mutator transaction binding the contract method 0x12e9ead6.
//
// Solidity: function unstakeRequest(uint128 dethAmount, uint128 minETHAmount, address l2Strategy, uint256 destChainId) returns()
func (_StakingManager *StakingManagerTransactor) UnstakeRequest(opts *bind.TransactOpts, dethAmount *big.Int, minETHAmount *big.Int, l2Strategy common.Address, destChainId *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "unstakeRequest", dethAmount, minETHAmount, l2Strategy, destChainId)
}

// UnstakeRequest is a paid mutator transaction binding the contract method 0x12e9ead6.
//
// Solidity: function unstakeRequest(uint128 dethAmount, uint128 minETHAmount, address l2Strategy, uint256 destChainId) returns()
func (_StakingManager *StakingManagerSession) UnstakeRequest(dethAmount *big.Int, minETHAmount *big.Int, l2Strategy common.Address, destChainId *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.UnstakeRequest(&_StakingManager.TransactOpts, dethAmount, minETHAmount, l2Strategy, destChainId)
}

// UnstakeRequest is a paid mutator transaction binding the contract method 0x12e9ead6.
//
// Solidity: function unstakeRequest(uint128 dethAmount, uint128 minETHAmount, address l2Strategy, uint256 destChainId) returns()
func (_StakingManager *StakingManagerTransactorSession) UnstakeRequest(dethAmount *big.Int, minETHAmount *big.Int, l2Strategy common.Address, destChainId *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.UnstakeRequest(&_StakingManager.TransactOpts, dethAmount, minETHAmount, l2Strategy, destChainId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StakingManager *StakingManagerTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StakingManager *StakingManagerSession) Withdraw() (*types.Transaction, error) {
	return _StakingManager.Contract.Withdraw(&_StakingManager.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StakingManager *StakingManagerTransactorSession) Withdraw() (*types.Transaction, error) {
	return _StakingManager.Contract.Withdraw(&_StakingManager.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakingManager *StakingManagerTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _StakingManager.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakingManager *StakingManagerSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StakingManager.Contract.Fallback(&_StakingManager.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakingManager *StakingManagerTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StakingManager.Contract.Fallback(&_StakingManager.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingManager *StakingManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingManager *StakingManagerSession) Receive() (*types.Transaction, error) {
	return _StakingManager.Contract.Receive(&_StakingManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingManager *StakingManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _StakingManager.Contract.Receive(&_StakingManager.TransactOpts)
}

// StakingManagerAllocatedETHToDepositsIterator is returned from FilterAllocatedETHToDeposits and is used to iterate over the raw logs and unpacked data for AllocatedETHToDeposits events raised by the StakingManager contract.
type StakingManagerAllocatedETHToDepositsIterator struct {
	Event *StakingManagerAllocatedETHToDeposits // Event containing the contract specifics and raw log

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
func (it *StakingManagerAllocatedETHToDepositsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerAllocatedETHToDeposits)
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
		it.Event = new(StakingManagerAllocatedETHToDeposits)
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
func (it *StakingManagerAllocatedETHToDepositsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerAllocatedETHToDepositsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerAllocatedETHToDeposits represents a AllocatedETHToDeposits event raised by the StakingManager contract.
type StakingManagerAllocatedETHToDeposits struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAllocatedETHToDeposits is a free log retrieval operation binding the contract event 0x9d04ecb465d2c8754acb798a22293dd26215a1c2f7a2a56607afa215c1aadc77.
//
// Solidity: event AllocatedETHToDeposits(uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterAllocatedETHToDeposits(opts *bind.FilterOpts) (*StakingManagerAllocatedETHToDepositsIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "AllocatedETHToDeposits")
	if err != nil {
		return nil, err
	}
	return &StakingManagerAllocatedETHToDepositsIterator{contract: _StakingManager.contract, event: "AllocatedETHToDeposits", logs: logs, sub: sub}, nil
}

// WatchAllocatedETHToDeposits is a free log subscription operation binding the contract event 0x9d04ecb465d2c8754acb798a22293dd26215a1c2f7a2a56607afa215c1aadc77.
//
// Solidity: event AllocatedETHToDeposits(uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchAllocatedETHToDeposits(opts *bind.WatchOpts, sink chan<- *StakingManagerAllocatedETHToDeposits) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "AllocatedETHToDeposits")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerAllocatedETHToDeposits)
				if err := _StakingManager.contract.UnpackLog(event, "AllocatedETHToDeposits", log); err != nil {
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

// ParseAllocatedETHToDeposits is a log parse operation binding the contract event 0x9d04ecb465d2c8754acb798a22293dd26215a1c2f7a2a56607afa215c1aadc77.
//
// Solidity: event AllocatedETHToDeposits(uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseAllocatedETHToDeposits(log types.Log) (*StakingManagerAllocatedETHToDeposits, error) {
	event := new(StakingManagerAllocatedETHToDeposits)
	if err := _StakingManager.contract.UnpackLog(event, "AllocatedETHToDeposits", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerAllocatedETHToUnstakeRequestsManagerIterator is returned from FilterAllocatedETHToUnstakeRequestsManager and is used to iterate over the raw logs and unpacked data for AllocatedETHToUnstakeRequestsManager events raised by the StakingManager contract.
type StakingManagerAllocatedETHToUnstakeRequestsManagerIterator struct {
	Event *StakingManagerAllocatedETHToUnstakeRequestsManager // Event containing the contract specifics and raw log

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
func (it *StakingManagerAllocatedETHToUnstakeRequestsManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerAllocatedETHToUnstakeRequestsManager)
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
		it.Event = new(StakingManagerAllocatedETHToUnstakeRequestsManager)
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
func (it *StakingManagerAllocatedETHToUnstakeRequestsManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerAllocatedETHToUnstakeRequestsManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerAllocatedETHToUnstakeRequestsManager represents a AllocatedETHToUnstakeRequestsManager event raised by the StakingManager contract.
type StakingManagerAllocatedETHToUnstakeRequestsManager struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAllocatedETHToUnstakeRequestsManager is a free log retrieval operation binding the contract event 0xfe89805cf5299ef9fbd1d1ddefb8dcc3fa9408064d2ea31e3fca6565768f5217.
//
// Solidity: event AllocatedETHToUnstakeRequestsManager(uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterAllocatedETHToUnstakeRequestsManager(opts *bind.FilterOpts) (*StakingManagerAllocatedETHToUnstakeRequestsManagerIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "AllocatedETHToUnstakeRequestsManager")
	if err != nil {
		return nil, err
	}
	return &StakingManagerAllocatedETHToUnstakeRequestsManagerIterator{contract: _StakingManager.contract, event: "AllocatedETHToUnstakeRequestsManager", logs: logs, sub: sub}, nil
}

// WatchAllocatedETHToUnstakeRequestsManager is a free log subscription operation binding the contract event 0xfe89805cf5299ef9fbd1d1ddefb8dcc3fa9408064d2ea31e3fca6565768f5217.
//
// Solidity: event AllocatedETHToUnstakeRequestsManager(uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchAllocatedETHToUnstakeRequestsManager(opts *bind.WatchOpts, sink chan<- *StakingManagerAllocatedETHToUnstakeRequestsManager) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "AllocatedETHToUnstakeRequestsManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerAllocatedETHToUnstakeRequestsManager)
				if err := _StakingManager.contract.UnpackLog(event, "AllocatedETHToUnstakeRequestsManager", log); err != nil {
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

// ParseAllocatedETHToUnstakeRequestsManager is a log parse operation binding the contract event 0xfe89805cf5299ef9fbd1d1ddefb8dcc3fa9408064d2ea31e3fca6565768f5217.
//
// Solidity: event AllocatedETHToUnstakeRequestsManager(uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseAllocatedETHToUnstakeRequestsManager(log types.Log) (*StakingManagerAllocatedETHToUnstakeRequestsManager, error) {
	event := new(StakingManagerAllocatedETHToUnstakeRequestsManager)
	if err := _StakingManager.contract.UnpackLog(event, "AllocatedETHToUnstakeRequestsManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakingManager contract.
type StakingManagerInitializedIterator struct {
	Event *StakingManagerInitialized // Event containing the contract specifics and raw log

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
func (it *StakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerInitialized)
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
		it.Event = new(StakingManagerInitialized)
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
func (it *StakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerInitialized represents a Initialized event raised by the StakingManager contract.
type StakingManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakingManager *StakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingManagerInitializedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingManagerInitializedIterator{contract: _StakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakingManager *StakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerInitialized)
				if err := _StakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseInitialized(log types.Log) (*StakingManagerInitialized, error) {
	event := new(StakingManagerInitialized)
	if err := _StakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerProtocolConfigChangedIterator is returned from FilterProtocolConfigChanged and is used to iterate over the raw logs and unpacked data for ProtocolConfigChanged events raised by the StakingManager contract.
type StakingManagerProtocolConfigChangedIterator struct {
	Event *StakingManagerProtocolConfigChanged // Event containing the contract specifics and raw log

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
func (it *StakingManagerProtocolConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerProtocolConfigChanged)
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
		it.Event = new(StakingManagerProtocolConfigChanged)
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
func (it *StakingManagerProtocolConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerProtocolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerProtocolConfigChanged represents a ProtocolConfigChanged event raised by the StakingManager contract.
type StakingManagerProtocolConfigChanged struct {
	SetterSelector  [4]byte
	SetterSignature string
	Value           []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProtocolConfigChanged is a free log retrieval operation binding the contract event 0x01d854e8dde9402801a4c6f2840193465752abfad61e0bb7c4258d526ae42e74.
//
// Solidity: event ProtocolConfigChanged(bytes4 indexed setterSelector, string setterSignature, bytes value)
func (_StakingManager *StakingManagerFilterer) FilterProtocolConfigChanged(opts *bind.FilterOpts, setterSelector [][4]byte) (*StakingManagerProtocolConfigChangedIterator, error) {

	var setterSelectorRule []interface{}
	for _, setterSelectorItem := range setterSelector {
		setterSelectorRule = append(setterSelectorRule, setterSelectorItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "ProtocolConfigChanged", setterSelectorRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerProtocolConfigChangedIterator{contract: _StakingManager.contract, event: "ProtocolConfigChanged", logs: logs, sub: sub}, nil
}

// WatchProtocolConfigChanged is a free log subscription operation binding the contract event 0x01d854e8dde9402801a4c6f2840193465752abfad61e0bb7c4258d526ae42e74.
//
// Solidity: event ProtocolConfigChanged(bytes4 indexed setterSelector, string setterSignature, bytes value)
func (_StakingManager *StakingManagerFilterer) WatchProtocolConfigChanged(opts *bind.WatchOpts, sink chan<- *StakingManagerProtocolConfigChanged, setterSelector [][4]byte) (event.Subscription, error) {

	var setterSelectorRule []interface{}
	for _, setterSelectorItem := range setterSelector {
		setterSelectorRule = append(setterSelectorRule, setterSelectorItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "ProtocolConfigChanged", setterSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerProtocolConfigChanged)
				if err := _StakingManager.contract.UnpackLog(event, "ProtocolConfigChanged", log); err != nil {
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

// ParseProtocolConfigChanged is a log parse operation binding the contract event 0x01d854e8dde9402801a4c6f2840193465752abfad61e0bb7c4258d526ae42e74.
//
// Solidity: event ProtocolConfigChanged(bytes4 indexed setterSelector, string setterSignature, bytes value)
func (_StakingManager *StakingManagerFilterer) ParseProtocolConfigChanged(log types.Log) (*StakingManagerProtocolConfigChanged, error) {
	event := new(StakingManagerProtocolConfigChanged)
	if err := _StakingManager.contract.UnpackLog(event, "ProtocolConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerReturnsReceivedIterator is returned from FilterReturnsReceived and is used to iterate over the raw logs and unpacked data for ReturnsReceived events raised by the StakingManager contract.
type StakingManagerReturnsReceivedIterator struct {
	Event *StakingManagerReturnsReceived // Event containing the contract specifics and raw log

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
func (it *StakingManagerReturnsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerReturnsReceived)
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
		it.Event = new(StakingManagerReturnsReceived)
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
func (it *StakingManagerReturnsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerReturnsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerReturnsReceived represents a ReturnsReceived event raised by the StakingManager contract.
type StakingManagerReturnsReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReturnsReceived is a free log retrieval operation binding the contract event 0x4cbb9d73b003a252cee3f2ee51d8d65a562af35eebb23730dd4a76d68127b370.
//
// Solidity: event ReturnsReceived(uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterReturnsReceived(opts *bind.FilterOpts) (*StakingManagerReturnsReceivedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "ReturnsReceived")
	if err != nil {
		return nil, err
	}
	return &StakingManagerReturnsReceivedIterator{contract: _StakingManager.contract, event: "ReturnsReceived", logs: logs, sub: sub}, nil
}

// WatchReturnsReceived is a free log subscription operation binding the contract event 0x4cbb9d73b003a252cee3f2ee51d8d65a562af35eebb23730dd4a76d68127b370.
//
// Solidity: event ReturnsReceived(uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchReturnsReceived(opts *bind.WatchOpts, sink chan<- *StakingManagerReturnsReceived) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "ReturnsReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerReturnsReceived)
				if err := _StakingManager.contract.UnpackLog(event, "ReturnsReceived", log); err != nil {
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

// ParseReturnsReceived is a log parse operation binding the contract event 0x4cbb9d73b003a252cee3f2ee51d8d65a562af35eebb23730dd4a76d68127b370.
//
// Solidity: event ReturnsReceived(uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseReturnsReceived(log types.Log) (*StakingManagerReturnsReceived, error) {
	event := new(StakingManagerReturnsReceived)
	if err := _StakingManager.contract.UnpackLog(event, "ReturnsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StakingManager contract.
type StakingManagerRoleAdminChangedIterator struct {
	Event *StakingManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerRoleAdminChanged)
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
		it.Event = new(StakingManagerRoleAdminChanged)
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
func (it *StakingManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerRoleAdminChanged represents a RoleAdminChanged event raised by the StakingManager contract.
type StakingManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingManager *StakingManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakingManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerRoleAdminChangedIterator{contract: _StakingManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingManager *StakingManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerRoleAdminChanged)
				if err := _StakingManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseRoleAdminChanged(log types.Log) (*StakingManagerRoleAdminChanged, error) {
	event := new(StakingManagerRoleAdminChanged)
	if err := _StakingManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StakingManager contract.
type StakingManagerRoleGrantedIterator struct {
	Event *StakingManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakingManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerRoleGranted)
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
		it.Event = new(StakingManagerRoleGranted)
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
func (it *StakingManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerRoleGranted represents a RoleGranted event raised by the StakingManager contract.
type StakingManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerRoleGrantedIterator{contract: _StakingManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakingManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerRoleGranted)
				if err := _StakingManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseRoleGranted(log types.Log) (*StakingManagerRoleGranted, error) {
	event := new(StakingManagerRoleGranted)
	if err := _StakingManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StakingManager contract.
type StakingManagerRoleRevokedIterator struct {
	Event *StakingManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakingManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerRoleRevoked)
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
		it.Event = new(StakingManagerRoleRevoked)
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
func (it *StakingManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerRoleRevoked represents a RoleRevoked event raised by the StakingManager contract.
type StakingManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerRoleRevokedIterator{contract: _StakingManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakingManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerRoleRevoked)
				if err := _StakingManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseRoleRevoked(log types.Log) (*StakingManagerRoleRevoked, error) {
	event := new(StakingManagerRoleRevoked)
	if err := _StakingManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the StakingManager contract.
type StakingManagerStakedIterator struct {
	Event *StakingManagerStaked // Event containing the contract specifics and raw log

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
func (it *StakingManagerStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerStaked)
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
		it.Event = new(StakingManagerStaked)
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
func (it *StakingManagerStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerStaked represents a Staked event raised by the StakingManager contract.
type StakingManagerStaked struct {
	Staker     common.Address
	EthAmount  *big.Int
	DETHAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed staker, uint256 ethAmount, uint256 dETHAmount)
func (_StakingManager *StakingManagerFilterer) FilterStaked(opts *bind.FilterOpts, staker []common.Address) (*StakingManagerStakedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerStakedIterator{contract: _StakingManager.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed staker, uint256 ethAmount, uint256 dETHAmount)
func (_StakingManager *StakingManagerFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingManagerStaked, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerStaked)
				if err := _StakingManager.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed staker, uint256 ethAmount, uint256 dETHAmount)
func (_StakingManager *StakingManagerFilterer) ParseStaked(log types.Log) (*StakingManagerStaked, error) {
	event := new(StakingManagerStaked)
	if err := _StakingManager.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUnstakeLaveAmountIterator is returned from FilterUnstakeLaveAmount and is used to iterate over the raw logs and unpacked data for UnstakeLaveAmount events raised by the StakingManager contract.
type StakingManagerUnstakeLaveAmountIterator struct {
	Event *StakingManagerUnstakeLaveAmount // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnstakeLaveAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnstakeLaveAmount)
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
		it.Event = new(StakingManagerUnstakeLaveAmount)
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
func (it *StakingManagerUnstakeLaveAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnstakeLaveAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnstakeLaveAmount represents a UnstakeLaveAmount event raised by the StakingManager contract.
type StakingManagerUnstakeLaveAmount struct {
	Staker     common.Address
	DETHLocked *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnstakeLaveAmount is a free log retrieval operation binding the contract event 0x1f4bdd7f902a9da109e5ee8424af475df1f64b7b7a8d36e07d5f97caaa5f19ec.
//
// Solidity: event UnstakeLaveAmount(address indexed staker, uint256 dETHLocked)
func (_StakingManager *StakingManagerFilterer) FilterUnstakeLaveAmount(opts *bind.FilterOpts, staker []common.Address) (*StakingManagerUnstakeLaveAmountIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnstakeLaveAmount", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnstakeLaveAmountIterator{contract: _StakingManager.contract, event: "UnstakeLaveAmount", logs: logs, sub: sub}, nil
}

// WatchUnstakeLaveAmount is a free log subscription operation binding the contract event 0x1f4bdd7f902a9da109e5ee8424af475df1f64b7b7a8d36e07d5f97caaa5f19ec.
//
// Solidity: event UnstakeLaveAmount(address indexed staker, uint256 dETHLocked)
func (_StakingManager *StakingManagerFilterer) WatchUnstakeLaveAmount(opts *bind.WatchOpts, sink chan<- *StakingManagerUnstakeLaveAmount, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnstakeLaveAmount", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnstakeLaveAmount)
				if err := _StakingManager.contract.UnpackLog(event, "UnstakeLaveAmount", log); err != nil {
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

// ParseUnstakeLaveAmount is a log parse operation binding the contract event 0x1f4bdd7f902a9da109e5ee8424af475df1f64b7b7a8d36e07d5f97caaa5f19ec.
//
// Solidity: event UnstakeLaveAmount(address indexed staker, uint256 dETHLocked)
func (_StakingManager *StakingManagerFilterer) ParseUnstakeLaveAmount(log types.Log) (*StakingManagerUnstakeLaveAmount, error) {
	event := new(StakingManagerUnstakeLaveAmount)
	if err := _StakingManager.contract.UnpackLog(event, "UnstakeLaveAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUnstakeRequestClaimedIterator is returned from FilterUnstakeRequestClaimed and is used to iterate over the raw logs and unpacked data for UnstakeRequestClaimed events raised by the StakingManager contract.
type StakingManagerUnstakeRequestClaimedIterator struct {
	Event *StakingManagerUnstakeRequestClaimed // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnstakeRequestClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnstakeRequestClaimed)
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
		it.Event = new(StakingManagerUnstakeRequestClaimed)
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
func (it *StakingManagerUnstakeRequestClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnstakeRequestClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnstakeRequestClaimed represents a UnstakeRequestClaimed event raised by the StakingManager contract.
type StakingManagerUnstakeRequestClaimed struct {
	Staker        common.Address
	L2Strategy    common.Address
	Bridge        common.Address
	SourceChainId *big.Int
	DestChainId   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnstakeRequestClaimed is a free log retrieval operation binding the contract event 0xdefb23229f98ca9823b290c7a4ca8908d69e2a25a6fa9e0b65b3ffad35798eb3.
//
// Solidity: event UnstakeRequestClaimed(address indexed staker, address indexed l2Strategy, address indexed bridge, uint256 sourceChainId, uint256 destChainId)
func (_StakingManager *StakingManagerFilterer) FilterUnstakeRequestClaimed(opts *bind.FilterOpts, staker []common.Address, l2Strategy []common.Address, bridge []common.Address) (*StakingManagerUnstakeRequestClaimedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var l2StrategyRule []interface{}
	for _, l2StrategyItem := range l2Strategy {
		l2StrategyRule = append(l2StrategyRule, l2StrategyItem)
	}
	var bridgeRule []interface{}
	for _, bridgeItem := range bridge {
		bridgeRule = append(bridgeRule, bridgeItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnstakeRequestClaimed", stakerRule, l2StrategyRule, bridgeRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnstakeRequestClaimedIterator{contract: _StakingManager.contract, event: "UnstakeRequestClaimed", logs: logs, sub: sub}, nil
}

// WatchUnstakeRequestClaimed is a free log subscription operation binding the contract event 0xdefb23229f98ca9823b290c7a4ca8908d69e2a25a6fa9e0b65b3ffad35798eb3.
//
// Solidity: event UnstakeRequestClaimed(address indexed staker, address indexed l2Strategy, address indexed bridge, uint256 sourceChainId, uint256 destChainId)
func (_StakingManager *StakingManagerFilterer) WatchUnstakeRequestClaimed(opts *bind.WatchOpts, sink chan<- *StakingManagerUnstakeRequestClaimed, staker []common.Address, l2Strategy []common.Address, bridge []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var l2StrategyRule []interface{}
	for _, l2StrategyItem := range l2Strategy {
		l2StrategyRule = append(l2StrategyRule, l2StrategyItem)
	}
	var bridgeRule []interface{}
	for _, bridgeItem := range bridge {
		bridgeRule = append(bridgeRule, bridgeItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnstakeRequestClaimed", stakerRule, l2StrategyRule, bridgeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnstakeRequestClaimed)
				if err := _StakingManager.contract.UnpackLog(event, "UnstakeRequestClaimed", log); err != nil {
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

// ParseUnstakeRequestClaimed is a log parse operation binding the contract event 0xdefb23229f98ca9823b290c7a4ca8908d69e2a25a6fa9e0b65b3ffad35798eb3.
//
// Solidity: event UnstakeRequestClaimed(address indexed staker, address indexed l2Strategy, address indexed bridge, uint256 sourceChainId, uint256 destChainId)
func (_StakingManager *StakingManagerFilterer) ParseUnstakeRequestClaimed(log types.Log) (*StakingManagerUnstakeRequestClaimed, error) {
	event := new(StakingManagerUnstakeRequestClaimed)
	if err := _StakingManager.contract.UnpackLog(event, "UnstakeRequestClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUnstakeRequestedIterator is returned from FilterUnstakeRequested and is used to iterate over the raw logs and unpacked data for UnstakeRequested events raised by the StakingManager contract.
type StakingManagerUnstakeRequestedIterator struct {
	Event *StakingManagerUnstakeRequested // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnstakeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnstakeRequested)
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
		it.Event = new(StakingManagerUnstakeRequested)
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
func (it *StakingManagerUnstakeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnstakeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnstakeRequested represents a UnstakeRequested event raised by the StakingManager contract.
type StakingManagerUnstakeRequested struct {
	Staker      common.Address
	L2Strategy  common.Address
	EthAmount   *big.Int
	DETHLocked  *big.Int
	DestChainId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnstakeRequested is a free log retrieval operation binding the contract event 0x7fd59a2c98b4d595be25d7976e588f7b6189ace6d41522f11868d951c2934367.
//
// Solidity: event UnstakeRequested(address indexed staker, address indexed l2Strategy, uint256 ethAmount, uint256 dETHLocked, uint256 destChainId)
func (_StakingManager *StakingManagerFilterer) FilterUnstakeRequested(opts *bind.FilterOpts, staker []common.Address, l2Strategy []common.Address) (*StakingManagerUnstakeRequestedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var l2StrategyRule []interface{}
	for _, l2StrategyItem := range l2Strategy {
		l2StrategyRule = append(l2StrategyRule, l2StrategyItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnstakeRequested", stakerRule, l2StrategyRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnstakeRequestedIterator{contract: _StakingManager.contract, event: "UnstakeRequested", logs: logs, sub: sub}, nil
}

// WatchUnstakeRequested is a free log subscription operation binding the contract event 0x7fd59a2c98b4d595be25d7976e588f7b6189ace6d41522f11868d951c2934367.
//
// Solidity: event UnstakeRequested(address indexed staker, address indexed l2Strategy, uint256 ethAmount, uint256 dETHLocked, uint256 destChainId)
func (_StakingManager *StakingManagerFilterer) WatchUnstakeRequested(opts *bind.WatchOpts, sink chan<- *StakingManagerUnstakeRequested, staker []common.Address, l2Strategy []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var l2StrategyRule []interface{}
	for _, l2StrategyItem := range l2Strategy {
		l2StrategyRule = append(l2StrategyRule, l2StrategyItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnstakeRequested", stakerRule, l2StrategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnstakeRequested)
				if err := _StakingManager.contract.UnpackLog(event, "UnstakeRequested", log); err != nil {
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

// ParseUnstakeRequested is a log parse operation binding the contract event 0x7fd59a2c98b4d595be25d7976e588f7b6189ace6d41522f11868d951c2934367.
//
// Solidity: event UnstakeRequested(address indexed staker, address indexed l2Strategy, uint256 ethAmount, uint256 dETHLocked, uint256 destChainId)
func (_StakingManager *StakingManagerFilterer) ParseUnstakeRequested(log types.Log) (*StakingManagerUnstakeRequested, error) {
	event := new(StakingManagerUnstakeRequested)
	if err := _StakingManager.contract.UnpackLog(event, "UnstakeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerValidatorInitiatedIterator is returned from FilterValidatorInitiated and is used to iterate over the raw logs and unpacked data for ValidatorInitiated events raised by the StakingManager contract.
type StakingManagerValidatorInitiatedIterator struct {
	Event *StakingManagerValidatorInitiated // Event containing the contract specifics and raw log

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
func (it *StakingManagerValidatorInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerValidatorInitiated)
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
		it.Event = new(StakingManagerValidatorInitiated)
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
func (it *StakingManagerValidatorInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerValidatorInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerValidatorInitiated represents a ValidatorInitiated event raised by the StakingManager contract.
type StakingManagerValidatorInitiated struct {
	Id              [32]byte
	OperatorID      *big.Int
	Pubkey          []byte
	AmountDeposited *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorInitiated is a free log retrieval operation binding the contract event 0x15f16c2e13e50235799a97b981bf4a66c8cd86051f06aca745c5ff26f39b330e.
//
// Solidity: event ValidatorInitiated(bytes32 indexed id, uint256 indexed operatorID, bytes pubkey, uint256 amountDeposited)
func (_StakingManager *StakingManagerFilterer) FilterValidatorInitiated(opts *bind.FilterOpts, id [][32]byte, operatorID []*big.Int) (*StakingManagerValidatorInitiatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var operatorIDRule []interface{}
	for _, operatorIDItem := range operatorID {
		operatorIDRule = append(operatorIDRule, operatorIDItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "ValidatorInitiated", idRule, operatorIDRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerValidatorInitiatedIterator{contract: _StakingManager.contract, event: "ValidatorInitiated", logs: logs, sub: sub}, nil
}

// WatchValidatorInitiated is a free log subscription operation binding the contract event 0x15f16c2e13e50235799a97b981bf4a66c8cd86051f06aca745c5ff26f39b330e.
//
// Solidity: event ValidatorInitiated(bytes32 indexed id, uint256 indexed operatorID, bytes pubkey, uint256 amountDeposited)
func (_StakingManager *StakingManagerFilterer) WatchValidatorInitiated(opts *bind.WatchOpts, sink chan<- *StakingManagerValidatorInitiated, id [][32]byte, operatorID []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var operatorIDRule []interface{}
	for _, operatorIDItem := range operatorID {
		operatorIDRule = append(operatorIDRule, operatorIDItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "ValidatorInitiated", idRule, operatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerValidatorInitiated)
				if err := _StakingManager.contract.UnpackLog(event, "ValidatorInitiated", log); err != nil {
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

// ParseValidatorInitiated is a log parse operation binding the contract event 0x15f16c2e13e50235799a97b981bf4a66c8cd86051f06aca745c5ff26f39b330e.
//
// Solidity: event ValidatorInitiated(bytes32 indexed id, uint256 indexed operatorID, bytes pubkey, uint256 amountDeposited)
func (_StakingManager *StakingManagerFilterer) ParseValidatorInitiated(log types.Log) (*StakingManagerValidatorInitiated, error) {
	event := new(StakingManagerValidatorInitiated)
	if err := _StakingManager.contract.UnpackLog(event, "ValidatorInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
