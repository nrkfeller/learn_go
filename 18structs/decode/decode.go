package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Person is a dude
type Person struct {
	Fname string
	Lname string
	Age   int
}

func main() {
	var p1 Person

	// `` turns string into a reader
	rdr := strings.NewReader(`{"nick", "feller", 26}`)

	// NewDecoder takes in a reader -> gives back pointer to decoder
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1)
}
