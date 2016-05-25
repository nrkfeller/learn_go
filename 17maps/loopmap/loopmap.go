package main

import "fmt"

func main() {
	ele := map[string]string{
		"O": "oxygen",
		"N": "nitrogen",
		"C": "carbon",
	}
	for k, v := range ele {
		fmt.Println(k, "is for", v)
	}
}
