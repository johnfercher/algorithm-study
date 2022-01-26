package main

import (
	"fmt"
)

/*
 * Complete the 'rotLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER d
 */

func rotLeft(a []int32, d int32) []int32 {
	newArray := []int32{}
	rotation := getNormalizedRotation(a, d)
	fmt.Println(rotation)

	for i := int(rotation); i < len(a); i++ {
		newArray = append(newArray, a[i])
	}

	for i := 0; i < int(rotation); i++ {
		newArray = append(newArray, a[i])
	}

	return newArray
}

func getNormalizedRotation(a []int32, d int32) int32 {
	return d % int32(len(a))
}

func main() {
	a := []int32{1, 2, 3, 4, 5}
	rotation := int32(4)

	result := rotLeft(a, rotation)

	fmt.Println(result)
}
