package structure

import (
	"fmt"
	"trello/config"
	"trello/utils"
)

// /"members/me/boards?"
func GettingBoard(b Board, configuration config.Configs, flag bool) {
	url := configuration.MAIN_END_POINT + b.Endpoint + fmt.Sprintf("key=%s&token=%s", configuration.API_KEY, configuration.API_TOKEN)
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
func GettingBoardLists(b Board, configuration config.Configs, flag bool) {
	url := configuration.MAIN_END_POINT + "boards/" + configuration.BOARD_ID + b.Endpoint + fmt.Sprintf("key=%s&token=%s", configuration.API_KEY, configuration.API_TOKEN)
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
