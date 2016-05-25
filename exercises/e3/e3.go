package main

import "fmt"

func main() {
	fmt.Println(largest(1, 23, 3, 4, 5, 54, 4, 33, 333, 23, 4))
}

func largest(x ...int) (l int) {
	l = x[0]
	for _, v := range x {
		if l < v {
			l = v
		}
	}
	return l
}
