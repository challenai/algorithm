package solution

import (
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	digits := "23"
	r := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	resu := letterCombinations(digits)

	if len(resu) != len(r) {
		t.Fail()
	}
}
