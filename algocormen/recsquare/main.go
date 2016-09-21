package main

import "fmt"

func main() {
	fmt.Println(sq(4, 4))
	fmt.Println(recsq(4, 4))
}

func sq(x, n int) int {
	ans := 1
	for i := 0; i < n; i++ {
		ans *= x
	}
	return ans
}

func recsq(x, n int) int {
	if n == 1 {
		return x
	}
	return x * recsq(x, n-1)
}
