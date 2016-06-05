package main

import (
	"fmt"
	"time"
)

func main() {

	// unbuffered channel --- c := make(chan int,10)
	c := make(chan int)

	go func() {
		for i := 0; i < 100000; i++ {
			// pack values into the c channel
			c <- i
			// this code stops until something else takes the value off the channel
		}
	}()

	// does a simple read of the channel
	go func() {
		for {
			// grab last item off the c channel
			fmt.Println(<-c)
			// after this, the free'd portion of the channel is free to take the next value
		}
	}()

	// this sleep is needed because once the main is over the program quits
	time.Sleep(time.Second)
}
