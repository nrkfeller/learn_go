package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"nick", "tu", "abe", "jin"}

	// using the itnerface to reverse sort strings
	// essentially first it is cast to StringSlice, then reverse gives the "sort" and interface by which it does a reversed sort
	sort.Sort(sort.Reverse(sort.StringSlice(names)))

	fmt.Println(names)

	nums := []int{12, 43, 254, 5, 6, 56, 5, 6, 4, 434}

	// using interface to reverse sort ints
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	fmt.Println(nums)
}
