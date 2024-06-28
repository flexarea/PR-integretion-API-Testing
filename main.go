package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Testing from cmd")

	//load .env
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env", err)
	}

	//loading environment variables for slack
	bot_token := os.Getenv("BOT_TOKEN")
	slack_main_end_point := os.Getenv("SLACK_MAIN_END_POINT")
	fmt.Println(bot_token)
	fmt.Println(slack_main_end_point)
}
