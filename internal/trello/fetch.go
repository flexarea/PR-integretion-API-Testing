package trello

import (
	"encoding/json"
	"log"
)

func Retrieve(cardShortID interface{}, data string) {

	var jsonData = []byte(data)

	var result []map[string]interface{}

	err := json.Unmarshal(jsonData, &result)

	if err != nil {
		log.Fatal(err)
	}

}
