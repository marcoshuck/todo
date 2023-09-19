package conf

import "github.com/gojaguar/jaguar/config"

type Config struct {
	config.Config
	DB config.Database `envPrefix:"DB_"`
}
