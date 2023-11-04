package conf

import (
	"fmt"
	"github.com/caarlos0/env/v9"
	"github.com/gojaguar/jaguar/config"
)

type Metrics struct {
	Enabled bool   `env:"ENABLED" envDefault:"true"`
	Host    string `env:"HOST" envDefault:"localhost"`
	Port    uint16 `env:"PORT" envDefault:"4317"`
}

func (m Metrics) Address() string {
	return fmt.Sprintf("%s:%d", m.Host, m.Port)
}

type Tracing struct {
	Enabled bool   `env:"ENABLED" envDefault:"true"`
	Host    string `env:"HOST" envDefault:"localhost"`
	Port    uint16 `env:"PORT" envDefault:"4317"`
}

func (t Tracing) Address() string {
	return fmt.Sprintf("%s:%d", t.Host, t.Port)
}

// ServerConfig holds the configuration for gRPC servers. It currently uses env vars, but it can eventually
// migrate to a different config provider.
type ServerConfig struct {
	config.Config
	DB      config.Database `envPrefix:"DATABASE_"`
	Metrics Metrics         `envPrefix:"METRICS_"`
	Tracing Tracing         `envPrefix:"TRACING_"`
}

// Listener returns the configuration needed to initialize a net.Listener instance.
func (c ServerConfig) Listener() (network string, address string) {
	return "tcp", fmt.Sprintf(":%d", c.Port)
}

// ReadServerConfig reads the ServerConfig from environment variables.
func ReadServerConfig() (ServerConfig, error) {
	var cfg ServerConfig
	err := env.Parse(&cfg)
	if err != nil {
		return ServerConfig{}, err
	}
	return cfg, nil
}

// ClientConfig holds the configuration for gRPC clients. It currently uses env vars, but it can eventually
// migrate to a different config provider.
type ClientConfig struct {
	config.Config
	ServerAddress string  `env:"SERVER_ADDRESS"`
	Metrics       Metrics `envPrefix:"METRICS_"`
	Tracing       Tracing `envPrefix:"TRACING_"`
}

// ReadClientConfig reads the ClientConfig from environment variables.
func ReadClientConfig() (ClientConfig, error) {
	var cfg ClientConfig
	err := env.Parse(&cfg)
	if err != nil {
		return ClientConfig{}, err
	}
	return cfg, nil
}
