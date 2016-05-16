package main

import "fmt"

var metersToYards = 1.09361

func main() {

	var meters float64

	fmt.Print("Enter meters swam : ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, "is in yards", yards)

}
