package main

import "fmt"

func main() {
	c := make(chan int)     // channel that hold information
	done := make(chan bool) // channel that controls flow

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// this needs to be in a go routine because otherwise we would never get to c and deadlock!
	go func() {
		<-done
		<-done // throw the bools away
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}
