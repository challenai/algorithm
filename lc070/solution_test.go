package solution

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	n1 := 2
	r1 := 2
	n2 := 3
	r2 := 3
	n3 := 7
	r3 := 21
	resu1 := climbStairs(n1)
	resu2 := climbStairs(n2)
	resu3 := climbStairs(n3)

	if resu1 != r1 || resu2 != r2 || resu3 != r3 {
		t.Fail()
	}
}
