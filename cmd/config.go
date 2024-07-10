package main

import (
	"fmt"
	"os"
	// "github.com/joho/godotenv"
)

func Test() {
	env, err := Load_config()
	if err != nil {
		return
	}
	fmt.Println(env)
}

type Configs struct {
	BOT_TOKEN             string
	SLACK_MAIN_END_POINT  string
	API_KEY               string
	API_TOKEN             string
	TRELLO_MAIN_END_POINT string
}

// load configuration variables
func Load_config() (Configs, error) {
	config := Configs{
		BOT_TOKEN:            os.Getenv("BOT_TOKEN"),
		SLACK_MAIN_END_POINT: os.Getenv("SLACK_MAIN_END_POINT"),
	}

	if config.BOT_TOKEN == "" {
		return Configs{}, fmt.Errorf("BOT_TOKEN env not set")
	}
	if config.SLACK_MAIN_END_POINT == "" {
		return Configs{}, fmt.Errorf("SLACK_MAIN_END_POINT env not set")
	}

	return config, nil
}
