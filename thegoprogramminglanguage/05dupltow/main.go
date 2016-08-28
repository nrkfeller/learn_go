package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	file := os.Args[1:]
	if len(file) != 1 {
		fmt.Println("Incorrect entry")
	} else {
		f, _ := os.Open(file[0])
		countLines(f, counts)
	}
}

func countLines(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}
	fmt.Println(count)
}
