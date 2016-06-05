package main

import (
	"fmt"
	"math"
)

func main() {
	s := Square{10}
	c := Circle{10}
	info(s)
	info(c)
}

// Shape is just an interface
type Shape interface {
	area() float64
}

// Square stands on its own
type Square struct {
	side float64
}

// Circle stands on its own, doesnt need to inherit or anything
type Circle struct {
	radius float64
}

// area() float64 - this means it can be called using Shape interface
func (c Circle) area() float64 {
	return c.radius * c.radius * 3.1415
}

// square has a function that returns its area
// area() float64 - this means it can be called using Shape interface
func (s Square) area() float64 {
	return math.Pow(s.side, 2)
}

// shapes have an info function that prints its information
func info(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}
