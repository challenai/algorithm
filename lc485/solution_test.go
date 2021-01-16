package solution

import (
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	nums := []int{1, 1, 0, 1, 1, 1}
	r := 3

	rsu := FindMaxConsecutiveOnes(nums)
	if r != rsu {
		t.Fail()
	}
}
