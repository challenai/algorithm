package solution

import (
	"testing"
)

func TestWordBreak(t *testing.T) {
	s1 := "leetcode"
	wordDict1 := []string{"leet", "code"}
	r1 := true
	s2 := "applepenapple"
	wordDict2 := []string{"apple", "pen"}
	r2 := true
	s3 := "catsandog"
	wordDict3 := []string{"cats", "dog", "sand", "and", "cat"}
	r3 := false

	rsu1 := wordBreak(s1, wordDict1)
	rsu2 := wordBreak(s2, wordDict2)
	rsu3 := wordBreak(s3, wordDict3)

	if rsu1 != r1 || rsu2 != r2 || rsu3 != r3 {
		t.Fail()
	}
}
