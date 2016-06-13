package main

import "fmt"

func main() {
	ans := performadd(func(a int, b int) int {
		return a + b
	})
	fmt.Println(ans)
}

type Add func(a int, b int) int

func performadd(adder Add) int {
	return adder(123, 654)
}
