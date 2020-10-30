package solution

import (
	"strconv"
	"testing"
)

func TestPacificAtlantic(t *testing.T) {
	matrix := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	r := map[string]bool{
		"0_4": false,
		"1_3": false,
		"1_4": false,
		"2_2": false,
		"3_0": false,
		"3_1": false,
		"4_0": false,
	}

	resu := pacificAtlantic(matrix)
	if len(r) != len(resu) {
		t.Fail()
	}
	for i := 0; i < len(resu); i++ {
		if _, ok := r[strconv.Itoa(resu[i][0])+"_"+strconv.Itoa(resu[i][1])]; !ok {
			t.Fail()
		}
	}
}
