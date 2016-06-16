package main

import (
	"fmt"
	"math"
)

func main() {
	input := [][]float64{{0, 0}, {3, 1}, {4, 2}, {1, 1}}

	fmt.Println(getfeweststeps(input))
}

func getfeweststeps(points [][]float64) float64 {
	x := points[0][0]
	y := points[0][1]
	var total, xdis, ydis float64
	for i := 1; i < len(points); i++ {
		xdis = math.Abs(x - points[i][0])
		ydis = math.Abs(y - points[i][1])
		x = points[i][0]
		y = points[i][1]
		total += largest(xdis, ydis)
	}

	return total
}

func largest(a float64, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
