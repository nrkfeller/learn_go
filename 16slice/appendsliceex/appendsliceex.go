package main

import "fmt"

//import "builtin"

func main() {

	sl1 := []int{1, 2, 3, 4, 5}
	sl2 := []int{6, 7, 8, 9}

	asl := append(sl1, sl2...)
	fmt.Println(asl) // we need to unpack the slice to append it

}
