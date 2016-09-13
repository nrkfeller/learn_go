package main

import "fmt"

func main() {
	input := []int{3, 4, 5, 2, 8}
	fmt.Println(search(input, 12))
}

func search(l []int, a int) interface{} {
	for k, v := range l {
		if a == v {
			return k
		}
	}
	return nil
}
