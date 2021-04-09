package app

import (
	logger "github.com/sirupsen/logrus"
	"github.com/tetovske/enrollments-parser/pkg/util"
	"os"
)

func Start() (err error) {
	logger.Println("Parser started")
	conf, err := util.ParseConfig(os.Getenv("PROJ_ROOT") + "/configs/config.yml")
	if err != nil {
		logger.Fatal(err)
		return err
	}

	logger.Println(conf)
	return err
}
