package main

import "fmt"

func main() {

	var person string

	fmt.Scan(&person)

	switch person {
	case "Adam":
		fmt.Println(" and Eve")
	case "Nick":
		fmt.Println(" father christmas")
	case "Sammy":
		fmt.Println(" Columbus")
	}
}
