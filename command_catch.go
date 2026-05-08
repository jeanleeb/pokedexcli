package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"
)

const catch_threshold = 50

func commandCatch(cfg *Config, arg string) error {
	if arg == "" {
		return errors.New("you need to pass a Pokemon name to catch")
	}
	if cfg == nil {
		return errors.New("config is nil")
	}

	pokemon, err := cfg.ApiClient.GetPokemon(strings.ToLower(arg))
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	catchRoll := rand.IntN(pokemon.BaseExperience)
	if catchRoll <= catch_threshold {
		fmt.Printf("%s was caught!\n", pokemon.Name)

		if _, exists := cfg.Pokedex[pokemon.Name]; !exists {
			cfg.Pokedex[pokemon.Name] = pokemon
		}

		return nil
	}

	fmt.Printf("%s escaped!\n", pokemon.Name)

	return nil
}
