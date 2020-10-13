package solution

import (
	"testing"
)

func TestSumRegion(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	numMatrix := Constructor(matrix)
	if numMatrix.SumRegion(2, 1, 4, 3) != 8 {
		t.Fail()
	}
	if numMatrix.SumRegion(1, 1, 2, 2) != 11 {
		t.Fail()
	}
	if numMatrix.SumRegion(1, 2, 2, 4) != 12 {
		t.Fail()
	}
}
