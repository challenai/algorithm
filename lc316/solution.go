package solution

func removeDuplicateLetters(s string) string {
	hash := map[byte][]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := hash[s[i]]; ok {
			hash[s[i]] = append(hash[s[i]], i)
		} else {
			hash[s[i]] = []int{i}
		}
	}
}
