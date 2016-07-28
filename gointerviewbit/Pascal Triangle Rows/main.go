package main

import "fmt"

func main() {
	input := 9
	fmt.Println(ptr(input))
}

func ptr(a int) [][]int {
	total := make([][]int, a)
	for k := range total {
		total[k] = pascalrow(k)
	}
	return total
}

func pascalrow(a int) []int {
	row := make([]int, a+1)

	for i := range row {
		row[i] = fact(a) / (fact(i) * fact(a-i))
	}
	return row
}

func fact(a int) int {
	if a == 0 {
		return 1
	}
	return fact(a-1) * a
}
