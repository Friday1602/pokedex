package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
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
	fmt.Println("Throwing a Pokeball at ", targetPokemon.Name, "...")
	time.Sleep(3 * time.Second)
	switch {
	case index < 50:
		fmt.Println(targetPokemon.Name, " was caught")
	default:
		fmt.Println(targetPokemon.Name, " escaped!")
	}

	return nil
}
