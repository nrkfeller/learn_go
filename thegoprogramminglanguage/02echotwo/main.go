package main

import (
	"fmt"
	"os"
)

func main() {
	sep := " "
	output := ""
	for _, i := range os.Args[1:] {
		output += i + sep
	}
	fmt.Println(output)
}
