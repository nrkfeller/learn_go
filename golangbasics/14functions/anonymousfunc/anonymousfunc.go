package main

import "fmt"

func main() {

	test := 123

	double := func(x int) int {
		return x * 2
	}

	fmt.Println(double(test))
	fmt.Printf("%T \n", double)

}
