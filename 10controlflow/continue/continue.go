package main

import "fmt"

func main() {

	var i byte

	for i < 10 {
		i++
		if i%2 == 0 {
			continue
		} else {
			fmt.Print(i)
		}
	}
}
