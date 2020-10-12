package solution

import (
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 5, 0, 6, 0, 3, 0, 12}
	r := []int{1, 5, 6, 3, 12, 0, 0, 0, 0}

	moveZeroes(nums)

	for i := 0; i < len(r); i++ {
		if nums[i] != r[i] {
			t.Fail()
		}
	}

}
