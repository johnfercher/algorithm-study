package sort

import (
	"algorithm-study/internal/domain"
)

// INSERTION
func InsertObject(arr domain.Objects, counter *domain.Counter) domain.Objects {
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

// MERGE
func MergeInt(arr []int, counter *domain.Counter) []int {
	if len(arr) < 2 {
		return arr
	}

	aUnsorted := arr[:len(arr)/2]
	bUnsorted := arr[len(arr)/2:]

	aSorted := MergeInt(aUnsorted, counter)
	bSorted := MergeInt(bUnsorted, counter)

	return mergeIntArray(aSorted, bSorted, counter)
}

func mergeIntArray(a []int, b []int, counter *domain.Counter) []int {
	var final []int
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		counter.Add()
		if a[i] < b[j] {
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

func MergeObject(arr domain.Objects, counter *domain.Counter) domain.Objects {
	if len(arr) < 2 {
		return arr
	}

	aUnsorted := arr[:len(arr)/2]
	bUnsorted := arr[len(arr)/2:]

	aSorted := MergeObject(aUnsorted, counter)
	bSorted := MergeObject(bUnsorted, counter)

	return mergeObjectArray(aSorted, bSorted, counter)
}

func mergeObjectArray(a domain.Objects, b domain.Objects, counter *domain.Counter) domain.Objects {
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
