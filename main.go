package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JP-Go/pokedex-go/internal"
	"github.com/JP-Go/pokedex-go/internal/commands"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	commandHandler := commands.NewCommandHandler()
	for {
		internal.PrintPrompt()
		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())
			handler, err := commandHandler.GetCommand(input)
			if err != nil {
				fmt.Println("Invalid command " + input)
			}
			err = handler.Callback()
			if err != nil {
				panic(fmt.Sprintf("Error with command %s: %v", input, err))
			}
		}
	}
}
