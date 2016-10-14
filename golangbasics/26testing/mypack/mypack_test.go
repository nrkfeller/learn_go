package mypack_test

import (
	"testing"

	"github.com/nicfeller/learn_go/golangbasics/26testing/mypack"
)

func TestCountby2(t *testing.T) {
	if mypack.Countby2() != 2 {
		t.Error("expected 2")
	}
	if mypack.Countby2() != 4 {
		t.Error("expected 4")
	}
	if mypack.Countby2() != 6 {
		t.Error("expected 6")
	}
	if mypack.Countby2() != 8 {
		t.Error("expected 8")
	}
}
