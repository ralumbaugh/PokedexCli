package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)


func runRepl() {


	scanner := bufio.NewScanner(os.Stdin)
	GlobalConfig := &Config{}
	GlobalConfig.Client = http.Client{}
	
	for {
		userInput, command := promptUser("Pokedex > ", scanner)

		if len(userInput) == 0 {
			continue
		}
		
		// Programmatically get all commands and their callbacks. If user command matches one of theirs, call it's function
		commands:= getCommands(GlobalConfig)

		if _, ok := commands[command]; ok {
			callback := commands[command].Callback
	
			callback(GlobalConfig)
		} else {
			fmt.Printf("I'm sorry, I don't know what %v means\n", command)
		}
	}
}