package main

import "fmt"

func main() {

	avg := average(213, 54, 6554, 6, 642, 5, 645, 6, 654, 6)

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
