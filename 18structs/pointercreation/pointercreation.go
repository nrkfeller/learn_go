package main

import "fmt"

func main() {
	vd := &MetroStation{"vendome", "orange"}
	fmt.Println(vd)
	fmt.Printf("%T \n", vd)
	fmt.Println(vd.name)
	fmt.Println(vd.line)

}

// MetroStation struct
type MetroStation struct {
	name string
	line string
}
