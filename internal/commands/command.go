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

type CommandConfig interface {
	Validate() error
}

type CliCommand struct {
	name        string
	description string
	Callback    func() error
}

type CommandHandler interface {
	GetCommand(command string) (CliCommand, error)
}

type CLICommandHandler struct {
	commands map[string]CliCommand
}

func (handler *CLICommandHandler) GetCommand(command string) (CliCommand, error) {
	commandHandle, ok := handler.commands[command]
	if !ok {
		return CliCommand{}, fmt.Errorf("Invalid command")
	}
	return commandHandle, nil
}

func (handler *CLICommandHandler) AddCommandHandler(name, description string, handlerFunc func() error) {
	handler.commands[name] = CliCommand{
		name,
		description,
		handlerFunc,
	}
}

func NewCommandHandler() CommandHandler {

	handler := CLICommandHandler{
		commands: map[string]CliCommand{},
	}
	commandMapConfig := CommandMapConfig{
		Next: api.FirstLocationPage,
		Prev: "",
	}
	handler.AddCommandHandler(CommandHelp, "Displays this help text", func() error {
		return CommandHelpHandler(&HelpCommandConfig{
			commands: handler.commands,
		})
	})
	handler.AddCommandHandler(CommandExit, "Exits the program", CommandExitHandler)
	handler.AddCommandHandler(CommandMap, "Shows next locations on the map", func() error {
		err := CommandMapHandler(&commandMapConfig, FetchDirectionForward)
		return err
	})
	handler.AddCommandHandler(CommandMapBack, "Shows previous locations on the map", func() error {
		err := CommandMapHandler(&commandMapConfig, FetchDirectionBackward)
		return err
	})

	return &handler
}
