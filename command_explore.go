package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *Config, arg string) error {
	if arg == "" {
		return errors.New("you need to pass an area to explore")
	}

	fmt.Printf("Exploring %s...\n", arg)
	areaDetails, err := cfg.ApiClient.GetAreaDetails(arg)
	if err != nil {
		return err
	}

	if len(areaDetails.PokemonEncounters) == 0 {
		return errors.New("no Pokemon found in this area")
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range areaDetails.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
