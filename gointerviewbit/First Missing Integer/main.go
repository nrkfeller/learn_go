package main

import "fmt"

func main() {
	input := []int{3, 4, -1, 2, 1}
	fmt.Println(fmi(input))
}

func fmi(a []int) int {
	l := len(a)
	j := 0
	for k, v := range a {
		if v <= 0 {
			temp := a[k]
			a[k] = a[j]
			a[j] = temp
			j++
		}
		fmt.Println(a)
	}
	for k, v := range a {
		if v < l && v > 0 {
			temp := a[v-1]
			a[v-1] = v
			a[k] = temp
		}
		fmt.Println(a)
	}
	for k, v := range a {
		if v != k+1 {
			return k + 1
		}
	}
	return l
}
