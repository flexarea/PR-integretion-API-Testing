package main

import (
	"bytes"
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
	//retrieve data from .env
	err := godotenv.Load()

	handleError(err)
	api_key := os.Getenv("API_KEY")
	//oauth := os.Getenv("OAUTH_TOKEN")
	api_token := os.Getenv("API_TOKEN")
	//gettingBoard(api_key, api_token)
	updatingBoard(api_key, api_token)

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

func updatingBoard(api_key string, api_token string) {
	fmt.Println("Updating board")

	board_id := "666c33eced5913ce5f990639"
	newname := "Testing_api"
	_url := "https://api.trello.com/1/boards/" + board_id + "?key=" + api_key + "&token=" + api_token
	//create JSON payload
	payload := map[string]string{"name": newname}

	payloadBytes, err := json.Marshal(payload)
	fmt.Println(payload)
	fmt.Println(payloadBytes)
	if err != nil {
		log.Fatal("Error marshalling JSON: ", err)
	}
	body := bytes.NewBuffer(payloadBytes) //converts payloadBytes into buffer so can be used as io.reader
	//create http request

	req, err := http.NewRequest("PUT", _url, body)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}

	//set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	// Read and print the response body

	bodybyte, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading body: ", err.Error())
	}
	// Print the response status code
	fmt.Println("Response Status:", resp.Status)

	// Convert response body to string
	bodyString := string(bodybyte)
	fmt.Println("Response body:\n" + bodyString)

}
