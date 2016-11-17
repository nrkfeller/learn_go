package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{1, 2, 3, 4}
	fmt.Println(allsubsets(input))
}

func allsubsets(a []int) int {
	num := int(math.Pow(2, float64(len(a))))
	return num
}
