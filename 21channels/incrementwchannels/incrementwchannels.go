package main

import "fmt"

func main() {
	c1 := inc("Foo:")
	c2 := inc("Bar:")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("Final Count : ", <-c3+<-c4)
}

// both calls are incrementing this function!
func inc(s string) chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 30; i++ {
			c <- 1
			fmt.Println(s, i)
		}
		close(c)
	}()
	return c
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
