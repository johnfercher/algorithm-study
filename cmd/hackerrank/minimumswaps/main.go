package main

import "fmt"

func main() {
	//arr := []int32{4, 3, 1, 2}
	arr := []int32{7, 5, 3, 2, 4, 1, 6}

	minSwaps := minimumSwaps(arr)
	fmt.Println(minSwaps)
}

func minimumSwaps(arr []int32) int32 {
	valueIndexMap := make(map[int32]int)

	for index, value := range arr {
		valueIndexMap[value] = index
	}

	swaps := int32(0)
	for i := int32(0); i < int32(len(arr)); i++ {
		rightValue := i + 1
		fmt.Println(rightValue)
		fmt.Println(arr)
		fmt.Println(valueIndexMap)

		if arr[i] != rightValue {
			fmt.Printf("Swap %d <-> %d\n", arr[i], rightValue)
			swaps++

			wrongValue := arr[i]
			arr[i] = rightValue
			oldRightValueIndex := valueIndexMap[rightValue]
			arr[oldRightValueIndex] = wrongValue

			valueIndexMap[wrongValue] = valueIndexMap[rightValue]
		}

		fmt.Println(arr)
		fmt.Println(valueIndexMap)
		fmt.Println()
	}

	return swaps
}
