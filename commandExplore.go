package main

import (
	"encoding/json"
	"fmt"
	"time"

	pokecache "github.com/friday1602/pokedex/internal"
)

// struct of location area endpoint
type locationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

var pokeLocationArea locationArea
var encounterCache = pokecache.NewCache(10 * time.Second)

// explore location(arg) print out encounter pokemon in that location
func commandExplore(arg ...string) error {
	// create val of API we want to request
	pokeEncounterAPI := "https://pokeapi.co/api/v2/location-area/" + arg[0]

	// check if there is data in the cache, print out data from cache and return
	if pokeEncounterCacheData, ok := encounterCache.Get(pokeEncounterAPI); ok {
		fmt.Println("Data from cache >>>>")
		err := printPokeEncounter(pokeEncounterCacheData)
		if err != nil {
			return err
		}
		return nil
	}


	// data not in cache so read from API directly store to cache and print out data
	body, err := readFromAPI(pokeEncounterAPI)
	if err != nil {
		return err
	}
	encounterCache.Add(pokeEncounterAPI, body)
	err = printPokeEncounter(body)
	if err != nil {
		return err
	}
	return nil
}

func printPokeEncounter(body []byte) error {
	err := json.Unmarshal(body, &pokeLocationArea)
	if err != nil {
		return err
	}
	if len(pokeLocationArea.PokemonEncounters) > 0 {
		for _, pokemonEncounter := range pokeLocationArea.PokemonEncounters {
			fmt.Println(pokemonEncounter.Pokemon.Name)
		}
	} else {
		fmt.Println("pokemon not found")
	}
	return nil
}
