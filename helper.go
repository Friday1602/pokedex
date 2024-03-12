package main

import "fmt"

type cliCommand struct {
	name            string
	description     string
	callback        func() error
	callbackWithArg func(string) error
}

func printPrompt() {
	fmt.Print("pokedex >")
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:            "explore",
			description:     "explore the pokemon in the specific location",
			callback:        nil,
			callbackWithArg: commandExplore,
		},
	}
}

func executeCommand(commandStruct cliCommand, arg string) {
	if arg == "" {
		err := commandStruct.callback()
		if err != nil {
			fmt.Printf("Error executing command %s\n", err)
		}
	}

}
