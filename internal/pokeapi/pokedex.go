package pokeapi

import "fmt"

type Pokedex struct {
	entries map[string]PokemonInfo
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		entries: map[string]PokemonInfo{},
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

func (p *Pokedex) GetPokemons() {
	for pokemon := range p.entries {
		fmt.Println(" - " + pokemon)
	}
}
