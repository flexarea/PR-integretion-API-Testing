package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Configs struct {
	API_KEY        string
	API_TOKEN      string
	BOARD_ID       string
	MAIN_END_POINT string
	LIST_ID        string
}

// load configuration variables
func Load_config() (*Configs, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}
	config := &Configs{
		API_KEY:        os.Getenv("API_KEY"),
		API_TOKEN:      os.Getenv("API_TOKEN"),
		BOARD_ID:       os.Getenv("BOARD_ID"),
		MAIN_END_POINT: os.Getenv("MAIN_END_POINT"),
		LIST_ID:        os.Getenv("LIST_ID"),
	}

	return config, err
}
