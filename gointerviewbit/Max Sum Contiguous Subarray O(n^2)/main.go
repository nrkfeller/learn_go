package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSum(input))
}

func maxSum(a []int) []int {
	ans := math.MinInt64
	var arr []int
	for startindex := 0; startindex < len(a); startindex++ {
		sum := 0
		for subarraysize := 1; subarraysize <= len(a); subarraysize++ {
			if startindex+subarraysize > len(a) {
				break
			}
			sum += a[startindex+subarraysize-1]
			if sum > ans {
				ans = sum
				arr = a[startindex : startindex+subarraysize]
			}
		}
	}
	return arr
}
