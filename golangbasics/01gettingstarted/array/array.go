package main

import "fmt"
import "unsafe"

func main() {

	var mylist [10]int
	mylist[0] = 1
	mylist[1] = 3
	mylist[4] = 5
	fmt.Println(mylist)

	caplist := [10]int{1, 2, 3}
	fmt.Println(caplist, "has capacity", len(caplist), cap(caplist))

	mylist2 := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(mylist2)

	mylist3 := mylist2[3:]
	fmt.Println(mylist3)

	mylist4 := make([]byte, 5)
	mylist4[0] = 254
	fmt.Println(mylist4[0] + 12)

	fmt.Printf("%d\n", unsafe.Sizeof(mylist3[0]))

}
