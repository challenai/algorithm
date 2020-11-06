package solution

import (
	"testing"
)

func TestCanIWin(t *testing.T) {
	maxChoosableInteger1 := 10
	desiredTotal1 := 11
	r1 := false
	maxChoosableInteger2 := 10
	desiredTotal2 := 0
	r2 := true

	rsu1 := canIWin(maxChoosableInteger1, desiredTotal1)
	rsu2 := canIWin(maxChoosableInteger2, desiredTotal2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
