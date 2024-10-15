package commands

import (
	"fmt"

	"github.com/JP-Go/pokedex-go/internal/pokeapi"
)

func createInspectCommand(cliConfig *CliConfig) commandCallback {
	return func(arguments ...string) error {
		return commandInspectHandler(cliConfig, arguments...)
	}
}

func commandInspectHandler(cliConfig *CliConfig, arguments ...string) error {
	if len(arguments) < 1 {
		fmt.Println("You must provide at least one pokemon")
		return nil
	}
	for _, pokemonName := range arguments {
		pokemon, wasCaught := cliConfig.pokedex.GetInfo(pokemonName)
		if !wasCaught {
			fmt.Println("You have not caught " + pokemonName)
		} else {
			showPokemonInfo(pokemon)
		}
	}
	return nil
}

func showPokemonInfo(pokemonInfo pokeapi.PokemonInfo) {
	fmt.Printf("Name: %s\n", pokemonInfo.Name)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemonInfo.Types {
		fmt.Printf("  -%s\n", pokemonType.Type.Name)
	}
}
