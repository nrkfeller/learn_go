package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	const input = "hello the north \n will not \n forget\n"
	scanner := bufio.NewScanner(strings.NewReader(input))

	// different split possibilities
	//scanner.Split(bufio.ScanBytes) // 36
	//scanner.Split(bufio.ScanWords) // 6
	scanner.Split(bufio.ScanLines) // 3

	count := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading output", err)
	}
	fmt.Println(count)
}
