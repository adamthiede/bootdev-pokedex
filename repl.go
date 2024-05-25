package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type pokemon struct {
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "get pokemon for a region",
			callback:    commandExplore,
		},
		"reap": {
			name:        "reap",
			description: "invalidate cache",
			callback:    commandReap,
		},
	}
}

func startRepl(cfg *config) {
	quit := false
	for !quit {
		// print prompt
		fmt.Print("pokedex > ")
		// read text
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		cleanedInput := cleanInput(line)
		if len(cleanedInput) == 0 {
			continue
		}

		// we don't yet use the rest of cleanedInput
		command := cleanedInput[0]
		args := []string{}
		if len(cleanedInput) > 1 {
		    args = cleanedInput[1:]
		}

		// get commands and check if input is in commands
		commands := getCommands()

		_, cmdInMap := commands[command]
		if !cmdInMap {
			fmt.Println("Invalid command:", command)
			continue
		}

		err = commands[command].callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func cleanInput(input string) []string {
	lowered := strings.ToLower(input)
	stripped := strings.TrimSpace(lowered)
	words := strings.Fields(stripped)
	return words
}
