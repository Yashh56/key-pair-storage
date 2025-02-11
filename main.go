package main

import (
	"fmt"
	"os"

	"github.com/Yashh56/keyValueStore/api"
	"github.com/Yashh56/keyValueStore/cli"
)

func main() {
	command := os.Args

	if len(command) > 1 && command[1] == "api" {
		api.Server()
	}
	if len(command) > 1 && command[1] == "cli" {
		cli.Cmd()
	} else {
		fmt.Println("Something went wrong")
	}

}
