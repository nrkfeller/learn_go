package main

import "fmt"

func main() {
	input := 9
	fmt.Println(krpt(input))
}

func krpt(a int) []int {
	ans := make([]int, a+1)

	for k := range ans {
		ans[k] = fact(a) / (fact(k) * fact(a-k))
	}

	return ans
}

func fact(i int) int {
	if i == 0 {
		return 1
	}
	return fact(i-1) * i
}
