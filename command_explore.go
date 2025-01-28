package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func commandExplore(cfg *Config, args ...string) error {

	if len(args) < 1 {
		fmt.Print("Please enter a location to explore. Example: explore canalave-city\n")
		return nil
	}
	location := args[0]

	newLocation := &encounterConfig{}

	cachedData, ok := cfg.Cache.Get(location)

	if ok {
		err := json.Unmarshal(cachedData, newLocation)

		if err != nil {
			return err
		}
	} else {
		data, err := callApi(cfg.Client, nil, "GET", "https://pokeapi.co/api/v2/location-area/"+location)

		if err != nil {
			if strings.Contains(err.Error(), "status code: 404") {
				fmt.Printf("Could not find location %v. Find locations to explore using the map and mapb command!\n", location)
			}
			return err
		}

		cfg.Cache.Add(location, data)

		err = json.Unmarshal(data, newLocation)

		if err != nil {
			return err
		}
	}

	fmt.Printf("Exploring %v...\nFound Pokemon:\n", location)

	for _, enc := range newLocation.Encounters {
		fmt.Printf("- %v\n", enc.Pokemon.Name)
	}

	return nil
}
