package main

import (
	"fmt"

	"github.com/kiganoakuma/pokedexcli/internal/pokeapi"
)

type Pokedex struct {
	Caught map[string]pokeapi.Pokemon
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		Caught: make(map[string]pokeapi.Pokemon),
	}
}

func (p *Pokedex) AddPokemon(pokemon pokeapi.Pokemon) {
	p.Caught[pokemon.Name] = pokemon
}

func (p *Pokedex) HasPokemon(name string) bool {
	_, ok := p.Caught[name]
	return ok
}

func (p *Pokedex) InspectPokemon(name string) {
	if p.HasPokemon(name) {
		pokemon := p.Caught[name]

		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Name: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types")
		for _, t := range pokemon.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}
	} else {
		fmt.Println("you have not caught this pokemon")
	}
}

func (p *Pokedex) DisplayPokedexContents() {
	for _, val := range p.Caught {
		fmt.Printf(" - %s\n", val.Name)
	}
}
