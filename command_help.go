package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	if len(args) > 0 {
		fmt.Println("Warning: 'help' command doesn't take any arguments. Ignoring extra input.")
	}

	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
