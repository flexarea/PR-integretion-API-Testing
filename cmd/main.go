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

	infoLog.Println("Running app")
	infoLog.Println("try: /newGitActionUpdate")

	Server()
}
