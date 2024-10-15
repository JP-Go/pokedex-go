package commands

import (
	"errors"
	"fmt"

	"github.com/JP-Go/pokedex-go/internal/pokeapi"
	"math/rand"
)

func createCatchCommand(cliConfig *CliConfig) commandCallback {
	pokedex := pokeapi.NewPokedex()

	return func(arguments ...string) error {
		return commandCatchHandler(cliConfig, pokedex, arguments...)
	}
}

func commandCatchHandler(cliconfig *CliConfig, pokedex *pokeapi.Pokedex, arguments ...string) error {

	if len(arguments) < 1 {
		return errors.New("You must provide a pokemon")
	}
	pokemon := arguments[0]
	locationInfo, err := pokeapi.FetchLocationArea(cliconfig.currentLocation, cliconfig.cache)
	if err != nil {
		return err
	}

	if !canCatchPokemon(locationInfo, pokemon) {
		return errors.New(fmt.Sprintf(
			"Pokemon not found in the current area of %s\n",
			cliconfig.currentLocation,
		))
	}
	pokemonInfo, err := pokeapi.FetchPokemonInfo(pokemon, cliconfig.cache)
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
	pokedex.AddToPokedex(pokemonInfo)
	pokedex.GetPokemons()

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
