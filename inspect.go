package main

import (
	"fmt"

	"pokedexcli/internal/pokecache"
)

func inspect(args []string, c *pokecache.Cache, p pokedex) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a pokemon name")
	}

	targetPokemon := args[0]

	entry, ok := p[targetPokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", entry.Name)
	fmt.Printf("Height: %d\n", entry.Height)
	fmt.Printf("Weight: %d\n", entry.Weight)
	fmt.Println("Stats:")
	for _, stat := range entry.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, p := range entry.Type {
		fmt.Printf("  -%s\n", p.TypeInfo.Name)
	}
	return nil
}
