package solution

import (
	"testing"
)

func TestMaxSumSubmatrix(t *testing.T) {
	matrix := [][]int{{1, 0, 1}, {0, -2, 3}}
	k := 2
	r := 2
	rsu := maxSumSubmatrix(matrix, k)

	if rsu != r {
		t.Fail()
	}
}
