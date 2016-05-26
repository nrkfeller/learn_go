package main

import "fmt"

func main() {

	p1 := person{"nick", "feller", 26}
	fmt.Println(p1)

	p2 := person{"kamal", "sehdev", 40}
	fmt.Println(p2)

}

type person struct {
	first  string
	second string
	age    int
}
