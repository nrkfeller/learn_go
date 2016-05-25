package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["nick"] = 100
	m["abel"] = 60
	m["emile"] = 40

	fmt.Println(m)

	delete(m, "abel")

	fmt.Println(m)

	// check is present and its value

	v, ok := m["nick"]
	fmt.Println(v, ok)
}
