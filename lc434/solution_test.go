package solution

import (
	"testing"
)

func TestCountSegments(t *testing.T) {
	s1 := "Hello, my name is John"
	r1 := 5
	s2 := "Hello"
	r2 := 1
	s3 := "love live! mu'sic forever"
	r3 := 4
	s4 := ""
	r4 := 0

	rsu1 := countSegments(s1)
	rsu2 := countSegments(s2)
	rsu3 := countSegments(s3)
	rsu4 := countSegments(s4)

	if r1 != rsu1 || r2 != rsu2 || r3 != rsu3 || r4 != rsu4 {
		t.Fail()
	}
}
