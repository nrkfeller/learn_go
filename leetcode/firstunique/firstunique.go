package main

import "fmt"

func main() {
	a := "leetcode"

	b := "loveleetcode"
	fmt.Println(fu(a))
	fmt.Println(fu(b))

}

func fu(s string) int {
	if s == "" {
		return -1
	}
	if len(s) == 1 {
		return 0
	}
	itemBank := make(map[rune]int)

	for _, v := range s {
		itemBank[v]++
	}

	for k, v := range s {
		if itemBank[v] == 1 {
			return k
		}
	}
	return -1
}
