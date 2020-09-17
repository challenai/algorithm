package solution

import (
	"testing"
)

func TestCombinationSum(t *testing.T) {
	candidates1 := []int{10, 1, 2, 7, 6, 1, 5}
	target1 := 8
	r1Len := 4
	candidates2 := []int{2, 5, 2, 1, 2}
	target2 := 5
	r2Len := 2

	resu1 := combinationSumz(candidates1, target1)
	resu2 := combinationSumz(candidates2, target2)

	if len(resu1) != r1Len || len(resu2) != r2Len {
		t.Fail()
	}
}
