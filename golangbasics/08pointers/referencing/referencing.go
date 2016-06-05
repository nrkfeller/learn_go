package main

import "fmt"

//import "unsafe"

func main() {
	var a int8 = 9
	b := 40

	var x = &a
	var y = &b

	// x and y are pointers to ints -- *int
	fmt.Printf("x and y are of type : %T - %T \n", x, y)

	fmt.Printf("x and y are pointing to address : %x - %x which contains the values : %d - %d", x, y, *&a, *y)
}
