package main

import (
	"errors"
	"fmt"
	//"encoding/json"
	//"github.com/adamthiede/bootdev-pokedex/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]
	locationArea, err := cfg.pokeApiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in \"%s\" area:\n", args[0])
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
