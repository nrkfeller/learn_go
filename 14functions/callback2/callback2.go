package main

import "fmt"

func filter(nums []int, callback func(int) bool) []int {
	var xs []int

	for _, n := range nums {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println((xs))
}
