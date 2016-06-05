package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	// multiple pulls from same channel - in
	f1 := factorial(in)
	f2 := factorial(in)
	f3 := factorial(in)
	f4 := factorial(in)
	f5 := factorial(in)
	f6 := factorial(in)
	f7 := factorial(in)
	f8 := factorial(in)
	f9 := factorial(in)
	f0 := factorial(in)

	// merging and printing - "range" also holds the main hanging
	for n := range merge(f1, f2, f3, f4, f5, f6, f7, f8, f9, f0) {
		fmt.Println(n)
	}
}

// cleaner merge function
func merge(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))

	for _, cee := range cs {
		go output(cee)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// func merge(chans ...chan int) chan int {
// 	out := make(chan int)
//
// 	var wg sync.WaitGroup
// 	wg.Add(len(chans))
//
// 	for _, c := range chans {
// 		go func(ch chan int) {
// 			for n := range ch {
// 				out <- n
// 			}
// 			wg.Done()
// 		}(c)
// 	}
//
// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()
// 	return out
// }

func gen() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100000; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(c chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(i int) int {
	sum := 1
	for i := i; i > 0; i-- {
		sum *= i
	}
	return sum
}
