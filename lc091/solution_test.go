package solution

import (
	"testing"
)

func TestNumDecodings(t *testing.T) {
	s1 := "882262"
	r1 := 3
	s2 := "226"
	r2 := 3

	resu1 := numDecodings(s1)
	resu2 := numDecodings(s2)

	if resu1 != r1 || resu2 != r2 {
		t.Fail()
	}
}
