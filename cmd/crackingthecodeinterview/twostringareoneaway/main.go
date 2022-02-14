package main

import (
	"fmt"
	"strings"
)

func main() {
	candidates := []Candidate{
		{A: "ple", B: "pale"},
		{A: "pale", B: "pales"},
		{A: "pale", B: "bale"},
		{A: "bae", B: "pale"},
	}

	for _, candidate := range candidates {
		oneAway := candidate.IsOneAway()
		fmt.Printf("%s = %s -> %v\n", candidate.A, candidate.B, oneAway)
	}
}

type Candidate struct {
	A string
	B string
}

func (s *Candidate) IsOneAway() bool {
	charactersCount := make(map[string]int)

	first := s.A
	second := s.B

	if len(s.B) < len(s.A) {
		first = s.B
		second = s.A
	}

	charactersFirst := strings.Split(first, "")
	for _, character := range charactersFirst {
		if _, ok := charactersCount[character]; ok {
			charactersCount[character] = charactersCount[character] + 1
		}

		charactersCount[character] = 1
	}

	charactersSecond := strings.Split(second, "")
	operationsAllowed := 1
	for _, character := range charactersSecond {
		_, ok := charactersCount[character]
		if !ok {
			if operationsAllowed == 0 {
				return false
			}

			operationsAllowed--
		}

		charactersCount[character] = charactersCount[character] - 1
	}

	return true
}
