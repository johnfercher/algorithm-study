package main

import (
	"algorithm-study/internal/domain"
	"algorithm-study/internal/generate"
	"algorithm-study/internal/sort"
	"fmt"
)

func main() {
	size := 10
	maxValue := 20

	arr := generate.UnsortedIntArray(size, maxValue)

	// Unsorted -> Time: O(n^2), Space: O(1)
	{
		fmt.Println(arr)
		counter := domain.NewCounter("unsorted o(n^2)")
		i, j, has := HasDuplicationQuadraticUnsorted(arr, counter)
		fmt.Println(i, j, has)
		counter.Print()
	}

	// Unsorted -> Time: O(n), Space: O(n)
	{
		fmt.Println(arr)
		counter := domain.NewCounter("unsorted o(n^2)")
		i, j, has := HasDuplicatedDynamicLinearUnsorted(arr, counter)
		fmt.Println(i, j, has)
		counter.Print()
	}

	// Sorted -> Time: O(n), Space: O(1)
	{
		sortCounter := domain.NewCounter("sort")
		arr := sort.MergeInt(arr, sortCounter)
		fmt.Println(arr)

		counter := domain.NewCounter("sorted o(n)")
		i, j, has := HasDuplicatedLinearSorted(arr, counter)
		fmt.Println(i, j, has)
		counter.Print()
	}

}

func HasDuplicationQuadraticUnsorted(arr []int, counter *domain.Counter) (int, int, bool) {
	for i, value1 := range arr {
		for j, value2 := range arr {
			counter.Increment()
			if i == j {
				continue
			}

			if value1 == value2 {
				return i, j, true
			}
		}
	}

	return 0, 0, false
}

func HasDuplicatedDynamicLinearUnsorted(arr []int, counter *domain.Counter) (int, int, bool) {
	foundElements := make(map[int]int)

	for i, value := range arr {
		counter.Increment()
		lastIndex, ok := foundElements[value]
		if ok {
			return lastIndex, i, true
		}

		foundElements[value] = i
	}

	return 0, 0, false
}

func HasDuplicatedLinearSorted(arr []int, counter *domain.Counter) (int, int, bool) {
	for i := 0; i < len(arr)-1; i++ {
		counter.Increment()
		if arr[i] == arr[i+1] {
			return i, i + 1, true
		}
	}

	return 0, 0, false
}
