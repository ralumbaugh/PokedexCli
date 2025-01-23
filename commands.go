package main

import (
	"fmt"
	"net/http"
	"os"
)

type CliCommand struct {
	Name		string
	Description	string
	Config		*Config
	Callback	func(*Config) error
}

type Config struct {
	Next		string `json:"next"`
	Previous	string `json:"previous"`
	Client		http.Client
	Results		[]Location `json:"results"`
}

type Location struct {
	Name	string `json:"name"`
	URL 	string `json:"url"`
}

func getCommands(cfg *Config) map[string]CliCommand {
	commands := map[string]CliCommand{
		"exit": {
			Name:        	"exit",
			Description:	"Exit the pokedex",
			Config:		 	cfg,
			Callback:    	commandExit,
		},
		"help": {
			Name:       	"help",
			Description:	"Displays a help message",
			Config:			cfg,
			Callback:   	commandHelp,
		},
		"map": {
			Name:			"map",
			Description:	"Displays the names of the next 20 regions",
			Config:			cfg,
			Callback:		commandNextMap,
		},
		"mapb": {
			Name:			"mapb",
			Description:	"Displays the names of the previous 20 regions",
			Config:			cfg,
			Callback:		commandPrevMap,
		},
	}

	return commands
}

func commandExit(config *Config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	
	os.Exit(0)

	return nil
}

func commandHelp(config *Config) error {
	commands := getCommands(config)

	fmt.Printf("Welcome to the Pokedex!\n\nUsage:\n%v\n", addBorder())
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.Name, command.Description)	
	}

	return nil
}

func commandNextMap(cfg *Config) error {
	err := getMap(cfg, getMapUrl(cfg.Next))
	return err
}

func commandPrevMap(cfg *Config) error {
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
	
	err := callApi(cfg.Client, nil, "GET", url, newConfig)
	if err != nil {
		return err
	}

	cfg.Next = newConfig.Next
	cfg.Previous = newConfig.Previous
	
	fmt.Printf("Map locations:\n%v\n", addBorder())
	for _, loc := range newConfig.Results {
		fmt.Printf("%v\n",loc.Name)
	}

	return nil
}
