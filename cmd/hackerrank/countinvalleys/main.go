package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

const (
	seaLevel = 0
	downHill = "D"
)

func countingValleys(steps int32, path string) int32 {
	valleys := int32(0)
	currentLevel := int32(0)
	paths := strings.Split(path, "")

	for i := 0; i < len(paths); i++ {
		currentLevel = computeLevel(currentLevel, paths[i])

		if isBellowSealLevel(currentLevel) {
			i++
			for i < len(paths) {
				currentLevel = computeLevel(currentLevel, paths[i])
				if isSealLevel(currentLevel) {
					valleys++
					break
				}
				i++
			}
		}
	}

	return valleys
}

func computeLevel(currentLevel int32, point string) int32 {
	if point == downHill {
		currentLevel--
		return currentLevel
	}

	currentLevel++
	return currentLevel
}

func isSealLevel(currentLevel int32) bool {
	return currentLevel == 0
}

func isBellowSealLevel(currentLevel int32) bool {
	return currentLevel < seaLevel
}

func main() {
	result := countingValleys(0, "UDDDUDUU")
	//result := countingValleys(0, "DUDDDUUDUU")
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
