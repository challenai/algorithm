package solution

func fullJustify(words []string, maxWidth int) []string {
	resu := []string{}
	current := []string{words[0]}
	currentLen := len(words[0])
	for i := 0; i < len(words); i++ {
		if (currentLen + 1 + len(words[i])) <= maxWidth {
			current = append(current, words[i])
			currentLen += len(words[i]) + 1
		} else {
			// for j := 0; j < len(current); j++ {
			// 	print(current[j], " ")
			// }
			// println()
			// clear current
			resu = append(resu, "")
			currentLen = maxWidth - currentLen
			if len(current) == 1 {
				resu[len(resu)-1] = current[i] + trimSpace(currentLen)
			} else {
				trimCnt := currentLen / (len(current) - 1)
				currentLen %= len(current) - 1
				for j := 0; j < len(current); j++ {
					if currentLen > 0 {
						resu[len(resu)-1] += current[j] + trimSpace(trimCnt+1)
						currentLen--
					} else if i == len(resu)-1 {
						resu[len(resu)-1] += current[j]
					} else {
						resu[len(resu)-1] += current[j] + trimSpace(trimCnt)
					}
				}
			}
			current = []string{words[i]}
			currentLen = len(words[i])
		}
	}
	for i := 0; i < len(resu); i++ {
		println(resu[i])
	}
	return resu
}

func trimSpace(n int) string {
	resu := ""
	for i := 0; i < n; i++ {
		resu += " "
	}
	return resu
}
