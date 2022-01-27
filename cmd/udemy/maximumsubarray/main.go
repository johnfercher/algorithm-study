package main

import (
	"algorithm-study/internal/generate"
	"fmt"
)

func main() {
	size := 10
	maxValue := 20

	arr := generate.UnsortedIntArray(size, maxValue)
	fmt.Println(arr)

	sum := findMaxSumSubArray(arr)
	fmt.Println(sum)
}

// Kadane Algorithm
func findMaxSumSubArray(arr []int) int {
	max := 0
	maxUntilNow := 0

	for _, value := range arr {
		maxUntilNow = maxUntilNow + value
		if maxUntilNow > max {
			max = maxUntilNow
		}

		if maxUntilNow < 0 {
			maxUntilNow = 0
		}
	}

	return max
}
