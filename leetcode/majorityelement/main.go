package main

import "fmt"

func main() {
	input := []int{5, 6, 6}

	fmt.Println(majority(input))
}

//func majority(nums []int) int {
//	majmap := make(map[int]int)
//
//	for _, v := range nums {
//		majmap[v]++
//	}
//
//	var maj int
//	var val int
//	for k, v := range majmap {
//		if v > maj {
//			maj = v
//			val = k
//		}
//	}
//	return val
//}

func majority(nums []int) int {
	count, major := 1, nums[0]

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			count++
			major = nums[i]
		} else if major == nums[i] {
			count++
		} else {
			count--
		}
	}
	return major
}
