package main

import (
	"errors"
	"fmt"
	"math/rand"
	//"encoding/json"
	//"github.com/adamthiede/bootdev-pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeApiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	chanceToCatch := rand.Intn(pokemon.BaseExperience)
	if chanceToCatch < 50 {
		// catch
		fmt.Println(pokemonName, "was caught!")
		cfg.caughtPokemon[pokemonName] = pokemon
	} else {
		fmt.Println(pokemonName, "escaped!")
	}

	return nil
}
