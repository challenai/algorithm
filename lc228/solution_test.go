package solution

import (
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	nums1 := []int{0, 1, 2, 4, 5, 7}
	r1 := []string{"0->2", "4->5", "7"}
	nums2 := []int{0, 2, 3, 4, 6, 8, 9}
	r2 := []string{"0", "2->4", "6", "8->9"}
	nums3 := []int{-1}
	r3 := []string{"-1"}

	rsu1 := summaryRanges(nums1)
	rsu2 := summaryRanges(nums2)
	rsu3 := summaryRanges(nums3)

	if len(rsu1) != len(r1) || len(rsu2) != len(r2) || len(rsu3) != len(r3) {
		t.Fail()
	}
}
