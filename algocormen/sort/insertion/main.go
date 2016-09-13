package main

import "fmt"

func main() {
	input := []int{6, 5, 3, 7, 8, 5, 3, 1, 2}
	fmt.Println(is(input))
}

func is(a []int) []int {
	for k, v := range a {
		if k == 0 {
			continue
		} else {
			current := v
			i := k - 1
			for i >= 0 && current > a[i] { // change < > for increasing and decreasing
				a[i+1] = a[i]
				i--
			}
			a[i+1] = current
		}
	}
	return a
}
