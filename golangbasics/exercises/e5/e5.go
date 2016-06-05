package main

import "fmt"

func f(n int) {
	for i := 0; i < 100; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)
	go f(1)
	go f(2)
	var input = "yao"
	fmt.Scanln(&input)
}
