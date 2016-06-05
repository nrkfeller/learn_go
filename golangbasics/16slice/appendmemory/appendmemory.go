package main

import "fmt"

func main() {
	// show we are just working with dynamic array
	sl := make([]int, 10, 20)

	for i := 0; i < 80; i++ {
		fmt.Println(cap(sl), len(sl))
		sl = append(sl, i)
	}
}
