package main

import "fmt"

func main() {

	input := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(shift(input))
}

func shift(s []int) []int {
	return append(s[1:len(s)], s[0])
}
