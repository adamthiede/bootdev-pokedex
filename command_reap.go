package main

import (
	"fmt"
)

func commandReap(cfg *config, args ...string) error {
	cfg.pokeApiClient.InvalidateCache()
	fmt.Println("invalidating cache.")
	return nil
}
