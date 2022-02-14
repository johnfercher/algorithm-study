package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{
		"use",
		"using",
		"usu",
	}

	matrix := [][]string{
		{"u", "s", "e"},
		{"a", "i", "h"},
		{"g", "n", "k"},
	}

	for _, word := range words {
		hasWord := hasWordInMatrix(word, matrix)
		fmt.Println(hasWord)
	}
}

func keyGenerator(i, j int) string {
	return fmt.Sprintf("%d-%d", i, j)
}

func hasWordInMatrix(word string, matrix [][]string) bool {
	if word == "" {
		return false
	}

	if len(matrix) == 0 {
		return false
	}

	letters := strings.Split(word, "")
	l := 0
	i := 0
	j := 0

	var visitedPositions = make(map[string]bool)

	hasLetter := hasLetterInMatrix(letters, l, i, j, matrix, visitedPositions)
	return hasLetter
}

func hasLetterInMatrix(word []string, l, i, j int, matrix [][]string, visitedPositions map[string]bool) bool {
	matchLetters := word[l] == matrix[i][j]
	if !matchLetters {
		return false
	}

	if visitedPositions[keyGenerator(i, j)] {
		return false
	}

	visitedPositions[keyGenerator(i, j)] = true

	fmt.Printf("(%d, %d): %s\n", i, j, word[l])
	isLastWord := l == len(word)-1

	if isLastWord {
		return true
	}

	// Up
	if i > 0 {
		has := hasLetterInMatrix(word, l+1, i-1, j, matrix, visitedPositions)
		if has {
			return true
		}
	}

	// Right
	if j < len(matrix[i])-1 {
		has := hasLetterInMatrix(word, l+1, i, j+1, matrix, visitedPositions)
		if has {
			return true
		}
	}

	// Down
	if i < len(matrix)-1 {
		has := hasLetterInMatrix(word, l+1, i+1, j, matrix, visitedPositions)
		if has {
			return true
		}
	}

	// Left
	if j > 0 {
		has := hasLetterInMatrix(word, l+1, i, j-1, matrix, visitedPositions)
		if has {
			return true
		}
	}

	return false
}
