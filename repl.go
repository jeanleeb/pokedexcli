package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jeanleeb/pokedexcli/internal/pokeapi"
)

func startRepl() {
	apiClient := pokeapi.NewClient()
	cfg := &Config{
		Next:      "",
		Previous:  "",
		ApiClient: apiClient,
		Pokedex:   map[string]pokeapi.PokemonResponse{},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		arg := ""
		if len(words) > 1 {
			arg = words[1]
		}
		err := command.callback(cfg, arg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	clean := []string{}
	for word := range strings.FieldsSeq(strings.ToLower(text)) {
		clean = append(clean, word)
	}

	return clean
}

type Config struct {
	Next      string
	Previous  string
	ApiClient pokeapi.Client
	Pokedex   map[string]pokeapi.PokemonResponse
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, arg string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex CLI",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List all Pokémon in an area. Usage: explore <area name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon. Usage: catch <pokemon name>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon in your Pokedex. Usage: inspect <pokemon name>",
			callback:    commandInspect,
		},
	}
}
