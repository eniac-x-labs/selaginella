package services

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"net"
	"sync"
	"sync/atomic"
	"time"

	kms "cloud.google.com/go/kms/apiv1"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	"github.com/evm-layer2/selaginella/bindings"
	"github.com/evm-layer2/selaginella/common/retry"
	"github.com/evm-layer2/selaginella/common/tasks"
	"github.com/evm-layer2/selaginella/config"
	"github.com/evm-layer2/selaginella/database"
	node "github.com/evm-layer2/selaginella/eth_client"
	"github.com/evm-layer2/selaginella/protobuf/pb"
	"github.com/evm-layer2/selaginella/sign"
)

const MaxRecvMessageSize = 1024 * 1024 * 300

type RpcServerConfig struct {
	GrpcHostname string
	GrpcPort     string
}

type HsmConfig struct {
	EnableHsm  bool
	HsmAPIName string
	HsmCreden  string
	HsmAddress string
}

type RpcServer struct {
	*RpcServerConfig
	*HsmConfig
	db                          *database.DB
	ethClients                  map[uint64]node.EthClient
	RawL1BridgeContract         *bind.BoundContract
	RawL2BridgeContract         map[uint64]*bind.BoundContract
	L1BridgeContract            *bindings.L1PoolManager
	L1BridgeContractAddress     common.Address
	L2BridgeContract            map[uint64]*bindings.L2PoolManager
	DAStrategyContract          map[uint64]*bindings.StrategyBase
	RawDAStrategyContract       map[uint64]*bind.BoundContract
	GamingStrategyContract      map[uint64]*bindings.StrategyBase
	RawGamingStrategyContract   map[uint64]*bind.BoundContract
	SocialStrategyContract      map[uint64]*bindings.StrategyBase
	RawSocialStrategyContract   map[uint64]*bind.BoundContract
	l1StakingManagerAddr        common.Address
	l1StakingManagerContract    *bindings.StakingManager
	RawL1StakingManagerContract *bind.BoundContract
	l1DETHContract              *bindings.DETH
	RawL1DETHContract           *bind.BoundContract
	StrategyManagerContract     map[uint64]*bindings.StrategyManager
	RawStrategyManagerContract  map[uint64]*bind.BoundContract
	EthAddress                  map[uint64]common.Address
	WEthAddress                 map[uint64]common.Address
	USDTAddress                 map[uint64]common.Address
	USDCAddress                 map[uint64]common.Address
	DAIAddress                  map[uint64]common.Address
	OKBAddress                  map[uint64]common.Address
	MNTAddress                  map[uint64]common.Address
	poolStartTimestamp          uint32
	poolEndTimestamp            uint32
	pb.UnimplementedBridgeServiceServer
	stopped atomic.Bool
	pb.BridgeServiceServer
	privateKey    *ecdsa.PrivateKey
	l1ChainID     uint64
	zkFairChainId uint64
	x1ChainId     uint64
	opChainId     uint64
	mantleChainId uint64
	tasks         tasks.Group
}

func (s *RpcServer) Stop(ctx context.Context) error {
	s.stopped.Store(true)
	return nil
}

func (s *RpcServer) Stopped() bool {
	return s.stopped.Load()
}

func NewRpcServer(ctx context.Context, db *database.DB, grpcCfg *RpcServerConfig, hsmCfg *HsmConfig, chainRpcCfg []*config.RPC, priKey *ecdsa.PrivateKey, chainIds config.ChainId, l1StakingManagerAddr string, l1DETHAddr string, shutdown context.CancelCauseFunc) (*RpcServer, error) {
	ethClients := make(map[uint64]node.EthClient)
	var rawL1BridgeContract *bind.BoundContract
	rawL2BridgeContracts := make(map[uint64]*bind.BoundContract)
	var l1BridgeContract *bindings.L1PoolManager
	l2BridgeContracts := make(map[uint64]*bindings.L2PoolManager)
	EthAddress := make(map[uint64]common.Address)
	WEthAddress := make(map[uint64]common.Address)
	USDTAddress := make(map[uint64]common.Address)
	USDCAddress := make(map[uint64]common.Address)
	DAIAddress := make(map[uint64]common.Address)
	OKBAddress := make(map[uint64]common.Address)
	MNTAddress := make(map[uint64]common.Address)
	rawDaStrategyContracts := make(map[uint64]*bind.BoundContract)
	rawGamingStrategyContracts := make(map[uint64]*bind.BoundContract)
	rawSocialStrategyContracts := make(map[uint64]*bind.BoundContract)
	daStrategyContracts := make(map[uint64]*bindings.StrategyBase)
	gamingStrategyContracts := make(map[uint64]*bindings.StrategyBase)
	socialStrategyContracts := make(map[uint64]*bindings.StrategyBase)
	strategyManagerContracts := make(map[uint64]*bindings.StrategyManager)
	rawStrategyManagerContracts := make(map[uint64]*bind.BoundContract)
	var l1PoolContractAddr common.Address
	var l1StakingManagerContract *bindings.StakingManager
	var rawL1StakingManagerContract *bind.BoundContract
	var l1DETHContract *bindings.DETH
	var rawL1DETHManagerContract *bind.BoundContract

	for i := range chainRpcCfg {
		if chainRpcCfg[i].ChainId == chainIds.L1ChainId {
			client, err := node.DialEthClient(ctx, chainRpcCfg[i].RpcUrl)
			if err != nil {
				log.Error("dial client fail", "error", err.Error())
				return nil, err
			}
			ethClients[chainRpcCfg[i].ChainId] = client

			l1Client, _ := node.DialEthClientWithTimeout(ctx, chainRpcCfg[i].RpcUrl, false)

			l1BridgeContract, rawL1BridgeContract, err = bindL1PoolManager(chainRpcCfg[i].FoundingPoolAddress, l1Client)
			if err != nil {
				return nil, err
			}

			l1PoolContractAddr = common.HexToAddress(chainRpcCfg[i].FoundingPoolAddress)
			l1StakingManagerContract, rawL1StakingManagerContract, err = bindL1StakingManager(l1StakingManagerAddr, l1Client)
			if err != nil {
				return nil, err
			}

			l1DETHContract, rawL1DETHManagerContract, err = bindDETH(l1DETHAddr, l1Client)
			if err != nil {
				return nil, err
			}

		} else {
			client, err := node.DialEthClient(ctx, chainRpcCfg[i].RpcUrl)
			if err != nil {
				log.Error("dial client fail", "error", err.Error())
				return nil, err
			}
			ethClients[chainRpcCfg[i].ChainId] = client

			l2Client, _ := node.DialEthClientWithTimeout(ctx, chainRpcCfg[i].RpcUrl, false)

			l2BridgeContracts[chainRpcCfg[i].ChainId], rawL2BridgeContracts[chainRpcCfg[i].ChainId], err = bindL2PoolManager(chainRpcCfg[i].FoundingPoolAddress, l2Client)

			daStrategyContracts[chainRpcCfg[i].ChainId], rawDaStrategyContracts[chainRpcCfg[i].ChainId], err = bindDaStrategyBase(chainRpcCfg[i].DaStrategyAddress, l2Client)
			gamingStrategyContracts[chainRpcCfg[i].ChainId], rawGamingStrategyContracts[chainRpcCfg[i].ChainId], err = bindGamingStrategyBase(chainRpcCfg[i].DaStrategyAddress, l2Client)
			socialStrategyContracts[chainRpcCfg[i].ChainId], rawSocialStrategyContracts[chainRpcCfg[i].ChainId], err = bindSocialStrategyBase(chainRpcCfg[i].DaStrategyAddress, l2Client)

			strategyManagerContracts[chainRpcCfg[i].ChainId], rawStrategyManagerContracts[chainRpcCfg[i].ChainId], err = bindStrategyManager(chainRpcCfg[i].StrategyManager, l2Client)

			if err != nil {
				return nil, err
			}
		}
		EthAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].EthAddress)
		WEthAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].WEthAddress)
		USDTAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].USDTAddress)
		USDCAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].USDCAddress)
		DAIAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].DAIAddress)
		OKBAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].OKBAddress)
		MNTAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].MNTAddress)

	}

	return &RpcServer{
		RpcServerConfig:             grpcCfg,
		db:                          db,
		HsmConfig:                   hsmCfg,
		ethClients:                  ethClients,
		RawL1BridgeContract:         rawL1BridgeContract,
		RawL2BridgeContract:         rawL2BridgeContracts,
		L1BridgeContract:            l1BridgeContract,
		L1BridgeContractAddress:     l1PoolContractAddr,
		L2BridgeContract:            l2BridgeContracts,
		DAStrategyContract:          daStrategyContracts,
		GamingStrategyContract:      gamingStrategyContracts,
		SocialStrategyContract:      socialStrategyContracts,
		RawDAStrategyContract:       rawDaStrategyContracts,
		RawGamingStrategyContract:   rawGamingStrategyContracts,
		RawSocialStrategyContract:   rawSocialStrategyContracts,
		l1StakingManagerAddr:        common.HexToAddress(l1StakingManagerAddr),
		l1StakingManagerContract:    l1StakingManagerContract,
		RawL1StakingManagerContract: rawL1StakingManagerContract,
		StrategyManagerContract:     strategyManagerContracts,
		RawStrategyManagerContract:  rawStrategyManagerContracts,
		l1DETHContract:              l1DETHContract,
		RawL1DETHContract:           rawL1DETHManagerContract,
		EthAddress:                  EthAddress,
		WEthAddress:                 WEthAddress,
		USDTAddress:                 USDTAddress,
		USDCAddress:                 USDCAddress,
		DAIAddress:                  DAIAddress,
		OKBAddress:                  OKBAddress,
		MNTAddress:                  MNTAddress,
		privateKey:                  priKey,
		l1ChainID:                   chainIds.L1ChainId,
		zkFairChainId:               chainIds.ZkfairChainId,
		x1ChainId:                   chainIds.X1ChainId,
		mantleChainId:               chainIds.MantleChainId,
		opChainId:                   chainIds.OpChainId,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in selaginella processor: %w", err))
		}},
	}, nil
}

