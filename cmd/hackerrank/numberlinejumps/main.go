package main

import (
	"fmt"
)

/*
 * Complete the 'kangaroo' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER x1
 *  2. INTEGER v1
 *  3. INTEGER x2
 *  4. INTEGER v2
 */

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if x1 == x2 {
		return "YES"
	}

	if x1 > x2 && v1 > v2 {
		return "NO"
	}

	if x2 > x1 && v2 > v1 {
		return "NO"
	}

	currentX1 := x1
	currentX2 := x2

	stopCondition := getStopCondition(currentX1, currentX2)
	for stopCondition(currentX1, currentX2) {
		newCurrentX1 := currentX1 + v1
		newCurrentX2 := currentX2 + v2

		if newCurrentX1 == newCurrentX2 {
			return "YES"
		}

		currentX1 = newCurrentX1
		currentX2 = newCurrentX2
	}

	return "NO"
}

func getStopCondition(x1, x2 int32) func(int32, int32) bool {
	if x1 > x2 {
		return func(i1 int32, i2 int32) bool {
			return i1 >= i2
		}
	}

	return func(i1 int32, i2 int32) bool {
		return i1 <= i2
	}
}

func main() {
	result := kangaroo(0, 3, 4, 2)
	fmt.Println(result)
}
