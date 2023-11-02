package server

import (
	"fmt"
	"github.com/caarlos0/env/v9"
	"github.com/gojaguar/jaguar/config"
)

type Metrics struct {
	Host string `env:"HOST" envDefault:"localhost"`
	Port uint16 `env:"PORT" envDefault:"4317"`
}

func (m Metrics) Address() string {
	return fmt.Sprintf("%s:%d", m.Host, m.Port)
}

type Tracing struct {
	Host string `env:"HOST" envDefault:"localhost"`
	Port uint16 `env:"PORT" envDefault:"4317"`
}

func (t Tracing) Address() string {
	return fmt.Sprintf("%s:%d", t.Host, t.Port)
}

// Config holds the configuration for this service. It currently uses env vars, but it can eventually
// migrate to a different config provider.
type Config struct {
	config.Config
	DB      config.Database `envPrefix:"DATABASE_"`
	Metrics Metrics         `envPrefix:"METRICS_"`
	Tracing Tracing         `envPrefix:"TRACING_"`
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
