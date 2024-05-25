package main

import (
	"fmt"
)

func commandList() error {
	pokemans := make([]pokemon, 0)
	fmt.Println("you have caught", len(pokemans), "pokemans")
	return nil
}

