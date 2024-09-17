package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

const FirstLocationPage = "/location-area/"

var PageLimitReached = errors.New("Page limit reached")

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
