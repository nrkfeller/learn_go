package main

import "fmt"

func main() {

	greet := makeGreeter()
	greet()

	greet2 := makeGreeter2()
	fmt.Println(greet2())

}

func makeGreeter() func() {
	return func() { fmt.Println("yoyoyo") }
}

func makeGreeter2() func() string {
	return func() string {
		return "hihihi"
	}
}
