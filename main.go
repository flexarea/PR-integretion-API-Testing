package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

//https://api.trello.com/1/actions/?key=APIKey&token=APIToken
//https://api.trello.com/1/boards/{idBoard}?key={yourKey}&token={yourToken}'

func main() {
	fmt.Println("Calling trello api")

	err := godotenv.Load()

	handleError(err)

	api_key := os.Getenv("API_KEY")
	oauth := os.Getenv("OAUTH_TOKEN")
	api_token := os.Getenv("API_TOKEN")

	fmt.Println("api_key: ", api_key)
	fmt.Println("oauth: ", oauth)
	//board_id := "666c33eced5913ce5f990639"

	//_url := "https://api.trello.com/1/actions/?key=" + api_key + "&token=" + oauth
	//xURL := "https://api.trello.com/1/members/me/boards?key=" + api_key + "&token=" + oauth
	xURL := "https://api.trello.com/1/members/me/?key=" + api_key + "&token=" + api_token

	req, err := http.NewRequest("GET", xURL, nil)

	if err != nil {
		fmt.Println(err)
	}

	//set headers

	req.Header.Add("Accept", "application/json")

	// Send the HTTP request
	client := &http.Client{} //memory address so we can reuse the the http.client instance for multiple request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
	}
	defer resp.Body.Close()

	bodybyte, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
	}
	// Print the response status code
	fmt.Println("Response Status:", resp.Status)

	// Convert response body to string
	bodyString := string(bodybyte)
	fmt.Println("API Response as String:\n" + bodyString)

	// If needed, you can unmarshal into a map for dynamic inspection
	var result map[string]interface{}
	json.Unmarshal(bodybyte, &result)
	fmt.Printf("API Response as map: %+v\n", result)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
