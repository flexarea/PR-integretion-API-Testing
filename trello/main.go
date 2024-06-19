package main

import (
	"fmt"
	"trello/config"
	"trello/structure"
	"trello/utils"
)

//https://api.trello.com/1/actions/?key=APIKey&token=APIToken
//https://api.trello.com/1/boards/{idBoard}?key={yourKey}&token={yourToken}'

func main() {
	fmt.Println("Calling trello api")
	//retrieve data from .env
	config, err := config.Load_config()
	utils.LoadConfigError(err)
	newBoar := structure.LoadEndpoint("members/me/boards?")
	structure.GettingBoard(*newBoar, *config)
}
