package commands

import (
	"errors"
	"fmt"

	"github.com/JP-Go/pokedex-go/internal/pokeapi"
)

type cursorDirection string

const FetchDirectionForward = "forward"
const FetchDirectionBackward = "backward"

var ErrorFirstPage = errors.New("Already on the first page")

func createMapHandler(config *CliConfig) commandCallback {
	return func(...string) error {
		return commandMapHandler(config)
	}
}

func createMapBHandler(config *CliConfig) commandCallback {
	return func(...string) error {
		return commandMapBHandler(config)
	}
}

func commandMapHandler(config *CliConfig) error {
	if config.next == "" {
		config.next = pokeapi.BaseURL + pokeapi.FirstLocationPage
	}
	locationRes := pokeapi.LocationResponse{}
	locationRes, err := pokeapi.FetchLocationAreas(config.next, config.cache)
	if err != nil {
		return err
	}
	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}

	nextURL, _ := locationRes.Next.(string)
	prevURL, _ := locationRes.Previous.(string)
	config.next = nextURL
	config.previous = prevURL
	return nil
}

func commandMapBHandler(config *CliConfig) error {
	locationRes := pokeapi.LocationResponse{}
	locationRes, err := pokeapi.FetchLocationAreas(config.previous, config.cache)
	if errors.Is(err, pokeapi.PageLimitReached) {
		return ErrorFirstPage
	}
	if err != nil {
		return err
	}
	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}

	nextURL, _ := locationRes.Next.(string)
	prevURL, _ := locationRes.Previous.(string)
	config.next = nextURL
	config.previous = prevURL
	return nil
}
