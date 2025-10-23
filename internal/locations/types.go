package locations

type LocationArea struct {
	Name string `json:"name"`
}

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
	}
}

type ExploreResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type LocationResponse struct {
	Results []LocationArea `json:"results"`
}
