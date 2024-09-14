package commands

import "fmt"

func CommandHelpHandler(commands map[string]CliCommand) error {

	fmt.Printf("Welcome to pokedex CLI\nUsage:\n\n")

	for _, c := range commands {
		fmt.Printf("%v: %v \n", c.name, c.description)
	}
	fmt.Println()
	return nil
}
