package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{
		"sword",
		"emblem",
	}

	for _, word := range words {
		isAllUnique := IsAllUnique(word)
		fmt.Printf("%s: %v\n", word, isAllUnique)
	}
}

func IsAllUnique(word string) bool {
	letters := strings.Split(word, "")
	uniqueLetter := make(map[string]bool)

	for _, letter := range letters {
		if uniqueLetter[letter] {
			return false
		}

		uniqueLetter[letter] = true
	}

	return true
}
