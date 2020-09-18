package solution

import (
	"testing"
)

func TestPermute(t *testing.T) {

	nums1 := []int{1, 2, 3}
	r1 := 6

	resu1 := permute(nums1)
	if len(resu1) != r1 {
		t.Fail()
	}
}
