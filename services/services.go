package services

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net"
	"strconv"
	"sync/atomic"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

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
	db                *database.DB
	ethClients        map[uint64]node.EthClient
	RawBridgeContract map[uint64]*bind.BoundContract
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
	for i := range chainRpcCfg {
		client, err := node.DialEthClient(ctx, chainRpcCfg[i].RpcUrl)
		if err != nil {
			log.Error("dial client fail", "error", err.Error())
			return nil, err
		}
		ethClients[chainRpcCfg[i].ChainId] = client
	}

	return &RpcServer{
		RpcServerConfig: grpcCfg,
		db:              db,
		HsmConfig:       hsmCfg,
		ethClients:      ethClients,
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
	var finalTx *types.Transaction
	for chainId, client := range s.ethClients {
		destChainId, _ := strconv.ParseUint(in.DestChainId, 10, 64)
		if chainId == destChainId {
			if s.EnableHsm {
				opts, err = sign.NewHSMTransactOpts(ctx, s.HsmAPIName,
					s.HsmAddress, new(big.Int).SetUint64(chainId), s.HsmCreden)
			}
			// todo get binding tx
			finalTx, err = s.RawBridgeContract[chainId].RawTransact(opts, tx)

			err = client.SendTransaction(ctx, finalTx)
			if err != nil {
				log.Error("send bridge transaction fail", "error", err)
			}
		}

	}

	return &pb.CrossChainTransferResponse{
		Success: true,
		Message: "call cross chain transfer success",
	}, nil
}

func (s *RpcServer) ChangeTransferStatus(ctx context.Context, in *pb.CrossChainTransferStatusRequest) (*pb.CrossChainTransferStatusResponse, error) {
	return &pb.CrossChainTransferStatusResponse{
		Success: true,
		Message: "call cross chain transfer success",
	}, nil
}
