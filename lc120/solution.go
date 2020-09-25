package solution

func minimumTotal(triangle [][]int) int {
	// just make a n length array to store the state
	// since the follow-up says Bonus point if you are able to do this using only O(n) extra space
	// I decide to try to modify the origin triangle and use O(1) sc :)
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
			}
		}
	}

	resu := triangle[len(triangle)-1][0]
	for i := 1; i < len(triangle[len(triangle)-1]); i++ {
		if triangle[len(triangle)-1][i] < resu {
			resu = triangle[len(triangle)-1][i]
		}
	}

	return resu
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
