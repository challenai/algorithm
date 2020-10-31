package solution

func countSegments(s string) int {
	resu := 0
	segementStart := false
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			segementStart = true
			continue
		}
		if segementStart {
			resu++
			segementStart = false
		}
	}
	if segementStart {
		return resu + 1
	}
	return resu
}
