package main

import "fmt"

func main() {

	// make unburrefer channel. meaning there is not buffer size. unlimited channel size
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// push i to the channel
			c <- i
		}
		// block other information from entering channel
		close(c)
		// not closing the channel created deadlock -- range wont be able to iterate
	}()

	// iterate through each element on the channel, save it to i
	for i := range c {
		// print i
		fmt.Println(i)
	}
}
