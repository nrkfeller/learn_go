package main

import "fmt"

func main() {
	input := [][]int{{1, 1, 1, 1}, {0, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}

	fmt.Println(smz(input))
}

func smz(a [][]int) [][]int {
	for y := 0; y < len(a); y++ {
		for x := 0; x < len(a); x++ {
			if a[y][x] == 0 {
				for horiz := 0; horiz < len(a); horiz++ {
					a[y][horiz] = 0
				}

				for vert := 0; vert < len(a); vert++ {
					a[vert][x] = 0
				}
				return a
			}
		}
	}
	return a
}
