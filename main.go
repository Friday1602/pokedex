package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		parts := strings.Fields(input)
		var arg string
		if len(parts) > 1 {
			arg = parts[1]
		}
	
		if commandStruct, ok := commands[parts[0]]; ok {
			executeCommand(commandStruct, arg)
		} else {
			fmt.Printf("%v: command not found\n", input)
		}

	}

}
