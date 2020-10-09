package solution

func findRepeatedDnaSequences(s string) []string {
	exist := map[string]bool{}
	resuMap := map[string]bool{}
	resu := []string{}
	for i := 9; i < len(s); i++ {
		if _, ok := exist[s[i-9:i+1]]; !ok {
			exist[s[i-9:i+1]] = false
		} else {
			resuMap[s[i-9:i+1]] = false
		}
	}
	for key := range resuMap {
		resu = append(resu, key)
	}
	return resu
}
