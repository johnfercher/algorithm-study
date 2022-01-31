package main

import (
	"fmt"
	"strings"
)

func main() {
	result := twoStrings("hi", "world")
	fmt.Println(result)
}

func twoStrings(s1 string, s2 string) string {
	lettersCache := make(map[string]bool)

	strings1 := strings.Split(s1, "")
	for _, value := range strings1 {
		lettersCache[value] = true
	}

	strings2 := strings.Split(s2, "")
	for _, value := range strings2 {
		thereIs := lettersCache[value]
		if thereIs {
			return "YES"
		}
	}

	return "NO"
}
