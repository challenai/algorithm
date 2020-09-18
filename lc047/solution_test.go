package solution

import (
	"testing"
)

func TestPermuteUnique(t *testing.T) {

	nums1 := []int{1, 1, 2}
	r1 := 3

	resu1 := permuteUnique(nums1)

	if len(resu1) != r1 {
		t.Fail()
	}
}
