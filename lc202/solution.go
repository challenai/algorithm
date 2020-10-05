package solution

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	path := map[int]bool{
		n: false,
	}
	var temp int
	currentNum := n
	ok := false
	for !ok {
		temp = 0
		for currentNum > 0 {
			temp += (currentNum % 10) * (currentNum % 10)
			currentNum /= 10
		}
		if temp == 1 {
			return true
		}
		currentNum = temp
		_, ok = path[currentNum]
	}

	return false
}
