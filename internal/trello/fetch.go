package trello

import (
	"encoding/json"
	"log"
)

func Retrieve(cardShortID interface{}, data string) string {

	var jsonData = []byte(data)

	var result []map[string]interface{}

	err := json.Unmarshal(jsonData, &result)

	if err != nil {
		log.Fatal(err)
	}

	for _, jsonObject := range result {
		cardId := jsonObject["id"].(string)
		idShort, ok := jsonObject["idShort"].(int)

		if !ok {
			return ""
		} else {
			if idShort == cardShortID.(int) {
				return cardId
			}
		}
	}

	return ""
}
