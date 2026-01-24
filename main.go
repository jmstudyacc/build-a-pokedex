package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// create a scanner to read from os.Stdin
	scanner := bufio.NewScanner(os.Stdin) // calling this will block & wait for user input and CR
	for {
		fmt.Print("Pokedex > ")

		// scanner.Scan() required to start populating a buffer with text
		scanner.Scan()
		userInput := scanner.Text() // you then call .Text() on the buffer

		midSlice := strings.ToLower(userInput)
		// fmt.Printf("DEBUG: %s\n", midSlice)

		inputSlice := strings.Fields(midSlice)
		// fmt.Printf("DEBUG: %#v (length: %d)\n", inputSlice, len(inputSlice))

		fmt.Printf("Your command was: %v\n", inputSlice[0])
	}
}
