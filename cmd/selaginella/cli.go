package selaginella

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/evm-layer2/selaginella"
	"github.com/evm-layer2/selaginella/config"
	"github.com/evm-layer2/selaginella/database"
	"github.com/urfave/cli/v2"
)

var (
	ConfigFlag = &cli.StringFlag{
		Name:    "config",
		Value:   "./selaginella.toml",
		Aliases: []string{"c"},
		Usage:   "path to config file",
		EnvVars: []string{"INDEXER_CONFIG"},
	}
	MigrationsFlag = &cli.StringFlag{
		Name:    "migrations-dir",
		Value:   "./migrations",
		Usage:   "path to migrations folder",
		EnvVars: []string{"INDEXER_MIGRATIONS_DIR"},
	}
)

func runSyncer(ctx *cli.Context) error {
	cfg, err := config.LoadConfig(ctx.String(ConfigFlag.Name))
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}

	db, err := database.NewDB(cfg.MasterDB)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	defer db.Close()

	redis, err := database.NewRedis(cfg.Redis)

	lithosphere, err := selaginella.NewSelaginella(db, redis, cfg.Chain, cfg.RPCs, cfg.HTTPServer, cfg.MetricsServer)
	if err != nil {
		log.Error("failed to create lithosphere", "err", err)
		return err
	}

	log.Info("running lithosphere...")
	return lithosphere.Run(ctx.Context)
}

func runApi(ctx *cli.Context) error {
	cfg, err := config.LoadConfig(ctx.String(ConfigFlag.Name))
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}
	var db *database.DB
	// slave database for api services
	if !cfg.SlaveDbEnable {
		db, err = database.NewDB(cfg.MasterDB)
		if err != nil {
			log.Error("failed to connect to master database", "err", err)
			return err
		}
	} else {
		db, err = database.NewDB(cfg.SlaveDB)
		if err != nil {
			log.Error("failed to connect to slave database", "err", err)
			return err
		}
	}
	defer db.Close()

	log.Info("running notify rpc ...")
	// todo: rpc new here
	return nil
}

func runMigrations(ctx *cli.Context) error {
	cfg, err := config.LoadConfig(ctx.String(ConfigFlag.Name))
	migrationsDir := ctx.String(MigrationsFlag.Name)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}

	// If slave database not enable, use master database
	if cfg.SlaveDbEnable {
		slaveDB, err := database.NewDB(cfg.SlaveDB)
		if err != nil {
			log.Error("failed to connect to slave database", "err", err)
			return err
		}
		defer slaveDB.Close()
		err = slaveDB.ExecuteSQLMigration(migrationsDir)
		if err != nil {
			log.Error("migrate slave database fail", "err", err)
			return err
		}
	}

	masterDB, err := database.NewDB(cfg.MasterDB)
	if err != nil {
		log.Error("failed to connect to master database", "err", err)
		return err
	}
	defer masterDB.Close()

	err = masterDB.ExecuteSQLMigration(migrationsDir)
	if err != nil {
		log.Error("migrate master database fail", "err", err)
		return err
	}
	return nil
}

func newCli(GitCommit string, GitDate string) *cli.App {
	flags := []cli.Flag{ConfigFlag}
	migrationFlags := []cli.Flag{MigrationsFlag, ConfigFlag}
	return &cli.App{
		Version:              params.VersionWithCommit(GitCommit, GitDate),
		Description:          "An lithosphere of all optimism events with a serving api layer",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "api",
				Flags:       flags,
				Description: "Runs the api service",
				Action:      runApi,
			},
			{
				Name:        "syncer",
				Flags:       flags,
				Description: "Runs the sycer service",
				Action:      runSyncer,
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
