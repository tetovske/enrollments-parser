package main

import (
	logger "github.com/sirupsen/logrus"
	"github.com/tetovske/enrollments-parser/internal/app"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		logger.Fatal("Invalid app arguments")
		return
	}

	for i, s := range os.Args {
		logger.Println(i, s)
	}
	app.Start(os.Args[1])
}
