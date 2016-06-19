package main

import "fmt"

func main() {
	input := []int{1, 1, 1}

	fmt.Println(checkFlip(input))
}

func checkFlip(a []int) []int {
	var ret []int
	var ans int

	n := len(a)

	for startIndex := 0; startIndex < n; startIndex++ {
		potential := 0
		for endIndex := 0; endIndex < n+1; endIndex++ {
			if startIndex+endIndex > n-1 {
				break
			}
			if a[startIndex+endIndex] == 0 {
				potential++
			} else {
				potential--
			}
			fmt.Println(potential, []int{startIndex, startIndex + endIndex})
			if potential > ans {
				ans = potential
				ret = []int{startIndex, startIndex + endIndex}
			}
		}
	}
	return ret
}
