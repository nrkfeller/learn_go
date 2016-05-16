package main

import "fmt"

func main() {
	// utf8 is a 4 byte encoding scheme
	word := "potato mafia"

	// "casts" the string into a byte array
	fmt.Println([]byte(word))

	bites := [...]byte{112, 111, 116, 97, 116, 111, 32, 109, 97, 102, 105, 97}
	for i := range bites {
		fmt.Println(string(bites[i]))
	}
}
