package commands

import (
	"fmt"

	"github.com/JP-Go/pokedex-go/internal/api"
)

const (
	CommandHelp    = "help"
	CommandExit    = "exit"
	CommandMap     = "map"
	CommandMapBack = "mapb"
)

type commandConfig interface {
	Validate() error
}

type CliConfig struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	Callback    func() error
}

type CommandHandler interface {
	GetCommand(command string) (cliCommand, error)
}

type CLICommandHandler struct {
	commands map[string]cliCommand
}

func (handler *CLICommandHandler) GetCommand(command string) (cliCommand, error) {
	commandHandle, ok := handler.commands[command]
	if !ok {
		return cliCommand{}, fmt.Errorf("Invalid command")
	}
	return commandHandle, nil
}

func (handler *CLICommandHandler) AddCommandHandler(name, description string, handlerFunc func() error) {
	handler.commands[name] = cliCommand{
		name,
		description,
		handlerFunc,
	}
}

func NewCliConfig() *CliConfig {

	return &CliConfig{
		next: api.BaseURL + api.FirstLocationPage,
	}
}

func NewCommandHandler(cfg *CliConfig) CommandHandler {

	handler := CLICommandHandler{
		commands: map[string]cliCommand{},
	}
	handler.AddCommandHandler(CommandHelp, "Displays this help text", func() error {
		return CommandHelpHandler(&HelpCommandConfig{
			commands: handler.commands,
		})
	})
	handler.AddCommandHandler(CommandExit, "Exits the program", CommandExitHandler)
	handler.AddCommandHandler(CommandMap, "Shows next locations on the map", func() error {
		err := CommandMapHandler(cfg)
		return err
	})
	handler.AddCommandHandler(CommandMapBack, "Shows previous locations on the map", func() error {
		err := CommandMapBHandler(cfg)
		return err
	})

	return &handler
}
