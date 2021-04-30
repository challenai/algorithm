package solution

func findLUSlength(strs []string) int {
	if len(strs) == 0 {
		return -1
	}
	if len(strs) == 1 {
		return len(strs[0])
	}
	strsMaxLen := []int{0}
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) > len(strs[strsMaxLen[0]]) {
			strsMaxLen = []int{i}
			continue
		}
		if len(strs[i]) == len(strs[strsMaxLen[0]]) {
			strsMaxLen = append(strsMaxLen, i)
		}
	}
	hash := map[string]int{}
	for i := 0; i < len(strsMaxLen); i++ {
		if _, ok := hash[strs[strsMaxLen[i]]]; ok {
			hash[strs[strsMaxLen[i]]] = 2
		} else {
			hash[strs[strsMaxLen[i]]] = 1
		}
	}
	for k, v := range hash {
		if v == 1 {
			return len(k)
		}
	}
	return -1
}
