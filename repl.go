package main

import (
	"PokedexCli/internal/pokecache"
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func runRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	GlobalConfig := &Config{}
	GlobalConfig.Client = http.Client{}
	cache := pokecache.NewCache(5 * time.Second)
	GlobalConfig.Cache = cache

	for {
		command, fullCommands := promptUser("Pokedex > ", scanner)

		if len(command) == 0 {
			continue
		}

		// Programmatically get all commands and their callbacks. If user command matches one of theirs, call it's function
		commands := getCommands(GlobalConfig)

		if _, ok := commands[command]; ok {
			callback := commands[command].Callback
			args := []string{}

			if len(fullCommands) > 1 {
				args = fullCommands[1:]
			}

			callback(GlobalConfig, args...)

		} else {
			fmt.Printf("I'm sorry, I don't know what %v means\n", command)
		}
	}
}
