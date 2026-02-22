package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"pokedexcli/internal/pokecache"
)

var ErrNotFound = fmt.Errorf("resource not found")

// creating a helper function for DRY
func fetchLocationData(url string, c *pokecache.Cache) (pokeAPI, error) {
	// generate pokeAPI struct to return
	var resp pokeAPI

	cacheData, ok := c.Get(url)
	if ok {
		if err := json.Unmarshal(cacheData, &resp); err != nil {
			return pokeAPI{}, fmt.Errorf("%w", err)
		}
		return resp, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return pokeAPI{}, fmt.Errorf("error accessing PokeAPI for location list: %w", err)
	}
	defer res.Body.Close()

	// get data from HTTP request
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokeAPI{}, fmt.Errorf("ERROR: %w", err)
	}

	// Add entry to Cache
	c.Add(url, data)

	// UNMARSHAL HTTP DATA
	if err := json.Unmarshal(data, &resp); err != nil {
		return pokeAPI{}, fmt.Errorf("ERROR - JSON Unmarshal Error: %w", err)
	}
	return resp, nil
}

func fetchData(url string, c *pokecache.Cache) ([]byte, error) {
	// CHECK Cache
	cacheData, ok := c.Get(url)
	if ok {
		return cacheData, nil
	}
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ERROR: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode == 404 {
		return nil, ErrNotFound
	}

	// GET DATA FROM Body
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("ERROR: %w", err)
	}
	// ADD TO CACHE
	c.Add(url, data)

	// RETURN DATA
	return data, nil
}
