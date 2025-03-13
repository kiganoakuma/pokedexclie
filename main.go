package main

import (
	"time"

	"github.com/kiganoakuma/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       *NewPokedex(),
	}
	startRepl(cfg)

}
