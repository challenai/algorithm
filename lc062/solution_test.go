package solution

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	m1 := 3
	n1 := 2
	r1 := 3
	m2 := 7
	n2 := 3
	r2 := 28

	resu1 := uniquePaths(m1, n1)
	resu2 := uniquePaths(m2, n2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
