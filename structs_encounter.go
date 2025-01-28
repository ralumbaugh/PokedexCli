package main

type encounterConfig struct {
	Encounters []Encounter `json:"pokemon_encounters"`
}

type Encounter struct {
	Pokemon Pokemon
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
