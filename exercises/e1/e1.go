package main

import "fmt"

func main() {

	m := 1236

	fmt.Println(half(m))
}

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}
