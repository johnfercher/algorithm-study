package main

import (
	"algorithm-study/internal/domain"
	"algorithm-study/internal/generate"
	"algorithm-study/internal/sort"
	"fmt"
)

type KeyValue struct {
	Key   int
	Value int
}

func main() {
	pairSum := 8
	size := 10000
	maxValue := 50

	counter := domain.NewCounter("sort")

	unsortedArray := generate.UnsortedIntArray(size, maxValue)
	fmt.Println(unsortedArray)
	sortedArray := sort.MergeInt(unsortedArray, counter)
	//counter.Print()

	// Unsorted Solutions
	{
		exponentialSearchConter := domain.NewCounter("unsorted exponential")
		a1, b1, find1 := getPairIfExistsExponential(pairSum, sortedArray, exponentialSearchConter)
		fmt.Println(a1, b1, find1)
		exponentialSearchConter.Print()

		dynamicProgrammingCounter := domain.NewCounter("unsorted dynamic programming")
		a2, b2, find2 := getPairIfExistsDynamicProgramming(pairSum, sortedArray, dynamicProgrammingCounter)
		fmt.Println(a2, b2, find2)
		dynamicProgrammingCounter.Print()
	}

	// Sorted Solutions
	{
		binarySearchCounter := domain.NewCounter("sorted binary search")
		keyValueArr := generateKeyValueArray(sortedArray, binarySearchCounter)
		a1, b1, find1 := getPairIfExistsBinarySearch(pairSum, keyValueArr, binarySearchCounter)
		fmt.Println(a1, b1, find1)
		binarySearchCounter.Print()

		linearBidirectionalCounter := domain.NewCounter("sorted bidirectional search")
		a2, b2, find2 := getPairIfExistsBidirectional(pairSum, sortedArray, linearBidirectionalCounter)
		fmt.Println(a2, b2, find2)
		linearBidirectionalCounter.Print()
	}
}

func getPairIfExistsDynamicProgramming(value int, arr []int, counter *domain.Counter) (int, int, bool) {
	sumPair := make(map[int]int)

	for _, element := range arr {
		counter.Increment()
		sum := value - element

		sumPairFromElement, ok := sumPair[element]
		if ok {
			return sumPairFromElement, element, true
		}

		sumPair[sum] = element
	}

	return 0, 0, false
}

func getPairIfExistsExponential(value int, arr []int, counter *domain.Counter) (int, int, bool) {
	for i, element1 := range arr {
		for j, element2 := range arr {
			counter.Increment()

			if i == j {
				continue
			}

			if element1+element2 == value {
				return element1, element2, true
			}
		}
	}

	return 0, 0, false
}

func getPairIfExistsBidirectional(value int, arr []int, counter *domain.Counter) (int, int, bool) {
	low := 0
	high := len(arr) - 1

	for low < high {
		counter.Increment()

		sum := arr[low] + arr[high]
		if sum == value {
			return arr[low], arr[high], true
		}

		if sum > value {
			high--
			continue
		}

		low++
	}

	return 0, 0, false
}

func getPairIfExistsBinarySearch(value int, arr []KeyValue, counter *domain.Counter) (int, int, bool) {
	exists := false
	iterationValue := 0
	pair := 0

	for i, element := range arr {
		pairToFind := getPair(value, element.Value)
		keyValueToFind := KeyValue{
			Key:   i,
			Value: pairToFind,
		}

		if thereIsValue(keyValueToFind, arr, counter) {
			iterationValue = arr[i].Value
			pair = pairToFind
			exists = true
			break
		}
	}

	return iterationValue, pair, exists
}

func thereIsValue(keyValue KeyValue, arr []KeyValue, counter *domain.Counter) bool {
	counter.Increment()
	if len(arr) <= 1 {
		return KeyAreDifferentAndValueIsEqual(keyValue, arr[0])
	}

	middleIndex := len(arr) / 2

	if KeyAreDifferentAndValueIsEqual(keyValue, arr[middleIndex]) {
		return true
	}

	if keyValue.Value < arr[middleIndex].Value {
		return thereIsValue(keyValue, arr[:middleIndex], counter)
	}

	if keyValue.Value > arr[middleIndex].Value {
		return thereIsValue(keyValue, arr[middleIndex:], counter)
	}

	return false
}

func generateKeyValueArray(arr []int, counter *domain.Counter) []KeyValue {
	var newArr []KeyValue
	for index, element := range arr {
		counter.Increment()
		newArr = append(newArr, KeyValue{
			Key:   index,
			Value: element,
		})
	}

	return newArr
}

func KeyAreDifferentAndValueIsEqual(a KeyValue, b KeyValue) bool {
	if a.Key == b.Key {
		return false
	}

	return a.Value == b.Value
}

func getPair(value int, current int) int {
	return value - current
}
