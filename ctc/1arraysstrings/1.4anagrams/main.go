package main

import "fmt"

func main() {
	i1 := "potatoc"
	i2 := "oopttac"
	fmt.Println(anagram(i1, i2))
}

func anagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	asc := make([]int, 256)
	asc2 := make([]int, 256)
	for i := 0; i < len(a); i++ {
		asc[a[i]]++
		asc2[b[i]]++
	}
	for i := 0; i < len(asc); i++ {
		if asc[i] != asc2[i] {
			return false
		}
	}
	return true
}
