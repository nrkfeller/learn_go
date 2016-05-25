package main

import "fmt"

func main() {
	myslice1 := make([]int, 10, 20)  // type, length, capacity
	equivArray := new([20]int)[0:10] // equivalent to above slice

	myslice2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(myslice1, equivArray)
	fmt.Println(myslice2[:3]) // slice up that slice
	fmt.Println(myslice2[3:6])
	fmt.Println(myslice2[len(myslice2)-1])
	fmt.Println(cap(myslice1), len(myslice1))
}
