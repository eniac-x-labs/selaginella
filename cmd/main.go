package cmd

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
	app := newCli(GitCommit, GitDate)
	ctx := opio.WithInterruptBlocker(context.Background())
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Error("application failed", "err", err)
		os.Exit(1)
	}
}
