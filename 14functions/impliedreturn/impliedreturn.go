package main

import "fmt"

func main() {
	fmt.Println(plural("potato"))
}

func plural(x ...interface{}) (s string) {
	s = fmt.Sprint(x[0]) + "s"
	return
}