func (s *RpcServer) Start(ctx context.Context) error {
	go func(s *RpcServer) {
		addr := fmt.Sprintf("%s:%s", s.GrpcHostname, s.GrpcPort)
		log.Info("start rpc server ", "addr", addr)

		listener, err := net.Listen("tcp", addr)
		if err != nil {
			log.Error("Could not start tcp listener")
		}

		opt := grpc.MaxRecvMsgSize(MaxRecvMessageSize)

		gs := grpc.NewServer(
			opt,
			grpc.ChainUnaryInterceptor(),
		)
		reflection.Register(gs)
		pb.RegisterBridgeServiceServer(gs, s)

		log.Info("grpc info ", "port", s.GrpcPort, "address", listener.Addr().String())
		if err := gs.Serve(listener); err != nil {
			log.Error("Could not GRPC server")
		}
	}(s)

	CTSTicker := time.NewTicker(10 * time.Second)
	s.tasks.Go(func() error {
		for range CTSTicker.C {
			err := s.ChangeBridgeTransactionStatus()
			if err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})

	CDUBTicker := time.NewTicker(10 * time.Second)
	s.tasks.Go(func() error {
		for range CDUBTicker.C {
			err := s.ChangeDepositUpdateBalanceTransactionStatus()
			if err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})

	CWUBTicker := time.NewTicker(10 * time.Second)
	s.tasks.Go(func() error {
		for range CWUBTicker.C {
			err := s.ChangeWithdrawUpdateBalanceTransactionStatus()
			if err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})

	CUBTicker := time.NewTicker(10 * time.Second)
	s.tasks.Go(func() error {
		for range CUBTicker.C {
			err := s.ChangeUnstakeBatchTransactionStatus()
			if err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})

	CUSTicker := time.NewTicker(10 * time.Second)
	s.tasks.Go(func() error {
		for range CUSTicker.C {
			err := s.ChangeUnstakeSingleTransactionStatus()
			if err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})

	CBMTicker := time.NewTicker(10 * time.Second)
	s.tasks.Go(func() error {
		for range CBMTicker.C {
			err := s.ChangeBatchMintTransactionStatus()
			if err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})

	SBTTicker := time.NewTicker(5 * time.Second)
	s.tasks.Go(func() error {
		for range SBTTicker.C {
			err := s.SendBridgeTransaction()
			if err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})

	SDUBTicker := time.NewTicker(3 * time.Second)
	s.tasks.Go(func() error {
		for range SDUBTicker.C {
			err := s.SendDepositUpdateBalanceTransaction()
			if err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})

	SWUBTicker := time.NewTicker(3 * time.Second)
	s.tasks.Go(func() error {
		for range SWUBTicker.C {
			err := s.SendWithdrawUpdateBalanceTransaction()
			if err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})

	SUBTicker := time.NewTicker(3 * time.Second)
	s.tasks.Go(func() error {
		for range SUBTicker.C {
			err := s.SendUnstakeBatchTransaction()
			if err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})

	SUSTicker := time.NewTicker(3 * time.Second)
	s.tasks.Go(func() error {
		for range SUSTicker.C {
			err := s.SendUnstakeSingleTransaction()
			if err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})

	SBMTicker := time.NewTicker(3 * time.Second)
	s.tasks.Go(func() error {
		for range SBMTicker.C {
			err := s.SendBatchMintTransaction()
			if err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})

	CPTicker := time.NewTicker(1 * time.Minute)
	s.tasks.Go(func() error {
		for range CPTicker.C {
			err := s.CompletePoolAndNew()
			if err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})

	//XSBTicker := time.NewTicker(3 * time.Second)
	//s.tasks.Go(func() error {
	//	for range XSBTicker.C {
	//		err := s.metricX1StrategyBalance()
	//		if err != nil {
	//			log.Error(err.Error())
	//		}
	//	}
	//	return nil
	//})

	OSBTicker := time.NewTicker(10 * time.Second)
	s.tasks.Go(func() error {
		for range OSBTicker.C {
			err := s.metricOpStrategyBalance()
			if err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})

	return nil
}

func (s *RpcServer) CrossChainTransfer(ctx context.Context, in *pb.CrossChainTransferRequest) (*pb.CrossChainTransferResponse, error) {
	if in == nil {
		log.Warn("invalid request: request body is empty")
		return nil, errors.New("invalid request: request body is empty")
	}

	var crossChainTransfers []database.CrossChainTransfer

	cCF, _ := s.db.CrossChainTransfer.CrossChainTransferBySourceHash(in.SourceHash)
	if cCF != nil {
		log.Error("cannot be called repeatedly!")
		return nil, errors.New("cannot be called repeatedly")
	}

	sourceHash := common.HexToHash(in.SourceHash)
	crossChainTransfer := s.db.CrossChainTransfer.BuildCrossChainTransfer(in, sourceHash)
	crossChainTransfers = append(crossChainTransfers, crossChainTransfer)

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
		if err := s.db.Transaction(func(tx *database.DB) error {
			if len(crossChainTransfers) > 0 {
				if err := s.db.CrossChainTransfer.StoreBatchCrossChainTransfer(crossChainTransfers); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			log.Error("unable to persist batch", "err", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return nil, err
	}

	return &pb.CrossChainTransferResponse{
		Success: true,
		Message: "call cross chain transfer success",
	}, nil
}

func (s *RpcServer) ChangeTransferStatus(ctx context.Context, in *pb.CrossChainTransferStatusRequest) (*pb.CrossChainTransferStatusResponse, error) {
	if in == nil {
		log.Warn("invalid request: request body is empty")
		return nil, errors.New("invalid request: request body is empty")
	}

	err := s.db.CrossChainTransfer.ChangeCrossChainTransferSuccessStatueByTxHash(in.TxHash)
	if err != nil {
		log.Error("change cross chain transfer status fail", "err", err)
		return nil, err
	}

	return &pb.CrossChainTransferStatusResponse{
		Success: true,
		Message: "call cross chain transfer success",
	}, nil
}

func (s *RpcServer) UpdateWithdrawFundingPoolBalance(ctx context.Context, in *pb.UpdateWithdrawFundingPoolBalanceRequest) (*pb.UpdateWithdrawFundingPoolBalanceResponse, error) {
	if in == nil {
		log.Warn("invalid request: request body is empty")
		return nil, errors.New("invalid request: request body is empty")
	}

	var updateFundingPoolBalances []database.UpdateWithdrawFundingPoolBalance

	uFPB, _ := s.db.UpdateWithdrawFundingPoolBalance.UpdateFundingPoolBalanceBySourceHash(in.SourceHash)
	if uFPB != nil {
		log.Error("cannot be called repeatedly!")
		return nil, errors.New("cannot be called repeatedly")
	}

	sourceHash := common.HexToHash(in.SourceHash)
	updateFundingPoolBalance := s.db.UpdateWithdrawFundingPoolBalance.BuildUpdateFundingPoolBalance(in, sourceHash)
	updateFundingPoolBalances = append(updateFundingPoolBalances, updateFundingPoolBalance)

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
		if err := s.db.Transaction(func(tx *database.DB) error {
			if len(updateFundingPoolBalances) > 0 {
				if err := s.db.UpdateWithdrawFundingPoolBalance.StoreBatchUpdateFundingPoolBalance(updateFundingPoolBalances); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			log.Error("unable to persist batch", "err", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return nil, err
	}

	return &pb.UpdateWithdrawFundingPoolBalanceResponse{
		Success: true,
		Message: "call update withdraw funding pool balance success",
	}, nil
}

func (s *RpcServer) UpdateDepositFundingPoolBalance(ctx context.Context, in *pb.UpdateDepositFundingPoolBalanceRequest) (*pb.UpdateDepositFundingPoolBalanceResponse, error) {
	if in == nil {
		log.Warn("invalid request: request body is empty")
		return nil, errors.New("invalid request: request body is empty")
	}

	var updateFundingPoolBalances []database.UpdateDepositFundingPoolBalance

	uFPB, _ := s.db.UpdateDepositFundingPoolBalance.UpdateFundingPoolBalanceBySourceHash(in.SourceHash)
	if uFPB != nil {
		log.Error("cannot be called repeatedly!")
		return nil, errors.New("cannot be called repeatedly")
	}

	sourceHash := common.HexToHash(in.SourceHash)
	updateFundingPoolBalance := s.db.UpdateDepositFundingPoolBalance.BuildUpdateFundingPoolBalance(in, sourceHash)
	updateFundingPoolBalances = append(updateFundingPoolBalances, updateFundingPoolBalance)

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
		if err := s.db.Transaction(func(tx *database.DB) error {
			if len(updateFundingPoolBalances) > 0 {
				if err := s.db.UpdateDepositFundingPoolBalance.StoreBatchUpdateFundingPoolBalance(updateFundingPoolBalances); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			log.Error("unable to persist batch", "err", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return nil, err
	}

	return &pb.UpdateDepositFundingPoolBalanceResponse{
		Success: true,
		Message: "call update deposit funding pool balance success",
	}, nil
}

func (s *RpcServer) UnstakeBatch(ctx context.Context, in *pb.UnstakeBatchRequest) (*pb.UnstakeBatchResponse, error) {
	if in == nil {
		log.Warn("invalid request: request body is empty")
		return nil, errors.New("invalid request: request body is empty")
	}

	var unstakeBatchs []database.UnstakeBatch

	uBSH, _ := s.db.UnstakeBatch.UnstakeBatchBySourceHash(in.SourceHash)
	if uBSH != nil {
		log.Error("cannot be called repeatedly!")
		return nil, errors.New("cannot be called repeatedly")
	}

	sourceHash := common.HexToHash(in.SourceHash)
	unstakeBatch := s.db.UnstakeBatch.BuildUnstakeBatch(in, sourceHash)
	unstakeBatchs = append(unstakeBatchs, unstakeBatch)

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
		if err := s.db.Transaction(func(tx *database.DB) error {
			if len(unstakeBatchs) > 0 {
				if err := s.db.UnstakeBatch.StoreUnstakeBatch(unstakeBatchs); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			log.Error("unable to persist batch", "err", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return nil, err
	}

	return &pb.UnstakeBatchResponse{
		Success: true,
		Message: "call unstake batch success",
	}, nil
}

func (s *RpcServer) UnstakeSingle(ctx context.Context, in *pb.UnstakeSingleRequest) (*pb.UnstakeSingleResponse, error) {
	if in == nil {
		log.Warn("invalid request: request body is empty")
		return nil, errors.New("invalid request: request body is empty")
	}

	var unstakeSingles []database.UnstakeSingle

	uSSH, _ := s.db.UnstakeSingle.UnstakeSingleBySourceHash(in.SourceHash)
	if uSSH != nil {
		log.Error("cannot be called repeatedly!")
		return nil, errors.New("cannot be called repeatedly")
	}

	sourceHash := common.HexToHash(in.SourceHash)
	unstakeSingle := s.db.UnstakeSingle.BuildUnstakeSingle(in, sourceHash)
	unstakeSingles = append(unstakeSingles, unstakeSingle)

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
		if err := s.db.Transaction(func(tx *database.DB) error {
			if len(unstakeSingles) > 0 {
				if err := s.db.UnstakeSingle.StoreUnstakeSingle(unstakeSingles); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			log.Error("unable to persist batch", "err", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return nil, err
	}

	return &pb.UnstakeSingleResponse{
		Success: true,
		Message: "call unstake single success",
	}, nil
}

func (s *RpcServer) BatchMint(ctx context.Context, in *pb.BatchMintRequest) (*pb.BatchMintResponse, error) {
	if in == nil {
		log.Warn("invalid request: request body is empty")
		return nil, errors.New("invalid request: request body is empty")
	}

	var batchMints []database.BatchMint

	bMBB, _ := s.db.BatchMint.BatchMintByBatch(in.Batch)
	if bMBB != nil {
		log.Error("cannot be called repeatedly!")
		return nil, errors.New("cannot be called repeatedly")
	}

	batchMints = s.db.BatchMint.BuildBatchMint(in)

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
		if err := s.db.Transaction(func(tx *database.DB) error {
			if len(batchMints) > 0 {
				if err := s.db.BatchMint.StoreBatchMint(batchMints); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			log.Error("unable to persist batch", "err", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return nil, err
	}

	return &pb.BatchMintResponse{
		Success: true,
		Message: "call batch mint success",
	}, nil
}

func (s *RpcServer) SendBridgeTransaction() error {

	bridgeTxs, _ := s.db.CrossChainTransfer.OldestPendingNoSentTransaction()
	if len(bridgeTxs) == 0 {
		log.Warn("no more bridge transaction need to be sent")
		return nil
	}

	var ctx context.Context
	for _, bridgeTx := range bridgeTxs {
		if bridgeTx.SourceChainId.Uint64() == 1442 || bridgeTx.DestChainId.Uint64() == 1442 {
			return nil
		}
		bridge, err := s.bridgeLogic(ctx, &bridgeTx)
		if err != nil {
			return err
		}

		retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
		if _, err := retry.Do[interface{}](context.Background(), 10, retryStrategy, func() (interface{}, error) {
			if err := s.db.Transaction(func(tx *database.DB) error {
				if bridge != nil {
					if err := s.db.CrossChainTransfer.UpdateCrossChainTransferTransactionHash(*bridge); err != nil {
						return err
					}
				}
				return nil
			}); err != nil {
				log.Error("unable to persist batch", "err", err)
				return nil, fmt.Errorf("unable to persist batch: %w", err)
			}
			return nil, nil
		}); err != nil {
			return err
		}

		log.Info("send bridge transaction success", "tx_hash", bridge.TxHash)

	}

	return nil
}

func (s *RpcServer) bridgeLogic(ctx context.Context, bridgeTx *database.CrossChainTransfer) (*database.CrossChainTransfer, error) {

	var opts *bind.TransactOpts
	var err error
	var tx *types.Transaction
	var finalTx *types.Transaction

	for chainId, client := range s.ethClients {
		if chainId == bridgeTx.DestChainId.Uint64() {
			opts, err = s.newTransactOpts(ctx, chainId)
			if err != nil {
				return nil, err
			}

			log.Info(fmt.Sprintf("get send transaction request: sourceChainId=%v, destChainId=%v, amount=%v, fee=%v, nonce=%v, tokenAddress=%v", bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce, bridgeTx.TokenAddress))

			if bridgeTx.DestReceiveAddress.String() == s.l1StakingManagerAddr.String() && bridgeTx.DestChainId.Uint64() == s.l1ChainID {
				tx, err = s.L1BridgeContract.BridgeFinalizeETHForStaking(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce, s.l1StakingManagerAddr)
				if err != nil {
					log.Error("transfer eth to l1 staking manager by abi fail", "error", err)
					return nil, err
				}

				finalTx, err = s.RawL1BridgeContract.RawTransact(opts, tx.Data())
				if err != nil {
					log.Error("raw send bridge transaction fail", "error", err)
					return nil, err
				}

				err = client.SendTransaction(ctx, finalTx)
				if err != nil {
					log.Error("send bridge transaction fail", "error", err)
					return nil, err
				}
				bridgeTx.TxHash = finalTx.Hash()

				retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
				if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
					if err := s.db.Transaction(func(tx *database.DB) error {
						if bridgeTx != nil {
							if err := s.db.CrossChainTransfer.UpdateCrossChainTransferTransactionHash(*bridgeTx); err != nil {
								return err
							}
						}
						return nil
					}); err != nil {
						log.Error("unable to persist batch", "err", err)
						return nil, fmt.Errorf("unable to persist batch: %w", err)
					}
					return nil, nil
				}); err != nil {
					return nil, err
				}

				return nil, nil
			}

			switch bridgeTx.TokenAddress.String() {
			case s.EthAddress[chainId].String():
				if chainId == s.l1ChainID {
					if bridgeTx.SourceChainId.Uint64() == s.zkFairChainId {
						tx, err = s.L1BridgeContract.BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.USDCAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("zkfair transfer usdc to l1 by abi fail", "error", err)
							return nil, err
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.x1ChainId {
						tx, err = s.L1BridgeContract.BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.OKBAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("x1 transfer okb to l1 by abi fail", "error", err)
							return nil, err
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.mantleChainId {
						tx, err = s.L1BridgeContract.BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.MNTAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("mantle transfer mnt to l1 by abi fail", "error", err)
							return nil, err
						}
					} else {
						tx, err = s.L1BridgeContract.BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("transfer eth to l1 by abi fail", "error", err)
							return nil, err
						}
					}
				} else if chainId == s.x1ChainId {
					if bridgeTx.SourceChainId.Uint64() == s.zkFairChainId {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.USDCAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("zkfair transfer usdc to x1 by abi fail", "error", err)
							return nil, err
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.mantleChainId {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.MNTAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("mantle transfer mnt to x1 by abi fail", "error", err)
							return nil, err
						}
					} else {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("transfer eth to x1 by abi fail", "error", err)
							return nil, err
						}
					}
				} else if chainId == s.zkFairChainId {
					if bridgeTx.SourceChainId.Uint64() == s.x1ChainId {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.OKBAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("x1 transfer okb to zkfair by abi fail", "error", err)
							return nil, err
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.mantleChainId {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.MNTAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("mantle transfer mnt to zkfair by abi fail", "error", err)
							return nil, err
						}
					} else {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("transfer eth to zkfair by abi fail", "error", err)
							return nil, err
						}
					}
				} else if chainId == s.mantleChainId {
					if bridgeTx.SourceChainId.Uint64() == s.x1ChainId {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.OKBAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("x1 transfer okb to mantle by abi fail", "error", err)
							return nil, err
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.zkFairChainId {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.USDCAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("mantle transfer mnt to mantle by abi fail", "error", err)
							return nil, err
						}
					} else {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("transfer eth to mantle by abi fail", "error", err)
							return nil, err
						}
					}
				} else {
					tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
					if err != nil {
						log.Error("transfer eth to l2 by abi fail", "error", err)
						return nil, err
					}
				}

			case s.WEthAddress[chainId].String():
				if chainId == s.l1ChainID {
					tx, err = s.L1BridgeContract.BridgeFinalizeWETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
					if err != nil {
						log.Error("get bridge finalize l1 weth by abi fail", "error", err)
						return nil, err
					}
				} else {
					tx, err = s.L2BridgeContract[chainId].BridgeFinalizeWETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
					if err != nil {
						log.Error("get bridge finalize l2 weth by abi fail", "error", err)
						return nil, err
					}
				}

			default:
				if chainId == s.l1ChainID {
					if bridgeTx.SourceChainId.Uint64() == s.x1ChainId {
						if bridgeTx.TokenAddress.String() == s.WEthAddress[s.x1ChainId].String() {
							tx, err = s.L1BridgeContract.BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("x1 transfer eth to l1 by abi fail", "error", err)
								return nil, err
							}
						} else {
							tx, err = s.L1BridgeContract.BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("x1 transfer erc20 to l1 by abi fail", "error", err)
								return nil, err
							}
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.zkFairChainId {
						if bridgeTx.TokenAddress.String() == s.WEthAddress[s.zkFairChainId].String() {
							tx, err = s.L1BridgeContract.BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("zkfair transfer eth to l1 by abi fail", "error", err)
								return nil, err
							}
						} else {
							tx, err = s.L1BridgeContract.BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("zkfair transfer erc20 to l1 by abi fail", "error", err)
								return nil, err
							}
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.mantleChainId {
						if bridgeTx.TokenAddress.String() == s.WEthAddress[s.mantleChainId].String() {
							tx, err = s.L1BridgeContract.BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer eth to l1 by abi fail", "error", err)
								return nil, err
							}
						} else {
							tx, err = s.L1BridgeContract.BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer erc20 to l1 by abi fail", "error", err)
								return nil, err
							}
						}
					} else {
						tx, err = s.L1BridgeContract.BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("transfer erc20 to l1 by abi fail", "error", err)
							return nil, err
						}
					}
				} else if chainId == s.zkFairChainId {
					if bridgeTx.SourceChainId.Uint64() == s.l1ChainID {
						if bridgeTx.TokenAddress.String() == s.EthAddress[s.l1ChainID].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("l1 transfer eth to zkfair by abi fail", "error", err)
								return nil, err
							}
						} else if bridgeTx.TokenAddress.String() == s.USDCAddress[s.l1ChainID].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("l1 transfer usdc to zkfair by abi fail", "error", err)
								return nil, err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("l1 transfer erc20 to zkfair by abi fail", "error", err)
								return nil, err
							}
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.x1ChainId {
						if bridgeTx.TokenAddress.String() == s.WEthAddress[s.x1ChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("x1 transfer eth to zkfair by abi fail", "error", err)
								return nil, err
							}
						} else if bridgeTx.TokenAddress.String() == s.USDCAddress[s.x1ChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("x1 transfer usdc to zkfair by abi fail", "error", err)
								return nil, err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("x1 transfer erc20 to zkfair by abi fail", "error", err)
								return nil, err
							}
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.mantleChainId {
						if bridgeTx.TokenAddress.String() == s.WEthAddress[s.mantleChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer eth to zkfair by abi fail", "error", err)
								return nil, err
							}
						} else if bridgeTx.TokenAddress.String() == s.USDCAddress[s.mantleChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer usdc to zkfair by abi fail", "error", err)
								return nil, err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer erc20 to zkfair by abi fail", "error", err)
								return nil, err
							}
						}
					} else {
						if bridgeTx.TokenAddress.String() == s.USDCAddress[bridgeTx.SourceChainId.Uint64()].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer usdc to zkfair by abi fail", "error", err)
								return nil, err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("transfer erc20 to zkfair by abi fail", "error", err)
								return nil, err
							}
						}
					}
				} else if chainId == s.x1ChainId {
					if bridgeTx.SourceChainId.Uint64() == s.l1ChainID {
						if bridgeTx.TokenAddress.String() == s.EthAddress[s.l1ChainID].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("l1 transfer eth to x1 by abi fail", "error", err)
								return nil, err
							}
						} else if bridgeTx.TokenAddress.String() == s.OKBAddress[s.l1ChainID].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("l1 transfer okb to x1 by abi fail", "error", err)
								return nil, err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("l1 transfer erc20 to x1 by abi fail", "error", err)
								return nil, err
							}
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.zkFairChainId {
						if bridgeTx.TokenAddress.String() == s.WEthAddress[s.zkFairChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("zkfair transfer eth to x1 by abi fail", "error", err)
								return nil, err
							}
						} else if bridgeTx.TokenAddress.String() == s.OKBAddress[s.zkFairChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("zkfair transfer okb to x1 by abi fail", "error", err)
								return nil, err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("zkfair transfer erc20 to x1 by abi fail", "error", err)
								return nil, err
							}
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.mantleChainId {
						if bridgeTx.TokenAddress.String() == s.WEthAddress[s.mantleChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer eth to x1 by abi fail", "error", err)
								return nil, err
							}
						} else if bridgeTx.TokenAddress.String() == s.OKBAddress[s.mantleChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer okb to x1 by abi fail", "error", err)
								return nil, err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer erc20 to x1 by abi fail", "error", err)
								return nil, err
							}
						}
					} else {
						if bridgeTx.TokenAddress.String() == s.OKBAddress[bridgeTx.SourceChainId.Uint64()].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("transfer okb to x1 by abi fail", "error", err)
								return nil, err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("transfer erc20 to x1 by abi fail", "error", err)
								return nil, err
							}
						}
					}
				} else {
					if bridgeTx.SourceChainId.Uint64() == s.x1ChainId && bridgeTx.TokenAddress == s.WEthAddress[s.x1ChainId] {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("x1 transfer eth to l2 by abi fail", "error", err)
							return nil, err
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.zkFairChainId && bridgeTx.TokenAddress == s.WEthAddress[s.zkFairChainId] {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("zkfair transfer eth to l2 by abi fail", "error", err)
							return nil, err
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.mantleChainId && bridgeTx.TokenAddress == s.WEthAddress[s.mantleChainId] {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("mantle transfer eth to l2 by abi fail", "error", err)
							return nil, err
						}
					} else {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("transfer erc20 to l2 by abi fail", "error", err)
							return nil, err
						}
					}
				}
			}

			log.Info("get bridge finalize by abi success")

			if chainId == s.l1ChainID {
				finalTx, err = s.RawL1BridgeContract.RawTransact(opts, tx.Data())
				if err != nil {
					log.Error("raw send bridge transaction fail", "error", err)
					return nil, err
				}
			} else {
				finalTx, err = s.RawL2BridgeContract[chainId].RawTransact(opts, tx.Data())
				if err != nil {
					log.Error("raw send bridge transaction fail", "error", err)
					return nil, err
				}
			}

			err = client.SendTransaction(ctx, finalTx)
			if err != nil {
				log.Error("send bridge transaction fail", "error", err)
				return nil, err
			}
			bridgeTx.TxHash = finalTx.Hash()
		}
	}
	return bridgeTx, nil
}

func (s *RpcServer) SendDepositUpdateBalanceTransaction() error {

	var cOpts *bind.CallOpts
	var tOpts *bind.TransactOpts
	var err error
	var tx *types.Transaction
	var finalTx *types.Transaction
	var ctx = context.Background()
	var nilUpdateFundingPoolBalance database.UpdateDepositFundingPoolBalance

	updateTx, _ := s.db.UpdateDepositFundingPoolBalance.OldestPendingNoSentTransaction()
	if *updateTx == nilUpdateFundingPoolBalance {
		log.Warn("no more update deposit balance transaction need to be sent")
		return nil
	}

	latestBlock, err := s.ethClients[updateTx.DestChainId.Uint64()].GetLatestBlock()
	if err != nil {
		log.Error("get latest block number fail", "err", err)
		return err
	}

	cOpts = &bind.CallOpts{
		BlockNumber: latestBlock,
		From:        crypto.PubkeyToAddress(s.privateKey.PublicKey),
	}

	balance, err := s.L2BridgeContract[updateTx.DestChainId.Uint64()].FundingPoolBalance(cOpts, updateTx.TokenAddress)
	if err != nil {
		log.Error("get l2 bridge funding pool balance fail", "err", err)
		return err
	}

	tOpts, err = s.newTransactOpts(ctx, updateTx.DestChainId.Uint64())
	if err != nil {
		log.Error("get transactOpts fail", "err", err)
		return err
	}

	tx, err = s.L2BridgeContract[updateTx.DestChainId.Uint64()].UpdateFundingPoolBalance(tOpts, updateTx.TokenAddress, new(big.Int).Add(balance, updateTx.Amount))
	finalTx, err = s.RawL2BridgeContract[updateTx.DestChainId.Uint64()].RawTransact(tOpts, tx.Data())
	if err != nil {
		log.Error("raw send update deposit funding pool balance transaction fail", "error", err)
		return err
	}
	err = s.ethClients[updateTx.DestChainId.Uint64()].SendTransaction(ctx, finalTx)
	if err != nil {
		log.Error("send update deposit funding pool balance transaction fail", "error", err)
		return err
	}

	log.Info("send update deposit funding pool balance transaction success", "tx_hash", finalTx.Hash())

	updateTx.TxHash = finalTx.Hash()

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
		if err := s.db.Transaction(func(tx *database.DB) error {
			if updateTx != nil {
				if err := s.db.UpdateDepositFundingPoolBalance.UpdateFundingPoolBalanceTransactionHash(*updateTx); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			log.Error("unable to persist batch", "err", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *RpcServer) SendWithdrawUpdateBalanceTransaction() error {

	var cOpts *bind.CallOpts
	var tOpts *bind.TransactOpts
	var err error
	var tx *types.Transaction
	var finalTx *types.Transaction
	var ctx = context.Background()
	var nilUpdateFundingPoolBalance database.UpdateWithdrawFundingPoolBalance

	updateTx, _ := s.db.UpdateWithdrawFundingPoolBalance.OldestPendingNoSentTransaction()
	if *updateTx == nilUpdateFundingPoolBalance {
		log.Warn("no more update withdraw balance transaction need to be sent")
		return nil
	}

	latestBlock, err := s.ethClients[updateTx.DestChainId.Uint64()].GetLatestBlock()
	if err != nil {
		log.Error("get latest block number fail", "err", err)
		return err
	}

	cOpts = &bind.CallOpts{
		BlockNumber: latestBlock,
		From:        crypto.PubkeyToAddress(s.privateKey.PublicKey),
	}

	balance, err := s.L1BridgeContract.FundingPoolBalance(cOpts, updateTx.TokenAddress)
	if err != nil {
		log.Error("get l1 bridge funding pool balance fail", "err", err)
		return err
	}

	tOpts, err = s.newTransactOpts(ctx, updateTx.DestChainId.Uint64())
	if err != nil {
		log.Error("get transactOpts fail", "err", err)
		return err
	}

	tx, err = s.L2BridgeContract[updateTx.DestChainId.Uint64()].UpdateFundingPoolBalance(tOpts, updateTx.TokenAddress, new(big.Int).Add(balance, updateTx.Amount))
	finalTx, err = s.RawL1BridgeContract.RawTransact(tOpts, tx.Data())
	if err != nil {
		log.Error("raw send update withdraw funding pool balance transaction fail", "error", err)
		return err
	}
	err = s.ethClients[updateTx.DestChainId.Uint64()].SendTransaction(ctx, finalTx)
	if err != nil {
		log.Error("send update withdraw funding pool balance transaction fail", "error", err)
		return err
	}

	log.Info("send update withdraw funding pool balance transaction success", "tx_hash", finalTx.Hash())

	updateTx.TxHash = finalTx.Hash()

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
		if err := s.db.Transaction(func(tx *database.DB) error {
			if updateTx != nil {
				if err := s.db.UpdateWithdrawFundingPoolBalance.UpdateFundingPoolBalanceTransactionHash(*updateTx); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			log.Error("unable to persist batch", "err", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *RpcServer) SendUnstakeBatchTransaction() error {

	var tOpts *bind.TransactOpts
	var err error
	var tx *types.Transaction
	var finalTx *types.Transaction
	var ctx = context.Background()
	var nilUnstakeBatch database.UnstakeBatch

	unStakeTx, _ := s.db.UnstakeBatch.OldestPendingNoSentTransaction()
	if *unStakeTx == nilUnstakeBatch {
		log.Warn("no more unstake batch tx need to be sent")
		return nil
	}

	tOpts, err = s.newTransactOpts(ctx, s.l1ChainID)
	if err != nil {
		log.Error("get transactOpts fail", "err", err)
		return err
	}

	tx, err = s.l1StakingManagerContract.ClaimUnstakeRequest(tOpts, unStakeTx.StrategyAddress, unStakeTx.BridgeAddress, unStakeTx.SourceChainId, unStakeTx.DestChainId, unStakeTx.GasLimit)

	finalTx, err = s.RawL1StakingManagerContract.RawTransact(tOpts, tx.Data())
	if err != nil {
		log.Error("raw send l1 staking manager claim unstake request transaction fail", "error", err)
		return err
	}
	err = s.ethClients[s.l1ChainID].SendTransaction(ctx, finalTx)
	if err != nil {
		log.Error("send l1 staking manager claim unstake request transaction fail", "error", err)
		return err
	}

	log.Info("send l1 staking manager claim unstake request transaction success", "tx_hash", finalTx.Hash())

	unStakeTx.TxHash = finalTx.Hash()

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
		if err := s.db.Transaction(func(tx *database.DB) error {
			if unStakeTx != nil {
				if err := s.db.UnstakeBatch.UpdateUnstakeBatchTransactionHash(*unStakeTx); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			log.Error("unable to persist batch", "err", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *RpcServer) SendUnstakeSingleTransaction() error {

	var tOpts *bind.TransactOpts
	var err error
	var tx *types.Transaction
	var finalTx *types.Transaction
	var ctx = context.Background()
	var nilUnstakeSingle database.UnstakeSingle

	unStakeSTx, _ := s.db.UnstakeSingle.OldestPendingNoSentTransaction()
	if *unStakeSTx == nilUnstakeSingle {
		log.Warn("no more unstake single tx need to be sent")
		return nil
	}

	tOpts, err = s.newTransactOpts(ctx, s.l1ChainID)
	if err != nil {
		log.Error("get transactOpts fail", "err", err)
		return err
	}

	tx, err = s.StrategyManagerContract[unStakeSTx.ChainId.Uint64()].MigrateRelatedL1StakerShares(tOpts, unStakeSTx.StakerAddress, unStakeSTx.StrategyAddress, unStakeSTx.SharesAmount)

	finalTx, err = s.RawStrategyManagerContract[unStakeSTx.ChainId.Uint64()].RawTransact(tOpts, tx.Data())
	if err != nil {
		log.Error("raw send strategy manager claim unstake single request transaction fail", "error", err)
		return err
	}
	err = s.ethClients[unStakeSTx.ChainId.Uint64()].SendTransaction(ctx, finalTx)
	if err != nil {
		log.Error("send strategy manager claim unstake single request transaction fail", "error", err)
		return err
	}

	log.Info("send strategy manager claim unstake single request transaction success", "tx_hash", finalTx.Hash())

	unStakeSTx.TxHash = finalTx.Hash()

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
		if err := s.db.Transaction(func(tx *database.DB) error {
			if unStakeSTx != nil {
				if err := s.db.UnstakeSingle.UpdateUnstakeSingleTransactionHash(*unStakeSTx); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			log.Error("unable to persist batch", "err", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *RpcServer) SendBatchMintTransaction() error {

	var tOpts *bind.TransactOpts
	var err error
	var tx *types.Transaction
	var finalTx *types.Transaction
	var ctx = context.Background()
	var nilBatchMint database.BatchMint
	var Batcher []bindings.IDETHBatchMint

	batchMintTx, _ := s.db.BatchMint.OldestPendingNoSentTransaction()
	if *batchMintTx == nilBatchMint {
		log.Warn("no more batch mint tx need to be sent")
		return nil
	}

	tOpts, err = s.newTransactOpts(ctx, s.l1ChainID)
	if err != nil {
		log.Error("get transactOpts failed", "err", err)
		return err
	}

	batchMints, err := s.db.BatchMint.BatchMintsByBatch(batchMintTx.Batch.Uint64())
	if err != nil {
		log.Error("get batch mints failed", "err", err)
		return err
	}
	for _, v := range batchMints {
		mint := bindings.IDETHBatchMint{
			Staker: v.StakerAddress,
			Amount: v.SharesAmount,
		}
		Batcher = append(Batcher, mint)
	}

	tx, err = s.l1DETHContract.BatchMint(tOpts, Batcher)

	finalTx, err = s.RawL1DETHContract.RawTransact(tOpts, tx.Data())
	if err != nil {
		log.Error("raw send batch mint transaction fail", "error", err)
		return err
	}
	err = s.ethClients[s.l1ChainID].SendTransaction(ctx, finalTx)
	if err != nil {
		log.Error("send batch mint transaction fail", "error", err)
		return err
	}

	log.Info("send batch mint transaction success", "tx_hash", finalTx.Hash())

	for _, v := range batchMints {
		v.TxHash = finalTx.Hash()
	}

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
		if err := s.db.Transaction(func(tx *database.DB) error {
			if batchMints != nil {
				if err := s.db.BatchMint.UpdateBatchMintTransactionHash(batchMints); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			log.Error("unable to persist batch", "err", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *RpcServer) ChangeBridgeTransactionStatus() error {
	var receipt *types.Receipt

	txs, err := s.db.CrossChainTransfer.OldestPendingSentTransaction()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("get oldest pending transaction fail", "err", err)
		return err
	}

	if len(txs) == 0 {
		log.Info("no more bridge pending transaction !")
	}

	var wg sync.WaitGroup
	for _, tx := range txs {
		wg.Add(1)
		go func(tx database.CrossChainTransfer) {
			defer wg.Done()
			if tx.DestChainId != nil {
				receipt, err = getTransactionReceipt(s.ethClients[tx.DestChainId.Uint64()], tx.TxHash)
				if err != nil {
					return
				}
				err = s.db.CrossChainTransfer.ChangeCrossChainTransferSentStatusByTxHash(receipt.TxHash.String())
				if err != nil {
					log.Error("change transaction status fail", "err", err)
					return
				}
			}

			log.Info("change bridge transaction status success", "txHash", receipt.TxHash)
		}(tx)
	}

	wg.Wait()

	return nil
}

func (s *RpcServer) ChangeDepositUpdateBalanceTransactionStatus() error {
	var receipt *types.Receipt

	tx, err := s.db.UpdateDepositFundingPoolBalance.OldestPendingSentTransaction()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("get oldest pending transaction fail", "err", err)
		return err
	}

	if tx.DestChainId != nil {
		receipt, err = s.ethClients[tx.DestChainId.Uint64()].TxReceiptDetailByHash(tx.TxHash)
		if errors.Is(err, ethereum.NotFound) {
			log.Warn("transaction not found")
			return nil
		}
	} else {
		log.Info("no more deposit update pool balance pending transaction !")
		return nil
	}

	err = s.db.UpdateDepositFundingPoolBalance.ChangeUpdateFundingPoolBalanceSentStatusByTxHash(receipt.TxHash.String())
	if err != nil {
		log.Error("change deposit update pool balance transaction status fail", "err", err)
		return err
	}

	log.Info("change update deposit pool balance transaction status success", "txHash", receipt.TxHash)

	return nil
}

func (s *RpcServer) ChangeWithdrawUpdateBalanceTransactionStatus() error {
	var receipt *types.Receipt

	tx, err := s.db.UpdateWithdrawFundingPoolBalance.OldestPendingSentTransaction()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("get oldest pending transaction fail", "err", err)
		return err
	}

	if tx.DestChainId != nil {
		receipt, err = s.ethClients[tx.DestChainId.Uint64()].TxReceiptDetailByHash(tx.TxHash)
		if errors.Is(err, ethereum.NotFound) {
			log.Warn("transaction not found")
			return nil
		}
	} else {
		log.Info("no more withdraw update pool balance pending transaction !")
		return nil
	}

	err = s.db.UpdateWithdrawFundingPoolBalance.ChangeUpdateFundingPoolBalanceSentStatusByTxHash(receipt.TxHash.String())
	if err != nil {
		log.Error("change withdraw update pool balance transaction status fail", "err", err)
		return err
	}

	log.Info("change update withdraw pool balance transaction status success", "txHash", receipt.TxHash)

	return nil
}

func (s *RpcServer) ChangeUnstakeBatchTransactionStatus() error {
	var receipt *types.Receipt

	tx, err := s.db.UnstakeBatch.OldestPendingSentTransaction()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("get oldest pending transaction fail", "err", err)
		return err
	}

	if tx.DestChainId != nil {
		receipt, err = s.ethClients[s.l1ChainID].TxReceiptDetailByHash(tx.TxHash)
		if errors.Is(err, ethereum.NotFound) {
			log.Warn("transaction not found")
			return nil
		}
	} else {
		log.Info("no more unstake batch pending transaction !")
		return nil
	}

	err = s.db.UnstakeBatch.ChangeUnstakeBatchSentStatusByTxHash(receipt.TxHash.String())
	if err != nil {
		log.Error("change l1 staking manager claim unstake request transaction status fail", "err", err)
		return err
	}

	log.Info("change l1 staking manager claim unstake request transaction status success", "txHash", receipt.TxHash)

	return nil
}

func (s *RpcServer) ChangeUnstakeSingleTransactionStatus() error {
	var receipt *types.Receipt

	tx, err := s.db.UnstakeSingle.OldestPendingSentTransaction()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("get oldest pending transaction fail", "err", err)
		return err
	}

	if tx.ChainId != nil {
		receipt, err = s.ethClients[tx.ChainId.Uint64()].TxReceiptDetailByHash(tx.TxHash)
		if errors.Is(err, ethereum.NotFound) {
			log.Warn("transaction not found")
			return nil
		}
	} else {
		log.Info("no more unstake single pending transaction !")
		return nil
	}

	err = s.db.UnstakeSingle.ChangeUnstakeSingleSentStatusByTxHash(receipt.TxHash.String())
	if err != nil {
		log.Error("change strategy manager claim unstake single request transaction status fail", "err", err)
		return err
	}

	log.Info("change strategy manager claim unstake single request transaction status success", "txHash", receipt.TxHash)

	return nil
}

func (s *RpcServer) ChangeBatchMintTransactionStatus() error {
	var receipt *types.Receipt

	tx, err := s.db.BatchMint.OldestPendingSentTransaction()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("get oldest pending transaction fail", "err", err)
		return err
	}

	if tx.Batch != nil {
		receipt, err = s.ethClients[s.l1ChainID].TxReceiptDetailByHash(tx.TxHash)
		if errors.Is(err, ethereum.NotFound) {
			log.Warn("transaction not found")
			return nil
		}
	} else {
		log.Info("no more batch mint pending transaction !")
		return nil
	}

	err = s.db.BatchMint.ChangeBatchMintSentStatusByTxHash(receipt.TxHash.String())
	if err != nil {
		log.Error("change batch mint transaction status fail", "err", err)
		return err
	}

	log.Info("change batch mint transaction status success", "txHash", receipt.TxHash)

	return nil
}

func (s *RpcServer) newTransactOpts(ctx context.Context, chainId uint64) (*bind.TransactOpts, error) {
	var opts *bind.TransactOpts
	var err error

	if s.EnableHsm {
		seqBytes, err := hex.DecodeString(s.HsmCreden)
		if err != nil {
			log.Error("selaginella", "decode hsm creden fail", err.Error())
		}
		apikey := option.WithCredentialsJSON(seqBytes)
		kClient, err := kms.NewKeyManagementClient(context.Background(), apikey)
		if err != nil {
			log.Error("selaginella", "create signer error", err.Error())
		}
		mk := &sign.ManagedKey{
			KeyName:      s.HsmAPIName,
			EthereumAddr: common.HexToAddress(s.HsmAddress),
			Gclient:      kClient,
		}
		opts, err = mk.NewEthereumTransactorWithChainID(context.Background(), new(big.Int).SetUint64(chainId))
		if err != nil {
			log.Error("selaginella", "create signer error", err.Error())
		}
	} else {
		if s.privateKey == nil {
			log.Error("selaginella", "create signer error", err.Error())
			return nil, errors.New("no private key provided")
		}

		opts, err = bind.NewKeyedTransactorWithChainID(s.privateKey, new(big.Int).SetUint64(chainId))
		if err != nil {
			log.Error("new keyed transactor fail", "err", err)
			return nil, err
		}
	}

	opts.Context = ctx
	opts.NoSend = true

	return opts, err
}

func (s *RpcServer) CompletePoolAndNew() error {
	var cOpts *bind.CallOpts
	var tOpts *bind.TransactOpts
	var err error
	var newCPools []bindings.IL1PoolManagerPool
	var newPools []*bindings.IL1PoolManagerPool
	var tx *types.Transaction
	var finalTx *types.Transaction
	var receipt *types.Receipt
	var ctx = context.Background()

	latestBlock, err := s.ethClients[s.l1ChainID].GetLatestBlock()
	if err != nil {
		log.Error("get latest block number fail", "err", err)
		return err
	}

	cOpts = &bind.CallOpts{
		BlockNumber: latestBlock,
		From:        crypto.PubkeyToAddress(s.privateKey.PublicKey),
	}

	ethPoolLength, err := s.L1BridgeContract.GetPoolLength(cOpts, s.EthAddress[s.l1ChainID])

	if err != nil {
		log.Error("get pool length fail", "err", err)
		return err
	}

	var l1EthPool bindings.IL1PoolManagerPool

	if ethPoolLength.Uint64() > 1 {
		l1EthPool, err = s.L1BridgeContract.GetPool(cOpts, s.EthAddress[s.l1ChainID], new(big.Int).Sub(ethPoolLength, new(big.Int).SetUint64(2)))
		l1EthPool.Token = s.EthAddress[s.l1ChainID]
		newPools = append(newPools, &l1EthPool)
	}

	if err != nil {
		log.Error("get pool fail", "err", err)
		return err
	}

	s.poolStartTimestamp = l1EthPool.StartTimestamp
	s.poolEndTimestamp = l1EthPool.EndTimestamp

	if time.Now().Unix() >= int64(s.poolEndTimestamp) {
		tOpts, err = s.newTransactOpts(ctx, s.l1ChainID)

		newCPools, err = s.newPools(newPools)
		if err != nil {
			return err
		}

		tx, err = s.L1BridgeContract.CompletePoolAndNew(tOpts, newCPools)
		if err != nil {
			log.Error("get l1 bridge complete pool and new by abi fail", "error", err)
			return err
		}
		finalTx, err = s.RawL1BridgeContract.RawTransact(tOpts, tx.Data())
		if err != nil {
			log.Error("raw send complete pool and new transaction fail", "error", err)
			return err
		}
		err = s.ethClients[s.l1ChainID].SendTransaction(ctx, finalTx)
		if err != nil {
			log.Error("send complete pool and new transaction fail", "error", err)
			return err
		}
		receipt, err = getTransactionReceipt(s.ethClients[s.l1ChainID], finalTx.Hash())
		if err != nil {
			log.Error("get complete pool and new receipt fail", "error", err)
			return err
		}

		log.Info("send complete pool and new transaction success", "tx_hash", receipt.TxHash)

	}

	return nil
}

func getTransactionReceipt(client node.EthClient, txHash common.Hash) (*types.Receipt, error) {
	var receipt *types.Receipt
	var err error

	ticker := time.NewTicker(30 * time.Second)
	for {
		<-ticker.C
		receipt, err = client.TxReceiptDetailByHash(txHash)
		if err != nil && !errors.Is(err, ethereum.NotFound) {
			log.Error("get transaction by hash fail", "error", err)
			return nil, err
		}

		if errors.Is(err, ethereum.NotFound) {
			continue
		}
		return receipt, nil
	}
}

func (s *RpcServer) newPools(newPools []*bindings.IL1PoolManagerPool) (newCPools []bindings.IL1PoolManagerPool, err error) {
	var totalFee *big.Int

	for _, newPool := range newPools {
		newPool.TotalFeeClaimed = new(big.Int).SetUint64(0)
		newPool.TotalAmount = new(big.Int).SetUint64(0)
		newPool.TotalFee = new(big.Int).SetUint64(0)
		totalFee, err = s.db.CrossChainTransfer.GetPeriodTotalFee(uint64(s.poolStartTimestamp), uint64(s.poolEndTimestamp), newPool.Token)
		newFee := new(big.Int).Mul(totalFee, big.NewInt(92))
		newFee.Div(newFee, big.NewInt(100))
		newPool.TotalFee = newFee

		newCPools = append(newCPools, *newPool)
	}

	return newCPools, err
}

func (s *RpcServer) metricX1StrategyBalance() error {
	if err := s.DaStrategyWETHToL2DappLinkBridge(s.x1ChainId); err != nil {
		return err
	}

	if err := s.GamingStrategyWETHToL2DappLinkBridge(s.x1ChainId); err != nil {
		return err
	}

	if err := s.SocialStrategyWETHToL2DappLinkBridge(s.x1ChainId); err != nil {
		return err
	}
	return nil
}

func (s *RpcServer) metricOpStrategyBalance() error {
	if err := s.DaStrategyETHToL2DappLinkBridge(s.opChainId); err != nil {
		return err
	}

	if err := s.GamingStrategyETHToL2DappLinkBridge(s.opChainId); err != nil {
		return err
	}

	if err := s.SocialStrategyETHToL2DappLinkBridge(s.opChainId); err != nil {
		return err
	}
	return nil
}

func (s *RpcServer) DaStrategyETHToL2DappLinkBridge(chainID uint64) error {
	var cOpts *bind.CallOpts
	var tOpts *bind.TransactOpts
	var tx *types.Transaction
	var finalTx *types.Transaction
	var receipt *types.Receipt
	var ctx = context.Background()

	latestBlock, err := s.ethClients[chainID].GetLatestBlock()
	if err != nil {
		log.Error("get latest block number fail", "err", err)
		return err
	}

	cOpts = &bind.CallOpts{
		BlockNumber: latestBlock,
		From:        crypto.PubkeyToAddress(s.privateKey.PublicKey),
	}

	ethBalance, err := s.DAStrategyContract[chainID].ETHBalance(cOpts)
	if err != nil {
		log.Error("get da strategy eth balance fail", "err", err)
		return err
	}
	if ethBalance.Cmp(new(big.Int).SetUint64(32)) > -1 {
		tOpts, err = s.newTransactOpts(ctx, chainID)

		tx, err = s.DAStrategyContract[chainID].TransferETHToL2DappLinkBridge(tOpts, new(big.Int).SetUint64(chainID), new(big.Int).SetUint64(s.l1ChainID), s.L1BridgeContractAddress, s.l1StakingManagerAddr, new(big.Int).SetUint64(21000))
		if err != nil {
			log.Error("transfer da eth to l2 dapp-link bridge by abi fail", "error", err)
			return err
		}
		finalTx, err = s.RawDAStrategyContract[chainID].RawTransact(tOpts, tx.Data())
		if err != nil {
			log.Error("raw da transfer eth to l2 dapp-link bridge fail", "error", err)
			return err
		}
		err = s.ethClients[chainID].SendTransaction(ctx, finalTx)
		if err != nil {
			log.Error("send da transfer eth to l2 dapp-link bridge transaction fail", "error", err)
			return err
		}
		receipt, err = getTransactionReceipt(s.ethClients[chainID], finalTx.Hash())
		if err != nil {
			log.Error("get da transfer eth to l2 dapp-link bridge receipt fail", "error", err)
			return err
		}

		log.Info("send da transfer eth to l2 dapp-link bridge transaction success", "tx_hash", receipt.TxHash.String())
	}
	return nil
}

func (s *RpcServer) GamingStrategyETHToL2DappLinkBridge(chainID uint64) error {
	var cOpts *bind.CallOpts
	var tOpts *bind.TransactOpts
	var tx *types.Transaction
	var finalTx *types.Transaction
	var receipt *types.Receipt
	var ctx = context.Background()

	latestBlock, err := s.ethClients[chainID].GetLatestBlock()
	if err != nil {
		log.Error("get latest block number fail", "err", err)
		return err
	}

	cOpts = &bind.CallOpts{
		BlockNumber: latestBlock,
		From:        crypto.PubkeyToAddress(s.privateKey.PublicKey),
	}

	ethBalance, err := s.GamingStrategyContract[chainID].ETHBalance(cOpts)
	if err != nil {
		log.Error("get gaming strategy eth balance fail", "err", err)
		return err
	}
	if ethBalance.Cmp(new(big.Int).SetUint64(32)) > -1 {
		tOpts, err = s.newTransactOpts(ctx, chainID)

		tx, err = s.GamingStrategyContract[chainID].TransferETHToL2DappLinkBridge(tOpts, new(big.Int).SetUint64(chainID), new(big.Int).SetUint64(s.l1ChainID), s.L1BridgeContractAddress, s.l1StakingManagerAddr, new(big.Int).SetUint64(21000))
		if err != nil {
			log.Error("transfer gaming eth to l2 dapp-link bridge by abi fail", "error", err)
			return err
		}
		finalTx, err = s.RawGamingStrategyContract[chainID].RawTransact(tOpts, tx.Data())
		if err != nil {
			log.Error("raw transfer gaming eth to l2 dapp-link bridge fail", "error", err)
			return err
		}
		err = s.ethClients[chainID].SendTransaction(ctx, finalTx)
		if err != nil {
			log.Error("send gaming transfer eth to l2 dapp-link bridge transaction fail", "error", err)
			return err
		}
		receipt, err = getTransactionReceipt(s.ethClients[chainID], finalTx.Hash())
		if err != nil {
			log.Error("get gaming transfer eth to l2 dapp-link bridge receipt fail", "error", err)
			return err
		}

		log.Info("send gaming transfer eth to l2 dapp-link bridge transaction success", "tx_hash", receipt.TxHash.String())
	}
	return nil
}

func (s *RpcServer) SocialStrategyETHToL2DappLinkBridge(chainID uint64) error {
	var cOpts *bind.CallOpts
	var tOpts *bind.TransactOpts
	var tx *types.Transaction
	var finalTx *types.Transaction
	var receipt *types.Receipt
	var ctx = context.Background()

	latestBlock, err := s.ethClients[chainID].GetLatestBlock()
	if err != nil {
		log.Error("get latest block number fail", "err", err)
		return err
	}

	cOpts = &bind.CallOpts{
		BlockNumber: latestBlock,
		From:        crypto.PubkeyToAddress(s.privateKey.PublicKey),
	}

	ethBalance, err := s.SocialStrategyContract[chainID].ETHBalance(cOpts)
	if err != nil {
		log.Error("get social strategy eth balance fail", "err", err)
		return err
	}
	if ethBalance.Cmp(new(big.Int).SetUint64(32)) > -1 {
		tOpts, err = s.newTransactOpts(ctx, chainID)

		tx, err = s.SocialStrategyContract[chainID].TransferETHToL2DappLinkBridge(tOpts, new(big.Int).SetUint64(chainID), new(big.Int).SetUint64(s.l1ChainID), s.L1BridgeContractAddress, s.l1StakingManagerAddr, new(big.Int).SetUint64(21000))
		if err != nil {
			log.Error("transfer social eth to l2 dapp-link bridge by abi fail", "error", err)
			return err
		}
		finalTx, err = s.RawSocialStrategyContract[chainID].RawTransact(tOpts, tx.Data())
		if err != nil {
			log.Error("raw transfer social eth to l2 dapp-link bridge fail", "error", err)
			return err
		}
		err = s.ethClients[chainID].SendTransaction(ctx, finalTx)
		if err != nil {
			log.Error("send social transfer eth to l2 dapp-link bridge transaction fail", "error", err)
			return err
		}
		receipt, err = getTransactionReceipt(s.ethClients[chainID], finalTx.Hash())
		if err != nil {
			log.Error("get social transfer eth to l2 dapp-link bridge receipt fail", "error", err)
			return err
		}

		log.Info("send social transfer eth to l2 dapp-link bridge transaction success", "tx_hash", receipt.TxHash.String())
	}
	return nil
}

func (s *RpcServer) DaStrategyWETHToL2DappLinkBridge(chainID uint64) error {
	var cOpts *bind.CallOpts
	var tOpts *bind.TransactOpts
	var tx *types.Transaction
	var finalTx *types.Transaction
	var receipt *types.Receipt
	var ctx = context.Background()

	latestBlock, err := s.ethClients[chainID].GetLatestBlock()
	if err != nil {
		log.Error("get latest block number fail", "err", err)
		return err
	}

	cOpts = &bind.CallOpts{
		BlockNumber: latestBlock,
		From:        crypto.PubkeyToAddress(s.privateKey.PublicKey),
	}

	wethBalance, err := s.DAStrategyContract[chainID].WETHBalance(cOpts)
	if err != nil {
		log.Error("get da strategy weth balance fail", "err", err)
		return err
	}
	if wethBalance.Cmp(new(big.Int).SetUint64(32)) > -1 {
		tOpts, err = s.newTransactOpts(ctx, chainID)

		tx, err = s.DAStrategyContract[chainID].TransferWETHToL2DappLinkBridge(tOpts, new(big.Int).SetUint64(chainID), new(big.Int).SetUint64(s.l1ChainID), s.L1BridgeContractAddress, s.l1StakingManagerAddr, s.WEthAddress[chainID], new(big.Int).SetUint64(21000))
		if err != nil {
			log.Error("transfer da weth to l2 dapp-link bridge by abi fail", "error", err)
			return err
		}
		finalTx, err = s.RawDAStrategyContract[chainID].RawTransact(tOpts, tx.Data())
		if err != nil {
			log.Error("raw da transfer weth to l2 dapp-link bridge fail", "error", err)
			return err
		}
		err = s.ethClients[chainID].SendTransaction(ctx, finalTx)
		if err != nil {
			log.Error("send da transfer weth to l2 dapp-link bridge transaction fail", "error", err)
			return err
		}
		receipt, err = getTransactionReceipt(s.ethClients[chainID], finalTx.Hash())
		if err != nil {
			log.Error("get da transfer weth to l2 dapp-link bridge receipt fail", "error", err)
			return err
		}

		log.Info("send da transfer weth to l2 dapp-link bridge transaction success", "tx_hash", receipt.TxHash.String())
	}
	return nil
}

func (s *RpcServer) GamingStrategyWETHToL2DappLinkBridge(chainID uint64) error {
	var cOpts *bind.CallOpts
	var tOpts *bind.TransactOpts
	var tx *types.Transaction
	var finalTx *types.Transaction
	var receipt *types.Receipt
	var ctx = context.Background()

	latestBlock, err := s.ethClients[chainID].GetLatestBlock()
	if err != nil {
		log.Error("get latest block number fail", "err", err)
		return err
	}

	cOpts = &bind.CallOpts{
		BlockNumber: latestBlock,
		From:        crypto.PubkeyToAddress(s.privateKey.PublicKey),
	}

	wethBalance, err := s.GamingStrategyContract[chainID].WETHBalance(cOpts)
	if err != nil {
		log.Error("get gaming strategy weth balance fail", "err", err)
		return err
	}
	if wethBalance.Cmp(new(big.Int).SetUint64(32)) > -1 {
		tOpts, err = s.newTransactOpts(ctx, chainID)

		tx, err = s.GamingStrategyContract[chainID].TransferWETHToL2DappLinkBridge(tOpts, new(big.Int).SetUint64(chainID), new(big.Int).SetUint64(s.l1ChainID), s.L1BridgeContractAddress, s.l1StakingManagerAddr, s.WEthAddress[chainID], new(big.Int).SetUint64(21000))
		if err != nil {
			log.Error("transfer gaming weth to l2 dapp-link bridge by abi fail", "error", err)
			return err
		}
		finalTx, err = s.RawGamingStrategyContract[chainID].RawTransact(tOpts, tx.Data())
		if err != nil {
			log.Error("raw gaming transfer weth to l2 dapp-link bridge fail", "error", err)
			return err
		}
		err = s.ethClients[chainID].SendTransaction(ctx, finalTx)
		if err != nil {
			log.Error("send gaming transfer weth to l2 dapp-link bridge transaction fail", "error", err)
			return err
		}
		receipt, err = getTransactionReceipt(s.ethClients[chainID], finalTx.Hash())
		if err != nil {
			log.Error("get gaming transfer weth to l2 dapp-link bridge receipt fail", "error", err)
			return err
		}

		log.Info("send da transfer weth to l2 dapp-link bridge transaction success", "tx_hash", receipt.TxHash.String())
	}
	return nil
}

func (s *RpcServer) SocialStrategyWETHToL2DappLinkBridge(chainID uint64) error {
	var cOpts *bind.CallOpts
	var tOpts *bind.TransactOpts
	var tx *types.Transaction
	var finalTx *types.Transaction
	var receipt *types.Receipt
	var ctx = context.Background()

	latestBlock, err := s.ethClients[chainID].GetLatestBlock()
	if err != nil {
		log.Error("get latest block number fail", "err", err)
		return err
	}

	cOpts = &bind.CallOpts{
		BlockNumber: latestBlock,
		From:        crypto.PubkeyToAddress(s.privateKey.PublicKey),
	}

	wethBalance, err := s.SocialStrategyContract[chainID].WETHBalance(cOpts)
	if err != nil {
		log.Error("get social strategy weth balance fail", "err", err)
		return err
	}
	if wethBalance.Cmp(new(big.Int).SetUint64(32)) > -1 {
		tOpts, err = s.newTransactOpts(ctx, chainID)

		tx, err = s.SocialStrategyContract[chainID].TransferWETHToL2DappLinkBridge(tOpts, new(big.Int).SetUint64(chainID), new(big.Int).SetUint64(s.l1ChainID), s.L1BridgeContractAddress, s.l1StakingManagerAddr, s.WEthAddress[chainID], new(big.Int).SetUint64(21000))
		if err != nil {
			log.Error("transfer social weth to l2 dapp-link bridge by abi fail", "error", err)
			return err
		}
		finalTx, err = s.RawSocialStrategyContract[chainID].RawTransact(tOpts, tx.Data())
		if err != nil {
			log.Error("raw social transfer weth to l2 dapp-link bridge fail", "error", err)
			return err
		}
		err = s.ethClients[chainID].SendTransaction(ctx, finalTx)
		if err != nil {
			log.Error("send social transfer weth to l2 dapp-link bridge transaction fail", "error", err)
			return err
		}
		receipt, err = getTransactionReceipt(s.ethClients[chainID], finalTx.Hash())
		if err != nil {
			log.Error("get social transfer weth to l2 dapp-link bridge receipt fail", "error", err)
			return err
		}

		log.Info("send social transfer weth to l2 dapp-link bridge transaction success", "tx_hash", receipt.TxHash.String())
	}
	return nil
}
