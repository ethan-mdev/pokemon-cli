package pokemon

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethan-mdev/pokemon-cli/internal/pokecache"
)

var cache = pokecache.NewCache()

// Fetches Pokemon data by name
func GetPokemonByName(name string) (Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)

	// Check cache first
	cachedPokemon, found := cache.Get(url)
	if found {
		var cachedPokemonData Pokemon
		if err := json.Unmarshal(cachedPokemon, &cachedPokemonData); err == nil {
			return cachedPokemonData, nil
		}
	}

	// Fetch from API if not found in cache
	resp, err := http.Get(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed to fetch pokemon: %v", err)
	}
	defer resp.Body.Close()

	var pokemon Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return Pokemon{}, fmt.Errorf("failed to decode pokemon response: %v", err)
	}

	return pokemon, nil
}
