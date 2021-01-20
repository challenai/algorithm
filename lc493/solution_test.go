package solution

import (
	"testing"
)

func TestReversePairs(t *testing.T) {

	nums1 := []int{1, 3, 2, 3, 1}
	r1 := 2
	nums2 := []int{2, 4, 3, 5, 1}
	r2 := 3
	rsu1 := ReversePairs(nums1)
	rsu2 := ReversePairs(nums2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
