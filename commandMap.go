package main

import (
	"encoding/json"
	"fmt"
	"time"

	pokecache "github.com/friday1602/pokedex/internal"
)

// location struct for location endpoint json
type location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var locationInfo location
var cache = pokecache.NewCache(5 * time.Second)

func commandMap(arg ...string) error {
	// check if Map is on the first page
	pokeLocationAPI := locationInfo.Next
	if locationInfo.Previous == nil && locationInfo.Next == "" {
		pokeLocationAPI = "https://pokeapi.co/api/v2/location-area/"
	}

	// check if data is already in cache.
	if cacheData, ok := cache.Get(pokeLocationAPI); ok {
		fmt.Println("Data from cache >>>>")
		err := printLocation(cacheData)
		if err != nil {
			return err
		}
		return nil
	}

	body, err := readFromAPI(pokeLocationAPI)
	if err != nil {
		return err
	}

	// store data in cache
	cache.Add(pokeLocationAPI, body)
	//process and print data loction
	err = printLocation(body)
	if err != nil {
		return err
	}
	return nil
}

func commandMapb(arg ...string) error {
	if locationInfo.Previous == nil {
		fmt.Println("Cannot go back: you are on the first page of information")
		return nil
	}
	if previousLocation, ok := locationInfo.Previous.(string); ok {
		locationInfo.Next = previousLocation
	} else {
		fmt.Println("assertion fail: cannot assert previousLocation to string type")
	}
	err := commandMap()
	if err != nil {
		return err
	}
	return nil
}

func printLocation(body []byte) error {
	// create locationInfo to store data from json
	err := json.Unmarshal(body, &locationInfo)
	if err != nil {
		return err
	}
	// check if data exist and print name of location
	if len(locationInfo.Results) > 0 {
		for _, name := range locationInfo.Results {
			fmt.Println(name.Name)
		}
		return nil
	} else {
		fmt.Println("no location found")
		return nil
	}
}
