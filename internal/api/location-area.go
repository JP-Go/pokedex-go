package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

const FirstLocationPage = "https://pokeapi.co/api/v2/location-area/"

var PageLimitReached = errors.New("Page limit reached")

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

func FetchLocationAreas(url string) (LocationResponse, error) {

	if url == "" {
		return LocationResponse{}, PageLimitReached
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationResponse{}, err
	}
	defer res.Body.Close()
	var locationResponse LocationResponse
	if err = json.NewDecoder(res.Body).Decode(&locationResponse); err != nil {
		return LocationResponse{}, err
	}
	return locationResponse, nil
}
