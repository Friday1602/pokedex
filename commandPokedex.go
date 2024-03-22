package main

import (
	"fmt"
)

func commandPokedex(arg ...string) error {
	for _, poke := range caughtPokemon {
		fmt.Printf(" -%s\n", poke)
	}
	return nil
}
