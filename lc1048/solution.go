package solution

import "sort"

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) <= len(words[j])
	})
	if len(words) == 0 {
		return 0
	}
	dp := make([]int, len(words))
	ptr := 0
	for ptr < len(words) {
		if len(words[ptr]) != len(words[0]) {
			break
		}
		dp[ptr] = 1
		ptr++
	}
	var temp, prev int
	temp = 0
	for ptr < len(words) {
		if len(words[ptr]) > len(words[ptr-1]) {
			prev = temp
			temp = ptr
		}
		for i := prev; i < temp; i++ {
			if diff1Char(words[i], words[ptr]) {
				if dp[ptr] < dp[prev]+1 {
					dp[ptr] = dp[prev] + 1
				}
			}
		}
		ptr++
	}
	rsu := dp[len(dp)-1]
	for i := len(dp) - 1; i >= 0; i-- {
		if rsu < dp[i] {
			rsu = dp[i]
		}
	}
	return rsu
}

func diff1Char(s1, s2 string) bool {
	insertUsed := false
	if len(s1)+1 != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if !insertUsed {
			if s1[i] != s2[i] {
				insertUsed = true
			}
		} else if i < len(s1)-1 {
			if s1[i+1] != s2[i] {
				return false
			}
		}
	}
	return true
}
