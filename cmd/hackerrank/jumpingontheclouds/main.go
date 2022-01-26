package main

import (
	"fmt"
)

/*
 * Complete the 'jumpingOnClouds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY c as parameter.
 */

func jumpingOnClouds(c []int32) int32 {
	jumps := int32(-1)
	i := 0

	for i < len(c) {
		fmt.Println(i)
		if i+2 < len(c) && c[i+2] == 0 {
			jumps++
			i += 2
			continue
		}

		jumps++
		i += 1
	}

	return jumps
}

func main() {
	//result := jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0})
	result := jumpingOnClouds([]int32{0, 0, 0, 1, 0, 0})
	fmt.Println(result)
}
