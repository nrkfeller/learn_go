package main

import "fmt"

func main() {
	a := 1
	b := 1.23
	c := "hello"
	d := false

	var x int
	var y float64
	var z bool
	var w string

	var i = "hello" // actually golang doesnt want var i string = "hello"
	var j int8 = 4
	var k float32 = 3.455

	var p, o int8 = 12, 43

	fmt.Println(a, b, c, d)
	fmt.Printf("Short handed assignment : %v %v %v %v \n'", a, b, c, d)
	fmt.Printf("Their respective types : %T %T %T %T \n", a, b, c, d)
	fmt.Println("--------------------------")
	fmt.Printf("Initial values : %v %v %v %v \n'", w, x, y, z)
	fmt.Printf("Their respective types : %T %T %T %T \n", w, x, y, z)
	fmt.Println("--------------------------")
	fmt.Printf("Specificly assigned values : %v %v %v \n'", i, j, k)
	fmt.Printf("Their respective types : %T %T %T \n", i, j, k)
	fmt.Println("--------------------------")
	fmt.Printf("Multiple assigned values : %v %v \n'", p, o)
	fmt.Printf("Their respective types : %T %T\n", p, o)
}
