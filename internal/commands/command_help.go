package commands

import "fmt"

type commandList = map[string]cliCommand

func commandHelpHandler(commands commandList) error {

	if commands == nil {
		return fmt.Errorf("Invalid command config. Missing commands")
	}

	fmt.Printf("Welcome to pokedex CLI\nUsage:\n\n")

	for _, c := range commands {
		fmt.Printf("%v: %v \n", c.name, c.description)
	}
	fmt.Println()
	return nil
}

func createHelpHandler(commands commandList) commandCallback {
	return func(_args ...string) error {
		return commandHelpHandler(commands)
	}
}
