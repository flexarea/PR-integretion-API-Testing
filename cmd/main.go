package main

import (
	"log"
	"os"

	"github.com/flexarea/PR-integration-API-Testing/pkg/models"
)

type Application struct {
	logs    *models.LogsModel
	infoLog *log.Logger
	errLog  *log.Logger
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime)

	app := &Application{
		infoLog: infoLog,
		errLog:  errorLog,
	}

	infoLog.Println("Running app")
	app.Server()
}
