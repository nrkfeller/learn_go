package main

import "fmt"

func f(n int) {
	for i := 0; i < 100; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)
	go f(1)
	go f(2)

	// this is here to prevent the main function from exiting
	var input string
	fmt.Scanln(&input)
}
