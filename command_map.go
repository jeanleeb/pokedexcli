package main

import (
	"fmt"
)

const LIMIT = 20
const LOCATION_AREAS_URL = "https://pokeapi.co/api/v2/location-area"

type LocationAreasResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *Config) error {
	var url string
	if cfg.Next != "" {
		url = cfg.Next
	} else {
		url = fmt.Sprintf("%s?limit=%d", LOCATION_AREAS_URL, LIMIT)
	}

	locationAreas, err := cfg.ApiClient.GetLocationAreas(url)
	if err != nil {
		return err
	}

	cfg.Next = locationAreas.Next
	cfg.Previous = locationAreas.Previous

	for _, item := range locationAreas.Results {
		fmt.Println(item.Name)
	}

	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreas, err := cfg.ApiClient.GetLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = locationAreas.Next
	cfg.Previous = locationAreas.Previous

	for _, item := range locationAreas.Results {
		fmt.Println(item.Name)
	}

	return nil
}
