package config

import (
	"os"
)

type Config struct {
	DBSource    string `mapstructure:"DB_SOURCE"`
	Environment string `mapstructure:"ENVIRONMENT"`
	Secret      string `mapstructure:"SECRET"`
}

func NewConfig() Config {
	return Config{
		DBSource:    os.Getenv("DB_SOURCE"),
		Environment: os.Getenv("ENVIRONMENT"),
		Secret:      os.Getenv("SECRET"),
	}
}
