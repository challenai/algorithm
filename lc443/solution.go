package solution

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	left := 1
	currentChar := chars[0]
	currentDup := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] == currentChar {
			currentDup++
		} else {
			// save dups
			if currentDup > 9 {
				temp := []byte{}
				for currentDup > 0 {
					temp = append(temp, byte('0'+currentDup%10))
					currentDup /= 10
				}
				for j := len(temp) - 1; j >= 0; j-- {
					chars[left] = temp[j]
					left++
				}
			} else if currentDup != 1 {
				chars[left] = byte('0' + currentDup)
				left++
			}
			chars[left] = chars[i]
			left++
			currentChar = chars[i]
			currentDup = 1
		}
	}

	if currentDup > 9 {
		temp := []byte{}
		for currentDup > 0 {
			temp = append(temp, byte('0'+currentDup%10))
			currentDup /= 10
		}
		for j := len(temp) - 1; j >= 0; j-- {
			chars[left] = temp[j]
			left++
		}
	} else if currentDup != 1 {
		chars[left] = byte('0' + currentDup)
		left++
	}

	return left
}
