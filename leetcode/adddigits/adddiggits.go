package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := 38
	fmt.Println(Ad(input))
}

// Ad adding digits
func Ad(a int) int {
	result := strconv.Itoa(a)
	for len(result) != 1 {
		total := 0
		for _, v := range result {
			i := int(v - '0')
			total += i
		}
		result = strconv.Itoa(total)
	}
	k, _ := strconv.Atoi(result)
	return k
}
