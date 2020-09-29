package solution

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	// initialize the state
	dp[0][0] = true
	for i := 1; i < len(p); i++ {
		// we assume that ** can match ""
		if p[i] == '*' {
			dp[0][i+1] = true
		} else {
			break
		}
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if s[i] == p[j] || p[j] == '?' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				if j == 0 {
					dp[i+1][j+1] = false
				} else if p[j-1] == '*' {
					dp[i+1][j+1] = dp[i+1][j]
				} else if p[j-1] == '.' {
					// todo
					if j == 1 {
						dp[i+1][j+1] = true
						continue
					}
					dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j-1]
				} else {
					if p[j-1] == s[i] {
						dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j]
					}
				}
			}
		}
	}

	return dp[len(s)][len(p)]
}
