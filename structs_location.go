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
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
