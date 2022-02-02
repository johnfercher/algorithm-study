package main

import "fmt"

/*
 * Complete the 'breakingRecords' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY scores as parameter.
 */

func breakingRecords(scores []int32) []int32 {
	if len(scores) <= 1 {
		return []int32{0, 0}
	}

	minPoint := scores[0]
	minRecordsBreaking := int32(0)
	maxPoint := scores[0]
	maxRecordsBreaking := int32(0)

	for i := 1; i < len(scores); i++ {
		if scores[i] < minPoint {
			minPoint = scores[i]
			minRecordsBreaking++
		} else if scores[i] > maxPoint {
			maxPoint = scores[i]
			maxRecordsBreaking++
		}
	}

	return []int32{maxRecordsBreaking, minRecordsBreaking}
}

func main() {
	result := breakingRecords([]int32{})
	fmt.Println(result)
}
