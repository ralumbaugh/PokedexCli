package main

import "fmt"

func commandPokedex(cfg *Config, args ...string) error {
	fmt.Print("Your Pokedex:\n")

	for _, pokemon := range cfg.Pokedex {
		fmt.Printf("  - %v\n", pokemon.Name)
	}
	return nil
}
