package main

import "fmt"

func main() {
	me := Person{firstname: "John", lastname: "Doe", age: 40}
	fmt.Println(fullname(&me))
}

// Person is just a dude
type Person struct {
	firstname string
	lastname  string
	age       int
}

// use pointer if the receriver needs to be modified
func fullname(p *Person) string {
	return p.firstname + p.lastname
}
