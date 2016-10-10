package main

import "fmt"

func main() {
	input := [][]int{[]int{1, 1, 1, 1, 1}, []int{2, 2, 0, 2, 2}, []int{3, 3, 3, 3, 3}, []int{4, 4, 4, 4, 4}, []int{5, 5, 5, 5, 0}}
	szm(input)
	fmt.Println(input)
}

func szm(a [][]int) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == 0 {
				defer szmhelper(a, i, j)
			}
		}
	}
}

func szmhelper(a [][]int, r, c int) {
	n := len(a)
	for i := 0; i < n; i++ {
		a[r][i] = 0
		for j := 0; j < n; j++ {
			a[j][c] = 0
		}
	}
}
