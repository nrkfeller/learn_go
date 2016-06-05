package main

// fallthrough makes the next case evaluate to true

import "fmt"

func main() {
	s := "hello"

	switch s {
	case "hello":
		fmt.Println("english")
		fallthrough
	case "hola":
		fmt.Println("spanish")
	case "guttentag":
		fmt.Println("german")
	}
}
