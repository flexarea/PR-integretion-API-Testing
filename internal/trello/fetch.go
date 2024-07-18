package trello

import (
	"encoding/json"
	"fmt"
	"log"
)

func Retrieve(data string) {

	var jsonData = []byte(data)

	var result []map[string]interface{}

	err := json.Unmarshal(jsonData, result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%T", result)

}
