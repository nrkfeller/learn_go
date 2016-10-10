package main

import (
	"fmt"
	"math"
)

func main() {
	// {5, 10, 3, 12, 5, 50, 6}
	input := []int{5, 10, 3, 12, 5}
	fmt.Println(mcm(input, 1, len(input)-1))
}

func mcm(a []int, start, end int) int {
	if start == end {
		return 0
	}

	min := math.MaxInt64

	for k := start; k < end; k++ {
		count := mcm(a, start, k) + mcm(a, k+1, end) + a[start-1]*a[k]*a[end]
		if count < min {
			min = count
		}
	}

	return min
}
