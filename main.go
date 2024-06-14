package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Trello API test")

	headers := map[string]string{"Accept": "application/json"}

	err := godotenv.Load()

	handleError(err)

	api_key := os.Getenv("API_KEY")
	oauth := os.Getenv("OAUTH_TOKEN")
	//prepare query parameters

	query := url.Values{}
	query.Set("key", api_key)
	query.Set("token", oauth)

	_url := "https://trello.com/b/GKkeJ61j/apitest/list" + query.Encode()

	req, err := http.NewRequest("GET", _url, nil)

	if err != nil {
		fmt.Println(err)
	}

	//set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the HTTP request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status code
	fmt.Println("Response Status:", resp.Status)

	// Print the response body
	// In a real application, you would typically read and process the body
	// For simplicity, we are just printing the response body as a string here
	// Note: This assumes the response body is textual, not binary
	body := make([]byte, 512)
	n, err := resp.Body.Read(body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println(n)

}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
