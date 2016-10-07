package main

import "fmt"

func main() {
	input := [][]int{[]int{1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2}, []int{3, 3, 3, 3, 3}, []int{4, 4, 4, 4, 4}, []int{5, 5, 5, 5, 5}}
	rot90(input)
	fmt.Println(input)
}

func rot90(a [][]int) {
	n := len(a)
	if n == 1 {
		return
	}
	for i := 0; i < n; i++ {
		a[0][i], a[n-1-i][n-1] = a[n-1-i][n-1], a[0][i]
	}
	rot90(a[1:])
}
