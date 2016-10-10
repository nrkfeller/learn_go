package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	start := 0
	end := 10

	for k := start; k < end; k++ {
		fmt.Print(a[k], " ")
	}
	fmt.Println()

	for _, v := range a[start:end] {
		fmt.Print(v, " ")
	}

}
