package main

import "fmt"

func main() {
	v := []int{1, 4, 5, 7}
	w := []int{1, 3, 4, 5}
	fmt.Println(zoknapsack(v, w, 7))
}

func zoknapsack(v, w []int, W int) int {
	dp := make([]int, W+1)
	for i := 0; i < len(v); i++ {
		for j := W; j > w[i]-1; j-- {
			dp[j] = max(dp[j], dp[j-w[i]], v[i])
		}
	}
	return dp[W]
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
