package solution

import "strconv"

func getPermutation(n, k int) string {
	resu := ""
	i := n - 1
	k %= factorial(n)
	opticals := make([]int, n)
	for i := 1; i <= n; i++ {
		opticals[i-1] = i
	}
	for k > 1 {
		current := 0
		for k >= factorial(i) {
			current++
			k -= factorial(i)
		}
		resu += strconv.Itoa(opticals[current])
		copy(opticals[current:], opticals[current+1:])
		opticals[len(opticals)-1] = 0
		opticals = opticals[:len(opticals)-1]
		i--
	}
	for i := 0; i < len(opticals); i++ {
		resu += strconv.Itoa(opticals[i])
	}
	return resu
}

func factorial(n int) int {
	resu := 1
	for i := 2; i <= n; i++ {
		resu *= i
	}
	return resu
}
