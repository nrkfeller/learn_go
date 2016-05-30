package main

import "fmt"

func main() {
	c := make(chan int)

	fact := func(n int, c chan int) {
		sum := 1
		for i := n; i > 1; i-- {
			sum *= i
		}
		c <- sum
		//close(c)
	}

	// runs on all cores
	go fact(20, c)

	fmt.Println(<-c)
}
