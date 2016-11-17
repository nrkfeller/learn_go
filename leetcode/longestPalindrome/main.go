package main

import "fmt"

func main() {
	input := "abccccdd"
	fmt.Println(longestPalindrome(input))
}

func longestPalindrome(s string) int {
	letterBank := make(map[rune]int)

	for _, v := range s {
		letterBank[v]++
	}

	var odd int
	var count int

	for _, v := range letterBank {
		if v%2 == 0 {
			count += v
		} else {
			odd = 1
			if v > 2 {
				count += v - 1
			}
		}
	}
	return count + odd
}
