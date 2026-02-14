package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"pokedexcli/internal/pokecache"
)

// constant created to store the Poke API Locations URL
const pokeAPILocations = "https://pokeapi.co/api/v2/location-area"

type cliCommand struct {
	name        string
	description string
	callback    func(*locationArea, *pokecache.Cache) error // callback is a function that returns an error
}

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

// Hashmap to store the commands
var commandMap map[string]cliCommand

// to avoid circular dependencies, init() used to assign the map before main runs
func init() {
	commandMap = map[string]cliCommand{ // map with keys that are strings, and values that are cliCommand structs
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		}, "exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		}, "map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    mapForward,
		}, "mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    mapBack,
		},
	}
}

func startRepl() {
	// generate internal struct to store next & previous location data
	la := &locationArea{}

	// instantiate the cache
	cache := pokecache.NewCache(5 * time.Minute)

	// create a scanner to read from os.Stdin
	scanner := bufio.NewScanner(os.Stdin) // calling this will block & wait for user input and CR
	for {
		fmt.Print("Pokedex > ")

		// scanner.Scan() required to start populating a buffer with text
		scanner.Scan()

		// generating a string slice by using cleanInput() on stdin
		userInput := cleanInput(scanner.Text())

		// checking length of input to avoid panic, just continue if no input received
		if len(userInput) == 0 {
			continue
		}

		command, exist := commandMap[userInput[0]]
		// if the user input is a command word and exists in the map
		if exist {
			// storing the return from command.callback()
			err := command.callback(la, cache)
			// if command.callback() does not return nil an error has occurred
			if err != nil {
				fmt.Println("Unknown Command")
			}
		}
		// print the first position word in the string slice
		// fmt.Printf("Your command was: %v\n", userInput[0])
	}
}

func cleanInput(text string) []string {
	// cleanInput must TRIM leading/trailing whitespaces
	// Convert the whole string to lowercase
	// split on whitespace into words
	// start with FIELDS and nest TRIMSPACE and TOLOWER
	return strings.Fields(
		strings.ToLower(
			strings.TrimSpace(text),
		),
	)
}

// function to exit out of the pokedex
func commandExit(la *locationArea, c *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	// send a successful exit code back to the function
	os.Exit(0)
	return nil
}

// function to display the available commands - error return value
func commandHelp(la *locationArea, c *pokecache.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	// range over the commandMap, you need the k&v for the loop
	for _, command := range commandMap {
		// access the fields for name and description
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

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
