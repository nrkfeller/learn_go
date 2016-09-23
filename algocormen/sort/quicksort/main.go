package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	input := randomSlice(20)

	fmt.Println(input)
	quicksort(input)
	fmt.Println(input)
}

func quicksort(A []int) []int {
	if len(A) < 2 {
		return A
	}

	left, right := 0, len(A)-1
	pivot := rand.Int() % len(A)

	A[pivot], A[right] = A[right], A[pivot]

	for i := range A {
		if A[i] > A[right] {
			A[i], A[left] = A[left], A[i]
			left++
		}
	}

	A[left], A[right] = A[right], A[left]

	quicksort(A[:left])
	quicksort(A[left+1:])

	return A
}

func randomSlice(l int) []int {
	bytes := make([]int, l)
	for i := 0; i < l; i++ {
		bytes[i] = randInt(0, int(1000))
	}
	return bytes
}
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
