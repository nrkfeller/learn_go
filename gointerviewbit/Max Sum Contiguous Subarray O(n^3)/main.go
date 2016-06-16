package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSum(input))
}

func maxSum(a []int) []int {
	answer := math.MinInt64
	var retslice []int
	for i := 0; i < len(a); i++ {
		slicelen := len(a) - i
		for slicelen > 0 {
			myslice := a[i : i+slicelen]
			total := sums(myslice)
			if total > answer {
				answer = total
				retslice = myslice
			}
			slicelen--
		}
	}
	return retslice
}

func sums(myS []int) int {
	total := 0
	for _, i := range myS {
		total += i
	}
	return total
}

// package main
//
// import "fmt"
//
// func main() {
// 	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
//
// 	fmt.Println(maxSum(input))
// }
//
// type ans struct {
// 	thetotal int
// 	theslice []int
// }
//
// func maxSum(a []int) []int {
// 	allsums := []ans{}
// 	for i := 0; i < len(a); i++ {
// 		slicelen := len(a) - i
// 		for slicelen > 0 {
// 			myslice := a[i : i+slicelen]
// 			total := sums(myslice)
// 			myans := ans{total, myslice}
// 			allsums = append(allsums, myans)
// 			slicelen--
// 		}
// 	}
// 	return max(allsums)
// }
//
// func sums(myS []int) int {
// 	total := 0
// 	for _, i := range myS {
// 		total += i
// 	}
// 	return total
// }
//
// func max(allsums []ans) []int {
// 	var final ans
// 	for _, i := range allsums {
// 		if i.thetotal > final.thetotal {
// 			final = i
// 		}
// 	}
// 	return final.theslice
// }
