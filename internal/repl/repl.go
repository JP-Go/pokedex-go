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
