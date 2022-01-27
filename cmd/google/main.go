package main

import (
	"algorithm-study/internal/domain"
	"algorithm-study/internal/generate"
	"algorithm-study/internal/sort"
	"fmt"
)

func main() {
	pairSum := 8
	size := 1000
	maxValue := 20
	counter := domain.NewCounter("sort")

	unsortedArray := generate.UnsortedIntArray(size, maxValue)
	sortedArray := sort.MergeInt(unsortedArray, counter)
	fmt.Println(sortedArray)
	counter.Print()

	binarySearchCounter := domain.NewCounter("binary search")
	a1, b1, find1 := getPairIfExistsBinarySearch(pairSum, sortedArray, binarySearchCounter)
	binarySearchCounter.Print()
	fmt.Println(a1, b1, find1)

	expenentialSearch := domain.NewCounter("exponential")
	a2, b2, find2 := getPairIfExistsExponential(pairSum, sortedArray, expenentialSearch)
	expenentialSearch.Print()
	fmt.Println(a2, b2, find2)
}

func getPairIfExistsExponential(value int, arr []int, counter *domain.Counter) (int, int, bool) {
	for _, element1 := range arr {
		for _, element2 := range arr {
			counter.Add()
			if element1+element2 == value {
				return element1, element2, true
			}
		}
	}

	return 0, 0, false
}

func getPairIfExistsBinarySearch(value int, arr []int, counter *domain.Counter) (int, int, bool) {
	exists := false
	iterationValue := 0
	pair := 0

	for i, element := range arr {
		pairToFind := getPair(value, element)

		if thereIsValue(pairToFind, arr, counter) {
			iterationValue = arr[i]
			pair = pairToFind
			exists = true
			break
		}
	}

	return iterationValue, pair, exists
}

func thereIsValue(value int, arr []int, counter *domain.Counter) bool {
	counter.Add()
	if len(arr) <= 1 {
		return arr[0] == value
	}

	middleIndex := len(arr) / 2

	if value == arr[middleIndex] {
		return true
	}

	if value < arr[middleIndex] {
		return thereIsValue(value, arr[:middleIndex], counter)
	}

	if value > arr[middleIndex] {
		return thereIsValue(value, arr[middleIndex:], counter)
	}

	return false
}

func getPair(value int, current int) int {
	return value - current
}
