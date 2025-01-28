package main

type encounterConfig struct {
	Encounters []Encounter `json:"pokemon_encounters"`
}

type Encounter struct {
	Pokemon Pokemon_encounter
}

type Pokemon_encounter struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
