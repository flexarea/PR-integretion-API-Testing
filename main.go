package main

import (
	"fmt"
	"github.com/flexarea/PR-integration-API-Testing/internal/config"
)

func main() {
	fmt.Println("Testing from cmd")

	//Testing load.env
	config.Test()
}
