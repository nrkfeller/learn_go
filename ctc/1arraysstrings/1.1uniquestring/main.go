package main

import "fmt"

func main() {
	input := "ahgsdafb"
	fmt.Println(unique(input))
}

func unique(a string) bool {
	allchar := make([]int, 256)
	for _, v := range a {
		allchar[v]++
		if allchar[v] > 1 {
			return false
		}
	}
	return true
}
