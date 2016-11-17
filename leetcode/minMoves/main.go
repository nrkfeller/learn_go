package main

import "fmt"

const MaxInt = int(^uint(0) >> 1)

func main() {
	input := []int{1, 2, 3}
	fmt.Println(minMoves(input))
}

func minMoves(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	fmt.Println(MaxInt)
	min := MaxInt
	var sum int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		min = minimum(min, nums[i])
	}
	return sum - len(nums)*min
}

func minimum(a, b int) int {
	if a > b {
		return b
	}
	return a
}
