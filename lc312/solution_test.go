package solution

import (
	"testing"
)

func TestMaxCoins(t *testing.T) {
	nums := []int{3, 1, 5, 8}
	r := 167
	resu := maxCoins(nums)

	if resu != r {
		t.Fail()
	}
}
