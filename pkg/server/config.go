package server

import (
	"fmt"
	"github.com/caarlos0/env/v9"
	"github.com/gojaguar/jaguar/config"
)

type Config struct {
	config.Config
	DB config.Database
}

func (c Config) Listener() (network string, address string) {
	return "tcp", fmt.Sprintf(":%d", c.Port)
}

func ReadConfig() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
