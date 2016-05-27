package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for k := 0; k < 10000; k++ {
			c <- k
		}
		close(c)
	}()

	go func() {
		for i := range c {
			fmt.Println(i)
		}
		done <- true
	}()

	go func() {
		for i := range c {
			fmt.Println(i)
		}
		done <- true
	}()

	<-done
	<-done
}
