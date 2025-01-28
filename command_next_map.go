package main

import (
	"encoding/json"
	"fmt"
)

func commandNextMap(cfg *Config, args ...string) error {
	err := getMap(cfg, getMapUrl(cfg.Next))
	return err
}

func commandPrevMap(cfg *Config, args ...string) error {
	if cfg.Previous == "" {
		fmt.Print("You're on the first page\n")
		return nil
	} else {
		err := getMap(cfg, getMapUrl(cfg.Previous))
		return err
	}
}

func getMapUrl(url string) string {
	if url == "" {
		return "https://pokeapi.co/api/v2/location-area"
	}
	return url
}

func getMap(cfg *Config, url string) error {
	newConfig := &Config{}

	// Check if it's in the cache first
	if cachedData, ok := cfg.Cache.Get(url); ok {
		err := json.Unmarshal(cachedData, newConfig)

		if err != nil {
			return err
		}
	} else {
		data, err := callApi(cfg.Client, nil, "GET", url)

		if err != nil {
			return err
		}

		cfg.Cache.Add(url, data)

		// We want to add the first page to the cache whether we're getting there moving forward or backward. First time it loads, the url is https://pokeapi.co/api/v2/location-area. Second time it loads it is https://pokeapi.co/api/v2/location-area?offset=0&limit=20
		if url == "https://pokeapi.co/api/v2/location-area" {
			cfg.Cache.Add("https://pokeapi.co/api/v2/location-area?offset=0&limit=20", data)
		}

		err = json.Unmarshal(data, newConfig)
		if err != nil {
			return err
		}
	}

	cfg.Next = newConfig.Next
	cfg.Previous = newConfig.Previous

	fmt.Printf("Map locations:\n%v\n", addBorder())
	for _, loc := range newConfig.Results {
		fmt.Printf("%v\n", loc.Name)
	}

	return nil
}
