package main

import "fmt"

func main() {
	input := []int{1, 2, 3}
	p(input, 0, len(input)-1)
}

func p(a []int, l, r int) {
	if l == r {
		fmt.Println(a)
	} else {
		for i := l; i < r+1; i++ {
			a[l], a[i] = a[i], a[l]
			p(a, l+1, r)
			a[l], a[i] = a[i], a[l]
		}
	}
}
