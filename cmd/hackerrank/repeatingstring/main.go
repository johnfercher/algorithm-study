package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
	if getMultiplier(s, n) == n {
		return n
	}

	repeated := int64(0)
	usefulString := stringNormalizer(s, n)

	fmt.Println(usefulString)

	for _, value := range usefulString {
		if value == "a" {
			repeated++
		}
	}

	return repeated
}

func stringNormalizer(s string, n int64) []string {
	usefulString := s
	if len(s) >= int(n) {
		usefulString = s[:n]
		return strings.Split(usefulString, "")
	}

	return strings.Split(merge(s, s, n), "")
}

func getMultiplier(subString string, n int64) int64 {
	return n / int64(len(subString))
}

func merge(subString string, superSubstring string, n int64) string {
	if len(superSubstring)*2 > int(n) {
		remainN := n - int64(len(superSubstring))
		if remainN < int64(len(subString)) {
			return subString + subString[:remainN]
		}

		return superSubstring + merge(subString, subString, remainN)
	}

	duplicated := superSubstring + superSubstring
	return merge(subString, duplicated, n)
}

func main() {
	//result := repeatedString("aba", 10)
	//result := repeatedString("a", 1000000000000)
	result := repeatedString("a", 100)
	fmt.Println(result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
