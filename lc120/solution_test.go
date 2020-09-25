package solution

import (
	"testing"
)

func TestMinimumTotal(t *testing.T) {

	nums := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	r := 11
	resu := minimumTotal(nums)

	if resu != r {
		t.Fail()
	}
}
