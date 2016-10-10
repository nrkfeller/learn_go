package main

import "fmt"

func main() {
	input := "hello its me adele"
	fmt.Println(replacespace(input))
}

func replacespace(a string) string {
	count := 0

	for _, v := range a {
		if v == 32 {
			count++
		}
	}

	ba := make([]byte, len(a)+count*2)

	for k := range a {
		if a[k] == 32 {
			ba = append(ba, []byte("%20")...)
		} else {
			ba = append(ba, a[k])
		}
	}

	return string(ba)
}
