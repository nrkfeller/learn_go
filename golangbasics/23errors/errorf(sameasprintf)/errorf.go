package main

import (
	"fmt"
	"log"
)

func main() {
	ans, err := roots(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(ans)
}

func roots(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("%v cannot be rooted", n)
	}
	return 42, nil
}
