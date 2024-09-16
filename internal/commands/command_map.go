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
var ErrorLastPage = errors.New("Already on the last page")

type CommandMapConfig struct {
	Next string
	Prev string
}

func (c *CommandMapConfig) Validate() error {
	if c.Next == "" && c.Prev == "" {
		return errors.New("Invalid config for map command")
	}

	return nil
}

func CommandMapHandler(config *CommandMapConfig, cursorDirection cursorDirection) error {
	err := config.Validate()
	if err != nil {
		return err
	}
	locationRes := api.LocationResponse{}
	if cursorDirection == FetchDirectionForward {
		locationRes, err = api.FetchLocationAreas(config.Next)
		if errors.Is(err, api.PageLimitReached) {
			return ErrorLastPage
		}
	} else {
		locationRes, err = api.FetchLocationAreas(config.Prev)
		if errors.Is(err, api.PageLimitReached) {
			return ErrorFirstPage
		}
	}
	if err != nil {
		return err
	}
	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}

	nextURL, _ := locationRes.Next.(string)
	prevURL, _ := locationRes.Previous.(string)
	config.Next = nextURL
	config.Prev = prevURL
	return nil
}
