package main

import "fmt"

func main() {
	input := []int{1, 1, 1, 1, 1, 2, 2, 2, 2}

	fmt.Println(removedupes(input))
}

func removedupes(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	r := 0
	for k, v := range a {
		for i, j := range a {
			if j == v && k != i {
				a[r], a[k] = a[k], a[r]
				r++
				break
			}
		}
	}
	return a[r:]
}
