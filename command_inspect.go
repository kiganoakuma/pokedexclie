package main

import "fmt"

func commandInpect(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("you  must provide a pokemon's name")
		return nil
	}
	pokeName := args[0]
	cfg.pokedex.InspectPokemon(pokeName)
	return nil
}
