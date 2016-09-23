package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	swap(a[0:2])
	fmt.Println(a)
}

func swap(a []int) {
	a[1], a[0] = a[0], a[1]
}
