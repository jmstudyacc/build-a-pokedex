package main

import (
	"fmt"

	"pokedexcli/internal/pokecache"
)

func inspect(args []string, c *pokecache.Cache, p pokedex) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a pokemon name")
	}

	for i := range args {
		targetPokemon := args[i]

		entry, ok := p[targetPokemon]
		if !ok {
			fmt.Println("\n!!!!!!!!!! POKEDEX ENTRY MISSING !!!!!!!!!!")
			fmt.Printf("you have not caught %s\n", targetPokemon)
			continue
			// return nil
		}

		fmt.Printf("\n########## POKEDEX ENTRY: %d ###########\n", entry.ID)
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
	}
	fmt.Println()
	return nil
}

func pokedexPrint(p pokedex) error {
	fmt.Println("Your Pokedex:")
	for _, entry := range p {
		fmt.Printf("  - %s\n", entry.Name)
	}

	return nil
}
