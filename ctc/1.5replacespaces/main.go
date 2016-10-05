package main

import "fmt"

func main() {
	input := "hello its me adele"
	fmt.Println(replacespace(input))
}

func replacespace(a string) string {
	ba := []byte(a)

	return string(ba)
}
