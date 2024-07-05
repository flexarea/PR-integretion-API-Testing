package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/flexarea/PR-integration-API-Testing/cmd/internal/slack"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime)
	//launch server
	configs, err := Load_config()

	if err != nil {
		errorLog.Fatal(err)
	}

	fmt.Println(configs)
	message := fmt.Sprint(time.Now())
	channelID := "C06KPMXQS4U"
	infoLog.Println("new slack message")
	slack.SendMessage(configs.BOT_TOKEN, configs.SLACK_MAIN_END_POINT, channelID, message)
}
