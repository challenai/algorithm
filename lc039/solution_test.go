package solution

import (
	"testing"
)

func TestCombinationSum(t *testing.T) {
	candidates1 := []int{2, 3, 6, 7}
	target1 := 7
	r1Len := 2
	candidates2 := []int{2, 3, 5}
	target2 := 8
	r2Len := 3

	resu1 := combinationSum(candidates1, target1)
	resu2 := combinationSum(candidates2, target2)

	if len(resu1) != r1Len || len(resu2) != r2Len {
		t.Fail()
	}
}
