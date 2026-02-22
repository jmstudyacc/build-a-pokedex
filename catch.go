package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"

	"pokedexcli/internal/pokecache"
)

// Catching pokemon
func catch(args []string, c *pokecache.Cache, p pokedex) error {
	var pokemon pokeInfo
	catchChance := rand.IntN(362)

	if len(args) < 1 {
		return fmt.Errorf("please enter a pokemon name")
	}

	targetPokemon := args[0]

	pokemonEndpointURL := "https://pokeapi.co/api/v2/pokemon/" + targetPokemon

	data, err := fetchData(pokemonEndpointURL, c)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return fmt.Errorf("%w", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", targetPokemon)

	// determine if pokemon can be caught
	// if experience is high, pokemon is harder to catch
	//	cc, b := c.Get(pokemon)
	// fmt.Println(catchChance)
	// fmt.Printf("DEBUG: Pokemon base experience: %d\n", pokemon.BaseExperience)

	if catchChance >= pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", targetPokemon)

		// adding caught pokemon to pokedex
		p[targetPokemon] = pokemon
		return nil
	}

	fmt.Printf("%s escaped!\n", targetPokemon)

	return nil
}
