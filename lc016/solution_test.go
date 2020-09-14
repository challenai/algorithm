package solution

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	target := 1
	r := 2
	resu := threeSumClosest(nums, target)
	println(resu)

	if resu != r {
		t.Fail()
	}
}
