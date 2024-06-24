package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Calling Slack API")
	//load .env
	err := godotenv.Load()

	if err != nil {
		return
	}
	bot_token := os.Getenv("BOT_TOKEN")
	main_end_point := os.Getenv("MAIN_END_POINT")
	fmt.Println(bot_token)
	fmt.Println(main_end_point)
	channelID := "C06KPMXQS4U"
	conversationHistory(bot_token, main_end_point, channelID)
}

func conversationHistory(bot_token, main_end_point, channelID string) {
	//create url

	url := fmt.Sprintf("%sconversations.history", main_end_point)
	fmt.Println(url)
	//create json payload for channelid
	payload := map[string]string{"channel": channelID}
	//marshal json
	payloadByte, err := json.Marshal(payload)

	if err != nil {
		return
	}
	//new request
	req, err := http.NewRequest("POST", url, bytes.NewReader(payloadByte))
	if err != nil {
		return
	}
	//set authorization header
	req.Header.Set("Authorization", "Bearer "+bot_token)
	req.Header.Set("Content-Type", "application/json")

	//hangle client response
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("error receiving response", err)
	}

	databyte, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	//format json body data
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, databyte, "", " ")
	if err != nil {
		log.Fatal("Error formatting JSON: ", err)
	}
	fmt.Println(prettyJSON.String())
}
