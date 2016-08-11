// Ref: https://blog.golang.org/pipelines
package main

import (
	"fmt"
	"sync"
	"time"
)

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				fmt.Print("gen -- done")
				return

			}
		}
	}()
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				fmt.Print("sq -- done")
				return
			}
		}
	}()
	return out
}

func merge(done <-chan struct{}, cs []<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {

	done := make(chan struct{})
	defer close(done)

	// Distribute the sq work across two goroutines that both read from in.
	var arr []int
	for i := 0; i < 100000; i++ {
		arr = append(arr, i)
	}

	in := gen(done, arr...)
	var allcs []<-chan int

	start := time.Now()
	for i := 0; i < 3; i++ {
		c := sq(done, in)
		allcs = append(allcs, c)

	}
	total := 0
	for n := range merge(done, allcs) {
		//fmt.Println(n)
		total += n
	}
	fmt.Println(time.Since(start), total)

}
