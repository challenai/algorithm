package solution

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	target := 49
	resu := maxArea(nums)

	if resu != target {
		t.Fail()
	}
}
