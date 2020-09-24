package solution

func isInterleave(s1 string, s2 string, s3 string) bool {
	dp := [][]bool{}
	for i := 0; i < len(s1)+1; i++ {
		dp = append(dp, make([]bool, len(s2)+1))
	}

	dp[0][0] = true
	for i := 1; i < len(s1)+1; i++ {
		if !dp[i-1][0] {
			break
		}
		dp[i][0] = s1[i-1] == s3[i-1]
	}
	for i := 1; i < len(s2)+1; i++ {
		if !dp[0][i-1] {
			break
		}
		dp[0][i] = s2[i-1] == s3[i-1]
	}

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			// bug appears here
			// use s3[i+j-1] instead of i+j-2
			// I solved by drawing the graph down to the paper
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	// for i := 0; i < len(s1)+1; i++ {
	// 	for j := 0; j < len(s2)+1; j++ {
	// 		print(dp[i][j], " ")
	// 	}
	// 	println()
	// }

	return dp[len(s1)][len(s2)]
}
