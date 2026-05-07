package main

import (
	"fmt"

	"github.com/jeanleeb/pokedexcli/internal/pokeapi"
)

const LIMIT = 20

type LocationAreasResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *Config, arg string) error {
	var url string
	if cfg.Next != "" {
		url = cfg.Next
	} else {
		url = fmt.Sprintf("%s?limit=%d", pokeapi.LOCATION_AREAS_URL, LIMIT)
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

func commandMapb(cfg *Config, arg string) error {
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
