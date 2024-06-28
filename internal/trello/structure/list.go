package structure

import (
	"fmt"
	"trello/config"
	"trello/utils"
)

// getting a single list information
func GettingListInfo(b Board, configuration config.Configs, flag bool) {
	url := configuration.MAIN_END_POINT + "lists/" + configuration.LIST_ID + b.Endpoint + fmt.Sprintf("key=%s&token=%s", configuration.API_KEY, configuration.API_TOKEN)
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
func GettingListCards(b Board, configuration config.Configs, flag bool) {
	url := configuration.MAIN_END_POINT + "lists/" + configuration.LIST_ID + b.Endpoint + fmt.Sprintf("key=%s&token=%s", configuration.API_KEY, configuration.API_TOKEN)
	fmt.Println(url)
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
