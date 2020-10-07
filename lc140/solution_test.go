package solution

import (
	"testing"
)

func TestWordBreak(t *testing.T) {
	s1 := "catsanddog"
	wordDict1 := []string{"cat", "cats", "and", "sand", "dog"}
	r1 := []string{"cats and dog", "cat sand dog"}
	s2 := "pineapplepenapple"
	wordDict2 := []string{"apple", "pen", "applepen", "pine", "pineapple"}
	r2 := []string{
		"pine apple pen apple",
		"pineapple pen apple",
		"pine applepen apple",
	}
	s3 := "catsandog"
	wordDict3 := []string{"cats", "dog", "sand", "and", "cat"}
	r3 := []string{}

	resu1 := wordBreak(s1, wordDict1)
	resu2 := wordBreak(s2, wordDict2)
	resu3 := wordBreak(s3, wordDict3)

	if len(resu1) != len(r1) || len(resu2) != len(r2) || len(resu3) != len(r3) {
		t.Fail()
	}
}
