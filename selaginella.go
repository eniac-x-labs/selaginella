package selaginella

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/evm-layer2/selaginella/config"
	"github.com/evm-layer2/selaginella/database"
)

type Selaginella struct {
	db              *database.DB
	redis           *redis.Client
	metricsConfig   config.ServerConfig
	metricsRegistry *prometheus.Registry
}

func NewSelaginella(
	db *database.DB,
	redis *redis.Client,
	chainConfig config.ChainConfig,
	rpcsConfig config.RPCsConfig,
	httpConfig config.ServerConfig,
	metricsConfig config.ServerConfig,
) (*Selaginella, error) {
	return nil, nil
}

func (s *Selaginella) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	errCh := make(chan error, 5)

	// if any goroutine halts, we stop the entire selaginella
	processCtx, processCancel := context.WithCancel(ctx)
	runProcess := func(start func(ctx context.Context) error) {
		wg.Add(1)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Error("halting selaginella on panic", "err", err)
					errCh <- fmt.Errorf("panic: %v", err)
				}
				processCancel()
				wg.Done()
			}()
			errCh <- start(processCtx)
		}()
	}

	// synchronizer engine
	runProcess(nil)

	// event engine
	runProcess(nil)

	// metrics server
	runProcess(nil)

	// api server
	runProcess(nil)

	wg.Wait()

	err := <-errCh
	if err != nil {
		log.Error("selaginella stopped", "err", err)
	} else {
		log.Error("selaginella stopped")
	}
	return err
}
