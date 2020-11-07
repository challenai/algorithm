package solution

import (
	"testing"
)

func TestTrapRainWater(t *testing.T) {
	heightMap := [][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}
	r := 4
	rsu := trapRainWater(heightMap)

	if rsu != r {
		t.Fail()
	}
}
