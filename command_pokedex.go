package main

import (
	"fmt"
	"strings"
)

func commandPokedex(cfg *Config, arg string) error {
	if cfg == nil {
		return nil
	}

	if len(cfg.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty. Try catching some Pokemon!")
		return nil
	}

	list := []string{"Your Pokedex:"}
	for _, item := range cfg.Pokedex {
		list = append(list, fmt.Sprintf("- %s", item.Name))
	}

	fmt.Println(strings.Join(list, "\n"))

	return nil
}
