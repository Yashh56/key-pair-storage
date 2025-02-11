package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Yashh56/keyValueStore/internal/store"
)

func Cmd() {
	store := store.NewKeyValueStore(100)
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
			key, value := args[1], args[2]
			ttl := 0

			if len(args) >= 4 {
				var err error
				ttl, err = strconv.Atoi(args[3])
				if err != nil {
					fmt.Println("Error: Invalid TTL value, must be an integer")
					return
				}
			}
			store.SetKeyValue(key, value, ttl)
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
