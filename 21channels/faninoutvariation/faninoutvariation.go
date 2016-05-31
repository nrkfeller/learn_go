package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3, 4, 5, 6, 6)

	// FAN OUT
	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// FAN IN
	// Consume the merged output from multiple channels.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

// Takes numbers and pushes them onto a channel
func gen(nums ...int) chan int {
	fmt.Printf("TYPE OF NUMS %T\n", nums) // just FYI

	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// this is the computation whatever it is
func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// takes in a variadic slice of channels and pushed them into one channel
// essentially consolidating the computed results
func merge(cs ...chan int) chan int {
	fmt.Printf("TYPE OF CS: %T\n", cs) // just FYI

	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	// get each channel
	for _, c := range cs {
		// get each value (in each channel)
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c) // takes in a channel!
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
