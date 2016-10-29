package lcs_test

import (
	"testing"

	"github.com/nicfeller/learn_go/algocormen/lcs"
)

func Testlcs(t *testing.T) {
	X := "AGGTAB"
	Y := "GXTXAYB"

	C := make([][]int, len(X))
	for i := range C {
		C[i] = make([]int, len(Y))
	}

	want := 4
	gets := lcs.Rlcs(X, Y, len(X), len(Y), C)

	if want != gets {
		t.Error("should get", gets)
	}
}
