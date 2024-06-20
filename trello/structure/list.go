package structure

import (
	"fmt"
	"trello/config"
	"trello/utils"
)

// /"members/me/boards?"
func GettingList(b Board, configuration config.Configs, flag bool) {
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
