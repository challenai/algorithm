package solution

func reverse(n int) int {
	resu := 0
	needReverseSymbol := false
	if n < 0 {
		needReverseSymbol = true
	}
	for n > 0 {
		resu *= 10
		resu += n % 10
		n /= 10
	}
	if needReverseSymbol {
		return -resu
	}
	return resu

}
