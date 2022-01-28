package main

import (
	"algorithm-study/internal/generate"
	"fmt"
	"strings"
)

func main() {
	txt := generate.RandomText()

	// O(n)
	wordArr := strings.Split(txt, " ")
	logestWord := FindLongest(wordArr)
	fmt.Println(logestWord)
}

func FindLongest(words []string) string {
	maxLength := 0
	longestWord := ""

	for _, word := range words {
		if len(word) > maxLength {
			maxLength = len(word)
			longestWord = word
		}
	}

	return longestWord
}
