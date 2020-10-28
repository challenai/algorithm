package solution

import (
	"testing"
)

func TestReconstructQueue(t *testing.T) {
	nums := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	r := [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}

	rsu := reconstructQueue(nums)

	for i := 0; i < len(r); i++ {
		if rsu[i][0] != r[i][0] || rsu[i][1] != r[i][1] {
			t.Fail()
		}
	}
}
