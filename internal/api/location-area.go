package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/JP-Go/pokedex-go/internal/cache"
)

const FirstLocationPage = "/location-area/"

var PageLimitReached = errors.New("Page limit reached")

func FetchLocationAreas(url string, cache *cache.Cache) (LocationResponse, error) {
	if url == "" {
		return LocationResponse{}, PageLimitReached
	}

	var locationResponse LocationResponse
	cached, exists := cache.Get(url)
	if exists {
		if err := json.NewDecoder(bytes.NewBuffer(cached.Val)).Decode(&locationResponse); err != nil {
			return LocationResponse{}, err
		}
		return locationResponse, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return LocationResponse{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationResponse{}, err
	}
	cache.Add(url, body)
	if err = json.Unmarshal(body, &locationResponse); err != nil {
		return LocationResponse{}, err
	}
	return locationResponse, nil

}
