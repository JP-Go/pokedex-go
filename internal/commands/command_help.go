package commands

import "fmt"

type HelpCommandConfig struct {
	commands map[string]cliCommand
}

func (c *HelpCommandConfig) Validate() error {
	if c.commands == nil {
		return fmt.Errorf("Invalid command config. Missing commands")
	}
	return nil
}

func CommandHelpHandler(config *HelpCommandConfig) error {

	err := config.Validate()
	if err != nil {
		return err
	}

	fmt.Printf("Welcome to pokedex CLI\nUsage:\n\n")

	for _, c := range config.commands {
		fmt.Printf("%v: %v \n", c.name, c.description)
	}
	fmt.Println()
	return nil
}
