package main

import "fmt"
import "unsafe"

const (
	// iota scope starts
	a = iota
	b = iota << 1
	c = iota << 1
)

func main() {
	// bit shifted iota
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%d\n", unsafe.Sizeof(a))
}
