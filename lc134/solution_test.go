package solution

import (
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	gas1 := []int{1, 2, 3, 4, 5}
	cost1 := []int{3, 4, 5, 1, 2}
	r1 := 3
	gas2 := []int{2, 3, 4}
	cost2 := []int{3, 4, 3}
	r2 := -1

	rsu1 := canCompleteCircuit(gas1, cost1)
	rsu2 := canCompleteCircuit(gas2, cost2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
