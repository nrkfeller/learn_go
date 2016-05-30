package main

import "fmt"

func main() {
	out := sq(gen(2, 3, 4, 5, 6))

	for i := range out {
		fmt.Println(i)
	}
}

// pushed the numbers into the channel
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// pops the data and squares it
func sq(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
