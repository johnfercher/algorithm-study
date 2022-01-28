package main

import (
	"algorithm-study/internal/generate"
	"fmt"
)

func main() {
	size := 70
	maxValue := 5

	arr := generate.UnsortedIntArray(size, maxValue)
	fmt.Println(arr)
	fmt.Println(moveZeros(arr))
}

func moveZeros(arr []int) []int {
	notZerosIndex := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			notZerosIndex = append(notZerosIndex, i)
		}
	}

	j := len(notZerosIndex) - 1 // star iteration from last element in not zero number array

	for i := 0; i < len(arr); i++ {
		// break if current iteration pass swapped value index (to not re-swap elements)
		if i >= notZerosIndex[j] {
			break
		}

		if arr[i] == 0 {
			arr[i], arr[notZerosIndex[j]] = arr[notZerosIndex[j]], arr[i] // swap based on not zero number array index
			j--                                                           // decrement zero number array index
		}
	}

	return arr
}
