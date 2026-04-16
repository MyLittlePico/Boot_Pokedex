package main

import(
	"strings"
)

func cleanInput(test string) []string{

	words := strings.Fields(strings.ToLower(test))
	
	return words
}