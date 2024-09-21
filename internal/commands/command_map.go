package commands

import (
	"errors"
	"fmt"

	"github.com/JP-Go/pokedex-go/internal/api"
)

type cursorDirection string

const FetchDirectionForward = "forward"
const FetchDirectionBackward = "backward"

var ErrorFirstPage = errors.New("Already on the first page")

func CommandMapHandler(config *CliConfig) error {
	if config.next == "" {
		config.next = api.BaseURL + api.FirstLocationPage
	}
	locationRes := api.LocationResponse{}
	locationRes, err := api.FetchLocationAreas(config.next, config.cache)
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

func CommandMapBHandler(config *CliConfig) error {
	locationRes := api.LocationResponse{}
	locationRes, err := api.FetchLocationAreas(config.previous, config.cache)
	if errors.Is(err, api.PageLimitReached) {
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
