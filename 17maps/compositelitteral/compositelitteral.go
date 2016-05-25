package main

import "fmt"

func main() {
	m := map[string]string{"nick": "boss", "others": "slaves"}

	fmt.Println(m)

	var m2 = make(map[string]string)
	m2["english"] = "hello"
	m2["german"] = "guttentag"
	m2["italian"] = "buongiorno"

	delete(m2, "english")

	fmt.Println(m2)

	// initializes to nil
	//m3 := make(map[int]float64)
	m4 := map[string]string{
		"potato": "yes",
		"tomato": "no",
	}
	fmt.Println(m4 == nil)
}
