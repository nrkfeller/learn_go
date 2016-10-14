package astar_test

import (
	"testing"

	"github.com/nicfeller/learn_go/golangbasics/26testing/astar"
)

func TestCheck(t *testing.T) {
	testinput := "a*a*a"
	get := astar.Check(testinput)
	want := true
	if get != want {
		t.Error(testinput, " should be valid ")
	}

	testinput = "a*a*a*"
	get = astar.Check(testinput)
	want = false
	if get != want {
		t.Error(testinput, " should be invalid ")
	}
}
