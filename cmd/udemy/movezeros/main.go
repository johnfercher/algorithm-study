package main

import (
	"algorithm-study/internal/generate"
	"fmt"
)

func main() {
	size := 50
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
	zerosQtd := len(arr) - len(notZerosIndex)
	zerosSwapped := 0

	for i := 0; i < len(arr); i++ {
		// if all zeros was swapped at least one time
		// necessary to avoid re-swap
		if zerosSwapped >= zerosQtd {
			break
		}

		if arr[i] == 0 {
			arr[i], arr[notZerosIndex[j]] = arr[notZerosIndex[j]], arr[i] // swap based on not zero number array index
			j--                                                           // decrement zero number array index
			zerosSwapped++                                                // increment zeros swapped
		}
	}

	return arr
}
