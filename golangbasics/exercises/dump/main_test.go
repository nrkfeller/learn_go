package basic

import "fmt"

func main() {
	a := []int{1, 2, 4, 55, 6}
	sum := 0
	sum += a[:2]
	fmt.Println()
}
