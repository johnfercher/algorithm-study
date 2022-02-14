package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "abbcccddddeeeeeffffff"
	wordCompressed := Compress(word)

	fmt.Println(wordCompressed)
}

func Compress(word string) string {
	compressedWord := ""

	wordString := strings.Split(word, "")

	for i := 0; i < len(wordString); i++ {
		letter := wordString[i]

		count := 0
		j := i
		for j < len(wordString) {
			if wordString[j] != letter {
				break
			}
			count++
			j++
		}

		if count < 3 {
			for k := 0; k < count; k++ {
				compressedWord += letter
			}
		} else {
			compressedWord += fmt.Sprintf("%d%s", count, letter)
		}

		i += count - 1
	}

	return compressedWord
}
