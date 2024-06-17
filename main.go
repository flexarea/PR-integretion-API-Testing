package main

import (
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
	//retrieve data from .env
	err := godotenv.Load()

	handleError(err)
	api_key := os.Getenv("API_KEY")
	//oauth := os.Getenv("OAUTH_TOKEN")
	api_token := os.Getenv("API_TOKEN")

	fmt.Println("api_key: ", api_key)
	fmt.Println("api_token: ", api_token)
	gettingBoard(api_key, api_token)
}
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func gettingBoard(api_key string, api_token string) {

	//board_id := "666c33eced5913ce5f990639"

	xURL := "https://api.trello.com/1/members/me/boards?fields=name,url&key=" + api_key + "&token=" + api_token

	req, err := http.NewRequest("GET", xURL, nil) //calling api, getting board info

	if err != nil {
		log.Fatal("Error creating request: ", err)
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
		fmt.Println("Error reading body: ", err.Error())
	}
	// Print the response status code
	fmt.Println("Response Status:", resp.Status)

	// Convert response body to string
	bodyString := string(bodybyte)
	fmt.Println("API Response as String:\n" + bodyString)

	// If needed, you can unmarshal into a map for dynamic inspection
}

func updatingBoard() {
	fmt.Println("Updating board")
}
