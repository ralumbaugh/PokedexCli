package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func runRepl() {


	scanner := bufio.NewScanner(os.Stdin)
	
	for ; ; {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		
		userInput := scanner.Text()
		command := cleanInput(userInput)[0]

		if len(userInput) == 0 {
			continue
		}
		
		if command == "exit" {
			commandExit()
		} else if command == "help" {
			commandHelp()
		}
		
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	
	os.Exit(0)

	return nil
}

func commandHelp() error {
	commands := map[string]cliCommand {
		"exit": {
			name:		"exit",
			description: "Exit the pokedex",
			callback:	commandExit,
		},
		"help": {
			name:		"help",
			description:	"Displays a help message",
			callback:	commandHelp,
		},
	}

	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.description)	
	}

	return nil
}

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}