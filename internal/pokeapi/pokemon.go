package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/JP-Go/pokedex-go/internal/cache"
)

var PokemonNotFoundError = errors.New("Pokemon not found")

func FetchPokemonInfo(pokemon string, cache *cache.Cache) (PokemonInfo, error) {

	entry, exists := cache.Get(pokemon)
	var pokemonResponse PokemonInfo

	if err := json.Unmarshal(entry.Val, &pokemonResponse); err != nil && exists {
		return pokemonResponse, nil
	}

	url := BaseURL + "/pokemon/" + pokemon
	res, err := http.Get(url)
	if err != nil {
		return PokemonInfo{}, err
	}

	defer res.Body.Close()
	if res.StatusCode == http.StatusNotFound {
		return PokemonInfo{}, PokemonNotFoundError
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonInfo{}, err
	}
	cache.Add(pokemon, body)
	if err = json.Unmarshal(body, &pokemonResponse); err != nil {
		return PokemonInfo{}, err
	}

	return pokemonResponse, nil
}
