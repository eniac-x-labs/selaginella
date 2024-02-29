package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DB struct {
	DbHost     string `yaml:"db_host"`
	DbPort     int    `yaml:"db_port"`
	DbName     string `yaml:"db_name"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
}

type RPC struct {
	RpcUrl              string `yaml:"rpc_url"`
	ChainId             uint64 `yaml:"chain_id"`
	FoundingPoolAddress string `yaml:"founding_pool_address"`
	EthAddress          string `yaml:"eth_address"`
	WEthAddress         string `yaml:"weth_address"`
	USDTAddress         string `yaml:"usdt_address"`
	USDCAddress         string `yaml:"usdc_address"`
	DAIAddress          string `yaml:"dai_address"`
}

type Exporter struct {
	ExporterAddress string `yaml:"exporter_address"`
}

type Config struct {
	Migrations string   `yaml:"migrations"`
	Server     Server   `yaml:"server"`
	RPCs       []*RPC   `yaml:"rpcs"`
	Metrics    Server   `yaml:"metrics"`
	Database   DB       `yaml:"database"`
	Exporter   Exporter `yaml:"exporter"`
}

func NewConfig(path string) (*Config, error) {
	var config = new(Config)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
