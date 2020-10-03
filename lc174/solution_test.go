package solution

import (
	"testing"
)

func TestCalculateMinimumHP(t *testing.T) {
	dugeon := [][]int{
		{-2, -3, 3},
		{-5, 10, 1},
		{10, 30, -5},
	}
	r := 7
	resu := calculateMinimumHP(dugeon)
	if resu != r {
		t.Fail()
	}
}
