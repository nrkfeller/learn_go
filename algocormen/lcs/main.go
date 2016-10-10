package main

import "fmt"

func main() {
	X := "AGGTAB"
	Y := "GXTXAYB"

	C := make([][]int, len(X))
	for i := range C {
		C[i] = make([]int, len(Y))
	}

	fmt.Println(rlcs(X, Y, len(X), len(Y), C))
}

func rlcs(a, b string, la, lb int, C [][]int) int {
	if la == 0 || lb == 0 {
		return 0
	}

	if a[la-1] == b[lb-1] {
		fmt.Println(string(a[la-1]), string(b[lb-1]))
		return 1 + rlcs(a, b, la-1, lb-1, C)
	}
	return max(rlcs(a, b, la-1, lb, C), rlcs(a, b, la, lb-1, C))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
