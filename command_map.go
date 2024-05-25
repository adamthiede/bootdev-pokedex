package main

import (
	"errors"
	"fmt"
	"log"
	//"encoding/json"
	//"github.com/adamthiede/bootdev-pokedex/internal/pokeapi"
)

func commandMap(cfg *config) error {
	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("no previous page")
	}
	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil
}
