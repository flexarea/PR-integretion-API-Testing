package trello

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

	//getting boards
	newBoar := structure.LoadEndpoint("members/me/boards?")
	structure.GettingBoard(*newBoar, *config, false)

	myBoardLists := structure.LoadEndpoint("/lists?")
	structure.GettingBoardLists(*myBoardLists, *config, false)
	// myList := structure.LoadEndpoint("?")
	structure.DeleteCard(*config, "6673f79f91ffeff4759232f1", false)
	myCard := structure.LoadEndpoint("/actions?")
	structure.GettingCardAction(*myCard, *config, false)
	structure.GettingCardInfo(*config, "66754f33929b363c060ae8a7", false)
	structure.MoveCardtoList(*config, "66754f3b97ef04adf644c5ec", "666c33ee225a312b369c2ad6", true)

}
