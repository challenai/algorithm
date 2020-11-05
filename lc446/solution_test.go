package solution

import (
	"testing"
)

func TestNumberOfArithmeticSlices(t *testing.T) {
	nums := []int{2, 4, 6, 8, 10}
	r := 7

	rsu := numberOfArithmeticSlices(nums)

	if rsu != r {
		t.Fail()
	}
}
