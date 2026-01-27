package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error // callback is a function that returns an error
}

type Config struct {
	nextURL     string
	previousURL string
}

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
			err := command.callback()
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

func mapForward(c *Config) error {
	return nil
}

func mapBack(c *Config) error {
	return nil
}
