package solution

import (
	"testing"
)

func TestDetectCapitalUse(t *testing.T) {

	word1 := "USA"
	r1 := true
	word2 := "leetcode"
	r2 := true
	word3 := "Google"
	r3 := true
	word4 := "FlaG"
	r4 := false

	rsu1 := detectCapitalUse(word1)
	rsu2 := detectCapitalUse(word2)
	rsu3 := detectCapitalUse(word3)
	rsu4 := detectCapitalUse(word4)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 || rsu4 != r4 {
		t.Fail()
	}
}
