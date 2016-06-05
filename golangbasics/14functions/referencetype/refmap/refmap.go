package main

import "fmt"

func main() {
	mmap := make(map[string]int)
	mmap["nick"] = 99
	fmt.Println(mmap)

	// same as string this actually changes!
	change(mmap)
	fmt.Println(mmap)
}

func change(m map[string]int) {
	m["nick"] = 123543
}
