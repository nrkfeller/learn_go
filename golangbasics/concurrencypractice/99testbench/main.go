package main

import (
	"fmt"
	"time"
)

func main() {
	var arr []int
	for i := 0; i < 10000000; i++ {
		arr = append(arr, i)
	}

	start := time.Now()
	total := 0

	for i := range arr {
		total += i * i
	}
	fmt.Println(time.Since(start), total)
}
