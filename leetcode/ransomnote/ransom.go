package main

import "fmt"

func main() {

	fmt.Println(canConstruct("aa", "aab"))
	fmt.Println(canConstruct("aa", "ab"))
}

func canConstruct(b, a string) bool {
	letterBank := make(map[rune]int)

	for _, v := range a {
		letterBank[v]++
	}

	for _, v := range b {
		letterBank[v]--
		if letterBank[v] < 0 {
			return false
		}
	}
	return true
}
