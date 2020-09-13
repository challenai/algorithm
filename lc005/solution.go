package solution

func longestPalindrome(s string) string {
	// idea: I can use dp to figure out if current string is a palindrome,
	// because we can always grow the array from a single character
	if s == "" {
		return ""
	}
	currentLongestStr := s[0:1]
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	// todo init dp
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}
	for i := 1; i < len(s); i++ {
		dp[i-1][i] = s[i-1] == s[i]
		if dp[i-1][i] {
			currentLongestStr = s[i-1 : i+1]
		}
	}

	// forward
	for left := 0; left < len(s)-3; left++ {
		for right := left + 2; right < len(s); right++ {
			dp[left][right] = dp[left+1][right-1] && s[left] == s[right]
			if dp[left][right] && right-left+1 > len(currentLongestStr) {
				currentLongestStr = s[left : right+1]
			}
		}
	}

	return currentLongestStr
}

func longestPalindrome2(s string) string {
	// catch up another idea that I can spread the palindrome from a single mid,
	// care that a palindrome may have two mid element
	return ""
}
