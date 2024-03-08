package main

import (
	"bufio"
	"fmt"
	"os"
)

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
