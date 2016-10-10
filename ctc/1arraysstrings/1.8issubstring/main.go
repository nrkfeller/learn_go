package main

import "fmt"

func main() {
	a := "potato"
	b := "topota"

	fmt.Println(isSubstring(a, b))
}

func isSubstring(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	concata := a + a
	for i := 0; i < len(concata)-len(a)+1; i++ {
		if concata[i:i+len(a)] == b {
			return true
		}
	}
	return false
}
