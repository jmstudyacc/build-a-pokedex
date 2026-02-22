package main

import (
	"encoding/json"
	"fmt"

	"pokedexcli/internal/pokecache"
)

func explore(args []string, c *pokecache.Cache) error {
	var encounters pokemonEncounter

	if len(args) == 0 {
		return fmt.Errorf("explore requires a location area name")
	}
	url := pokeAPILocations + args[0]
	data, err := fetchData(url, c)
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	// UNMARSHAL THE DATA INTO POKEMON ENCOUNTER TOP LEVEL STRUCT
	if err := json.Unmarshal(data, &encounters); err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	// DISPLAY RESULTING POKEMON ENCOUNTERS
	for _, e := range encounters.Encounter {
		fmt.Printf("- %s\n", e.Pokemon.Name)
	}
	return nil
}
