package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
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
	DB_HOST               string
	DB_DATABASE           string
	DB_USERNAME           string
	DB_PASSWORD           string
}

// load configuration variables
func Load_config() (Configs, error) {
	godotenv.Load(".env")
	config := Configs{
		BOT_TOKEN:            os.Getenv("BOT_TOKEN"),
		SLACK_MAIN_END_POINT: os.Getenv("SLACK_MAIN_END_POINT"),
		DB_HOST:              os.Getenv("DB_HOST"),
		DB_DATABASE:          os.Getenv("DB_DATABASE"),
		DB_USERNAME:          os.Getenv("DB_USERNAME"),
		DB_PASSWORD:          os.Getenv("DB_PASSWORD"),
	}

	if config.BOT_TOKEN == "" {
		return Configs{}, fmt.Errorf("BOT_TOKEN env not set")
	}
	if config.SLACK_MAIN_END_POINT == "" {
		return Configs{}, fmt.Errorf("SLACK_MAIN_END_POINT env not set")
	}

	return config, nil
}
