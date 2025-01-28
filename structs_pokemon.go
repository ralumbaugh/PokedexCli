package main

type Pokemon struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Xp     int    `json:"base_experience"`
	Stats  []Stat `json:"stats"`
	Types  []Type `json:"types"`
}

type Type struct {
	Name string `json:"name.name"`
}

type Stat struct {
	Name string `json:"stat.name"`
	Val  int    `json:"base_stat"`
}
