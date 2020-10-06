package solution

import "strings"

var str1 []string
var resu bool

func isScramble(s1, s2 string) bool {
	resu := false
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) <= 2 {
		return true
	}
	str1 = strings.Split(s1, "")
	return resu
}
