package main

import (
	"log"
	"os"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	infoLog.Println("Running app")
	Server()
}
