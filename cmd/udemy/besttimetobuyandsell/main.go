package main

import (
	"fmt"
	"math"
)

func main() {
	//nums := generate.UnsortedPositiveIntArray(10, 100)
	//nums := []int{2,1,1,2}
	//nums := []int{1,2,3,1}
	nums := []int{7, 1, 5, 3, 6, 4}
	valueRobbed := maxProfit(nums)
	fmt.Println(valueRobbed)
}

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := math.MaxInt

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}
