package main

import (
	"fmt"
	"io"
	"net/http"
)

type cliCommand struct {
	name        string
	description string
	callback    func(...string) error
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
			name:        "explore",
			description: "explore the pokemon in the specific location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch the pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect the pokemon caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list of the pokemon caught",
			callback:    commandPokedex,
		},
	}
}

func executeCommand(commandStruct cliCommand, arg string) {
	err := commandStruct.callback(arg)
	if err != nil {
		fmt.Printf("Error executing command %s\n", err)

	}

}

// data not in cache, make API request
// get from locaction endpoint and send response to resp
func readFromAPI(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and body: %s", resp.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}
