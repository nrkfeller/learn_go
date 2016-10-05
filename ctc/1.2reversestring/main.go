package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5, 6}
	reverse(input)
	fmt.Println(input)
}

func reverse(a []int) {
	l := len(a)
	for i := 0; i < l/2; i++ {
		a[i], a[l-i-1] = a[l-i-1], a[i]
	}
}
