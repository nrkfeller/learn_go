package main

import (
	"fmt"

	"github.com/nicfeller/learn_go/golangbasics/26testing/astar"
)

func main() {
	testinput := "a*a*a*a"
	fmt.Println(astar.Check(testinput))
}
