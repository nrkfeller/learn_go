package main

import "fmt"

func main() {
	var name interface{} = "Nicolas"

	_, ok := name.(string)
	if ok {
		fmt.Println("its a string")
	} else {
		fmt.Println("not a string")
	}
}
