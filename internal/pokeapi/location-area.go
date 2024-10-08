package pokeapi

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
var UnkownLocation = errors.New("Unkown location")

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

func FetchLocationArea(location string, cache *cache.Cache) (LocationAreaInfo, error) {
	url := BaseURL + FirstLocationPage + location
	if data, exists := cache.Get(url); exists {
		var info LocationAreaInfo
		if err := json.Unmarshal(data.Val, &info); err != nil {
			return LocationAreaInfo{}, err
		}
		return info, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaInfo{}, err
	}
	if res.StatusCode == http.StatusNotFound {
		return LocationAreaInfo{}, UnkownLocation
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	var info LocationAreaInfo
	if err = json.Unmarshal(body, &info); err != nil {
		return LocationAreaInfo{}, err
	}
	cache.Add(url, body)
	return info, err
}
