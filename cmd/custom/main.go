package main

import (
	"crypto/rand"
	"custom/domain"
	"errors"
	"fmt"
	"math"
	"math/big"
)

func main() {
	size := 160
	idToFind := 169

	unsortedArray := initUnsortedArrayWithWorstCase(size, idToFind)

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
	sortedArray := insertionSort(unsortedArray, sortCounter)
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
	sortedArray := mergeSort(unsortedArray, sortCounter)
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
	counter.Add()
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
		opCounter.Add()

		if arr[i].ID == id {
			return arr[i], nil
		}
	}

	return nil, errors.New(fmt.Sprintf("%d not found", id))
}

func insertionSort(arr domain.Objects, counter *domain.Counter) domain.Objects {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			counter.Add()
			if arr[j].ID < arr[j-1].ID {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	return arr
}

func mergeSort(arr domain.Objects, counter *domain.Counter) domain.Objects {
	if len(arr) < 2 {
		return arr
	}

	aUnsorted := arr[:len(arr)/2]
	bUnsorted := arr[len(arr)/2:]

	aSorted := mergeSort(aUnsorted, counter)
	bSorted := mergeSort(bUnsorted, counter)

	return merge(aSorted, bSorted, counter)
}

func merge(a domain.Objects, b domain.Objects, counter *domain.Counter) domain.Objects {
	var final domain.Objects
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		counter.Add()
		if a[i].ID < b[j].ID {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		counter.Add()
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		counter.Add()
		final = append(final, b[j])
	}

	return final
}

func initUnsortedArrayWithWorstCase(size int, rightID int) domain.Objects {
	var arr domain.Objects
	for i := 0; i < size-1; i++ {
		arr = append(arr, generateRandomObject(math.MaxInt8, rightID))
	}

	arr = append(arr, &domain.Object{
		ID:   rightID,
		Name: "wrong",
	})

	return arr
}

func generateRandomObject(max, avoidID int) *domain.Object {
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	randNumber := int(randomNumber.Int64())

	if randNumber == avoidID {
		return &domain.Object{
			ID:   avoidID + 1,
			Name: "wrong",
		}
	}

	return &domain.Object{
		ID:   randNumber,
		Name: "wrong",
	}
}
