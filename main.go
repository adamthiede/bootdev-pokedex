package main

import (
    "fmt"
    "github.com/adamthiede/bootdev-pokedex/internal/pokeapi"
)

type config struct {
    pokeApiClient pokeapi.Client
    nextLocationAreaURL *string
    prevLocationAreaURL *string
}

func main() {
    fmt.Println("Welcome to PokeApi")
    cfg := config {
	pokeApiClient: pokeapi.NewClient(),
    }
    startRepl(&cfg)
}
