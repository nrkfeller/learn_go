package main

import (
	"fmt"
	"math"
)

func main() {
	a := "AA"
	b := "AB"
	c := "AC"
	d := "BA"
	fmt.Println(cn(a))
	fmt.Println(cn(b))
	fmt.Println(cn(c))
	fmt.Println(cn(d))
}

func cn(s string) int {

	num := 0
	l := float64(len(s))

	for _, v := range s {
		num += (int(v) - 64) * int(math.Pow(26.0, l-1))
		l--
	}

	return num
}
