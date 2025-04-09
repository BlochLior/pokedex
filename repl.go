package main

import (
	"strings"
)

// split to words by whitespace
// trim leading or trailing whitespace
// lowercase input
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
