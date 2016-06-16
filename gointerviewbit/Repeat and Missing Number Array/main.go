package main

import "fmt"

func main() {
	input := []int{1, 10, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(findboth(input))
}

func findboth(a []int) (int, int) {
	n := len(a)
	supposed := (n * (n + 1)) / 2
	actualsq, actual := fullsum(a)
	diff := supposed - actual
	over := (supposed*(2*n+1)/3 - actualsq) / diff
	ret := (over - diff) / 2
	return ret, ret + diff
}

func fullsum(a []int) (int, int) {
	totalsq := 0
	total := 0
	for _, v := range a {
		totalsq += v * v
		total += v
	}
	return totalsq, total
}
