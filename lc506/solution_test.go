package solution

import (
	"testing"
)

func TestFindRelativeRanks(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	r := []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}

	rsu := findRelativeRanks(nums)

	for i := 0; i < len(r); i++ {
		if rsu[i] != r[i] {
			t.Fail()
		}
	}
}
