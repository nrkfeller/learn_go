package main

import "fmt"

func main() {
	input := [][]int{{1, 3}, {2, 6}, {2, 10}, {15, 18}}

	fmt.Println(moi(input))
}

func moi(a [][]int) [][]int {
	for block := 0; block < len(a)-1; block++ {
		for a[block][1] > a[block+1][0] {
			a[block][1] = a[block+1][1]
			a = append(a[:block+1], a[block+2:]...)
		}
	}
	return a
}
