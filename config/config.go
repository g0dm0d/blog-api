package config

import (
	"github.com/pelletier/go-toml"
	"os"
)

type ServerConfig struct {
	Addr   string
	Port   string
	Secret string
}

type Database struct {
	DSN string
}

type Config struct {
	Server   *ServerConfig
	Database *Database
}

func Load() (*Config, error) {
	var cfg Config

	dat, err := os.ReadFile("./config.toml")
	if err != nil {
		return &cfg, err
	}
	err = toml.Unmarshal(dat, &cfg)
	if err != nil {
		return &cfg, err
	}

	return &cfg, nil
}
