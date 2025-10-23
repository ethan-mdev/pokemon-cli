package locations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethan-mdev/pokemon-cli/internal/pokecache"
)

var cache = pokecache.NewCache()

// Fetches 20 location areas starting from the given offset
func GetLocationAreas(offset int) ([]LocationArea, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%d&limit=20", offset)

	// Check cache first
	cachedLocations, found := cache.Get(url)
	if found {
		var cachedResults []LocationArea
		if err := json.Unmarshal(cachedLocations, &cachedResults); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cached locations: %v", err)
		}
		return cachedResults, nil
	}

	// Cache miss - make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch locations: %v", err)
	}
	defer resp.Body.Close()

	var locationsResponse LocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationsResponse); err != nil {
		return nil, fmt.Errorf("failed to decode locations response: %v", err)
	}

	// Cache the results
	resultsJSON, err := json.Marshal(locationsResponse.Results)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal locations response: %v", err)
	}
	cache.Add(url, resultsJSON)

	return locationsResponse.Results, nil
}

// Fetches the list of Pokemon in a given location area by its URL
func GetPokemonInLocationArea(url string) ([]string, error) {

	// Check cache first
	cachedArea, found := cache.Get(url)
	if found {
		var cachedResponse ExploreResponse
		if err := json.Unmarshal(cachedArea, &cachedResponse); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cached location area: %v", err)
		}
		var pokemonNames []string
		for _, encounter := range cachedResponse.PokemonEncounters {
			pokemonNames = append(pokemonNames, encounter.Pokemon.Name)
		}
		return pokemonNames, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch location area: %v", err)
	}
	defer resp.Body.Close()

	var exploreResponse ExploreResponse
	if err := json.NewDecoder(resp.Body).Decode(&exploreResponse); err != nil {
		return nil, fmt.Errorf("failed to decode location area response: %v", err)
	}

	// Cache the results
	resultsJSON, err := json.Marshal(exploreResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal location area response: %v", err)
	}
	cache.Add(url, resultsJSON)

	var pokemonNames []string
	for _, encounter := range exploreResponse.PokemonEncounters {
		pokemonNames = append(pokemonNames, encounter.Pokemon.Name)
	}

	return pokemonNames, nil
}
