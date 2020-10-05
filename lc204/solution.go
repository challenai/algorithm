package solution

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	var base, resu int
	resu = 1
	for i := 3; i <= n; i += 2 {
		base = 3
		for base < i {
			if i%base == 0 {
				break
			}
			base += 2
		}
		if base >= i {
			resu++
		}
	}
	return resu
}
