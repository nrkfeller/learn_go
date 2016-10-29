package main

import "fmt"

func main() {
	n := []int{1, 2, 2, 1}
	m := []int{2, 2}

	fmt.Println(arrint(n, m))
}

func arrint(a, b []int) []int {
	var answer []int

	itemBank := make(map[int]int)

	for _, v := range a {
		itemBank[v] = 1
	}

	for _, v := range b {
		if itemBank[v] == 1 {
			answer = append(answer, v)
			itemBank[v] = 0
		}
	}
	return answer
}
