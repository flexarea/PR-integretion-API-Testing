package main

import (
	"fmt"
	"github.com/flexarea/PR-integration-API-Testing/internal/config"
	"github.com/flexarea/PR-integration-API-Testing/internal/slack"
)

func main() {
	fmt.Println("Testing from cmd")

	//Testing slack api
	config.Load_config()

}
