package solution

import (
	"testing"
)

func TestOriginalDigits(t *testing.T) {
	s1 := "owoztneoer"
	r1 := "012"
	s2 := "fviefuro"
	r2 := "45"
	s3 := "ffoourrutreehtwtwoo"
	r3 := "22344"

	rsu1 := originalDigits(s1)
	rsu2 := originalDigits(s2)
	rsu3 := originalDigits(s3)

	if r1 != rsu1 || r2 != rsu2 || r3 != rsu3 {
		t.Fail()
	}
}
