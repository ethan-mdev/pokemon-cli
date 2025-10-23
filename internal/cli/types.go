package cli

import "github.com/ethan-mdev/pokemon-cli/internal/user"

type cliCommand struct {
	name        string
	description string
	parameters  []string
	callback    func(*config, []string) error
}

type config struct {
	locationOffset int
	pokedex        *user.Pokedex
}
