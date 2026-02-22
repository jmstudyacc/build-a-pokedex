package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"pokedexcli/internal/pokecache"
)

// constant created to store the Poke API Locations URL
const pokeAPILocations = "https://pokeapi.co/api/v2/location-area/"

type cliCommand struct {
	name            string
	description     string
	simpleCallback  func() error
	navCallback     func(*locationArea, *pokecache.Cache) error
	argsCallback    func([]string, *pokecache.Cache) error
	pokedexCallback func([]string, *pokecache.Cache, pokedex) error
}

// Hashmap to store the commands
var commandMap map[string]cliCommand

// to avoid circular dependencies, init() used to assign the map before main runs
func init() {
	commandMap = map[string]cliCommand{ // map with keys that are strings, and values that are cliCommand structs
		"help": {
			name:           "help",
			description:    "Displays a help message",
			simpleCallback: commandHelp,
		}, "exit": {
			name:           "exit",
			description:    "Exit the Pokedex",
			simpleCallback: commandExit,
		}, "map": {
			name:        "map",
			description: "Displays the next 20 locations",
			navCallback: mapForward,
		}, "mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			navCallback: mapBack,
		}, "explore": {
			name:         "explore",
			description:  "Displays the Pokemon encounters in that area",
			argsCallback: explore,
		}, "catch": {
			name:            "catch",
			description:     "Catch the target Pokemon",
			pokedexCallback: catch,
		}, "inspect": {
			name:            "inspect",
			description:     "Displays information of the target Pokemon",
			pokedexCallback: inspect,
		},
	}
}

func startRepl() {
	// generate internal struct to store next & previous location data
	la := &locationArea{}

	// instantiate the cache
	cache := pokecache.NewCache(5 * time.Minute)

	// create pokedex
	pokedexMap := pokedex{}

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
			var err error

			// Check which callback type is set and call it
			if command.simpleCallback != nil {
				err = command.simpleCallback()
			} else if command.navCallback != nil {
				err = command.navCallback(la, cache)
			} else if command.argsCallback != nil {
				// userInput[0] is the command name, so userInput[1:] are the arguments
				err = command.argsCallback(userInput[1:], cache)
			} else if command.pokedexCallback != nil {
				err = command.pokedexCallback(userInput[1:], cache, pokedexMap)
			}

			if err != nil {
				fmt.Println("Error:", err)
			}
		}
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
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	// send a successful exit code back to the function
	os.Exit(0)
	return nil
}

// function to display the available commands - error return value
func commandHelp() error {
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
