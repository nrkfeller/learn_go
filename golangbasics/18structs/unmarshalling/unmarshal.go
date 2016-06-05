package main

import (
	"encoding/json"
	"fmt"
)

// MetroStation yo
type MetroStation struct {
	Line    string
	Zipcode int
	Name    string
}

func main() {
	var m1 MetroStation
	fmt.Println(m1)

	bs := []byte(`{"Name":"de la savane", "Zipcode":1234,"Line":"orange"}`)

	// takes a byte slice and places it into an unset type
	json.Unmarshal(bs, &m1)

	fmt.Println(m1)
}
