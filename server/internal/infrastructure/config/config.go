package config

import (
	"os"
)

type Config struct {
	DBSource            string `mapstructure:"DB_SOURCE"`
	Environment         string `mapstructure:"ENVIRONMENT"`
	Secret              string `mapstructure:"SECRET"`
	RandomServiceURL    string `mapstructure:"RANDOM_STRING_SERVICE_URL"`
	RandomServiceApiKey string `mapstructure:"RANDOM_STRING_SERVICE_API_KEY"`
}

func NewConfig() *Config {
	return &Config{
		DBSource:            os.Getenv("DB_SOURCE"),
		Environment:         os.Getenv("ENVIRONMENT"),
		Secret:              os.Getenv("SECRET"),
		RandomServiceURL:    os.Getenv("RANDOM_STRING_SERVICE_URL"),
		RandomServiceApiKey: os.Getenv("RANDOM_STRING_SERVICE_API_KEY"),
	}
}
