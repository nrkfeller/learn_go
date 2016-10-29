package main

import "fmt"

func main() {
	input := []int{1, 1, 3, 2, 3, 2, 4, 5, 5, 6, 7, 6, 7}
	fmt.Println(singleNumber(input))
}

func singleNumber(nums []int) int {
	mm := make(map[int]int)

	for _, v := range nums {
		mm[v]++
	}
	for k, v := range mm {
		if v == 1 {
			return k
		}
	}
	return -1
}
