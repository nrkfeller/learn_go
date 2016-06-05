package main

import "fmt"

func main() {
	in := gen()

	f := factorial(in)

	for n := range f {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 20; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range c {
			out <- fact(i)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	sum := 1
	for i := n; i > 0; i-- {
		sum *= i
	}
	return sum
}
