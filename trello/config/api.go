package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"trello/utils"
)

func NewRequest(method, url string, payload interface{}) (*http.Request, error) {
	var body io.Reader
	if payload != nil {
		databyte, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(databyte)
	}

	req, err := NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// handle client response
func ClientResponse(req http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(&req)
}

// parse response data
func ParseResponse(resp http.Response) string {
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	databyte, err := io.ReadAll(resp.Body)
	utils.ReadBodyError(err)

	//format json body data
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, databyte, "", " ")
	if err != nil {
		log.Fatal("Error formatting JSON:", err)
	}
	return prettyJSON.String()
}
