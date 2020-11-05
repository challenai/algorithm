package solution

import (
	"testing"
)

func TestFrequencySort(t *testing.T) {
	s1 := "tree"
	r1 := "eetr" // or eert ...
	s2 := "cccaaca"
	r2 := "ccccaaa"
	s3 := "Aabb"
	r3 := "bbAa"
	s4 := "ababdbsdbsd"
	r4 := "bbbbdddaass"

	rsu1 := frequencySort(s1)
	rsu2 := frequencySort(s2)
	rsu3 := frequencySort(s3)
	rsu4 := frequencySort(s4)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 || rsu4 != r4 {
		t.Fail()
	}
}
