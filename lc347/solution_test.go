package solution

import (
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums1 := []int{6, 4, 1, 1, 1, 2, 2, 3, 6, 6, 6, 6, 1, 2, 2, 2, 4, 5, 6, 8, 9, 7, 6, 4}
	k1 := 3
	r1 := []int{1, 2, 6}
	nums2 := []int{1}
	k2 := 1
	r2 := []int{1}

	rsu1 := topKFrequent(nums1, k1)
	rsu2 := topKFrequent(nums2, k2)
	// for i := 0; i < len(rsu1); i++ {
	// 	print(rsu1[i], " ")
	// }
	// println()
	// for i := 0; i < len(rsu2); i++ {
	// 	print(rsu2[i])
	// }
	// println()

	if len(rsu1) != len(r1) || len(rsu2) != len(r2) {
		t.Fail()
	}
}
