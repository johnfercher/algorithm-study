package main

import (
	"algorithm-study/internal/generate"
	"fmt"
)

func main() {
	size := 9
	maxValue := 10

	unsortedArray := generate.UnsortedIntArray(size, maxValue)
	fmt.Println(unsortedArray)
}
