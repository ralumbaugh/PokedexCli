package main

import "fmt"

func commandInspect(cfg *Config, args ...string) error {
	if len(args) < 1 {
		fmt.Print("Please include the name of the pokemon you wish to look up\n")
		return nil
	}

	pokemonName := args[0]

	if pokemon, ok := cfg.Pokedex[pokemonName]; ok {
		fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)

		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%v: %v\n", stat.Name.Name, stat.Val)
		}
		fmt.Print("Types:\n")
		for _, stat := range pokemon.Types {
			fmt.Printf("  -%v\n", stat.Name.Name)
		}
	} else {
		fmt.Printf("You have not caught a %v\n", pokemonName)
	}

	return nil
}
