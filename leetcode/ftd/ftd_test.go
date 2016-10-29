package ftd_test

import "testing"

func Testftd(t *testing.T) {
	s := "abcd"
	a := "abcde"

	wants := "e"
	gets := ftd.Ftd(s, a)

	if wants != gets {
		t.Error("should get", gets)
	}
}
