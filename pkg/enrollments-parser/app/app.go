package app

import (
	logger "github.com/sirupsen/logrus"
	"github.com/tetovske/enrollments-parser/config"
	"github.com/tetovske/enrollments-parser/pkg/enrollments-parser/services"
)

func Start() (err error) {
	logger.Info("Parser started")
	err = config.Init()
	if err != nil {
		logger.Fatal(err)
		return
	}

	parsingProcessor, err := services.NewProcessor()
	if err != nil {
		logger.Fatal(err)
		return err
	}

	err = parsingProcessor.ProcessDocuments()
	if err != nil {
		return
	}
	logger.Info(parsingProcessor)

	return
}
