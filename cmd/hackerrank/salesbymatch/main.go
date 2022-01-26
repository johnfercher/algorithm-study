package main

import (
	"fmt"
)

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
	pairs := int32(0)
	socksMap := make(map[int32]int)

	for _, value := range ar {
		_, ok := socksMap[value]
		if ok {
			socksMap[value] += 1
		} else {
			socksMap[value] = 1
		}
	}

	for _, value := range socksMap {
		pairs += int32(value / 2)
	}

	return pairs
}

func main() {
	result := sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20})
	fmt.Println(result)
}
