package commands

import "fmt"

func createPokedexCommand(config *CliConfig) commandCallback {

	return func(_arguments ...string) error {
		return commandPokedexHandler(config)
	}
}

func commandPokedexHandler(config *CliConfig) error {
	pokemons := config.pokedex.GetPokemons()
	fmt.Println("Your pokemons:")
	for _, pokemon := range pokemons {
		fmt.Printf("  -%s\n", pokemon)
	}
	return nil
}
