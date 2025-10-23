package user

import "github.com/ethan-mdev/pokemon-cli/internal/pokemon"

type Pokedex struct {
	CaughtPokemon map[string]pokemon.Pokemon
}
