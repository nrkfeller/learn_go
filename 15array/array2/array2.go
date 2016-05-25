package main

import "fmt"

func main() {
	var x [256]byte

	var i byte = 97

	for k := range x {
		x[k] = i
		i++
	}

	for _, v := range x {
		fmt.Printf("%c - %T - %b \n", v, v, v)
	}
}
