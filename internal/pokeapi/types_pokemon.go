package pokeapi

type PokemonInfo struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	BaseExperience int         `json:"base_experience"`
	Order          int         `json:"order"`
	Abilities      []Abilities `json:"abilities"`
	Forms          []Forms     `json:"forms"`
	Moves          []Moves     `json:"moves"`
	Stats          []Stats     `json:"stats"`
	Types          []Types     `json:"types"`
}
type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Abilities struct {
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
	Ability  Ability `json:"ability"`
}
type Forms struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type GameIndices struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}
type Item struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Move struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Moves struct {
	Move Move `json:"move"`
}
type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}
type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}
