package selaginlla

import (
	"context"
	"github.com/evm-layer2/selaginella/config"
	"google.golang.org/grpc"

	"github.com/evm-layer2/selaginella/driver/bridge"
)

type SelaginellaConfig struct {
}

type Selaginella struct {
	ctx          context.Context
	cfg          *SelaginellaConfig
	bridgeDriver bridge.Bridge
	gpcService   *grpc.Server
}

func NewMantleBatch(cfg config.Config) (*Selaginella, error) {
	return nil, nil
}

func (sg *Selaginella) Start() error {
	return nil
}

func (sg *Selaginella) Stop() {

}
