package main

import "fmt"

func main() {
	defer world()
	hello()
}

func hello() {
	fmt.Print("hello")
}

func world() {
	fmt.Print("world")
}
