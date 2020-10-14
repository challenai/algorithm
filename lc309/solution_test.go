package solution

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	nums := []int{1, 2, 3, 0, 2}
	r := 3

	resu := maxProfit(nums)

	if resu != r {
		t.Fail()
	}
}
