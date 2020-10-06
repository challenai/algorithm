package solution

import (
	"testing"
)

func TestFindSubstring(t *testing.T) {
	s1 := "barfoothefoobarman"
	words1 := []string{"foo", "bar"}
	s2 := "wordgoodgoodgoodbestword"
	words2 := []string{"word", "good", "best", "word"}
	s3 := "barfoofoobarthefoobarman"
	words3 := []string{"bar", "foo", "the"}
	r1 := []int{0, 9}
	r2 := []int{}
	r3 := []int{6, 9, 12}

	rsu1 := findSubstring(s1, words1)
	rsu2 := findSubstring(s2, words2)
	rsu3 := findSubstring(s3, words3)

	if len(rsu1) != len(r1) || len(rsu2) != len(r2) || len(rsu3) != len(r3) {
		t.Fail()
	}
}
