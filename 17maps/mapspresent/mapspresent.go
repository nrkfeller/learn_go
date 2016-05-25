package main

import "fmt"

func main() {
	m1 := map[int]string{
		1: "nick",
		2: "abel",
		3: "emile",
	}

	delete(m1, 1)

	if v, ok := m1[1]; ok {
		fmt.Println("first position exists, is : ", v)
	} else {
		fmt.Println("there is no 1")
	}
}
