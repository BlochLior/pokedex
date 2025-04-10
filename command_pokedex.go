package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(args) > 0 {
		fmt.Println("Warning: 'pokedex' command doesn't take any arguments. Ignoring extra input.")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}
