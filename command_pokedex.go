package main

import (
	"fmt"
	//"encoding/json"
	//"github.com/adamthiede/bootdev-pokedex/internal/pokeapi"
)

func commandPokedex(cfg *config, args ...string) error {
	pokemons := cfg.caughtPokemon
	if len(pokemons) == 0 {
		fmt.Println("You've caught no pokemon. Get out there!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for p := range pokemons {
		fmt.Printf(" - %s\n", p)
	}
	return nil
}
