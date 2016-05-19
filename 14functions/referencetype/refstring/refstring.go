package main

import "fmt"

func main() {

	m := make([]string, 1, 25)
	fmt.Println(m)

	// actually changes m!
	change(m)
	fmt.Println(m)

}

func change(s []string) {
	s[0] = "potato"
}
