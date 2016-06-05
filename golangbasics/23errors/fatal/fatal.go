package main

import (
	"log"
	"os"
)

// os exit, give and exit status, 0 means success, non zero means an error, the number associates with a specific error
func main() {
	_, err := os.Open("doesntexist.txt")
	if err != nil {
		// exit status 1!
		log.Fatal(err)
	}
}
