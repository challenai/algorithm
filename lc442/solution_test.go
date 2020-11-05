package solution

import (
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	// find all Duplicates in an array,
	// could you find all duplicates that appears twice?
	// yes, but I cant solve in O(n)time and O(1) space at the same time,
	// because I'm not a embed engineer, I seldom need to use bitwise operation.
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	r := []int{2, 3}

	rsu := findDuplicates(nums)

	if len(rsu) != len(r) {
		t.Fail()
	}
}
