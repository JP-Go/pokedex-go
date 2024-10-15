package commands

import (
	"errors"
	"fmt"

	"github.com/JP-Go/pokedex-go/internal/pokeapi"
	"math/rand"
)

func createCatchCommand(cliConfig *CliConfig) commandCallback {
	return func(arguments ...string) error {
		return commandCatchHandler(cliConfig, arguments...)
	}
}

func commandCatchHandler(cliConfig *CliConfig, arguments ...string) error {

	if len(arguments) < 1 {
		return errors.New("You must provide a pokemon")
	}
	pokemon := arguments[0]
	locationInfo, err := pokeapi.FetchLocationArea(cliConfig.currentLocation, cliConfig.cache)
	if err != nil {
		return err
	}

	if !canCatchPokemon(locationInfo, pokemon) {
		return errors.New(fmt.Sprintf(
			"Pokemon not found in the current area of %s\n",
			cliConfig.currentLocation,
		))
	}
	pokemonInfo, err := pokeapi.FetchPokemonInfo(pokemon, cliConfig.cache)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a pokeball at %s...\n", pokemonInfo.Name)
	chance := rand.Intn(pokemonInfo.BaseExperience)
	if chance < pokemonInfo.BaseExperience/2 {
		fmt.Printf("%s escaped!\n", pokemonInfo.Name)
		return nil
	}
	fmt.Printf("%s was caugth!\n", pokemonInfo.Name)
	fmt.Println("You can now inspect it with the inspect command.")
	cliConfig.pokedex.AddToPokedex(pokemonInfo)

	return nil
}

func canCatchPokemon(locationInfo pokeapi.LocationAreaInfo, pokemon string) bool {
	isPokemonFoundInCurrentLocation := false
	for _, encounter := range locationInfo.PokemonEncounters {
		if encounter.Pokemon.Name == pokemon {
			isPokemonFoundInCurrentLocation = true
		}
	}
	return isPokemonFoundInCurrentLocation
}
