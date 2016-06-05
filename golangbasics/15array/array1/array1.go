package main

import "fmt"

func main() {
	var x [56]string
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(len(x))

	for i := 65; i < len(x)+65; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)

	var xx [2][3]int
	fmt.Println(xx)
	fmt.Println(&xx)
	fmt.Println(len(xx))

}
