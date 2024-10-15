package pokeapi

import (
	"maps"
	"slices"
)

type Pokedex struct {
	entries map[string]PokemonInfo
}

func NewPokedex() *Pokedex {
	entries := make(map[string]PokemonInfo)
	return &Pokedex{
		entries: entries,
	}
}

func (p *Pokedex) AddToPokedex(pokemonInfo PokemonInfo) {
	p.entries[pokemonInfo.Name] = pokemonInfo
}

func (p *Pokedex) GetInfo(pokemon string) (pokemonInfo PokemonInfo, exists bool) {
	info, exists := p.entries[pokemon]
	if !exists {
		return PokemonInfo{}, false
	}
	return info, true
}

func (p *Pokedex) GetPokemons() []string {
	return slices.Collect(maps.Keys(p.entries))

}
