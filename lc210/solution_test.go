package solution

import (
	"testing"
)

func TestFindOrder(t *testing.T) {
	numCourses1 := 4
	prerequisites1 := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	r1 := []int{0, 1, 2, 3}
	numCourses2 := 1
	prerequisites2 := [][]int{}
	r2 := []int{0}

	resu1 := findOrder(numCourses1, prerequisites1)
	resu2 := findOrder(numCourses2, prerequisites2)

	for i := 0; i < len(r1); i++ {
		if r1[i] != resu1[i] {
			t.Fail()
		}
	}

	if resu2[0] != r2[0] {
		t.Fail()
	}
}
