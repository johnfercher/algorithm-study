package main

import (
	"algorithm-study/internal/domain"
	"algorithm-study/internal/generate"
	"algorithm-study/internal/sort"
	"errors"
	"fmt"
)

func main() {
	size := 10
	maxValue := 200
	idToFind := 169

	unsortedArray := generate.UnsortedObjectArrayWithWorstCase(size, maxValue, idToFind)

	// First Solution
	unSortedSolution(unsortedArray, idToFind)

	// Second Solution
	insertSortedSolution(unsortedArray, idToFind)

	// Third Solution
	mergeSortedSolution(unsortedArray, idToFind)
}

func unSortedSolution(unsortedArray domain.Objects, idToFind int) {
	counterUnsortedOp := domain.NewCounter("find")
	objectUnsorted, err := findUnsorted(unsortedArray, idToFind, counterUnsortedOp)
	if err != nil {
		fmt.Println(err.Error())
	}

	counterUnsortedOp.Print()
	objectUnsorted.Print()
	fmt.Println()
}

func insertSortedSolution(unsortedArray domain.Objects, idToFind int) {
	sortCounter := domain.NewCounter("insertion sort")
	findCounter := domain.NewCounter("find")
	sortedArray := sort.InsertObject(unsortedArray, sortCounter)
	objectSorted, err := findIndexOfSorted(sortedArray, idToFind, findCounter)
	if err != nil {
		fmt.Println(err.Error())
	}

	sortCounter.Print()
	findCounter.Print()
	objectSorted.Print()
	fmt.Println()
}

func mergeSortedSolution(unsortedArray domain.Objects, idToFind int) {
	sortCounter := domain.NewCounter("merge sort")
	findCounter := domain.NewCounter("find")
	sortedArray := sort.MergeObject(unsortedArray, sortCounter)
	objectSorted, err := findIndexOfSorted(sortedArray, idToFind, findCounter)
	if err != nil {
		fmt.Println(err.Error())
	}

	sortCounter.Print()
	findCounter.Print()
	objectSorted.Print()
	fmt.Println()
}

func findIndexOfSorted(arr domain.Objects, id int, counter *domain.Counter) (*domain.Object, error) {
	counter.Increment()
	middlePosition := len(arr) / 2
	if id == arr[middlePosition].ID {
		return arr[middlePosition], nil
	}

	if id < arr[middlePosition].ID {
		return findIndexOfSorted(arr[:len(arr)/2], id, counter)
	}

	if id > arr[middlePosition].ID {
		return findIndexOfSorted(arr[len(arr)/2:], id, counter)
	}

	return nil, errors.New(fmt.Sprintf("%d not found", id))
}

func findUnsorted(arr domain.Objects, id int, opCounter *domain.Counter) (*domain.Object, error) {
	for i := 0; i < len(arr); i++ {
		opCounter.Increment()

		if arr[i].ID == id {
			return arr[i], nil
		}
	}

	return nil, errors.New(fmt.Sprintf("%d not found", id))
}
