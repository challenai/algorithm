package solution

func Fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	resu := 1
	temp := 0
	for i := 2; i <= n; i++ {
		resu, temp = resu+temp, resu
	}

	return resu
}
