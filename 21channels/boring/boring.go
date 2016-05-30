package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Samy"), boring("Nick"), boring("Toms"), boring("Paul"))

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Donedonedone!")
}
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(i1, i2, i3, i4 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-i1
		}
	}()
	go func() {
		for {
			c <- <-i4
		}
	}()
	go func() {
		for {
			c <- <-i3
		}
	}()
	go func() {
		for {
			c <- <-i2
		}
	}()
	return c
}
