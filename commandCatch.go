package main

import "encoding/json"

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
	return nil
}
