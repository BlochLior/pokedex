package main

import (
	"errors"
	"fmt"
)

// commandExplore -
func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]

	// fetch pokemon data for the given location
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	// print the pokemon found in this location
	fmt.Println("Exploring " + location.Name + "...")
	fmt.Println("Found Pokemon:")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
