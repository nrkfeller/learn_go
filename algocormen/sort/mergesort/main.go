package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	input := randomSlice(21)
	fmt.Println(input)
	input = mergesort(input)
	fmt.Println(input)
}

func mergesort(A []int) []int {
	if len(A) == 1 {
		return A
	}
	start := mergesort(A[:len(A)/2])
	end := mergesort(A[len(A)/2:])

	A = combine(start, end)

	return A
}

func combine(start, end []int) []int {
	var A []int
	s := 0
	e := 0
	for {
		if start[s] <= end[e] {
			A = append(A, start[s])
			s++
			if s == len(start) {
				A = append(A, end[e:]...)
				break
			}
		} else {
			A = append(A, end[e])
			e++
			if e == len(end) {
				A = append(A, start[s:]...)
				break
			}
		}
	}
	return A
}

func randomSlice(l int) []int {
	s := make([]int, l)
	for i := range s {
		s[i] = rand.Int() % 1000
	}
	return s
}
