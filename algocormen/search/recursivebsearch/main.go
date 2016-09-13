package main

import "fmt"

func main() {
	input := []int{1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10, 11}
	recbin(input, 2)
	recbin(input, 3)
	recbin(input, 4)
	recbin(input, 22)
	recbin(input, 0)
	recbin(input, 10)
}

func recbin(l []int, s int) {
	if len(l) == 1 && l[0] != s {
		fmt.Println(false)
	} else if l[len(l)/2] == s {
		fmt.Println(true)
	} else if l[len(l)/2] > s {
		recbin(l[:len(l)/2], s)
	} else {
		recbin(l[len(l)/2:], s)
	}
}
