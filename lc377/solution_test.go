package solution

import (
	"testing"
)

func TestCombinationSum4(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 7}
	target := 9
	r := 25

	rsu := combinationSum4(nums, target)
	println(rsu)

	if rsu != r {
		t.Fail()
	}
}
