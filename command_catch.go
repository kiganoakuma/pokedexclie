package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("invalid location name")
	}
	name := args[0]
	pokemon, err := cfg.pokeapiClient.CatchPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	caughtStatus := pokemon.AttemptCatch()
	if caughtStatus {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex.AddPokemon(pokemon)

	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}
