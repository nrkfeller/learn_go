package main

import "fmt"

func main() {
	input := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

	direction := "right"
	stride := len(input)

	total := len(input) * len(input[0])

	current := []int{0, 0}

	for total > 0 {
		switch {
		case direction == "right":
			for i := 0; i < stride; i++ {
				//fmt.Println(direction, total)
				fmt.Println(input[current[0]][current[1]])
				current[1]++
				total--
			}
			current[1]--
			stride--
			direction = "down"

		case direction == "down":
			current[0]++
			for i := 0; i < stride; i++ {
				//fmt.Println(direction, total)
				fmt.Println(input[current[0]][current[1]])
				current[0]++
				total--
			}
			current[0]--
			direction = "left"

		case direction == "left":
			current[1]--
			for i := 0; i < stride; i++ {
				//fmt.Println(direction, total)
				fmt.Println(input[current[0]][current[1]])
				current[1]--
				total--
			}
			current[1]++
			stride--
			direction = "up"

		case direction == "up":
			for i := 0; i < stride; i++ {
				//fmt.Println(direction, total)
				fmt.Println(input[current[0]][current[1]])
				current[0]--
				total--
			}
			direction = "right"

		}
	}
}
