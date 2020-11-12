package solution

// func longestDupSubstring(s string) string {
// 	// I cant figure out the dp O(n) solution after a 10 min exploration
// 	// my solution is O(n^2)
// 	dp := make([][]int, len(s))
// 	for i := 0; i < len(s); i++ {
// 		dp[i] = make([]int, len(s))
// 	}
// 	for i := 0; i < len(s); i++ {
// 		for j := i + 1; j < len(s); j++ {
// 			// if s[i] == s[j]
// 		}
// 	}

// }

func longestDupSubstring(s string) string {
	var m, n int
	var current, resu string
	resu = ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				m = i
				n = j
				for m >= 0 && s[m] == s[n] {
					current = s[m : i+1]
					m--
					n--
				}
				if len(current) > len(resu) {
					resu = current
				}
			}
		}
	}
	return resu
}
