package main

import "fmt"

func main() {

	x := 5
	// we are only passing a value!
	zero(x)
	fmt.Println(x)

}

func zero(x int) {
	x = 0
}
