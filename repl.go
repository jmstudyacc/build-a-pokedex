package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	// create a scanner to read from os.Stdin
	scanner := bufio.NewScanner(os.Stdin) // calling this will block & wait for user input and CR
	for {
		fmt.Print("Pokedex > ")

		// scanner.Scan() required to start populating a buffer with text
		scanner.Scan()

		// generating a string slice by using cleanInput() on stdin
		userInput := cleanInput(scanner.Text())

		// checking length of input to avoid panic
		if len(userInput) == 0 {
			continue
		}
		// print the first position word in the string slice
		fmt.Printf("Your command was: %v\n", userInput[0])
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
