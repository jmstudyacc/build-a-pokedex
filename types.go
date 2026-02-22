package main

// Struct created to STORE data FROM the PokeAPI
type pokeAPI struct {
	// NOTE: to export JSON fields you need to capitalise the Fields
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []struct { // anonymous Struct to store the names of the locations on the map
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// internal Struct to store the next & previous values from the PokeAPI (Next & Previous respectively)
type locationArea struct {
	next     string
	previous string
}

type pokemonEncounter struct {
	Encounter []encounter `json:"pokemon_encounters"`
}

type encounter struct {
	Pokemon pokemon `json:"pokemon"`
}

type pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type pokeInfo struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	BaseExperience int         `json:"base_experience"`
	Height         int         `json:"height"`
	Weight         int         `json:"weight"`
	Stats          []statEntry `json:"stats"`
	Type           []types     `json:"types"`
	// seen & caught will be needed
}

type statEntry struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"stat"`
}

type types struct {
	Slot     int `json:"slot"`
	TypeInfo struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}

type pokedex map[string]pokeInfo
