package util

import (
	"github.com/spf13/viper"
)

const ConfigPath = "configs/config.yml"

type Config struct {
	App struct {
		loggerPath string `yaml: "logger_location"`
	} `yaml: "app"`

	Parser struct {
		DocumentsLocation string `yaml: "docs_location"`
	} `yaml: "enrl_parser"`
}

func ParseConfig(configPath string) (config Config, err error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	config = Config{}
	err = viper.Unmarshal(&config)
	return
}
