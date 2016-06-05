package main

import (
	"errors"
	"fmt"
	"log"
)

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
		fmt.Printf("%T", errors.New("negative number cannot be rooted"))
		return 0, errors.New("negative number cannot be rooted")
	}
	return 42, nil
}
