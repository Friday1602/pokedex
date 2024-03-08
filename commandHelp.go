package main

import "fmt"

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
