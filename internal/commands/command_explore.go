package commands

import (
	"errors"
	"fmt"

	"github.com/JP-Go/pokedex-go/internal/pokeapi"
)

func createExploreHandler(cliConfig *CliConfig) commandCallback {
	return func(arguments ...string) error {
		return commandExploreHandler(cliConfig, arguments...)
	}
}

func commandExploreHandler(cliConfig *CliConfig, arguments ...string) error {
	if len(arguments) < 1 {
		return errors.New("You must provide a location")
	}
	location := arguments[0]
	locationInfo, err := pokeapi.FetchLocationArea(location, cliConfig.cache)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Exploring %s", locationInfo.Name))
	fmt.Println("Found pokemon")
	for _, pokemon := range locationInfo.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
