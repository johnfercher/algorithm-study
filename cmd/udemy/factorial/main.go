package main

import "fmt"

func main() {
	result := getFactorialRecursive(4)
	fmt.Println(result)

	result = getFactorialInteractive(4)
	fmt.Println(result)
}

func getFactorialRecursive(value int64) int64 {
	if value <= 1 {
		return 1
	}

	return value * getFactorialRecursive(value-1)
}

func getFactorialInteractive(value int64) int64 {
	result := int64(1)
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}
