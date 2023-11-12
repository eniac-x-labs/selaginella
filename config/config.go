package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/log"
)

const defaultLoopInterval = 5000

type Config struct {
	Chain         ChainConfig  `toml:"chain"`
	RPCs          RPCsConfig   `toml:"rpcs"`
	Redis         RedisConfig  `toml:"redis"`
	MasterDB      DBConfig     `toml:"master_db"`
	SlaveDB       DBConfig     `toml:"slave_db"`
	SlaveDbEnable bool         `toml:"slave_db_enable"`
	HTTPServer    ServerConfig `toml:"http"`
	MetricsServer ServerConfig `toml:"metrics"`
}

type ChainConfig struct {
	StartingHeight    uint `toml:"starting-height"`
	ConfirmationDepth uint `toml:"l1-confirmation-depth"`
	PollingInterval   uint `toml:"l1-polling-interval"`
}

type RPCsConfig struct {
	EvmRPC string `toml:"evm-rpc"`
}

type DBConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Name     string `toml:"name"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

type ServerConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type RedisConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

func LoadConfig(path string) (Config, error) {
	log.Debug("loading config", "path", path)

	var cfg Config
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	data = []byte(os.ExpandEnv(string(data)))
	log.Debug("parsed config file", "data", string(data))

	if _, err := toml.Decode(string(data), &cfg); err != nil {
		log.Error("failed to decode config file", "err", err)
		return cfg, err
	}

	if cfg.Chain.PollingInterval == 0 {
		cfg.Chain.PollingInterval = defaultLoopInterval
	}

	log.Info("loaded chain config", "config", cfg.Chain)
	return cfg, nil
}
