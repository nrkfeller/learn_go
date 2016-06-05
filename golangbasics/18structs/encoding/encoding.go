package main

import "os"
import "encoding/json"

// Person is a dude
type Person struct {
	Fname string
	Lname string
	Age   int
}

func main() {
	p1 := Person{"nick", "feller", 26}

	// writing it to standard out (just an open file that we can always write to)!
	// creates new pointer to encoder -- then encodes passed argument
	json.NewEncoder(os.Stdout).Encode(p1)
}
