package solution

func getRow(n int) []int {
	resu := []int{1}
	if n == 0 {
		return resu
	}
	resu = append(resu, 1)
	n--
	if n == 0 {
		return resu
	}
	for n > 0 {
		nextMid := len(resu) / 2
		for i := 1; i <= nextMid; i++ {
			resu[i] += resu[len(resu)-1-i]
		}
		resu = append(resu, 1)
		for i := len(resu) - 2; i > nextMid; i-- {
			resu[i] = resu[len(resu)-1-i]
		}
		n--
	}
	return resu
}
