package main

import "fmt"

func main() {
	input := []int{1, 2, 5, -1, 1, 2, 3, 4, 5, 5, -7, 3, 5, 6}
	fmt.Println(mnnsa(input))
}

func mnnsa(a []int) []int {
	largest := 0
	var ret []int

	for i := 0; i < len(a); i++ {
		sum := 0
		start := i
		for a[i] >= 0 {
			sum += a[i]
			i++
			if i >= len(a) {
				break
			}
		}
		if sum > largest {
			largest = sum
			ret = a[start:i]
		}
	}
	return ret
}
