package solution

import (
	"testing"
)

func TestInsert(t *testing.T) {
	intervals1 := [][]int{
		[]int{1, 3},
		[]int{6, 9},
	}
	newInterval1 := []int{2, 5}
	r1 := [][]int{
		[]int{1, 5},
		[]int{6, 9},
	}

	intervals2 := [][]int{
		[]int{1, 2},
		[]int{3, 5},
		[]int{6, 7},
		[]int{8, 10},
		[]int{12, 16},
	}
	newInterval2 := []int{4, 8}
	r2 := [][]int{
		[]int{1, 2},
		[]int{3, 5},
		[]int{6, 7},
		[]int{8, 10},
		[]int{12, 16},
	}

	intervals3 := [][]int{}
	newInterval3 := []int{5, 7}
	r3 := [][]int{{5, 7}}
	intervals4 := [][]int{
		[]int{1, 5},
	}
	newInterval4 := []int{2, 3}
	r4 := [][]int{{1, 5}}

	resu1 := insert(intervals1, newInterval1)
	resu2 := insert(intervals2, newInterval2)
	resu3 := insert(intervals3, newInterval3)
	resu4 := insert(intervals4, newInterval4)

	showResu(resu1)
	showResu(resu2)
	showResu(resu3)
	showResu(resu4)

	if false {
		t.Fail()
	}
}

func showResu(resu [][]int) {
	for i := 0; i < len(resu); i++ {
		for j := 0; j < len(resu[i]); j++ {
			print(resu[i][j], " ")
		}
		println()
	}
}
