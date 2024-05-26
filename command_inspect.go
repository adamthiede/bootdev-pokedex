package main

import (
	"errors"
	"fmt"
	//"encoding/json"
	//"github.com/adamthiede/bootdev-pokedex/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		fmt.Printf("You have not caught %s\n", pokemonName)
	} else {
		fmt.Printf("You've caught %s\n", pokemonName)
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("ID: %v\n", pokemon.ID)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, x := range pokemon.Stats {
			fmt.Printf(" - %s: %v\n", x.Stat.Name, x.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, x := range pokemon.Types {
			fmt.Printf(" - %s\n", x.Type.Name)
		}
	}

	return nil
}
