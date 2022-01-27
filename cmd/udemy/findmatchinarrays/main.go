package main

import (
	"algorithm-study/internal/domain"
	"algorithm-study/internal/generate"
	"algorithm-study/internal/mapper"
	"fmt"
)

func main() {
	size1 := 50
	size2 := 100
	maxValue := 2000

	arr1 := generate.UnsortedIntArray(size1, maxValue)
	fmt.Println(arr1)
	arr2 := generate.UnsortedIntArray(size1, size2)
	fmt.Println(arr2)

	// O(a+b)
	{
		counter := domain.NewCounter("O(a+b)")
		matchedNumber, match := hasMatchElementsWithMap(arr1, arr2, counter)
		fmt.Println(matchedNumber, match)
		counter.Print()
	}

	// O(a*b)
	{
		counter := domain.NewCounter("O(a*b)")
		matchedNumber, match := hasMatchElementsWithNestedLoops(arr1, arr2, counter)
		fmt.Println(matchedNumber, match)
		counter.Print()
	}
}

func hasMatchElementsWithMap(arr1, arr2 []int, counter *domain.Counter) (int, bool) {
	m1 := mapper.ArrayToMap(arr1, counter)

	for _, value := range arr2 {
		counter.Increment()
		if m1[value] {
			return value, true
		}
	}

	return 0, false
}

func hasMatchElementsWithNestedLoops(arr1, arr2 []int, counter *domain.Counter) (int, bool) {
	for _, a := range arr1 {
		for _, b := range arr2 {
			counter.Increment()
			if a == b {
				return a, true
			}
		}
	}

	return 0, false
}
