package commands

import (
	"fmt"
	"time"

	"github.com/JP-Go/pokedex-go/internal/cache"
	"github.com/JP-Go/pokedex-go/internal/pokeapi"
)

const (
	CommandHelp    = "help"
	CommandExit    = "exit"
	CommandMap     = "map"
	CommandMapBack = "mapb"
	CommandExplore = "explore"
	CommandCatch   = "catch"
	CommandInspect = "inspect"
	CommandPokedex = "pokedex"
)

type commandCallback = func(arguments ...string) error

type CliConfig struct {
	next            string
	previous        string
	cache           *cache.Cache
	pokedex         *pokeapi.Pokedex
	currentLocation string
}

type cliCommand struct {
	name        string
	usage       string
	description string
	Callback    commandCallback
}

type CommandHandler interface {
	GetCommand(command string) (cliCommand, error)
}

func NewCliConfig(cacheCfg *cache.Cache) CliConfig {
	if cacheCfg != nil {
		return CliConfig{
			cache:   cacheCfg,
			pokedex: pokeapi.NewPokedex(),
		}
	}
	return CliConfig{
		cache:   cache.NewCache(20 * time.Second),
		pokedex: pokeapi.NewPokedex(),
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

func (handler *CLICommandHandler) AddCommandHandler(cliCommand cliCommand) {
	handler.commands[cliCommand.name] = cliCommand
}

func NewCommandHandler(config *CliConfig) CommandHandler {

	handler := CLICommandHandler{
		commands: map[string]cliCommand{},
	}
	handler.AddCommandHandler(
		cliCommand{
			name:        CommandHelp,
			description: "Displays this help text",
			usage:       "help",
			Callback:    createHelpHandler(handler.commands),
		},
	)
	handler.AddCommandHandler(cliCommand{
		name:        CommandExit,
		description: "Exits the program",
		Callback:    createExitHandler(),
		usage:       "exit",
	})
	handler.AddCommandHandler(
		cliCommand{
			name:        CommandMap,
			description: "Shows next locations on the map",
			Callback:    createMapHandler(config),
			usage:       "map",
		},
	)
	handler.AddCommandHandler(
		cliCommand{
			name:        CommandMapBack,
			description: "Shows previous locations on the map",
			Callback:    createMapBHandler(config),
			usage:       "mapb",
		},
	)
	handler.AddCommandHandler(
		cliCommand{
			name:        CommandExplore,
			description: "Displays pokemon encounters available at the location",
			Callback:    createExploreHandler(config),
			usage:       "explore <location>",
		},
	)
	handler.AddCommandHandler(
		cliCommand{
			name:        CommandCatch,
			description: "Tries to catch the pokemon in the current area",
			Callback:    createCatchCommand(config),
			usage:       "catch <pokemon-name>",
		},
	)
	handler.AddCommandHandler(
		cliCommand{
			name:        CommandInspect,
			description: "Inspects a pokemon in your pokedex",
			Callback:    createInspectCommand(config),
			usage:       "inspect <pokemon-name>",
		},
	)
	handler.AddCommandHandler(
		cliCommand{
			name:        CommandPokedex,
			description: "List your pokedex",
			Callback:    createPokedexCommand(config),
			usage:       "pokedex",
		},
	)

	return &handler
}
