package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBSource    string `mapstructure:"DB_SOURCE"`
	Environment string `mapstructure:"ENVIRONMENT"`
	Secret      string `mapstructure:"SECRET"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("local")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
