package solution

import (
	"testing"
)

func TestMaxSumAfterPartitioning(t *testing.T) {
	arr1 := []int{1, 15, 7, 9, 2, 5, 10}
	k1 := 3
	r1 := 84
	arr2 := []int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}
	k2 := 4
	r2 := 83

	rsu1 := maxSumAfterPartitioning(arr1, k1)
	rsu2 := maxSumAfterPartitioning(arr2, k2)
	println(rsu1)
	println(rsu2)
	t.Fail()

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
