package main

func commandPokedex(cfg *config, args ...string) error {
	cfg.pokedex.DisplayPokedexContents()
	return nil
}
