package services

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"net"
	"strings"
	"sync/atomic"
	"time"

	kms "cloud.google.com/go/kms/apiv1"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
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
	db                  *database.DB
	ethClients          map[uint64]node.EthClient
	RawL1BridgeContract *bind.BoundContract
	RawL2BridgeContract map[uint64]*bind.BoundContract
	L1BridgeContract    *bindings.L1PoolManager
	L2BridgeContract    map[uint64]*bindings.L2PoolManager
	EthAddress          map[uint64]common.Address
	WEthAddress         map[uint64]common.Address
	USDTAddress         map[uint64]common.Address
	USDCAddress         map[uint64]common.Address
	DAIAddress          map[uint64]common.Address
	OKBAddress          map[uint64]common.Address
	MNTAddress          map[uint64]common.Address
	poolStartTimestamp  uint32
	poolEndTimestamp    uint32
	pb.UnimplementedBridgeServiceServer
	stopped atomic.Bool
	pb.BridgeServiceServer
	privateKey    *ecdsa.PrivateKey
	l1ChainID     uint64
	zkFairChainId uint64
	x1ChainId     uint64
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

func NewRpcServer(ctx context.Context, db *database.DB, grpcCfg *RpcServerConfig, hsmCfg *HsmConfig, chainRpcCfg []*config.RPC, priKey *ecdsa.PrivateKey, l1ChainID uint64, zkFairChainID uint64, x1ChainID uint64, mantleChainID uint64, shutdown context.CancelCauseFunc) (*RpcServer, error) {
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

	for i := range chainRpcCfg {
		if chainRpcCfg[i].ChainId == l1ChainID {
			client, err := node.DialEthClient(ctx, chainRpcCfg[i].RpcUrl)
			if err != nil {
				log.Error("dial client fail", "error", err.Error())
				return nil, err
			}
			ethClients[chainRpcCfg[i].ChainId] = client

			l1Parsed, err := abi.JSON(strings.NewReader(
				bindings.L1PoolManagerABI,
			))
			if err != nil {
				log.Error("selaginella parse l1 pool contract abi fail", "err", err)
				return nil, err
			}

			l1Client, _ := node.DialEthClientWithTimeout(ctx, chainRpcCfg[i].RpcUrl, false)
			rawL1PoolContract := bind.NewBoundContract(
				common.HexToAddress(chainRpcCfg[i].FoundingPoolAddress), l1Parsed, l1Client, l1Client,
				l1Client,
			)
			rawL1BridgeContract = rawL1PoolContract

			l1PoolContract, err := bindings.NewL1PoolManager(common.HexToAddress(chainRpcCfg[i].FoundingPoolAddress), l1Client)
			l1BridgeContract = l1PoolContract
			EthAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].EthAddress)
			WEthAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].WEthAddress)
			USDTAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].USDTAddress)
			USDCAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].USDCAddress)
			DAIAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].DAIAddress)
			OKBAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].OKBAddress)
			MNTAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].MNTAddress)

		} else {
			client, err := node.DialEthClient(ctx, chainRpcCfg[i].RpcUrl)
			if err != nil {
				log.Error("dial client fail", "error", err.Error())
				return nil, err
			}
			ethClients[chainRpcCfg[i].ChainId] = client

			l2Parsed, err := abi.JSON(strings.NewReader(
				bindings.L2PoolManagerABI,
			))
			if err != nil {
				log.Error("selaginella parse l2 pool contract abi fail", "err", err)
				return nil, err
			}

			l2Client, _ := node.DialEthClientWithTimeout(ctx, chainRpcCfg[i].RpcUrl, false)
			rawL2PoolContract := bind.NewBoundContract(
				common.HexToAddress(chainRpcCfg[i].FoundingPoolAddress), l2Parsed, l2Client, l2Client,
				l2Client,
			)
			rawL2BridgeContracts[chainRpcCfg[i].ChainId] = rawL2PoolContract

			l2PoolContract, err := bindings.NewL2PoolManager(common.HexToAddress(chainRpcCfg[i].FoundingPoolAddress), l2Client)
			l2BridgeContracts[chainRpcCfg[i].ChainId] = l2PoolContract
			EthAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].EthAddress)
			WEthAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].WEthAddress)
			USDTAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].USDTAddress)
			USDCAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].USDCAddress)
			DAIAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].DAIAddress)
			OKBAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].OKBAddress)
			MNTAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].MNTAddress)
		}

	}

	return &RpcServer{
		RpcServerConfig:     grpcCfg,
		db:                  db,
		HsmConfig:           hsmCfg,
		ethClients:          ethClients,
		RawL1BridgeContract: rawL1BridgeContract,
		RawL2BridgeContract: rawL2BridgeContracts,
		L1BridgeContract:    l1BridgeContract,
		L2BridgeContract:    l2BridgeContracts,
		EthAddress:          EthAddress,
		WEthAddress:         WEthAddress,
		USDTAddress:         USDTAddress,
		USDCAddress:         USDCAddress,
		DAIAddress:          DAIAddress,
		OKBAddress:          OKBAddress,
		MNTAddress:          MNTAddress,
		privateKey:          priKey,
		l1ChainID:           l1ChainID,
		zkFairChainId:       zkFairChainID,
		x1ChainId:           x1ChainID,
		mantleChainId:       mantleChainID,
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

	SBTTicker := time.NewTicker(3 * time.Second)
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

func (s *RpcServer) SendBridgeTransaction() error {

	var opts *bind.TransactOpts
	var err error
	var tx *types.Transaction
	var finalTx *types.Transaction
	var ctx = context.Background()
	var nilCrossChainTransfer database.CrossChainTransfer

	bridgeTx, _ := s.db.CrossChainTransfer.OldestPendingNoSentTransaction()
	if *bridgeTx == nilCrossChainTransfer {
		log.Warn("no more bridge transaction need to be sent")
		return nil
	}

	for chainId, client := range s.ethClients {
		if chainId == bridgeTx.DestChainId.Uint64() {
			opts, err = s.newTransactOpts(ctx, chainId)
			if err != nil {
				return err
			}

			log.Info(fmt.Sprintf("get send transaction request: sourceChainId=%v, destChainId=%v, amount=%v, fee=%v, nonce=%v, tokenAddress=%v", bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce, bridgeTx.TokenAddress))

			switch bridgeTx.TokenAddress.String() {
			case s.EthAddress[chainId].String():
				if chainId == s.l1ChainID {
					if bridgeTx.SourceChainId.Uint64() == s.zkFairChainId {
						tx, err = s.L1BridgeContract.BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.USDCAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("zkfair transfer usdc to l1 by abi fail", "error", err)
							return err
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.x1ChainId {
						tx, err = s.L1BridgeContract.BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.OKBAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("x1 transfer okb to l1 by abi fail", "error", err)
							return err
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.mantleChainId {
						tx, err = s.L1BridgeContract.BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.MNTAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("mantle transfer mnt to l1 by abi fail", "error", err)
							return err
						}
					} else {
						tx, err = s.L1BridgeContract.BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("transfer eth to l1 by abi fail", "error", err)
							return err
						}
					}
				} else if chainId == s.x1ChainId {
					if bridgeTx.SourceChainId.Uint64() == s.zkFairChainId {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.USDCAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("zkfair transfer usdc to x1 by abi fail", "error", err)
							return err
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.mantleChainId {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.MNTAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("mantle transfer mnt to x1 by abi fail", "error", err)
							return err
						}
					} else {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("transfer eth to x1 by abi fail", "error", err)
							return err
						}
					}
				} else if chainId == s.zkFairChainId {
					if bridgeTx.SourceChainId.Uint64() == s.x1ChainId {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.OKBAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("x1 transfer okb to zkfair by abi fail", "error", err)
							return err
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.mantleChainId {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.MNTAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("mantle transfer mnt to zkfair by abi fail", "error", err)
							return err
						}
					} else {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("transfer eth to zkfair by abi fail", "error", err)
							return err
						}
					}
				} else if chainId == s.mantleChainId {
					if bridgeTx.SourceChainId.Uint64() == s.x1ChainId {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.OKBAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("x1 transfer okb to mantle by abi fail", "error", err)
							return err
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.zkFairChainId {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.USDCAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("mantle transfer mnt to mantle by abi fail", "error", err)
							return err
						}
					} else {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("transfer eth to mantle by abi fail", "error", err)
							return err
						}
					}
				} else {
					tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
					if err != nil {
						log.Error("transfer eth to l2 by abi fail", "error", err)
						return err
					}
				}

			case s.WEthAddress[chainId].String():
				if chainId == s.l1ChainID {
					tx, err = s.L1BridgeContract.BridgeFinalizeWETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
					if err != nil {
						log.Error("get bridge finalize l1 weth by abi fail", "error", err)
						return err
					}
				} else {
					tx, err = s.L2BridgeContract[chainId].BridgeFinalizeWETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
					if err != nil {
						log.Error("get bridge finalize l2 weth by abi fail", "error", err)
						return err
					}
				}

			default:
				if chainId == s.l1ChainID {
					if bridgeTx.SourceChainId.Uint64() == s.x1ChainId {
						if bridgeTx.TokenAddress.String() == s.WEthAddress[s.x1ChainId].String() {
							tx, err = s.L1BridgeContract.BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("x1 transfer eth to l1 by abi fail", "error", err)
								return err
							}
						} else {
							tx, err = s.L1BridgeContract.BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("x1 transfer erc20 to l1 by abi fail", "error", err)
								return err
							}
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.zkFairChainId {
						if bridgeTx.TokenAddress.String() == s.WEthAddress[s.zkFairChainId].String() {
							tx, err = s.L1BridgeContract.BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("zkfair transfer eth to l1 by abi fail", "error", err)
								return err
							}
						} else {
							tx, err = s.L1BridgeContract.BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("zkfair transfer erc20 to l1 by abi fail", "error", err)
								return err
							}
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.mantleChainId {
						if bridgeTx.TokenAddress.String() == s.WEthAddress[s.mantleChainId].String() {
							tx, err = s.L1BridgeContract.BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer eth to l1 by abi fail", "error", err)
								return err
							}
						} else {
							tx, err = s.L1BridgeContract.BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer erc20 to l1 by abi fail", "error", err)
								return err
							}
						}
					} else {
						tx, err = s.L1BridgeContract.BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("transfer erc20 to l1 by abi fail", "error", err)
							return err
						}
					}
				} else if chainId == s.zkFairChainId {
					if bridgeTx.SourceChainId.Uint64() == s.l1ChainID {
						if bridgeTx.TokenAddress.String() == s.EthAddress[s.l1ChainID].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("l1 transfer eth to zkfair by abi fail", "error", err)
								return err
							}
						} else if bridgeTx.TokenAddress.String() == s.USDCAddress[s.l1ChainID].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("l1 transfer usdc to zkfair by abi fail", "error", err)
								return err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("l1 transfer erc20 to zkfair by abi fail", "error", err)
								return err
							}
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.x1ChainId {
						if bridgeTx.TokenAddress.String() == s.WEthAddress[s.x1ChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("x1 transfer eth to zkfair by abi fail", "error", err)
								return err
							}
						} else if bridgeTx.TokenAddress.String() == s.USDCAddress[s.x1ChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("x1 transfer usdc to zkfair by abi fail", "error", err)
								return err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("x1 transfer erc20 to zkfair by abi fail", "error", err)
								return err
							}
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.mantleChainId {
						if bridgeTx.TokenAddress.String() == s.WEthAddress[s.mantleChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer eth to zkfair by abi fail", "error", err)
								return err
							}
						} else if bridgeTx.TokenAddress.String() == s.USDCAddress[s.mantleChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer usdc to zkfair by abi fail", "error", err)
								return err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer erc20 to zkfair by abi fail", "error", err)
								return err
							}
						}
					} else {
						if bridgeTx.TokenAddress.String() == s.USDCAddress[bridgeTx.SourceChainId.Uint64()].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer usdc to zkfair by abi fail", "error", err)
								return err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("transfer erc20 to zkfair by abi fail", "error", err)
								return err
							}
						}
					}
				} else if chainId == s.x1ChainId {
					if bridgeTx.SourceChainId.Uint64() == s.l1ChainID {
						if bridgeTx.TokenAddress.String() == s.EthAddress[s.l1ChainID].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("l1 transfer eth to x1 by abi fail", "error", err)
								return err
							}
						} else if bridgeTx.TokenAddress.String() == s.OKBAddress[s.l1ChainID].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("l1 transfer okb to x1 by abi fail", "error", err)
								return err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("l1 transfer erc20 to x1 by abi fail", "error", err)
								return err
							}
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.zkFairChainId {
						if bridgeTx.TokenAddress.String() == s.WEthAddress[s.zkFairChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("zkfair transfer eth to x1 by abi fail", "error", err)
								return err
							}
						} else if bridgeTx.TokenAddress.String() == s.OKBAddress[s.zkFairChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("zkfair transfer okb to x1 by abi fail", "error", err)
								return err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("zkfair transfer erc20 to x1 by abi fail", "error", err)
								return err
							}
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.mantleChainId {
						if bridgeTx.TokenAddress.String() == s.WEthAddress[s.mantleChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, s.WEthAddress[chainId], bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer eth to x1 by abi fail", "error", err)
								return err
							}
						} else if bridgeTx.TokenAddress.String() == s.OKBAddress[s.mantleChainId].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer okb to x1 by abi fail", "error", err)
								return err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("mantle transfer erc20 to x1 by abi fail", "error", err)
								return err
							}
						}
					} else {
						if bridgeTx.TokenAddress.String() == s.OKBAddress[bridgeTx.SourceChainId.Uint64()].String() {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("transfer okb to x1 by abi fail", "error", err)
								return err
							}
						} else {
							tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
							if err != nil {
								log.Error("transfer erc20 to x1 by abi fail", "error", err)
								return err
							}
						}
					}
				} else {
					if bridgeTx.SourceChainId.Uint64() == s.x1ChainId && bridgeTx.TokenAddress == s.WEthAddress[s.x1ChainId] {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("x1 transfer eth to l2 by abi fail", "error", err)
							return err
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.zkFairChainId && bridgeTx.TokenAddress == s.WEthAddress[s.zkFairChainId] {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("zkfair transfer eth to l2 by abi fail", "error", err)
							return err
						}
					} else if bridgeTx.SourceChainId.Uint64() == s.mantleChainId && bridgeTx.TokenAddress == s.WEthAddress[s.mantleChainId] {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeETH(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("mantle transfer eth to l2 by abi fail", "error", err)
							return err
						}
					} else {
						tx, err = s.L2BridgeContract[chainId].BridgeFinalizeERC20(opts, bridgeTx.SourceChainId, bridgeTx.DestChainId, bridgeTx.DestReceiveAddress, bridgeTx.TokenAddress, bridgeTx.Amount, bridgeTx.Fee, bridgeTx.Nonce)
						if err != nil {
							log.Error("transfer erc20 to l2 by abi fail", "error", err)
							return err
						}
					}
				}
			}

			log.Info("get bridge finalize by abi success")

			if chainId == s.l1ChainID {
				finalTx, err = s.RawL1BridgeContract.RawTransact(opts, tx.Data())
				if err != nil {
					log.Error("raw send bridge transaction fail", "error", err)
					return err
				}
			} else {
				finalTx, err = s.RawL2BridgeContract[chainId].RawTransact(opts, tx.Data())
				if err != nil {
					log.Error("raw send bridge transaction fail", "error", err)
					return err
				}
			}

			err = client.SendTransaction(ctx, finalTx)
			if err != nil {
				log.Error("send bridge transaction fail", "error", err)
				return err
			}
			bridgeTx.TxHash = finalTx.Hash()
		}
	}

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
		return err
	}

	return nil
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

