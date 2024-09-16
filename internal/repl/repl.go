package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JP-Go/pokedex-go/internal/commands"
)

func StartRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	commandHandler := commands.NewCommandHandler()
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := processInput(scanner.Text())
			if len(input) == 0 {
				fmt.Println("No command")
				continue
			}
			command := input[0]
			handler, err := commandHandler.GetCommand(command)
			if err != nil {
				fmt.Printf("Invalid command: %v \n", command)
				continue
			}
			err = handler.Callback()
			if err != nil {
				fmt.Printf("Error with command %s: %v\n", command, err)
			}
		}
	}
}

func processInput(input string) []string {
	lowerInput := strings.ToLower(strings.TrimSpace(input))
	commandAndArgs := strings.Fields(lowerInput)
	return commandAndArgs
}
