package solution

import (
	"testing"
)

func TestMinMutation(t *testing.T) {
	start1 := "AACCGGTT"
	end1 := "AAACGGTA"
	bank1 := []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}
	r1 := 2
	start2 := "AAAAACCC"
	end2 := "AACCCCCC"
	bank2 := []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}
	r2 := 3

	rsu1 := minMutation(start1, end1, bank1)
	rsu2 := minMutation(start2, end2, bank2)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
