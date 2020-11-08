package solution

func findMaxForm(strs []string, m, n int) int {
	rsu := 0
	for i := 0; i < len(strs); i++ {
		counter := 0
		for j := 0; j < len(strs[i]); j++ {
			if strs[i][j] == '1' {
				counter++
			}
		}
		if counter <= n && len(strs[i])-counter <= m {
			rsu++
		}
	}
	return rsu
}
