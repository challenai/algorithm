package solution

import (
	"testing"
)

func TestLargestNumber(t *testing.T) {
	nums1 := []int{10, 2}
	r1 := "210"
	nums2 := []int{3, 30, 34, 5, 9}
	r2 := "9534330"
	nums3 := []int{1}
	r3 := "1"
	nums4 := []int{10}
	r4 := "10"

	resu1 := largestNumber(nums1)
	resu2 := largestNumber(nums2)
	resu3 := largestNumber(nums3)
	resu4 := largestNumber(nums4)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 || resu4 != r4 {
		t.Fail()
	}
}
