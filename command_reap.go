package main

import (
	"fmt"
)

func commandReap(cfg *config) error {
	cfg.pokeApiClient.InvalidateCache()
	fmt.Println("invalidating cache.")
	return nil
}
