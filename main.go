package main

import (
	"fmt"
	"github.com/adamthiede/bootdev-pokedex/internal/pokeapi"
	"time"
)

type config struct {
	pokeApiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	fmt.Println("Welcome to PokeApi")
	cfg := config{
		pokeApiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
