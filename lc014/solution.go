package solution

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var resu string
	resu = strs[0]
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < len(resu) {
			resu = resu[:len(strs[i])]
		}
		for j := 0; j < len(resu); j++ {
			if resu[j] != strs[i][j] {
				resu = resu[:j]
			}
		}
		if len(resu) == 0 {
			break
		}
	}

	return resu
}
