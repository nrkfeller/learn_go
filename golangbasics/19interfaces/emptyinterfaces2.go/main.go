package main

import "fmt"

func main() {
	add(3, 6, 8)
}

func add(a interface{}, b interface{}, c interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("a is an int")
	default:
		fmt.Println("invalid argument")
	}
	switch b.(type) {
	case int:
		fmt.Println("b is an int")
	default:
		fmt.Println("invalid argument")
	}
	switch c.(type) {
	case int:
		fmt.Println("c is an int")
	default:
		fmt.Println("invalid argument")
	}
}
