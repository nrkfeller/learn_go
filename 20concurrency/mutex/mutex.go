package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex
var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(5)
	go inc("foo")
	go inc("bar")
	go inc("man")
	go inc("bun")
	go inc("zoo")
	wg.Wait()
	fmt.Println("final Value", counter)
}

func inc(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, counter)
		mutex.Unlock()
	}
	wg.Done()
}
