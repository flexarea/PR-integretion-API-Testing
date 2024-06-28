package utils

import (
	"fmt"
	"log"
)

func ErrLogTest() {
	fmt.Println("printing from ErrLog.go")
}

func ReqError(err error) {
	if err != nil {
		log.Fatal("Error creating request: ", err)
	}
}
func RespError(err error) {
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}
}
func LoadConfigError(err error) {
	if err != nil {
		log.Fatal("Error loading .env: ", err)
	}
}
func JsonMarshalError(err error) {
	if err != nil {
		log.Fatal("Error marshalling JSON: ", err)
	}
}
func ReadBodyError(err error) {
	if err != nil {
		log.Fatal("Error reading body: ", err)
	}
}
