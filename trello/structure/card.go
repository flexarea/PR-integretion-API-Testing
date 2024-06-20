package structure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"trello/config"
	"trello/utils"
)

// getting a single list information
func GettingCardInfo(b Board, configuration config.Configs, flag bool) {
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

	if !flag {
		return nil
	}

	url := fmt.Sprintf("%scards/%s?key=%s&token=%s", configuration.MAIN_END_POINT, cardID, configuration.API_KEY, configuration.API_TOKEN)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	p, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	response := struct {
		Limits map[string]interface{} `json:limits`
	}{}

	res.Body.Close()

	err = json.Unmarshal(p, &response)
	if err != nil {
		return err
	}

	fmt.Println(response)

	return nil
}
