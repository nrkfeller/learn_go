package main

import "fmt"

func main() {
	input := 4
	fmt.Println(som2(input))
}

func som2(a int) [][]int {
	answer := make([][]int, a)
	for row := range answer {
		answer[row] = make([]int, a)
	}

	status := 0 // right
	current := 1
	rowindex := 0
	columnindex := 0
	end := a
	for current < a*a {
		switch {
		case status == 0:
			for i := columnindex; i < end; i++ {
				answer[columnindex][i] = current
				current++
			}
			columnindex = end - 1
			rowindex++
			status++

		case status == 1:
			for i := rowindex; i < end; i++ {
				answer[i][columnindex] = current
				current++
			}
			columnindex

		}
	}
	return answer
}
