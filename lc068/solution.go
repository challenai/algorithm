package solution

// func fullJustify(words []string, maxWidth int) []string {
// 	resu := []string{}
// 	current := []string{words[0]}
// 	currentLen := len(words[0])
// 	for i := 0; i < len(words); i++ {
// 		if (currentLen + 1 + len(words[i])) <= maxWidth {
// 			current = append(current, words[i])
// 			currentLen += len(words[i]) + 1
// 		} else {
// 			// for j := 0; j < len(current); j++ {
// 			// 	print(current[j], " ")
// 			// }
// 			// println()
// 			// clear current
// 			resu = append(resu, "")
// 			currentLen = maxWidth - currentLen
// 			if len(current) == 1 {
// 				resu[len(resu)-1] = current[i] + trimSpace(currentLen)
// 			} else {
// 				trimCnt := currentLen / (len(current) - 1)
// 				currentLen %= len(current) - 1
// 				for j := 0; j < len(current); j++ {
// 					if currentLen > 0 {
// 						resu[len(resu)-1] += current[j] + trimSpace(trimCnt+1)
// 						currentLen--
// 					} else if i == len(resu)-1 {
// 						resu[len(resu)-1] += current[j]
// 					} else {
// 						resu[len(resu)-1] += current[j] + trimSpace(trimCnt)
// 					}
// 				}
// 			}
// 			current = []string{words[i]}
// 			currentLen = len(words[i])
// 		}
// 	}
// 	for i := 0; i < len(resu); i++ {
// 		println(resu[i])
// 	}
// 	return resu
// }

func fullJustify(words []string, maxWidth int) []string {
	resu := []string{}
	var currentLine []string
	var currentLen, ptr, restSpace int
	ptr = 0
	for ptr < len(words) {
		if currentLen == 0 {
			currentLine = append(currentLine, words[ptr])
			currentLen = len(words[ptr])
			ptr++
		} else if currentLen+1+len(words[ptr]) <= maxWidth {
			currentLine = append(currentLine, words[ptr])
			currentLen += 1 + len(words[ptr])
			ptr++
		} else {
			if len(currentLine) == 1 {
				resu = append(resu, currentLine[0]+trimSpace(maxWidth-len(currentLine[0])))
			} else {
				// generate new line
				spaceNum := maxWidth - currentLen + len(currentLine) - 1
				trimSize := spaceNum / (len(currentLine) - 1)
				restSpace = spaceNum % (len(currentLine) - 1)
				for i := 0; i < len(currentLine); i++ {
					if i == 0 {
						resu = append(resu, currentLine[i])
						continue
					}
					if restSpace > 0 {
						resu[len(resu)-1] += trimSpace(trimSize+1) + currentLine[i]
						restSpace--
						continue
					}
					resu[len(resu)-1] += trimSpace(trimSize) + currentLine[i]
				}
			}
			currentLen = 0
			currentLine = []string{}
		}
	}
	if len(currentLine) > 0 {
		resu = append(resu, currentLine[0])
		for i := 1; i < len(currentLine); i++ {
			resu[len(resu)-1] += " " + currentLine[i]
		}
	}
	// for i := 0; i < len(resu); i++ {
	// 	println(resu[i])
	// }
	// println()
	return resu
}

func trimSpace(n int) string {
	resu := ""
	for i := 0; i < n; i++ {
		resu += " "
	}
	return resu
}
