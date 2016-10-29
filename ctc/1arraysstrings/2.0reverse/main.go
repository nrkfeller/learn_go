package main

import "fmt"

func main() {
	input := "hello"
	fmt.Println(reverseString(input))
}

func reverseString(s string) string {
	var ans string
	for i := len(s) - 1; i >= 0; i-- {
		ans += string(s[i])
	}
	return ans
}
