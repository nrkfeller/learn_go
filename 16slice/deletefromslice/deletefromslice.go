package main

import "fmt"

func main() {
	sl1 := []int{1, 2, 3, 3, 4, 5, 6, 7, 8}
	fmt.Println(append(sl1[:3], sl1[4:]...))
}
