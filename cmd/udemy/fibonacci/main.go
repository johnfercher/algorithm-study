package main

import "fmt"

func main() {
	valueToFind := int64(89)
	index, ok := fibonacci(1, 0, valueToFind, 0)
	fmt.Println(index)
	fmt.Println(ok)
}

func fibonacci(current int64, previous int64, valueToFind int64, i int) (int, bool) {
	if current == valueToFind {
		return i + 1, true
	}

	if current > valueToFind {
		return 0, false
	}

	fmt.Println(current)
	newPrevious := current
	return fibonacci(current+previous, newPrevious, valueToFind, i+1)
}
