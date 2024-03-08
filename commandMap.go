package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func commandMap() error {
	pokeLocationAPI := locationInfo.Next
	if locationInfo.Previous == nil && locationInfo.Next == "" {
		pokeLocationAPI = "https://pokeapi.co/api/v2/location/"
	}
	// get from locaction endpoint and send response to resp
	resp, err := http.Get(pokeLocationAPI)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d and body: %s", resp.StatusCode, body)
	}
	if err != nil {
		return err
	}

	// create locationInfo to store data from json
	err = json.Unmarshal(body, &locationInfo)
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

func commandMapb() error {
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
