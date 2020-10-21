package solution

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	r1 := []byte{'o', 'l', 'l', 'e', 'h'}
	s2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	r2 := []byte{'h', 'a', 'n', 'n', 'a', 'H'}

	rsu1 := reverseString(s1)
	rsu2 := reverseString(s2)

	for i := 0; i < len(r1); i++ {
		if r1[i] != rsu1[i] {
			t.Fail()
		}
	}
	for i := 0; i < len(r2); i++ {
		if r2[i] != rsu2[i] {
			t.Fail()
		}
	}
}
