package api

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
