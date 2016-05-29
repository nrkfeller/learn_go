package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)

	go takefromchan(c)

	for i := 0; i < 2; i++ {
		go closechan(c, done)

	}

	<-done
	<-done
}

func closechan(c chan int, done chan bool) {
	for i := range c {
		fmt.Println(i)
	}
	done <- true
}

func takefromchan(c chan int) {
	for k := 0; k < 10000; k++ {
		c <- k
	}
	close(c)
}
