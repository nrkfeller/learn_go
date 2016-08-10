package main

import "fmt"

func main() {
	out := make(chan string)
	done := make(chan bool)

	go lame(out)
	go closechan(out, done)

	<-done

}

func lame(c chan string) {
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("%d", i)
		c <- res
	}
	close(c)

}

func closechan(c chan string, done chan bool) {
	for i := range c {
		fmt.Println(i)
	}
	done <- true
}
