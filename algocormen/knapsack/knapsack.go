package knapsack

// func main() {
// 	values :=
// 	weights := []int{1, 3, 4, 5}
// 	fmt.Println(zoknapsack(7, weights, values, len(values)))
// }

// Knapsack is the zero one knapsack solved recursively
func Knapsack(W int, weights, values []int, n int) int {
	if n == 0 || W == 0 {
		return 0
	}

	if weights[n-1] > W {
		return Knapsack(W, weights, values, n-1)
	}
	return max(values[n-1]+Knapsack(W-weights[n-1], weights, values, n-1), Knapsack(W, weights, values, n-1))

}

func max(nums ...int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
