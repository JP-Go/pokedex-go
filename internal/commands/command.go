package commands

import (
	"fmt"
)

const (
	CommandHelp = "help"
	CommandExit = "exit"
)

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
	handler.AddCommandHandler(CommandHelp, "Displays this help text", func() error {
		return CommandHelpHandler(handler.commands)
	})
	handler.AddCommandHandler(CommandExit, "Exits the program", CommandExitHandler)

	return &handler
}
