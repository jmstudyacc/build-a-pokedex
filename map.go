package main

import (
	"fmt"

	"pokedexcli/internal/pokecache"
)

func mapForward(la *locationArea, c *pokecache.Cache) error {
	// set local value for url to constant for PokeAPI
	url := pokeAPILocations

	if la.next != "" {
		url = la.next // call to PokeAPI has occurred, value will be present in internal struct for next map locations
	}
	// begin HTTP request
	resp, err := fetchLocationData(url, c)
	if err != nil {
		return fmt.Errorf("error accessing PokeAPI for Locations list: %w", err)
	}

	// DISPLAY RESULTS
	// range over the Results[] in resp (pokeAPI)
	for _, item := range resp.Results {
		fmt.Println(item.Name)
	}

	// UPDATE STATE
	// now API call is made, we can set the internal struct 'next' & 'previous' fields to match the response stored in PokeAPI
	la.next = resp.Next
	la.previous = resp.Previous

	return nil
}

func mapBack(la *locationArea, c *pokecache.Cache) error {
	// if c.previous is empty - print a message and return
	if la.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	url := la.previous // call to PokeAPI has occurred, value will be present in internal struct for next map locations

	resp, err := fetchLocationData(url, c)
	if err != nil {
		return fmt.Errorf("error accessing PokeAPI URL for Locations list: %w", err)
	}

	// DISPLAY RESULTS
	for _, item := range resp.Results {
		fmt.Println(item.Name)
	}

	// UPDATE STATE
	la.previous = resp.Previous
	la.next = resp.Next

	return nil
}
