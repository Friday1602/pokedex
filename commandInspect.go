package main

import (
	"encoding/json"
	"fmt"
)

func commandInspect(pokemonName ...string) error {
	pokemonAPI := "https://pokeapi.co/api/v2/pokemon/" + pokemonName[0]
	if pokemonData, ok := pokedexCache.Get(pokemonAPI); ok {
		err := printPokemonInfo(pokemonData)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("You have not caught that pokemon")
	}
	return nil
}

func printPokemonInfo(body []byte) error {
	err := json.Unmarshal(body, &targetPokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Name: %s\n", targetPokemon.Name)
	fmt.Printf("Height: %d\n", targetPokemon.Height)
	fmt.Printf("Weight: %d\n", targetPokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range targetPokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range targetPokemon.Types {
		fmt.Printf("  -%s\n", pokeType.Type.Name)
	}

	return nil
}
