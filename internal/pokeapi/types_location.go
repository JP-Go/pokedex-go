package pokeapi

type LocationResponse struct {
	Count    int            `json:"count"`
	Next     any            `json:"next"`
	Previous any            `json:"previous"`
	Results  []LocationArea `json:"results"`
}
type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type LocationAreaInfo struct {
	ID                int                 `json:"id"`
	Name              string              `json:"name"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type EncounterDetails struct {
	MinLevel        int   `json:"min_level"`
	MaxLevel        int   `json:"max_level"`
	ConditionValues []any `json:"condition_values"`
	Chance          int   `json:"chance"`
}
type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}
