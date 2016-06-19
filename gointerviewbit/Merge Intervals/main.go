package main

import "fmt"

func main() {
	//input := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	//mergeVal := []int{4, 9}
	input := [][]int{{1, 3}, {6, 9}}
	mergeVal := []int{2, 5}

	fmt.Println(mergeIntervals(input, mergeVal))
}

func mergeIntervals(a [][]int, mi []int) [][]int {
	swap := false
	for block := 0; block < len(a); block++ {

		if a[block][0] < mi[0] && a[block][1] > mi[0] {
			swap = true
			a[block][1] = mi[1]
		}

		if a[block][1] < mi[1] && swap {
			a = append(a[:block], a[block+1:]...)
		}

		if a[block][0] < mi[1] && a[block][1] > mi[1] && swap {
			temp := a[block][1]
			a[block-1][1] = temp
			a = append(a[:block], a[block+1:]...)
		}

	}

	return a

}
