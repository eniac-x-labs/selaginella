package services

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ethereum/go-ethereum/log"

	"github.com/evm-layer2/selaginella/protobuf/pb"
)

const localhost = "0.0.0.0"

type Config struct {
	Hostname string
	GrpcPort string
}

type Server struct {
	*Config
	pb.BridgeServiceServer
}

func NewServer(config *Config) *Server {
	return &Server{
		Config: config,
	}
}

func (s *Server) Start() {
	go func(s *Server) {
		addr := fmt.Sprintf("%s:%s", localhost, s.GrpcPort)
		log.Info("start rpc server", "addr", addr)

		listener, err := net.Listen("tcp", addr)
		if err != nil {
			log.Error("Could not start tcp listener. ")
		}

		opt := grpc.MaxRecvMsgSize(1024 * 1024 * 300)

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
}

func (s *Server) crossChainTransfer(ctx context.Context, in *pb.BridgeRequest) (*pb.BridgeResponse, error) {
	return nil, nil
}
