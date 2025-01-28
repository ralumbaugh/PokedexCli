package main

import (
	"fmt"
	"os"
)

func commandExit(config *Config, args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")

	config.Cache.Close()
	os.Exit(0)

	return nil
}
