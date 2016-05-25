package main

import "fmt"

func main() {

	m := 1236

	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	fmt.Println(half(m))
}
