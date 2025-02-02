package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Yashh56/keyValueStore/cmd/api"
	"github.com/Yashh56/keyValueStore/cmd/internal/store"
)

func Cmd() {
	store := store.NewKeyValueStore()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Simple Key-Value Store CLI")
	fmt.Println("Commands:")
	fmt.Println("  store <key> <value>  - Store a key-value pair")
	fmt.Println("  get <key>            - Retrieve a value")
	fmt.Println("  delete <key>         - Delete a key")
	fmt.Println("  exit                 - Quit the CLI")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		command := args[0]

		switch command {
		case "store":
			if len(args) < 3 {
				fmt.Println("Usage: store <key> <value>")
				continue
			}
			key, value := args[1], strings.Join(args[2:], " ")
			store.SetKeyValue(key, value)
			fmt.Printf("Stored: %s = %s\n", key, value)

		case "get":
			if len(args) != 2 {
				fmt.Println("Usage: get <key>")
				continue
			}
			key := args[1]
			value, found := store.GetKeyValue(key)
			if found {
				fmt.Printf("Value: %s\n", value)
			} else {
				fmt.Println("Key not found")
			}

		case "delete":
			if len(args) != 2 {
				fmt.Println("Usage: delete <key>")
				continue
			}
			key := args[1]
			if store.DeleteKeyValue(key) {
				fmt.Println("Deleted successfully")
			} else {
				fmt.Println("Key not found")
			}

		case "exit":
			fmt.Println("Exiting CLI...")
			return

		default:
			fmt.Println("Unknown command. Available commands: store, get, delete, exit")
		}
	}

}

func main() {
	command := os.Args

	if len(command) > 1 && command[1] == "api" {
		api.Server()
	}
	if len(command) > 1 && command[1] == "cli" {
		Cmd()
	} else {
		fmt.Println("Something went wrong")
	}

}
