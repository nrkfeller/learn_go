package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{123, 45, 6, 56, 54, 6, 7, 7, 7, 575, 4, 6}

	sort.Ints(nums)
	fmt.Println(nums)
}
