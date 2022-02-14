package main

import (
	"fmt"
	"strings"
)

func main() {
	candidates := []Candidate{
		{A: "dog god", B: "dog god"},
		{A: "focus", B: "coffin"},
		{A: "dog", B: "boy"},
	}

	for _, candidate := range candidates {
		palindrome := candidate.IsPalindrome()
		fmt.Printf("%s = %s -> %v\n", candidate.A, candidate.B, palindrome)
	}
}

type Candidate struct {
	A string
	B string
}

func (s *Candidate) IsPalindrome() bool {
	if len(s.A) != len(s.B) {
		return false
	}

	charactersA := strings.Split(s.A, "")
	charactersB := strings.Split(s.B, "")

	for i := 0; i < len(charactersA); i++ {
		fmt.Printf("%s - %s\n", charactersA[i], charactersB[len(charactersB)-i-1])
		if charactersA[i] != charactersB[len(charactersB)-i-1] {
			return false
		}
	}

	return true
}
