package structure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"trello/config"
	"trello/utils"
)

// getting a single list information
func GettingCardInfo(configuration config.Configs, cardID string, flag bool) {
	if !flag {
		return
	}
	url := fmt.Sprintf("%scards/%s?key=%s&token=%s", configuration.MAIN_END_POINT, cardID, configuration.API_KEY, configuration.API_TOKEN)

	//make new request
	req, err := http.NewRequest("GET", url, nil)
	utils.ReqError(err)
	resp, err := http.DefaultClient.Do(req)
	utils.RespError(err)
	databyte, err := io.ReadAll(resp.Body)
	utils.ReadBodyError(err)

	//format json body data
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, databyte, "", " ")
	if err != nil {
		log.Fatal("Error formatting JSON: ", err)
	}
	fmt.Println(prettyJSON.String())
}

// getting all cards from a single list
func GettingCardAction(b Board, configuration config.Configs, flag bool) {
	url := configuration.MAIN_END_POINT + "cards/" + configuration.CARD_ID + b.Endpoint + fmt.Sprintf("key=%s&token=%s", configuration.API_KEY, configuration.API_TOKEN)
	req, err := config.NewRequest("GET", url, nil)
	utils.ReqError(err)
	resp, err := config.ClientResponse(req)
	utils.RespError(err)

	if !flag {
		fmt.Println("")
	} else {
		fmt.Println(config.ParseResponse(resp))
	}
}
func DeleteCard(configuration config.Configs, cardID string, flag bool) error {
	//code modified by CTO
	if !flag {
		return nil
	}

	url := fmt.Sprintf("%scards/%s?key=%s&token=%s", configuration.MAIN_END_POINT, cardID, configuration.API_KEY, configuration.API_TOKEN)

	req, err := http.NewRequest("DELETE", url, nil)
	utils.ReqError(err)

	res, err := http.DefaultClient.Do(req)
	utils.ReqError(err)
	p, err := io.ReadAll(res.Body)
	utils.ReadBodyError(err)

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
