package main

/*
	O(n*log(n))
*/

import "fmt"

const (
	formater = "-"
)

func main() {
	levelFormater := formater
	_ = mergeSort(levelFormater, 12, 11, 13, 5, 6)
}

func mergeSort(levelFormater string, arr ...int) []int {
	fmt.Printf("%s input %v\n", levelFormater, arr)
	if len(arr) < 2 {
		fmt.Printf("%s output %v\n", levelFormater, arr)
		return arr
	}

	first := mergeSort(appendLevel(levelFormater), arr[:len(arr)/2]...)
	second := mergeSort(appendLevel(levelFormater), arr[len(arr)/2:]...)

	sorted := merge(appendLevel(levelFormater), first, second)
	fmt.Printf("%s output %v\n", levelFormater, sorted)
	return sorted
}

func merge(levelFormater string, first []int, second []int) []int {
	fmt.Printf("%s merge %v, %v\n", levelFormater, first, second)

	final := []int{}
	i := 0
	j := 0

	for i < len(first) && j < len(second) {
		fmt.Printf("%s i: %d, j: %d\n", levelFormater, i, j)
		if first[i] < second[j] {
			fmt.Printf("%s got i %d[%d]\n", levelFormater, first[i], i)
			final = append(final, first[i])
			i++
		} else {
			fmt.Printf("%s got j %d[%d]\n", levelFormater, second[j], j)
			final = append(final, second[j])
			j++
		}
	}

	for ; i < len(first); i++ {
		final = append(final, first[i])
	}

	for ; j < len(second); j++ {
		final = append(final, second[j])
	}

	return final
}

func appendLevel(levelFormater string) string {
	return levelFormater + formater
}
