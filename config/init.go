package config

import (
	"github.com/spf13/viper"
)

const ConfigName string = "config"

func Init() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName(ConfigName)

	return viper.ReadInConfig()
}
