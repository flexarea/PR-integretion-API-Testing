package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Test() {
	env, err := Load_config()
	if err != nil {
		return
	}
	fmt.Println(env)
}

type Configs struct {
	BOT_TOKEN            string
	SLACK_MAIN_END_POINT string
}

// load configuration variables
func Load_config() (Configs, error) {
	err := godotenv.Load()

	if err != nil {
		return Configs{}, err
	}
	config := Configs{
		BOT_TOKEN:            os.Getenv("BOT_TOKEN"),
		SLACK_MAIN_END_POINT: os.Getenv("SLACK_MAIN_END_POINT"),
	}

	return config, err
}
