package main

import (
	"log"
	"os"

	"github.com/flexarea/PR-integration-API-Testing/pkg/models"
)

type Application struct {
	logs *models.LogsModel
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	infoLog.Println("Running app")
	app := &Application{
		logs: &models.LogsModel{DB: db},
	}

	Server()
}
