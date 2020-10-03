package solution

func convertToTitle(num int) string {
	resu := ""
	const digitsCnt = 26
	for num > 0 {
		if num%digitsCnt == 0 {
			resu = "Z" + resu
			num /= digitsCnt
			num--
		} else {
			resu = string(byte('A'+num%digitsCnt-1)) + resu
			num /= digitsCnt
		}
	}
	return resu
}
