package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []float64{6, 7, 8}
	foo(aSlice...) // unpack the slice
	foo()
}

func foo(x ...float64) (avg float64) {
	avg = 0
	l := float64(len(x))
	for _, v := range x {
		avg += v / l
	}
	fmt.Println(avg)
	return
}
