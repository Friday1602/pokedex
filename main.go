package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func printPrompt() {
	fmt.Println("pokedex >")
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
	}
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("usage: ")
	fmt.Println()
	commands := getCommands()

	for key, value := range commands {
		fmt.Println(key, " : ", value.description)
	}

	return nil
}

func commandExit() error {
	fmt.Println("Exiting the Pokedex!")
	os.Exit(0)
	return nil
}
func executeCommand(commandStruct cliCommand) {
	err := commandStruct.callback()
	if err != nil {
		fmt.Printf("Error executing command %s\n", err)
	}
}
func main() {

	commands := getCommands()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		printPrompt()
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if commandStruct, ok := commands[input]; ok {
			executeCommand(commandStruct)
		} else {
			fmt.Printf("%v: command not found\n", input)
		}

	}

}
