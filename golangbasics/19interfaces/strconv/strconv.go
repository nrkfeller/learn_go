package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 32

	b := "55"

	i, _ := strconv.Atoi(b)

	fmt.Println(i + a)

	k := strconv.Itoa(a)

	fmt.Println(b + k)
}
