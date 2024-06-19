package structure

import (
	"fmt"
	"trello/config"
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
func GettingBoard(b Board, config config.Configs) {
	url := config.MAIN_END_POINT + b.Endpoint + fmt.Sprintf("key=%s&token=%s", config.API_KEY, config.API_TOKEN)
	fmt.Println(url)
}
