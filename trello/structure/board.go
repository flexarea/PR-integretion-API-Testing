package structure

import (
	"fmt"
	"trello/config"
	"trello/utils"
)

type Board struct {
	Endpoint string
}

func LoadEndpoint(endpt string) *Board {
	newBoard := &Board{
		Endpoint: endpt,
	}
	return newBoard
}

// /"members/me/boards?"
func GettingBoard(b Board, configuration config.Configs, flag bool) {
	url := configuration.MAIN_END_POINT + b.Endpoint + fmt.Sprintf("key=%s&token=%s", configuration.API_KEY, configuration.API_TOKEN)
	req, err := config.NewRequest("GET", url, nil)
	utils.ReqError(err)
	resp, err := config.ClientResponse(req)
	utils.RespError(err)

	if flag == false {
		fmt.Println("")
	} else {
		fmt.Println(config.ParseResponse(resp))
	}

}
func GettingBoardLists(b Board, configuration config.Configs) {
	url := configuration.MAIN_END_POINT + "boards/" + configuration.BOARD_ID + b.Endpoint + fmt.Sprintf("key=%s&token=%s", configuration.API_KEY, configuration.API_TOKEN)
	req, err := config.NewRequest("GET", url, nil)
	utils.ReqError(err)
	resp, err := config.ClientResponse(req)
	utils.RespError(err)
	fmt.Println(config.ParseResponse(resp))
}
