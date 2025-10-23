package pokemon

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []Stat `json:"stats"`
	Types          []Type `json:"types"`
	Sprites        struct {
		FrontDefault string `json:"front_default"`
		FrontShiny   string `json:"front_shiny"`
		BackDefault  string `json:"back_default"`
		BackShiny    string `json:"back_shiny"`
	} `json:"sprites"`
}

type Stat struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"stat"`
}

type Type struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}
