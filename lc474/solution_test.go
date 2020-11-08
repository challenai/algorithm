package solution

import (
	"testing"
)

func TestFindMaxForm(t *testing.T) {
	strs1 := []string{"10", "0001", "111001", "1", "0"}
	m1 := 5
	n1 := 3
	strs2 := []string{"10", "0", "1"}
	m2 := 1
	n2 := 1
	r1 := 4
	r2 := 2

	rsu1 := findMaxForm(strs1, m1, n1)
	rsu2 := findMaxForm(strs2, m2, n2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
