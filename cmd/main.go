package main

import (
	configs "Zhantasov/config"
	"Zhantasov/internal/handler"
	"Zhantasov/internal/pkg"
	"Zhantasov/internal/server"
	"Zhantasov/internal/service"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var filePaths = []string{

	"./exam_data/5-сынып.json",
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			os.Exit(1)
		}
	}()

	logrus.SetFormatter(new(logrus.JSONFormatter))

	cnf, err := configs.New()
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	testFiles, err := pkg.GetTestFiles(filePaths)
	if err != nil {
		logrus.Fatal(err)
	}
	service := service.New(cnf, testFiles)

	handler := handler.New(service)

	srv := new(server.Server)

	if err := srv.Run(cnf.Port, handler.InitRouters()); err != nil {
		logrus.Fatalf("err of listening server, %s", err.Error())
	}

}
