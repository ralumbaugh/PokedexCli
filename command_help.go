package main

import "fmt"

func commandHelp(config *Config, args ...string) error {
	commands := getCommands(config)

	fmt.Printf("Welcome to the Pokedex!\n\nUsage:\n%v\n", addBorder())

	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.Name, command.Description)
	}

	return nil
}
