package main

import "fmt"

func main() {

	x := 5
	// pass in address
	zero(&x)
	fmt.Println(x)

}

// *int is an address
func zero(p *int) {
	*p = 0
}
