package main

import "fmt"

const (
	// iota scope starts
	a = iota
	b = iota
	c = iota
)

func main() {
	// iota scope restarts
	const (
		d = iota
		_ = iota
		e = iota
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("threw away iota of 1", e)
}
