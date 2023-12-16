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
	// Root context is used to propagate cancellations to all spawned goroutines
	ctx, cancel := context.WithCancel(context.Background())

	// Register cancellation on interrupts
	go func() {
		opio.BlockOnInterrupts()
		cancel()
	}()

	// Configure logging
	logConfig := log.LvlFilterHandler(log.LvlInfo, log.StderrHandler)
	log.Root().SetHandler(logConfig)

	// Create and run the CLI application
	app := newCli(GitCommit, GitDate)
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Error("application failed", "err", err)
		os.Exit(1) // Return a non-zero exit code to indicate failure
	}
}
