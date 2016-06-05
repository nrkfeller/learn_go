package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// attempt to open a file that does not exist
	_, err := os.Open("notrealfile.txt")
	if err != nil {
		// print error message
		fmt.Println("error occured", err)

		// log package writes standard error and logs time/date
		log.Println("error occured", err)
	}
}
