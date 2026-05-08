package main

import (
	"errors"
	"fmt"
	"strings"
)

const inspect_template = `Name: %s
Height: %d
Weight: %d
Status:
- hp: %d
- attack: %d
- defense: %d
- special-attack: %d
- special-defense: %d
- speed: %d
Types:`

const type_line_template = "\n- %s"

type PokemonStats struct {
	hp             int
	attack         int
	defense        int
	specialAttack  int
	specialDefense int
	speed          int
}

func commandInspect(cfg *Config, arg string) error {
	if arg == "" {
		return errors.New("you need to pass a Pokemon name to inspect")
	}

	pokemon, exists := cfg.Pokedex[strings.ToLower(arg)]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	stats := PokemonStats{}
	for _, item := range pokemon.Stats {
		switch item.Stat.Name {
		case "hp":
			stats.hp = item.BaseStat
		case "attack":
			stats.attack = item.BaseStat
		case "defense":
			stats.defense = item.BaseStat
		case "special-attack":
			stats.specialAttack = item.BaseStat
		case "special-defense":
			stats.specialDefense = item.BaseStat
		case "speed":
			stats.speed = item.BaseStat
		}
	}

	output := fmt.Sprintf(inspect_template, pokemon.Name, pokemon.Height, pokemon.Weight, stats.hp, stats.attack, stats.defense, stats.specialAttack, stats.specialDefense, stats.speed)
	for _, t := range pokemon.Types {
		output += fmt.Sprintf(type_line_template, t.Type.Name)
	}

	fmt.Println(output)

	return nil
}
