package solution

import (
	"testing"
)

func TestCanFinish(t *testing.T) {
	numCourses1 := 4
	prerequisites1 := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	r1 := true
	numCourses2 := 1
	prerequisites2 := [][]int{}
	r2 := true
	numCourses3 := 2
	prerequisites3 := [][]int{{1, 0}, {0, 1}}
	r3 := false

	rsu1 := canFinish(numCourses1, prerequisites1)
	rsu2 := canFinish(numCourses2, prerequisites2)
	rsu3 := canFinish(numCourses3, prerequisites3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
