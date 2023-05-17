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

	bribes := 0

	for i := int32(len(q)) - 1; i > 0; i-- {
		if q[i] != i+1 {
			if q[i-1] == i+1 {
				bribes += 1
				q[i-1], q[i] = q[i], q[i-1]
			} else if q[i-2] == i+1 {
				bribes += 2
				q[i-2], q[i-1], q[i] = q[i-1], q[i], q[i-2]
			} else {
				fmt.Println("Too chaotic")
				return
			}
		}
	}

	fmt.Println(bribes)
}

func main() {
	//array := []int32{2}
	array := []int32{2, 1, 5, 3, 4}
	//array := []int32{2, 5, 1, 3, 4}
	//array := []int32{1, 2, 5, 3, 7, 8, 6, 4}
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
