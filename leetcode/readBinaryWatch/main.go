package main

import "fmt"

func main() {
	input := 1
	fmt.Println(readBinaryWatch(input))
}

func readBinaryWatch(n int) string {
	possibleConfigs := make([]int, 10)
	
