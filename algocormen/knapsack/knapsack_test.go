package knapsack_test

import (
	"testing"

	"github.com/nicfeller/learn_go/algocormen/knapsack"
)

func TestKnapsack(t *testing.T) {
	inputValues := []int{1, 4, 5, 7}
	inputWeights := []int{1, 3, 4, 5}
	gets := knapsack.Knapsack(7, inputWeights, inputValues, len(inputValues))
	wants := 9
	if gets != wants {
		t.Error(inputValues, inputWeights, "is optimal at", wants)
	}
}
