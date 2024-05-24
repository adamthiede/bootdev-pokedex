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
	callback    func() error
}

func getCommands() map[string]cliCommand {

    commandHelp:=func() error {
	for _,command:= range getCommands() {
	    fmt.Println(">",command.name, ":", command.description)
	}
	return nil
    }
    commandExit:=func() error {
	os.Exit(0)
	return nil
    }

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
    }
}

func main() {

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
		line=strings.TrimSpace(line)

		// get commands and check if input is in commands
		commands := getCommands()
		_, cmdInMap := commands[line]
		if cmdInMap {
		    // if the command is found, run it; functions are passed as values
			commands[line].callback()
		} else {
		    fmt.Println("Did not find your command", line, "in commands")
		}

	}
}
