package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	AppName  string `toml:"app_name"`
	Database struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Name     string `toml:"name"`
	} `toml:"database"`
	Rpc struct {
		Port int `toml:"port"`
	}
}

// New Setup init config
func New(path string) (*Config, error) {
	// config global config instance
	var config = new(Config)

	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(f, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
