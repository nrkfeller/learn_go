package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, _ := strconv.ParseBool("true")
	b, _ := strconv.ParseFloat("123.432", 64)
	c, _ := strconv.ParseInt("432", 10, 32)
	d, _ := strconv.ParseUint("-992", 10, 32) // this parses to 0 because there can not be negative uint

	fmt.Println(a, b, c, d)

	w := strconv.FormatBool(false)
	x := strconv.FormatFloat(12324.543, 'E', -1, 64)
	y := strconv.FormatInt(666, 8)
	z := strconv.FormatUint(33, 16)
	fmt.Println(w, x, y, z)
}
