package solution

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	r := 6
	resu := maxSubArray(nums)
	if resu != r {
		t.Fail()
	}
}
