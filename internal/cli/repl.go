package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Starts the REPL for the Pokemon CLI
func Start() {
	cfg := &config{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		command := input[0]
		args := input[1:]

		if cmd, exists := commands[command]; exists {
			if err := cmd.callback(cfg, args); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

// Cleans and splits input text into lowercase words
func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	words := strings.Fields(lowercase)
	return words
}
