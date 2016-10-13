package main

import "fmt"

func main() {
	input := []int{10, 22, 9, 33, 21, 50, 41, 60}
	fmt.Println(lis(input))
}

func lis(a []int) int {

	output := make([]int, len(a))

	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i] > a[j] && output[i] < output[j]+1 {
				output[i] = output[j] + 1
			}
		}
	}

	maximum := 0

	for i := 0; i < len(a); i++ {
		if maximum < output[i] {
			maximum = output[i]
		}
	}

	return maximum
}
