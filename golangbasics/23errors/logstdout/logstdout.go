package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("doesntexist.txt")
	if err != nil {
		fmt.Println("Writing to log file")
		log.Println("File not found", err)
	}
}

// func init runs before main, kind of to set up the program
func init() {
	nf, err := os.Create("potato.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("created the log file")
	}

	// this part tells the log to put the logs somewhere else, IE the potato.txt stored in MM as nf
	log.SetOutput(nf)
}
