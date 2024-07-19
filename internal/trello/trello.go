package trello

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/flexarea/PR-integration-API-Testing/configs"
)

//https://api.trello.com/1/actions/?key=APIKey&token=APIToken
//https://api.trello.com/1/boards/{idBoard}?key={yourKey}&token={yourToken}'

// getting a single list information
func GettingCardsInList(configuration configs.Configs, listID string) (*string, error) {

	url := fmt.Sprintf("%slists/%s/cards?key=%s&token=%s", configuration.TRELLO_MAIN_END_POINT, listID, configuration.API_KEY, configuration.API_TOKEN)

	//make new request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	databyte, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	//format json body data
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, databyte, "", " ")
	if err != nil {
		log.Fatal("Error formatting JSON: ", err)
	}
	str := prettyJSON.String()
	return &str, nil
}

// getting all cards from a single list
// ** To update
func DeleteCard(configuration configs.Configs, cardID string, flag bool) error {
	//code modified by CTO
	if !flag {
		return nil
	}

	url := fmt.Sprintf("%scards/%s?key=%s&token=%s", configuration.TRELLO_MAIN_END_POINT, cardID, configuration.API_KEY, configuration.API_TOKEN)

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	p, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	response := struct {
		Limits map[string]interface{} `json:limits`
	}{}

	defer res.Body.Close()

	err = json.Unmarshal(p, &response)
	if err != nil {
		return err
	}

	fmt.Println(response)

	return nil
}

func MoveCardtoList(configuration configs.Configs, cardID string, targetListId string) error {
	url := fmt.Sprintf("%scards/%s?key=%s&token=%s", configuration.TRELLO_MAIN_END_POINT, cardID, configuration.API_KEY, configuration.API_TOKEN)
	//create json payload
	payload := map[string]string{"idList": targetListId}
	//make new request

	dataByte, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(dataByte))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	//send http request
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
