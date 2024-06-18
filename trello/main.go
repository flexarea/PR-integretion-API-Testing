package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"trello/utils"

	"github.com/joho/godotenv"
)

//https://api.trello.com/1/actions/?key=APIKey&token=APIToken
//https://api.trello.com/1/boards/{idBoard}?key={yourKey}&token={yourToken}'

func main() {
	fmt.Println("Calling trello api")
	//retrieve data from .env
	err := godotenv.Load()

	utils.HandleError(err)
	api_key := os.Getenv("API_KEY")
	board_id := os.Getenv("BOARD_ID")
	api_token := os.Getenv("API_TOKEN")
	//gettingBoardLists(api_key, api_token, board_id)
	//updatingBoard(api_key, api_token)
	//gettingActions(api_key, api_token)
	creatingList(api_key, api_token, board_id)

}

func gettingBoard(api_key string, api_token string) {

	xURL := "https://api.trello.com/1/members/me/boards?key=" + api_key + "&token=" + api_token

	req, err := http.NewRequest("GET", xURL, nil) //calling api, getting board info

	utils.ReqError(err)
	//set headers

	req.Header.Add("Accept", "application/json")

	// Send the HTTP request
	client := &http.Client{} //memory address so we can reuse the the http.client instance for multiple request
	resp, err := client.Do(req)
	utils.RespError(err)
	defer resp.Body.Close()

	bodybyte, err := io.ReadAll(resp.Body)

	utils.ReadBodyError(err)
	// Print the response status code
	fmt.Println("Response Status:", resp.Status)

	// Convert response body to string
	bodyString := string(bodybyte)
	fmt.Println("API Response as String:\n" + bodyString)

}

func updatingBoard(api_key string, api_token string, board_id string) {
	fmt.Println("Updating board")

	newname := "Testing_api"
	_url := "https://api.trello.com/1/boards/" + board_id + "?key=" + api_key + "&token=" + api_token
	//create JSON payload
	payload := map[string]string{"name": newname}

	payloadBytes, err := json.Marshal(payload)
	fmt.Println(payload)
	fmt.Println(payloadBytes)

	body := bytes.NewBuffer(payloadBytes) //converts payloadBytes into buffer so can be used as io.reader
	//create http request

	req, err := http.NewRequest("PUT", _url, body)
	utils.ReqError(err)
	//set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	utils.RespError(err)
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	// Read and print the response body

	bodybyte, err := io.ReadAll(resp.Body)

	utils.ReadBodyError(err)
	// Print the response status code
	fmt.Println("Response Status:", resp.Status)

	// Convert response body to string
	bodyString := string(bodybyte)
	fmt.Println("Response body:\n" + bodyString)
}

func gettingBoardLists(api_key string, api_token string, board_id string) {

	xURL := "https://api.trello.com/1/boards/" + board_id + "/lists?key=" + api_key + "&token=" + api_token

	req, err := http.NewRequest("GET", xURL, nil) //calling api, getting board info

	utils.ReqError(err)
	//set headers

	req.Header.Add("Accept", "application/json")

	// Send the HTTP request
	client := &http.Client{} //memory address so we can reuse the the http.client instance for multiple request
	resp, err := client.Do(req)
	utils.RespError(err)
	defer resp.Body.Close()

	bodybyte, err := io.ReadAll(resp.Body)

	utils.ReadBodyError(err)
	// Print the response status code
	fmt.Println("Response Status:", resp.Status)

	// Convert response body to string
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, bodybyte, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println("API Response as String:\n" + prettyJSON.String())

}

func creatingList(api_key string, api_token string, board_id string) {
	//url
	newListName := "oreo"
	_url := "https://api.trello.com/1/boards/" + board_id + "/lists?&key=" + api_key + "&token=" + api_token
	//create json payload
	payload := map[string]string{"name": newListName}
	//convert payload into byte and marshal it
	dataByte, err := json.Marshal(payload)
	utils.JsonMarshalError(err)

	//make http request
	req, err := http.NewRequest("POST", _url, bytes.NewBuffer(dataByte))
	utils.ReqError(err)
	//set header

	req.Header.Add("Content-Type", "application/json")

	//send http request
	client := &http.Client{}
	resp, err := client.Do(req)

	utils.RespError(err)
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	//Read and print response body

	bodybyte, err := io.ReadAll(resp.Body)
	utils.ReadBodyError(err)

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, bodybyte, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println("API Response as String:\n" + prettyJSON.String())
}
