package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 30; i++ {
		fmt.Print("foo")
		//time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 30; i++ {
		fmt.Print("bar")
		//time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}
