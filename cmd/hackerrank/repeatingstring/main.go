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
	usefulString := strings.Split(s, "")
	subStringOccurrences := getSubstringOccurrences(usefulString)
	multiplier, rest := getMultiplierAndRest(usefulString, n)

	entireOccurrences := subStringOccurrences * multiplier

	for i := 0; i < int(rest); i++ {
		if usefulString[i] == "a" {
			entireOccurrences++
		}
	}

	return entireOccurrences
}

func getSubstringOccurrences(s []string) int64 {
	repeated := int64(0)
	for _, value := range s {
		if value == "a" {
			repeated++
		}
	}
	return repeated
}

func getMultiplierAndRest(subString []string, n int64) (int64, int64) {
	return n / int64(len(subString)), n % int64(len(subString))
}

func main() {
	//result := repeatedString("aba", 10)
	result := repeatedString("a", 1000000000000)
	//result := repeatedString("a", 100)
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
