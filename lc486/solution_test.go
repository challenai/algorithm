package solution

import (
	"testing"
)

func TestPredictTheWinner(t *testing.T) {

	nums1 := []int{1, 5, 2}
	r1 := false
	nums2 := []int{1, 5, 233, 7}
	r2 := true

	rsu1 := PredictTheWinner(nums1)
	rsu2 := PredictTheWinner(nums2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
