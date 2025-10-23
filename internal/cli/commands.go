package cli

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/ethan-mdev/pokemon-cli/internal/display"
	"github.com/ethan-mdev/pokemon-cli/internal/locations"
	"github.com/ethan-mdev/pokemon-cli/internal/pokemon"
	"github.com/ethan-mdev/pokemon-cli/internal/user"
)

var commands map[string]cliCommand

// Initialize the commands map with predefined commands
func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			parameters:  []string{},
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			parameters:  []string{},
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations in the Pokedex",
			parameters:  []string{},
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the Pokedex",
			parameters:  []string{},
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Shows a list of all pokemon in a location",
			parameters:  []string{"location"},
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catches a pokemon by name",
			parameters:  []string{"pokemon"},
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects a caught pokemon",
			parameters:  []string{"pokemon"},
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all caught pokemon",
			parameters:  []string{},
			callback:    commandPokedex,
		},
		"image": {
			name:        "image",
			description: "Displays an image of a caught pokemon",
			parameters:  []string{"pokemon"},
			callback:    commandImage,
		},
	}
}

// Exits the application
func commandExit(*config, []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// Displays help information about available commands
func commandHelp(*config, []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

// Displays the next 20 locations in the Pokedex
func commandMap(cfg *config, args []string) error {
	locationAreas, err := locations.GetLocationAreas(cfg.locationOffset)
	if err != nil {
		return err
	}

	for _, location := range locationAreas {
		fmt.Printf("%s\n", location.Name)
	}
	cfg.locationOffset += 20
	return nil
}

// Displays the previous 20 locations in the Pokedex
func commandMapBack(cfg *config, args []string) error {
	if cfg.locationOffset <= 0 {
		fmt.Println("You're on the first page!")
		return nil
	}

	cfg.locationOffset -= 20

	locationAreas, err := locations.GetLocationAreas(cfg.locationOffset)
	if err != nil {
		return err
	}

	for _, location := range locationAreas {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}

// Shows a list of all pokemon in a location
func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: explore <location>")
	}
	location := args[0]

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", location)
	pokemonNames, err := locations.GetPokemonInLocationArea(url)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring location: %s\n", location)
	fmt.Println("Found Pokemon:")
	for _, name := range pokemonNames {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}

// Catches a pokemon by name
func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: catch <pokemon>")
	}
	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonData, err := pokemon.GetPokemonByName(pokemonName)
	if err != nil {
		return err
	}

	chance := 75 - (pokemonData.BaseExperience / 10)
	if chance < 0 {
		chance = 0
	}

	if rand.Intn(100) > chance {
		fmt.Printf("%s escaped! Better luck next time.\n", pokemonName)
		return nil
	}

	// Initialize pokedex if nil
	if cfg.pokedex == nil {
		cfg.pokedex = &user.Pokedex{CaughtPokemon: make(map[string]pokemon.Pokemon)}
	}
	cfg.pokedex.CaughtPokemon[pokemonName] = pokemonData
	fmt.Printf("Congratulations! You caught %s!\n", pokemonName)

	return nil
}

// Inspects a caught pokemon
func commandInspect(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: inspect <pokemon>")
	}
	pokemonName := args[0]

	if cfg.pokedex == nil {
		return fmt.Errorf("you haven't caught any pokemon yet")
	}

	pokemonData, found := cfg.pokedex.CaughtPokemon[pokemonName]
	if !found {
		return fmt.Errorf("pokemon %s not found in your pokedex", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %d\n", pokemonData.Height)
	fmt.Printf("Weight: %d\n", pokemonData.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonData.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemonData.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}

// Displays all caught pokemon
func commandPokedex(cfg *config, args []string) error {
	if cfg.pokedex == nil || len(cfg.pokedex.CaughtPokemon) == 0 {
		fmt.Println("Your pokedex is empty. Catch some pokemon!")
		return nil
	}
	fmt.Println("Caught Pokemon:")
	for name := range cfg.pokedex.CaughtPokemon {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}

// Displays an image of a caught pokemon
func commandImage(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: image <pokemon>")
	}
	pokemonName := args[0]
	if cfg.pokedex == nil {
		return fmt.Errorf("you haven't caught any pokemon yet")
	}
	pokemonData, found := cfg.pokedex.CaughtPokemon[pokemonName]
	if !found {
		return fmt.Errorf("pokemon %s not found in your pokedex", pokemonName)
	}

	fmt.Printf("Displaying %s:\n", pokemonName)
	return display.DisplayImage(pokemonData.Sprites.FrontDefault)
}
