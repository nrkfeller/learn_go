package main

import "fmt"

func main() {
	sl := make([]int, 1, 1)
	sl[0] = 10
	sl[0] += 10 // python inc
	sl[0]++     // other langs inc
	fmt.Print(sl)
}
