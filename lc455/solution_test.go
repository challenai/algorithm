package solution

import (
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	g1 := []int{1, 2, 3}
	s1 := []int{1, 1}
	r1 := 1
	g2 := []int{1, 2}
	s2 := []int{1, 2, 3}
	r2 := 2

	rsu1 := findContentChildren(g1, s1)
	rsu2 := findContentChildren(g2, s2)
	// rsu3 := findContentChildren(g3, s3)

	if rsu1 != r1 || rsu2 != r2 {
		t.Fail()
	}
}
