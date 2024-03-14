package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

var targetPokemon pokemon

func commandCatch(arg ...string) error {
	pokemonAPI := "https://pokeapi.co/api/v2/pokemon/" + arg[0]
	body, err := readFromAPI(pokemonAPI)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &targetPokemon)
	if err != nil {
		return err
	}

	index := rand.Float64() * float64(targetPokemon.BaseExperience)

	switch {
	case index < 50:
		fmt.Println("You've caught ", targetPokemon.Name)
	default:
		fmt.Println("Catch failed")
	}

	return nil
}
