package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	if len(args) > 0 {
		fmt.Println("Warning: 'exit' command doesn't take any arguments. Ignoring extra input.")
	}

	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
