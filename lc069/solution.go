package solution

func sqrt(x int) int {
	if x < 1 {
		return 0
	}
	var resu int
	counter := 1
	for square(counter) <= x {
		counter *= 2
	}
	counter /= 2
	resu = counter
	if square(resu) == x {
		return resu
	}

	for counter > 0 {
		counter /= 2
		if square(resu+counter) <= x {
			resu += counter
		}
	}
	return resu
}

func index2(n int) int {
	resu := 1
	for i := 0; i < n; i++ {
		resu *= 2
	}
	return resu
}

func square(n int) int {
	return n * n
}
