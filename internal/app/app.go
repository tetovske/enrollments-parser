package app

import (
	logger "github.com/sirupsen/logrus"
	"github.com/tetovske/enrollments-parser/pkg/config"
)

func Start(configPath string) (err error) {
	logger.Println("Parser started")
	conf, err := config.ParseConfig(configPath)
	if err != nil {
		logger.Fatal(err)
		return err
	}

	logger.Println(conf)
	return err
}
