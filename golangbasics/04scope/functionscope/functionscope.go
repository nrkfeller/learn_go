package main

import "fmt"

var x int

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())

	var y int
	decrement := func() int {
		y--
		return y
	}

	fmt.Println(decrement())
	fmt.Println(decrement())

	returnedfunc := funcreturn()
	fmt.Println(returnedfunc())
	fmt.Println(returnedfunc())

}

// Function that returns a func that returns a float64
func funcreturn() func() float64 {
	x := 2.5
	return func() float64 {
		x = x * x
		return x
	}
}
