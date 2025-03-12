package main

import "fmt"

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Useage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s : %s\n", command.name, command.description)
	}
	return nil
}
