package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
    fmt.Println("Available commands:")
    for _,command:= range getCommands() {
	fmt.Println(">",command.name, ":", command.description)
    }
    return nil
}

