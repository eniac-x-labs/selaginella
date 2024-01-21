package main

import (
	"context"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/evm-layer2/selaginella/common/cliapp"
	"github.com/evm-layer2/selaginella/config"
	"github.com/evm-layer2/selaginella/database"
	"github.com/evm-layer2/selaginella/services"
)

var (
	ConfigFlag = &cli.StringFlag{
		Name:    "config",
		Value:   "./selaginella.yaml",
		Aliases: []string{"c"},
		Usage:   "path to config file",
		EnvVars: []string{"ACORUS_CONFIG"},
	}
	MigrationsFlag = &cli.StringFlag{
		Name:    "migrations-dir",
		Value:   "./migrations",
		Usage:   "path to migrations folder",
		EnvVars: []string{"ACORUS_MIGRATIONS_DIR"},
	}
)

func runGrpcServer(ctx *cli.Context, _ context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	cfg, err := config.NewConfig(ctx.String(ConfigFlag.Name))
	if err != nil {
		log.Error("failed to load config", "err", err)
		return nil, err
	}
	grpcServerCfg := &services.RpcServerConfig{
		GrpcHostname: cfg.Server.Host,
		GrpcPort:     strconv.Itoa(cfg.Server.Port),
	}
	db, err := database.NewDB(ctx.Context, cfg.Database)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return nil, err
	}
	return services.NewRpcServer(db, grpcServerCfg)
}

func runMigrations(ctx *cli.Context) error {
	log.Info("running migrations...")
	cfg, err := config.NewConfig(ctx.String(ConfigFlag.Name))
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}
	db, err := database.NewDB(ctx.Context, cfg.Database)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	return db.ExecuteSQLMigration(ctx.String(MigrationsFlag.Name))
}

func newCli(GitCommit string, GitDate string) *cli.App {
	flags := []cli.Flag{ConfigFlag}
	migrationFlags := []cli.Flag{MigrationsFlag, ConfigFlag}
	return &cli.App{
		Version:              params.VersionWithCommit(GitCommit, GitDate),
		Description:          "An indexer of all mantle da block data with a serving grpc api layer",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "grpc",
				Flags:       flags,
				Description: "Runs the grpc service",
				Action:      cliapp.LifecycleCmd(runGrpcServer),
			},
			{
				Name:        "migrate",
				Flags:       migrationFlags,
				Description: "Runs the database migrations",
				Action:      runMigrations,
			},
			{
				Name:        "version",
				Description: "print version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}
