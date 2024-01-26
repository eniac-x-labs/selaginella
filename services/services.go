package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evm-layer2/selaginella/bindings"
	"github.com/evm-layer2/selaginella/common/retry"
	"github.com/evm-layer2/selaginella/protobuf/pb"
	"math/big"
	"net"
	"strconv"
	"strings"
	"sync/atomic"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	common2 "github.com/evm-layer2/selaginella/common"
	"github.com/evm-layer2/selaginella/config"
	"github.com/evm-layer2/selaginella/database"
	node "github.com/evm-layer2/selaginella/eth_client"
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
	L1BridgeContract    *bindings.L1Pool
	L2BridgeContract    map[uint64]*bindings.L2Pool
	pb.UnimplementedBridgeServiceServer
	stopped atomic.Bool
	pb.BridgeServiceServer
}

func (s *RpcServer) Stop(ctx context.Context) error {
	s.stopped.Store(true)
	return nil
}

func (s *RpcServer) Stopped() bool {
	return s.stopped.Load()
}

func NewRpcServer(ctx context.Context, db *database.DB, grpcCfg *RpcServerConfig, hsmCfg *HsmConfig, chainRpcCfg []*config.RPC) (*RpcServer, error) {
	ethClients := make(map[uint64]node.EthClient)
	var rawL1BridgeContract *bind.BoundContract
	rawL2BridgeContracts := make(map[uint64]*bind.BoundContract)
	var l1BridgeContract *bindings.L1Pool
	l2BridgeContracts := make(map[uint64]*bindings.L2Pool)
	for i := range chainRpcCfg {
		if chainRpcCfg[i].ChainId == 1 {
			client, err := node.DialEthClient(ctx, chainRpcCfg[i].RpcUrl)
			if err != nil {
				log.Error("dial client fail", "error", err.Error())
				return nil, err
			}
			ethClients[chainRpcCfg[i].ChainId] = client

			l1Parsed, err := abi.JSON(strings.NewReader(
				bindings.L1PoolABI,
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

			l1PoolContract, err := bindings.NewL1Pool(common.HexToAddress(chainRpcCfg[i].FoundingPoolAddress), l1Client)
			l1BridgeContract = l1PoolContract

		} else {
			client, err := node.DialEthClient(ctx, chainRpcCfg[i].RpcUrl)
			if err != nil {
				log.Error("dial client fail", "error", err.Error())
				return nil, err
			}
			ethClients[chainRpcCfg[i].ChainId] = client

			l2Parsed, err := abi.JSON(strings.NewReader(
				bindings.L2PoolABI,
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

			l2PoolContract, err := bindings.NewL2Pool(common.HexToAddress(chainRpcCfg[i].FoundingPoolAddress), l2Client)
			l2BridgeContracts[chainRpcCfg[i].ChainId] = l2PoolContract
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
	}, nil
}

func (s *RpcServer) Start(ctx context.Context) error {
	go func(s *RpcServer) {
		addr := fmt.Sprintf("%s:%s", s.GrpcHostname, s.GrpcPort)
		log.Info("start rpc server", "addr", addr)

		listener, err := net.Listen("tcp", addr)
		if err != nil {
			log.Error("Could not start tcp listener. ")
		}

		opt := grpc.MaxRecvMsgSize(MaxRecvMessageSize)

		gs := grpc.NewServer(
			opt,
			grpc.ChainUnaryInterceptor(),
		)
		reflection.Register(gs)
		pb.RegisterBridgeServiceServer(gs, s)

		log.Info("grp info", "port", s.GrpcPort, "address", listener.Addr().String())
		if err := gs.Serve(listener); err != nil {
			log.Error("Could not GRPC server")
		}
	}(s)
	return nil
}

func (s *RpcServer) CrossChainTransfer(ctx context.Context, in *pb.CrossChainTransferRequest) (*pb.CrossChainTransferResponse, error) {
	if in == nil {
		return nil, errors.New("invalid request: request body is empty")
	}

	var opts *bind.TransactOpts
	var err error
	var tx *types.Transaction
	var finalTx *types.Transaction
	var crossChainTransfers []database.CrossChainTransfer

	for chainId, client := range s.ethClients {
		destChainId, _ := strconv.ParseUint(in.DestChainId, 10, 64)
		if chainId == destChainId {
			if s.EnableHsm {
				opts, err = sign.NewHSMTransactOpts(ctx, s.HsmAPIName,
					s.HsmAddress, new(big.Int).SetUint64(chainId), s.HsmCreden)
			}

			switch in.TokenAddress {
			case common2.ETH:
				tx, err = s.L1BridgeContract.DepositAndStakingETH(opts)
				if err != nil {
					log.Error("get bridge transaction by abi fail", "error", err)
					return nil, err
				}
			default:
				amount, _ := new(big.Int).SetString(in.Amount, 10)
				tx, err = s.L1BridgeContract.DepositAndStarkingERC20(opts, common.HexToAddress(in.TokenAddress), amount)
				if err != nil {
					log.Error("get bridge transaction by abi fail", "error", err)
					return nil, err
				}
			}

			finalTx, err = s.RawL2BridgeContract[chainId].RawTransact(opts, tx.Data())
			if err != nil {
				log.Error("raw send bridge transaction fail", "error", err)
				return nil, err
			}

			err = client.SendTransaction(ctx, finalTx)
			if err != nil {
				log.Error("send bridge transaction fail", "error", err)
			}
			crossChainTransfer := s.db.CrossChainTransfer.BuildCrossChainTransfer(in)
			crossChainTransfers = append(crossChainTransfers, crossChainTransfer)
		}

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
	}

	return &pb.CrossChainTransferResponse{
		Success: true,
		Message: "call cross chain transfer success",
	}, nil
}

func (s *RpcServer) ChangeTransferStatus(ctx context.Context, in *pb.CrossChainTransferStatusRequest) (*pb.CrossChainTransferStatusResponse, error) {
	if in == nil {
		return nil, errors.New("invalid request: request body is empty")
	}

	err := s.db.CrossChainTransfer.ChangeCrossChainTransferStatueByTxHash(in.TxHash)
	if err != nil {
		log.Error("change cross chain transfer status fail", "err", err)
		return nil, err
	}

	return &pb.CrossChainTransferStatusResponse{
		Success: true,
		Message: "call cross chain transfer success",
	}, nil
}
