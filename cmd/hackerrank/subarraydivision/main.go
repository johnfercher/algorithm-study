package main

import "fmt"

/*
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */

func birthday(s []int32, d int32, m int32) int32 {
	if len(s) <= 0 {
		return 0
	}

	if len(s) == 1 && s[0] == d {
		return 1
	}

	wayToGetCount := int32(0)

	for i := 0; i < len(s)-int(m)+1; i++ {
		sum := int32(0)
		for j := 0; j < int(m); j++ {
			k := i + j
			valueK := s[k]
			fmt.Printf("(%d %d) ", k, valueK)
			sum += valueK
		}
		fmt.Printf(" -> %d", sum)
		fmt.Println()
		if sum == d {
			wayToGetCount++
		}
	}

	return wayToGetCount
}

func main() {
	/*s := []int32{4}
	d := int32(4)
	m := int32(1)*/

	/*s := []int32{1,2,1,3,2}
	d := int32(3)
	m := int32(2)*/

	s := []int32{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1}
	d := int32(18)
	m := int32(7)

	result := birthday(s, d, m)
	fmt.Println(result)
}
