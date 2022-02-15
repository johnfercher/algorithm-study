package main

import "fmt"

func main() {
	m := [][]int{
		{0, 1, 2, 0},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
		{12, 13, 14, 15},
	}

	PrintMatrix(m)
	ZeroMatrix(m)
	PrintMatrix(m)
}

func PrintMatrix(m [][]int) {
	for _, line := range m {
		for _, element := range line {
			fmt.Printf("%d\t", element)
		}
		fmt.Println()
	}
	fmt.Println()
}

func ZeroMatrix(m [][]int) {
	manipulatedValues := make(map[string]bool)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if m[i][j] == 0 && !manipulatedValues[KeyGenerator(i, j)] {
				for k := 0; k < len(m); k++ {
					if m[i][k] != 0 {
						manipulatedValues[KeyGenerator(i, k)] = true
						m[i][k] = 0
					}

					if m[k][j] != 0 {
						manipulatedValues[KeyGenerator(k, j)] = true
						m[k][j] = 0
					}
				}
			}
		}
	}
}

func KeyGenerator(i, j int) string {
	return fmt.Sprintf("%d-%d", i, j)
}
