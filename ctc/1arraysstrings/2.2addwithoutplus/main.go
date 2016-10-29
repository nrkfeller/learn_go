package main

import "fmt"

func main() {
	fmt.Println(add(5, 12))
}

func add(x, y int) int {
	if y == 0 {
		return x
	}
	return add(x^y, (x&y)<<1)
}
