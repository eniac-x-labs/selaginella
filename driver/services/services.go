package services

import (
	"context"
	"fmt"
	"net"
	"sync/atomic"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ethereum/go-ethereum/log"

	"github.com/evm-layer2/selaginella/database"
	"github.com/evm-layer2/selaginella/protobuf/pb"
)

const MaxRecvMessageSize = 1024 * 1024 * 300

type RpcServerConfig struct {
	GrpcHostname string
	GrpcPort     string
}

type RpcServer struct {
	*RpcServerConfig
	db *database.DB
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

func NewRpcServer(db *database.DB, config *RpcServerConfig) (*RpcServer, error) {
	return &RpcServer{
		RpcServerConfig: config,
		db:              db,
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

func (s *RpcServer) CrossChainTransact(ctx context.Context, in *pb.BridgeRequest) (*pb.BridgeResponse, error) {
	return &pb.BridgeResponse{
		Success: true,
		Message: "call cross chain transfer success",
	}, nil
}
