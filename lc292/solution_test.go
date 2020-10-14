package solution

import (
	"testing"
)

func TestCanWinNim(t *testing.T) {
	n1 := 4
	r1 := false
	n2 := 1
	r2 := true
	n3 := 2
	r3 := true

	rsu1 := canWinNim(n1)
	rsu2 := canWinNim(n2)
	rsu3 := canWinNim(n3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
