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

type commandConfig interface {
	Validate() error
}

type CliConfig struct {
	next     string
	previous string
	cache    *cache.Cache
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

type cliCommand struct {
	name        string
	description string
	Callback    func(arguments ...string) error
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

func (handler *CLICommandHandler) AddCommandHandler(name, description string, handlerFunc func(...string) error) {
	handler.commands[name] = cliCommand{
		name,
		description,
		handlerFunc,
	}
}

func NewCommandHandler(cfg *CliConfig) CommandHandler {

	handler := CLICommandHandler{
		commands: map[string]cliCommand{},
	}
	handler.AddCommandHandler(CommandHelp, "Displays this help text", func(_args ...string) error {
		return CommandHelpHandler(&HelpCommandConfig{
			commands: handler.commands,
		})
	})
	handler.AddCommandHandler(CommandExit, "Exits the program", CommandExitHandler)
	handler.AddCommandHandler(CommandMap, "Shows next locations on the map", func(_args ...string) error {
		err := CommandMapHandler(cfg)
		return err
	})
	handler.AddCommandHandler(CommandMapBack, "Shows previous locations on the map", func(_args ...string) error {
		err := CommandMapBHandler(cfg)
		return err
	})

	return &handler
}
