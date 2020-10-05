package solution

func rangeBitwiseAnd(m, n int) int {
	if m > n {
		return 0
	}
	resu := m
	for i := m + 1; i <= n; i++ {
		resu &= i
	}
	return resu
}
