package main

import "fmt"

func main() {
	a := 333
	var b = &a
	*b = 222
	fmt.Println(a)
}
