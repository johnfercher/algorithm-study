package main

import (
	"fmt"
	"strings"
)

func main() {
	candidates := []Candidate{
		{A: "dog", B: "god"},
		{A: "focus", B: "coffin"},
		{A: "dog", B: "boy"},
	}

	for _, candidate := range candidates {
		permutation := candidate.IsPermutation()
		fmt.Printf("%s = %s -> %v\n", candidate.A, candidate.B, permutation)
	}
}

type Candidate struct {
	A string
	B string
}

func (s *Candidate) IsPermutation() bool {
	if len(s.A) != len(s.B) {
		return false
	}

	charactersCount := make(map[string]int)

	charactersA := strings.Split(s.A, "")
	for _, character := range charactersA {
		if _, ok := charactersCount[character]; ok {
			charactersCount[character] = charactersCount[character] + 1
		}

		charactersCount[character] = 1
	}

	charactersB := strings.Split(s.B, "")
	for _, character := range charactersB {
		_, ok := charactersCount[character]
		if !ok {
			return false
		}

		charactersCount[character] = charactersCount[character] - 1
	}

	return true
}
