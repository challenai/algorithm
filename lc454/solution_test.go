package solution

import (
	"testing"
)

func TestFourSumCount(t *testing.T) {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}

	r := 2
	rsu := fourSumCount(A, B, C, D)

	if rsu != r {
		t.Fail()
	}
}
