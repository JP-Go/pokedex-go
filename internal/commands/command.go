package commands

import (
	"fmt"
	"time"

	"github.com/JP-Go/pokedex-go/internal/cache"
)

const (
	CommandHelp    = "help"
	CommandExit    = "exit"
	CommandMap     = "map"
	CommandMapBack = "mapb"
)

type commandCallback = func(arguments ...string) error

type CliConfig struct {
	next     string
	previous string
	cache    *cache.Cache
}

type cliCommand struct {
	name        string
	description string
	Callback    commandCallback
}

type CommandHandler interface {
	GetCommand(command string) (cliCommand, error)
}

func NewCliConfig(cacheCfg *cache.Cache) CliConfig {
	if cacheCfg != nil {
		return CliConfig{
			cache: cacheCfg,
		}
	}
	return CliConfig{
		cache: cache.NewCache(20 * time.Second),
	}
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

func (handler *CLICommandHandler) AddCommandHandler(name, description string, callback commandCallback) {
	handler.commands[name] = cliCommand{
		name,
		description,
		callback,
	}
}

func NewCommandHandler(config *CliConfig) CommandHandler {

	handler := CLICommandHandler{
		commands: map[string]cliCommand{},
	}
	handler.AddCommandHandler(CommandHelp, "Displays this help text", createHelpHandler(handler.commands))
	handler.AddCommandHandler(CommandExit, "Exits the program", createExitHandler())
	handler.AddCommandHandler(CommandMap, "Shows next locations on the map", createMapHandler(config))
	handler.AddCommandHandler(CommandMapBack, "Shows previous locations on the map", createMapBHandler(config))

	return &handler
}
