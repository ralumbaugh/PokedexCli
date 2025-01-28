package main

type CliCommand struct {
	Name        string
	Description string
	Config      *Config
	Callback    func(*Config, ...string) error
}

func getCommands(cfg *Config) map[string]CliCommand {
	commands := map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Config:      cfg,
			Callback:    commandHelp,
		},
		"catch": {
			Name:        "catch",
			Description: "Attempts to catch a pokemon. The higher experience it has, the harder it will be to catch. Example usage: catch <pokemon-name>",
			Config:      cfg,
			Callback:    commandCatch,
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
		"inspect": {
			Name:        "inspect",
			Description: "Inspects the pokedex entry for a caught pokemon",
			Config:      cfg,
			Callback:    commandInspect,
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
