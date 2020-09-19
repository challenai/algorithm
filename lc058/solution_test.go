package solution

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	s := "hello world"
	resu := lengthOfLastWord(s)
	r := 5

	if resu != r {
		t.Fail()
	}
}
