package main

import "fmt"

func main() {

	mya := []float64{213, 54, 6554, 6, 642, 5, 645, 6, 654, 6}

	// unpack this array and pass it
	avg := average(mya...)

	fmt.Println(avg)

}

func average(x ...float64) (f float64) {

	fmt.Println(x) // became a slice

	f = 0

	// use of underscore bc we dont care about the index
	for _, v := range x {
		f += v / float64(len(x))
	}

	return
}
