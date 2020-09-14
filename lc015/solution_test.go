package solution

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-20, -50, -1, 0, 1, 2, -1, -4, -21, -16, -32, 200}
	r := [][]int{
		{-1, 0, 1},
		{-1, -1, 2},
	}

	resu := threeSum(nums)

	if len(resu) != len(r) {
		t.Fail()
	}
}
