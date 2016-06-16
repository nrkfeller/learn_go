package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 9}

	fmt.Println(addOne(input))
}

func addOne(a []int) []string {
	order := len(a) - 1
	var num int
	for i := 0; i < len(a); i++ {
		num += a[i] * intPow(10, order)
		order--
	}
	num++

	ans := strconv.Itoa(num)
	ret := strings.Split(ans, "")

	return ret

}

func intPow(num int, power int) int {
	ans := 1
	for i := 0; i < power; i++ {
		ans = ans * num
	}
	return ans
}