func (s *RpcServer) ChangeBridgeTransactionStatus() error {
	var receipt *types.Receipt

	tx, err := s.db.CrossChainTransfer.OldestPendingSentTransaction()
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
		log.Info("no more pending transaction !")
		return nil
	}

	err = s.db.CrossChainTransfer.ChangeCrossChainTransferSentStatueByTxHash(receipt.TxHash.String())
	if err != nil {
		log.Error("change transaction status fail", "err", err)
		return err
	}

	log.Info("change bridge transaction status success", "txHash", receipt.TxHash)

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
		log.Info("no more pending transaction !")
		return nil
	}

	err = s.db.UpdateDepositFundingPoolBalance.ChangeUpdateFundingPoolBalanceSentStatueByTxHash(receipt.TxHash.String())
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
		log.Info("no more pending transaction !")
		return nil
	}

	err = s.db.UpdateWithdrawFundingPoolBalance.ChangeUpdateFundingPoolBalanceSentStatueByTxHash(receipt.TxHash.String())
	if err != nil {
		log.Error("change withdraw update pool balance transaction status fail", "err", err)
		return err
	}

	log.Info("change update withdraw pool balance transaction status success", "txHash", receipt.TxHash)

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
		receipt, err = getTransactionReceipt(s.ethClients[s.l1ChainID], *finalTx)
		if err != nil {
			log.Error("get complete pool and new receipt fail", "error", err)
			return err
		}

		log.Info("send complete pool and new transaction success", "tx_hash", receipt.TxHash)

	}

	return nil
}

func getTransactionReceipt(client node.EthClient, tx types.Transaction) (*types.Receipt, error) {
	var receipt *types.Receipt
	var err error

	ticker := time.NewTicker(30 * time.Second)
	for {
		<-ticker.C
		receipt, err = client.TxReceiptDetailByHash(tx.Hash())
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
