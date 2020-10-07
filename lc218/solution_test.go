package solution

import (
	"testing"
)

func TestGetSkyline(t *testing.T) {
	buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	r := [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}}

	resu := getSkyline(buildings)

	if len(resu) != len(r) {
		t.Fail()
	}
	if resu[2][2] != r[2][2] {
		t.Fail()
	}
}
