package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * Complete the 'minimumBribes' function below.
 *
 * The function accepts INTEGER_ARRAY q as parameter.
 */

func minimumBribes(q []int32) {
	if len(q) == 1 {
		return
	}

	bribers := make(map[int32]int)

	// O(N^2)
	for i := 0; i < len(q); i++ {
		for j := i; j < len(q); j++ {
			if q[i] > q[j] {
				bribers[q[i]] = bribers[q[i]] + 1
				q[i], q[j] = q[j], q[i]
			}
		}
	}

	// O(n log n)?

	bribes := 0

	for _, b := range bribers {
		if b > 2 {
			fmt.Println("Too chaotic")
			return
		}
		bribes += b
	}

	fmt.Println(bribes)
}

func main() {
	//array := []int32{2}
	//array := []int32{2, 1, 5, 3, 4}
	//array := []int32{2, 5, 1, 3, 4}
	array := []int32{1, 2, 5, 3, 7, 8, 6, 4}
	minimumBribes(array)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
