package solution

import (
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	nums := []int{2, 1, 5, 6, 2, 3}
	r := 10

	resu := largestRectangleArea(nums)

	if resu != r {
		t.Fail()
	}
}
