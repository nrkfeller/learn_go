package main

import (
	"fmt"
	"time"
)

func main() {
	inc := make(chan int)

	go func() {
		inc <- 1
	}()

	go func() {
		count := 0
		for i := 0; i < 100000; i++ {
			count++
			go increment(inc, count)
		}
	}()

	time.Sleep(time.Second)
}

func increment(c chan int, count int) chan int {
	num := <-c
	num++
	fmt.Println(count, ":", num)
	c <- num
	return c
}
