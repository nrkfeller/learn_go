package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {

	wg.Add(4)
	go inc("foo")
	go inc("bar")
	go inc("bar")
	go inc("bar")
	wg.Wait()
	fmt.Println(counter)

}

func inc(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, counter)
	}
	wg.Done()
}
