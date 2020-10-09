package solution

import (
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	r := []string{"AAAAACCCCC", "CCCCCAAAAA"}

	resu := findRepeatedDnaSequences(s)

	if len(resu) != len(r) {
		t.Fail()
	}
}
