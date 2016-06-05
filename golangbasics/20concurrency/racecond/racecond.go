package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {

	wg.Add(2)
	go inc("foo")
	go inc("bar")
	wg.Wait()
	fmt.Println(counter)

}

func inc(s string) {
	for i := 0; i < 10; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, counter)
	}
	wg.Done()
}
