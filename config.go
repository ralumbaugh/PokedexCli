package main

import (
	"PokedexCli/internal/pokecache"
	"net/http"
)

type Config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Client   http.Client
	Results  []Location `json:"results"`
	Cache    *pokecache.Cache
	Pokedex  map[string]Pokemon
	PlayerXp int
}
