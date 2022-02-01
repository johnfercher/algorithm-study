package main

import (
	"fmt"
)

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {
	scaleFactor := int32(5)

	for i := 0; i < len(grades); i++ {
		if grades[i] < 38 {
			continue
		}

		if grades[i] >= 100 {
			continue
		}

		nextGrade := (int32(grades[i]/scaleFactor) * 5) + 5
		if nextGrade-grades[i] < 3 {
			grades[i] = nextGrade
		}
	}

	return grades
}

func main() {
	grades := []int32{84, 94, 21, 0, 18, 100, 18, 62, 30, 61, 53, 0, 43, 2, 29, 53, 61, 40}
	//grades := []int32{73, 67, 38, 33}
	result := gradingStudents(grades)
	fmt.Println(result)
}
