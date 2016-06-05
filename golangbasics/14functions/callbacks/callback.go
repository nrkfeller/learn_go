package main

import "fmt"

// callback is passing a func as an argument

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n * 2)
	})
}

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}
