package main

import "fmt"
import "strings"

func main() {

	fmt.Println(fullname("nick", "feller"))

}

func fullname(f, l string) (string, string) {
	return strings.ToUpper(f), strings.ToUpper(l)
}
