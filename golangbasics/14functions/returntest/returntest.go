package main

import "fmt"

func main() {
	fmt.Println(gretting("nick", "feller"))
}

func gretting(x ...interface{}) string {
	return fmt.Sprint(x[0], x[1])
}
