package solution

import "strings"

func wordPattern(pattern, s string) bool {
	hashB2S := map[byte]string{}
	hashS2B := map[string]byte{}
	sList := strings.Split(s, " ")
	if len(sList) != len(pattern) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		if temp, ok := hashB2S[pattern[i]]; ok {
			if temp != sList[i] {
				return false
			}
		} else {
			hashB2S[pattern[i]] = sList[i]
			hashS2B[sList[i]] = pattern[i]
		}
	}
	for i := 0; i < len(sList); i++ {
		if temp, ok := hashS2B[sList[i]]; ok {
			if temp != pattern[i] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
