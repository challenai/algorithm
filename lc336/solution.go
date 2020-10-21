package solution

func palindromePairs(words []string) [][]int {
	rsu := [][]int{}
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if checkPalindrome(words[j] + words[i]) {
				rsu = append(rsu, []int{j, i})
			}
			if checkPalindrome(words[i] + words[j]) {
				rsu = append(rsu, []int{i, j})
			}
		}
	}
	return rsu
}

func checkPalindrome(s string) bool {
	var left, right int
	left = 0
	right = len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
