package main

import (
	"errors"
	"fmt"
	"log"
)

// ErrNegRoot for negative root attempts
var ErrNegRoot = errors.New("negative number cannot be rooted")

func main() {
	ans, err := rootroot(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(ans)
}

func rootroot(n float64) (float64, error) {
	if n < 0 {
		// just create a custom error
		fmt.Printf("%T", ErrNegRoot)
		return 0, ErrNegRoot
	}
	return 42, nil
}
