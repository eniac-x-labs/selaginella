package main

import (
	"context"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/evm-layer2/selaginella/common/opio"
)

var (
	GitCommit = ""
	GitDate   = ""
)

func main() {
	// This is the most root context, used to propagate
	// cancellations to all spawned application-level goroutines
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		opio.BlockOnInterrupts()
		cancel()
	}()

	app := newCli(GitCommit, GitDate)
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Error("application failed", "err", err)
	}
}
