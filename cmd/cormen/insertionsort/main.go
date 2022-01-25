package main

/*
	O(n^2)
*/

import "fmt"

func main() {
	insertionSort(12, 11, 13, 5, 6)
}

func insertionSort(arr ...int) []int {
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("key: %d[%d]\n", arr[i], i)
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			fmt.Printf("Swap: %d[%d] <-> %d[%d]\n", arr[j], j, arr[j-1], j-1)
			arr[j], arr[j-1] = arr[j-1], arr[j]
			fmt.Printf("After J iteration: %v\n", arr)
		}
		fmt.Printf("After I iteration: %v\n", arr)
	}
	return arr
}
