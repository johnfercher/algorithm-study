package main

import (
	"algorithm-study/internal/domain"
	"algorithm-study/internal/generate"
	"algorithm-study/internal/sort"
	"fmt"
)

func main() {
	size1 := 10
	size2 := 20
	maxValue := 100
	counter := domain.NewCounter("sort")

	unsorted1 := generate.UnsortedIntArray(size1, maxValue)
	sorted1 := sort.MergeInt(unsorted1, counter)

	unsorted2 := generate.UnsortedIntArray(size2, maxValue)
	sorted2 := sort.MergeInt(unsorted2, counter)

	fmt.Println(mergeSorted(sorted1, sorted2))
}

func mergeSorted(arr1, arr2 []int) []int {
	var final []int

	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			final = append(final, arr1[i])
			i++
		} else {
			final = append(final, arr2[j])
			j++
		}
	}

	for ; i < len(arr1); i++ {
		final = append(final, arr1[i])
	}

	for ; j < len(arr2); j++ {
		final = append(final, arr2[j])
	}

	return final
}
