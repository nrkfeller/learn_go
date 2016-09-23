package main

import "fmt"

func main() {
	large := []int{1, 2, 3, 4, 5}
	small := []int{8, 9, 10}
	fmt.Println(copy(large[0:4], small[:]))
}
