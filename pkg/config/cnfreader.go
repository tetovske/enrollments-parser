package config

import (
	logger "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	App struct {
		loggerPath string `yaml: "logger_location"`
	} `yaml: "app"`

	Parser struct {
		DocumentsLocation string `yaml: "docs_location"`
	} `yaml: "enrl_parser"`
}

func ParseConfig(configLocation string) (*Config, error) {
	f, err := ioutil.ReadFile(configLocation)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	confData := Config{}
	err = yaml.Unmarshal(f, &confData)
	if err != nil {
		logger.Fatal(err)
	}

	return &confData, err
}
