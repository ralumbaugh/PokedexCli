package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) > 1 {
		fmt.Print("Please include the name of a pokemon to catch")
		return nil
	}

	pokemon := args[0]

	if newPokemon, ok := cfg.Pokedex[pokemon]; ok {
		// Pokemon is in pokedex already. No need for api call
		attemptCatch(cfg, newPokemon)
	} else {
		newPokemon := &Pokemon{}
		resp, err := callApi(cfg.Client, nil, "GET", "https://pokeapi.co/api/v2/pokemon/"+pokemon)

		if err != nil {
			if strings.Contains(err.Error(), "status code: 404") {
				fmt.Printf("Could not find pokemon %v. Try exploring areas with the explore <area-name> command to find pokemon to catch!\n", pokemon)
			}
			return err
		}

		err = json.Unmarshal(resp, newPokemon)

		if err != nil {
			return err
		}

		fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)

		attemptCatch(cfg, *newPokemon)
	}

	return nil
}

func attemptCatch(cfg *Config, pokemon Pokemon) {
	// Players gain xp while catching pokemon. The more pokemon they catch, the easier it will be to catch harder pokemon.
	// Players should also gain some experience from attempting a catch, regardless of success.
	cfg.PlayerXp += 10
	caughtPokemon := calculateChance(cfg.PlayerXp, pokemon.Xp)

	if caughtPokemon {
		cfg.PlayerXp += pokemon.Xp / 4
		cfg.Pokedex[pokemon.Name] = pokemon
		fmt.Printf("%v was caught! Player XP: %v\n", pokemon.Name, cfg.PlayerXp)
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}
}

func calculateChance(playerXp int, pokemonXp int) bool {
	pokemonRoll := rand.Intn(pokemonXp)
	playerRoll := rand.Intn(playerXp)

	fmt.Printf("Pokemon roll: %v Player Roll: %v\n", pokemonRoll, playerRoll)

	return playerRoll >= pokemonRoll
}
