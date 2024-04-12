package services

import (
	"github.com/evm-layer2/selaginella/bindings/bridge"
	"github.com/evm-layer2/selaginella/bindings/staking"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

func bindL1PoolManager(FoundingPoolAddress string, l1Client *ethclient.Client) (*bridge.L1PoolManager, *bind.BoundContract, error) {
	l1PoolParsed, err := abi.JSON(strings.NewReader(
		bridge.L1PoolManagerABI,
	))
	if err != nil {
		log.Error("selaginella parse l1 pool contract abi fail", "err", err)
		return nil, nil, err
	}

	rawL1PoolContract := bind.NewBoundContract(
		common.HexToAddress(FoundingPoolAddress), l1PoolParsed, l1Client, l1Client,
		l1Client,
	)

	l1PoolContract, err := bridge.NewL1PoolManager(common.HexToAddress(FoundingPoolAddress), l1Client)

	return l1PoolContract, rawL1PoolContract, nil
}

func bindL1StakingManager(l1StakingManagerAddr string, l1Client *ethclient.Client) (*staking.StakingManager, *bind.BoundContract, error) {
	l1StakingManagerParsed, err := abi.JSON(strings.NewReader(
		staking.StakingManagerABI,
	))
	if err != nil {
		log.Error("selaginella parse l1 staking manager contract abi fail", "err", err)
		return nil, nil, err
	}
	rawL1StakingManagerContract := bind.NewBoundContract(
		common.HexToAddress(l1StakingManagerAddr), l1StakingManagerParsed, l1Client, l1Client,
		l1Client,
	)
	l1StakingManagerContract, err := staking.NewStakingManager(common.HexToAddress(l1StakingManagerAddr), l1Client)

	return l1StakingManagerContract, rawL1StakingManagerContract, nil
}

func bindL2PoolManager(FoundingPoolAddress string, l2Client *ethclient.Client) (*bridge.L2PoolManager, *bind.BoundContract, error) {
	l2Parsed, err := abi.JSON(strings.NewReader(
		bridge.L2PoolManagerABI,
	))
	if err != nil {
		log.Error("selaginella parse l2 pool contract abi fail", "err", err)
		return nil, nil, err
	}

	rawL2PoolContract := bind.NewBoundContract(
		common.HexToAddress(FoundingPoolAddress), l2Parsed, l2Client, l2Client,
		l2Client,
	)

	l2PoolContract, err := bridge.NewL2PoolManager(common.HexToAddress(FoundingPoolAddress), l2Client)

	return l2PoolContract, rawL2PoolContract, nil
}

func bindDaStrategyBase(DaStrategyAddress string, l2Client *ethclient.Client) (*staking.StrategyBase, *bind.BoundContract, error) {
	strategyParsed, err := abi.JSON(strings.NewReader(
		staking.StrategyBaseABI,
	))
	if err != nil {
		log.Error("selaginella parse strategy contract abi fail", "err", err)
		return nil, nil, err
	}

	rawDaStrategyContract := bind.NewBoundContract(
		common.HexToAddress(DaStrategyAddress), strategyParsed, l2Client, l2Client,
		l2Client,
	)

	daStrategyContract, err := staking.NewStrategyBase(common.HexToAddress(DaStrategyAddress), l2Client)

	return daStrategyContract, rawDaStrategyContract, nil
}

func bindGamingStrategyBase(GamingStrategyAddress string, l2Client *ethclient.Client) (*staking.StrategyBase, *bind.BoundContract, error) {
	strategyParsed, err := abi.JSON(strings.NewReader(
		staking.StrategyBaseABI,
	))
	if err != nil {
		log.Error("selaginella parse strategy contract abi fail", "err", err)
		return nil, nil, err
	}

	rawGamingStrategyContract := bind.NewBoundContract(
		common.HexToAddress(GamingStrategyAddress), strategyParsed, l2Client, l2Client,
		l2Client,
	)

	gamingStrategyContract, err := staking.NewStrategyBase(common.HexToAddress(GamingStrategyAddress), l2Client)

	return gamingStrategyContract, rawGamingStrategyContract, nil
}

func bindSocialStrategyBase(SocialStrategyAddress string, l2Client *ethclient.Client) (*staking.StrategyBase, *bind.BoundContract, error) {
	strategyParsed, err := abi.JSON(strings.NewReader(
		staking.StrategyBaseABI,
	))
	if err != nil {
		log.Error("selaginella parse strategy contract abi fail", "err", err)
		return nil, nil, err
	}

	rawSocialStrategyContract := bind.NewBoundContract(
		common.HexToAddress(SocialStrategyAddress), strategyParsed, l2Client, l2Client,
		l2Client,
	)

	socialStrategyContract, err := staking.NewStrategyBase(common.HexToAddress(SocialStrategyAddress), l2Client)

	return socialStrategyContract, rawSocialStrategyContract, nil
}

func bindStrategyManager(StrategyManagerAddress string, l2Client *ethclient.Client) (*staking.StrategyManager, *bind.BoundContract, error) {
	strategyManagerParsed, err := abi.JSON(strings.NewReader(
		staking.StrategyManagerABI,
	))
	if err != nil {
		log.Error("selaginella parse strategy manager contract abi fail", "err", err)
		return nil, nil, err
	}

	rawStrategyManagerContract := bind.NewBoundContract(
		common.HexToAddress(StrategyManagerAddress), strategyManagerParsed, l2Client, l2Client,
		l2Client,
	)

	strategyManagerContract, err := staking.NewStrategyManager(common.HexToAddress(StrategyManagerAddress), l2Client)

	return strategyManagerContract, rawStrategyManagerContract, nil
}
