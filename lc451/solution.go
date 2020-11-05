package solution

func frequencySort(s string) string {
	rsu := ""
	hash := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := hash[s[i]]; ok {
			hash[s[i]]++
		} else {
			hash[s[i]] = 1
		}
	}
	var maxKey byte
	for len(hash) > 0 {
		maxKey = '.'
		for k, v := range hash {
			if maxKey == '.' || hash[maxKey] < v {
				maxKey = k
			}
		}
		for i := 0; i < hash[maxKey]; i++ {
			rsu += string(maxKey)
		}
		delete(hash, maxKey)
	}
	// println(rsu)
	return rsu
}
