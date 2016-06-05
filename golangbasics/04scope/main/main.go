package main

import "fmt"

// fully qualified package name!
import "github.com/udemycourse/04scope/doublepackage"

var potato = 123

func main() {

	fmt.Println(doublepackage.Double(potato))
	fmt.Printf(doublepackage.Dubs)

}
