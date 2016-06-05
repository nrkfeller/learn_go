package main

import "os"

func main() {
	_, err := os.Open("doesnotexists.txt")
	if err != nil {
		// gives the stack trace, information

		// stops normal execution, panic/recover is like a try-catch block
		panic(err)
	}
}
