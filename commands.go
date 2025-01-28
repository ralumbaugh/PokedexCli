package main

import (
	"PokedexCli/internal/pokecache"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type CliCommand struct {
	Name        string
	Description string
	Config      *Config
	Callback    func(*Config, ...string) error
}

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

func getCommands(cfg *Config) map[string]CliCommand {
	commands := map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Config:      cfg,
			Callback:    commandHelp,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays the names of the pokemon available in a region. Example usage: explore <area-name>",
			Config:      cfg,
			Callback:    commandExplore,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of the next 20 regions",
			Config:      cfg,
			Callback:    commandNextMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the names of the previous 20 regions",
			Config:      cfg,
			Callback:    commandPrevMap,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Config:      cfg,
			Callback:    commandExit,
		},
	}

	return commands
}

func commandExit(config *Config, args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")

	config.Cache.Close()
	os.Exit(0)

	return nil
}

func commandHelp(config *Config, args ...string) error {
	commands := getCommands(config)

	fmt.Printf("Welcome to the Pokedex!\n\nUsage:\n%v\n", addBorder())

	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.Name, command.Description)
	}

	return nil
}

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

func commandExplore(cfg *Config, args ...string) error {

	if len(args) < 1 {
		fmt.Print("Please enter a location to explore. Example: explore canalave-city\n")
		return nil
	}
	location := args[0]

	newLocation := &encounterConfig{}

	cachedData, ok := cfg.Cache.Get(location)
	if ok {
		// Data is in cache. Let's goooooo!
		// fmt.Printf("%v was found in cache. Loading from Cache.\n", location)
		err := json.Unmarshal(cachedData, newLocation)

		if err != nil {
			return err
		}
	} else {
		// fmt.Printf("%v was not found in cache. Calling api.\n", location)

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

// Helper Functions below here

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
		fmt.Printf("\n%v\n%v already exists in our cache. Loading from cache.\n%v\n", addBorder(), url, addBorder())

		err := json.Unmarshal(cachedData, newConfig)

		if err != nil {
			return err
		}
	} else {
		fmt.Printf("\n%v\n%v doesn't exist in our cache yet. Making api call.\n%v\n", addBorder(), url, addBorder())
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
