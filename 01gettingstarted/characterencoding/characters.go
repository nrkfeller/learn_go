package main

import "fmt"

func main() {

	words := "hello its me"

	for i := 0; i < len(words); i++ {
		fmt.Printf("%q", words[i])
	}
}
