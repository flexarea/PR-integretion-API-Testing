package main

import (
	"fmt"
	"github.com/flexarea/PR-integration-API-Testing/internal/config"
	"github.com/flexarea/PR-integration-API-Testing/internal/slack"
)

func main() {
	fmt.Println("Testing from cmd")

	//load .env
	configs, err := config.Load_config()
	if err != nil {
		return
	}
	//loading environment variables for slack
	channelID := "C06KPMXQS4U"
	newMessage := "sending from branch v.1.0.1"
	slack.SendMessage(configs.BOT_TOKEN, configs.SLACK_MAIN_END_POINT, channelID, newMessage)

}
