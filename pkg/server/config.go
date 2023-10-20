package server

import (
	"fmt"
	"github.com/caarlos0/env/v9"
	"github.com/gojaguar/jaguar/config"
)

// Config holds the configuration for this service. It currently uses env vars, but it can eventually
// migrate to a different config provider.
type Config struct {
	config.Config
	DB config.Database `envPrefix:"DATABASE_"`
}

// Listener returns the configuration needed to initialize a net.Listener instance.
func (c Config) Listener() (network string, address string) {
	return "tcp", fmt.Sprintf(":%d", c.Port)
}

// ReadConfig reads the Config from environment variables.
func ReadConfig() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
