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
	Name TypeName `json:"type"`
}

type TypeName struct {
	Name string `json:"name"`
}

type Stat struct {
	Name StatName `json:"stat"`
	Val  int      `json:"base_stat"`
}

type StatName struct {
	Name string `json:"name"`
}
