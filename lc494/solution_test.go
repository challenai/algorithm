package solution

import (
	"testing"
)

func TestFindTargetSumWays(t *testing.T) {

	nums := []int{1, 1, 1, 1, 1}
	target := 3
	r := 5
	rsu := FindTargetSumWays(nums, target)

	if rsu != r {
		t.Fail()
	}
}
